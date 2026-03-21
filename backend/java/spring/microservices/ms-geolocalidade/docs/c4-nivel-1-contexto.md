# C4 — Nível 1 (Contexto) — ms-geolocalidade

## Objetivo

Descrever o **contexto** do microserviço `ms-geolocalidade`: quem o consome, quais sistemas externos ele integra e quais responsabilidades ele possui.

## Responsabilidade do sistema

O `ms-geolocalidade` é um microserviço REST que:

- Consulta **geolocalização por CEP** e **vizinhança em raio (km)** via **AwesomeAPI**.
- Enriquece respostas com dados oficiais do **DTB/IBGE** persistidos em **PostgreSQL**.
- Aplica cache em chamadas de geocodificação por CEP (AwesomeAPI `/json/{cep}`) para reduzir latência/custo.

## Diagrama (Contexto)

```mermaid
flowchart LR
  %% C4-L1: System Context (representado em Mermaid)

  person["Pessoa/Sistema Consumidor\n(Outros microsserviços, Front-ends, Integrações)"]

  subgraph fbso["Plataforma FBSO"]
    geo["Software System\nms-geolocalidade\nREST API de geolocalização por CEP"]
    batch["Software System\nbatch-geolocalidade\nCarga/atualização das tabelas DTB/IBGE"]
  end

  awesome["External Software System\nAwesomeAPI (API CEP)\n/json/{cep} e /search"]
  dtb[("Data Store\nPostgreSQL\nTabelas DTB/IBGE (schema localidade)")]

  person -->|"HTTP/JSON"| geo
  geo -->|"Consulta CEP/raio"| awesome
  geo -->|"Leitura enriquecimento"| dtb
  batch -->|"Escrita/atualização"| dtb
```

## Principais entradas/saídas

- Entradas: chamadas HTTP para endpoints `/api/v1/localidades/**`.
- Saídas:
  - HTTP para AwesomeAPI.
  - SQL (leitura) para PostgreSQL com dados DTB/IBGE.

## Premissas e limites

- Autenticação/autorização não está representada aqui (não há evidência no projeto de um mecanismo de auth no nível de controller).
- `batch-geolocalidade` é tratado como sistema externo no contexto do `ms-geolocalidade` (ele é responsável por popular as tabelas DTB/IBGE). 
