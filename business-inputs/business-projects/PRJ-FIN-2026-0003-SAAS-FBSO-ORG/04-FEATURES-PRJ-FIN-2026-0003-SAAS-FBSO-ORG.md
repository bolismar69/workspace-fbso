# Funcionalidades do Projeto: FBSO Platform — Portal Administrativo SaaS

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 2.0 — Revisada com Estrutura Modular v2.0 (26/07/2026) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` v2.0 |
| **Status** | Em Revisão / Aguardando Validação |

---

## Visão Geral das Funcionalidades

| ID | Funcionalidade | Épico | Prioridade | User Stories | Data-Alvo |
|----|---------------|-------|------------|-------------|-----------|
| **FEAT-EP-0001-0001** | [Dashboard de Métricas Operacionais](features/FEAT-EP-0001-0001-dashboard-metricas-operacionais.md) | EP-0001 | Must Have | 3 | **15/08/2026** |
| **FEAT-EP-0001-0002** | [Visão de Contas com Filtros](features/FEAT-EP-0001-0002-visao-contas-com-filtros.md) | EP-0001 | Must Have | 2 | **15/08/2026** |
| **FEAT-EP-0001-0003** | [Alertas e Indicadores de Atenção](features/FEAT-EP-0001-0003-alertas-indicadores-atencao.md) | EP-0001 | Should Have | 2 | **15/08/2026** |
| **FEAT-EP-0002-0001** | [Cadastro e Ativação de Contas de Clientes](features/FEAT-EP-0002-0001-cadastro-ativacao-contas-clientes.md) | EP-0002 | Must Have | 4 | **31/08/2026** |
| **FEAT-EP-0002-0002** | [Gestão de Status do Tenant](features/FEAT-EP-0002-0002-gestao-status-tenant.md) | EP-0002 | Must Have | 3 | **31/08/2026** |
| **FEAT-EP-0002-0003** | [Configuração de Planos Comerciais](features/FEAT-EP-0002-0003-configuracao-planos-comerciais.md) | EP-0002 | Must Have | 4 | **31/08/2026** |
| **FEAT-EP-0002-0004** | [Vinculação e Gestão de Assinaturas](features/FEAT-EP-0002-0004-vinculacao-gestao-assinaturas.md) | EP-0002 | Must Have | 3 | **31/08/2026** |
| **FEAT-EP-0002-0005** | [Histórico de Auditoria Administrativa](features/FEAT-EP-0002-0005-historico-auditoria-administrativa.md) | EP-0002 | Must Have | 2 | **31/08/2026** |
| **FEAT-EP-0003-0001** | [Cadastro e Convite de Usuários](features/FEAT-EP-0003-0001-cadastro-convite-usuarios.md) | EP-0003 | Must Have | 6 | **15/09/2026** |
| **FEAT-EP-0003-0002** | [Definição de Papéis e Permissões (RBAC)](features/FEAT-EP-0003-0002-definicao-papeis-permissoes-rbac.md) | EP-0003 | Must Have | 4 | **15/09/2026** |
| **FEAT-EP-0003-0003** | [Vinculação Usuário × Unidade × Módulo](features/FEAT-EP-0003-0003-vinculacao-usuario-unidade-modulo.md) | EP-0003 | Must Have | 3 | **15/09/2026** |
| **FEAT-EP-0003-0004** | [Controle de Visibilidade de Menus e Ações](features/FEAT-EP-0003-0004-controle-visibilidade-menus-acoes.md) | EP-0003 | Must Have | 3 | **15/09/2026** |
| **FEAT-EP-0004-0001** | [Autenticação e Recuperação de Senha](features/FEAT-EP-0004-0001-autenticacao-recuperacao-senha.md) | EP-0004 | Must Have | 3 | **30/09/2026** |
| **FEAT-EP-0004-0002** | [Onboarding Guiado de Primeiro Acesso](features/FEAT-EP-0004-0002-onboarding-guiado-primeiro-acesso.md) | EP-0004 | Must Have | 5 | **30/09/2026** |
| **FEAT-EP-0004-0003** | [Dashboard do Cliente](features/FEAT-EP-0004-0003-dashboard-cliente.md) | EP-0004 | Should Have | 3 | **30/09/2026** |
| **FEAT-EP-0004-0004** | [App Switcher (Seletor de Módulos)](features/FEAT-EP-0004-0004-app-switcher-seletor-modulos.md) | EP-0004 | Must Have | 3 | **30/09/2026** |
| **FEAT-EP-0004-0005** | [Gestão de Unidades de Negócio](features/FEAT-EP-0004-0005-gestao-unidades-negocio.md) | EP-0004 | Must Have | 5 | **15/10/2026** |
| **FEAT-EP-0004-0006** | [Catálogo de Produtos e Serviços](features/FEAT-EP-0004-0006-catalogo-produtos-servicos.md) | EP-0004 | Must Have | 4 | **15/10/2026** |

**Total: 18 funcionalidades | 62 user stories**

> 📄 **Detalhamento completo** de cada feature — incluindo objetivo de negócio, user stories, critérios de aceitação e regras de negócio — está nos arquivos individuais da pasta [`features/`](features/).

### Mapeamento de Numeração: Antigo → Novo

| Antigo | Novo | Épico |
|--------|------|-------|
| F01-01 | FEAT-EP-0001-0001 | EP-0001 |
| F01-02 | FEAT-EP-0001-0002 | EP-0001 |
| F01-03 | FEAT-EP-0001-0003 | EP-0001 |
| F02-01 | FEAT-EP-0002-0001 | EP-0002 |
| F02-02 | FEAT-EP-0002-0002 | EP-0002 |
| F02-03 | FEAT-EP-0002-0003 | EP-0002 |
| F02-04 | FEAT-EP-0002-0004 | EP-0002 |
| F02-05 | FEAT-EP-0002-0005 | EP-0002 |
| F03-01 | FEAT-EP-0003-0001 | EP-0003 |
| F03-02 | FEAT-EP-0003-0002 | EP-0003 |
| F03-03 | FEAT-EP-0003-0003 | EP-0003 |
| F03-04 | FEAT-EP-0003-0004 | EP-0003 |
| F04-01 | FEAT-EP-0004-0001 | EP-0004 |
| F04-02 | FEAT-EP-0004-0002 | EP-0004 |
| F04-03 | FEAT-EP-0004-0003 | EP-0004 |
| F04-04 | FEAT-EP-0004-0004 | EP-0004 |
| F04-05 | FEAT-EP-0004-0005 | EP-0004 |
| F04-06 | FEAT-EP-0004-0006 | EP-0004 |

### Cronograma de Entregas por Funcionalidade

| Data-Alvo | Marco | Épico | Funcionalidades |
|-----------|-------|-------|----------------|
| **15/08/2026** | M2 | EP-0001 | FEAT-EP-0001-0001, FEAT-EP-0001-0002, FEAT-EP-0001-0003 |
| **31/08/2026** | M3 | EP-0002 | FEAT-EP-0002-0001, FEAT-EP-0002-0002, FEAT-EP-0002-0003, FEAT-EP-0002-0004, FEAT-EP-0002-0005 |
| **15/09/2026** | M4 | EP-0003 | FEAT-EP-0003-0001, FEAT-EP-0003-0002, FEAT-EP-0003-0003, FEAT-EP-0003-0004 |
| **30/09/2026** | M5 | EP-0004a | FEAT-EP-0004-0001, FEAT-EP-0004-0002, FEAT-EP-0004-0003, FEAT-EP-0004-0004 |
| **15/10/2026** | M6 | EP-0004b | FEAT-EP-0004-0005, FEAT-EP-0004-0006 |
| **30/10/2026** | M7 | Todos | Homologação final (D1-D7) |

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Features

| BRD | Requisito Funcional | Épico/Jornada | Feature |
|:---|:---|:---|:---|
| **BR-01** | Dashboard Administrativo | EP-0001 / J1: Acompanhamento diário · J2: Análise de crescimento | [FEAT-EP-0001-0001](features/FEAT-EP-0001-0001-dashboard-metricas-operacionais.md), [FEAT-EP-0001-0002](features/FEAT-EP-0001-0002-visao-contas-com-filtros.md), [FEAT-EP-0001-0003](features/FEAT-EP-0001-0003-alertas-indicadores-atencao.md) |
| **BR-02** | Ativação e Gestão de Contas | EP-0002 / J1: Ativação de novo cliente · J2: Suspensão | [FEAT-EP-0002-0001](features/FEAT-EP-0002-0001-cadastro-ativacao-contas-clientes.md), [FEAT-EP-0002-0002](features/FEAT-EP-0002-0002-gestao-status-tenant.md), [FEAT-EP-0002-0005](features/FEAT-EP-0002-0005-historico-auditoria-administrativa.md) |
| **BR-03** | Configuração de Planos Comerciais | EP-0002 / Requisitos Funcionais §5 | [FEAT-EP-0002-0003](features/FEAT-EP-0002-0003-configuracao-planos-comerciais.md) |
| **BR-04** | Vinculação de Assinaturas | EP-0002 / J1: Ativação · J3: Upgrade de plano | [FEAT-EP-0002-0004](features/FEAT-EP-0002-0004-vinculacao-gestao-assinaturas.md) |
| **BR-05** | Gestão de Usuários e Permissões | EP-0003 / J1: Convite · J2: Restrição · J3: Revogação | [FEAT-EP-0003-0001](features/FEAT-EP-0003-0001-cadastro-convite-usuarios.md), [FEAT-EP-0003-0002](features/FEAT-EP-0003-0002-definicao-papeis-permissoes-rbac.md), [FEAT-EP-0003-0003](features/FEAT-EP-0003-0003-vinculacao-usuario-unidade-modulo.md), [FEAT-EP-0003-0004](features/FEAT-EP-0003-0004-controle-visibilidade-menus-acoes.md) |
| **BR-06** | Portal do Cliente com Autenticação | EP-0004 / J1: Primeiro acesso e onboarding | [FEAT-EP-0004-0001](features/FEAT-EP-0004-0001-autenticacao-recuperacao-senha.md), [FEAT-EP-0004-0002](features/FEAT-EP-0004-0002-onboarding-guiado-primeiro-acesso.md), [FEAT-EP-0004-0003](features/FEAT-EP-0004-0003-dashboard-cliente.md) |
| **BR-07** | Onboarding Guiado de Primeiro Acesso | EP-0004 / J1: Primeiro acesso e onboarding | [FEAT-EP-0004-0002](features/FEAT-EP-0004-0002-onboarding-guiado-primeiro-acesso.md) |
| **BR-08** | App Switcher (Seletor de Aplicativos) | EP-0004 / J4: Navegação com App Switcher | [FEAT-EP-0004-0004](features/FEAT-EP-0004-0004-app-switcher-seletor-modulos.md) |
| **BR-09** | Cadastro de Unidades de Negócio | EP-0004 / J2: Cadastro de filiais | [FEAT-EP-0004-0005](features/FEAT-EP-0004-0005-gestao-unidades-negocio.md) |
| **BR-10** | Catálogo de Produtos/Serviços | EP-0004 / J3: Cadastro de portfólio | [FEAT-EP-0004-0006](features/FEAT-EP-0004-0006-catalogo-produtos-servicos.md) |

---

## Matriz de Cobertura: Entregas do Project Charter × Features

| Entrega (Project Charter) | Funcionalidades Relacionadas |
|---------------------------|------------------------------|
| D1 — Portal Administrativo Interno | [FEAT-EP-0001-0001](features/FEAT-EP-0001-0001-dashboard-metricas-operacionais.md), [FEAT-EP-0001-0002](features/FEAT-EP-0001-0002-visao-contas-com-filtros.md), [FEAT-EP-0001-0003](features/FEAT-EP-0001-0003-alertas-indicadores-atencao.md) |
| D2 — Módulo de Gestão de Contas | [FEAT-EP-0002-0001](features/FEAT-EP-0002-0001-cadastro-ativacao-contas-clientes.md), [FEAT-EP-0002-0002](features/FEAT-EP-0002-0002-gestao-status-tenant.md), [FEAT-EP-0002-0005](features/FEAT-EP-0002-0005-historico-auditoria-administrativa.md) |
| D3 — Módulo de Planos e Assinaturas | [FEAT-EP-0002-0003](features/FEAT-EP-0002-0003-configuracao-planos-comerciais.md), [FEAT-EP-0002-0004](features/FEAT-EP-0002-0004-vinculacao-gestao-assinaturas.md) |
| D4 — Módulo de Usuários e Permissões | [FEAT-EP-0003-0001](features/FEAT-EP-0003-0001-cadastro-convite-usuarios.md), [FEAT-EP-0003-0002](features/FEAT-EP-0003-0002-definicao-papeis-permissoes-rbac.md), [FEAT-EP-0003-0003](features/FEAT-EP-0003-0003-vinculacao-usuario-unidade-modulo.md), [FEAT-EP-0003-0004](features/FEAT-EP-0003-0004-controle-visibilidade-menus-acoes.md) |
| D5 — Portal do Cliente | [FEAT-EP-0004-0001](features/FEAT-EP-0004-0001-autenticacao-recuperacao-senha.md), [FEAT-EP-0004-0002](features/FEAT-EP-0004-0002-onboarding-guiado-primeiro-acesso.md), [FEAT-EP-0004-0003](features/FEAT-EP-0004-0003-dashboard-cliente.md), [FEAT-EP-0004-0004](features/FEAT-EP-0004-0004-app-switcher-seletor-modulos.md) |
| D6 — Cadastro de Unidades de Negócio | [FEAT-EP-0004-0005](features/FEAT-EP-0004-0005-gestao-unidades-negocio.md) |
| D7 — Catálogo de Produtos/Serviços | [FEAT-EP-0004-0006](features/FEAT-EP-0004-0006-catalogo-produtos-servicos.md) |

```
24/07    15/08     31/08     15/09     30/09     15/10     30/10
  │────────│─────────│─────────│─────────│─────────│─────────│
  M1       M2        M3        M4        M5        M6        M7
  ▼        ▼         ▼         ▼         ▼         ▼         ▼
Kickoff  EP-0001   EP-0002   EP-0003   EP-0004a  EP-0004b  Aceite
         Portal    Clientes  Acessos   Portal    BUs+     Final
         Admin     Planos    Permiss   Cliente   Catálogo
         EP-0001    EP-0002    EP-0003    EP-0004    EP-0004
          -0001     -0001     -0001     -0001     -0005
          -0002     -0002     -0002     -0002     -0006
          -0003     -0003     -0003     -0003
                    -0004     -0004     -0004
                    -0005
```

---

## Matriz de Priorização (MoSCoW)

| Prioridade | Funcionalidades | Quantidade |
|-----------|----------------|------------|
| **Must Have** | FEAT-EP-0001-0001, FEAT-EP-0001-0002, FEAT-EP-0002-0001, FEAT-EP-0002-0002, FEAT-EP-0002-0003, FEAT-EP-0002-0004, FEAT-EP-0002-0005, FEAT-EP-0003-0001, FEAT-EP-0003-0002, FEAT-EP-0003-0003, FEAT-EP-0003-0004, FEAT-EP-0004-0001, FEAT-EP-0004-0002, FEAT-EP-0004-0004, FEAT-EP-0004-0005, FEAT-EP-0004-0006 | 16 |
| **Should Have** | FEAT-EP-0001-0003, FEAT-EP-0004-0003 | 2 |
| **Could Have** | — | 0 |
| **Won't Have (esta fase)** | Funcionalidades dos módulos Tributali-Engine e Storekeeper Portal | — |

---

> **Este documento é um índice resumido.** O detalhamento completo de cada feature — incluindo objetivo de negócio, user stories, critérios de aceitação e regras de negócio — está nos arquivos individuais da pasta [`features/`](features/).

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: agile-ba-practices, acceptance-criteria, breakdown-feature-prd. Estrutura modular v2.0.*

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
