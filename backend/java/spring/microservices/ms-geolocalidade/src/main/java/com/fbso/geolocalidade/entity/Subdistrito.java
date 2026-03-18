package com.fbso.geolocalidade.entity;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

@Entity
@Table(name = "subdistritos")
public class Subdistrito {

  @Id
  @Column(name = "codigo_subdistrito11", length = 11, nullable = false)
  private String codigoSubdistrito11;

  @Column(name = "nome_subdistrito", nullable = false)
  private String nomeSubdistrito;

  protected Subdistrito() {}

  public Subdistrito(String codigoSubdistrito11, String nomeSubdistrito) {
    this.codigoSubdistrito11 = codigoSubdistrito11;
    this.nomeSubdistrito = nomeSubdistrito;
  }

  public String getCodigoSubdistrito11() {
    return codigoSubdistrito11;
  }

  public String getNomeSubdistrito() {
    return nomeSubdistrito;
  }
}
