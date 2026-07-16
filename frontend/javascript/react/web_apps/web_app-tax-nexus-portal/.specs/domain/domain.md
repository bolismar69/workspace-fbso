---
title: "Domínio — web_app-tax-nexus-portal"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["domain", "glossary", "ubiquitous-language", "tax", "reforma-tributaria"]
---

# Glossário de Domínio — TaxNexus Portal (TaaS)

## 1. Introdução

Este glossário define a linguagem ubíqua do domínio de **simulação tributária da Reforma Tributária 2026** no Brasil. O portal TaxNexus TaaS permite que contribuintes (pessoas jurídicas) simulem o impacto da transição do sistema tributário atual (PIS, COFINS, ICMS, IPI, ISS) para o novo modelo de IVA Dual (CBS + IBS) com Imposto Seletivo.

O domínio abrange: tributos, regimes fiscais, classificações fiscais (NCM), localização (UF/município via código IBGE), e regras de transição.

---

## 2. Termos de Domínio

### 2.1. Identificação do Contribuinte

| Termo | Definição | Sinônimos | Código (interface/type) |
|---|---|---|---|
| **CNPJ** | Cadastro Nacional da Pessoa Jurídica. Identificador único da empresa contribuinte no Brasil (14 dígitos). | Contribuinte, Empresa | `TaxRequest.cnpj: string` |
| **IBGE** | Código de 7 dígitos do município segundo o IBGE. Usado para identificar a jurisdição fiscal do cálculo. | Código IBGE, Cidade | `TaxRequest.ibge: string` |
| **UF** | Unidade Federativa (estado). Retornado pelo backend na resposta do cálculo. | Estado | `TaxResponse.calculation.uf: string` |

### 2.2. Classificação Fiscal

| Termo | Definição | Sinônimos | Código (interface/type) |
|---|---|---|---|
| **NCM** | Nomenclatura Comum do Mercosul. Código de 8 dígitos que classifica mercadorias para fins fiscais. Ex: `62011100` (sobretudos masculinos). | Código NCM, Classificação fiscal | `TaxRequest.ncm: string` |

### 2.3. Tributos do Sistema Atual (Legado)

| Termo | Definição | Sinônimos | Código (interface/type) |
|---|---|---|---|
| **PIS** | Programa de Integração Social. Tributo federal sobre faturamento (regime cumulativo ou não-cumulativo). | PIS/PASEP | `TaxResponse.calculation.pis: number` |
| **COFINS** | Contribuição para o Financiamento da Seguridade Social. Tributo federal sobre faturamento. | — | `TaxResponse.calculation.cofins: number` |
| **ICMS** | Imposto sobre Circulação de Mercadorias e Serviços. Tributo estadual. | — | `TaxResponse.calculation.icms: number` |
| **IPI** | Imposto sobre Produtos Industrializados. Tributo federal. | — | `TaxResponse.calculation.ipi: number` |
| **ISS** | Imposto Sobre Serviços. Tributo municipal. | ISSQN | `TaxResponse.calculation.iss: number` |

### 2.4. Tributos da Reforma (IVA Dual)

| Termo | Definição | Sinônimos | Código (interface/type) |
|---|---|---|---|
| **CBS** | Contribuição sobre Bens e Serviços. Tributo federal do novo IVA Dual (substitui PIS, COFINS, IPI). | CBS calculada | `TaxResponse.calculation.cbs_calculada: number` |
| **IBS** | Imposto sobre Bens e Serviços. Tributo estadual/municipal do novo IVA Dual (substitui ICMS, ISS). | IBS calculado | `TaxResponse.calculation.ibs_calculado: number` |
| **Imposto Seletivo** | Tributo extrafiscal sobre bens e serviços prejudiciais à saúde ou meio ambiente ("sin tax"). | IS, Seletivo | `TaxResponse.calculation.imposto_seletivo: number` |

### 2.5. Tributos de Transição (Regra 2027+)

| Termo | Definição | Sinônimos | Código (interface/type) |
|---|---|---|---|
| **Saldo Remanescente** | Crédito tributário acumulado de períodos anteriores, usado como redutor na transição para o novo sistema. | Crédito acumulado | `TaxRequest.saldo_remanescente: number` |
| **IPVA Novo** | Projeção de IPVA no contexto da reforma — ampliação de base para veículos de luxo (aquáticos e aéreos). | — | `TaxResponse.calculation.ipva_novo: number` |
| **ITCMD Novo** | Projeção de ITCMD (Imposto sobre Transmissão Causa Mortis e Doação) com alíquota progressiva na reforma. | — | `TaxResponse.calculation.itcmd_novo: number` |

### 2.6. Integração

| Termo | Definição | Sinônimos | Código (interface/type) |
|---|---|---|---|
| **Cadastro Único** | Identificador de cadastro retornado pelo backend para tracking da simulação no sistema corporativo. | Callback, ID de rastreio | `TaxResponse.callback.id_cadastro_unico: string` |
| **Transaction Status** | Status da transação de cálculo retornada pelo backend. | Status | `TaxResponse.transaction_status: string` |

---

## 3. Relações entre Conceitos

```
Contribuinte (CNPJ)
  └─ Localizado em: UF + Município (IBGE)
  └─ Classifica mercadoria: NCM
  └─ Submete simulação:
       ├─ Entrada: CNPJ + NCM + IBGE + Saldo Remanescente
       └─ Saída:
            ├─ Tributos Legados: PIS, COFINS, ICMS, ISS, IPI
            ├─ Tributos Reforma: CBS, IBS, Imposto Seletivo
            ├─ Projeções: IPVA Novo, ITCMD Novo
            └─ Callback: ID Cadastro Único + Status Integração
```

---

## 4. Regras de Negócio Fundamentais

1. **CNPJ como identificador único**: Toda simulação é vinculada a um CNPJ de 14 dígitos.
2. **Jurisdição fiscal**: O cálculo tributário depende do município (código IBGE) — diferentes municípios têm diferentes alíquotas.
3. **NCM determina classificação**: O código NCM define a categoria fiscal da mercadoria e influencia alíquotas aplicáveis.
4. **Transição 2026→2027**: A simulação projeta dois anos — 2026 (transição) e 2027 (projeção com fator 1.02x sobre os valores de 2026).
5. **Saldo remanescente como redutor**: Créditos acumulados de períodos anteriores reduzem o imposto devido no novo sistema.
6. **Imposto Seletivo é extrafiscal**: Aplica-se apenas a bens específicos (tabaco, bebidas, combustíveis), não a todos os NCMs.
