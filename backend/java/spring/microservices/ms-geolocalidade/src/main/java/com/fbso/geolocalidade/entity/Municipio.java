package com.fbso.geolocalidade.entity;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

@Entity
@Table(name = "municipios")
public class Municipio {

  @Id
  @Column(name = "codigo_ibge7", length = 7, nullable = false)
  private String codigoIbge7;

  @Column(name = "nome_municipio", nullable = false)
  private String nomeMunicipio;

  @Column(name = "uf_sigla", length = 2, nullable = false)
  private String ufSigla;

  protected Municipio() {}

  public Municipio(String codigoIbge7, String nomeMunicipio, String ufSigla) {
    this.codigoIbge7 = codigoIbge7;
    this.nomeMunicipio = nomeMunicipio;
    this.ufSigla = ufSigla;
  }

  public String getCodigoIbge7() {
    return codigoIbge7;
  }

  public String getNomeMunicipio() {
    return nomeMunicipio;
  }

  public String getUfSigla() {
    return ufSigla;
  }
}
