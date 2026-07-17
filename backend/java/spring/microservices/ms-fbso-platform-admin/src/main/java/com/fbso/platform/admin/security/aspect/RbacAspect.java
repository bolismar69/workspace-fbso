package com.fbso.platform.admin.security.aspect;

import com.fbso.platform.admin.exception.PermissionDeniedException;
import com.fbso.platform.admin.security.TenantContext;
import com.fbso.platform.admin.security.annotation.RequiresPermission;
import org.aspectj.lang.ProceedingJoinPoint;
import org.aspectj.lang.annotation.Around;
import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.reflect.MethodSignature;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Component;

import java.util.Set;

/**
 * Aspecto RBAC — intercepta métodos anotados com {@link RequiresPermission}
 * e verifica se o papel do usuário tem acesso ao recurso + ação.
 * <p>
 * Pipeline: JWT Filter → RbacAspect → Controller
 * <p>
 * Matriz de permissões (RN10-01):
 * <ul>
 *   <li>ADMIN_TENANT — acesso total (todos resources × actions)</li>
 *   <li>MANAGER_BU — edita BUs e Produtos da sua unidade</li>
 *   <li>OPERATOR_BU — apenas leitura de BUs e Produtos</li>
 *   <li>AUDITOR — apenas leitura de Auditoria</li>
 * </ul>
 * <p>
 * Na Fase 0, a matriz é simplificada. O carregamento completo do banco
 * (tabelas resource_action + role_resource) é implementado na Sprint 4 (M4 — RBAC).
 * <p>
 * // ponytail: matriz hardcoded até Sprint 4 — substituir por consulta ao banco com cache
 *
 * @see RequiresPermission
 * @see <a href="ARCHITECTURE.md#4.1">ARCHITECTURE.md §4.1</a>
 */
@Aspect
@Component
public class RbacAspect {

    private static final Logger log = LoggerFactory.getLogger(RbacAspect.class);

    /**
     * Matriz simplificada de permissões (Fase 0).
     * // ponytail: ceiling = Sprint 4 carrega do banco (RoleResource + ResourceAction)
     */
    private static final Set<String> ADMIN_ALL_ACCESS = Set.of("ADMIN_TENANT");
    private static final Set<String> MANAGER_EDIT_RESOURCES = Set.of("BUSINESS_UNIT", "PRODUCT_SERVICE");
    private static final Set<String> MANAGER_VIEW_RESOURCES = Set.of(
            "BUSINESS_UNIT", "PRODUCT_SERVICE",
            "TENANT", "PLAN", "SUBSCRIPTION", "DASHBOARD");
    private static final Set<String> MANAGER_EDIT_ACTIONS = Set.of("view", "create", "edit");
    private static final Set<String> MANAGER_VIEW_ACTIONS = Set.of("view");
    private static final Set<String> OPERATOR_RESOURCES = Set.of(
            "BUSINESS_UNIT", "PRODUCT_SERVICE",
            "TENANT", "PLAN", "SUBSCRIPTION", "DASHBOARD");
    private static final Set<String> OPERATOR_ACTIONS = Set.of("view");
    private static final Set<String> AUDITOR_RESOURCES = Set.of("AUDIT");
    private static final Set<String> AUDITOR_ACTIONS = Set.of("view");

    @Around("@annotation(requiresPermission)")
    public Object checkPermission(ProceedingJoinPoint joinPoint,
                                   RequiresPermission requiresPermission) throws Throwable {

        String resource = requiresPermission.resource();
        String action = requiresPermission.action();
        var roles = TenantContext.getRoles();

        if (roles.isEmpty()) {
            log.warn("RBAC: sem roles no contexto — acesso negado a {}:{}", resource, action);
            throw new PermissionDeniedException();
        }

        // Admin tem acesso total
        if (roles.stream().anyMatch(ADMIN_ALL_ACCESS::contains)) {
            log.debug("RBAC: admin acessa {}:{}", resource, action);
            return joinPoint.proceed();
        }

        // Verificar permissão por papel
        boolean granted = roles.stream().anyMatch(role -> switch (role) {
            case "MANAGER_BU" -> (MANAGER_EDIT_RESOURCES.contains(resource)
                               && MANAGER_EDIT_ACTIONS.contains(action))
                              || (MANAGER_VIEW_RESOURCES.contains(resource)
                               && MANAGER_VIEW_ACTIONS.contains(action));
            case "OPERATOR_BU" -> OPERATOR_RESOURCES.contains(resource)
                              && OPERATOR_ACTIONS.contains(action);
            case "AUDITOR"     -> AUDITOR_RESOURCES.contains(resource)
                              && AUDITOR_ACTIONS.contains(action);
            default -> false;
        });

        if (!granted) {
            log.warn("RBAC: acesso negado — role={}, resource={}, action={}",
                    roles, resource, action);
            throw new PermissionDeniedException();
        }

        log.debug("RBAC: acesso permitido — role={}, resource={}, action={}",
                roles, resource, action);
        return joinPoint.proceed();
    }
}
