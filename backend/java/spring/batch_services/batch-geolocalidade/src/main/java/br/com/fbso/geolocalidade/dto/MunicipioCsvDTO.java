package br.com.fbso.geolocalidade.dto;

public record MunicipioCsvDTO(
    String ufId,
    String ufNome,
    String regiaoInterId,
    String regiaoInterNome,
    String regiaoImedId,
    String regiaoImedNome,
    String municipioCodCurto,
    String municipioIdCompleto,
    String municipioNome
) {}