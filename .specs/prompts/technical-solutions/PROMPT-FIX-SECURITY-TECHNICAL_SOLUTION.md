# PROMPT-FIX-SECURITY-FROM-GATE

## Contexto

Este prompt é acionado quando o `PROMPT-GATE-SECURITY-TECHNICAL_SOLUTION.md` reprova o artefato `TECHNICAL-SOLUTION-SECURITY.md` com veredito `REPROVADO` (≥1 dimensão com menos de 75% de aprovação, ou violação de qualquer Regra de Ouro do GLOBAL-SECURITY.md).

O agente corretor atua como um **cirurgião de segurança** — aplica correções pontuais e cirúrgicas no TECHNICAL-SOLUTION-SECURITY.md com base no relatório de falha (`SECURITY_SCOPE_FAIL_REPORT.md`), preservando todas as seções e controles que foram aprovados. **Nunca reescreve o documento do zero.**

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
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-SECURITY.md
    │   (artefato a ser corrigido — NÃO reescrever do zero)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-PRD.md
    │   (baseline de escopo — pode conter requisitos de segurança não capturados)
    ├── {SOLUTION_PATH}/.specs/business-projects/{PROJECT_NAME}/TECHNICAL-SOLUTION-ARCHITECTURE.md
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
| **P0** (Crítica) | Cobertura zero em dimensão de validação | Criar/reescrever a seção correspondente do TECHNICAL-SOLUTION-SECURITY.md do zero, usando TECHNICAL-SOLUTION-PRD.md e TECHNICAL-SOLUTION-ARCHITECTURE.md como guia. |
| **P1** (Alta) | Dimensão reprovada com < 75% | Corrigir cada verificação falha na dimensão, uma por uma, conforme critério do gate. |
| **P1** (Alta) | Pipeline de segurança inconsistente com TECHNICAL-SOLUTION-ARCHITECTURE.md | Alinhar o TECHNICAL-SOLUTION-SECURITY.md com o pipeline de segurança documentado no TECHNICAL-SOLUTION-ARCHITECTURE.md. |
| **P2** (Média) | Tabela de compliance GLOBAL-SECURITY.md incompleta | Completar a tabela: Regra Global → Controle Implementado → Evidência. |
| **P2** (Média) | OWASP Top 10 com cobertura parcial | Para cada categoria com status ⚠️ Parcial, detalhar o controle ou justificar a não-aplicabilidade. |
| **P3** (Baixa) | Security ADRs insuficientes (< 3) | Adicionar ADRs faltantes com: ID, Decisão, Alternativas, Justificativa. |
| **P3** (Baixa) | Documentação incompleta (ex: faltam ferramentas nomeadas) | Complementar com nomes de ferramentas, frequências, responsáveis. |

### Passo 3 — Aplicar Correções Cirúrgicas

Para CADA NC listada no relatório de falha, seguindo a ordem de prioridade:

1. **Localizar** a seção do TECHNICAL-SOLUTION-SECURITY.md afetada pela NC
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

Adicionar entrada no changelog do TECHNICAL-SOLUTION-SECURITY.md (Seção 12):

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
| 5.5 | Consistência cross-documento | As correções não criaram inconsistências com TECHNICAL-SOLUTION-PRD.md e TECHNICAL-SOLUTION-ARCHITECTURE.md |
| 5.6 | Changelog atualizado | Seção 12 registra a correção com referência ao relatório de falha |
| 5.7 | Nenhuma seção aprovada foi alterada desnecessariamente | Conteúdo das dimensões APROVADAS e RESSALVAS permanece intacto |

---

## Skills Utilizados

> **📌 Nota sobre Skills:** A tabela abaixo lista os skills **recomendados** para a correção cirúrgica do TECHNICAL-SOLUTION-SECURITY.md. O agente corretor deve usá-los como ponto de partida, mas tem autonomia para selecionar outros skills identificados como mais aderentes às necessidades específicas das não-conformidades, stack tecnológica ou domínio de negócio. A ordem sugerida reflete o fluxo natural de correção: analisar falhas → corrigir por categoria → validar.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório de falha (`SECURITY_SCOPE_FAIL_REPORT.md`) e priorizar correções por gravidade (P0→P3) | Análise |
| 2 | `security-auditor` | Corrigir não-conformidades de conformidade com GLOBAL-SECURITY.md, especialmente Regras de Ouro (P0) | Conformidade |
| 3 | `security-audit` | Reavaliar gaps de segurança após cada correção; verificar se novas NCs não foram introduzidas | Conformidade |
| 4 | `skill-security-auditor` | Validação cruzada pós-correção: verificar se as correções mantêm consistência com TECHNICAL-SOLUTION-PRD.md e TECHNICAL-SOLUTION-ARCHITECTURE.md | Conformidade |
| 5 | `threat-modeling-expert` | Refinar threat model STRIDE: adicionar controles mitigadores faltantes, recalibrar matriz de risco | Threat Model |
| 6 | `senior-security` | Supervisão sênior das correções: validar que controles adicionados são Proporcionais, eficazes e não introduzem novos riscos | Estratégia |
| 7 | `security-best-practices` | Garantir que correções seguem boas práticas de segurança para a stack `{STACK}` | Boas Práticas |
| 8 | `api-security-best-practices` | Corrigir controles de API: rate limiting, CORS, input validation, autenticação/autorização | API Security |
| 9 | `api-security-testing` | Adicionar cenários de teste de segurança de API faltantes identificados pelo gate | API Security |
| 10 | `security-scanning-security-sast` | Corrigir seção de SAST no pipeline: adicionar ferramenta, regras, frequência, gates no CI/CD | Pipeline |
| 11 | `security-scanning-security-dependencies` | Corrigir seção de SCA/gestão de dependências: ferramenta, política de atualização, SLSA | Pipeline |
| 12 | `security-scanning-security-hardening` | Corrigir seção de hardening de infraestrutura: secrets management, containers, WAF | Infra |
| 13 | `security-reviewer` | Revisão final pós-correção: verificar cobertura das regras de ouro, checklist SDD e OWASP Top 10 | Revisão |
| 14 | `security-review` | Revisão complementar: verificar consistência cross-documento e aderência ao TECHNICAL-SOLUTION-PRD.md e TECHNICAL-SOLUTION-ARCHITECTURE.md | Revisão |
| 15 | `engineering-skills` | Garantir que todas as correções são acionáveis e verificáveis pela equipe de engenharia | Qualidade |
| 16 | `documentation-writer` | Atualizar TECHNICAL-SOLUTION-SECURITY.md com correções mantendo consistência documental e atualizando changelog (Seção 12) | Documentação |

> **🔄 Flexibilidade:** Se durante a correção o agente identificar que um skill diferente dos listados acima é mais adequado para resolver uma NC específica (ex: um skill de conformidade regulatória LGPD/PCI para correções na Dimensão 4, um skill de segurança em nuvem para hardening de infraestrutura, ou um skill especializado na stack `{STACK}` para controles técnicos), ele deve substituí-lo e justificar a escolha no changelog do TECHNICAL-SOLUTION-SECURITY.md.

---

## Observações

- As 3 Regras de Ouro do GLOBAL-SECURITY.md são **inegociáveis**. Nenhuma correção pode enfraquecê-las.
- Correções são **cirúrgicas** — modifica-se apenas o necessário, preserva-se todo o resto.
- O changelog (Seção 12) deve ser **sempre** incrementado com referência ao `SECURITY_SCOPE_FAIL_REPORT.md`.
- Após correção, o fluxo natural é: **FIX → GATE → (se APROVADO ou RESSALVA) prosseguir para Fase 4 (TECHNICAL-SOLUTION-SPECS.md).**
- Se TECHNICAL-SOLUTION-PRD.md ou TECHNICAL-SOLUTION-ARCHITECTURE.md forem alterados concomitantemente, reavaliar a necessidade de correções adicionais por efeito cascata.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 21/07/2026 | Criação inicial: prompt de correção cirúrgica do TECHNICAL-SOLUTION-SECURITY.md com matriz de prioridades P0-P3 e validação pós-correção | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Skills de referência: gap-analysis, security-auditor, security-audit, skill-security-auditor, threat-modeling-expert, senior-security, security-best-practices, api-security-best-practices, api-security-testing, security-scanning-security-sast, security-scanning-security-dependencies, security-scanning-security-hardening, security-reviewer, security-review, engineering-skills, documentation-writer. Outros skills podem ser utilizados conforme aderência à necessidade específica.*
