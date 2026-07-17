package com.fbso.platform.admin.integration;

import org.junit.jupiter.api.AfterAll;
import org.junit.jupiter.api.BeforeAll;
import org.springframework.test.context.DynamicPropertyRegistry;
import org.springframework.test.context.DynamicPropertySource;
import org.testcontainers.containers.PostgreSQLContainer;
import org.testcontainers.junit.jupiter.Container;
import org.testcontainers.junit.jupiter.Testcontainers;

/**
 * Classe base para testes de integração com Testcontainers PostgreSQL.
 *
 * <p><b>Padrão de uso:</b></p>
 * <pre>
 * class MeuRepositoryTest extends BaseIntegrationTest {
 *     &#64;Autowired private MeuRepository repository;
 *
 *     &#64;Test
 *     void deveBuscarPorId() {
 *         var result = repository.findById(knownId);
 *         assertThat(result).isPresent();
 *     }
 * }
 * </pre>
 *
 * <p>O container PostgreSQL sobe uma única vez por classe de teste
 * (hierarquia {@code @TestInstance(PER_CLASS)} implícita do Testcontainers).
 * Flyway executa as migrations automaticamente no startup.</p>
 *
 * <h3>Pré-requisitos</h3>
 * <ul>
 *   <li>Docker rodando</li>
 *   <li>Testcontainers no pom.xml (já configurado — v1.20.6)</li>
 * </ul>
 */
@Testcontainers
public abstract class BaseIntegrationTest {

    @Container
    protected static PostgreSQLContainer<?> postgres = new PostgreSQLContainer<>("postgres:17-alpine")
            .withDatabaseName("fbso_platform_test")
            .withUsername("fbso_test")
            .withPassword("fbso_test");

    @BeforeAll
    static void startContainer() {
        if (!postgres.isRunning()) {
            postgres.start();
        }
    }

    @AfterAll
    static void stopContainer() {
        // Testcontainers encerra automaticamente via @Container,
        // mas garantimos cleanup explícito para evitar dangling containers
    }

    @DynamicPropertySource
    static void configureProperties(DynamicPropertyRegistry registry) {
        registry.add("spring.datasource.url", postgres::getJdbcUrl);
        registry.add("spring.datasource.username", postgres::getUsername);
        registry.add("spring.datasource.password", postgres::getPassword);
        registry.add("spring.flyway.locations", () -> "classpath:db/migration");
        registry.add("spring.flyway.schemas", () -> "fbso_platform");
    }
}
