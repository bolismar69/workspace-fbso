package com.fbso.platform.admin.dto.request;

import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.NotNull;

import java.util.UUID;

/**
 * Request para criação de Produto/Serviço (F04-06).
 *
 * <p>RN18-01: Vinculado à BU ativa no seletor.
 * RN18-02: SKU opcional, se informado único por BU.</p>
 */
public record ProductCreateRequest(
        @NotNull(message = "ID da Unidade de Negócio é obrigatório")
        UUID businessUnitId,

        @NotBlank(message = "Nome do produto é obrigatório")
        String name,

        String sku,

        @NotBlank(message = "Tipo é obrigatório (PRODUCT ou SERVICE)")
        String type,

        String description
) {}
