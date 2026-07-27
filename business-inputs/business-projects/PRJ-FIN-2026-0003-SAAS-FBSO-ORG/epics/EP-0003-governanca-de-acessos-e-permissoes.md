# EPICO - EP-0003: Governança de Acessos e Permissões

| Campo | Detalhe |
|-------|---------|
| **Épico** | EP-0003 — Governança de Acessos e Permissões |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Épicos (Estrutura Modular v4.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` v1.1 e `01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` v1.1 |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Épicos:** [`03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Anterior:** [EP-0002 — Gestão de Clientes e Assinaturas](../EP-0002-gestao-de-clientes-e-assinaturas.md) | **Próximo:** [EP-0004 — Experiência do Cliente e Autoatendimento](../EP-0004-experiencia-do-cliente-e-autoatendimento.md)

---

## 1. Nome do Épico
**Governança de Acessos e Permissões — RBAC Multi-Tenant e Multi-Unidade**

**Requisitos BRD Vinculados:** [BR-05](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Gestão de Usuários e Permissões

## 2. Objetivo (Goal)

- **Problema:** Em um SaaS multi-produto com múltiplos clientes (Tenants) e múltiplas filiais (Unidades de Negócio) por cliente, a segurança dos dados é um requisito inegociável. Um operador de caixa de uma filial não pode acessar dados fiscais de outra filial; um auditor externo não pode alterar configurações; e um cliente não pode, sob nenhuma hipótese, visualizar dados de outro cliente. Sem um sistema robusto de permissões, a FBSO Platform estaria exposta a riscos legais, fiscais e de reputação.
- **Solução:** Implementar controle de acesso baseado em papéis (RBAC) que permita ao administrador do tenant definir precisamente o que cada usuário pode ver e fazer — em qual módulo, em qual unidade de negócio e com qual nível de permissão (Admin, Gerente, Operador, Auditor).
- **Impacto:** Segurança e conformidade; isolamento total de dados entre tenants e entre unidades de negócio; flexibilidade para o cliente gerenciar sua própria equipe; base para o modelo de App Switcher com visibilidade condicional por módulo.

## 3. Personas de Usuário (User Personas)

| Persona | Descrição | Necessidades |
|---------|-----------|-------------|
| **Administrador do Tenant** | Dono da conta ou contador master | Convidar usuários; definir quem acessa o quê; revogar acessos; visão completa de todas as unidades |
| **Gerente de Unidade (Manager BU)** | Responsável por uma filial específica | Gerenciar produtos e operações apenas na sua unidade; não pode alterar regras fiscais |
| **Operador de Unidade (Operator BU)** | Funcionário operacional (ex: faturamento, caixa) | Executar tarefas na sua unidade específica; sem acesso a configurações sensíveis |
| **Auditor (fase futura)** | Auditor externo ou interno — previsto para fase posterior ao MVP | Visualizar dados das unidades autorizadas; não pode criar, editar ou excluir nada |

## 4. Jornadas de Usuário de Alto Nível (High-Level User Journeys)

**Jornada 1: Administrador convida novo usuário para o time**
1. Administrador do tenant acessa "Gestão de Usuários"
2. Clica em "Convidar Usuário" e preenche e-mail, nome e papel (ex: Gerente)
3. Seleciona quais Unidades de Negócio o usuário poderá acessar (ex: Filial SP e Filial RJ)
4. Seleciona quais módulos o usuário pode acessar (ex: apenas Storekeeper Portal)
5. Sistema envia convite por e-mail; usuário define senha no primeiro acesso
6. Ao fazer login, o usuário vê apenas as unidades e módulos autorizados
> 🏷️ Atende [BR-05](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

**Jornada 2: Restrição de acesso entre filiais**
1. João (Operador da Filial SP) faz login no portal
2. O seletor de Unidade de Negócio exibe apenas "Filial SP"
3. João acessa "Catálogo de Produtos" e vê apenas produtos da Filial SP
4. João não vê a opção "Configurações Fiscais" pois não tem permissão
5. João não vê o App Switcher (só tem acesso a um módulo)
> 🏷️ Atende [BR-05](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

**Jornada 3: Revogação de acesso**
1. Administrador remove o acesso de um usuário que saiu da empresa
2. Usuário não consegue mais fazer login (acesso bloqueado imediatamente)
3. Ação registrada em auditoria
> 🏷️ Atende [BR-05](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

## 5. Requisitos de Negócio (Business Requirements)

### Requisitos Funcionais

- Cadastro de usuários vinculados a um Tenant (convite por e-mail)
- Definição de papéis de acesso (MVP): Admin do Tenant, Gerente de Unidade, Operador. Papel "Auditor" documentado e com schema previsto para fase futura.
- Vinculação de usuário a uma ou mais Unidades de Negócio
- Vinculação de usuário a um ou mais módulos/produtos da plataforma
- Cada papel define um conjunto de permissões: visualizar, criar, editar, excluir por área funcional
- Usuário sem permissão para um módulo não vê o módulo no App Switcher
- Usuário sem permissão para uma funcionalidade não vê o menu ou botão correspondente
- Usuário sem permissão para uma Unidade de Negócio não acessa seus dados
- Administrador do tenant pode revogar acesso a qualquer momento
- Registro de auditoria para criação, alteração e revogação de permissões

### Requisitos Não-Funcionais

- Verificação de permissão em todas as operações — não basta esconder o menu, o acesso deve ser barrado também se o usuário tentar acessar diretamente
- Bloqueio de acesso em até 5 minutos após revogação
- Senhas e credenciais gerenciadas com segurança (política de complexidade mínima)

## 6. Métricas de Sucesso (Success Metrics)

| KPI | Meta |
|-----|------|
| Incidentes de acesso não autorizado entre unidades de negócio | Zero |
| Incidentes de acesso não autorizado entre tenants | Zero |
| Tempo para configurar permissões de um novo usuário | ≤ 3 minutos |
| Cobertura de verificação de permissão | 100% das operações |

## 7. Fora do Escopo (Out of Scope)

- Integração com provedores de identidade corporativos (SSO / SAML / Azure AD) — funcionalidade futura
- Autenticação em duas etapas (MFA) — funcionalidade futura
- Permissões customizáveis por cliente (além dos papéis padrão) — funcionalidade futura
- Login social (Google, LinkedIn) — funcionalidade futura
- Delegação temporária de permissões — funcionalidade futura

## 8. Valor de Negócio (Business Value)

| Critério | Avaliação | Justificativa |
|----------|-----------|---------------|
| Valor de Negócio | **Crítico** | Segurança de dados é pré-requisito legal e de mercado. Sem RBAC robusto, o SaaS não pode operar com clientes reais. |

---

## Matriz de Rastreabilidade BRD → Este Épico

| BRD | Requisito Funcional | Este Épico | Jornada(s) que Realizam |
|:---|:---|:---|:---|
| **BR-05** | Gestão de Usuários e Permissões | **EP-0003** — Governança de Acessos e Permissões | J1: Convite de novo usuário · J2: Restrição de acesso entre filiais · J3: Revogação de acesso |

---

> 📄 **Índice de Épicos:** [`03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Anterior:** [EP-0002 — Gestão de Clientes e Assinaturas](../EP-0002-gestao-de-clientes-e-assinaturas.md) | **Próximo:** [EP-0004 — Experiência do Cliente e Autoatendimento](../EP-0004-experiencia-do-cliente-e-autoatendimento.md)

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: breakdown-epic-pm, agile-ba-practices. Estrutura modular v4.0.*

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE ÉPICOS]
