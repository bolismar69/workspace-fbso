package com.fbso.platform.admin.security.annotation;

import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

/**
 * Anotação de controle de acesso RBAC.
 * <p>
 * Aplica-se a métodos de Controller. O {@code RbacAspect} intercepta
 * métodos anotados e verifica se o papel do usuário (do JWT) tem
 * permissão para o recurso e ação especificados.
 * <p>
 * Exemplo de uso:
 * <pre>
 * @RequiresPermission(resource = "PRODUCT_SERVICE", action = "edit")
 * public ResponseEntity&lt;ProductResponse&gt; update(...) { }
 * </pre>
 *
 * @see com.fbso.platform.admin.security.aspect.RbacAspect
 * @see <a href="ARCHITECTURE.md#4.1">ARCHITECTURE.md §4.1</a>
 */
@Target(ElementType.METHOD)
@Retention(RetentionPolicy.RUNTIME)
public @interface RequiresPermission {

    /**
     * Nome do recurso protegido.
     * <p>
     * Exemplos: {@code DASHBOARD}, {@code TENANT}, {@code PLAN},
     * {@code PRODUCT_SERVICE}, {@code BUSINESS_UNIT}, {@code USER},
     * {@code PERMISSION}, {@code AUDIT}.
     */
    String resource();

    /**
     * Ação sobre o recurso.
     * <p>
     * Valores típicos: {@code view}, {@code create}, {@code edit}, {@code delete}.
     */
    String action();
}
