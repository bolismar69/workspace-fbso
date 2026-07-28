# User Story: US-FEAT-EP-0004-0002-0042 — cadastrar minha primeira Unidade de Negócio (CNPJ matriz, regime tribu

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0004](../epics/EP-0004-portal-do-cliente-auto-servico.md) ➔ Feature [FEAT-EP-0004-0002](../features/FEAT-EP-0004-0002-onboarding-guiado-de-primeiro-acesso.md) ➔ User Story US-FEAT-EP-0004-0002-0042
- **Prioridade:** Must Have
- **Data-Alvo:** 30/09/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Cliente no onboarding,
- **quero** cadastrar minha primeira Unidade de Negócio (CNPJ matriz, regime tributário, endereço) para começar a operar na plataforma,
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Cliente no onboarding está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Formulário de cadastro de Unidade de Negócio integrado ao fluxo**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Cliente no onboarding está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Campos: CNPJ, razão social, regime tributário (Simples, Lucro Real, Lucro Presumido), endereço**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Cliente no onboarding está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Validação de CNPJ informa se formato é válido (validação de dígitos)**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0004-0002 — Onboarding Guiado de Primeiro Acesso](../features/FEAT-EP-0004-0002-onboarding-guiado-de-primeiro-acesso.md) | **Épico:** [EP-0004 — Portal do Cliente Auto Servico](../epics/EP-0004-portal-do-cliente-auto-servico.md)
