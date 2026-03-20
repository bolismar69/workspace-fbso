package br.com.fbso.geolocalidade.entity;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.Id;
import jakarta.persistence.Table;

@Entity
@Table(name = "uf")
public class Uf {

  @Id
  @Column(name = "id", length = 2, nullable = false)
  private String id;

  @Column(name = "sigla", length = 5, nullable = false)
  private String sigla;

  @Column(name = "nome", length = 50, nullable = false)
  private String nome;

  protected Uf() {}

  public Uf(String id, String sigla, String nome) {
    this.id = id;
    this.sigla = sigla;
    this.nome = nome;
  }

  public String getId() {
    return id;
  }

  public String getSigla() {
    return sigla;
  }

  public String getNome() {
    return nome;
  }
}
