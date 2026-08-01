# PROMPT: GENERATE — DOWNSTREAM-ARCHITECTURE-REFINEMENT — DATA-ARCHITECTURE-DEFINITION (F4)
## Versão: 1.1 — Arquitetura de Dados Detail-Level (ERD + Query Patterns + Migrations) — Independente de Tecnologia

Atue como um Data Architect e Especialista em Modelagem de Dados, com expertise em sistemas multi-tenant e estratégias de isolamento de dados.

## OBJETIVO

Produzir a definição de arquitetura de dados em nível de implementação: modelo de dados completo, estratégia de isolamento multi-tenant, query patterns otimizadas, estratégia de versionamento de schema.

**Este documento é independente de tecnologias específicas de banco de dados.** Durante a análise da stack do projeto, identifique o banco de dados utilizado e busque skills relacionados a esse banco para aprimorar as especificações. Caso não encontre skills específicos, utilize skills generalistas de arquitetura de dados.

## INPUTS

1. **Arquitetura Detail-Level** (F2)
2. **Segurança Detail-Level** (F3) — políticas de isolamento
3. **Features relevantes:** dashboards, tenants, RBAC, auditoria

## ESTRUTURA DO DOCUMENTO

```markdown
# DETAIL-LEVEL-DATA-ARCHITECTURE-DEFINITION — Arquitetura de Dados Detail-Level

## 1. Modelo de Dados Completo
[Todas as entidades: atributos, tipos, constraints, índices, relacionamentos]
Esquemas e domínios de dados

## 2. Estratégia de Multi-Tenancy
- Abordagem de isolamento (discriminator column, RLS, schema-per-tenant, database-per-tenant)
- Políticas de isolamento por tabela
- Verificação de isolamento

## 3. Estratégia de Particionamento e Arquivamento
- Tabelas de alto volume: estratégia de particionamento
- Política de retenção e arquivamento

## 4. Query Patterns Otimizadas
[Índices, views materializadas, estratégias de cache para queries críticas]

## 5. Estratégia de Versionamento de Schema
- Ferramenta de migration
- Versionamento e baseline
- Estratégia de rollback

## 6. Volumes e Crescimento
[Projeções: year 1-3, tenants, users, transactions, storage]

## 7. Riscos de Dados
```

### Skills Recomendados

**Skills generalistas de dados (sempre aplicáveis):**
- `senior-data-engineer`, `database-architect`, `database-design`
- `data-engineer`, `data-modeling`, `database`
- `database-migrations`, `database-migrations-sql-migrations`

**Skills tecnológicos de banco de dados (condicionais — buscar ao identificar a stack):**
- Ao identificar um banco de dados específico durante a análise da stack, busque skills relacionados a esse banco para aprimorar as especificações e estimativas
- Caso não encontre skills específicos para o banco identificado, utilize os skills generalistas listados acima como referência

🤖 *Prompt gerador — Fase 4 do Downstream Architecture Refinement · Independente de Tecnologia*
