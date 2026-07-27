# FEATURE - FEAT-EP-0002-0003: Configuração de Planos Comerciais

| Campo | Detalhe |
|-------|---------|
| **Feature** | FEAT-EP-0002-0003 — Configuração de Planos Comerciais |
| **Épico** | [EP-0002 — Gestão de Clientes e Assinaturas](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0002](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md) | **Anterior:** [FEAT-EP-0002-0002 — Gestão de Status](../FEAT-EP-0002-0002-gestao-status-tenant.md) | **Próximo:** [FEAT-EP-0002-0004 — Vinculação de Assinaturas](../FEAT-EP-0002-0004-vinculacao-gestao-assinaturas.md)

**Requisitos BRD Vinculados:** [BR-03](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Configuração de Planos Comerciais

---

## Objetivo de Negócio
Permitir que o time de produto crie e gerencie os planos comerciais do SaaS de forma autônoma, definindo preços, módulos incluídos e recorrências.

**Prioridade:** Must Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-015 | Como **Gestor de Produto**, quero cadastrar um novo plano comercial definindo nome, descrição, valor mensal e recorrência disponível (mensal, trimestral, anual) | • Formulário de cadastro com todos os campos • Valores monetários formatados em Real (R$) • Recorrências selecionáveis via checkboxes (pode marcar mais de uma) |
| US-016 | Como **Gestor de Produto**, quero definir quais módulos/produtos da plataforma um plano inclui (ex: Tributali-Engine, Storekeeper Portal) para controlar o que cada cliente pode acessar | • Lista de módulos disponíveis com checkbox ao lado de cada um • Módulos marcados são incluídos no plano • Plano "Full Suite" inclui todos os módulos automaticamente |
| US-017 | Como **Gestor de Produto**, quero editar um plano existente (nome, preço, módulos) mantendo o histórico de versões anteriores | • Edição de plano gera nova versão • Clientes já vinculados permanecem na versão contratada até upgrade • Não é possível excluir plano, apenas desativá-lo |
| US-018 | Como **Administrador FBSO**, quero desativar um plano comercial para que ele não esteja mais disponível para novas contratações, sem afetar clientes que já o possuem | • Plano desativado não aparece como opção em novas assinaturas • Clientes ativos no plano desativado continuam com acesso normal • Plano aparece na lista administrativa com indicador "Descontinuado" |

## Regras de Negócio

- **RN-FEAT-EP-0002-0003-0001:** Um plano não pode ser excluído se houver clientes ativos vinculados a ele
- **RN-FEAT-EP-0002-0003-0002:** Alteração de preço de plano não afeta assinaturas já contratadas (vale o preço da data de contratação)
- **RN-FEAT-EP-0002-0003-0003:** Deve existir pelo menos um plano ativo no sistema

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-03** | Configuração de Planos Comerciais | [EP-0002](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md) / Requisitos Funcionais §5 | **FEAT-EP-0002-0003** — Configuração de Planos Comerciais |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0002](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md)

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
