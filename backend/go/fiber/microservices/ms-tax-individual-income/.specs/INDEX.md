# 🗺️ Mapa de Especificações — ms-tax-individual-income

Diretório centralizado de arquitetura de contexto e especificações do microsserviço de cálculo de tributos individuais.

## 🧱 Engenharia e Arquitetura
- **Visão Arquitetural e Regras de Solução:** Consulte [architecture/architecture.md](architecture/architecture.md)
- **Modelo de Dados (ERD Completo):** Consulte [architecture/erd.md](architecture/erd.md)
- **Diagrama de Fluxo e C4 Context:** Consulte [architecture/c4-context.md](architecture/c4-context.md)
- **Integrações, Dependências e Libs Locais:** Consulte [architecture/integrations.md](architecture/integrations.md)
- **Análise Técnica de Fluxo de Código (Handlers/Services):** Consulte [engineering/code-analysis.md](engineering/code-analysis.md)

## 🎯 Requisitos e Escopo de Produto
- **Especificação de Requisitos (RF-01 a RF-06 / RNF / MoSCoW):** Consulte [product/requirements.md](product/requirements.md)
- **Roadmap de Features Planejadas:** Consulte [product/feature-roadmap.md](product/feature-roadmap.md)
- **Contrato OpenAPI da API (Rotas, Schemas, Erros, Health):** Consulte [api/tax-api.yaml](api/tax-api.yaml)

## ⚖️ Regras de Negócio e Domínio Fiscal
- **Glossário Geral de Domínio (Inclui INSS e Motor Progressivo):** Consulte [domain/domain.md](domain/domain.md)
- **Regra de Transição da Reforma Tributária 2026:** Consulte [domain/reform-2026.md](domain/reform-2026.md)

## 📋 Governança e Qualidade das Specs
- **Inventário do Projeto e Cobertura de Testes:** Consulte [governance/inventory.md](governance/inventory.md)
- **Relatório de Confiança e Score da Documentação (🟢 92%):** Consulte [governance/confidence-report.md](governance/confidence-report.md)

## ❓ Lacunas e Perguntas
- **Histórico de Lacunas Resolvidas:** Consulte [questions/questions_01.md](questions/questions_01.md)

## 📝 Logs de Implementação
- **Cache Redis (nil-safety, logging, TTL configurável):** Consulte [skill-output/2026-06-20_174207.md](skill-output/2026-06-20_174207.md)
- **Health Check Endpoints (/healthz, /health):** Consulte [skill-output/2026-06-20_182911.md](skill-output/2026-06-20_182911.md)
