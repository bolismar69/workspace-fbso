# PROMPT-FIX-ARCHITECTURE-FROM-GATE

## Contexto

Este prompt é acionado quando o **Gate de Alinhamento de Escopo** (PROMPT-GATE-ARCHITECTURE-SCOPE) reprova o `ARCHITECTURE.md` e gera o relatório `ARCHITECTURE_SCOPE_FAIL_REPORT.md`.

O agente atua como **corretor de arquitetura** — lê o relatório de falha e aplica correções cirúrgicas no ARCHITECTURE.md, com foco especial em **simplificação** (remoção de complexidade desnecessária) e **alinhamento ao PRD.md**.

**Princípio fundamental:** Scope creep técnico é tão perigoso quanto scope creep funcional. A correção prioriza REMOVER complexidade sobre ADICIONAR componentes.

---

## Parâmetros de Entrada

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{STACK}` | Stack tecnológica principal | `Java 25 + Spring Boot + PostgreSQL` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar TODOS os 5 parâmetros.

### Passo 1 — Carregar Artefatos e Relatório de Falha

```
Ler obrigatoriamente:
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/ARCHITECTURE_SCOPE_FAIL_REPORT.md
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/ARCHITECTURE.md (a corrigir)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/PRD.md (baseline de escoco)
    └── {PROJECT_PATH}/TECHNICAL-PLAN.md
```

### Passo 2 — Processar Não-Conformidades por Prioridade

Extrair do `ARCHITECTURE_SCOPE_FAIL_REPORT.md`:

| Prioridade | Tipo de NC | Ação |
|---|---|---|
| P0 (Crítica) | Scope creep técnico (componentes/serviços sem feature no PRD.md) | **Remover** componente da arquitetura |
| P0 (Crítica) | Complexidade injustificada (Clean/Hexagonal sem justificativa, over-engineering) | **Simplificar** para package-by-layer |
| P1 (Alta) | Infraestrutura não justificada (Kafka, Redis, Elasticsearch sem requisito) | **Remover** ou **mover para "Futuro"** |
| P1 (Alta) | Entidade/componente sem correspondência no PRD.md | **Remover** ou **justificar no PRD.md** |
| P2 (Média) | Inconsistência com TECHNICAL-PLAN.md ou ADRs globais | **Alinhar** com documento de referência |
| P2 (Média) | Seção incompleta (pipeline de segurança, estratégia de testes) | **Complementar** seção |
| P3 (Baixa) | Documentação incompleta (ADR local, nomenclatura) | **Documentar** decisão ou padronizar nome |

### Passo 3 — Aplicar Correções no ARCHITECTURE.md

```
Para cada NC no relatório:
    │
    ├── NC de SIMPLIFICAÇÃO (Dimensão 2 — KISS)?
    │     → Executar skill ponytail na seção problemática
    │     → Se Clean Architecture sem justificativa → migrar estrutura para package-by-layer
    │     → Se padrão complexo sem justificativa → substituir por alternativa mais simples
    │     → Atualizar ADR de simplificação documentando a mudança
    │
    ├── NC de REMOÇÃO (scope creep técnico)?
    │     → Remover componente/serviço/integração não justificado
    │     → Atualizar diagrama de estrutura de pacotes
    │     → Atualizar lista de dependências
    │
    ├── NC de ALINHAMENTO (inconsistência)?
    │     → Consultar documento de referência (PRD.md, TECHNICAL-PLAN.md)
    │     → Ajustar stack, padrões ou nomenclatura
    │     → Atualizar ADRs locais se necessário
    │
    └── NC de COMPLEMENTO (incompleto)?
          → Identificar seção faltante ou incompleta
          → Gerar conteúdo seguindo padrão do ARCHITECTURE.md
          → Garantir coesão com restante do documento
```

### Passo 4 — Revisão Final com Ponytail

Após aplicar todas as correções, **obrigatoriamente** invocar o skill `ponytail` como revisor final:

```
Invocar ponytail sobre o ARCHITECTURE.md corrigido:
    - "Esta arquitetura está simples o suficiente?"
    - "Há algo mais que pode ser removido?"
    - "A complexidade restante é estritamente necessária?"

Se o ponytail sugerir novas simplificações → aplicar antes de finalizar.
```

### Passo 5 — Atualizar Registro de Alterações

```markdown
| v{X+1} | {data} | Correção pós-gate: {N} não-conformidades resolvidas do ARCHITECTURE_SCOPE_FAIL_REPORT.md v{Y}. NCs: {lista}. Simplificações aplicadas: {resumo}. | Agente Corretor Arquitetura/IA |
```

### Passo 6 — Validar Correções

| # | Verificação | Critério |
|---|---|---|
| 1 | Todas as NCs endereçadas | Cada NC do relatório tem correção ou justificativa |
| 2 | Scope creep técnico removido | Componentes/serviços não justificados foram removidos |
| 3 | Simplificações aplicadas | Complexidade desnecessária foi eliminada |
| 4 | Revisão do ponytail concluída | Skill ponytail executado e sugestões incorporadas |
| 5 | Consistência com PRD.md | Arquitetura cobre exatamente o escopo do PRD.md |
| 6 | Registro de alterações | Nova versão documentada |

---

## Skills Orquestradas

| Ordem | Skill | Propósito |
|---|---|---|
| 1ª | `ponytail` | Simplificar — remover complexidade desnecessária (executado no início E no final) |
| 2ª | `gap-analysis` | Analisar relatório de falha e planejar correções |
| 3ª | `architecture-patterns` | Garantir que correções usam padrões adequados ao escopo |
| 4ª | `engineering-skills` | Validar estrutura de pacotes, tratamento de erros, logging |
| 5ª | `documentation-writer` | Qualidade do ARCHITECTURE.md corrigido |

---

## Observações

1. **Simplificar é a ação padrão.** Na dúvida entre adicionar complexidade ou simplificar, SIMPLIFICAR. O skill `ponytail` é executado duas vezes: no início (para avaliar o estado atual) e no final (para revisar as correções).

2. **Remover > Mover para "Futuro".** Se um componente não é justificado pelo PRD.md atual, ele deve ser REMOVIDO, não movido para uma seção de "trabalho futuro". A seção "Futuro" só deve conter itens explicitamente mencionados no PRD.md como fases posteriores.

3. **Cada remoção deve ser documentada.** Quando um componente é removido, documentar no ADR de simplificação o motivo da remoção e a referência ao PRD.md que justifica a decisão.

4. **Incrementar versão SEMPRE.** Toda correção gera nova versão com entrada no registro de alterações.

5. **Após correção, reexecutar o gate.** FIX → GATE → (se APROVADO) prosseguir. Se o escopo foi alterado, reexecutar também o gate de PRD.md.

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 13/07/2026 | Criação inicial: correção de ARCHITECTURE.md baseada em ARCHITECTURE_SCOPE_FAIL_REPORT.md, foco em simplificação e remoção de scope creep técnico | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, architecture-patterns, ponytail, gap-analysis.*
