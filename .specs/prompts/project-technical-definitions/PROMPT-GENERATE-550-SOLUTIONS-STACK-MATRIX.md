# PROMPT-GENERATE-550-SOLUTIONS-STACK-MATRIX

## Contexto

Este prompt gera o artefato `550-SOLUTIONS-STACK-MATRIX.md` — a **matriz de stacks tecnológicas** do projeto. Para cada solução catalogada na Fase 2, define a stack precisa: linguagem, framework, banco de dados, mensageria, containerização, CI/CD, com versões específicas e justificativas.

**Inputs upstream (Bloco C — F15):** Este artefato é o terceiro elo da cadeia progressiva do Bloco C, recebendo insumos cumulativos:
- **Inputs Globais do Roadmap** (todas as 8 variáveis)
- **Bloco 0 (Product Def & Backlog & PRD):** `INTAKE-LOG.md`, `DOR-ASSESSMENT.md`, `PRODUCT-BACKLOG-LIST.md`, `PRD-DEFINITION.md`
- **Bloco A (People & Solutions):** `TEAM-SKILLS-MAP.md` + `TEAM-CAPACITY.md`
- **Bloco B (6 Disciplinas Técnicas):** `ARCHITECTURE-DEFINITION.md` + `SECURITY-DEFINITION.md` + `DATA-ARCHITECTURE-DEFINITION.md` + `DEVOPS-SRE-DEFINITION.md` + `TEST-STRATEGY-DEFINITION.md` + `INFRA-CLOUD-DEFINITION.md`
- **F13 (Bloco C):** `530-SOLUTIONS-CATALOG.md` — catálogo de soluções
- **F14 (Bloco C):** `540-SOLUTIONS-MATRIX.md` — matriz solução×stack×owner
- **Arquitetura Global:** ADRs, blueprints da pasta `architecture/`

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Opcional) Lista de caminhos para prompts auxiliares ou contextos adicionais |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros

### Passo 1 — Carregar Documentos Base
Ler `530-SOLUTIONS-CATALOG.md`, ADRs globais, blueprints (pom.xml templates, Dockerfiles, application.yml) e TECHNICAL-PLAN.md.

### Passo 2 — Invocar Skills Especializadas
Para cada solução do catálogo, invocar skills específicos da stack para definir versões, justificar escolhas e referenciar ADRs/blueprints.

### Passo 3 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/550-SOLUTIONS-STACK-MATRIX.md` com:
- Matriz: Solução × Linguagem × Framework × Banco × Mensageria × Containerização × CI/CD
- Versões específicas de cada tecnologia
- Justificativa para escolhas não-óbvias
- Referências aos ADRs e blueprints da pasta `architecture/`
- Compatibilidade cross-solution (ex: versão Java consistente entre serviços)

### Passo 4 — Validação Pós-Geração
Verificar: todas as soluções do catálogo cobertas, versões especificadas, justificativas documentadas, referências aos ADRs.

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados. O agente tem autonomia para selecionar outros mais aderentes à stack específica.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `tech-stack-evaluator` | Avaliar e recomendar stacks para cada solução | Avaliação |
| 2 | `technology-stack-blueprint-generator` | Gerar blueprint de stack | Avaliação |
| 3 | `java-architect` | Definir stack Java/Spring precisa | Java |
| 4 | `java-spring-boot` | Especificar versão Spring Boot e módulos | Java |
| 5 | `react-best-practices` | Definir stack React/Next.js | Frontend |
| 6 | `nextjs-best-practices` | Especificar versão Next.js e configurações | Frontend |
| 7 | `database-architect` | Definir bancos de dados e versões | Dados |
| 8 | `postgres-pro` | Especificar PostgreSQL e extensões | Dados |
| 9 | `kubernetes-specialist` | Definir stack K8s e Helm | Infra |
| 10 | `docker-expert` | Especificar Dockerfiles e imagens base | Infra |
| 11 | `devops-engineer` | Definir pipeline CI/CD e ferramentas | DevOps |
| 12 | `documentation-writer` | Redigir a matriz consolidada | Documentação |

> **🔄 Flexibilidade:** Skills de linguagens/frameworks específicos devem ser trocados conforme a stack real do projeto (Go, Python, Flutter, etc.). Justificar no changelog.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt gerador da matriz de stacks | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
