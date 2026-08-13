# PROMPT: GERADOR DE SOFTWARE ARCHITECTURE DOCUMENT (SAD)
## Versão: 1.0 — WATERFALL Orchestrator

Atue como Arquiteto de Soluções Sênior, especializado em documentação de arquitetura de software com 6 visões.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo onde o arquivo será criado (formato: `030-SAD-{PROJECT_ID_NAME}.md`) |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `ARCHITECTURE_GLOBAL` | Caminho para a pasta de arquitetura global — ADRs, blueprints, padrões corporativos |
| `SECURITY_GLOBAL` | Caminho para o GLOBAL-SECURITY.md — regras de ouro, checklist SDD, threat model global |
| `TECHNICAL_SOLUTIONS` | Lista de soluções técnicas do projeto (`TECHNICAL_SOLUTION_NAMES`) — nomes dos microsserviços, frontends, batches |
| `PROJECT-STACK` | Stack tecnológica validada contra baseline corporativa em `.specs/standards/STACK-PADROES-CORPORATIVOS-FBSO-ORG.md` |
| `UPSTREAM_DOCS` | Lista: `[001-PROJECT-CHARTER, 010-FRD, 020-SRS, 025-RTM-FASE-2]` |
| `EXTRA_INPUTS` | Documentos brutos de entrada adicionais fornecidos pelo humano (`PROJECT_DOCUMENTS_INPUTS`) |
| `SKILLS` | Lista de skills: ["software-architecture", "architecture-designer", "architecture-patterns"] |

## Regras

1. **NUNCA** procure por inputs em diretórios — use apenas o que foi passado nos parâmetros acima. Os parâmetros listados na tabela de Inputs são a única fonte de dados — não leia outros arquivos além dos explicitamente fornecidos
2. **LEIA** os documentos em `UPSTREAM_DOCS` — todos os artefatos devem rastrear de volta a eles
3. Skills: tente usar as skills listadas em `SKILLS` via `Skill` tool. Se falharem, use o template de fallback abaixo
4. Crie o arquivo em `DOC_PATH` com o status inicial `[STATUS: Em análise]`
5. Ao final, retorne `{DOC_PATH}` confirmando a criação

## Template de Fallback

```
# SOFTWARE ARCHITECTURE DOCUMENT (SAD): {NOME DO PROJETO}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documentos Base** | 001-PROJECT-CHARTER, 010-FRD, 020-SRS, 025-RTM-FASE-2 |
| **Data de Elaboração** | {DATA ATUAL} |
| **Versão** | 1.0 |
| **Metodologia** | WATERFALL |

---
## SAD — Software Architecture Document (Documento de Arquitetura de Software)
O **SAD** é o blueprint arquitetural da plataforma: consolida os requisitos funcionais do **010-FRD**, as especificações técnicas do **020-SRS** e a rastreabilidade técnica do **025-RTM-FASE-2** em decisões de arquitetura (estilos, padrões, ADRs) e nas seis visões — lógica, dados, segurança, DevOps/SRE, infraestrutura e testes. Serve como ponte entre a especificação de requisitos e os desenhos de alto nível (**035-HLD**) e baixo nível (**040-LLD**).

### 1. Architectural Overview
[Estilo arquitetural, ADRs chave, diagrama de contexto C4 Level 1]

### 2. Solution Architecture
[Componentes, camadas, fluxos principais, padrões de design]

### 3. Data Architecture
[Modelo conceitual de dados, storage strategy, data governance, data flow diagrams]

### 4. Security Architecture
[Threat model (STRIDE), authN/authZ (OAuth2, RBAC), data protection (encryption at rest/transit), secrets management, OWASP coverage]

### 5. DevOps/SRE Architecture
[CI/CD pipeline, observability (logs, metrics, traces, alerts), SLOs/SLIs, incident response, IaC strategy]

### 6. Infrastructure/Cloud Architecture
[Deployment topology, scaling strategy, disaster recovery (RPO/RTO), network architecture, environment inventory]

### 7. Testing Architecture
[Test pyramid, strategy per layer (unit/integration/e2e), quality gates, performance/security testing integration]

### 8. Cross-cutting Concerns
[Logging, error handling, i18n/l10n, feature flags, caching strategy]

### 9. ADR Registry
| ADR ID | Título | Decisão | Status | NFR Vinculado (SRS) |
|--------|--------|---------|--------|---------------------|
| ADR-001 | ... | ... | Accepted/Proposed/Superseded | NFR-XX |

### 10. Traceability SAD → SRS → BRD → Charter
| Decisão Arquitetural (SAD) | Requisito Funcional (SRS) | Requisito Negócio (BRD) | Objetivo (Charter) | Status |
|---------------------------|--------------------------|------------------------|-------------------|--------|
| ADR-001 | FR-01 | REQ-01 | OBJ-01 | ✅ |
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` se o documento estiver completo.
