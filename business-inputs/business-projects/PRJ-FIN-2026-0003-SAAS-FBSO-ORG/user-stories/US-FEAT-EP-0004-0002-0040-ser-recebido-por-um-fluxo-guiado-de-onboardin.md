# User Story: US-FEAT-EP-0004-0002-0040 — ser recebido por um fluxo guiado de onboarding que me conduza passo a 

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0004](../epics/EP-0004-portal-do-cliente-auto-servico.md) ➔ Feature [FEAT-EP-0004-0002](../features/FEAT-EP-0004-0002-onboarding-guiado-de-primeiro-acesso.md) ➔ User Story US-FEAT-EP-0004-0002-0040
- **Prioridade:** Must Have
- **Data-Alvo:** 30/09/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Cliente no primeiro acesso,
- **quero** ser recebido por um fluxo guiado de onboarding que me conduza passo a passo pelas configurações iniciais obrigatórias,
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Cliente no primeiro acesso está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Ao detectar primeiro login, sistema inicia automaticamente o onboarding**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Cliente no primeiro acesso está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Barra de progresso visível (Passo 1 de 4, Passo 2 de 4...)**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Cliente no primeiro acesso está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Não é possível pular etapas obrigatórias**.

### Cenário 4: [Fluxo Alternativo 4]
- **Dado que** Cliente no primeiro acesso está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Cliente pode salvar e continuar depois (retoma de onde parou)**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0004-0002 — Onboarding Guiado de Primeiro Acesso](../features/FEAT-EP-0004-0002-onboarding-guiado-de-primeiro-acesso.md) | **Épico:** [EP-0004 — Portal do Cliente Auto Servico](../epics/EP-0004-portal-do-cliente-auto-servico.md)
