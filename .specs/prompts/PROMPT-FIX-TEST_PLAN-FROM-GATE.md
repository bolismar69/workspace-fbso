# PROMPT-FIX-TEST_PLAN-FROM-GATE

## Contexto

Este prompt é acionado quando o **Gate de Viabilidade e Completude Técnica** (PROMPT-GATE-TEST_PLAN-TECHNICAL) reprova o `TEST_PLAN.md` e gera o relatório `TECHNICAL_TEST_PLAN_FAIL_REPORT.md`.

O agente atua como **corretor de plano de testes** — lê o relatório de falha e aplica correções no TEST_PLAN.md, com foco em **cobertura de features**, **qualidade dos cenários** e **cobertura de segurança**.

**Princípio fundamental:** Um plano de testes com gaps resulta em bugs não detectados. A correção garante que cada feature tem cenários adequados em todos os níveis da pirâmide.

---

## Parâmetros de Entrada

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{STACK}` | Stack tecnológica (para ferramentas de teste) | `Java 25 + Spring Boot + PostgreSQL` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar TODOS os 5 parâmetros.

### Passo 1 — Carregar Artefatos e Relatório de Falha

```
Ler obrigatoriamente:
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL_TEST_PLAN_FAIL_REPORT.md
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TEST_PLAN.md (a corrigir)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/SPECS.md
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TASKS.md
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/ARCHITECTURE.md
    └── {PROJECT_PATH}/04-FEATURES.md (critérios de aceitação)
```

### Passo 2 — Processar Não-Conformidades por Prioridade

Extrair do `TECHNICAL_TEST_PLAN_FAIL_REPORT.md`, com atenção ao §4 (Mapa de Cobertura) e §5 (Gaps de Segurança):

| Prioridade | Tipo de NC | Ação |
|---|---|---|
| P0 (Bloqueante) | Feature sem nenhum cenário de teste (0% cobertura) | **Criar** cenários em todos os níveis aplicáveis |
| P0 (Bloqueante) | Feature sem cenários de segurança (RBAC, Multi-Tenant) | **Adicionar** cenários de segurança |
| P1 (Alta) | Feature com cobertura parcial (só unitário, sem integração) | **Completar** níveis faltantes da pirâmide |
| P1 (Alta) | Cenário com estrutura incompleta (sem pré-condição ou resultado) | **Completar** estrutura do cenário |
| P2 (Média) | Pirâmide desbalanceada (muitos E2E, poucos unitários) | **Adicionar** cenários nos níveis subrepresentados |
| P2 (Média) | Cenário com resultado esperado ambíguo | **Reescrever** resultado como verificação objetiva |
| P3 (Baixa) | Ferramenta de teste incompatível com a stack | **Substituir** ferramenta |
| P3 (Baixa) | Cenário sem status ou sem agrupamento por feature | **Organizar** e **classificar** |

### Passo 3 — Aplicar Correções no TEST_PLAN.md

```
Para cada NC no relatório:
    │
    ├── NC de COBERTURA ZERO (P0)?
    │     → Feature sem cenários:
    │         → Consultar feature no SPECS.md e 04-FEATURES.md
    │         → Criar cenários em TODOS os níveis aplicáveis:
    │             - Unitário: testar lógica de negócio, validações, edge cases
    │             - Integração: testar API, persistência, integrações
    │             - E2E: testar fluxo completo do usuário
    │             - Segurança: testar RBAC (autorizado + negado), tenant isolation
    │         → Adicionar ao marco de teste correspondente no TASKS.md
    │
    ├── NC de SEGURANÇA (P0)?
    │     → Para cada endpoint da feature:
    │         → Cenário RBAC: papel autorizado acessa (200/201)
    │         → Cenário RBAC: papel NÃO autorizado é rejeitado (403)
    │         → Cenário Multi-Tenant: Tenant A não acessa dados do Tenant B
    │     → Se aplicável: cenários OWASP (SQL Injection, XSS, CSRF)
    │
    ├── NC de COMPLETUDE (P1)?
    │     → Feature com testes unitários mas sem integração:
    │         → Adicionar cenário de integração: chamada HTTP completa
    │     → Cenário incompleto:
    │         → Adicionar pré-condição: "Usuário autenticado como ADMIN no tenant X"
    │         → Adicionar resultado esperado: "HTTP 200 com corpo JSON contendo 'id', 'name'"
    │         → Adicionar passos numerados
    │
    ├── NC de PIRÂMIDE (P2)?
    │     → Muitos E2E, poucos unitários → adicionar cenários unitários por serviço
    │     → Muitos unitários, zero E2E → adicionar 1-2 cenários E2E para fluxos críticos
    │
    └── NC de QUALIDADE (P2/P3)?
          → Reescrever resultado ambíguo: "funciona" → "retorna lista paginada com 10 itens"
          → Substituir ferramenta incompatível
          → Organizar cenários por feature e atualizar status
```

### Passo 4 — Atualizar Mapa de Cobertura

Gerar/atualizar tabela de cobertura no próprio TEST_PLAN.md:

```markdown
## X. Mapa de Cobertura (Feature × Nível de Teste)
| Feature | Unit | Integração | E2E | Segurança | Cobertura |
|---|---|---|---|---|---|
| F01-01 | 3 | 2 | 1 | 2 | 100% |
| F02-03 | 2 | 1 | 0 | 2 | 75% |
```

### Passo 5 — Atualizar Registro de Alterações

```markdown
| v{X+1} | {data} | Correção pós-gate: {N} não-conformidades resolvidas do TECHNICAL_TEST_PLAN_FAIL_REPORT.md v{Y}. {C} cenários criados, {S} cenários de segurança adicionados, {R} cenários reescritos. Cobertura geral: {antiga}% → {nova}%. | Agente Corretor TEST_PLAN/IA |
```

### Passo 6 — Validar Correções

| # | Verificação | Critério |
|---|---|---|
| 1 | Cobertura total | Nenhuma feature com 0% de cobertura |
| 2 | Segurança coberta | Toda feature com endpoints tem cenários RBAC + Multi-Tenant |
| 3 | Pirâmide balanceada | Proporção aproximada: 60% unit, 25% int, 10% E2E, 5% sec |
| 4 | Cenários completos | Todo cenário tem: ID, descrição, nível, pré-condição, passos, resultado |
| 5 | Resultados objetivos | Nenhum resultado esperado ambíguo ("funciona", "está correto") |
| 6 | Ferramentas compatíveis | Ferramentas listadas são compatíveis com {STACK} |
| 7 | Registro de alterações | Nova versão documentada com métricas de cobertura |

---

## Skills Orquestradas

| Ordem | Skill | Propósito |
|---|---|---|
| 1ª | `test-strategy-design` | Rebalancear pirâmide e estratégia de testes |
| 2ª | `qa-test-planner` | Criar cenários de teste para features com cobertura zero |
| 3ª | `security-reviewer` | Adicionar cenários de segurança (RBAC, Multi-Tenant, OWASP) |
| 4ª | `acceptance-criteria` | Alinhar cenários com critérios de aceite |
| 5ª | `gap-analysis` | Analisar relatório e priorizar correções |
| 6ª | `documentation-writer` | Qualidade do TEST_PLAN.md corrigido |

---

## Observações

1. **Cobertura zero é bloqueante.** Se o mapa de cobertura (§4 do relatório) mostra features com 0%, essas têm prioridade máxima. Uma feature sem testes é um bug esperando para acontecer.

2. **Segurança não é opcional.** Mesmo que o PRD.md não explicite requisitos de segurança, cenários básicos de RBAC e Multi-Tenant são obrigatórios em toda feature com endpoints.

3. **Cenários de segurança seguem um padrão.** Para cada endpoint: 1 cenário de sucesso (autorizado), 1 cenário de negação (não autorizado), 1 cenário de isolamento (tenant isolation). Isso garante cobertura mínima de segurança.

4. **A pirâmide é um guia, não uma camisa de força.** A proporção 60/25/10/5 é uma referência. Features exclusivamente de backend podem ter 0% E2E. Features de frontend podem ter mais E2E. Use bom senso.

5. **Incrementar versão SEMPRE.** Após correção, reexecutar o gate (FIX → GATE → aprovação). Este é o último gate antes do desenvolvimento — garantir que ele passe é crítico.

---

## Registro de Alterações do Prompt

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 13/07/2026 | Criação inicial: correção de TEST_PLAN.md baseada em TECHNICAL_TEST_PLAN_FAIL_REPORT.md, foco em cobertura e segurança | Time de Arquitetura |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: writing-skills, agile-ba-practices, test-strategy-design, security-reviewer, qa-test-planner.*
