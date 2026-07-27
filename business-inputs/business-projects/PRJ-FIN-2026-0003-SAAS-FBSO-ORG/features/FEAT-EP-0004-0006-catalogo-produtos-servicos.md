# FEATURE - FEAT-EP-0004-0006: Catálogo de Produtos e Serviços

| Campo | Detalhe |
|-------|---------|
| **Feature** | FEAT-EP-0004-0006 — Catálogo de Produtos e Serviços |
| **Épico** | [EP-0004 — Experiência do Cliente e Autoatendimento](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) |
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Versão** | 1.0 — Documento Inicial de Funcionalidades (Estrutura Modular v2.0) |
| **Data** | 26 de julho de 2026 |
| **Origem** | `03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` e `02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) | **Anterior:** [FEAT-EP-0004-0005 — Unidades de Negócio](../FEAT-EP-0004-0005-gestao-unidades-negocio.md)

**Requisitos BRD Vinculados:** [BR-10](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Catálogo de Produtos/Serviços

---

## Objetivo de Negócio
Permitir que o cliente cadastre e gerencie seu portfólio comercial (produtos e serviços), preparando a base de dados para o futuro mapeamento fiscal do Tributali-Engine.

**Prioridade:** Must Have

## User Stories

| # | User Story | Critérios de Aceitação |
|---|-----------|----------------------|
| US-055 | Como **Cliente**, quero cadastrar um novo produto ou serviço informando nome, SKU/código interno, tipo (Produto ou Serviço) e descrição | • Formulário com campos: nome (obrigatório), SKU (opcional), tipo (obrigatório: Produto ou Serviço), descrição (opcional) • Item criado com status "Ativo" por padrão • Vinculado automaticamente à Unidade de Negócio ativa no seletor |
| US-056 | Como **Cliente**, quero visualizar a lista de produtos e serviços cadastrados para minha Unidade de Negócio, com busca por nome ou SKU | • Lista filtrada pela Unidade de Negócio selecionada no seletor • Campo de busca textual que filtra por nome ou SKU • Colunas: Nome, SKU, Tipo, Status, Indicador de Mapeamento Fiscal |
| US-057 | Como **Cliente**, quero editar as informações de um produto ou serviço (nome, SKU, tipo, descrição) para manter o catálogo atualizado | • Tela de edição acessível a partir da lista • Todos os campos do cadastro são editáveis • Alterações aplicadas imediatamente após salvar |
| US-058 | Como **Cliente**, quero ativar ou desativar um produto do catálogo sem excluí-lo definitivamente, para controlar quais itens estão em uso | • Botão "Desativar" na lista (para itens ativos) • Botão "Ativar" na lista (para itens inativos) • Item desativado não aparece em cadastros futuros, mas mantém histórico |

## Regras de Negócio

- **RN-FEAT-EP-0004-0006-0001:** Catálogo é segmentado por Unidade de Negócio — cada unidade tem seu próprio catálogo
- **RN-FEAT-EP-0004-0006-0002:** SKU é opcional, mas se informado deve ser único por Unidade de Negócio
- **RN-FEAT-EP-0004-0006-0003:** Indicador de "Mapeamento Fiscal" nesta fase exibe "Não mapeado" para todos os itens
- **RN-FEAT-EP-0004-0006-0004:** Exclusão de produtos segue política de soft delete (desativação lógica, não remoção física)

---

## Matriz de Rastreabilidade BRD → Épico/Jornada → Esta Feature

| BRD | Requisito Funcional | Épico/Jornada | Esta Feature |
|:---|:---|:---|:---|
| **BR-10** | Catálogo de Produtos/Serviços | [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md) / J3: Cadastro de portfólio de produtos | **FEAT-EP-0004-0006** — Catálogo de Produtos e Serviços |

---

> 📄 **Índice de Features:** [`04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | **Épico:** [EP-0004](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md)

[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE FEATURES]
