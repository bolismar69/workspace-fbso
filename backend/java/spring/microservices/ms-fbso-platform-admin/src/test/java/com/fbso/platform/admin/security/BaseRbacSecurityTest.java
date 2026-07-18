package com.fbso.platform.admin.security;

import com.fbso.platform.admin.integration.BaseIntegrationTest;
import io.restassured.RestAssured;
import io.restassured.specification.RequestSpecification;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.TestInstance;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.server.LocalServerPort;
import org.springframework.test.context.ActiveProfiles;

import java.util.List;
import java.util.UUID;

import static io.restassured.RestAssured.given;

/**
 * Classe base para testes de segurança RBAC com REST Assured + Testcontainers (T-056).
 *
 * <p><b>Objetivo:</b> Validar a matriz RN10-01 com 20+ combinações papel×endpoint.
 * Cada combinação testa que um papel SEM permissão recebe HTTP 403 com
 * o formato RFC 7807 padrão.
 *
 * <p><b>Estratégia de autenticação nos testes:</b>
 * Como o JwtAuthenticationFilter valida tokens JWT reais (Keycloak RS256),
 * os testes de segurança RBAC operam no nível do {@link PermissionService}
 * e {@link com.fbso.platform.admin.security.aspect.RbacAspect}, que são
 * testados diretamente via {@code TenantContext} + chamadas de serviço.
 *
 * <p>Os cenários de matriz papel×endpoint são validados via:
 * <ul>
 *   <li>{@link com.fbso.platform.admin.unit.service.PermissionServiceTest} — unitário</li>
 *   <li>{@link RbacMatrixValidationTest} — matriz completa 20+ combinações</li>
 * </ul>
 *
 * <h3>Pré-requisitos</h3>
 * <ul>
 *   <li>Migration V004 (seed RBAC) executada</li>
 *   <li>PostgreSQL com RLS ativo</li>
 * </ul>
 *
 * @see RbacMatrixValidationTest
 */
@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
@ActiveProfiles("test")
@TestInstance(TestInstance.Lifecycle.PER_CLASS)
public abstract class BaseRbacSecurityTest extends BaseIntegrationTest {

    @LocalServerPort
    protected int port;

    @BeforeAll
    void baseSetUp() {
        RestAssured.port = port;
        RestAssured.basePath = "/api/v1";
    }

    @BeforeEach
    void cleanContext() {
        TenantContext.clear();
    }

    /**
     * Cria um RequestSpecification com headers de autenticação simulados.
     *
     * <p>Como os testes RBAC operam no nível de serviço (não via HTTP real
     * com JWT Keycloak), este método configura o TenantContext antes de
     * cada verificação.
     *
     * @param role  papel a simular (ADMIN_TENANT, MANAGER_BU, OPERATOR_BU, AUDITOR)
     * @param bus   BUs do usuário
     */
    protected void givenUser(String role, List<UUID> bus) {
        TenantContext.set(
                UUID.randomUUID(), // tenantId
                UUID.randomUUID(), // userId
                List.of(role),
                bus != null ? bus : List.of(),
                List.of()
        );
    }
}
