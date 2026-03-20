package br.com.fbso.geolocalidade.entity;

import jakarta.persistence.CascadeType;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.Table;

@Entity
@Table(name = "regiao_intermediaria")
public class RegiaoIntermediaria {

  @Id
  @Column(name = "id", length = 4, nullable = false)
  private String id;

  @Column(name = "nome", length = 100, nullable = false)
  private String nome;

  @ManyToOne(fetch = FetchType.LAZY, optional = false, cascade = CascadeType.MERGE)
  @JoinColumn(name = "uf_id", referencedColumnName = "id", nullable = false)
  private Uf uf;

  protected RegiaoIntermediaria() {}

  public RegiaoIntermediaria(String id, String nome, Uf uf) {
    this.id = id;
    this.nome = nome;
    this.uf = uf;
  }

  public void setId(String id) {
    this.id = id;
  }

  public void setNome(String nome) {
    this.nome = nome;
  }

  public void setUf(Uf uf) {
    this.uf = uf;
  }

  public String getId() {
    return id;
  }

  public String getNome() {
    return nome;
  }

  public Uf getUf() {
    return uf;
  }
}
