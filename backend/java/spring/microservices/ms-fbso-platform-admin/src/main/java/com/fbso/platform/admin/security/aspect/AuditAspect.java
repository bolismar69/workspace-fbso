package com.fbso.platform.admin.security.aspect;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.security.annotation.Auditable;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.annotation.Around;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.reflect.MethodSignature;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.core.task.TaskExecutor;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.stereotype.Component;

import java.lang.reflect.Parameter;
import java.time.OffsetDateTime;
import java.util.Map;
import java.util.UUID;

/**
 * Aspecto de auditoria — grava operações anotadas com {@link Auditable}
 * na tabela {@code audit_log} de forma ASSÍNCRONA.
 *
 * <p>ADR-L03: Auditoria assíncrona para não bloquear a operação principal.</p>
 *
 * <p>Pipeline: @Around → captura estado ANTES (main thread) → proceed() → captura estado DEPOIS
 * → TaskExecutor (async) → INSERT audit_log com previous_value/new_value (DT-021).</p>
 *
 * <p>IMPORTANTE (T-015.3.DT-002): tenantId/userId capturados na thread principal
 * ANTES do dispatch assíncrono (ThreadLocal é limpo no finally do Filter).</p>
 *
 * @see Auditable
 */
@Aspect
@Component
public class AuditAspect {

    private static final Logger log = LoggerFactory.getLogger(AuditAspect.class);
    private static final ObjectMapper objectMapper = new ObjectMapper();

    private final JdbcTemplate jdbc;
    private final TaskExecutor taskExecutor;

    public AuditAspect(JdbcTemplate jdbc, TaskExecutor taskExecutor) {
        this.jdbc = jdbc;
        this.taskExecutor = taskExecutor;
    }

    /**
     * Intercepta métodos anotados com @Auditable usando @Around (DT-021).
     * <p>
     * Captura estado da entidade ANTES da execução, executa o método,
     * captura estado DEPOIS, e despacha auditoria assíncrona com diff.
     */
    @Around("@annotation(auditable)")
    public Object audit(ProceedingJoinPoint joinPoint, Auditable auditable) throws Throwable {
        // ---- Capturar contexto na THREAD PRINCIPAL ----
        UUID tenantId;
        UUID userId;
        try {
            tenantId = TenantContext.getTenantId();
            userId = TenantContext.getUserId();
        } catch (Exception e) {
            log.warn("Audit: TenantContext não disponível — auditoria ignorada");
            return joinPoint.proceed();
        }

        String entityType = auditable.entityType();
        String action = auditable.action();
        String entityId = extractEntityId(joinPoint, auditable);

        // ---- DT-021: Capturar estado ANTES ----
        String previousValue = captureEntityState(entityType, entityId);

        // ---- Executar método original ----
        Object result;
        try {
            result = joinPoint.proceed();
        } catch (Throwable ex) {
            // Propagar exceção — não auditar operações que falharam
            throw ex;
        }

        // ---- DT-021: Capturar estado DEPOIS ----
        String newValue = captureEntityState(entityType, entityId);

        // ---- Despachar para thread assíncrona ----
        final UUID finalTenantId = tenantId;
        final UUID finalUserId = userId;
        final String finalEntityId = entityId;
        final String finalPrevious = previousValue;
        final String finalNew = newValue;

        taskExecutor.execute(() -> writeAuditLog(finalTenantId, finalUserId, entityType,
                action, finalEntityId, finalPrevious, finalNew, auditable));

        return result;
    }

    /**
     * Grava o registro de auditoria com previous_value e new_value (DT-021).
     */
    private void writeAuditLog(UUID tenantId, UUID userId, String entityType,
                                String action, String entityId,
                                String previousValue, String newValue,
                                Auditable auditable) {
        try {
            UUID parsedEntityId = parseEntityId(entityId);

            String sql = """
                INSERT INTO fbso_platform.audit_log
                    (timestamp, tenant_id, action, entity_type, entity_id, actor_id,
                     previous_value, new_value)
                VALUES (?, ?, ?, ?, ?, ?, ?::jsonb, ?::jsonb)
                """;

            jdbc.update(sql,
                    OffsetDateTime.now(),
                    tenantId,
                    action,
                    entityType,
                    parsedEntityId,
                    userId,
                    previousValue,
                    newValue);

            log.debug("Audit: {} {} — entity={}, tenant={}, user={}, hasPrevious={}, hasNew={}",
                    entityType, action, entityId, tenantId, userId,
                    previousValue != null, newValue != null);

        } catch (Exception e) {
            log.warn("Audit: falha ao gravar registro — {}: {}", auditable.entityType(), e.getMessage());
        }
    }

    /**
     * DT-021: Consulta o estado atual da entidade no banco e serializa como JSON.
     *
     * @return JSON string com os dados da entidade, ou null se não encontrada
     */
    private String captureEntityState(String entityType, String entityId) {
        if (entityId == null || entityId.isBlank()) {
            return null;
        }
        try {
            UUID id = parseEntityId(entityId);
            String tableName = resolveTableName(entityType);
            if (tableName == null) {
                return null;
            }

            String sql = "SELECT * FROM fbso_platform." + tableName + " WHERE id = ?";
            Map<String, Object> row = null;
            try {
                var rows = jdbc.queryForList(sql, id);
                if (!rows.isEmpty()) {
                    row = rows.get(0);
                }
            } catch (Exception e) {
                log.debug("Audit: não foi possível consultar estado de {}.{} — {}",
                        entityType, entityId, e.getMessage());
                return null;
            }

            if (row == null) {
                return null; // entidade não existe (ex: antes de CREATE)
            }

            // Remover campos binários/pesados
            row.remove("created_dt");
            row.remove("updated_dt");
            row.remove("created_by");
            row.remove("updated_by");
            row.remove("deleted_dt");
            row.remove("deleted_by");

            return objectMapper.writeValueAsString(row);
        } catch (Exception e) {
            log.debug("Audit: falha ao serializar estado de {} — {}", entityType, e.getMessage());
            return null;
        }
    }

    /**
     * Mapeia entityType para nome da tabela no banco.
     */
    private String resolveTableName(String entityType) {
        return switch (entityType) {
            case "TENANT" -> "tenant";
            case "PLAN" -> "plan";
            case "SUBSCRIPTION" -> "subscription";
            case "USER" -> "\"user\"";
            case "BUSINESS_UNIT" -> "business_unit";
            case "PRODUCT_SERVICE" -> "product_service";
            default -> null;
        };
    }

    private UUID parseEntityId(String entityId) {
        if (entityId == null || entityId.isBlank()) {
            return UUID.randomUUID();
        }
        try {
            return UUID.fromString(entityId);
        } catch (IllegalArgumentException e) {
            log.warn("Audit: entity_id '{}' não é UUID válido", entityId);
            return UUID.randomUUID();
        }
    }

    /**
     * Extrai o ID da entidade dos argumentos do método (T-015.9.DT-008).
     */
    private String extractEntityId(ProceedingJoinPoint joinPoint, Auditable auditable) {
        Object[] args = joinPoint.getArgs();
        if (args.length == 0) {
            return null;
        }

        // Estratégia 1: parâmetro nomeado via anotação
        String idParamName = auditable.idParamName();
        if (!idParamName.isBlank()) {
            return extractByParameterName(joinPoint, idParamName);
        }

        // Estratégia 2: primeiro argumento UUID ou String
        Object firstArg = args[0];
        if (firstArg instanceof UUID uuid) {
            return uuid.toString();
        }
        if (firstArg instanceof String str && !str.isBlank()) {
            return str;
        }

        // Estratégia 3: reflexão getId()
        try {
            var getIdMethod = firstArg.getClass().getMethod("getId");
            Object id = getIdMethod.invoke(firstArg);
            if (id instanceof UUID uuid) return uuid.toString();
            if (id != null) return id.toString();
        } catch (Exception e) {
            log.debug("Audit: reflection getId() falhou para {}", firstArg.getClass().getSimpleName());
        }

        return null;
    }

    private String extractByParameterName(ProceedingJoinPoint joinPoint, String paramName) {
        MethodSignature signature = (MethodSignature) joinPoint.getSignature();
        Parameter[] parameters = signature.getMethod().getParameters();
        Object[] args = joinPoint.getArgs();

        for (int i = 0; i < parameters.length; i++) {
            if (paramName.equals(parameters[i].getName())) {
                Object value = args[i];
                if (value instanceof UUID uuid) return uuid.toString();
                if (value != null) return value.toString();
            }
        }
        return null;
    }
}
