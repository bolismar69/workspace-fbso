# PROMPT-FIX-595-TECHLEAD-RETURN-PACKAGE

## Contexto

Este prompt implementa o **corretor cirúrgico do Pacote de Retorno do TECHLEAD** para o artefato `595-RETURN-PACKAGE-{FILA-NN}.md` (Bloco E — modo waterfall-discovery). Ele corrige APENAS os pontos conflitantes reportados pelo GATE, sem regerar o documento.

**Princípio fundamental:** correção cirúrgica — editar somente as seções com violações reportadas, preservando o que passou no GATE e mantendo o contrato "TECHLEAD propõe, PM/PO aplica".

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{FILA_NN}` | Identificador da fila/ciclo de entrega do 092 |
| `{GATE_REPORT}` | Relatório de auditoria do GATE-595 com os conflitos `[595-NN]` e as respostas do humano |

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `595-RETURN-PACKAGE-{FILA_NN}.md` (artefato a corrigir) e `{GATE_REPORT}` (conflitos + respostas do humano).

### Passo 2 — Corrigir Apenas as Violações

1. Edite **APENAS** as seções listadas nos conflitos `[595-NN]` do relatório do GATE
2. **NÃO** recrie, regenere ou reescreva o documento inteiro
3. **NÃO** altere seções que passaram no GATE
4. **NÃO** edite arquivos do PM/PO (092/093/095/085/088) — violação de ownership `[595-07]`
5. Mantenha o vocabulário WATERFALL (veto a termos ágeis — `[595-06]`)
6. Após cada correção, adicione comentário inline `<!-- FIX: [595-NN] — corrigido -->` na seção reparada
7. Máximo de 3 loops de correção; se persistir após 3 loops, parar e acionar o humano

### Passo 3 — Re-Gate
Solicitar nova execução do `PROMPT-GATE-595-TECHLEAD-RETURN-PACKAGE.md`.

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Correção de vínculos e rastreabilidade | Análise |
| 2 | `requirements-validation` | Revalidar cobertura do snapshot 092 | Requisitos |
| 3 | `documentation-writer` | Redigir as correções no pacote | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no relatório.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 16/08/2026 | Criação inicial: corretor cirúrgico do pacote de retorno (Bloco E — modo waterfall-discovery) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
