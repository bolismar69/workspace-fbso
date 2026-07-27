# User Stories: Configuração de Planos Comerciais

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** FEAT-EP-0002-0003 — Configuração de Planos Comerciais
- **Épico:** EP-0002 — Gestão de Clientes e Assinaturas
- **Prioridade:** Must Have
- **Data-Alvo:** 31/08/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

---

## Objetivo de Negócio

Permitir que o time de produto crie e gerencie os planos comerciais do SaaS de forma autônoma, definindo preços, módulos incluídos e recorrências.

---

## User Stories

### US-015 — Cadastro de Novo Plano Comercial

**Como** Administrador FBSO, **quero** cadastrar um novo plano comercial definindo nome, descrição, valor mensal e recorrência disponível (mensal, trimestral, anual). (Na prática, funcionalidade será utilizada pelo Gestor de Produto — perfil de stakeholder, não papel RBAC do sistema.)

**Critérios de Aceitação:**
- Formulário de cadastro com todos os campos
- Valores monetários formatados em Real (R$)
- Recorrências selecionáveis via checkboxes (pode marcar mais de uma)
- Valor do plano deve ser maior que zero. Sistema valida e exibe mensagem de erro para valores inválidos.
- Pelo menos uma opção de recorrência deve ser selecionada. Sistema valida antes de salvar.

### US-016 — Definição de Módulos Incluídos no Plano

**Como** Administrador FBSO, **quero** definir quais módulos/produtos da plataforma um plano inclui (ex: Tributali-Engine, Storekeeper Portal) **para** controlar o que cada cliente pode acessar. (Na prática, funcionalidade será utilizada pelo Gestor de Produto — perfil de stakeholder, não papel RBAC do sistema.)

**Critérios de Aceitação:**
- Lista de módulos disponíveis com checkbox ao lado de cada um
- Módulos marcados são incluídos no plano
- Plano "Full Suite" inclui todos os módulos automaticamente. O comportamento 'Full Suite' é acionado quando todos os módulos disponíveis são marcados no plano. Quando novos módulos forem adicionados à plataforma no futuro, planos 'Full Suite' os incluirão automaticamente.

### US-017 — Edição de Plano com Versionamento

**Como** Administrador FBSO, **quero** editar um plano existente (nome, preço, módulos) mantendo o histórico de versões anteriores. (Na prática, funcionalidade será utilizada pelo Gestor de Produto — perfil de stakeholder, não papel RBAC do sistema.)

**Critérios de Aceitação:**
- Edição de plano gera nova versão
- Clientes já vinculados permanecem na versão contratada até upgrade
- Plano pode ser excluído apenas se não houver clientes ativos vinculados a ele (RN-FEAT-EP-0002-0003-0001). Caso haja clientes ativos, o plano pode apenas ser desativado.

### US-018 — Desativação de Plano Comercial

**Como** Administrador FBSO, **quero** desativar um plano comercial para que ele não esteja mais disponível para novas contratações, sem afetar clientes que já o possuem.

**Critérios de Aceitação:**
- Plano desativado não aparece como opção em novas assinaturas
- Clientes ativos no plano desativado continuam com acesso normal
- Plano aparece na lista administrativa com indicador "Descontinuado"

---

## Regras de Negócio

| ID | Regra |
|----|-------|
| **RN-FEAT-EP-0002-0003-0001** | Um plano não pode ser excluído se houver clientes ativos vinculados a ele |
| **RN-FEAT-EP-0002-0003-0002** | Alteração de preço de plano não afeta assinaturas já contratadas (vale o preço da data de contratação) |
| **RN-FEAT-EP-0002-0003-0003** | Deve existir pelo menos um plano ativo no sistema |

---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | Plano cadastrado e disponível para vinculação em tempo real | Plano visível no seletor de nova assinatura |
| F2 | Edição gera nova versão sem afetar clientes existentes | Cliente em versão anterior com acesso preservado |
| F3 | Plano desativado indisponível para novas assinaturas | Seletor de planos sem o plano desativado |

---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*
