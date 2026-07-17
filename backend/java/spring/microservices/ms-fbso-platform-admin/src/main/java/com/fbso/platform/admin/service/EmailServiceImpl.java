package com.fbso.platform.admin.service;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.mail.SimpleMailMessage;
import org.springframework.mail.javamail.JavaMailSender;
import org.springframework.stereotype.Service;

/**
 * Implementação de {@link EmailService} com JavaMailSender (T-028).
 *
 * <p>Em ambiente dev, o SMTP aponta para Mailhog (localhost:1025).
 * Em staging/prod, usar SMTP real configurado em application-{profile}.yml.</p>
 */
@Service
public class EmailServiceImpl implements EmailService {

    private static final Logger log = LoggerFactory.getLogger(EmailServiceImpl.class);

    private final JavaMailSender mailSender;

    public EmailServiceImpl(JavaMailSender mailSender) {
        this.mailSender = mailSender;
    }

    @Override
    public void sendInvite(String tenantId, String toEmail, String tenantName) {
        // Em dev, se o SMTP não estiver configurado, apenas loga
        if (toEmail == null || toEmail.isBlank()) {
            log.warn("Email de convite não enviado — destinatário vazio para tenant={}", tenantId);
            return;
        }

        try {
            SimpleMailMessage message = new SimpleMailMessage();
            message.setTo(toEmail);
            message.setSubject("Bem-vindo à FBSO Platform — Ative sua conta");
            message.setText(String.format("""
                    Olá, %s!

                    Sua conta na FBSO Platform foi criada com sucesso.
                    Acesse o link abaixo para iniciar o onboarding:

                    https://platform.fbso.org/onboarding/%s

                    Este link expira em 7 dias.

                    Atenciosamente,
                    Equipe FBSO Platform
                    """, tenantName, tenantId));

            mailSender.send(message);
            log.info("Email de convite enviado: tenant={}, to={}", tenantId,
                    maskEmail(toEmail));
        } catch (Exception e) {
            log.error("Falha ao enviar email de convite: tenant={}, to={}", tenantId,
                    maskEmail(toEmail), e);
        }
    }

    /**
     * Mascara email para logs: exibe apenas primeiros 2 caracteres da parte local.
     * Ex: "usuario@exemplo.com" → "us***@exemplo.com"
     */
    static String maskEmail(String email) {
        if (email == null || !email.contains("@")) return "***";
        return email.replaceAll("(?<=.{2}).(?=.*@)", "*");
    }
}
