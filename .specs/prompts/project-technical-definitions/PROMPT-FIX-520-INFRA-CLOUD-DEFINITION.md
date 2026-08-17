# PROMPT-FIX-520-INFRA-CLOUD-DEFINITION

## Contexto

Este prompt é acionado quando o gate reprova `520-INFRA-CLOUD-DEFINITION.md`. O agente corretor aplica correções cirúrgicas com base no relatório inline do gate. **Nunca reescreve o documento do zero. Modifique estritamente as seções, tabelas ou linhas apontadas como Não Compliance.**

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{ARCHITECTURE_GLOBAL}` | Caminho da pasta de arquitetura global |

---

## Fluxo de Execução

### Passo 1 — Carregar Relatório do Gate e Artefatos
Ler o **Relatório de Auditoria** emitido pelo gate (relatório inline com os IDs de conflito e respostas do humano), o documento atual, Architecture Definition (F7), Security Definition (F8), Data Architecture (F9), DevOps SRE (F10), ADRs globais.

### Passo 2 — Processar NCs por Prioridade
| Prioridade | Tipo de NC | Ação Corretiva |
|---|---|---|
| P0 | Ambiente sem topologia definida | Adicionar diagrama de topologia para o ambiente |
| P0 | Solução sem compute/sizing | Dimensionar compute com justificativa |
| P0 | DR sem RPO/RTO | Definir RPO e RTO por solução |
| P1 | Networking incompleto | Adicionar VPC, subnets, DNS, CDN faltantes |
| P2 | Estimativa de custos ausente | Calcular custo mensal por ambiente e serviço |
| P3 | Auto-scaling sem política | Definir políticas de escalabilidade |

### Passo 3 — Aplicar Correções Cirúrgicas
### Passo 4 — Validar Correções

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `gap-analysis` | Analisar relatório e priorizar | Análise |
| 2 | `cloud-architect` | Corrigir arquitetura cloud | Cloud |
| 3 | `aws-solution-architect` | Corrigir topologia AWS | AWS |
| 4 | `network-engineer` | Corrigir redes | Rede |
| 5 | `disaster-recovery` | Corrigir estratégia de DR | DR |
| 6 | `cost-optimization` | Corrigir estimativa de custos | Custos |
| 7 | `documentation-writer` | Atualizar documento | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 30/07/2026 | Criação inicial: prompt de correção da definição de infraestrutura cloud | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
