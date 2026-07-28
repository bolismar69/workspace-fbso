# User Story: US-FEAT-EP-0002-0003-0017 — editar um plano existente (nome, preço, módulos) mantendo o histórico 

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Mapeamento Ágil:** Épico [EP-0002](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md) ➔ Feature [FEAT-EP-0002-0003](../features/FEAT-EP-0002-0003-configuracao-de-planos-comerciais.md) ➔ User Story US-FEAT-EP-0002-0003-0017
- **Prioridade:** Must Have
- **Data-Alvo:** 31/08/2026
- **Versão:** 1.0 — Especificação Modular Base
- **Status:** NEW — Aguardando refinamento de negócio

---

## 1. Declaração da História (User Story Statement)

- **Como** Gestor de Produto,
- **quero** editar um plano existente (nome, preço, módulos) mantendo o histórico de versões anteriores,
- **para** atender a necessidade de negocio descrita.

---

## 2. Cenários Comportamentais de Aceite (Gherkin Format)

### Cenário 1: [Fluxo Principal]
- **Dado que** Gestor de Produto está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Edição de plano gera nova versão**.

### Cenário 2: [Fluxo Alternativo 2]
- **Dado que** Gestor de Produto está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Clientes já vinculados permanecem na versão contratada até upgrade**.

### Cenário 3: [Fluxo Alternativo 3]
- **Dado que** Gestor de Produto está autenticado(a) no portal com as permissões adequadas,
- **Quando** executa a ação correspondente a esta funcionalidade,
- **Então** o sistema deve: **Não é possível excluir plano, apenas desativá-lo**.

---

## 3. Regras de Negócio de Tela Relacionadas

— (herdadas da feature)

---

> 📄 **Índice RTM:** [05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../05-USER-STORIES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Feature:** [FEAT-EP-0002-0003 — Configuração de Planos Comerciais](../features/FEAT-EP-0002-0003-configuracao-de-planos-comerciais.md) | **Épico:** [EP-0002 — Gestao de Clientes e Assinaturas](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md)
