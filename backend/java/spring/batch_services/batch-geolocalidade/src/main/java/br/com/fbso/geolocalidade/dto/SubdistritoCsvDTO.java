package br.com.fbso.geolocalidade.dto;

public record SubdistritoCsvDTO(
    String ufId, String ufNome, String regInterId, String regInterNome,
    String regImedId, String regImedNome, String munCod, String municipioId,
    String munNome, String distCod, String distritoId, String distNome,
    String subdistritoCodCurto, String subdistritoIdCompleto, String subdistritoNome
) {}