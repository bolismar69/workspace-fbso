# Stakeholder Map

- **Projeto:** Portal Corporativo de Gestão Tributária — Autonomia do Time de Finanças na Administração de Impostos
- **Código:** PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS
- **Programa Pai:** [PRJ-FIN-2026-0001](../PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/)
- **Versão:** 1.0
- **Atualizado:** 2026-07-08

---

## 1. Identificação das Partes Interessadas

### 1.1 Patrocinadores Executivos (Sponsors)

| Papel | Nome | Decide sobre | Contato |
|-------|------|--------------|---------|
| Sponsor — Diretor Financeiro (CFO) | `<nome>` | Liberação de recursos orçamentários, aprovação de diretrizes de governança fiscal, validação do alinhamento estratégico com o programa PRJ-FIN-2026-0001 | `<email>` |

### 1.2 Governança Fiscal e Controles Internos

| Papel | Nome | Decide sobre | Contato |
|-------|------|--------------|---------|
| Comitê Fiscal e Jurídico de Negócios | `<nome>` | Regras de validação de negócio incorporadas ao portal (conflitos de vigência, transição de regimes, integridade entre tributos), validação das alíquotas carregadas na base, interpretação de mudanças regulatórias aplicáveis | `<email>` |
| Gerência de Controladoria e Compliance | `<nome>` | Requisitos de trilha de auditoria, aprovação dos relatórios de governança, critérios de conformidade com Lei das S.A. e framework COSO | `<email>` |
| Auditoria Interna | `<nome>` | Validação da integridade e imutabilidade da trilha de auditoria, escopo e frequência das auditorias amostrais | `<email>` |

### 1.3 Time de Finanças (Usuários Finais)

| Papel | Nome | Decide sobre | Contato |
|-------|------|--------------|---------|
| Gerente de Fiscal/Controller | `<nome>` | Aprovação de alterações de alto impacto (fluxo de duas etapas), priorização de funcionalidades do portal, validação dos critérios de aceite de negócio (UAT) | `<email>` |
| Analista Fiscal Sênior | `<nome>` | Operação diária do portal (cadastro e manutenção de alíquotas), elicitação de regras de validação tácitas, treinamento de novos usuários | `<email>` |
| Analistas Fiscais (equipe) | `<nome>` | Operação diária do portal, feedback de usabilidade, identificação de necessidades de evolução | `<email>` |
| Contadores | `<nome>` | Consulta de tabelas fiscais para conciliação contábil, validação de alíquotas para fechamento mensal | `<email>` |

### 1.4 Lideranças de Negócio Impactadas

| Papel | Nome | Decide sobre | Contato |
|-------|------|--------------|---------|
| Gerência de Tesouraria | `<nome>` | Impacto das alíquotas geridas no portal sobre projeções de fluxo de caixa (split payment), necessidades de informação para conciliação bancária | `<email>` |
| Lideranças de Canais e Mercado (Diretores/Gerentes Comerciais) | `<nome>` | Feedback sobre impacto das alíquotas cadastradas nos preços praticados nos canais de venda, rapidez necessária na atualização de alíquotas que afetam competitividade | `<email>` |

### 1.5 Execução e Governança do Projeto

| Papel | Nome | Decide sobre | Contato |
|-------|------|--------------|---------|
| Product Owner (PO) — Portal de Gestão Tributária | `<nome>` | Priorização de funcionalidades, definição de critérios de aceite, trade-offs de escopo vs. prazo, aceitação formal das entregas | `<email>` |
| PMO Corporativo (Programa PRJ-FIN-2026-0001) | `<nome>` | Alinhamento do cronograma com o programa pai, arbitragem de conflitos de priorização, escalação de impedimentos aos sponsors | `<email>` |

---

## 2. Matriz RACI por Fase do Projeto

**Legenda:** R = Responsible (executa) | A = Accountable (aprova/responde) | C = Consulted (consultado) | I = Informed (informado)

### 2.1 Fase 0 — Fundamentação (Mapeamento e Prototipação)

| Atividade | CFO | Comitê Fiscal | Controladoria | Gerente Fiscal | Analistas Fiscais | Tesouraria | Lideranças Comerciais | PO | PMO |
|-----------|-----|---------------|---------------|----------------|--------------------|------------|----------------------|----|-----|
| Mapeamento das tabelas fiscais existentes | — | C | C | R | R | — | — | **A** | I |
| Definição das regras de validação de negócio | — | **A** | C | R | C | — | — | I | — |
| Levantamento de necessidades do time fiscal | — | — | — | C | R | — | — | **A** | — |
| Prototipação das telas do portal | — | C | — | C | C | — | — | **A** | I |
| Definição de perfis de acesso e segregação | — | C | **A** | C | — | — | — | R | I |

### 2.2 Fase 1 — Portal: Gestão Básica de Alíquotas (Entrega 1)

| Atividade | CFO | Comitê Fiscal | Controladoria | Gerente Fiscal | Analistas Fiscais | Tesouraria | Lideranças Comerciais | PO | PMO |
|-----------|-----|---------------|---------------|----------------|--------------------|------------|----------------------|----|-----|
| Carga inicial das tabelas vigentes | — | C | C | R | R | — | — | **A** | I |
| Validação das alíquotas carregadas | — | **A** | C | R | C | — | — | I | — |
| Treinamento do time fiscal no portal | — | — | — | R | R | — | — | **A** | — |
| Aceitação do Módulo 1 (Painel de Alíquotas) | — | C | — | **A** | R | — | — | C | I |
| Aceitação do Módulo 2 (Cadastro de Alíquotas) | — | C | — | **A** | R | — | — | C | I |
| Aceitação do Módulo 3 (Classificações e Regimes) | — | **A** | C | R | C | — | — | C | I |

### 2.3 Fase 2 — Governança e Auditoria Fiscal (Entrega 2)

| Atividade | CFO | Comitê Fiscal | Controladoria | Gerente Fiscal | Analistas Fiscais | Tesouraria | Lideranças Comerciais | PO | PMO |
|-----------|-----|---------------|---------------|----------------|--------------------|------------|----------------------|----|-----|
| Aceitação do Módulo 6 (Administração de Acessos e Perfis) | — | — | **A** | C | — | — | — | R | I |
| Validação da trilha de auditoria | — | C | **A** | R | — | — | — | C | I |
| Aceitação do Módulo 4 (Linha do Tempo e Auditoria) | — | C | **A** | R | C | — | — | C | I |
| Validação cruzada portal × base ativa (reconciliação) | — | C | **A** | R | C | — | — | I | — |
| Descontinuação do acesso à ferramenta legada | — | — | C | **A** | R | — | — | I | I |

### 2.4 Fase 3 — Operações Fiscais em Escala (Entrega 3)

| Atividade | CFO | Comitê Fiscal | Controladoria | Gerente Fiscal | Analistas Fiscais | Tesouraria | Lideranças Comerciais | PO | PMO |
|-----------|-----|---------------|---------------|----------------|--------------------|------------|----------------------|----|-----|
| Definição de patamares de materialidade (fluxo de aprovação) | C | **A** | C | R | — | C | — | I | — |
| Aceitação dos fluxos de aprovação (Feature 03.1) | — | C | **A** | R | C | — | — | C | I |
| Teste de carga em lote (simulação 5.570 municípios) | — | C | — | R | R | — | — | **A** | I |
| Aceitação do Módulo 5 (Importação/Exportação em Lote) | — | C | — | **A** | R | — | — | C | I |

### 2.5 Fase 4 — Inteligência Fiscal e Analytics (Entrega 4)

| Atividade | CFO | Comitê Fiscal | Controladoria | Gerente Fiscal | Analistas Fiscais | Tesouraria | Lideranças Comerciais | PO | PMO |
|-----------|-----|---------------|---------------|----------------|--------------------|------------|----------------------|----|-----|
| Aprovação dos relatórios de governança | C | C | **A** | R | — | — | — | I | — |
| Validação dos dashboards de KPIs fiscais | **A** | C | C | R | — | C | C | I | — |
| Teste do cenário híbrido completo | — | **A** | C | R | C | — | — | C | I |
| Treinamento final e handover para operação autônoma | — | — | — | R | R | — | — | **A** | I |

### 2.6 Período Híbrido (2029–2032: Suporte Continuado)

| Atividade | CFO | Comitê Fiscal | Controladoria | Gerente Fiscal | Analistas Fiscais | Tesouraria | Lideranças Comerciais | PO | PMO |
|-----------|-----|---------------|---------------|----------------|--------------------|------------|----------------------|----|-----|
| Gestão dual de tabelas (antigo + novo regime) | — | C | C | **A** | R | — | C | I | I |
| Desativação progressiva de tributos extintos | — | **A** | C | R | R | — | — | I | I |
| Monitoramento de publicações do Comitê Gestor do IBS | I | **A** | C | R | — | — | — | I | I |

---

## 3. Canais de Comunicação e Frequência

| Fórum | Participantes | Frequência | Objetivo | Artefato de saída |
|-------|---------------|------------|----------|-------------------|
| **Comitê Executivo do Projeto** | CFO, PMO, Comitê Fiscal, PO | Mensal | Aprovar direcionamento, liberar recursos, revisar KPIs | Dashboard Executivo (MATRIZ-KPI.md) |
| **Reunião de Governança Fiscal** | Comitê Fiscal, Controladoria, Gerente Fiscal, PO | Quinzenal | Validar regras de negócio do portal, interpretar mudanças regulatórias, revisar relatórios de governança | Ata + atualização de regras de validação |
| **Demo e Validação com Time Fiscal** | PO, Gerente Fiscal, Analistas Fiscais | Quinzenal (por ciclo de desenvolvimento) | Demonstrar funcionalidades concluídas, coletar feedback de usabilidade, validar critérios de aceite | Funcionalidades aprovadas + log de feedback |
| **Reconciliação Fiscal** | Controladoria, Gerente Fiscal, Analistas | Mensal | Cruzar alíquotas do portal com base ativa e publicações oficiais, investigar divergências | Relatório de reconciliação (KPI G3) |
| **Alinhamento com Programa Pai** | PMO (PRJ-FIN-2026-0001), PO | Quinzenal | Sincronizar cronogramas, escalar impedimentos, alinhar dependências | Status report integrado |

---

## 4. Caminho de Escalação (Escalation Path)

```
[IMPEDIMENTO DE NEGÓCIO]
    │
    ├─ 1. Resolver com Product Owner (PO)
    │
    ├─ 2. Se requer decisão fiscal:
    │      Escalar para Comitê Fiscal → Controladoria
    │
    ├─ 3. Se impacta orçamento ou diretriz estratégica:
    │      Escalar para PMO → CFO
    │
    └─ 4. Se impacta programa pai (PRJ-FIN-2026-0001):
           PMO corporativo convoca alinhamento extraordinário

[IMPEDIMENTO OPERACIONAL]
    │
    ├─ 1. Resolver com Gerente Fiscal
    │
    ├─ 2. Se requer mudança de processo ou perfil de acesso:
    │      Escalar para Controladoria → Comitê Fiscal
    │
    └─ 3. Se impacta adoção ou satisfação do time:
           Escalar para PO → PMO → CFO

[IMPEDIMENTO REGULATÓRIO]
    │
    ├─ 1. Comitê Fiscal analisa impacto da mudança legislativa
    │
    ├─ 2. Comitê Fiscal emite parecer com recomendação
    │
    └─ 3. Se impacto crítico (alíquotas, prazos, obrigações):
           PO convoca Comitê Executivo extraordinário
```

---

## 5. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | 2026-07-08 | Criação inicial: stakeholders de negócio, matriz RACI por fase, canais de comunicação, escalation path | Time de Negócios |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: stakeholder-analysis, agile-ba-practices.*
