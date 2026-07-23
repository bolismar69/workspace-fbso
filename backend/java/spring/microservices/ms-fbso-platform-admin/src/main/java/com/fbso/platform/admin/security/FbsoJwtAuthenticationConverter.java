package com.fbso.platform.admin.security;

import com.fbso.platform.admin.utils.JwtUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.core.convert.converter.Converter;
import org.springframework.security.authentication.AbstractAuthenticationToken;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.oauth2.jwt.Jwt;
import org.springframework.security.oauth2.server.resource.authentication.JwtAuthenticationToken;

import java.util.List;
import java.util.UUID;

/**
 * Converter customizado que extrai claims do JWT durante a decodificação
 * do Resource Server (Spring Security), eliminando a dupla decodificação.
 * <p>
 * <b>Problema (DT-076/DT-102):</b> Antes deste converter, cada requisição
 * decodificava o JWT duas vezes:
 * <ol>
 *   <li>{@code BearerTokenAuthenticationFilter} (Spring Security) —
 *       valida assinatura e expiração</li>
 *   <li>{@code JwtAuthenticationFilter} (customizado) —
 *       extrai claims e popula {@link TenantContext}</li>
 * </ol>
 * <p>
 * <b>Solução:</b> Este converter é injetado no Resource Server via
 * {@code .jwtAuthenticationConverter()}. Durante a primeira (e única)
 * decodificação, ele extrai todas as claims customizadas
 * ({@code tenant_id, user_id, roles, business_unit_ids, modules})
 * e popula {@link TenantContext} e as authorities do SecurityContext.
 * <p>
 * O {@link JwtAuthenticationFilter} continua responsável por:
 * <ul>
 *   <li>Retornar 401 para tokens ausentes (antes do Resource Server)</li>
 *   <li>Limpar {@link TenantContext} no {@code finally}</li>
 * </ul>
 * Mas NÃO decodifica mais o JWT — usa o {@link Jwt} já validado
 * pelo Resource Server.
 *
 * @see <a href="ARCHITECTURE.md#4">ARCHITECTURE.md §4 — Pipeline de Segurança</a>
 */
public class FbsoJwtAuthenticationConverter
        implements Converter<Jwt, AbstractAuthenticationToken> {

    private static final Logger log = LoggerFactory.getLogger(
            FbsoJwtAuthenticationConverter.class);

    @Override
    public AbstractAuthenticationToken convert(Jwt jwt) {
        // Extrair claims customizadas do JWT (já validado pelo Resource Server)
        UUID tenantId = JwtUtils.getTenantId(jwt);
        UUID userId = JwtUtils.getUserId(jwt);
        List<String> roles = JwtUtils.getRoles(jwt);
        List<UUID> businessUnitIds = JwtUtils.getBusinessUnitIds(jwt);
        List<String> modules = JwtUtils.getModules(jwt);

        // Popular TenantContext (ThreadLocal) para isolamento multi-tenant
        TenantContext.set(tenantId, userId, roles, businessUnitIds, modules);

        // Construir authorities a partir das roles do JWT
        var authorities = roles.stream()
                .map(r -> new SimpleGrantedAuthority("ROLE_" + r))
                .toList();

        log.debug("JWT convertido: tenant={}, user={}, roles={}, bus={}, modules={}",
                tenantId, userId, roles.size(), businessUnitIds.size(), modules);

        return new JwtAuthenticationToken(jwt, authorities);
    }
}
