# EP-0002: Gestão de Clientes e Assinaturas

| Campo | Detalhe |
|-------|---------|
| **Épico** | EP-0002 — Gestão de Clientes e Assinaturas |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Épicos (Estrutura Modular v4.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` v1.1 e `01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` v1.1 |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Épicos:** [`03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Anterior:** [EP-0001 — Portal Administrativo Interno](../EP-0001-portal-administrativo-interno.md) | **Próximo:** [EP-0003 — Governança de Acessos e Permissões](../EP-0003-governanca-de-acessos-e-permissoes.md)

---

## 1. Nome do Épico
**Gestão de Clientes e Assinaturas — Ciclo de Vida do Tenant**

**Requisitos BRD Vinculados:** [BR-02](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) (Ativação e Gestão de Contas), [BR-03](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) (Configuração de Planos Comerciais), [BR-04](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) (Vinculação de Assinaturas)

## 2. Objetivo (Goal)

- **Problema:** A FBSO.ORG precisa gerenciar o ciclo de vida completo dos seus clientes SaaS — da ativação inicial à suspensão por inadimplência ou upgrade de plano. Sem uma ferramenta centralizada, essas operações seriam manuais, propensas a erros e impossíveis de auditar. Além disso, a oferta comercial (planos e preços) precisa ser configurável pelo time de produto sem dependência de desenvolvimento técnico.
- **Solução:** Criar os módulos de gestão de contas de clientes (Tenants) e de planos/assinaturas, permitindo que o time administrativo ative, suspenda e reative contas, enquanto o time de produto configura os planos comerciais e o time comercial vincula clientes aos planos contratados.
- **Impacto:** Operação comercial estruturada e auditável; redução de erros manuais; autonomia do time de produto para criar e ajustar ofertas; base de dados de clientes consolidada para futura automação de faturamento.

## 3. Personas de Usuário (User Personas)

| Persona | Descrição | Necessidades |
|---------|-----------|-------------|
| **Administrador FBSO.ORG** | Gerencia a base de clientes no dia a dia | Ativar/suspender contas rapidamente; ver histórico de ações; identificar contas problemáticas |
| **Gestor de Produto** | Define e ajusta a oferta comercial | Criar novos planos; alterar preços e módulos incluídos; versionar planos |
| **Líder Comercial** | Acompanha a carteira de clientes | Verificar plano de cada cliente; identificar oportunidades de upgrade; acompanhar renovações |

## 4. Jornadas de Usuário de Alto Nível (High-Level User Journeys)

**Jornada 1: Ativação de novo cliente (venda consultiva — fase futura)**
1. Vendedor fecha contrato com cliente corporativo e gera uma Ordem de Serviço (processo comercial fora do escopo desta fase; simulado manualmente pelo time administrativo)
2. Administrador acessa o módulo de gestão de contas e cria o Tenant vinculado ao plano contratado
3. Sistema gera link de onboarding e o dispara para o e-mail do cliente
4. Administrador acompanha status: "Aguardando Onboarding" → "Ativo"
5. Caso o cliente não complete o onboarding em X dias, sistema alerta o administrador
> 🏷️ Atende [BR-02](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md), [BR-04](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

**Jornada 2: Suspensão por inadimplência**
1. Administrador identifica cliente com pagamento pendente (via alerta ou busca)
2. Acessa a conta do cliente e altera status para "Suspenso"
3. Sistema bloqueia acesso dos usuários do tenant a partir daquele momento
4. Ação registrada em auditoria com identificação do administrador, data e motivo
5. Futuramente, processo será automatizado via integração com gateway de pagamento (fora do escopo)
> 🏷️ Atende [BR-02](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

**Jornada 3: Upgrade de plano**
1. Cliente solicita upgrade do plano Básico para o Core
2. Comercial acessa a conta, seleciona novo plano e define nova data de vigência
3. Sistema atualiza os módulos disponíveis para o tenant imediatamente
4. No próximo login, o cliente visualiza os novos módulos no App Switcher
> 🏷️ Atende [BR-04](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

## 5. Requisitos de Negócio (Business Requirements)

### Requisitos Funcionais

- Criação de conta de cliente (Tenant) com dados corporativos: razão social, nome fantasia, segmento
- Ativação, suspensão e reativação de contas com registro de motivo
- Visualização do status de cada conta e tempo em cada status
- Cadastro de planos comerciais: nome, descrição, valor, recorrência (mensal, trimestral, anual)
- Definição de módulos/produtos incluídos em cada plano
- Vinculação de cliente a plano com data de início, vigência e status da assinatura
- Troca de plano (upgrade/downgrade) com registro de data de alteração
- Histórico de todas as ações administrativas: quem fez, o que fez, quando fez
- Estrutura de plano preparada para acoplar faturamento real no futuro (sem processar cobranças)

### Requisitos Não-Funcionais

- Registro de auditoria imutável para todas as ações administrativas (audit trail)
- Bloqueio de acesso ao portal do cliente em até 5 minutos após suspensão da conta
- Validação de unicidade de razão social por tenant
- Planos inativos (descontinuados) não podem ser vinculados a novas assinaturas

## 6. Métricas de Sucesso (Success Metrics)

| KPI | Meta |
|-----|------|
| Tempo para ativar uma nova conta (da decisão ao link de onboarding) | ≤ 2 minutos |
| Tempo para suspender conta e bloquear acesso | ≤ 5 minutos |
| Erros em configuração de plano (ex: módulo errado liberado) | Zero não conformidades |
| Cobertura de auditoria | 100% das ações administrativas |

## 7. Fora do Escopo (Out of Scope)

- Processamento de cobranças e faturamento real — estrutura de dados preparada, sem execução financeira
- Renovação automática de assinaturas — será manual nesta fase
- Integração com gateways de pagamento
- Período de trial gratuito automático — pode ser simulado via ativação manual
- Gestão de contratos e documentos legais — funcionalidade futura

## 8. Valor de Negócio (Business Value)

| Critério | Avaliação | Justificativa |
|----------|-----------|---------------|
| Valor de Negócio | **Crítico** | Sem gestão de clientes e assinaturas, não há operação SaaS. É o coração do modelo de negócio. |

---

## Matriz de Rastreabilidade BRD → Este Épico

| BRD | Requisito Funcional | Este Épico | Jornada(s) que Realizam |
|:---|:---|:---|:---|
| **BR-02** | Ativação e Gestão de Contas | **EP-0002** — Gestão de Clientes e Assinaturas | J1: Ativação de novo cliente · J2: Suspensão por inadimplência |
| **BR-03** | Configuração de Planos Comerciais | **EP-0002** — Gestão de Clientes e Assinaturas | Requisitos Funcionais §5 (Cadastro de planos comerciais) |
| **BR-04** | Vinculação de Assinaturas | **EP-0002** — Gestão de Clientes e Assinaturas | J1: Ativação de novo cliente · J3: Upgrade de plano |

---

> 📄 **Índice de Épicos:** [`03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Anterior:** [EP-0001 — Portal Administrativo Interno](../EP-0001-portal-administrativo-interno.md) | **Próximo:** [EP-0003 — Governança de Acessos e Permissões](../EP-0003-governanca-de-acessos-e-permissoes.md)

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: breakdown-epic-pm, agile-ba-practices. Estrutura modular v4.0.*

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE ÉPICOS]
