---
title: "Índice Central — web_app-tax-nexus-portal"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["index", "documentation"]
---

# 📑 Índice da Documentação — TaxNexus Portal (TaaS)

**Solução:** `web_app-tax-nexus-portal` — Portal frontend React para simulação da Reforma Tributária 2026 (Tax as a Service).  
**Stack:** React 19 + TypeScript 5.9 + Vite 8 + Recharts 3.8  
**Última atualização:** 2026-07-08  
**Score de confiança:** 82% (ver `governance/confidence-report.md`)  
**Pipeline:** FRONTEND-SPEC-MINING (passos 5→11 executados)

---

## 📁 Domain
| Arquivo | Descrição |
|---|---|
| [`domain/domain.md`](domain/domain.md) | Glossário de domínio tributário + ubiquitous language |

## 📁 API
| Arquivo | Descrição |
|---|---|
| [`api/tax-nexus-api.yaml`](api/tax-nexus-api.yaml) | Contrato OpenAPI 3.1 da API de cálculo tributário consumida pelo portal |

## 📁 Architecture
| Arquivo | Descrição |
|---|---|
| [`architecture/architecture.md`](architecture/architecture.md) | Visão geral da arquitetura do portal |
| [`architecture/c4-context.md`](architecture/c4-context.md) | C4 — Nível 1 (Contexto) |
| [`architecture/c4-containers.md`](architecture/c4-containers.md) | C4 — Nível 2 (Containers) |
| [`architecture/integrations.md`](architecture/integrations.md) | Integrações externas e dependências |

## 📁 Engineering
| Arquivo | Descrição |
|---|---|
| [`engineering/code-analysis.md`](engineering/code-analysis.md) | Análise técnica de fluxo de código |
| [`engineering/api-guidelines.md`](engineering/api-guidelines.md) | Padrões de consumo de API, erros, formatos |

## 📁 Product
| Arquivo | Descrição |
|---|---|
| [`product/product.md`](product/product.md) | Visão do produto e proposta de valor |
| [`product/requirements.md`](product/requirements.md) | Requisitos funcionais e não-funcionais (EARS + MoSCoW) |
| [`product/feature-roadmap.md`](product/feature-roadmap.md) | Roadmap de features e dívidas técnicas |

## 📁 Governance
| Arquivo | Descrição |
|---|---|
| [`governance/inventory.md`](governance/inventory.md) | Inventário do projeto e cobertura |
| [`governance/confidence-report.md`](governance/confidence-report.md) | Relatório de confiança da documentação |

## 📁 Security
| Arquivo | Descrição |
|---|---|
| [`security/SECURITY.md`](security/SECURITY.md) | Definições de segurança do frontend |

## 📁 Design (Design System)
| Arquivo | Descrição |
|---|---|
| [`design/DESIGN.md`](design/DESIGN.md) | Design system document (cores, tipografia, componentes, regras) |
| [`design/design-tokens.md`](design/design-tokens.md) | Tokens de design extraídos (cores, spacing, radius, shadows, breakpoints) |
| [`design/components.md`](design/components.md) | Catálogo de componentes React com variantes, props e estados |

## 📁 Frontend (Análises Específicas)
| Arquivo | Descrição |
|---|---|
| [`frontend/audit-report.md`](frontend/audit-report.md) | Auditoria técnica: a11y, performance, responsividade, anti-padrões (13/20) |
| [`frontend/ux-critique.md`](frontend/ux-critique.md) | Avaliação heurística de UX (Nielsen, 25/40) |
| [`frontend/visual-inspection-report.md`](frontend/visual-inspection-report.md) | Inspeção visual: layout, contraste, responsividade |
| [`frontend/accessibility-audit.md`](frontend/accessibility-audit.md) | Auditoria WCAG 2.1/2.2 (13 findings, ~40% conformidade AA) |
| [`frontend/functional-verification.md`](frontend/functional-verification.md) | Verificação funcional: 5 fluxos de usuário mapeados |

## 📁 Business Projects
| Arquivo | Descrição |
|---|---|
| [`business-projects/README.md`](business-projects/README.md) | Índice de projetos corporativos vinculados |

## 📁 Skill Output
| Arquivo | Descrição |
|---|---|
| [`skill-output/2026-07-08-183455-spec-mining.md`](skill-output/2026-07-08-183455-spec-mining.md) | Relatório da mineração inicial (spec-miner, 17 artefatos) |
| [`skill-output/2026-07-08-195500-frontend-spec-mining.md`](skill-output/2026-07-08-195500-frontend-spec-mining.md) | Relatório da mineração frontend (passos 5→11, 8 artefatos) |

---

## 📊 Cobertura por Pasta

| Pasta | Arquivos Esperados | Arquivos Existentes | Cobertura |
|---|---|---|---|
| domain/ | 1 | 1 | 100% |
| api/ | 1 | 1 | 100% |
| architecture/ | 6+ | 4 | 67% (sem c4-components, adrs/) |
| engineering/ | 2 | 2 | 100% |
| product/ | 3 | 3 | 100% |
| design/ | 3 | 3 | 100% |
| frontend/ | 6 | 5 | 83% (sem performance-report.md) |
| governance/ | 2 | 2 | 100% |
| security/ | 1 | 1 | 100% |
| business-projects/ | 1+N | 1 | 50% (sem subpasta PRJ-FIN-2026-0001) |
| features/ | 0+N | 0 | N/A |
| skill-output/ | 0+N | 2 | N/A |

🤖 *Documentação gerada por mineração reversa de especificações: spec-miner (artefatos base) + FRONTEND-SPEC-MINING (passos 5→11, artefatos de design e frontend).*
