package com.fbso.geolocalidade.entity;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.Table;

@Entity
@Table(name = "regiao_imediata")
public class RegiaoImediata {

  @Id
  @Column(name = "id", length = 6, nullable = false)
  private String id;

  @Column(name = "nome", length = 100, nullable = false)
  private String nome;

  @ManyToOne(fetch = FetchType.LAZY, optional = false)
  @JoinColumn(name = "regiao_intermediaria_id", referencedColumnName = "id", nullable = false)
  private RegiaoIntermediaria regiaoIntermediaria;

  protected RegiaoImediata() {}

  public RegiaoImediata(String id, String nome, RegiaoIntermediaria regiaoIntermediaria) {
    this.id = id;
    this.nome = nome;
    this.regiaoIntermediaria = regiaoIntermediaria;
  }

  public String getId() {
    return id;
  }

  public String getNome() {
    return nome;
  }

  public RegiaoIntermediaria getRegiaoIntermediaria() {
    return regiaoIntermediaria;
  }
}
