# DETAIL-LEVEL-PRD — Product Requirements Document (Detail-Level)

- **Autor:** Negócio (PO/PM/Funcional)
- **Audiência:** Time de TI
- **Data:** 31/07/2026
- **Documentos Referenciados:** [Project Charter](../01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [BRD](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [Épicos](../03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [Features](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [User Stories](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [Glossário](../GLOSSARY.md), [Matriz KPI](../MATRIZ-KPI.md), [Stakeholder Map](../STAKEHOLDER-MAP.md), [DoD](../DEFINITION_OF_DONE.md)

---

## 1. Visão do Produto

A **FBSO Platform** é uma plataforma SaaS multi-produto que opera no modelo de Suíte com módulos ativáveis por plano contratado. O Core administrativo — escopo deste projeto — é a camada fundamental que gerencia contas de clientes (Tenants), planos comerciais, assinaturas, usuários e permissões de acesso (RBAC). É o alicerce sobre o qual todos os módulos-produto futuros (Tributali-Engine, Storekeeper Portal) serão acoplados.

**Por que este projeto existe:** A FBSO.ORG precisa de uma plataforma administrativa unificada para gerenciar seus clientes SaaS, permitindo que o time interno tenha visibilidade e controle operacional, e que os próprios clientes tenham autonomia para gerenciar suas contas, usuários e catálogos via portal de auto-serviço.

**Para quem:** Dois públicos principais:
1. **Time Administrativo FBSO.ORG** — operadores internos que gerenciam contas, ativam clientes, configuram planos e controlam acessos
2. **Clientes da Plataforma** — administradores dos tenants que gerenciam suas unidades de negócio, usuários, produtos e assinaturas via portal

**O que NÃO é escopo:** Os módulos-produto Tributali-Engine e Storekeeper Portal serão desenvolvidos em fases futuras. O escopo atual é o Core Administrativo que suportará esses módulos.

### Critérios de Sucesso (do Project Charter)

| ID | Critério | Métrica |
|:---|:---|:---|
| C1 | Portal Admin operacional | 100% das funcionalidades D1-D4 em produção |
| C2 | Portal Cliente operacional | 100% das funcionalidades D5-D7 em produção |
| C3 | Autonomia do cliente | ≥80% dos onboardings sem intervenção do suporte |
| C4 | Tempo de ativação | ≤10 minutos para conclusão do onboarding |
| C5 | Segurança multi-tenant | Zero incidentes de cross-tenant data leakage |
| C6 | Disponibilidade | 99.5% uptime do portal |
| C7 | Escalabilidade | Suporte a 500+ tenants no primeiro ano |
| C8 | Time-to-market | 7 entregas em 14 semanas (24/07–30/10/2026) |

---

## 2. Personas e Jornadas

### 2.1 Personas

| Persona | Papel | Necessidades Principais | Dores |
|:---|:---|:---|:---|
| **Operador Admin FBSO** | Time interno que gerencia a plataforma | Visão consolidada de todas as contas, status e métricas; capacidade de ativar/suspender clientes rapidamente | Falta de visibilidade operacional; processos manuais de ativação |
| **Gestor Comercial FBSO** | Define planos e preços | Configurar planos comerciais com diferentes módulos e preços; visualizar adoção por plano | Rigidez na definição de ofertas comerciais |
| **Admin do Tenant (Cliente)** | Administrador da conta do cliente | Gerenciar usuários, unidades de negócio e produtos da sua empresa; ter autonomia sem depender do suporte FBSO | Dependência do suporte para tarefas simples; falta de visão da sua conta |
| **Usuário Operacional (Cliente)** | Usuário final do portal do cliente | Acessar os módulos contratados; navegar entre unidades de negócio | Acesso complexo; não saber quais módulos estão disponíveis |
| **Auditor** | Revisor de conformidade | Visualizar histórico de ações e mudanças | Falta de rastreabilidade; logs incompletos |

### 2.2 Jornadas Principais

#### J1 — Acompanhamento Diário da Operação (Operador Admin)
O operador acessa o dashboard administrativo, visualiza métricas em tempo real (total de contas ativas, por status, por plano), identifica contas que precisam de atenção (alertas), e navega pela lista de contas com filtros para análise detalhada.
- **US relacionadas:** US-0001 a US-0007 (7 US)

#### J2 — Ativação de Novo Cliente (Operador Admin)
O time comercial fecha contrato. O operador cadastra a nova conta (razão social, segmento), define o plano e vigência, e o sistema envia automaticamente e-mail de boas-vindas com link de ativação. O cliente recebe o link e inicia o onboarding.
- **US relacionadas:** US-0008 a US-0011, US-0015 a US-0021 (11 US)

#### J3 — Gestão de Acessos e Permissões (Operador Admin + Admin Tenant)
O Admin do Tenant solicita novos usuários. O operador FBSO ou o Admin do Tenant convida usuários, define papéis (Admin/Gerente/Operador/Auditor), vincula a unidades de negócio e módulos específicos. O sistema aplica as permissões automaticamente nos menus e ações.
- **US relacionadas:** US-0024 a US-0036, US-0059 a US-0061 (16 US)

#### J4 — Primeiro Acesso e Onboarding (Admin Tenant — Cliente)
O Admin do Tenant recebe e-mail de boas-vindas, clica no link, cria senha, e é guiado por um fluxo passo a passo: confirmar dados cadastrais → cadastrar primeira Unidade de Negócio → visualizar plano contratado → tela de boas-vindas com próximos passos.
- **US relacionadas:** US-0037 a US-0044 (8 US)

#### J5 — Gestão do Catálogo de Produtos (Usuário Operacional — Cliente)
O usuário acessa o portal, usa o App Switcher para navegar entre módulos, seleciona a Unidade de Negócio, e gerencia o catálogo de produtos/serviços daquela unidade (cadastrar, editar, ativar/desativar).
- **US relacionadas:** US-0047 a US-0058 (12 US)

#### J6 — Auto-Serviço e Upgrade (Admin Tenant — Cliente)
O Admin do Tenant visualiza seu dashboard com resumo da conta, vê notificações relevantes, e pode realizar upgrade de plano diretamente pelo portal (self-service), sem precisar contatar o suporte FBSO.
- **US relacionadas:** US-0045, US-0046, US-0062 (3 US)

---

## 3. Escopo por Entrega (D1-D7)

### D1 — Portal Administrativo Interno (M2: 15/08)
**Valor de Negócio:** Time FBSO.ORG tem visibilidade operacional em tempo real de todas as contas, status e métricas.
| Feature | US | Prioridade |
|:---|---:|:---|
| FEAT-EP-0001-0001 — Dashboard Métricas Operacionais | 3 | Must |
| FEAT-EP-0001-0002 — Visão de Contas com Filtros | 2 | Must |
| FEAT-EP-0001-0003 — Alertas e Indicadores de Atenção | 2 | Should |
| **Total** | **7 US** | |

### D2 — Gestão de Contas de Clientes (M3: 31/08)
**Valor de Negócio:** Time FBSO.ORG cadastra e gerencia contas de clientes com todo o ciclo de vida (criação → ativação → suspensão → reativação).
| Feature | US | Prioridade |
|:---|---:|:---|
| FEAT-EP-0002-0001 — Cadastro e Ativação de Contas | 4 | Must |
| FEAT-EP-0002-0002 — Gestão de Status do Tenant | 3 | Must |
| FEAT-EP-0002-0005 — Histórico de Auditoria | 2 | Must |
| **Total** | **9 US** | |

### D3 — Gestão de Planos e Assinaturas (M3: 31/08)
**Valor de Negócio:** Time FBSO.ORG configura ofertas comerciais e vincula clientes a planos com upgrade/downgrade.
| Feature | US | Prioridade |
|:---|---:|:---|
| FEAT-EP-0002-0003 — Configuração de Planos Comerciais | 4 | Must |
| FEAT-EP-0002-0004 — Vinculação e Gestão de Assinaturas | 3 | Must |
| **Total** | **7 US** | |

### D4 — Gestão de Usuários e Permissões (M4: 15/09)
**Valor de Negócio:** Controle granular de acesso — cada usuário vê apenas o que seu papel e permissões permitem, no escopo das unidades e módulos autorizados.
| Feature | US | Prioridade |
|:---|---:|:---|
| FEAT-EP-0003-0001 — Cadastro e Convite de Usuários | 6 | Must |
| FEAT-EP-0003-0002 — Papéis e Permissões (RBAC) | 4 | Must |
| FEAT-EP-0003-0003 — Vinc. Usuário×Unidade×Módulo | 3 | Must |
| FEAT-EP-0003-0004 — Controle de Visibilidade de Menus | 3 | Must |
| **Total** | **16 US** | |

### D5 — Portal do Cliente (M5: 30/09)
**Valor de Negócio:** Clientes têm autonomia para acessar a plataforma, completar onboarding guiado e navegar entre módulos.
| Feature | US | Prioridade |
|:---|---:|:---|
| FEAT-EP-0004-0001 — Autenticação e Recuperação de Senha | 3 | Must |
| FEAT-EP-0004-0002 — Onboarding Guiado de Primeiro Acesso | 5 | Must |
| FEAT-EP-0004-0003 — Dashboard do Cliente | 3 | Should |
| FEAT-EP-0004-0004 — App Switcher | 3 | Must |
| **Total** | **14 US** | |

### D6 — Cadastro de Unidades de Negócio (M6: 15/10)
**Valor de Negócio:** Clientes gerenciam sua estrutura organizacional (matriz e filiais) com hierarquia.
| Feature | US | Prioridade |
|:---|---:|:---|
| FEAT-EP-0004-0005 — Gestão de Unidades de Negócio | 5 | Must |
| **Total** | **5 US** | |

### D7 — Catálogo de Produtos e Serviços (M6: 15/10)
**Valor de Negócio:** Clientes mantêm seu portfólio de produtos ativo e atualizado por unidade de negócio.
| Feature | US | Prioridade |
|:---|---:|:---|
| FEAT-EP-0004-0006 — Catálogo de Produtos e Serviços | 4 | Must |
| **Total** | **4 US** | |

### Bônus — Dashboard do Cliente (Should Have)
| Feature | US | Prioridade |
|:---|---:|:---|
| FEAT-EP-0004-0003 — Dashboard do Cliente (parcial) | 3 | Should |
| **Total** | **3 US** (US-0045, US-0046, US-0062) | |

---

## 4. Matriz US × Jornada

| Jornada | US | Quantidade |
|:---|---:|---:|
| J1 — Acompanhamento Diário da Operação | US-0001 a US-0007 | 7 |
| J2 — Ativação de Novo Cliente | US-0008 a US-0011, US-0015 a US-0021 | 11 |
| J3 — Gestão de Acessos e Permissões | US-0024 a US-0036, US-0059 a US-0061 | 16 |
| J4 — Primeiro Acesso e Onboarding | US-0037 a US-0044 | 8 |
| J5 — Gestão do Catálogo de Produtos | US-0047 a US-0058 | 12 |
| J6 — Auto-Serviço e Upgrade | US-0045, US-0046, US-0062 | 3 |
| Transversal — Auditoria | US-0012 a US-0014, US-0022, US-0023 | 5 |

---

## 5. Restrições de Negócio

| Restrição | Descrição | Origem |
|:---|:---|:---|
| **LGPD** | Dados de usuários e clientes devem ser tratados conforme LGPD (consentimento, direito ao esquecimento, portabilidade) | BRD, Charter |
| **SLA de Disponibilidade** | Portal deve ter 99.5% de uptime (Critério C6 do Charter) | Charter |
| **Isolamento Multi-Tenant** | Zero cross-tenant data leakage — cada tenant só acessa seus próprios dados (Critério C5) | Charter, BRD |
| **Segurança** | Autenticação robusta (bloqueio por tentativas, recuperação segura de senha), autorização granular (RBAC), criptografia em trânsito e repouso | BRD |
| **Time-to-Market** | 7 entregas em 14 semanas (24/07 a 30/10/2026) — cronograma agressivo | Charter |
| **Time** | Equipe de 9 profissionais, alguns em carga parcial. Frontend dedicado somente a partir de 01/11 | Charter |
| **Orçamento** | Definido no Project Charter — ver Seção 12 do Charter | Charter |

### Métricas de Sucesso (KPIs)
Conforme [Matriz KPI](../MATRIZ-KPI.md), as seguintes dimensões são monitoradas:
- **Adoção e Autonomia:** Taxa de Onboarding Autônomo ≥80%, Tempo Médio de Onboarding ≤10min, Taxa de Abandono ≤15%
- **Operação e Governança:** Tempo Médio de Ativação, Taxa de Utilização do Portal Admin, SLA de Disponibilidade
- **Satisfação e Qualidade:** NPS do Portal, Taxa de Chamados de Suporte, Tempo Médio de Resolução

### Definition of Done
Conforme [DoD](../DEFINITION_OF_DONE.md), cada user story, feature e entrega deve satisfazer critérios objetivos de conclusão em 3 níveis cumulativos (US → Feature → Entrega). Ver DoD para checklist completo.

---

## 6. Glossário de Negócio

| Termo | Definição |
|:---|:---|
| **Tenant** | Conta de cliente na plataforma SaaS. Cada tenant tem seus próprios usuários, unidades de negócio, assinaturas e dados isolados |
| **Unidade de Negócio (BU)** | Filial ou divisão de uma empresa cliente. Estruturada hierarquicamente (Matriz → Filiais) |
| **Plano Comercial** | Oferta contratável que define quais módulos/produtos o cliente pode acessar, com valor e recorrência |
| **Assinatura** | Vínculo entre um tenant e um plano comercial, com data de início, vigência e status |
| **RBAC** | Role-Based Access Control — controle de acesso baseado em papéis predefinidos |
| **Onboarding** | Fluxo guiado de primeiro acesso onde o cliente configura sua conta, cadastra a primeira unidade de negócio e conhece a plataforma |
| **App Switcher** | Seletor de módulos no topo do portal que permite ao usuário navegar entre os produtos contratados |
| **Core Administrativo** | Camada fundamental da plataforma — gerencia tenants, planos, assinaturas, usuários e permissões. É o alicerce para módulos futuros |
| **Módulo-Produto** | Aplicação de negócio acoplada ao Core (ex: Tributali-Engine, Storekeeper Portal). Fora do escopo atual |

Para glossário completo, ver [GLOSSARY.md](../GLOSSARY.md).

---

## 7. Referências aos Documentos de Projeto

| Documento | Path | Conteúdo Relevante |
|:---|:---|:---|
| Project Charter | `01-PROJECT-CHARTER-...md` | Escopo, entregas D1-D7, milestones M1-M7, critérios de sucesso, riscos, orçamento |
| BRD | `02-BRD-...md` | Requisitos de negócio BR-01 a BR-10, regras de validação |
| Épicos | `03-EPICS-...md` + `epics/` | 4 épicos com escopo macro e objetivos de negócio |
| Features | `04-FEATURES-...md` + `features/` | 18 features detalhadas com US, critérios de aceitação e regras de negócio |
| User Stories | `05-USER-STORIES-...md` + `user-stories/` | 62 US com cenários Gherkin e rastreabilidade completa (RTM) |
| Glossário | `GLOSSARY.md` | Terminologia completa da plataforma e do domínio |
| Matriz KPI | `MATRIZ-KPI.md` | Indicadores de adoção, operação e satisfação para reporte executivo |
| Stakeholder Map | `STAKEHOLDER-MAP.md` | Partes interessadas e responsabilidades |
| DoD | `DEFINITION_OF_DONE.md` | Critérios de conclusão em 3 níveis (US, Feature, Entrega) |

---

🤖 *Documento gerado pelo Negócio (PO/PM/Funcional) para o time de TI — Fase 1 do Downstream Architecture Refinement. Independente de upstream discovery.*
