# PROMPT-FIX-SECURITY-FROM-GATE

## Contexto

Este prompt é acionado quando o `PROMPT-GATE-SECURITY-TECHNICAL_SOLUTION.md` reprova o artefato `SECURITY.md` com veredito `REPROVADO` (≥1 dimensão com menos de 75% de aprovação, ou violação de qualquer Regra de Ouro do GLOBAL-SECURITY.md).

O agente corretor atua como um **cirurgião de segurança** — aplica correções pontuais e cirúrgicas no SECURITY.md com base no relatório de falha (`SECURITY_SCOPE_FAIL_REPORT.md`), preservando todas as seções e controles que foram aprovados. **Nunca reescreve o documento do zero.**

**Princípio fundamental:** As Regras de Ouro do GLOBAL-SECURITY.md são inegociáveis. Se uma NC for de violação de Regra de Ouro, a correção é obrigatória e deve referenciar explicitamente o GLOBAL-SECURITY.md como fonte normativa.

---

## Parâmetros de Entrada

> **Instrução:** No momento de invocar este prompt, o agente deve solicitar ao humano os valores abaixo. Se algum não for informado, perguntar antes de prosseguir.

| Parâmetro | Descrição | Exemplo |
|---|---|---|
| `{SOLUTION_PATH}` | Caminho absoluto da pasta da solução técnica | `/home/user/work/backend/java/spring/microservices/ms-fbso-platform-admin` |
| `{PROJECT_PATH}` | Caminho absoluto da pasta do projeto de negócio | `/home/user/work/business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{PROJECT_NAME}` | Nome/código do projeto | `PRJ-FIN-2026-0003-SAAS-FBSO-ORG` |
| `{SOLUTION_NAME}` | Nome da solução/microsserviço | `ms-fbso-platform-admin` |
| `{STACK}` | Stack tecnológica principal | `Java 25 + Spring Boot + PostgreSQL` |
| `{SECURITY_GLOBAL}` | Caminho absoluto para o GLOBAL-SECURITY.md | `/home/user/work/.specs/security/GLOBAL-SECURITY.md` |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

Verificar se TODOS os 6 parâmetros foram informados. Se algum estiver ausente, perguntar antes de prosseguir.

### Passo 1 — Carregar Artefatos e Relatório de Falha

```
Ler obrigatoriamente:
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/SECURITY_SCOPE_FAIL_REPORT.md
    │   (relatório de falha do gate — entrada primária para as correções)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/SECURITY.md
    │   (artefato a ser corrigido — NÃO reescrever do zero)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/PRD.md
    │   (baseline de escopo — pode conter requisitos de segurança não capturados)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/ARCHITECTURE.md
    │   (pipeline de segurança, cross-cutting concerns, estrutura de pacotes)
    └── {SECURITY_GLOBAL} (GLOBAL-SECURITY.md)
        (regras de ouro e checklist SDD — referência normativa mestra)

Se SECURITY_SCOPE_FAIL_REPORT.md não existir → ERRO: "Relatório de falha não encontrado.
Execute primeiro o PROMPT-GATE-SECURITY-TECHNICAL_SOLUTION.md para gerar o diagnóstico."
Se GLOBAL-SECURITY.md não existir → ERRO CRÍTICO: "GLOBAL-SECURITY.md não encontrado.
As correções das Regras de Ouro exigem a referência normativa."
```

### Passo 2 — Processar Não-Conformidades por Prioridade

As NCs do relatório de falha são processadas em ordem de prioridade. NCs de mesma prioridade são processadas na ordem em que aparecem no relatório.

| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| **P0** (Crítica) | Violação de Regra de Ouro do GLOBAL-SECURITY.md | Corrigir obrigatoriamente. Adicionar controle específico na seção correspondente, com citação explícita da regra do GLOBAL-SECURITY.md. |
| **P0** (Crítica) | Ameaça STRIDE sem controle mitigador | Adicionar controle mitigador para a ameaça, com: o quê, quem implementa, como verificar. |
| **P0** (Crítica) | Cobertura zero em dimensão de validação | Criar/reescrever a seção correspondente do SECURITY.md do zero, usando PRD.md e ARCHITECTURE.md como guia. |
| **P1** (Alta) | Dimensão reprovada com < 75% | Corrigir cada verificação falha na dimensão, uma por uma, conforme critério do gate. |
| **P1** (Alta) | Pipeline de segurança inconsistente com ARCHITECTURE.md | Alinhar o SECURITY.md com o pipeline de segurança documentado no ARCHITECTURE.md. |
| **P2** (Média) | Tabela de compliance GLOBAL-SECURITY.md incompleta | Completar a tabela: Regra Global → Controle Implementado → Evidência. |
| **P2** (Média) | OWASP Top 10 com cobertura parcial | Para cada categoria com status ⚠️ Parcial, detalhar o controle ou justificar a não-aplicabilidade. |
| **P3** (Baixa) | Security ADRs insuficientes (< 3) | Adicionar ADRs faltantes com: ID, Decisão, Alternativas, Justificativa. |
| **P3** (Baixa) | Documentação incompleta (ex: faltam ferramentas nomeadas) | Complementar com nomes de ferramentas, frequências, responsáveis. |

### Passo 3 — Aplicar Correções Cirúrgicas

Para CADA NC listada no relatório de falha, seguindo a ordem de prioridade:

1. **Localizar** a seção do SECURITY.md afetada pela NC
2. **Aplicar** a ação corretiva específica (conforme tabela do Passo 2)
3. **Preservar** todo o conteúdo das seções não afetadas — correção é aditiva ou substitutiva pontual, nunca destrutiva
4. **Referenciar** o GLOBAL-SECURITY.md como fonte normativa sempre que a correção envolver Regras de Ouro ou Checklist SDD
5. **Marcar** a NC como resolvida

**Regra especial — Violações de Regra de Ouro (P0):**

Se qualquer NC for de violação de Regra de Ouro (NC-1.1, NC-1.2 ou NC-1.3), a correção DEVE:
- Citar textualmente a regra do GLOBAL-SECURITY.md
- Descrever o controle concreto que implementa a regra na stack `{STACK}`
- Incluir evidência de conformidade (ex: configuração do Spring Security, vault config, schema validation)

### Passo 4 — Atualizar Registro de Alterações

Adicionar entrada no changelog do SECURITY.md (Seção 12):

```
| v{X+1} | [AAAA-MM-DD] | Correção pós-gate: {N} NCs resolvidas ({P0}x P0, {P1}x P1, {P2}x P2, {P3}x P3). Relatório: SECURITY_SCOPE_FAIL_REPORT.md | Agente Corretor SEC/IA |
```

### Passo 5 — Validar Correções

Antes de reportar sucesso, executar as verificações abaixo:

| # | Verificação | Critério de Sucesso |
|---|---|---|
| 5.1 | NCs P0 resolvidas | 100% das NCs de prioridade P0 foram endereçadas |
| 5.2 | Regras de Ouro em conformidade | As 3 Regras de Ouro do GLOBAL-SECURITY.md estão explicitamente cobertas |
| 5.3 | Threat model completo | Toda ameaça de alta severidade tem controle mitigador documentado |
| 5.4 | Tabela de compliance completa | Tabela Regra Global → Controle → Evidência preenchida para todas as regras |
| 5.5 | Consistência cross-documento | As correções não criaram inconsistências com PRD.md e ARCHITECTURE.md |
| 5.6 | Changelog atualizado | Seção 12 registra a correção com referência ao relatório de falha |
| 5.7 | Nenhuma seção aprovada foi alterada desnecessariamente | Conteúdo das dimensões APROVADAS e RESSALVAS permanece intacto |

---

## Skills Utilizados

| Ordem | Skill | Propósito |
|---|---|---|
| 1 | `gap-analysis` | Analisar relatório de falha e priorizar correções |
| 2 | `security-auditor` | Corrigir não-conformidades de segurança conforme GLOBAL-SECURITY.md |
| 3 | `threat-modeling-expert` | Refinar threat model STRIDE e adicionar controles faltantes |
| 4 | `security-best-practices` | Garantir que correções seguem boas práticas para a stack `{STACK}` |
| 5 | `documentation-writer` | Atualizar SECURITY.md com correções mantendo consistência documental |

---

## Observações

- As 3 Regras de Ouro do GLOBAL-SECURITY.md são **inegociáveis**. Nenhuma correção pode enfraquecê-las.
- Correções são **cirúrgicas** — modifica-se apenas o necessário, preserva-se todo o resto.
- O changelog (Seção 12) deve ser **sempre** incrementado com referência ao `SECURITY_SCOPE_FAIL_REPORT.md`.
- Após correção, o fluxo natural é: **FIX → GATE → (se APROVADO ou RESSALVA) prosseguir para Fase 4 (SPECS.md).**
- Se PRD.md ou ARCHITECTURE.md forem alterados concomitantemente, reavaliar a necessidade de correções adicionais por efeito cascata.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 21/07/2026 | Criação inicial: prompt de correção cirúrgica do SECURITY.md com matriz de prioridades P0-P3 e validação pós-correção | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Foram utilizados os skills: gap-analysis, security-auditor, threat-modeling-expert, security-best-practices, documentation-writer.*
