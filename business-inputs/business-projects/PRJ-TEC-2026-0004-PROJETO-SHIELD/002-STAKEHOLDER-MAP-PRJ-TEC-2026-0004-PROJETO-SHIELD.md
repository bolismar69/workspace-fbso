# Stakeholder Map: PROJETO SHIELD — Plataforma de Identidade e Segurança
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documento Base** | 001-PROJECT-CHARTER-PRJ-TEC-2026-0004-PROJETO-SHIELD.md |
| **Data de Elaboração** | 07/08/2026 |
| **Versão** | 1.2 — Revisão de correção (19/08/2026): referência de entregas do PO corrigida para (D1-D2), conforme Charter (001) |
| **Metodologia** | WATERFALL |

---

## O que é um Stakeholder Map?

O **Stakeholder Map (Mapa de Partes Interessadas)** é o registro canônico de todas as pessoas, grupos e organizações que têm interesse ou influência sobre o projeto. Ele expande a Seção 5 do Project Charter com informações detalhadas de contato, responsabilidades decisórias, canais de comunicação e caminhos de escalação.

> **Nota:** Este documento lista exclusivamente **Sponsors, Stakeholders e Lideranças do Projeto** — pessoas com poder de decisão sobre escopo, orçamento, prazo e direcionamento estratégico. O time técnico (engenharia, arquitetura, QA, DevOps) será referenciado em documentos técnicos futuros (030-SAD, 035-HLD, 040-LLD, 060-EAP-WBS).

### Conexão com o Pipeline

- **UPSTREAM:** Consome a lista de stakeholders e RACI macro do 001-PROJECT-CHARTER
- **DOWNSTREAM:** Alimenta 005-BRD (mapeamento detalhado de stakeholders), 075-PLANO-COMUNICACAO (canais e frequência)

---

## 1. Identificação das Partes Interessadas (Sponsors, Stakeholders e Liderança do Projeto)

| # | Papel | Nome | Posição | Decide sobre | Contato |
|---|-------|------|---------|--------------|---------|
| **1.1** | Sponsor — Diretoria de Tecnologia | A definir | C-Level / Diretoria | Liberação de recursos orçamentários, aprovação de diretrizes estratégicas, GO/NO-GO no GATE 1 (ROM ±50%), validação do alinhamento com portfólio de produtos | A definir |
| **1.2** | Gerência Comercial | A definir | Diretor Comercial | Principal impactado — aprovação de alterações de alto impacto, priorização de funcionalidades do ponto de vista do negócio, validação dos critérios de aceite de negócio (UAT) | A definir |
| **1.2** | Gerência de Tecnologia | A definir | Diretor de Tecnologia | Responsável pela entrega do projeto — aprovação e priorização de todas as frentes do projeto para evitar atrasos e garantir qualidade | A definir |
| **1.2** | Gerência de Finanças | A definir | Diretor Financeiro | Responsável pela liberação do budget para pagamento das despesas com fornecedores e custos de infraestrutura | A definir |
| **1.3** | Product Owner (PO) | A definir | Gerente Funcional de TI | Priorização de funcionalidades, definição de critérios de aceite, trade-offs de escopo vs. prazo, aceitação formal das entregas (D1-D2) | A definir |
| **1.3** | PMO Corporativo | A definir | Gerente de Projetos | Alinhamento do cronograma com portfólio, arbitragem de conflitos de priorização, escalação de impedimentos aos sponsors e stakeholders, status reporting | A definir |

**Legenda da coluna #:**
- **1.1** = Patrocinadores Executivos (Sponsors)
- **1.2** = Lideranças de Negócio Impactadas (Stakeholders)
- **1.3** = Execução e Governança do Projeto

---

## 2. Matriz RACI por Fase do Projeto

**Legenda:** R = Responsible (executa) | A = Accountable (aprova/responde) | C = Consulted (consultado) | I = Informed (informado)

### 2.1 Fase 1 — Iniciação e Requisitos de Negócio

| Atividade | Sponsor (Dir. Tec.) | Ger. Comercial | Ger. Tecnologia | Ger. Finanças | PO | PMO |
|-----------|---------------------|----------------|-----------------|---------------|-----|------|
| Aprovação do Project Charter (001) | **A** | C | C | C | C | R |
| Identificação e registro de stakeholders (002) | C | C | C | I | **A** | R |
| Levantamento de requisitos de negócio (005-BRD) | C | **A** | C | C | R | I |
| Especificação funcional (010-FRD) | I | C | C | I | **A** | I |
| Validação da rastreabilidade de negócio (015-RTM-F1) | I | C | C | I | C | **A** |

### 2.2 Fase 2 — Especificação de Sistema e Arquitetura Macro

| Atividade | Sponsor (Dir. Tec.) | Ger. Comercial | Ger. Tecnologia | Ger. Finanças | PO | PMO |
|-----------|---------------------|----------------|-----------------|---------------|-----|------|
| Especificação de requisitos de sistema (020-SRS) | I | C | **A** | I | C | I |
| Rastreabilidade de sistema (025-RTM-F2) | I | I | C | I | I | **A** |
| Documento de Arquitetura (030-SAD) | I | I | **A** | I | C | I |
| High-Level Design (035-HLD) | I | I | **A** | I | I | I |
| **GATE 1:** Decisão GO/NO-GO (ROM ±50%) | **A** | C | C | C | C | R |

### 2.3 Fase 3 — Engenharia Detalhada e Qualidade

| Atividade | Sponsor (Dir. Tec.) | Ger. Comercial | Ger. Tecnologia | Ger. Finanças | PO | PMO |
|-----------|---------------------|----------------|-----------------|---------------|-----|------|
| Low-Level Design (040-LLD) | I | I | **A** | I | I | I |
| Estratégia de Testes (045-EST-PLAN) | I | C | C | I | C | **A** |
| Casos de Teste (050-EST-CASES) | I | C | C | I | **A** | C |
| Relatório de Qualidade (055) | I | I | C | I | C | **A** |
| EAP/WBS (060) | I | C | C | I | C | **A** |
| **GATE 2:** Decisão PERT (±15-25%) | **A** | C | C | C | C | R |

### 2.4 Fase 4 — Planejamento e Baseline

| Atividade | Sponsor (Dir. Tec.) | Ger. Comercial | Ger. Tecnologia | Ger. Finanças | PO | PMO |
|-----------|---------------------|----------------|-----------------|---------------|-----|------|
| Cronograma (065) | C | C | C | I | C | **A** |
| Orçamento (070) | **A** | C | C | C | I | R |
| Plano de Comunicação (075) | C | **A** | C | I | R | C |
| Plano de Riscos (080) | C | C | **A** | C | C | R |
| Plano de Gestão de Mudanças (085) | **A** | C | C | I | C | R |
| Deployment Plan (090) | I | I | **A** | I | C | C |

### 2.5 Fase 5 — Execução e Construção

| Atividade | Sponsor (Dir. Tec.) | Ger. Comercial | Ger. Tecnologia | Ger. Finanças | PO | PMO |
|-----------|---------------------|----------------|-----------------|---------------|-----|------|
| Backlog-Kanban (092) | I | C | C | I | **A** | R |
| Gestão de Times (093) | I | I | **A** | C | C | R |
| Definição de Janelas de Entrega (096) | C | C | **A** | I | C | R |
| Manuais de Usuário (097) | I | **A** | C | I | R | I |
| Manuais Operacionais (100) | I | I | **A** | I | C | C |

### 2.6 Fase 6 — Encerramento e Operação

| Atividade | Sponsor (Dir. Tec.) | Ger. Comercial | Ger. Tecnologia | Ger. Finanças | PO | PMO |
|-----------|---------------------|----------------|-----------------|---------------|-----|------|
| Termo de Aceite (105) | **A** | C | C | I | R | C |
| Lições Aprendidas (110) | C | C | C | I | R | **A** |
| Termo de Encerramento (115) | **A** | C | C | C | R | R |


---

## 3. Canais de Comunicação e Frequência

| Fórum | Participantes | Frequência | Objetivo | Artefato de Saída |
|-------|---------------|------------|----------|-------------------|
| **Comitê Executivo do Projeto** | Sponsor, Ger. Tecnologia, Ger. Comercial, Ger. Finanças, PMO, PO | Mensal | Aprovar direcionamento, liberar recursos orçamentários, revisar KPIs e marcos, decidir GO/NO-GO | Dashboard Executivo + Ata |
| **Reunião de Alinhamento de Negócio** | Ger. Comercial, PO, Ger. Tecnologia | Quinzenal | Validar funcionalidades concluídas contra expectativas de negócio, revisar critérios de aceite | Funcionalidades aprovadas + ajustes de prioridade |
| **Reunião de Governança e Budget** | Ger. Finanças, Sponsor, PMO | Mensal | Acompanhar execução orçamentária, aprovar desembolsos, revisar Budget/Limite | Relatório financeiro |
| **Alinhamento com PMO** | PMO, PO, Ger. Tecnologia | Quinzenal | Sincronizar cronogramas, escalar impedimentos, alinhar dependências com portfólio | Status report integrado |

---

## 4. Caminho de Escalação (Escalation Path)

```
[IMPEDIMENTO DE NEGÓCIO / ESCOPO]
    │
    ├─ 1. Resolver com Product Owner (PO)
    │
    ├─ 2. Se requer mudança de escopo ou impacto no negócio:
    │      Escalar para Gerência Comercial → Gerência de Tecnologia
    │
    ├─ 3. Se requer formalização de Change Request:
    │      CCB (085-PLANO-GESTAO-MUDANCAS) → Sponsor
    │
    └─ 4. Se impacta orçamento:
           Gerência de Finanças → Sponsor (Diretoria de Tecnologia)

[IMPEDIMENTO ORÇAMENTÁRIO / FINANCEIRO]
    │
    ├─ 1. Resolver com Gerência de Finanças + PMO
    │
    ├─ 2. Se requer realocação ou budget adicional:
    │      Escalar para Sponsor (Diretoria de Tecnologia)
    │
    └─ 3. Se impacto crítico no projeto:
           Comitê Executivo extraordinário

[IMPEDIMENTO DE CRONOGRAMA / ENTREGA]
    │
    ├─ 1. Resolver com PMO + PO + Gerência de Tecnologia
    │
    ├─ 2. Se requer repactuação de prazo:
    │      Escalar para Gerência Comercial (impacto no negócio) → Sponsor
    │
    └─ 3. Se impacta portfólio corporativo:
           PMO convoca alinhamento extraordinário com todos os sponsors

[IMPEDIMENTO REGULATÓRIO / LGPD]
    │
    ├─ 1. Gerência de Tecnologia + PO analisam impacto técnico e de negócio
    │
    ├─ 2. Emitem parecer com recomendação de adequação
    │
    └─ 3. Se impacto crítico (risco de sanção):
           PO convoca Comitê Executivo extraordinário (Sponsor + Ger. Comercial + Ger. Finanças)
```

---

## 5. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | 07/08/2026 | Criação inicial a partir do Project Charter (001) — stakeholders de negócio | Time de Negócios / Orquestrador WATERFALL v2.0 |
| 1.1 | 19/08/2026 | Atualização (skill waterfall-business-documents): nomes/contatos marcados como "A definir" (decisão HITL); Matriz RACI reestruturada para as 6 fases do roadmap v3.3 (F5 Execução + F6 Encerramento) | Time de Negócios / skill waterfall-business-documents |
| 1.2 | 19/08/2026 | Correção cirúrgica (review FASE 1): referência de entregas do PO corrigida para (D1-D2), conforme entregas definidas no Charter (001, Seção 4) | Time de Negócios / skill waterfall-business-documents |
| 1.3 | 19/08/2026 | Aprovação humana (P1=SIM, P2/P3/P4=NÃO — atalho OK) — documento congelado em COMPLIANCE | Orquestrador / skill waterfall-business-documents |
