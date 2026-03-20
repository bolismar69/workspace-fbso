package com.fbso.geolocalidade.entity;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.Table;

@Entity
@Table(name = "municipio")
public class Municipio {

  @Id
  @Column(name = "id", length = 7, nullable = false)
  private String id;

  @Column(name = "codigo", length = 5, nullable = false)
  private String codigo;

  @Column(name = "nome", length = 100, nullable = false)
  private String nome;

  @ManyToOne(fetch = FetchType.LAZY, optional = false)
  @JoinColumn(name = "regiao_imediata_id", referencedColumnName = "id", nullable = false)
  private RegiaoImediata regiaoImediata;

  protected Municipio() {}

  public Municipio(String id, String codigo, String nome, RegiaoImediata regiaoImediata) {
    this.id = id;
    this.codigo = codigo;
    this.nome = nome;
    this.regiaoImediata = regiaoImediata;
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

  public RegiaoImediata getRegiaoImediata() {
    return regiaoImediata;
  }

  // ---- Legacy getters (mantidos para não quebrar o endpoint existente) ----
  public String getCodigoIbge7() {
    return id;
  }

  public String getNomeMunicipio() {
    return nome;
  }

  public String getUfSigla() {
    if (regiaoImediata == null) {
      return null;
    }
    RegiaoIntermediaria ri = regiaoImediata.getRegiaoIntermediaria();
    if (ri == null) {
      return null;
    }
    Uf uf = ri.getUf();
    return uf == null ? null : uf.getSigla();
  }
}
