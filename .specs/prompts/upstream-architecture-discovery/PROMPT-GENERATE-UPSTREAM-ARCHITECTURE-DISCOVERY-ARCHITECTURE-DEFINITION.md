# PROMPT-GENERATE-UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION

## Contexto

Este prompt gera o artefato `UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION.md` — a **definição de arquitetura do projeto** que especifica como todas as soluções técnicas se integram.

**Este documento é independente de tecnologias específicas.** Durante a análise da stack tecnológica do projeto, identifique as tecnologias a serem utilizadas através do contexto do projeto e aplicando questionário para o usuário, e busque skills relacionados a essas tecnologias para aprimorar as estimativas. Caso não encontre skills específicos, utilize skills generalistas de arquitetura e engenharia de sistemas, e tambem utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills.

**Inputs upstream:**
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD Discovery-Level (F1)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` — Catálogo de Soluções (F8)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md` — Matriz Solução×Disciplina (F9)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SPECS.md` — Consolidação Técnica (F10)
5. `{ARCHITECTURE_GLOBAL}/` — ADRs, blueprints, data standards globais

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Diretiva) Checkpoint HITL: sempre solicitar ao usuário se deseja fornecer informações adicionais ou novos direcionamentos via prompt |
| `{PROJECT-TEAM-SKILLS-MAP}` | Skills necessários para o time de implementação (obter e validar com usuário) |
| `{PROJECT-TEAM-CAPACITY}` | Capacidade esperada do time — seniores, plenos, juniores, duração (obter e validar com usuário) |
| `{PROJECT-STACK}` | Stack tecnológica da solução. Baseline corporativa: `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`. Tecnologias fora do padrão exigem justificativa |

### Variáveis Derivadas (calculadas automaticamente)

```
PROJECT_COMPLETE_PATH_NAME    = PROJECT_PATH + "/" + PROJECT_ID_NAME
UPSTREAM_DISCOVERY_PATH       = PROJECT_COMPLETE_PATH_NAME + "/upstream-architecture-discovery"
```

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Confirmar os parâmetros de entrada recebidos e seu foco:
- `PROJECT_PATH={PROJECT_PATH}` — base dos projetos de negócio
- `PROJECT_ID_NAME={PROJECT_ID_NAME}` — identificador do projeto
- `TECHNICAL_SOLUTION_PATH={TECHNICAL_SOLUTION_PATH}` — base das soluções técnicas
- `TECHNICAL_SOLUTION_NAMES={TECHNICAL_SOLUTION_NAMES}` — soluções do projeto
- `ARCHITECTURE_GLOBAL={ARCHITECTURE_GLOBAL}` — ADRs e blueprints globais
- `SECURITY_GLOBAL={SECURITY_GLOBAL}` — documento de segurança global
- `PROJECT_DOCUMENTS_INPUTS` — documentos adicionais (se fornecidos)
- `PROJECT_PROMPT_INPUTS` — solicitar input adicional do usuário (checkpoint HITL)
- `PROJECT-TEAM-SKILLS-MAP` — skills do time (se fornecidos)
- `PROJECT-TEAM-CAPACITY` — capacidade do time (se fornecida)
- `PROJECT-STACK` — stack tecnológica; validar contra STACK-PADROES-CORPORATIVOS-FBSO-ORG.md
Validar que `{UPSTREAM_DISCOVERY_PATH}` existe e contém os artefatos upstream.

### Passo 1 — Carregar Documentos Base
Confirmar leitura dos seguintes artefatos upstream:
1. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` — PRD Discovery-Level (F1)
2. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` — Catálogo de Soluções (F8)
3. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md` — Matriz Solução×Disciplina (F9)
4. `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SPECS.md` — Consolidação Técnica (F10)
5. `{ARCHITECTURE_GLOBAL}/` — ADRs globais, blueprints (Java, Go, React), data standards (Protobuf, Avro)

### Passo 2 — Invocar Skills Especializadas
Invocar skills de C4, arquitetura, integração, mensageria e infra para projetar a arquitetura do sistema completo.

### Passo 2.5 — Apresentar Skills, Capacidade e Stack para Validação Humana

Avaliar e apresentar ao usuário para validação:

1. **PROJECT-TEAM-SKILLS-MAP:** Skills identificados como necessários para implementar a solução nesta disciplina.
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

2. **PROJECT-TEAM-CAPACITY:** Capacidade estimada do time nesta disciplina (ex: 2 seniores, 3 plenos).
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

3. **PROJECT-STACK:** Tecnologias identificadas para esta disciplina. Verificar conformidade com `STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`. Tecnologias fora do padrão corporativo DEVEM ser listadas com justificativa técnica e requerem aprovação explícita do usuário.
   ⏸️ **Solicitar validação do usuário e aguardar confirmação.**

### Passo 3 — Gerar o Artefato
Gerar `{UPSTREAM_DISCOVERY_PATH}/UPSTREAM-ARCHITECTURE-DISCOVERY-ARCHITECTURE-DEFINITION.md` com:
- Diagrama C4 Level 1 (System Context): todas as soluções + sistemas externos
- Diagrama C4 Level 2 (Container): comunicação entre serviços
- Matriz de integração: origem → destino → protocolo → autenticação
- Topologia de deploy (containers, rede, ambientes Dev/Staging/Prod)
- Estratégia de comunicação: síncrono (REST/gRPC) vs. assíncrono (mensageria/eventos)
- Diagramas de sequência para fluxos cross-solution críticos
- ADRs de integração (decisões que afetam múltiplas soluções)

---
## Layout do Documento (Modelo Estrutural)

> 📐 **Modelo de referência:** O documento `business-inputs/business-projects/PRJ-FIN-2026-0003-SAAS-FBSO-ORG/upstream-architecture-discovery/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md` ilustra a estrutura esperada. Use-o como referência de **formato e organização**, NÃO como fonte de dados — todo conteúdo deve ser gerado a partir dos artefatos do projeto corrente (`{PROJECT_ID_NAME}`).

### Estrutura Esperada do `DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md`

```markdown
# DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md
## Fase 2 — Bloco B: Architecture & Security & Specialists (Discovery-Level)

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documento** | DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION-v1.0 |
| **Versão** | 1.0 — Discovery-Level (Análise de Viabilidade) |
| **Data** | {DATA_ATUAL} |
| **Autor** | Solution Architect / Tech Lead |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em {DATA} |

**Documentos Vinculados:**
- [`DISCOVERY-LEVEL-PRD.md`](DISCOVERY-LEVEL-PRD.md) — PRD Discovery-Level (F1)
- [`STACK-PADROES-CORPORATIVOS-FBSO-ORG.md`](../../../.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md) — Padrões Corporativos
- [`GLOBAL-SECURITY.md`](../../../.specs/security/GLOBAL-SECURITY.md) — Política de Segurança Global
- [`architecture/adr/`](../../../architecture/adr/) — ADRs globais

---

## 1. Visão Arquitetural Macro
- **1.1 Abordagem Arquitetural:** Princípios arquiteturais (Discovery-Level) e justificativa
- **1.2 Soluções Técnicas do Projeto:** Tabela: Solução | Tipo | Propósito | Repositório | Status

## 2. Diagrama C4 Level 1 — System Context
- Diagrama Mermaid C4Context com todos os atores, sistemas externos, e boundary principal
- Pontos de controle de segurança numerados no contexto

## 3. Diagrama C4 Level 2 — Containers
- Diagrama Mermaid C4Container detalhando containers dentro de cada solução
- Container boundaries: frontend, gateway, backend, observabilidade, infraestrutura

## 4. Matriz de Integração
- Tabela: # | Origem | Destino | Protocolo | Autenticação | Propósito

## 5. Arquitetura de Módulos — Monólito Modular (ou arquitetura escolhida)
- **5.1 Módulos Internos:** Árvore de diretórios com módulos e shared
- **5.2 Regras de Modularização:** Tabela: Regra | Descrição

## 6. ADRs — Decisões Arquiteturais (Discovery-Level)
- Cada ADR com: Contexto, Decisão, Rationale, Consequências
- Ex: ADR-0001: Gateway como Trust Boundary, ADR-0002: Multi-Tenant via RLS, ADR-0003: Soft Delete, ADR-0004: Comunicação REST

## 7. Estratégia de Comunicação
- **7.1 Síncrono (REST/HTTPS):** Tabela: Caso de Uso | Endpoint Pattern | Exemplo
- **7.2 Assíncrono (Mensageria — Fase Futura):** Tabela: Cenário Futuro | Gatilho | Consumidor

## 8. Topologia de Deploy (Discovery-Level)
- Diagrama ASCII da topologia de rede (Cloudflare → Load Balancer → K8s Cluster → DBs)
- **8.1 Recursos Estimados:** Tabela por ambiente (Dev/Staging/Prod)

## 9. Stack Tecnológica — Validação contra Padrões Corporativos
- Tabela: Camada | Padrão Corporativo | Stack do Projeto | Conformidade
- Subseção: 🆕 Tecnologias Adicionais (Fora do Padrão — Justificativas)

## 10. Validação de Capacidade do Time (Step 2.5)
- **PROJECT-TEAM-SKILLS-MAP:** Tabela: Papel | Profissional | Nível | Skills-Chave
- **PROJECT-TEAM-CAPACITY:** Tabela: Indicador | Valor
- **PROJECT-STACK (Validada):** Resumo de conformidade com padrões corporativos

## 11. Riscos Arquiteturais e Estimativa de Esforço
- **11.1 Riscos Técnicos:** Tabela: ID | Risco | Prob. | Impacto | Mitigação
- **11.2 Estimativa de Esforço (Discovery-Level):** Tabela: Área | Complexidade | Esforço Estimado (dias) | Responsável

---

## Registro de Alterações
| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | {DATA} | Criação inicial: Architecture Definition Discovery-Level | Solution Architect |
```

### Passo 4 — Validação Pós-Geração

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `c4-architecture-c4-architecture` | Diagramas C4 Level 1-2 | C4 |
| 2 | `c4-context` | System Context diagram | C4 |
| 3 | `c4-container` | Container diagram | C4 |
| 4 | `architecture-patterns` | Padrões arquiteturais cross-solution | Arquitetura |
| 5 | `architecture-decision-records` | ADRs de integração | ADR |
| 6 | `create-architectural-decision-record` | Criar ADRs formais | ADR |
| 7 | `api-design-principles` | Design de APIs entre soluções | Integração |
| 8 | `api-patterns` | Padrões de API REST/gRPC | Integração |
| 9 | `openapi-spec-generation` | Contratos de API cross-solution | Integração |
| 10 | `event-sourcing-architect` | Estratégia de eventos entre serviços | Mensageria |
| 11 | `saga-orchestration` | Orquestração de sagas cross-solution | Mensageria |
| 12 | `kubernetes-architect` | Topologia de deploy K8s | Infra |
| 13 | `docker-expert` | Containerização das soluções | Infra |
| 14 | `cloud-design-patterns` | Padrões de cloud | Cloud |
| 15 | `domain-driven-design` | Bounded contexts cross-solution | DDD |
| 16 | `deployment-pipeline-design` | Pipeline de deploy cross-solution | DevOps |
| 17 | `mermaid-expert` | Diagramas Mermaid para o documento | Diagramas |
| 18 | `documentation-writer` | Redigir o Architecture Definition | Documentação |

**Skills generalistas de arquitetura (sempre aplicáveis):**
- `senior-architect`, `engineering-skills`, `engineering-advanced-skills`
- `architecture`, `software-architecture`, `architecture-patterns`
- `c4-container`, `c4-component`, `architecture-decision-records`
- `system-design`, `microservices-patterns`

**Skills tecnológicos (condicionais — buscar ao identificar a stack):**
- Ao identificar uma tecnologia específica durante a análise da stack, procure skills relacionados a essa tecnologia para aprimorar as estimativas
- Caso não encontre skills específicos para a tecnologia identificada, utilize os skills generalistas listados acima como referência, e tambem utilize as skills `find-skills`, `skill-router`, `antigravity-skill-orchestrator` passando informações das necessidades para tambem ajudar na busca de skills.

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt gerador da definição de arquitetura | Time de Arquitetura |
| 2.0 | 30/07/2026 | Atualização F3→F7: adicionada referência cruzada aos artefatos F9-F12 (Bloco B) | Time de Arquitetura |

---

## Arquivos Utilizados na Tarefa

| # | Arquivo | Propósito |
|---|---|---|
| 1 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-PRD.md` | PRD Discovery-Level (F1) — visão do produto |
| 2 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md` | Catálogo de Soluções (F8) — inventário de soluções |
| 3 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md` | Matriz Solução×Disciplina (F9) |
| 4 | `{UPSTREAM_DISCOVERY_PATH}/DISCOVERY-LEVEL-SPECS.md` | Consolidação Técnica (F10) |
| 5 | `{ARCHITECTURE_GLOBAL}/` | ADRs, blueprints, data standards globais |
| 6 | `{PROJECT_DOCUMENTS_INPUTS}` | Documentos adicionais (se fornecidos) |
| 7 | `{PROJECT-TEAM-SKILLS-MAP}` | Skills do time (se fornecidos) |
| 8 | `{PROJECT-TEAM-CAPACITY}` | Capacidade do time (se fornecida) |
| 9 | `{PROJECT-STACK}` | Stack tecnológica (validar contra padrões corporativos) |
| 10 | `{PROJECT_PROMPT_INPUTS}` | Checkpoint HITL — input adicional do usuário |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
