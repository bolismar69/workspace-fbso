# Política Organizacional: Cálculo de Impostos Corporativos

**Código:** POLICE-FIN-00001  
**Versão:** 1.1  
**Classificação:** Confidencial — Uso Interno  
**Owner:** Tax Compliance Officer (Diretoria Financeira)  
**Última Atualização:** 2026-06-21  
**Cadência de Revisão:** Semestral (Jun/Dez) ou por evento (nova legislação)  
**Aprovadores:** CFO · Controller · Head of Engineering · Comitê Fiscal  

---

## 1. Propósito

Esta política estabelece o framework normativo para apuração, cálculo, retenção e recolhimento de tributos incidentes sobre as operações da empresa no setor de **Telecomunicações**, abrangendo:

- **Regime atual** (Legacy): IRPJ, CSLL, PIS/Pasep, COFINS, IPI, CPP, FGTS, ICMS e ISS.
- **Regime de transição** (Reforma Tributária 2026–2033): CBS, IBS e IS.
- **Convivência** entre os dois modelos durante o período de _hybrid mode_.

O descumprimento desta política expõe a organização a passivos fiscais, multas regulatórias (até 150% do tributo devido), responsabilização criminal dos administradores (Lei 8.137/90) e riscos reputacionais.

---

## 2. Escopo

### 2.1 Incluído

| Dimensão | Abrangência |
|:---|:---|
| **Entidades legais** | Todas as PJ do grupo com CNPJ ativo no Brasil. |
| **Regimes tributários** | Lucro Real, Lucro Presumido e Simples Nacional (quando aplicável). |
| **Operações** | Vendas de mercadorias, prestação de serviços de telecom, importação, transferência entre estabelecimentos, devoluções e remessas. |
| **Sistemas** | Microserviços `ms-billing-engine-tax-rates`, `ms-tax-nexus-taas-calc-engine`, `ms-tax-individual-income` e lib `taxnexus-billing-core-lib`. |
| **Horizonte temporal** | Janeiro/2026 a Dezembro/2033 (ciclo completo da Reforma Tributária). |

### 2.2 Excluído

- Tributos sobre pessoa física (IRPF, INSS empregado) — cobertos por `ms-tax-individual-income`.
- Obrigações acessórias de entrega (SPED, EFD, DCTF) — escopo do ERP corporativo.
- Tributos de importação (II, AFRMM, Taxa Siscomex) — cobertos por política específica de comércio exterior.

---

## 3. Governança Tributária

### 3.1 Estrutura de Responsabilidades

| Papel | Responsável | Autoridade |
|:---|:---|:---|
| **Tax Compliance Officer** | Diretor Financeiro ou designado | Aprovar alíquotas, mudanças de interpretação e o plano de contas tributário. |
| **Comitê Fiscal** | CFO + Controller + Head of Engineering + Tax Compliance Officer | Reunir-se trimestralmente; validar shadow runs; aprovar alterações na política. |
| **Engineering Lead — Tax Engine** | Tech Lead do ms-billing-engine-tax-rates | Manter o código dos cálculos; garantir rastreabilidade; revisar PRs de alteração tributária. |
| **Controller** | Controller Corporativo | Conciliar apuração × contabilidade; prover dados de lucro e folha para o engine de período. |
| **Auditoria Interna** | Time de Compliance | Testar anualmente a efetividade dos controles (SOX-like); reportar ao Conselho. |

### 3.2 Matriz RACI

| Atividade | R (Responsável) | A (Accountable) | C (Consultado) | I (Informado) |
|:---|:---|:---|:---|:---|
| Definição de alíquotas | Tax Compliance Officer | CFO | Consultoria Tributária externa | Comitê Fiscal |
| Implementação no código | Engineering Lead | Tax Compliance Officer | Controller | Head of Engineering |
| Shadow run CBS/IBS | Engineering Lead | Comitê Fiscal | Controller | CFO |
| Reconciliação mensal | Controller | Tax Compliance Officer | Engineering Lead | Comitê Fiscal |
| Atualização de NCM e CEST | Engenharia de Dados | Tax Compliance Officer | Consultoria Tributária | Controller |
| Gestão de créditos tributários | Controller | Tax Compliance Officer | Engineering Lead | CFO |
| Resposta a fiscalização | Tax Compliance Officer | CFO | Jurídico Corporativo | Conselho de Administração |
| Revisão semestral da política | Tax Compliance Officer | Comitê Fiscal | Auditoria Interna | Conselho de Administração |

### 3.3 Escalonamento (Escalation Path)

```
Nível 1 — Divergência de cálculo ≤ R$ 10.000
  → Engineering Lead resolve com Controller em 48h.

Nível 2 — Divergência > R$ 10.000 ou ≤ R$ 100.000
  → Tax Compliance Officer convoca reunião extraordinária em 24h.

Nível 3 — Divergência > R$ 100.000 ou risco de autuação
  → Comitê Fiscal reporta ao CFO e ao Jurídico em 12h.
  → Comunicação ao Conselho de Administração em 5 dias úteis.

Nível 4 — Risco de responsabilização criminal (Lei 8.137/90)
  → CFO aciona Conselho e assessoria jurídica externa imediatamente.
```

---

## 4. Inventário Tributário

### 4.1 Regime Atual (Legacy)

| # | Esfera | Imposto | Sigla | Base de Cálculo | Alíquota de Referência | Frequência | Microserviço |
|:---|:---|:---|:---|:---|:---|:---|:---|
| 1 | Federal | Imposto de Renda Pessoa Jurídica | **IRPJ** | Lucro Real / Presumido | 15% + 10% adicional > R$ 20.000/mês | Trimestral / Anual | Period Engine (Lucratividade) |
| 2 | Federal | Contribuição Social sobre o Lucro Líquido | **CSLL** | Lucro Real / Presumido | 9% (geral) / 20% (instituições financeiras) | Trimestral / Anual | Period Engine (Lucratividade) |
| 3 | Federal | Programa de Integração Social | **PIS/Pasep** | Receita Bruta (Cumulativo: 0,65%) ou Base de Créditos (Não-Cumulativo: 1,65%) | 0,65% ou 1,65% | Mensal | Billing Engine (Faturamento) |
| 4 | Federal | Contribuição para o Financiamento da Seguridade Social | **COFINS** | Receita Bruta (Cumulativo: 3%) ou Base de Créditos (Não-Cumulativo: 7,6%) | 3% ou 7,6% | Mensal | Billing Engine (Faturamento) |
| 5 | Federal | Imposto sobre Produtos Industrializados | **IPI** | Valor do produto industrializado (TIPI) | Variável por NCM (0% a 300%) | Por operação | Billing Engine (Faturamento) |
| 6 | Federal | Contribuição Patronal Previdenciária | **CPP** | Folha de pagamento bruta | 20% + RAT (1-3%) + Terceiros (5,8%) | Mensal | Period Engine (Folha) |
| 7 | Federal | Fundo de Garantia do Tempo de Serviço | **FGTS** | Folha de pagamento (base FGTS) | 8% | Mensal | Period Engine (Folha) |
| 8 | Estadual | Imposto sobre Circulação de Mercadorias e Serviços | **ICMS** | Valor da operação + frete + seguro + IPI | 7% a 22% (varia por UF e tipo de operação) | Mensal | Billing Engine (Faturamento) |
| 9 | Municipal | Imposto Sobre Serviços de Qualquer Natureza | **ISS/ISSQN** | Preço do serviço | 2% a 5% (legislação municipal) | Mensal | Billing Engine (Faturamento) |
| 10 | Federal | Fundo de Universalização dos Serviços de Telecomunicações | **FUST** | Receita Operacional Líquida (Valor Bruto − ICMS − PIS − COFINS) | 1% | Mensal | Billing Engine (Faturamento) |
| 11 | Federal | Fundo para o Desenvolvimento Tecnológico das Telecomunicações | **FUNTTEL** | Receita Operacional Líquida (Valor Bruto − ICMS − PIS − COFINS) | 0,5% | Mensal | Billing Engine (Faturamento) |

### 4.2 Regime Futuro (Reforma Tributária)

| # | Esfera | Imposto | Sigla | Substitui | Base de Cálculo | Alíquota Estimada | Cálculo | Microserviço |
|:---|:---|:---|:---|:---|:---|:---|:---|:---|
| 10 | Federal | Contribuição sobre Bens e Serviços | **CBS** | PIS + COFINS | Receita Bruta (sem compor a própria base — "por fora") | 8,8% a 12% (setorial) | Por operação | Billing Engine (Reforma) |
| 11 | Subnacional | Imposto sobre Bens e Serviços | **IBS** | ICMS + ISS | Receita Bruta (alíquota do destino) | Variável por município/UF de destino | Por operação | Billing Engine (Reforma) |
| 12 | Federal | Imposto Seletivo | **IS** | — (novo, extrafiscal) | Valor do produto com NCM restrito | Variável por categoria | Por operação (antes da CBS) | Billing Engine (Reforma) |

### 4.3 Regras Específicas por Imposto

#### IRPJ / CSLL
- **Lucro Real:** Base = Lucro Contábil + Adições LALUR − Exclusões LALUR − Compensação de Prejuízos (até 30% do lucro).
- **Lucro Presumido:** Base = Receita Bruta × Percentual de Presunção (8% para comércio/indústria; 32% para serviços).
- **Adicional IRPJ:** 10% sobre a parcela do lucro que exceder R$ 20.000 × número de meses do período.
- **Retenções na Fonte:** Abater IRRF/CSLL retidos por clientes do total apurado no período.

#### PIS / COFINS
- **Regime Cumulativo:** Alíquota total 3,65% sobre receita bruta, sem direito a crédito.
- **Regime Não-Cumulativo:** Alíquota total 9,25% sobre receita bruta, com crédito sobre insumos (CST 50).
- **CST PIS:** 01 (tributado), 02 (tributado com isenção), 50 (crédito), 73 (não tributado).
- **CST COFINS:** 01 (tributado), 02 (tributado com isenção), 50 (crédito), 73 (não tributado).
- **Setor TELECOM:** Serviços de telecomunicação estão sujeitos ao regime não-cumulativo.

#### IPI
- Alíquota determinada pela TIPI (Tabela de Incidência do IPI) conforme NCM do produto.
- Base de cálculo: valor do produto industrializado.
- **CST IPI:** 00 (tributado), 01 (tributado com alíquota zero), 02 (isento), 50 (crédito).
- Não incide sobre serviços de telecomunicação puros; incide sobre equipamentos/similares fornecidos.

#### ICMS
- **Alíquota interestadual:** 4% (produtos importados com conteúdo de importação > 40%, conforme Resolução Senado 13/2012); 7% (origem Sul/Sudeste — exceto Espírito Santo — para outras regiões); 12% (demais casos).
- **DIFAL (EC 87/2015):** Vendas interestaduais para consumidor final não contribuinte → diferença entre alíquota interna do destino e alíquota interestadual.
- **ICMS-ST (Substituição Tributária):** Recolhido antecipadamente pelo remetente; base = (Valor + MVA) × alíquota interna.
- **CST ICMS:** 00 (tributado integralmente), 10 (tributado com ST), 40 (isento), 60 (ST retido anteriormente).
- **FCP (Fundo de Combate à Pobreza):** Adicional de até 2% sobre operações específicas em alguns estados.
- **ICMS Desonerado (Benefício Fiscal):**
  - **Motivo da Desoneração (motDesICMS):** Códigos 1 a 12 e 90, conforme tabela oficial da SEFAZ (ex: 3-Uso na agropecuária, 7-SUFRAMA, 9-Outros, 12-Órgão Público).
  - **CSTs que permitem desoneração:** 20, 30, 40, 41, 50, 70, 90. CST 00 (tributado integralmente) NÃO permite desoneração.
  - **Cálculo com redução de base:** `Base_Reduzida = Valor_Item × (1 − Percentual_Redução)`; `ICMS = Base_Reduzida × Alíquota`; `vICMSDeson = (Valor_Item × Alíquota) − ICMS`.
  - **Cálculo com limitação de alíquota efetiva:** `Índice_Redução = 1 − (Alíquota_Alvo / Alíquota_Nominal)`; `Base_Reduzida = Valor_Item × (Alíquota_Alvo / Alíquota_Nominal)`.
  - **Abatimento no total:** O valor do ICMS desonerado (`vICMSDeson`) deve ser subtraído do valor total da nota fiscal.
  - **Simples Nacional:** A desoneração "clássica" (com destaque na NF-e) é voltada para o Regime Normal (Lucro Real/Presumido). Empresas do Simples Nacional tratam isenções/reduções diretamente na apuração mensal (PGDAS-D) e geralmente não preenchem o campo de ICMS desonerado na nota.

#### ISS
- Alíquota definida por legislação municipal (2% a 5%).
- **Lista de Serviços (LC 116/2003):** Serviços de telecomunicação = item 1.05.
- Município competente: local do estabelecimento prestador (regra geral) ou local da prestação.
- **Retenção de ISS:** O tomador pode reter o ISS na fonte se previsto em legislação municipal.

#### CBS (Reforma)
- **Cálculo "por fora":** `CBS_Valor = Base × Aliquota_CBS` — a CBS não integra a própria base.
- **Não-cumulatividade plena:** Crédito financeiro sobre todos os insumos (diferente do modelo PIS/COFINS que exige pertinência).
- **Alíquota de referência em 2026 (Beta):** 0,1% (teste).
- **Alíquota estimada setorial (TELECOM):** 10,5% a 12% (a confirmar pelo Governo Federal).
- **Alíquota plena (2027+):** Definida pelo Ministério da Fazenda.

#### IBS (Reforma)
- **Princípio do destino:** A alíquota aplicável é a do município/estado onde o serviço é consumido.
- **Consulta ao Comitê Gestor:** Obrigatória via API para cada operação (alíquota em tempo real).
- **Cache permitido:** Até 24h para mesma jurisdição de destino. Expirado o cache, reconsultar.
- **Alíquota de referência em 2026 (Beta):** 0,9% (teste).
- **Split IBS:** A alíquota total se divide em parcela estadual + parcela municipal.

#### IS (Imposto Seletivo)
- **Gatilho obrigatório:** Antes do cálculo da CBS, verificar se o NCM do item está na lista de produtos sujeitos ao IS.
- **Produtos sujeitos (lista preliminar):** Bebidas alcoólicas, tabaco, açúcar, veículos poluentes, combustíveis fósseis.
- **Impacto em TELECOM:** Normalmente NÃO incide sobre serviços de telecomunicação puros. Incide se houver fornecimento de equipamentos com NCM restrito.
- **Flag técnica:** `isento_is` no modelo de dados — se `true`, o IS é zero independente do NCM.

#### FUST / FUNTTEL (Contribuições Setoriais de Telecom)

**FUST (Fundo de Universalização dos Serviços de Telecomunicações):**
- **Base de cálculo:** Receita Operacional Líquida = Valor Bruto do Serviço − ICMS − PIS − COFINS.
- **Alíquota:** 1%.
- **Incidência:** Exclusivamente sobre serviços de telecomunicação (SCM, STFC). **Não incide** sobre SVA (Serviço de Valor Adicionado) como streaming, antivírus, suporte técnico.
- **Fundamento:** Lei 9.998/2000.

**FUNTTEL (Fundo para o Desenvolvimento Tecnológico das Telecomunicações):**
- **Base de cálculo:** A mesma do FUST (Receita Operacional Líquida).
- **Alíquota:** 0,5%.
- **Incidência:** Idêntica ao FUST — apenas SCM e STFC.
- **Fundamento:** Lei 10.052/2000.

**Ordem de cálculo:** ICMS → PIS → COFINS → FUST → FUNTTEL. FUST e FUNTTEL são calculados em cascata porque dependem do valor líquido após os impostos principais.

**Segregação SCM × SVA:** Serviços de Valor Adicionado (SVA) não compõem a base de FUST/FUNTTEL. O motor deve segregar itens de SVA (natureza = SVA) antes do cálculo destas contribuições.

---

## 5. Regras de Convivência entre Regimes

### 5.1 Roadmap de Transição

| Fase | Período | Regime PIS/COFINS | Regime CBS | Regime ICMS | Regime ISS | Regime IBS | Status do Código |
|:---|:---|:---|:---|:---|:---|:---|:---|
| **Fase 1 — Shadow Run** | 2026 | Produção (Main) | Beta (0,1% teste) | Produção (Main) | Produção (Main) | Beta (0,9% teste) | **Hybrid — Beta** |
| **Fase 2 — CBS Plena** | 2027 | Extinção gradual | Produção (Main) | Produção (Main) | Produção (Main) | Beta | **Hybrid — Main CBS** |
| **Fase 3 — Transição Subnacional** | 2029–2032 | Extinto | Produção (Main) | Redução gradual (2% a.a.) | Redução gradual | Produção (Main) | **Hybrid Mode** |
| **Fase 4 — IVA Dual** | 2033+ | Extinto | Produção (Main) | Extinto | Extinto | Produção (Main) | **Full Reform** |

### 5.2 Regra de Cálculo por Fase

```
SE fase = "Shadow Run" (2026):
    Imposto_Total = PIS + COFINS + ICMS + ISS + IPI + IRPJ + CSLL + CPP + FGTS + FUST + FUNTTEL
    Shadow_CBS = Base × 0,001     // Cálculo paralelo — não compõe Total a Pagar
    Shadow_IBS = Base × 0,009     // Cálculo paralelo — não compõe Total a Pagar

SE fase = "CBS Plena" (2027):
    Imposto_Total = CBS + ICMS + ISS + IPI + IRPJ + CSLL + CPP + FGTS + FUST + FUNTTEL
    // PIS/COFINS extintos; ICMS e ISS ainda em Produção

SE fase = "Transição Subnacional" (2029–2032):
    Imposto_Total = CBS + (ICMS × fator_reducao_ICMS) + (ISS × fator_reducao_ISS)
                  + IBS + IPI + IRPJ + CSLL + CPP + FGTS + FUST + FUNTTEL
    // fator_reducao_ICMS diminui 2 pontos percentuais ao ano até zero
    // fator_reducao_ISS diminui proporcionalmente

SE fase = "IVA Dual Pleno" (2033+):
    Imposto_Total = CBS + IBS + IPI + IRPJ + CSLL + CPP + FGTS + IS + FUST + FUNTTEL
    // ICMS e ISS totalmente extintos
    // FUST e FUNTTEL seguem incidindo sobre a base após CBS e IBS (a confirmar pelo regulador)
```

### 5.3 Matriz de Convência para Telecom

| Natureza da Operação | 2026 (Beta) | 2027 | 2029–2032 | 2033+ |
|:---|:---|:---|:---|:---|
| Serviço de telefonia (intra-estadual) | PIS+COFINS+ICMS+ISS+FUST+FUNTTEL | CBS+ICMS+ISS+FUST+FUNTTEL | CBS+(ICMS×fator)+(ISS×fator)+IBS+FUST+FUNTTEL | CBS+IBS+FUST+FUNTTEL |
| Serviço de telefonia (interestadual) | PIS+COFINS+ICMS+DIFAL+ISS+FUST+FUNTTEL | CBS+ICMS+DIFAL+ISS+FUST+FUNTTEL | CBS+(ICMS_DIFAL×fator)+(ISS×fator)+IBS+FUST+FUNTTEL | CBS+IBS+FUST+FUNTTEL |
| Venda de equipamento (intra-estadual) | PIS+COFINS+ICMS+IPI | CBS+ICMS+IPI | CBS+(ICMS×fator)+IPI+IBS+IS(se NCM restrito) | CBS+IBS+IPI+IS(se NCM restrito) |
| Serviço de internet/dados | PIS+COFINS+ISS+FUST+FUNTTEL | CBS+ISS+FUST+FUNTTEL | CBS+(ISS×fator)+IBS+FUST+FUNTTEL | CBS+IBS+FUST+FUNTTEL |

---

## 6. Matriz de Risco Tributário

Aplicando o modelo **Probabilidade × Impacto = Risk Score** (DOJ Compliance Framework):

| # | Risco | Probabilidade | Impacto Financeiro | Score | Resposta | Controle |
|:---|:---|:---|:---|:---|:---|:---|
| R1 | Divergência CBS × PIS/COFINS na Fase 1 (Beta) | **Alta** — alíquotas de teste vs. produção | **Alto** — R$ 50K–500K/mês | 🔴 **9** | Shadow run com reconciliação mensal | `reconcile_shadow_run()` — output em log imutável |
| R2 | API do Comitê Gestor do IBS indisponível | **Média** — sistema novo, SLAs incertos | **Alto** — impossibilidade de calcular IBS | 🟠 **6** | Fallback com última alíquota cacheada + notificação imediata | Circuit breaker + retry 3× com backoff exponencial |
| R3 | Classificação NCM incorreta para IS | **Média** — tabela NCM em migração | **Alto** — IS indevido ou omitido | 🟠 **6** | Dupla verificação: taxonomia oficial + consultoria tributária | Validação pré-commit da tabela `ncm_seletivo` |
| R4 | Alíquota CBS setorial (TELECOM) maior que a provisionada | **Baixa** — divulgação prévia esperada | **Alto** — impacto no fluxo de caixa | 🟡 **3** | Provisionamento de margem de 200bps | Revisão trimestral da provisão pelo Comitê Fiscal |
| R5 | Falha no Hybrid Mode (ICMS coexistindo com IBS) | **Alta** — complexidade de dois regimes simultâneos | **Médio** — double-counting ou sub-apuração | 🟠 **6** | Testes de regressão bimestrais com massa histórica | Suite de testes `hybrid_mode_test.go` |
| R6 | Alteração unilateral de alíquota por UF/município | **Alta** — 27 UFs + 5.570 municípios | **Médio** — R$ 5K–50K por evento | 🟠 **6** | Monitoramento de diários oficiais + feed do Comitê Gestor | Job diário `sync_aliquota_ibc.go` |
| R7 | Interpretação divergente de crédito de PIS/COFINS em insumos de TELECOM | **Média** — jurisprudência instável | **Médio** — autuação retroativa | 🟡 **4** | Documentar interpretação adotada + jurisprudência de suporte | Parecer tributário anual |
| R8 | Erro de arredondamento acumulado em grandes volumes | **Média** — milhares de notas/dia | **Baixo** — centavos por nota | 🟢 **2** | Precisão decimal com 4 casas; arredondamento só no total final | `utils.RoundDecimalWithUintPrecision()` |
| R9 | Falha na compensação cruzada de tributos (PER/DCOMP) | **Baixa** — processo maduro | **Médio** — perda de crédito legítimo | 🟡 **3** | Validação de elegibilidade antes de cada PER/DCOMP | Regras em `compensacao_tributaria.go` |
| R10 | Vazamento de dados fiscais via API IBS | **Baixa** — TLS 1.3 + mTLS | **Crítico** — LGPD + sigilo fiscal | 🟡 **3** | Criptografia em trânsito e em repouso; logs anonimizados | Revisão de segurança trimestral |

---

## 7. Controles Internos (SOX-like)

### 7.1 Inventário de Controles

| ID | Controle | Tipo | Frequência | Evidência |
|:---|:---|:---|:---|:---|
| CTL-001 | Shadow run CBS/IBS com reconciliação vs. PIS/COFINS/ICMS/ISS | Detectivo + Automatizado | Mensal | Relatório `shadow_run_reconcile_YYYY-MM.csv` assinado pelo Controller |
| CTL-002 | Validação de alíquotas antes de deploy em produção | Preventivo + Manual | Por release | PR aprovado pelo Tax Compliance Officer |
| CTL-003 | Teste de regressão do Hybrid Mode | Detectivo + Automatizado | Bimestral | Log da suite `hybrid_mode_test.go` com 100% pass |
| CTL-004 | Verificação de integridade da tabela NCM Seletivo | Preventivo + Automatizado | Semanal | Hash SHA-256 do arquivo `ncm_seletivo.csv` |
| CTL-005 | Circuit breaker para API do Comitê Gestor IBS | Automatizado | Contínuo | Métrica `ibs_api_circuit_state` no dashboard de monitoração |
| CTL-006 | Conciliação IRPJ/CSLL × Balancete Contábil | Detectivo + Manual | Trimestral | Planilha de conciliação com assinatura do Controller + Tax Compliance Officer |
| CTL-007 | Verificação de retenções na fonte (IRRF/CSLL/PIS/COFINS/ISS) | Detectivo + Automatizado | Mensal | Relatório `retencoes_na_fonte_YYYY-MM.csv` |
| CTL-008 | Aprovação de mudança de alíquota no código | Preventivo + Manual | Por evento | PR com dupla aprovação (Engineering Lead + Tax Compliance Officer) |

### 7.2 Trilha de Auditoria

Todo cálculo de imposto deve produzir registro imutável contendo:

```
{
  "timestamp": "2026-06-21T15:30:00Z",
  "transaction_id": "NF-2026-000001",
  "tax_type": "CBS",
  "phase": "SHADOW_RUN",
  "input_base": 1000.00,
  "rate_applied": 0.001,
  "calculated_amount": 1.00,
  "rate_source": "STATIC_TEST",
  "engine_version": "v2.3.1",
  "git_commit": "a1b2c3d4",
  "calculated_by": "ms-billing-engine-tax-rates"
}
```

Logs de auditoria devem ser:
- **Imutáveis:** Append-only, sem update ou delete.
- **Retidos:** Prazo mínimo de 6 anos (prazo decadencial do CTN + 1 ano de margem).
- **Indexados:** Por `transaction_id`, `tax_type` e `timestamp`.

---

## 8. Procedimentos Operacionais Padrão (SOPs)

### SOP-001: Cálculo de CBS "por fora" em Serviços de Telecom

- **Quem:** `ms-billing-engine-tax-rates` (Billing Engine — Reforma)
- **Quando:** Para cada nota fiscal / operação de saída a partir de Janeiro/2026.
- **Como:**
  1. Receber `TaxDocumentInput` com `valor_bruto`, `c_class_trib` e fase ativa.
  2. Determinar alíquota CBS: se fase = "Shadow", usar `0,001`; se fase = "CBS Plena", consultar tabela `cbs_rates` por `c_class_trib`.
  3. Calcular: `Valor_CBS = Base × Aliquota_CBS` (cálculo "por fora" — a CBS não integra a própria base).
  4. Se fase = "Shadow", armazenar em `shadow_tax_result` (não compor `total_a_pagar`).
  5. Registrar em trilha de auditoria.
- **Output:** Valor CBS calculado, registrado em `TaxResponse.tributos[]`.

### SOP-002: Consulta e Cache de Alíquotas IBS por Destino

- **Quem:** `ms-billing-engine-tax-rates` (Billing Engine — Reforma)
- **Quando:** Para cada operação com `destino.municipio_codigo_ibge` diferente de vazio.
- **Como:**
  1. Verificar cache Redis com chave `ibs:rate:{ibge_code}:{date}`. TTL = 24h.
  2. Cache MISS → chamar API do Comitê Gestor IBS: `GET /api/v1/rates?ibge_code={code}`.
  3. Se API retornar erro, aplicar circuit breaker (3 falhas em 60s → usar última taxa cacheada + alerta).
  4. Split: `aliquota_estadual` + `aliquota_municipal` = `aliquota_total_ibs`.
  5. Armazenar no cache com TTL e em trilha de auditoria.
- **Output:** Alíquotas IBS (estadual + municipal) disponíveis para o cálculo.

### SOP-003: Classificação NCM e Avaliação de Incidência do IS

- **Quem:** Pipeline de dados + `ms-billing-engine-tax-rates`
- **Quando:** Antes do cálculo da CBS (o IS é pré-filtro).
- **Como:**
  1. Receber `NCM` do item.
  2. Consultar tabela `ncm_seletivo` (sincronizada semanalmente com a lista oficial).
  3. Se NCM constar na tabela E `isento_is = false` → IS aplicável.
  4. Se IS aplicável: `Valor_IS = Base × Aliquota_IS_Categoria`.
  5. Registrar em trilha de auditoria mesmo se IS = 0 (com flag `is_exempt`).
- **Output:** Valor IS calculado ou zero, com motivo documentado.

### SOP-004: Reconciliação Mensal Shadow Run (CBS/IBS vs. PIS/COFINS/ICMS/ISS)

- **Quem:** Controller + Engineering Lead
- **Quando:** Dia 5 de cada mês, referente ao mês anterior.
- **Como:**
  1. Extrair totais mensais do Legacy: PIS, COFINS, ICMS, ISS.
  2. Extrair totais mensais do Shadow: CBS, IBS.
  3. Somar Legacy (PIS+COFINS) vs. CBS; (ICMS+ISS) vs. IBS.
  4. Calcular `variação_percentual = (Shadow − Legacy) / Legacy × 100`.
  5. Se variação > 10%, investigar item a item.
  6. Gerar relatório `shadow_run_reconcile_YYYY-MM.csv`.
  7. Controller assina o relatório; Tax Compliance Officer revisa.
- **Output:** Relatório de reconciliação assinado, arquivado na pasta de compliance.

### SOP-005: Atualização de Alíquotas por Nova Legislação

- **Quem:** Tax Compliance Officer + Engineering Lead
- **Quando:** Até 5 dias úteis após publicação de decreto/lei que altere alíquotas.
- **Como:**
  1. Tax Compliance Officer identifica a alteração e registra em `tax_rate_change_log`.
  2. Engineering Lead cria branch `tax-update/YYYY-MM-DD-description`.
  3. Atualizar tabela SQL de alíquotas (`billing_tax_rates.icms_rules`, `cbs_rates`, etc.).
  4. Atualizar constantes em `models/constants.go` se aplicável.
  5. PR com dupla aprovação (Engineering + Tax).
  6. Deploy em staging → teste de regressão → deploy em produção.
  7. Comunicar financeiro sobre vigência e impacto estimado.
- **Output:** Alíquotas atualizadas em produção com registro em `tax_rate_change_log`.

### SOP-006: Tratamento de Exceções e Creditamento

- **Quem:** `ms-billing-engine-tax-rates` + Controller
- **Quando:** Na apuração mensal/trimestral.
- **Como:**
  1. Consolidar créditos do período (entradas/compras) por tributo.
  2. Consolidar débitos do período (saídas/vendas) por tributo.
  3. Para cada tributo: `Saldo = Créditos − Débitos`.
  4. Se `Saldo > 0` → Crédito acumulado (Saldo a Recuperar para próximo período).
  5. Se `Saldo < 0` → Valor a Recolher (DARF/GNRE/GPS).
  6. Retenções na fonte abatem diretamente o valor a recolher.
  7. Compensação cruzada (PER/DCOMP): apenas se aprovado pelo Tax Compliance Officer.
- **Output:** `TaxPeriodResult` com `total_a_pagar` e `saldos_a_recuperar[]`.

---

## 9. Compensação e Creditamento Tributário

### 9.1 Tipos de Compensação

| Tipo | Tributos Envolvidos | Mecanismo |
|:---|:---|:---|
| **Não-Cumulatividade** | ICMS, IPI, PIS, COFINS | Crédito sobre insumos → abate débito do período |
| **Retenção na Fonte** | IRRF, CSLL, PIS, COFINS, ISS | Cliente retém → abate do total apurado |
| **ICMS-ST** | ICMS | Indústria recolhe antecipado → varejo não paga na revenda |
| **Compensação Cruzada (PER/DCOMP)** | Tributos Federais (PIS/COFINS → IRPJ/CSLL/CPP) | Crédito excedente de um tributo federal abate outro |

### 9.2 Regras de Creditamento

- **Crédito de PIS/COFINS (Não-Cumulativo):** Somente sobre insumos diretamente relacionados à atividade-fim. Bens de capital: crédito em 12 parcelas mensais.
- **Crédito de ICMS:** Bens de uso/consumo e ativo imobilizado: crédito em 48 parcelas mensais (a partir de 2033). Energia elétrica e serviços de comunicação: crédito integral se comprovadamente insumo.
- **Crédito de IPI:** Matéria-prima, produto intermediário e material de embalagem adquiridos para industrialização.
- **Crédito de CBS (Reforma):** Não-cumulatividade plena — todo insumo gera crédito financeiro, independentemente de pertinência ao objeto social.

---

## 10. Obrigações de Conformidade e Sanções

### 10.1 Consequências por Não Conformidade

| Natureza | Gravidade | Consequência |
|:---|:---|:---|
| Erro de cálculo sem dolo, valor ≤ R$ 10.000 | Leve | Correção + treinamento adicional |
| Erro de cálculo sem dolo, valor > R$ 10.000 | Média | Investigação de causa raiz + ação corretiva documentada |
| Omissão de imposto com evidência de negligência | Grave | Sanção administrativa + reporte ao Comitê Fiscal |
| Sonegação fiscal dolosa (Lei 8.137/90) | Crítica | Demissão por justa causa + comunicação ao Ministério Público |
| Falha de controle interno não reportada em 48h | Grave | Sanção ao gestor da área + revisão do ambiente de controle |

### 10.2 Canal de Denúncia

Reportar suspeitas de irregularidades tributárias via:
- **Canal confidencial:** `compliance@empresa.com.br`
- **Política de não retaliação:** Aderente à Lei 13.608/2018 e Dodd-Frank (whistleblower).
- **Investigação:** Auditoria Interna em até 15 dias úteis.

---

## 11. Gestão de Terceiros

### 11.1 Comitê Gestor do IBS como Fornecedor Crítico

| Dimensão | Requisito |
|:---|:---|
| **SLA de disponibilidade** | 99,9% (contratual) |
| **Fallback** | Última alíquota cacheada por 24h + notificação ao Tax Compliance Officer |
| **Homologação** | Ambiente de sandbox do Comitê Gestor testado antes de cada release |
| **Monitoração** | Métrica `ibs_api_latency_p99` < 500ms; `ibs_api_error_rate` < 0,1% |
| **Due diligence** | Revisão anual da qualidade dos dados fornecidos pelo Comitê Gestor |

### 11.2 Consultoria Tributária Externa

- **Escopo:** Interpretação de legislação, classificação NCM/CEST, suporte em fiscalizações.
- **Contrato:** Deve incluir cláusula de confidencialidade e obrigação de atualização em até 48h após nova legislação.
- **Revisão anual:** Avaliar qualidade das entregas e renovar/competir contrato.

---

## 12. Métricas de Performance do Motor de Cálculo

| Métrica | Alvo | Como Medir |
|:---|:---|:---|
| Precisão de cálculo (Shadow Run) | ≥ 99,5% de acerto vs. cálculo manual | Relatório de reconciliação mensal |
| Latência de cálculo por nota | p99 < 200ms | APM (Application Performance Monitoring) |
| Disponibilidade da API IBS | ≥ 99,9% | Dashboard de monitoração |
| Tempo para atualizar alíquota após nova lei | ≤ 5 dias úteis | `tax_rate_change_log` |
| Cobertura de testes do Hybrid Mode | 100% dos cenários mapeados | Relatório de cobertura `go test -cover` |
| Volume de exceções (erros de cálculo) | ≤ 0,01% das operações | Log aggregation |
| Tempo de resposta a divergências > R$ 10K | ≤ 48h | SLA tracker no Jira de Compliance |

---

## 13. Referências

### 13.1 Legislação

| Norma | Assunto |
|:---|:---|
| Lei 8.137/1990 | Crimes contra a ordem tributária |
| Lei Complementar 116/2003 | ISS |
| Lei 10.833/2003 | PIS/COFINS Não-Cumulativo |
| Lei 9.718/1998 | PIS/COFINS Cumulativo |
| Lei 8.212/1991 | CPP (INSS Patronal) |
| Lei 8.036/1990 | FGTS |
| Emenda Constitucional 87/2015 | DIFAL |
| Emenda Constitucional 132/2023 | Reforma Tributária (CBS, IBS, IS) |
| Lei Complementar 214/2025 (prevista) | Regulamentação do IBS |
| Lei 9.998/2000 | FUST — Fundo de Universalização dos Serviços de Telecomunicações |
| Lei 10.052/2000 | FUNTTEL — Fundo para o Desenvolvimento Tecnológico das Telecomunicações |
| Resolução Senado 13/2012 | Alíquota ICMS 4% para produtos importados |
| Código Tributário Nacional (Lei 5.172/1966) | Normas gerais de direito tributário |

### 13.2 Documentos Relacionados

| Documento | Código / Local |
|:---|:---|
| Escopo do Projeto de Microserviços Tributários | `docs/README-ESCOPO.md` |
| Detalhamento ICMS | `docs/README-ICMS.md` |
| Extensão de Regras ICMS | `docs/README-ICMS-EXTENSAO-REGRAS.md` |
| Detalhamento ICMS-ST e Desoneração | `docs/README-ICSM-TAXA-DESONERACAO.md` |
| Detalhamento PIS/COFINS | `docs/README-PIS-COFINS.md` |
| Adendo PIS/COFINS | `docs/README-PIS-COFINS-ADENDO.md` |
| Desoneração PIS/COFINS | `docs/README-PIS-COFINS-DESONERACAO.md` |
| Detalhamento IPI | `docs/README-IPI.md` |
| Tabela CST/CSOSN | `docs/README-TABELA-CST-CSON.md` |
| Simples Nacional | `docs/README-SIMPLES-NACIONAL.md` |
| Constantes do Motor | `docs/README-CONSTANTS.md` |
| Brainstorm de Arquitetura | `docs/README-BRAINSTORM.md` |
| Especificação técnica do modelo de dados | `backend/go/libs/go-native/taxnexus-billing-core-lib/models/` |
| Código-fonte do motor de cálculo | `backend/go/fiber/microservices/ms-billing-engine-tax-rates/` |

### 13.3 Ferramentas Externas

| Recurso | URL |
|:---|:---|
| Calculadora CBS (Piloto — Governo Federal) | https://piloto-cbs.tributos.gov.br/servico/calculadora-consumo/calculadora |
| Guia de Integração API CBS | https://piloto-cbs.tributos.gov.br/servico/calculadora-consumo/calculadora/documentacao/guia-integracao |
| Swagger API Calculadora CBS | https://piloto-cbs.tributos.gov.br/servico/calculadora-consumo/api/swagger-ui/index.html |
| Tabela NBS (Nomenclatura de Bens e Serviços) | https://www.unimake.com.br/downloads/tabela_nbs.json |
| Portal NFE (Nota Fiscal Eletrônica) | https://www.nfe.fazenda.gov.br/ |
| Portal Único Siscomex (Consulta NCM) | https://portalunico.siscomex.gov.br/classif/ |
| TIPI (Tabela IPI) | https://www.gov.br/receitafederal/pt-br/acesso-a-informacao/legislacao/documentos-e-arquivos/tipi.xlsx |
| Projeto ACBr (código de referência) | https://github.com/vilsonneto/tributos-br |
| Discussão OCA/l10n-brazil | https://github.com/OCA/l10n-brazil/discussions/4237 |

---

## 14. Controle de Versão

| Versão | Data | Autor | Alterações |
|:---|:---|:---|:---|
| 1.0 | 2026-06-21 | Comitê Fiscal | Versão inicial. Cobre regime Legacy, Reforma Tributária, matriz de riscos, SOPs, governança e transição 2026–2033. |
| 1.1 | 2026-06-21 | Comitê Fiscal + Engineering Lead | Adicionados FUST (1%) e FUNTTEL (0,5%) como contribuições setoriais de TELECOM (§4.1, §4.3, §5.2, §5.3). Adicionada alíquota ICMS 4% para importados (§4.3). Adicionada seção de ICMS Desonerado (§4.3). Referências atualizadas (§13.1). |

---

**Próxima Revisão Obrigatória:** 2026-12-21 (ou antes, se publicada a Lei Complementar do IBS ou definida a alíquota CBS setorial para TELECOM).

---

_Documento classificado como **Confidencial — Uso Interno**. Distribuição controlada pelo Tax Compliance Officer. Cópia não autorizada constitui violação da política de segurança da informação._
