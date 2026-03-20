package com.fbso.geolocalidade.entity;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.Table;

@Entity
@Table(name = "distrito")
public class Distrito {

  @Id
  @Column(name = "id", length = 9, nullable = false)
  private String id;

  @Column(name = "codigo", length = 2, nullable = false)
  private String codigo;

  @Column(name = "nome", length = 100, nullable = false)
  private String nome;

  @ManyToOne(fetch = FetchType.LAZY, optional = false)
  @JoinColumn(name = "municipio_id", referencedColumnName = "id", nullable = false)
  private Municipio municipio;

  protected Distrito() {}

  public Distrito(String id, String codigo, String nome, Municipio municipio) {
    this.id = id;
    this.codigo = codigo;
    this.nome = nome;
    this.municipio = municipio;
  }

  public String getId() {
    return id;
  }

  public String getCodigo() {
    return codigo;
  }

  public String getNome() {
    return nome;
  }

  public Municipio getMunicipio() {
    return municipio;
  }
}
