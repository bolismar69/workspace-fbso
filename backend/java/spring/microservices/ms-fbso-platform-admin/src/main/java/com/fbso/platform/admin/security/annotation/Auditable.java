package com.fbso.platform.admin.security.annotation;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

/**
 * Anotação de auditoria — registra automaticamente operações em {@code audit_log}.
 * <p>
 * Aplica-se a métodos de Service. O {@code AuditAspect} intercepta
 * métodos anotados e grava um registro em {@code audit_log} de forma
 * ASSÍNCRONA (não bloqueia a thread principal).
 * <p>
 * Exemplo de uso:
 * <pre>
 * @Auditable(entityType = "TENANT", action = "SUSPENDED")
 * public TenantResponse suspend(String tenantId, String reason) { }
 * </pre>
 *
 * @see com.fbso.platform.admin.security.aspect.AuditAspect
 * @see <a href="ARCHITECTURE.md#4.2">ARCHITECTURE.md §4.2</a>
 */
@Target(ElementType.METHOD)
@Retention(RetentionPolicy.RUNTIME)
public @interface Auditable {

    /**
     * Tipo da entidade auditada.
     * <p>
     * Exemplos: {@code TENANT}, {@code USER}, {@code PLAN},
     * {@code SUBSCRIPTION}, {@code BUSINESS_UNIT}, {@code PRODUCT_SERVICE}.
     */
    String entityType();

    /**
     * Ação executada.
     * <p>
     * Exemplos: {@code CREATED}, {@code UPDATED}, {@code SUSPENDED},
     * {@code REACTIVATED}, {@code PLAN_CHANGED}, {@code DEACTIVATED}.
     */
    String action();

    /**
     * Nome do parâmetro do método que contém o ID da entidade a auditar.
     * <p>
     * Se não informado, o aspecto tenta extrair o ID do primeiro argumento.
     * Exemplo: para {@code suspend(UUID tenantId, String reason)},
     * use {@code idParamName = "tenantId"}.
     */
    String idParamName() default "";
}
