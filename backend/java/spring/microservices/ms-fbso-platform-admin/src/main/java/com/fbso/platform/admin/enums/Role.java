package com.fbso.platform.admin.enums;

/**
 * Papéis RBAC conforme matriz RN10-01.
 * <p>
 * <pre>
 * ADMIN_TENANT  — acesso total ao tenant
 * MANAGER_BU    — gerencia sua(s) Unidade(s) de Negócio
 * OPERATOR_BU   — apenas leitura de BUs e Produtos
 * AUDITOR       — apenas leitura de Auditoria
 * </pre>
 */
public enum Role {
    ADMIN_TENANT,
    MANAGER_BU,
    OPERATOR_BU,
    AUDITOR
}
