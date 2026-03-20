package com.fbso.geolocalidade.config;

import jakarta.servlet.Servlet;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.autoconfigure.condition.ConditionalOnProperty;
import org.springframework.boot.web.servlet.ServletRegistrationBean;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
@ConditionalOnProperty(prefix = "spring.h2.console", name = "enabled", havingValue = "true")
public class H2ConsoleConfig {

  @Bean
  public ServletRegistrationBean<Servlet> h2ConsoleServlet(
      @Value("${spring.h2.console.path:/h2-console}") String consolePath) {

    String normalizedPath = consolePath;
    if (normalizedPath == null || normalizedPath.isBlank()) {
      normalizedPath = "/h2-console";
    }
    if (!normalizedPath.startsWith("/")) {
      normalizedPath = "/" + normalizedPath;
    }

    String mapping = normalizedPath.endsWith("/*") ? normalizedPath : normalizedPath + "/*";

    Servlet servlet = instantiateH2Servlet();
    ServletRegistrationBean<Servlet> bean = new ServletRegistrationBean<>(servlet, mapping);

    // Security defaults: do NOT allow remote access unless explicitly changed.
    bean.addInitParameter("webAllowOthers", "false");
    bean.addInitParameter("trace", "false");

    return bean;
  }

  private static Servlet instantiateH2Servlet() {
    try {
      Class<?> servletClass = Class.forName("org.h2.server.web.JakartaWebServlet");
      return (Servlet) servletClass.getDeclaredConstructor().newInstance();
    } catch (Exception exception) {
      throw new IllegalStateException(
          "H2 Console está habilitado, mas o servlet 'org.h2.server.web.JakartaWebServlet' não foi encontrado. "
              + "Verifique se a dependência 'com.h2database:h2' está no classpath em runtime.",
          exception);
    }
  }
}
