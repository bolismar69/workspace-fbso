# ADR 0001: Orquestração de Workflows com Kestra no Monorepo

## 📅 Data
2026-01-22

## 💡 Status
Aceito

## 🎯 Contexto

A organização precisa orquestrar processos de dados e integrações que atravessam múltiplos domínios, incluindo:

- ETL/ELT (pipelines batch e event-driven)
- Hub de Integrações (parceiros, integrações externas, integrações internas)
- Integrações enterprise e migração/convivência com stacks legadas (ex.: Oracle Suite)
- Automações operacionais (rotinas, jobs, backfills, reconciliações)

A ausência de uma padronização para definição, versionamento, validação e implantação de workflows tende a gerar:

- Definições dispersas (YAML/scripts em repositórios diferentes, sem rastreabilidade)
- Fragilidade em deploy (mudanças manuais em ambiente)
- Falta de governança (ownership, revisão e padrões de secrets)

## ✅ Decisão

Adotar **Kestra** como ferramenta padrão de orquestração de workflows e definir uma estrutura clara no monorepo para:

1. **Definições de fluxo e ativos de execução** (YAML, SQL, Python, Shell etc.) em `orchestration/kestra/`.
2. **Deploy/runtime de instâncias Kestra** (infra, helm/terraform, observabilidade, sizing) em `devops/`.
3. **Implantação (subida) de flows e ativos** executada **preferencialmente via CI/CD**, evitando operações manuais.

### Estrutura de Pastas

- `orchestration/`
  - `kestra/`
    - `enterprise-integrations/`
    - `integration-hub/`
    - `etl/`
    - `automation/`

### Convenções

**1) Namespace e ID**

- `namespace`: `fbso.<dominio>.<sistema>.<contexto>`
  - exemplos: `fbso.etl.salesforce.orders`, `fbso.integration_hub.partner_x.onboarding`
- `id`: `snake_case`, descritivo e curto
  - exemplo: `sync_orders_to_dwh`

**2) Scripts e SQL versionados**

- Flows YAML devem focar em orquestração.
- Lógica pesada deve residir em `scripts/` (ex.: SQL/Python/Shell) e ser referenciada pelo flow.

**3) Secrets e Configuração**

- Segredos **não** devem ser versionados no git.
- Flows devem referenciar secrets/variáveis de ambiente gerenciadas no ambiente Kestra.
- Cada domínio/projeto deve documentar no README quais chaves são obrigatórias.

**4) CI/CD como caminho padrão**

- Mudanças em `orchestration/kestra/**` devem disparar pipeline de publicação.
- Publicação deve ser idempotente (atualiza se já existir) e falhar o pipeline em caso de erro de validação do Kestra.

## 📐 Consequências

### Positivas

- **Rastreabilidade e auditabilidade**: toda mudança de flow/script passa por PR.
- **Governança clara**: ownership e convenções padronizadas.
- **Menos risco operacional**: publicação automatizada via CI/CD.

### Negativas

- **Curva de aprendizado**: time precisa padronizar authoring e operação do Kestra.
- **Disciplina de organização**: exige namespaces consistentes e revisão contínua.

## 🔗 Referências

- Estrutura de orquestração: `orchestration/kestra/`
- Deploy/infra: `devops/orchestration/kestra/`
