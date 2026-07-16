# Stakeholder Map

- **Projeto:** FBSO Platform — Portal Administrativo SaaS
- **Código:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Versão:** 1.0
- **Atualizado:** 2026-07-13
- **Status:** Aguardando preenchimento de nomes e contatos

---

## 1. Identificação das Partes Interessadas

### 1.1 Patrocinadores Executivos (Sponsors)

| Papel | Nome | Decide sobre | Contato |
|-------|------|--------------|---------|
| Sponsor — Diretoria FBSO.ORG | `<nome>` | Aprovação do Project Charter, liberação de recursos orçamentários, validação do alinhamento estratégico, decisões de escopo que impactem prazo ou orçamento | `<email>` |

### 1.2 Produto e Negócio

| Papel | Nome | Decide sobre | Contato |
|-------|------|--------------|---------|
| Dono do Produto (Product Owner) | `<nome>` | Priorização de funcionalidades, definição de critérios de aceite, trade-offs de escopo vs. prazo, aceitação formal das entregas D1-D7 | `<email>` |
| Analista de Negócios | `<nome>` | Requisitos de negócio, regras de validação (RNs do FEATURES.md), documentação funcional, critérios de aceitação das user stories | `<email>` |
| Coordenador do Projeto | `<nome>` | Cronograma, alocação de recursos, gestão de riscos, comunicação entre stakeholders, escalação de impedimentos | `<email>` |

### 1.3 Time Comercial e de Mercado

| Papel | Nome | Decide sobre | Contato |
|-------|------|--------------|---------|
| Líder Comercial | `<nome>` | Definição do portfólio de planos (nomes, preços, módulos incluídos), estratégia de precificação, seleção de early adopters para validação do portal, feedback de mercado | `<email>` |
| Time de Vendas | `<nome>` | Operação de ativação de clientes no portal, feedback sobre usabilidade do fluxo comercial, identificação de necessidades de evolução do portal | `<email>` |

### 1.4 Time Administrativo e Operações

| Papel | Nome | Decide sobre | Contato |
|-------|------|--------------|---------|
| Líder Administrativo | `<nome>` | Processos de ativação/suspensão de contas, critérios de gestão de status de tenants, validação dos fluxos administrativos do portal, treinamento do time interno | `<email>` |
| Equipe Administrativa | `<nome>` | Operação diária do portal administrativo (gestão de contas, planos, usuários), feedback de usabilidade, identificação de necessidades de melhoria | `<email>` |

### 1.5 Clientes e Usuários Finais

| Papel | Nome | Decide sobre | Contato |
|-------|------|--------------|---------|
| Early Adopter 1 — Cliente parceiro | `<nome>` | Validação do Portal do Cliente (D5), feedback de usabilidade do onboarding, percepção de valor da plataforma | `<email>` |
| Early Adopter 2 — Cliente parceiro | `<nome>` | Validação do Portal do Cliente (D5), feedback sobre cadastro de Unidades de Negócio e Catálogo de Produtos | `<email>` |
| Cliente Final (futuro) | — | Usuário final do portal; consultado sobre experiência de onboarding e usabilidade | — |

---

## 2. Matriz RACI por Entrega do Projeto

**Legenda:** R = Responsible (executa) | A = Accountable (aprova/responde) | C = Consulted (consultado) | I = Informed (informado)

### 2.1 D1 — Portal Administrativo Interno (M2: 15/08/2026)

| Atividade | Diretoria | Coord. Projeto | PO | Analista Neg. | Líder Comercial | Líder Admin | Time Vendas | Early Adopters |
|-----------|-----------|---------------|----|--------------|-----------------|-------------|-------------|----------------|
| Definição de métricas do dashboard | I | C | C | R | C | R | I | — |
| Validação do dashboard | I | I | C | C | I | **A** | C | — |
| Aceitação da entrega D1 | **A** | R | C | C | I | C | I | — |

### 2.2 D2 — Gestão de Contas + D3 — Planos e Assinaturas (M3: 31/08/2026)

| Atividade | Diretoria | Coord. Projeto | PO | Analista Neg. | Líder Comercial | Líder Admin | Time Vendas | Early Adopters |
|-----------|-----------|---------------|----|--------------|-----------------|-------------|-------------|----------------|
| Definição do portfólio de planos (nomes, preços, módulos) | C | I | C | C | **A** | C | C | — |
| Regras de transição de status de tenant | I | I | C | R | C | **A** | I | — |
| Fluxo de ativação e suspensão de contas | I | I | C | R | I | **A** | C | — |
| Aceitação da entrega D2+D3 | **A** | R | C | C | C | C | I | — |

### 2.3 D4 — Usuários e Permissões / RBAC (M4: 15/09/2026)

| Atividade | Diretoria | Coord. Projeto | PO | Analista Neg. | Líder Comercial | Líder Admin | Time Vendas | Early Adopters |
|-----------|-----------|---------------|----|--------------|-----------------|-------------|-------------|----------------|
| Definição da matriz de permissões (papéis × funcionalidades) | I | I | **A** | R | C | C | — | — |
| Validação de isolamento entre Unidades de Negócio | I | I | C | R | — | **A** | — | — |
| Aceitação da entrega D4 | **A** | R | C | C | I | C | — | — |

### 2.4 D5 — Portal do Cliente (M5: 30/09/2026)

| Atividade | Diretoria | Coord. Projeto | PO | Analista Neg. | Líder Comercial | Líder Admin | Time Vendas | Early Adopters |
|-----------|-----------|---------------|----|--------------|-----------------|-------------|-------------|----------------|
| Definição do fluxo de onboarding (passos, textos, mensagens) | I | I | C | R | C | I | — | C |
| Teste e validação do onboarding | — | I | C | C | C | — | — | **A** |
| Validação do App Switcher | I | I | **A** | C | C | — | — | C |
| Aceitação da entrega D5 | **A** | R | C | C | C | I | — | C |

### 2.5 D6 — Unidades de Negócio + D7 — Catálogo de Produtos (M6: 15/10/2026)

| Atividade | Diretoria | Coord. Projeto | PO | Analista Neg. | Líder Comercial | Líder Admin | Time Vendas | Early Adopters |
|-----------|-----------|---------------|----|--------------|-----------------|-------------|-------------|----------------|
| Validação do cadastro de Unidades de Negócio | — | I | C | R | C | I | — | **A** |
| Validação do Catálogo de Produtos | — | I | C | R | C | — | — | **A** |
| Aceitação da entrega D6+D7 | **A** | R | C | C | C | — | — | C |

### 2.6 M7 — Aceite Final (30/10/2026)

| Atividade | Diretoria | Coord. Projeto | PO | Analista Neg. | Líder Comercial | Líder Admin | Time Vendas | Early Adopters |
|-----------|-----------|---------------|----|--------------|-----------------|-------------|-------------|----------------|
| Homologação completa D1-D7 | **A** | R | C | C | C | C | C | C |
| Revisão da Matriz de Riscos | I | R | C | C | C | C | — | — |
| Encerramento formal do projeto | **A** | R | C | I | I | I | I | I |

---

## 3. Canais de Comunicação e Frequência

| Fórum | Participantes | Frequência | Objetivo | Artefato de saída |
|-------|---------------|------------|----------|-------------------|
| **Comitê Executivo do Projeto** | Diretoria, Coord. Projeto, PO | Mensal | Aprovar direcionamento, liberar recursos, revisar KPIs | Dashboard Executivo (MATRIZ-KPI.md) |
| **Reunião de Alinhamento de Produto** | PO, Analista de Negócios, Coord. Projeto | Semanal | Revisar progresso das entregas, destravar impedimentos, refinar backlog | Ata + backlog atualizado |
| **Demo e Validação Administrativa** | PO, Líder Admin, Equipe Admin | Quinzenal | Demonstrar funcionalidades concluídas do portal interno, coletar feedback | Funcionalidades aprovadas + log de feedback |
| **Alinhamento Comercial** | PO, Líder Comercial, Coord. Projeto | Quinzenal | Revisar portfólio de planos, alinhar estratégia de early adopters | Atualização de planos e preços |
| **Sessão com Early Adopters** | PO, Analista Neg., Early Adopters | Quinzenal | Validar Portal do Cliente e onboarding com clientes reais | Feedback registrado + ajustes no backlog |
| **Daily Técnica (Standup)** | Time de Desenvolvimento, QA, Tech Leads | Diária (15 min) | Sincronizar progresso técnico, destravar impedimentos de código | Board de tarefas atualizado |
| **Revisão de Arquitetura** | Arquiteto, Tech Leads, DevOps | Quinzenal | Validar decisões técnicas, revisar ADRs, antecipar riscos de integração | ADRs atualizados, log de decisões técnicas |

---

## 4. Caminho de Escalação (Escalation Path)

```
[IMPEDIMENTO TÉCNICO]
    │
    ├─ 1. Resolver com Tech Lead (Backend ou Frontend)
    │
    ├─ 2. Se requer decisão arquitetural:
    │      Escalar para Arquiteto de Solução → Coordenador do Projeto
    │
    ├─ 3. Se impacta infraestrutura ou segurança:
    │      DevOps → Arquiteto → Coordenador → Diretoria
    │
    └─ 4. Se é incidente crítico (produção fora do ar, vazamento de dados):
           Acionar plano de resposta a incidentes. Coordenador notifica Diretoria imediatamente

[IMPEDIMENTO DE NEGÓCIO]
    │
    ├─ 1. Resolver com Coordenador do Projeto
    │
    ├─ 2. Se requer decisão de produto:
    │      Escalar para Product Owner (PO)
    │
    ├─ 3. Se impacta orçamento ou diretriz estratégica:
    │      Escalar para Diretoria FBSO.ORG
    │
    └─ 4. Se impacta cronograma de módulos futuros:
           Alinhar com roadmap de produto (pós-projeto)

[IMPEDIMENTO OPERACIONAL]
    │
    ├─ 1. Resolver com Líder Administrativo (time interno) ou PO (portal do cliente)
    │
    ├─ 2. Se requer mudança de processo ou regra de negócio:
    │      Escalar para PO → Analista de Negócios
    │
    └─ 3. Se impacta adoção ou satisfação:
           Escalar para Coordenador → Diretoria

[IMPEDIMENTO COMERCIAL]
    │
    ├─ 1. Resolver com Líder Comercial
    │
    ├─ 2. Se requer alteração de portfólio de planos:
    │      Líder Comercial → PO → Diretoria
    │
    └─ 3. Se impacta relacionamento com early adopters:
           Coordenador + Líder Comercial definem plano de ação
```

---

## 5. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | 2026-07-13 | Criação inicial: stakeholders de negócio, matriz RACI por entrega (D1-D7, M7), canais de comunicação, escalation path. Nomes e contatos a preencher. | Time de Negócios |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: stakeholder-analysis, agile-ba-practices.*
🔍 *Revisado pelo skill caveman-review em 15/07/2026. Ajustes aplicados: escalation path técnico adicionado (DevOps→Arquiteto→Coordenador→Diretoria), canais de comunicação técnica (Daily + Revisão de Arquitetura).*
