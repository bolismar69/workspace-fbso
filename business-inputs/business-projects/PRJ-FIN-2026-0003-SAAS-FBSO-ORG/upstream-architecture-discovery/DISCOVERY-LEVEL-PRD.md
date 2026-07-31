# DISCOVERY-LEVEL-PRD — Product Definition Discovery-Level

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Fase:** F1 — Bloco 0 (Product Definition Discovery-Level)
- **Versão:** 1.0 · **Data:** 30/07/2026 · **Status:** CREATED
- **Elaborado por:** PM / PO / Analista de Negócios
- **Público-alvo:** Time de TI (Discovery Team — Upstream Architecture)

---

## 1. Visão do Produto (Resumo Executivo)

A **FBSO Platform** é o portal administrativo SaaS multi-produto da FBSO.ORG — a camada fundamental ("Core") sobre a qual todos os produtos futuros da empresa serão acoplados como módulos ativáveis por plano de assinatura.

> Para a FBSO.ORG, que precisa de uma plataforma unificada para gerenciar clientes, planos, permissões e acesso aos seus produtos SaaS, a FBSO Platform oferece um portal administrativo multi-tenant que permite ativar, suspender e gerenciar contas de clientes, configurar planos comerciais com módulos ativáveis e controlar acesso granular por usuário e unidade de negócio — tudo em uma única plataforma.

**Diferencial de negócio:** Quando um cliente do Storekeeper Portal precisar do Split Payment da Reforma Tributária, o upgrade será uma simples ativação de módulo na mesma plataforma, sem migração de sistema.

---

## 2. Épicos de Negócio (Alto Nível)

### EP-0001 — Portal Administrativo Interno
- **Objetivo de negócio:** Fornecer ao time interno da FBSO uma visão consolidada da operação: quantos clientes ativos, qual o crescimento da base, status das contas e indicadores de atenção.

- **Funcionalidades esperadas:** Dashboard com métricas operacionais, lista de contas com busca e filtros, indicadores visuais de atenção (pendências, atrasos).

- **Valor para o negócio:** Centralizar a gestão operacional em um único portal, substituindo planilhas e controles manuais.

### EP-0002 — Gestão de Clientes e Assinaturas
- **Objetivo de negócio:** Permitir que o time Comercial e Administrativo cadastre novos clientes, configure planos, gerencie assinaturas e acompanhe o ciclo de vida de cada conta.

- **Funcionalidades esperadas:** Cadastro de conta com dados do cliente, ativação com envio de credenciais, configuração de planos comerciais, vinculação de assinaturas, histórico de auditoria.

- **Valor para o negócio:** Processo comercial estruturado — da prospecção à ativação do cliente — com rastreabilidade completa.

### EP-0003 — Governança de Acessos e Permissões (RBAC)
- **Objetivo de negócio:** Garantir que cada usuário do sistema acesse apenas o que seu perfil permite — administradores FBSO gerenciam tudo, administradores do cliente gerenciam seu próprio tenant, operadores têm acesso restrito.

- **Funcionalidades esperadas:** Cadastro e convite de usuários, definição de papéis e permissões, vinculação usuário × unidade de negócio × módulo, controle de visibilidade de menus e ações.

- **Valor para o negócio:** Segurança e compliance — cada cliente vê apenas seus dados; cada usuário acessa apenas suas funcionalidades.

### EP-0004 — Experiência do Cliente e Autoatendimento
- **Objetivo de negócio:** Oferecer ao cliente final um portal de autoatendimento onde ele gerencia sua conta, usuários e unidades de negócio sem depender do suporte da FBSO.

- **Funcionalidades esperadas:** Login seguro, recuperação de senha, onboarding guiado de primeiro acesso, dashboard do cliente, seletor de módulos (App Switcher), gestão de unidades de negócio.

- **Valor para o negócio:** Redução de chamados de suporte, autonomia do cliente, experiência profissional que posiciona a FBSO como plataforma enterprise.

---

## 3. MVP Macro — Primeiro Delivery

| Entrega de Negócio | Épicos envolvidos | Valor para o negócio |
|:---|:---|:---|
| **Operação interna funcionando** | EP-0001 + EP-0002 | Time FBSO gerencia clientes e planos pelo portal |
| **Acesso seguro e controlado** | EP-0003 | Autenticação dos usuários com papéis definidos |
| **Cliente faz autoatendimento** | EP-0004 | Cliente gerencia sua conta sem depender do suporte |

**Expectativa de prazo (negócio):** 3-4 meses para o MVP completo.

---

## 4. Restrições de Negócio

| Restrição | Impacto |
|:---|:---|
| Time enxuto (11 profissionais, alguns em carga parcial) | Velocidade de entrega proporcional à capacidade disponível |
| Profissional dedicado à experiência do cliente (frontend) disponível apenas a partir de 01/11 | Portal do Cliente pode iniciar após esta data |
| Ambiente de produção precisa estar operacional antes do primeiro cliente externo | Infraestrutura é pré-requisito para o MVP |
| Conformidade com LGPD para dados de clientes | Todas as funcionalidades devem contemplar privacidade desde a concepção |

---

## 5. Glossário de Negócio

| Termo | Definição |
|:---|:---|
| **Tenant** | Conta Master / Pagadora do cliente no SaaS |
| **Business Unit (Unidade de Negócio)** | CNPJ / Filial vinculada a um Tenant |
| **Plano Comercial** | Pacote de módulos e funcionalidades contratado pelo cliente |
| **Assinatura** | Vínculo ativo entre um Tenant e um Plano Comercial |
| **RBAC** | Modelo de controle de acesso baseado em papéis (Role-Based Access Control) |
| **App Switcher** | Seletor que permite ao cliente navegar entre diferentes produtos/módulos |

---

## 6. Documentos de Negócio Vinculados

| Documento | Conteúdo |
|:---|:---|
| [Project Charter](../01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | Escopo de alto nível, objetivos, premissas, riscos, governança |
| [BRD](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | Requisitos de negócio detalhados e regras de atendimento |
| [Épicos (índice)](../03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | Agrupamento de requisitos em grandes blocos de entrega de valor |
| [Épicos (detalhes)](../epics/) | 4 arquivos individuais com detalhamento completo de cada épico |

---

## Histórico de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: PRD Discovery-Level — visão de negócio baseada nos 4 Épicos | PM / PO / Time de Negócios |

🤖 *Upstream Architecture Discovery — Fase 1 · Bloco 0 · Documento de Negócio — briefing executivo para o time de TI iniciar o Discovery Técnico*
