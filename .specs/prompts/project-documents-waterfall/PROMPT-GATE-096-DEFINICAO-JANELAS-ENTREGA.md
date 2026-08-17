# PROMPT: PORTÃO DE VALIDAÇÃO DE DEFINIÇÃO DE JANELAS DE ENTREGA (096)
## Versão: 1.0 — WATERFALL Orchestrator v3.1 (6 Fases, 39 Documentos)

Atue como um Auditor de Qualidade de Documentação, especializado em Gestão de Entregas e metodologia WATERFALL.

## Inputs (recebidos explicitamente — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo do arquivo a ser validado |

## Regras

1. Leia **APENAS** o arquivo em `DOC_PATH` — não busque outros arquivos
2. Execute cada item do CHECKLIST abaixo contra o conteúdo do documento
3. Se TODOS os checks passarem: altere o status para `[STATUS: Em revisão]` e retorne `{PASS}`
4. Se houver falhas: NÃO altere o status; retorne `{FAIL, VIOLATIONS: [{section, description, severity}]}`

## Checklist de Compliance

1. **Cabeçalho e Metadados:** Projeto, Documentos Base (092, 085, 045, 050, 095, 105, 090, 087), Data, Versão e Metodologia preenchidos? Status é "Em análise" ou "Em revisão"?
2. **Cobertura das 4 Janelas:** DEV/QA/UAT/DEPLOY definidas com objetivo, na ordem correta do loop? — `[096-01]`
3. **Donos por Janela:** cada janela tem dono da execução e dono do aceite, respeitando "TECHLEAD propõe, PM/PO aplica" (QA aceite PM/PO via 095; UAT aceite por DE-ACORDO/APROVAÇÃO por entrega; DEPLOY go/no-go PM/PO)? — `[096-02]`
4. **Critérios de Entrada:** cada janela tem critérios de entrada explícitos, derivados dos upstreams (gate DEV, 095 GO, 105 assinado, GMUD)? — `[096-03]`
5. **Critérios de Saída (gate):** cada janela tem gate de saída verificável (PR+CI verde, 095 GO, DE-ACORDO/APROVAÇÃO da entrega, pós-deploy validado)? — `[096-04]`
6. **Matriz de Transição:** o loop DEV→QA→UAT→DEPLOY→DEV por FILA-NN está diagramado com as 3 tratativas de retorno (QA NO-GO, UAT divergência via 085, DEPLOY bloqueado via IMP-NN)? — `[096-05]`
7. **Rastreabilidade:** prefixos JAN-*-NN usados; registros apontam para 600/595/095/105; nenhum órfão? — `[096-06]`
8. **Evidências por Janela:** cada janela tem evidências esperadas com localização (paths/repositórios)? — `[096-07]`
9. **Vocabulário WATERFALL:** nenhum termo ágil (Sprint/User Story/DoR/Epic)? IDs apenas JAN-*-NN? — `[096-08]`
10. **Limite de Escopo (fronteira dupla):** o documento NÃO define Filas/Ciclos (`FILA-NN` — exclusivo do 092)? — `[096-09]`
11. **Alinhamento Upstream:** critérios são consistentes com 092/085/045/050/095/090/087 (ex.: GMUD em PROD, 095 GO como pré-requisito do UAT)? — `[096-10]`
12. **Consistência Interna:** janelas da Seção 1 são as mesmas da matriz (Seção 2) e da rastreabilidade (Seção 3)?
13. **105 fora das janelas:** o 105-TERMO-ACEITE NÃO é usado como gate por entrega — é citado apenas como aceite FINAL do projeto (FASE 6)? — `[096-11]`
