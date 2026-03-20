package br.com.fbso.geolocalidade.dto;

public record DistritoCsvDTO(
    String ufId, String ufNome, String regInterId, String regInterNome,
    String regImedId, String regImedNome, String munCod, String municipioId,
    String munNome, String distritoCodCurto, String distritoIdCompleto, String distritoNome
) {}