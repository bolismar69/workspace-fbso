package com.fbso.platform.admin.dto.request;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Pattern;

import java.util.UUID;

/**
 * Request para criação de Unidade de Negócio (F04-05).
 *
 * <p>RN17-01: CNPJ validado via {@code CnpjValidator}. Único entre ativos do tenant.
 * RN17-03: Primeira BU do onboarding é Matriz (isMatrix=true, parentId=null).</p>
 */
public record BusinessUnitCreateRequest(
        @NotBlank(message = "CNPJ é obrigatório")
        @Pattern(regexp = "^[A-Z0-9]{2}\\.[A-Z0-9]{3}\\.[A-Z0-9]{3}/[A-Z0-9]{4}-[A-Z0-9]{2}$",
                 message = "CNPJ em formato inválido. Use: XX.XXX.XXX/XXXX-XX")
        String cnpj,

        @NotBlank(message = "Razão social é obrigatória")
        String corporateName,

        @NotBlank(message = "Regime tributário é obrigatório")
        String taxRegime,

        UUID parentId,

        // Endereço (opcional no create)
        String street,
        String number,
        String complement,
        String neighborhood,
        String city,
        String state,
        String zipCode,

        boolean isMatrix
) {}
