# PROMPT-FIX-SPECS-FROM-GATE

## Contexto

Este prompt é acionado quando o **Gate de Viabilidade e Completude Técnica** (PROMPT-GATE-SPECS-TECHNICAL) reprova o `TECHNICAL-SOLUTION-SPECS.md` e gera o relatório `TECHNICAL_SPECS_FAIL_REPORT.md`.

O agente atua como **corretor de especificações** — lê o relatório de falha e aplica correções cirúrgicas no TECHNICAL-SOLUTION-SPECS.md, priorizando itens **bloqueantes** (que impedem a geração de TECHNICAL-SOLUTION-TASKS.md) sobre não-bloqueantes.

**Princípio fundamental:** O TECHNICAL-SOLUTION-SPECS.md é a ponte negócio→técnico. Correções aqui impactam diretamente a qualidade das tarefas e dos testes downstream.

---

## Parâmetros de Entrada

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{SOLUTION_TYPE}` | Tipo da solução | `backend`, `frontend`, `batch`, `mobile` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar TODOS os 5 parâmetros.

### Passo 1 — Carregar Artefatos e Relatório de Falha

```
Ler obrigatoriamente:
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL_SPECS_FAIL_REPORT.md
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-SPECS.md (a corrigir)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-PRD.md
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-ARCHITECTURE.md
    └── {PROJECT_PATH}/04-FEATURES.md
```

### Passo 2 — Processar Não-Conformidades por Prioridade

Extrair do `TECHNICAL_SPECS_FAIL_REPORT.md`, com atenção especial ao §4 (Bloqueios Técnicos):

| Prioridade | Tipo de NC | Ação |
|---|---|---|
| P0 (Bloqueante) | Feature sem especificação suficiente para implementação | **Especificar** feature completamente |
| P0 (Bloqueante) | API sem contrato completo (faltando request/response schema) | **Completar** contrato da API |
| P0 (Bloqueante) | RN ambígua ou contraditória | **Clarificar** regra de negócio com exemplos e casos de borda |
| P1 (Alta) | Feature/US sem mapeamento BR (rastreabilidade quebrada) | **Adicionar** mapeamento na matriz |
| P1 (Alta) | NFR sem métrica objetiva ou método de verificação | **Adicionar** métrica quantificável e método de verificação |
| P2 (Média) | Critério de aceitação não testável ou sem evidência esperada | **Reescrever** critério como verificação objetiva |
| P2 (Média) | Inconsistência terminológica ou referência cruzada quebrada | **Corrigir** termo ou link |
| P3 (Baixa) | Glossário incompleto, premissas não documentadas | **Complementar** seção |

### Passo 3 — Aplicar Correções no TECHNICAL-SOLUTION-SPECS.md

```
Para cada NC no relatório:
    │
    ├── NC BLOQUEANTE (P0)?
    │     → Feature sem especificação:
    │         → Consultar feature no 04-FEATURES.md e TECHNICAL-SOLUTION-PRD.md
    │         → Escrever especificação completa: APIs, RNs, critérios de aceite
    │     → API sem contrato:
    │         → Definir método, path, RBAC, request schema, response schema, status codes
    │         → Documentar regras de validação por campo
    │     → RN ambígua:
    │         → Reescrever com: descrição formal, 3 exemplos, 3 casos de borda
    │
    ├── NC ALTA (P1)?
    │     → Completar rastreabilidade: adicionar entradas na matriz BR→Feature→US
    │     → Adicionar métricas aos NFRs: substituir "rápido" por "p95 < 200ms"
    │
    ├── NC MÉDIA (P2)?
    │     → Reescrever critérios de aceite: "funciona" → "API retorna 200 com corpo contendo X"
    │     → Corrigir termos inconsistentes e links quebrados
    │
    └── NC BAIXA (P3)?
          → Completar glossário com termos usados no documento
          → Documentar premissas e restrições implícitas
```

### Passo 4 — Atualizar Registro de Alterações

```markdown
| v{X+1} | {data} | Correção pós-gate: {N} não-conformidades resolvidas do TECHNICAL_SPECS_FAIL_REPORT.md v{Y}. {M} bloqueantes, {P} não-bloqueantes. NCs: {lista}. | Agente Corretor SPECS/IA |
```

### Passo 5 — Validar Correções

| # | Verificação | Critério |
|---|---|---|
| 1 | Bloqueantes resolvidos | Todas as NCs P0 do §4 do relatório foram corrigidas |
| 2 | Rastreabilidade restaurada | Matriz BR→Feature→US está completa |
| 3 | APIs completas | Endpoints têm request/response schema, RBAC e status codes |
| 4 | NFRs quantificáveis | Cada NFR tem métrica objetiva e método de verificação |
| 5 | Critérios testáveis | Critérios de aceite são objetivos e vinculados ao DoD |
| 6 | Consistência cross-documento | TECHNICAL-SOLUTION-SPECS.md alinhado com TECHNICAL-SOLUTION-PRD.md e TECHNICAL-SOLUTION-ARCHITECTURE.md |
| 7 | Registro de alterações | Nova versão documentada |

---

## Skills Orquestradas

| Ordem | Skill | Propósito |
|---|---|---|
| 1ª | `gap-analysis` | Analisar relatório de falha e priorizar correções |
| 2ª | `spec-miner` | Extrair especificações faltantes dos docs de negócio |
| 3ª | `acceptance-criteria` | Reescrever critérios de aceite não-testáveis |
| 4ª | `domain-modeling` | Refinar regras de negócio ambíguas |
| 5ª | `documentation-writer` | Qualidade e consistência do TECHNICAL-SOLUTION-SPECS.md corrigido |

---

## Observações

1. **Bloqueantes primeiro.** O §4 do relatório classifica NCs em bloqueantes vs não-bloqueantes. Resolver TODAS as bloqueantes antes de abordar as demais.

2. **Correções cirúrgicas, mantendo o escopo.** Não adicionar features, APIs ou regras de negócio que não estejam no TECHNICAL-SOLUTION-PRD.md. Se o TECHNICAL-SOLUTION-PRD.md estiver desatualizado, corrigi-lo primeiro (via PROMPT-FIX-PRD-FROM-GATE).

3. **Métricas de NFR são obrigatórias.** "Rápido", "seguro", "disponível" não são métricas. Substituir por valores quantificáveis: "p95 < 200ms", "RBAC com 4 papéis", "99.9% uptime".

4. **Incrementar versão SEMPRE.** Toda correção gera nova versão com entrada no registro de alterações.

5. **Após correção, reexecutar o gate.** FIX SPECS → GATE SPECS → (se APROVADO) prosseguir para TECHNICAL-SOLUTION-TASKS.md.

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 13/07/2026 | Criação inicial: correção de TECHNICAL-SOLUTION-SPECS.md baseada em TECHNICAL_SPECS_FAIL_REPORT.md, 4 níveis de prioridade, foco em itens bloqueantes | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, agile-ba-practices, spec-miner, gap-analysis.*
