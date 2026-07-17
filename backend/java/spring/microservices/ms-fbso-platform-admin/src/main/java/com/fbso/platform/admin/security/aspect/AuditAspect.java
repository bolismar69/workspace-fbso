package com.fbso.platform.admin.security.aspect;

import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.security.annotation.Auditable;
import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.annotation.AfterReturning;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.reflect.MethodSignature;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.core.task.TaskExecutor;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Component;

import java.lang.reflect.Parameter;
import java.time.OffsetDateTime;
import java.util.UUID;

/**
 * Aspecto de auditoria — grava operações anotadas com {@link Auditable}
 * na tabela {@code audit_log} de forma ASSÍNCRONA.
 * <p>
 * ADR-L03: Auditoria assíncrona para não bloquear a operação principal.
 * Trade-off: perda de registros em crash (aceitável para Fase 0).
 * <p>
 * Pipeline: Service → AuditAspect (captura main thread) → TaskExecutor (async) → audit_log
 * <p>
 * IMPORTANTE (T-015.3.DT-002): Os valores de tenantId/userId são capturados
 * na thread principal ANTES do dispatch assíncrono, pois o
 * {@link TenantContext} é um ThreadLocal que é limpo no finally do
 * {@code JwtAuthenticationFilter}.
 *
 * @see Auditable
 * @see <a href="ARCHITECTURE.md#4.2">ARCHITECTURE.md §4.2</a>
 */
@Aspect
@Component
public class AuditAspect {

    private static final Logger log = LoggerFactory.getLogger(AuditAspect.class);

    private final JdbcTemplate jdbc;
    private final TaskExecutor taskExecutor;

    public AuditAspect(JdbcTemplate jdbc, TaskExecutor taskExecutor) {
        this.jdbc = jdbc;
        this.taskExecutor = taskExecutor;
    }

    /**
     * Intercepta métodos anotados com @Auditable APÓS execução bem-sucedida.
     * <p>
     * Captura tenantId/userId/entityId na thread principal e despacha
     * a gravação para um executor assíncrono (não bloqueia a requisição).
     */
    @AfterReturning("@annotation(auditable)")
    public void audit(JoinPoint joinPoint, Auditable auditable) {
        // ---- Capturar contexto na THREAD PRINCIPAL (T-015.3.DT-002) ----
        // O TenantContext é um ThreadLocal que será limpo no finally do
        // JwtAuthenticationFilter. Precisamos capturar os valores aqui,
        // ANTES que o TaskExecutor execute em outra thread.
        UUID tenantId;
        UUID userId;
        try {
            tenantId = TenantContext.getTenantId();
            userId = TenantContext.getUserId();
        } catch (Exception e) {
            log.warn("Audit: TenantContext não disponível — auditoria ignorada");
            return;
        }

        String entityType = auditable.entityType();
        String action = auditable.action();
        String entityId = extractEntityId(joinPoint, auditable);

        // ---- Despachar para thread assíncrona com valores capturados ----
        taskExecutor.execute(() -> writeAuditLog(tenantId, userId, entityType,
                action, entityId, auditable));
    }

    /**
     * Grava o registro de auditoria no banco (executado em thread separada).
     */
    private void writeAuditLog(UUID tenantId, UUID userId, String entityType,
                                String action, String entityId, Auditable auditable) {
        try {
            UUID parsedEntityId = parseEntityId(entityId);

            String sql = """
                INSERT INTO fbso_platform.audit_log
                    (timestamp, tenant_id, action, entity_type, entity_id, actor_id)
                VALUES (?, ?, ?, ?, ?, ?)
                """;

            jdbc.update(sql,
                    OffsetDateTime.now(),
                    tenantId,
                    action,
                    entityType,
                    parsedEntityId,
                    userId);

            log.debug("Audit: {} {} — entity={}, tenant={}, user={}",
                    entityType, action, entityId, tenantId, userId);

        } catch (Exception e) {
            // ADR-L03: falha na auditoria NÃO interrompe a operação principal
            log.warn("Audit: falha ao gravar registro — {}: {}", auditable.entityType(), e.getMessage());
        }
    }

    /**
     * Converte o entityId bruto em UUID válido, ou gera fallback.
     */
    private UUID parseEntityId(String entityId) {
        if (entityId == null || entityId.isBlank()) {
            log.warn("Audit: entity_id não encontrado — usando fallback UUID");
            return UUID.randomUUID();
        }
        try {
            return UUID.fromString(entityId);
        } catch (IllegalArgumentException e) {
            log.warn("Audit: entity_id '{}' não é UUID válido — usando fallback", entityId);
            return UUID.randomUUID();
        }
    }

    /**
     * Extrai o ID da entidade dos argumentos do método.
     * <p>
     * Estratégia (T-015.9.DT-008):
     * <ol>
     *   <li>Se {@code @Auditable(idParamName)} foi especificado, busca o parâmetro pelo nome</li>
     *   <li>Caso contrário, tenta o primeiro argumento como UUID/String</li>
     *   <li>Se o argumento for um objeto com método {@code getId()}, usa reflection</li>
     * </ol>
     */
    private String extractEntityId(JoinPoint joinPoint, Auditable auditable) {
        Object[] args = joinPoint.getArgs();
        if (args.length == 0) {
            return null;
        }

        // Estratégia 1: parâmetro nomeado via anotação
        String idParamName = auditable.idParamName();
        if (!idParamName.isBlank()) {
            return extractByParameterName(joinPoint, idParamName);
        }

        // Estratégia 2: primeiro argumento que seja UUID ou String
        Object firstArg = args[0];
        if (firstArg instanceof UUID uuid) {
            return uuid.toString();
        }
        if (firstArg instanceof String str && !str.isBlank()) {
            return str;
        }

        // Estratégia 3: reflexão — tenta getId() no primeiro argumento
        try {
            var getIdMethod = firstArg.getClass().getMethod("getId");
            Object id = getIdMethod.invoke(firstArg);
            if (id instanceof UUID uuid) {
                return uuid.toString();
            }
            if (id != null) {
                return id.toString();
            }
        } catch (Exception e) {
            log.debug("Audit: não foi possível extrair ID via reflection de {}",
                    firstArg.getClass().getSimpleName());
        }

        return null;
    }

    /**
     * Extrai o valor de um parâmetro pelo nome (usa reflection no bytecode).
     */
    private String extractByParameterName(JoinPoint joinPoint, String paramName) {
        MethodSignature signature = (MethodSignature) joinPoint.getSignature();
        Parameter[] parameters = signature.getMethod().getParameters();
        Object[] args = joinPoint.getArgs();

        for (int i = 0; i < parameters.length; i++) {
            if (paramName.equals(parameters[i].getName())) {
                Object value = args[i];
                if (value instanceof UUID uuid) {
                    return uuid.toString();
                }
                if (value != null) {
                    return value.toString();
                }
            }
        }
        log.debug("Audit: parâmetro '{}' não encontrado no método", paramName);
        return null;
    }
}
