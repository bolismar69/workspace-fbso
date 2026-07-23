package com.fbso.platform.admin.dto.request;

import java.util.UUID;

/**
 * Request para atualização de Unidade de Negócio (F04-05).
 *
 * <p>RN17-01: CNPJ NÃO pode ser alterado após cadastro — campo não incluso.
 * RN17-02: parentId deve referir BU ativa.</p>
 */
public record BusinessUnitUpdateRequest(
        String corporateName,
        String taxRegime,
        UUID parentId,
        String street,
        String number,
        String complement,
        String neighborhood,
        String city,
        String state,
        String zipCode
) {}
