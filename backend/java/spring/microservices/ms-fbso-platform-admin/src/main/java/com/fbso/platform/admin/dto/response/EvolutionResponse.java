package com.fbso.platform.admin.dto.response;

import com.fasterxml.jackson.annotation.JsonFormat;
import com.fasterxml.jackson.annotation.JsonInclude;

import java.time.LocalDate;
import java.util.List;

/**
 * Evolução temporal da base de tenants.
 * <p>
 * F01-01: Períodos suportados: 7d, 30d, 90d, mês_atual, ano_atual.
 * Padrão: mês atual (RN01-02).
 */
@JsonInclude(JsonInclude.Include.NON_NULL)
public record EvolutionResponse(
        String period,
        List<DataPoint> dataPoints
) {
    @JsonInclude(JsonInclude.Include.NON_NULL)
    public record DataPoint(
            @JsonFormat(pattern = "yyyy-MM-dd")
            LocalDate date,
            int count
    ) {}

    public static EvolutionResponse of(String period, List<DataPoint> dataPoints) {
        return new EvolutionResponse(period, dataPoints);
    }
}
