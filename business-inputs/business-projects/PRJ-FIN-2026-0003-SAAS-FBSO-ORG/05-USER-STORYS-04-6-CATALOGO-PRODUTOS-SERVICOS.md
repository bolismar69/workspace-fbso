# User Stories: Catálogo de Produtos e Serviços

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** F04-06 — Catálogo de Produtos e Serviços
- **Épico:** EP-04 — Experiência do Cliente e Autoatendimento
- **Prioridade:** Must Have
- **Data-Alvo:** 15/10/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES.md](../04-FEATURES.md)

---

## Objetivo de Negócio

Permitir que o cliente cadastre e gerencie seu portfólio comercial (produtos e serviços), preparando a base de dados para o futuro mapeamento fiscal do Tributali-Engine.

---

## User Stories

### US-055 — Cadastro de Novo Produto ou Serviço

**Como** Cliente, **quero** cadastrar um novo produto ou serviço informando nome, SKU/código interno, tipo (Produto ou Serviço) e descrição.

**Critérios de Aceitação:**
- Formulário com campos: nome (obrigatório), SKU (opcional), tipo (obrigatório: Produto ou Serviço), descrição (opcional)
- Item criado com status "Ativo" por padrão
- Vinculado automaticamente à Unidade de Negócio selecionada no Seletor de Unidade de Negócio (US-054)

### US-056 — Lista de Produtos com Busca

**Como** Cliente, **quero** visualizar a lista de produtos e serviços cadastrados para minha Unidade de Negócio, com busca por nome ou SKU.

**Critérios de Aceitação:**
- Lista filtrada pela Unidade de Negócio selecionada no seletor
- Campo de busca textual que filtra por nome ou SKU
- Colunas: Nome, SKU, Tipo, Status, Indicador de Mapeamento Fiscal

### US-057 — Edição de Produto ou Serviço

**Como** Cliente, **quero** editar as informações de um produto ou serviço (nome, SKU, tipo, descrição) **para** manter o catálogo atualizado.

**Critérios de Aceitação:**
- Tela de edição acessível a partir da lista
- Nome, SKU, descrição e status são editáveis. Tipo (Produto/Serviço) NÃO pode ser alterado após criação, devido a implicações para mapeamento fiscal futuro (NCM para produtos, NBS para serviços).
- Alterações aplicadas imediatamente após salvar

### US-058 — Ativação e Desativação de Produto

**Como** Cliente, **quero** ativar ou desativar um produto do catálogo sem excluí-lo definitivamente, **para** controlar quais itens estão em uso.

**Critérios de Aceitação:**
- Botão "Desativar" na lista (para itens ativos)
- Botão "Ativar" na lista (para itens inativos)
- Item desativado não aparece em cadastros futuros, mas mantém histórico

---

## Regras de Negócio

| ID | Regra |
|----|-------|
| **RN18-01** | Catálogo é segmentado por Unidade de Negócio — cada unidade tem seu próprio catálogo |
| **RN18-02** | SKU é opcional, mas se informado deve ser único por Unidade de Negócio |
| **RN18-03** | Indicador de "Mapeamento Fiscal" nesta fase exibe "Não mapeado" para todos os itens (placeholder para integração futura com Tributali-Engine) |
| **RN18-04** | Exclusão de produtos segue política de soft delete (desativação lógica, não remoção física) |

---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | Produto cadastrado e visível na lista da Unidade de Negócio correta | Cadastro em BU-A e verificação de ausência em BU-B |
| F2 | Busca por nome e SKU funcional | Teste com termo parcial |
| F3 | Indicador de Mapeamento Fiscal exibe "Não mapeado" para todos os itens | Verificação visual na lista |
| F4 | Desativação preserva histórico e remove item dos seletores | Verificação pós-desativação |

## Casos de Borda

- Catálogo vazio exibe mensagem 'Nenhum produto cadastrado' com botão 'Cadastrar Primeiro Produto'.
- Busca sem resultados exibe 'Nenhum produto encontrado para [termo buscado].'


---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*

---
👷 *Revisão técnica realizada pelo Agente: CaveMan em 15/07/2026, conforme User Story Review. Foram utilizados os skills: caveman-review.*
