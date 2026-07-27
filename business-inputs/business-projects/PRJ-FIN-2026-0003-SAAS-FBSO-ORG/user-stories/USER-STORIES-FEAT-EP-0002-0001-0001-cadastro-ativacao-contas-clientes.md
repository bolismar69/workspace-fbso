# User Stories: Cadastro e Ativação de Contas de Clientes

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** FEAT-EP-0002-0001 — Cadastro e Ativação de Contas de Clientes
- **Épico:** EP-0002 — Gestão de Clientes e Assinaturas
- **Prioridade:** Must Have
- **Data-Alvo:** 31/08/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

---

## Objetivo de Negócio

Permitir que o time interno crie contas de clientes de forma estruturada, com todos os dados corporativos necessários, gerando automaticamente o convite de acesso ao portal.

---

## User Stories

### US-008 — Criação de Nova Conta de Cliente (Tenant)

**Como** Administrador FBSO, **quero** criar uma nova conta de cliente (Tenant) preenchendo razão social, nome fantasia e segmento de mercado **para** registrar o cliente na plataforma.

**Critérios de Aceitação:**
- Formulário com campos obrigatórios: razão social, segmento. Segmento selecionado de lista predefinida: Varejo, Atacado, Serviços, Indústria, Tecnologia, Outros.
- Campos opcionais: nome fantasia, observações
- Ao salvar, Tenant é criado com status "Pendente Onboarding"
- Sistema gera link único de ativação para o cliente

### US-009 — Envio Automático de E-mail de Boas-Vindas

**Como** Administrador FBSO, **quero** que o sistema envie automaticamente um e-mail de boas-vindas ao cliente com o link de ativação da conta após a criação do Tenant.

**Critérios de Aceitação:**
- E-mail disparado automaticamente após criação do Tenant
- E-mail contém link único e instruções de primeiro acesso
- Link expira em 7 dias. Administrador pode gerar novo link via botão 'Reenviar Convite' (US-011), que cria novo link com novo prazo de 7 dias.

### US-010 — Edição de Dados Cadastrais do Cliente

**Como** Administrador FBSO, **quero** editar os dados cadastrais de um cliente (razão social, nome fantasia, segmento) **para** manter as informações sempre atualizadas.

**Critérios de Aceitação:**
- Tela de edição acessível a partir da lista de contas
- Campos editáveis: razão social, nome fantasia, segmento, observações
- Alterações registradas no histórico de auditoria

### US-011 — Reenvio de E-mail de Ativação

**Como** Administrador FBSO, **quero** reenviar o e-mail de ativação caso o cliente não tenha recebido ou o link tenha expirado.

**Critérios de Aceitação:**
- Botão "Reenviar Convite" na tela de detalhes do Tenant
- Disponível apenas para contas com status "Pendente Onboarding"
- Gera novo link e novo prazo de 7 dias

---

## Regras de Negócio

| ID | Regra |
|----|-------|
| **RN-FEAT-EP-0002-0001-0001** | Toda criação de Tenant gera registro de auditoria com: administrador responsável, data/hora, dados iniciais |
| **RN-FEAT-EP-0002-0001-0002** | Razão social é validada como obrigatória; sistema alerta se já existir Tenant ativo com mesma razão social |
| **RN-FEAT-EP-0002-0001-0003** | Link de ativação é único e de uso único por Tenant |

---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | Tenant criado com status "Pendente Onboarding" e link gerado | Registro no sistema + e-mail recebido |
| F2 | Edição de dados com registro de auditoria | Registro no histórico |
| F3 | Reenvio de convite funcional apenas para status Pendente | Teste com tenant em cada status |

---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*
