package com.fbso.platform.admin;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.scheduling.annotation.EnableAsync;

/**
 * Ponto de entrada do microsserviço ms-fbso-platform-admin.
 * <p>
 * Backend do Core Administrativo da FBSO Platform — API REST com Spring Boot.
 * <p>
 * Stack: Java 25 + Spring Boot + PostgreSQL + Keycloak JWT.
 */
@SpringBootApplication
@EnableAsync
public class FbsoPlatformAdminApplication {

    public static void main(String[] args) {
        SpringApplication.run(FbsoPlatformAdminApplication.class, args);
    }
}
