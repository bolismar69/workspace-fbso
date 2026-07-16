package com.fbso.platform.admin.security.aspect;

import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.security.annotation.Auditable;
import org.aspectj.lang.JoinPoint;
import org.aspectj.lang.annotation.AfterReturning;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.reflect.MethodSignature;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.scheduling.annotation.Async;
import org.springframework.stereotype.Component;

import java.time.OffsetDateTime;
import java.util.UUID;

/**
 * Aspecto de auditoria — grava operações anotadas com {@link Auditable}
 * na tabela {@code audit_log} de forma ASSÍNCRONA.
 * <p>
 * ADR-L03: Auditoria assíncrona para não bloquear a operação principal.
 * Trade-off: perda de registros em crash (aceitável para Fase 0).
 * <p>
 * Pipeline: Service → AuditAspect (async) → audit_log
 *
 * @see Auditable
 * @see <a href="ARCHITECTURE.md#4.2">ARCHITECTURE.md §4.2</a>
 */
@Aspect
@Component
public class AuditAspect {

    private static final Logger log = LoggerFactory.getLogger(AuditAspect.class);

    private final JdbcTemplate jdbc;

    public AuditAspect(JdbcTemplate jdbc) {
        this.jdbc = jdbc;
    }

    /**
     * Intercepta métodos anotados com @Auditable APÓS execução bem-sucedida.
     * A gravação é assíncrona (@Async) — não bloqueia a thread principal.
     */
    @AfterReturning("@annotation(auditable)")
    @Async
    public void audit(JoinPoint joinPoint, Auditable auditable) {
        try {
            UUID tenantId = TenantContext.getTenantId();
            UUID userId = TenantContext.getUserId();
            String entityType = auditable.entityType();
            String action = auditable.action();

            // Extrair ID da entidade dos argumentos do método (primeiro argumento String/UUID)
            String entityId = extractEntityId(joinPoint);

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
                    entityId != null ? UUID.fromString(entityId) : UUID.randomUUID(),
                    userId);

            log.debug("Audit: {} {} — tenant={}, user={}",
                    entityType, action, tenantId, userId);

        } catch (Exception e) {
            // ADR-L03: falha na auditoria NÃO interrompe a operação principal
            log.warn("Audit: falha ao gravar registro — {}: {}", auditable.entityType(), e.getMessage());
        }
    }

    /**
     * Extrai o ID da entidade dos argumentos do método.
     * Procura o primeiro argumento do tipo String ou UUID.
     */
    private String extractEntityId(JoinPoint joinPoint) {
        Object[] args = joinPoint.getArgs();
        if (args.length > 0 && args[0] != null) {
            return args[0].toString();
        }
        return null;
    }
}
