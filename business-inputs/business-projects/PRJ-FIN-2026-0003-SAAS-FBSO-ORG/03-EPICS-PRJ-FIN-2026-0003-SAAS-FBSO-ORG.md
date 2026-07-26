# Épicos do Projeto: FBSO Platform — Portal Administrativo SaaS

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 2.0 — Revisada com Estrutura Modular v4.0 (26/07/2026) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` v1.1 e `01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` v1.1 |
| **Status** | Em Revisão / Aguardando Validação |

---

## Visão Geral dos Épicos

| ID | Épico | Objetivo de Negócio | Func. | Prioridade | Data-Alvo |
|----|-------|---------------------|-------|------------|-----------|
| **EP-0001** | [Portal Administrativo Interno](epics/EP-0001-portal-administrativo-interno.md) | Time FBSO.ORG gerencia a operação SaaS com visibilidade em tempo real | 3 | Must Have | **15/08/2026** |
| **EP-0002** | [Gestão de Clientes e Assinaturas](epics/EP-0002-gestao-de-clientes-e-assinaturas.md) | Estruturar a operação comercial: contas, planos e ciclo de vida do cliente | 5 | Must Have | **31/08/2026** |
| **EP-0003** | [Governança de Acessos e Permissões](epics/EP-0003-governanca-de-acessos-e-permissoes.md) | Garantir segurança e isolamento de dados entre clientes e entre filiais | 4 | Must Have | **15/09/2026** |
| **EP-0004** | [Experiência do Cliente e Autoatendimento](epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) | Cliente realiza onboarding, gerencia seus dados e navega entre módulos | 6 | Must Have | **15/10/2026** |

> 📄 **Detalhamento completo** de cada épico disponível na pasta [`epics/`](epics/):
> - [EP-0001 — Portal Administrativo Interno](epics/EP-0001-portal-administrativo-interno.md)
> - [EP-0002 — Gestão de Clientes e Assinaturas](epics/EP-0002-gestao-de-clientes-e-assinaturas.md)
> - [EP-0003 — Governança de Acessos e Permissões](epics/EP-0003-governanca-de-acessos-e-permissoes.md)
> - [EP-0004 — Experiência do Cliente e Autoatendimento](epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md)

### Cronograma de Épicos

| Data-Alvo | Marco | Épicos | Funcionalidades |
|-----------|-------|--------|----------------|
| 15/08/2026 | M2 | EP-0001 — Portal Administrativo Interno | F01-01, F01-02, F01-03 |
| 31/08/2026 | M3 | EP-0002 — Gestão de Clientes e Assinaturas | F02-01 a F02-05 |
| 15/09/2026 | M4 | EP-0003 — Governança de Acessos e Permissões | F03-01 a F03-04 |
| 30/09/2026 | M5 | EP-0004 — Portal do Cliente e Onboarding | F04-01 a F04-04 |
| 15/10/2026 | M6 | EP-0004 — Unidades de Negócio e Catálogo | F04-05, F04-06 |
| 30/10/2026 | M7 | Aceite Final — homologação completa | Todas |

### Mapa de Dependências entre Épicos

```
24/07    15/08     31/08     15/09     30/09     15/10     30/10
  │────────│─────────│─────────│─────────│─────────│─────────│
  M1       M2        M3        M4        M5        M6        M7
  ▼        ▼         ▼         ▼         ▼         ▼         ▼
Kickoff  EP-0001   EP-0002   EP-0003   EP-0004a  EP-0004b  Aceite
         Portal    Clientes  Acessos   Portal    BUs+     Final
         Admin     Planos    Permiss   Cliente   Catálogo
```

> **Nota:** EP-0001 é pré-requisito para EP-0002 (dependência sequencial conforme milestones M2→M3 do Project Charter). EP-0003 depende da existência de Tenants e Usuários. EP-0004 depende de toda a camada de governança (EP-0003). EP-0004 é entregue em duas etapas: Portal do Cliente e Onboarding (M5) e Unidades de Negócio + Catálogo (M6). D6 e D7 podem parcialmente paralelizar com validação de D5.

---

## Matriz de Rastreabilidade BRD → Épicos

Todo requisito funcional do BRD deve estar coberto por pelo menos um épico. A matriz abaixo audita essa cobertura e vincula cada BR às jornadas de usuário que o realizam.

| BRD | Requisito Funcional | Épico | Jornada(s) que Realizam |
|:---|:---|:---|:---|
| **BR-01** | Dashboard Administrativo | [**EP-0001** — Portal Administrativo Interno](epics/EP-0001-portal-administrativo-interno.md) | J1: Acompanhamento diário da operação · J2: Análise de crescimento por plano |
| **BR-02** | Ativação e Gestão de Contas | [**EP-0002** — Gestão de Clientes e Assinaturas](epics/EP-0002-gestao-de-clientes-e-assinaturas.md) | J1: Ativação de novo cliente · J2: Suspensão por inadimplência |
| **BR-03** | Configuração de Planos Comerciais | [**EP-0002** — Gestão de Clientes e Assinaturas](epics/EP-0002-gestao-de-clientes-e-assinaturas.md) | Requisitos Funcionais §5 (Cadastro de planos comerciais) |
| **BR-04** | Vinculação de Assinaturas | [**EP-0002** — Gestão de Clientes e Assinaturas](epics/EP-0002-gestao-de-clientes-e-assinaturas.md) | J1: Ativação de novo cliente · J3: Upgrade de plano |
| **BR-05** | Gestão de Usuários e Permissões | [**EP-0003** — Governança de Acessos e Permissões](epics/EP-0003-governanca-de-acessos-e-permissoes.md) | J1: Convite de novo usuário · J2: Restrição de acesso entre filiais · J3: Revogação de acesso |
| **BR-06** | Portal do Cliente com Autenticação | [**EP-0004** — Experiência do Cliente](epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) | J1: Primeiro acesso e onboarding |
| **BR-07** | Onboarding Guiado de Primeiro Acesso | [**EP-0004** — Experiência do Cliente](epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) | J1: Primeiro acesso e onboarding |
| **BR-08** | App Switcher (Seletor de Aplicativos) | [**EP-0004** — Experiência do Cliente](epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) | J4: Navegação com App Switcher (visão futura) |
| **BR-09** | Cadastro de Unidades de Negócio | [**EP-0004** — Experiência do Cliente](epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) | J2: Cadastro de filiais (Unidades de Negócio) |
| **BR-10** | Catálogo de Produtos/Serviços | [**EP-0004** — Experiência do Cliente](epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) | J3: Cadastro de portfólio de produtos |

### Auditoria de Cobertura

| Métrica | Resultado |
|---|---|
| BRs cobertos por épicos | 10/10 (100%) |
| BRs cobertos por jornadas | 9/10 — BR-03 coberto via requisitos funcionais (sem jornada narrativa explícita) |
| Épicos com rastreabilidade BRD | 4/4 (100%) |
| Épicos com BRs novos (não mapeados) | 0 — Nenhum escopo extra detectado |
| Requisitos órfãos (sem épico) | 0 |

> 💡 **BR-03** (Configuração de Planos Comerciais) não possui uma jornada narrativa explícita porque é uma funcionalidade de configuração realizada pelo Gestor de Produto — não é uma jornada de "cliente" ou "administrador". Está coberto pelos Requisitos Funcionais da seção §5 do [EP-0002](epics/EP-0002-gestao-de-clientes-e-assinaturas.md). Se desejado, uma jornada "Gestor de Produto cadastra novo plano" pode ser adicionada futuramente.

---

## Sumário de Cobertura do Escopo

| Entrega do Project Charter | Épico(s) que cobrem |
|---------------------------|---------------------|
| D1 — Portal Administrativo Interno | [EP-0001](epics/EP-0001-portal-administrativo-interno.md) |
| D2 — Módulo de Gestão de Contas | [EP-0002](epics/EP-0002-gestao-de-clientes-e-assinaturas.md) |
| D3 — Módulo de Planos e Assinaturas | [EP-0002](epics/EP-0002-gestao-de-clientes-e-assinaturas.md) |
| D4 — Módulo de Usuários e Permissões | [EP-0003](epics/EP-0003-governanca-de-acessos-e-permissoes.md) |
| D5 — Portal do Cliente | [EP-0004](epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) |
| D6 — Cadastro de Unidades de Negócio | [EP-0004](epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) |
| D7 — Catálogo de Produtos/Serviços | [EP-0004](epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) |

---

> **Este documento é um índice resumido.** O detalhamento completo de cada épico — incluindo objetivo, personas, jornadas, requisitos de negócio, métricas de sucesso, escopo excluído e valor de negócio — está nos arquivos individuais da pasta [`epics/`](epics/).

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: breakdown-epic-pm, agile-ba-practices. Estrutura modular v4.0.*

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE ÉPICOS]
