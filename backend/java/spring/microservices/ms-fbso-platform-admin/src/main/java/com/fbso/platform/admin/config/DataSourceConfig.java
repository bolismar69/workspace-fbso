package com.fbso.platform.admin.config;

import org.springframework.beans.BeansException;
import org.springframework.beans.factory.config.BeanPostProcessor;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import javax.sql.DataSource;

/**
 * Configuração do DataSource com suporte a Multi-Tenant via PostgreSQL RLS.
 * <p>
 * O {@link BeanPostProcessor} abaixo intercepta o DataSource auto-configurado
 * pelo Spring Boot (HikariCP) e o envolve com {@link TenantAwareDataSource}.
 * <p>
 * O {@link TenantAwareDataSource} configura {@code app.current_tenant_id}
 * na sessão PostgreSQL em cada {@code getConnection()}, permitindo que as
 * políticas de Row-Level Security (ADR-L07) funcionem corretamente com
 * connection pooling.
 *
 * @see TenantAwareDataSource
 * @see <a href="ARCHITECTURE.md#4.3">ARCHITECTURE.md §4.3</a>
 */
@Configuration
public class DataSourceConfig {

    /**
     * Post-processor que envolve o DataSource auto-configurado com
     * {@link TenantAwareDataSource} após sua inicialização pelo Spring Boot.
     * <p>
     * {@code static} → não depende de nenhum outro bean, evitando
     * dependência circular com o auto-configure do HikariCP.
     */
    @Bean
    static BeanPostProcessor tenantDataSourcePostProcessor() {
        return new BeanPostProcessor() {
            @Override
            public Object postProcessAfterInitialization(Object bean, String beanName)
                    throws BeansException {
                if (bean instanceof DataSource
                        && !(bean instanceof TenantAwareDataSource)) {
                    return new TenantAwareDataSource((DataSource) bean);
                }
                return bean;
            }
        };
    }
}
