package com.fbso.platform.admin.dto.request;

/**
 * Request para atualização de Produto/Serviço (F04-06).
 */
public record ProductUpdateRequest(
        String name,
        String sku,
        String type,
        String description
) {}
