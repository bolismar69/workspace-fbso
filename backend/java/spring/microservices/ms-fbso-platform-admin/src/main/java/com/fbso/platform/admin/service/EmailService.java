package com.fbso.platform.admin.service;

/**
 * Interface de envio de email (F02-01, T-028).
 *
 * <p>Implementação real usa {@code JavaMailSender} (spring-boot-starter-mail).
 * Para dev/testes, usar mock com Mailhog/GreenMail.</p>
 */
public interface EmailService {

    /**
     * Envia email de convite de ativação para o tenant.
     *
     * @param tenantId      ID do tenant
     * @param toEmail       email de destino
     * @param tenantName    nome do tenant (razão social)
     */
    void sendInvite(String tenantId, String toEmail, String tenantName);
}
