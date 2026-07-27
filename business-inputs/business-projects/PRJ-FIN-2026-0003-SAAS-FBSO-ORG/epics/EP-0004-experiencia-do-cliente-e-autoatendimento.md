# EPICO - EP-0004: Experiência do Cliente e Autoatendimento

| Campo | Detalhe |
|-------|---------|
| **Épico** | EP-0004 — Experiência do Cliente e Autoatendimento |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Épicos (Estrutura Modular v4.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` v1.1 e `01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` v1.1 |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Épicos:** [`03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Anterior:** [EP-0003 — Governança de Acessos e Permissões](../EP-0003-governanca-de-acessos-e-permissoes.md)

---

## 1. Nome do Épico
**Experiência do Cliente e Autoatendimento — Portal, Onboarding e Catálogo**

**Requisitos BRD Vinculados:** [BR-06](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) (Portal do Cliente), [BR-07](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) (Onboarding Guiado), [BR-08](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) (App Switcher), [BR-09](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) (Unidades de Negócio), [BR-10](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) (Catálogo de Produtos/Serviços)

## 2. Objetivo (Goal)

- **Problema:** Para que a FBSO Platform seja percebida como um produto profissional e escalável, o cliente precisa de uma experiência de entrada (onboarding) fluida e autônoma, um portal onde possa gerenciar seus dados sem depender do time FBSO.ORG, e ferramentas para cadastrar o que é essencial ao seu negócio: suas filiais (Unidades de Negócio) e seu portfólio de produtos. Sem essa camada de autoatendimento, cada novo cliente geraria uma carga manual insustentável para o time interno.
- **Solução:** Construir o portal do cliente com fluxo de onboarding guiado, App Switcher para navegação entre módulos, gestão de unidades de negócio com estrutura hierárquica (Matriz/Filial) e catálogo de produtos/serviços com classificação por tipo e status — tudo adaptado ao plano e às permissões do usuário.
- **Impacto:** Redução drástica do esforço de ativação de clientes; experiência de produto profissional desde o primeiro acesso; base de dados comercial estruturada para futura ativação dos módulos fiscais e de varejo.

## 3. Personas de Usuário (User Personas)

| Persona | Descrição | Necessidades |
|---------|-----------|-------------|
| **Cliente Administrador** | Dono ou responsável pela conta da empresa no SaaS | Fazer onboarding rápido; cadastrar filiais; gerenciar catálogo; convidar usuários do time |
| **Cliente Operador** | Funcionário que usa o portal no dia a dia | Acessar funcionalidades da sua unidade; realizar tarefas operacionais; ver apenas o que lhe compete |
| **Contador (futuro)** | Profissional contábil que gerencia múltiplos clientes — funcionalidade multi-tenant NÃO está no escopo desta fase | Alternar entre clientes (unidades de negócio de diferentes tenants); visão multi-empresa (fase futura) |

## 4. Jornadas de Usuário de Alto Nível (High-Level User Journeys)

**Jornada 1: Primeiro acesso e onboarding**
1. Cliente recebe e-mail com link de ativação da conta
2. Clica no link, define sua senha e faz login pela primeira vez
3. Sistema identifica que é primeiro acesso e inicia o fluxo de onboarding
4. **Passo 1:** Confirma dados cadastrais da empresa
5. **Passo 2:** Cadastra a primeira Unidade de Negócio (CNPJ matriz, regime tributário)
6. **Passo 3:** Recebe orientação sobre os módulos disponíveis no seu plano
7. **Passo 4:** É direcionado ao portal com uma tela de boas-vindas
8. Portal liberado; menus adaptados ao plano e permissões
> 🏷️ Atende [BR-06](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [BR-07](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

**Jornada 2: Cadastro de filiais (Unidades de Negócio)**
1. Administrador do tenant acessa "Unidades de Negócio"
2. Visualiza a matriz cadastrada no onboarding
3. Clica em "+ Nova Unidade" para cadastrar uma filial
4. Preenche CNPJ, endereço, regime tributário e define o vínculo hierárquico (Filial de: Matriz)
5. Sistema valida CNPJ duplicado ativo e exibe confirmação
6. Nova unidade aparece na lista, com indicador visual de hierarquia (recuo para filiais)
> 🏷️ Atende [BR-09](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

**Jornada 3: Cadastro de portfólio de produtos**
1. Gerente de unidade acessa "Catálogo de Produtos e Serviços"
2. Visualiza lista de itens cadastrados para sua unidade
3. Clica em "+ Novo Produto" e preenche: nome, SKU, tipo (Produto ou Serviço), status
4. Sistema salva o produto e o exibe na lista
5. Indicador visual mostra que o produto ainda não possui mapeamento fiscal (preparação para Tributali-Engine)
> 🏷️ Atende [BR-10](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

**Jornada 4: Navegação com App Switcher (visão futura)**
1. Cliente com plano Full Suite faz login
2. App Switcher no topo mostra: "Storekeeper Portal" (selecionado) e "Tributali-Engine"
3. Clica em "Tributali-Engine" e o menu lateral se transforma para mostrar opções fiscais
4. Volta ao Storekeeper e o menu retorna para opções de varejo
> 🏷️ Atende [BR-08](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

## 5. Requisitos de Negócio (Business Requirements)

### Requisitos Funcionais — Portal do Cliente e Onboarding

- Tela de login e recuperação de senha
- Fluxo de primeiro acesso (onboarding): detecção automática de primeiro login
- Passos do onboarding: confirmação de dados, cadastro da primeira Unidade de Negócio, orientação sobre o plano
- Dashboard do cliente pós-onboarding com visão adaptada ao plano e ao módulo ativo
- Área de perfil do usuário: nome, e-mail, alteração de senha

### Requisitos Funcionais — App Switcher

- Seletor de módulos visível no topo do portal, ao lado da identidade visual da FBSO Platform
- Exibe apenas os módulos que o usuário tem permissão para acessar
- Ao alternar de módulo, o menu lateral e o conteúdo se adaptam ao módulo selecionado
- Na Fase 0 (este projeto), haverá apenas um módulo visível (placeholder para os futuros)

### Requisitos Funcionais — Unidades de Negócio

- Cadastro de Unidade de Negócio: CNPJ, razão social, nome fantasia, regime tributário, endereço
- Estrutura hierárquica Matriz/Filial com indicador visual de relacionamento
- Validação de CNPJ duplicado por tenant (CNPJ ativo não pode ser cadastrado duas vezes)
- Listagem de unidades vinculadas ao tenant, agrupadas por hierarquia
- Indicador de status (ativo/inativo) por unidade

### Requisitos Funcionais — Catálogo de Produtos/Serviços

- Cadastro de item: nome, SKU/código, tipo (Produto, Serviço), descrição, status (ativo/inativo)
- Listagem de itens por Unidade de Negócio com busca por nome ou SKU
- Indicador visual de status: Ativo, Inativo
- Indicador visual "Preparado para mapeamento fiscal" (placeholder para fase futura)
- Edição e ativação/desativação de itens

### Requisitos Não-Funcionais

- Onboarding completo em até 4 passos; tempo total ≤ 10 minutos
- Portal do cliente responsivo (desktop principal; tablet aceitável)
- Interface em português (Brasil)
- Mensagens de erro em linguagem clara e não-técnica

## 6. Métricas de Sucesso (Success Metrics)

| KPI | Meta |
|-----|------|
| Clientes que completam onboarding sem ajuda do suporte | ≥ 80% |
| Tempo médio de onboarding completo (login → portal liberado) | ≤ 10 minutos |
| Abandono durante onboarding | ≤ 15% |
| Satisfação com a experiência de primeiro acesso | Nota ≥ 4,0 / 5,0 |
| Unidades de Negócio cadastradas com sucesso na primeira tentativa | ≥ 95% |

## 7. Fora do Escopo (Out of Scope)

- Funcionalidades específicas dos módulos Tributali-Engine e Storekeeper Portal
- Mapeamento fiscal de produtos (NCM, NBS, CNAE, alíquotas IBS/CBS)
- Emissão de pedidos e faturas
- Customização visual do portal por tenant (white label)
- Aplicativo mobile
- Portal em outros idiomas além de português (Brasil)

## 8. Valor de Negócio (Business Value)

| Critério | Avaliação | Justificativa |
|----------|-----------|---------------|
| Valor de Negócio | **Crítico** | O portal do cliente é a face visível do SaaS. A experiência de onboarding é o primeiro contato do cliente com o produto. Um onboarding ruim gera abandono e sobrecarrega o suporte. |

---

## Matriz de Rastreabilidade BRD → Este Épico

| BRD | Requisito Funcional | Este Épico | Jornada(s) que Realizam |
|:---|:---|:---|:---|
| **BR-06** | Portal do Cliente com Autenticação | **EP-0004** — Experiência do Cliente | J1: Primeiro acesso e onboarding |
| **BR-07** | Onboarding Guiado de Primeiro Acesso | **EP-0004** — Experiência do Cliente | J1: Primeiro acesso e onboarding |
| **BR-08** | App Switcher (Seletor de Aplicativos) | **EP-0004** — Experiência do Cliente | J4: Navegação com App Switcher (visão futura) |
| **BR-09** | Cadastro de Unidades de Negócio | **EP-0004** — Experiência do Cliente | J2: Cadastro de filiais (Unidades de Negócio) |
| **BR-10** | Catálogo de Produtos/Serviços | **EP-0004** — Experiência do Cliente | J3: Cadastro de portfólio de produtos |

---

> 📄 **Índice de Épicos:** [`03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Anterior:** [EP-0003 — Governança de Acessos e Permissões](../EP-0003-governanca-de-acessos-e-permissoes.md)

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: breakdown-epic-pm, agile-ba-practices. Estrutura modular v4.0.*

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE ÉPICOS]
