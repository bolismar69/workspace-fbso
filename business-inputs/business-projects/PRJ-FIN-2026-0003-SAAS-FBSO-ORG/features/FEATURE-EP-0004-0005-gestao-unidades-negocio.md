# FEATURE - EP-0004-0005: Gestão de Unidades de Negócio

| Campo | Detalhe |
|-------|---------|
| **Feature** | EP-0004-0005 — Gestão de Unidades de Negócio |
| **Épico** | [EP-0004 — Experiência do Cliente e Autoatendimento](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) | **Anterior:** [EP-0004-0004 — App Switcher](../FEATURE-EP-0004-0004-app-switcher-seletor-modulos.md) | **Próximo:** [EP-0004-0006 — Catálogo](../FEATURE-EP-0004-0006-catalogo-produtos-servicos.md)

**Requisitos BRD Vinculados:** [BR-09](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Cadastro de Unidades de Negócio

---

## Objetivo de Negócio
Permitir que o cliente cadastre e gerencie suas filiais e CNPJs de forma autônoma, com estrutura hierárquica clara (Matriz/Filial).

**Prioridade:** Must Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-050 | Como **Cliente**, quero visualizar a lista das minhas Unidades de Negócio organizadas hierarquicamente (Matriz no topo, filiais recuadas abaixo) para entender a estrutura da minha empresa | • Visualização em cards ou lista com recuo visual para filiais • Cada card exibe: razão social, CNPJ (mascarado), regime tributário, status • Indicador visual de Matriz vs. Filial |
| US-051 | Como **Cliente**, quero cadastrar uma nova Unidade de Negócio (filial) informando CNPJ, razão social, regime tributário e definindo a qual unidade ela se vincula (Matriz ou outra filial) | • Formulário com campos obrigatórios: CNPJ, razão social, regime tributário • Seletor de unidade pai (Matriz ou filial existente) • Validação de duplicidade de CNPJ ativo para o mesmo tenant |
| US-052 | Como **Cliente**, quero editar os dados de uma Unidade de Negócio (razão social, regime tributário, endereço) para manter as informações atualizadas | • Tela de edição acessível a partir do card da unidade • Campos editáveis: razão social, regime tributário, endereço • CNPJ não pode ser alterado após o cadastro |
| US-053 | Como **Cliente**, quero desativar uma Unidade de Negócio que não está mais em operação, mantendo seu histórico no sistema | • Botão "Desativar" no card da unidade • Confirmação exigida antes da desativação • Unidade desativada não aparece nos seletores para novos cadastros • Dados históricos permanecem acessíveis para consulta |
| US-054 | Como **Cliente**, quero usar o seletor de Unidade de Negócio no topo do portal para alternar entre minhas filiais e visualizar os dados específicos de cada uma | • Seletor dropdown no topo do portal, ao lado do App Switcher • Exibe apenas unidades que o usuário tem permissão para acessar • Ao trocar de unidade, dados exibidos nas telas são filtrados automaticamente |

## Regras de Negócio

- **RN-FEAT-EP-0004-0005-0001:** CNPJ deve ser único entre Unidades de Negócio ativas do mesmo tenant
- **RN-FEAT-EP-0004-0005-0002:** Uma unidade desativada não pode ser definida como "pai" de novas filiais
- **RN-FEAT-EP-0004-0005-0003:** A primeira unidade cadastrada (durante onboarding) é automaticamente a Matriz
- **RN-FEAT-EP-0004-0005-0004:** Não há limite de níveis hierárquicos (Matriz → Filial → Sub-filial)
- **RN-FEAT-EP-0004-0005-0005:** Seletor de Unidade de Negócio reflete apenas as unidades que o usuário tem permissão

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-09** | Cadastro de Unidades de Negócio | [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) / J2: Cadastro de filiais | **EP-0004-0005** — Gestão de Unidades de Negócio |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md)

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
