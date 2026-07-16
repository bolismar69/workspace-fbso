# Mapa de Especificações — ms-billing-engine-tax-rates

Diretório centralizado de arquitetura de contexto e especificações do microsserviço de cálculo de tributos sobre faturamento (IPI, ICMS, PIS, COFINS, ISS, CBS, IBS, IS, FUST, FUNTTEL).

**Confiança da documentação:** 99% (VERDE) — veja [governance/confidence-report.md](governance/confidence-report.md)
**Última atualização:** 2026-06-30 (PR #6 merge — Fases 0-1-2 Reforma Tributária: Admin Fiscal, Créditos, TaxToken, Simulação, Fornecedores; 211+ testes)

## Engenharia e Arquitetura
- **Visão Arquitetural e Regras de Solução:** Consulte [architecture/architecture.md](architecture/architecture.md)
- **Registro de Decisões Arquiteturais (ADRs):** Consulte [architecture/adrs/INDEX.md](architecture/adrs/INDEX.md)
- **Modelo de Dados (ERD Completo):** Consulte [architecture/erd.md](architecture/erd.md)
- **Dicionário de Dados (Função e Uso das Tabelas):** Consulte [architecture/data-dictionary.md](architecture/data-dictionary.md) 🆕
- **Diagrama de Fluxo e C4 Context:** Consulte [architecture/c4-context.md](architecture/c4-context.md)
- **Integrações, Dependências e Libs Locais:** Consulte [architecture/integrations.md](architecture/integrations.md)
- **Análise Técnica de Fluxo de Código (Handlers/Services):** Consulte [engineering/code-analysis.md](engineering/code-analysis.md)
- **Padrões de API, Erros e Observabilidade:** Consulte [engineering/api-guidelines.md](engineering/api-guidelines.md)

## Requisitos e Escopo de Produto
- **Especificação de Requisitos (RF-01 a RF-10 / RNF / MoSCoW):** Consulte [product/requirements.md](product/requirements.md)
- **Roadmap de Features e Dívidas Técnicas (lista canônica):** Consulte [product/feature-roadmap.md](product/feature-roadmap.md)
- **Contrato OpenAPI da API (Rotas, Schemas, Erros):** Consulte [api/tax-rates-api.yaml](api/tax-rates-api.yaml)

## Regras de Negócio e Domínio Fiscal
- **Glossário Geral de Domínio (todos os tributos, CSTs, CFOPs, NCMs):** Consulte [domain/domain.md](domain/domain.md)

## Governança e Qualidade das Specs
- **Inventário do Projeto e Cobertura de Testes:** Consulte [governance/inventory.md](governance/inventory.md)
- **Dívidas Técnicas (lista canônica):** Consulte [product/feature-roadmap.md](product/feature-roadmap.md#d%C3%ADvidas-t%C3%A9cnicas) (DT-01 a DT-11 com prioridades e localização)
- **Relatório de Confiança e Score da Documentação (99%):** Consulte [governance/confidence-report.md](governance/confidence-report.md)

## Lacunas e Perguntas
- **Histórico de Lacunas Resolvidas:** Consulte [questions/questions_01.md](questions/questions_01.md)

## Esquema de Banco e Dados
- **Schema DDL + Triggers (15 tabelas, PL/pgSQL):** `data/init.sql`
- **Dicionário de Dados (propósito, uso e regras de cada tabela):** Consulte [architecture/data-dictionary.md](architecture/data-dictionary.md) 🆕

## Domínio e Calculadoras
- **Inventário completo de módulos, calculadoras e cobertura de testes:** Consulte [governance/inventory.md](governance/inventory.md)
- **Análise de fluxo de código (handlers, services, patterns):** Consulte [engineering/code-analysis.md](engineering/code-analysis.md)

## Middleware e Infraestrutura
- **W3C Trace Context, JWT Auth, Métricas Prometheus:** Consulte [engineering/code-analysis.md](engineering/code-analysis.md) e [architecture/integrations.md](architecture/integrations.md)

## Prompts Reutilizáveis (Workflows)
- **Prompt #1 — Processar TASKS + Gerar documento de execução:** Consulte `../../../../../.specs/prompts/PROMPT-01-PROCESSAR-TASKS-E-GERAR-DOCUMENTO-DE-EXECUCAO.md`
- **Prompt #2 — Atualizar repositório + Abrir Pull Request:** Consulte `../../../../../.specs/prompts/PROMPT-02-ATUALIZAR-REPOSITORIO-E-ABRIR-PULL-REQUEST.md`
- **Prompt #3 — Minerar e Criar Documentação Técnica (Spec-Mining):** Consulte `../../../../../.specs/prompts/PROMPT-MINING-SPECIFICATION.md` — orquestra 4 skills (domain-modeling → api-designer → architecture-designer → documentation-writer) para criar/atualizar toda a documentação `.specs/`

## Histórico de Pull Requests
- **Histórico de Pull Requests:** Consulte [pull-requests/](pull-requests/)
- **PRs abertas no GitHub:** [github.com/bolismar69/workspace-fbso/pulls](https://github.com/bolismar69/workspace-fbso/pulls)

## Histórico de Implementações (skill-output)
- **16 registros de sessão** documentando cada feature implementada em [skill-output/](skill-output/)
- **9 features implementadas** (F-001 a C-002) — ver [features/FEATURE-2026-06-21.md](features/FEATURE-2026-06-21.md)
- **Gap Analysis (conformidade spec→código):** [features/FEATURE-2026-06-21-GAP-ANALISYS.md](features/FEATURE-2026-06-21-GAP-ANALISYS.md) — 54/55 requisitos ✅
- **Schema DDL (15 tabelas + índices):** `data/init.sql` — listagem completa em [architecture/erd.md](architecture/erd.md) e [architecture/data-dictionary.md](architecture/data-dictionary.md)
