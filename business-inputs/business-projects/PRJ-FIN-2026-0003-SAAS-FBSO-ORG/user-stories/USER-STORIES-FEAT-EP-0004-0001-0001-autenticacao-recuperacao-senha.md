# User Stories: Autenticação e Recuperação de Senha

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** FEAT-EP-0004-0001 — Autenticação e Recuperação de Senha
- **Épico:** EP-0004 — Experiência do Cliente e Autoatendimento
- **Prioridade:** Must Have
- **Data-Alvo:** 30/09/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

---

## Objetivo de Negócio

Oferecer ao cliente uma experiência segura e fluida de acesso ao portal, com fluxo de recuperação de senha autônomo.

---

## User Stories

### US-037 — Login no Portal

**Como** Cliente, **quero** fazer login no portal usando meu e-mail e senha **para** acessar minha conta.

**Critérios de Aceitação:**
- Tela de login com campos: e-mail e senha
- Mensagens de erro: 'E-mail ou senha inválidos.' (genérica, não revela existência do e-mail). 'Conta temporariamente bloqueada. Tente novamente em X minutos.' (quando conta bloqueada por tentativas). 'Sua conta foi suspensa. Entre em contato com o suporte.' (quando tenant ou usuário suspenso).
- Após login bem-sucedido, redirecionar para o dashboard ou onboarding (se primeiro acesso)

### US-038 — Recuperação de Senha

**Como** Cliente, **quero** recuperar minha senha caso eu a esqueça, recebendo um link de redefinição por e-mail.

**Critérios de Aceitação:**
- Opção "Esqueci minha senha" na tela de login
- Usuário informa e-mail e recebe link de redefinição
- Link expira em 1 hora
- Nova senha deve atender critérios mínimos de complexidade

### US-039 — Bloqueio por Tentativas Consecutivas com Erro

**Como** Cliente, **quero** que ao errar a senha repetidas vezes, minha conta seja temporariamente bloqueada por segurança.

**Critérios de Aceitação:**
- Após 5 tentativas consecutivas com erro, conta bloqueada por 15 minutos
- Mensagem informa o tempo restante de bloqueio
- Administrador do tenant pode desbloquear usuário na tela de detalhes do usuário (US-025), via botão 'Desbloquear Acesso'.

---

## Regras de Negócio

| ID | Regra |
|----|-------|
| **RN-FEAT-EP-0004-0001-0001** | Senha deve ter no mínimo 8 caracteres, incluindo letra e número |
| **RN-FEAT-EP-0004-0001-0002** | Sessão expira após 60 minutos de inatividade |
| **RN-FEAT-EP-0004-0001-0003** | Link de redefinição de senha é de uso único |

---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | Login funcional com redirecionamento correto (dashboard vs. onboarding) | Teste com conta nova e conta existente |
| F2 | Fluxo completo de recuperação de senha (solicitação → e-mail → redefinição → novo login) | Percurso completo executado |
| F3 | Bloqueio após 5 tentativas e desbloqueio após 15 minutos ou via admin | Log de tentativas e timestamps |

---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*

---
👷 *Revisão técnica realizada pelo Agente: CaveMan em 15/07/2026, conforme User Story Review. Foram utilizados os skills: caveman-review.*
