# PROMPT-FIX-PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION

## Contexto

Este prompt é acionado quando o gate reprova `PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md`. O agente corretor aplica correções cirúrgicas com base no relatório de falha, com atenção especial às Regras de Ouro (P0 — inegociáveis).

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{SECURITY_GLOBAL}` | Caminho do GLOBAL-SECURITY.md |

---

## Fluxo de Execução

### Passo 1 — Carregar Relatório de Falha e Artefatos
Ler `SECURITY_DEFINITION_SCOPE_FAIL_REPORT.md`, o documento atual, GLOBAL-SECURITY.md, Architecture Definition.

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Violação de Regra de Ouro | Corrigir obrigatoriamente com citação explícita da regra do GLOBAL-SECURITY.md |
| P0 | Threat model sem controles | Adicionar controles mitigadores para cada ameaça |
| P1 | Pipeline DevSecOps incompleto | Completar SAST, SCA, Secret Scanning |
| P2 | Matriz de compliance incompleta | Mapear LGPD/PCI/SOC2 por solução |
| P3 | Controle não acionável | Detalhar: o quê, quem, como verificar |

### Passo 3 — Aplicar Correções Cirúrgicas
Para violações de Regra de Ouro (P0): citar textualmente a regra, descrever controle concreto, incluir evidência de conformidade.

### Passo 4 — Validar Correções
100% P0 resolvidas, 3 Regras de Ouro em conformidade, threat model completo.

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar | Análise |
| 2 | `security-auditor` | Corrigir NCs de conformidade | Auditoria |
| 3 | `threat-modeling-expert` | Refinar threat model | Threat Model |
| 4 | `senior-security` | Validar correções | Estratégia |
| 5 | `security-best-practices` | Garantir boas práticas | Boas Práticas |
| 6 | `documentation-writer` | Atualizar documento | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt de correção da definição de segurança | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
