# PROMPT-GENERATE-595-TECHLEAD-RETURN-PACKAGE

## Contexto

Este prompt gera o artefato `595-RETURN-PACKAGE-{FILA-NN}.md` 🆕 — o **pacote de retorno do TECHLEAD ao PM/PO** ao final de cada ciclo de entrega da FASE 5 do WATERFALL. Ele consolida, em um único documento estruturado, tudo o que o TECHLEAD **propõe** ao PM/PO: status dos itens, débito técnico, evidências, impedimentos e pedidos de mudança.

**Regras de negócio (ownership — "TECHLEAD propõe, PM/PO aplica"):**
- Este prompt **NUNCA edita** os arquivos do PM/PO: 092-BACKLOG-KANBAN, 093-GESTAO-TIMES, 095-RELATORIO-QUALIDADE, 085-PLANO-GESTAO-MUDANCAS e 088-PRODUCT-BACKLOG-LIST. Ele apenas gera o pacote com **propostas**
- O PM/PO aplica o pacote via `PROMPT-GENERATE-092-BACKLOG-KANBAN.md` em modo atualização (com o pacote em `EXTRA_INPUTS`), revalida com o GATE-092 e registra impedimentos (`IMP-NN`) no 093
- Vocabulário WATERFALL obrigatório (veto a termos ágeis: Sprint, User Story, DoR, Epic)
- Status propostos seguem a máquina de estados do 092: `A Fazer → Em Execução → Em Revisão → Concluído/Impedido`

**Papel no Bloco E (Esteira de Construção — modo waterfall-discovery):** passo final de cada ciclo. Consome os artefatos técnicos do ciclo e entrega o pacote ao PM/PO.

**Inputs upstream:** snapshot do 092-BACKLOG-KANBAN recebido no pacote de demanda, 580-SPRINT-BACKLOG, `590-sprint-NNN/` (5 contratos), `SPRINT-REVIEW.md`, `TASK-EXECUTED`, `IDENTIFIED-TECHNICAL-DEBT.md`, relatório do `PROMPT-EXECUTE-SPRINT-TASKS.md`, 600-EXECUTION-HISTORY.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{FILA_NN}` | Identificador da fila/ciclo de entrega do 092 (ex.: `FILA-01`) |
| `{SNAPSHOT_092}` | Caminho do snapshot do 092-BACKLOG-KANBAN recebido no pacote de demanda (NUNCA inferir — vem do handoff PM/PO) |
| `{SPRINT_NUMBER}` | Número do `590-sprint-NNN` correspondente à `{FILA_NN}` (regra: `590-sprint-NN ↔ FILA-NN`) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Verificar se TODOS os parâmetros obrigatórios foram informados, em especial `{FILA_NN}` e `{SNAPSHOT_092}`.

### Passo 1 — Carregar Documentos Base
Ler `{SNAPSHOT_092}` (itens `BL-NN` da fila, `CR-NN` aprovadas), `580-SPRINT-BACKLOG.md` (tarefas `T-NNN`), `technical-discovery/590-sprint-{SPRINT_NUMBER}/` (5 contratos), `SPRINT-REVIEW.md`, `TASK-EXECUTED`, `IDENTIFIED-TECHNICAL-DEBT.md`, o relatório do `PROMPT-EXECUTE-SPRINT-TASKS.md` e `600-EXECUTION-HISTORY.md` (estado do pipeline). NÃO ler ou buscar outros arquivos.

### Passo 2 — Invocar Skills Especializadas
Invocar skills de gestão de execução, gap analysis e documentação.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/595-RETURN-PACKAGE-{FILA_NN}.md` com **7 seções**:

1. **Identificação do Ciclo** — `FILA-NN`, solução técnica, `590-sprint-NNN`, datas de início/fim do ciclo
2. **Status Proposto por Item** — tabela: `BL-NN` (do snapshot 092), status proposto (Em Revisão/Concluído/Impedido), % concluído, tarefas `T-NNN` executadas, **Janela atual (`JAN-DEV-NN`/`JAN-QA-NN`/`JAN-UAT-NN`/`JAN-DEPLOY-NN`, coerente com o 600)** e **Aceite UAT (DE-ACORDO/APROVAÇÃO por entrega — Key Users + PM/PO)** para itens que saíram da UAT, observações
3. **Débito Técnico** — `DT-XXX` (fonte: `IDENTIFIED-TECHNICAL-DEBT.md`/`SPRINT-REVIEW.md`): descrição, impacto, tarefas afetadas + **proposta de CR Técnico** (`CR-NN` com justificativa para o processo do 085)
4. **Evidências para o 095** — caminhos absolutos, **agrupadas por janela** (DEV: PRs/relatórios; QA: resultados de testes funcionais/carga/pentest; UAT: registro de DE-ACORDO por entrega; DEPLOY: runbooks do 090 + evidências CI/CD do 087): `TASK-EXECUTED`, `SPRINT-REVIEW.md`, PRs (nº e link), relatórios de execução, resultados de testes
5. **Impedimentos Propostos** — `IMP-NN`: nome, descrição, impacto no ciclo, solução sugerida (o PM/PO registra no 093)
6. **Pedidos de Mudança** — divergências vs baseline M4: escopo (processo 085), cronograma (PERT/065/070), stack
7. **Recomendação de Próxima FILA-NN** — priorização sugerida (base: 580 + capacidade do 460) — **não vinculante**

### Passo 4 — Validação Pós-Geração
Verificar: 100% dos `BL-NN` do snapshot com status proposto; coluna "Janela" preenchida e coerente com o 600-EXECUTION-HISTORY; DE-ACORDO presente para itens que saíram da UAT; todo `DT-XXX` com CR Técnico vinculada; todas as evidências com caminhos absolutos existentes, agrupadas por janela; nenhum arquivo do PM/PO editado; vocabulário WATERFALL.

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados para esta fase. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `delivery-manager` | Consolidação da entrega do ciclo | Gestão |
| 2 | `gap-analysis` | Débito técnico e pedidos de mudança | Análise |
| 3 | `backlog-management` | Status e priorização por BL-NN/FILA-NN | Gestão |
| 4 | `documentation-writer` | Redigir o pacote de retorno | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog do artefato.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 16/08/2026 | Criação inicial: pacote de retorno TECHLEAD→PM/PO (Bloco E — modo waterfall-discovery) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Skills de referência listados acima. Outros skills podem ser utilizados conforme aderência.*
