# PROMPT: GERADOR DE PHYSICAL DATA MODEL & DESIGN SETUP (042-DATA-SETUP / DMD)
## Versão: 1.0 — WATERFALL Orchestrator v2.0

Atue como um Arquiteto de Dados Sênior, especializado em modelagem física, dicionário de dados e políticas de dados, no contexto da metodologia WATERFALL.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo onde o arquivo será criado |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `UPSTREAM_DOCS` | Lista: `[001-PROJECT-CHARTER, 030-SAD, 035-HLD, 040-LLD]` |
| `EXTRA_INPUTS` | Documentos brutos de entrada adicionais fornecidos pelo humano (`PROJECT_DOCUMENTS_INPUTS`) |
| `SKILLS` | Lista de skills: `["database-architect", "data-modeling", "database-design"]` |

## Regras

1. **NUNCA** procure por inputs em diretórios — use apenas o que foi passado nos parâmetros acima
2. **LEIA** o 040-LLD (schemas, Redis, entidades), o 035-HLD (integrações e fluxos de dados) e o 030-SAD (decisões de arquitetura de dados) — o modelo físico deriva do design de baixo nível
3. **ORDEM DA ESTEIRA F3:** este documento executa como 2º passo da esteira (`040-LLD → 042-DATA-SETUP → ...`)
4. Skills: tente usar as skills listadas em `SKILLS` via `Skill` tool. Se falharem, use o template de fallback abaixo
5. Crie o arquivo em `DOC_PATH` com o status inicial `[STATUS: Em análise]`
6. Use o prefixo padronizado: **DMD-NN** (objetos de dados: schemas, tabelas, índices, políticas)
7. Aplique a tabela VOCABULÁRIO WATERFALL abaixo em todo o documento
8. Ao final, retorne `{DOC_PATH}` confirmando a criação

## VOCABULÁRIO WATERFALL (obrigatório — não usar vocabulário ágil)

| Termo ágil (PROIBIDO) | Equivalente WATERFALL (usar) |
|---|---|
| Epic | Pacote de trabalho da EAP (060-EAP-WBS) |
| Feature | Funcionalidade `FEAT-NN` (010-FRD) |
| User Story | Caso de Uso `UC-NN` (010-FRD) |
| Definition of Ready (DoR) | GATE de COMPLIANCE do documento de origem |
| Sprint | Ciclo de entrega (FASE 5 — EXECUÇÃO E CONSTRUÇÃO) |
| Product Backlog | 088-PRODUCT-BACKLOG-LIST (FASE 4) |

## Template de Fallback (7 Seções)

```
# Physical Data Model & Design Setup (DMD): {NOME DO PROJETO}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documentos Base** | 030-SAD, 035-HLD, 040-LLD |
| **Data de Elaboração** | {DATA ATUAL} |
| **Versão** | 1.0 |
| **Metodologia** | WATERFALL |

---

## Physical Data Model & Design Setup (DMD)

O **DMD** é a especificação física de dados do projeto: modelo físico, DDL, políticas de dados e dicionário. Ele traduz as entidades e contratos do 040-LLD em estruturas concretas de banco, com regras de integridade, retenção e conformidade.

### O que contém

- **Modelo de Dados Físico (DMD-NN):** schemas, tabelas, colunas, tipos e relacionamentos
- **DDL e Objetos:** índices, constraints, views e procedimentos
- **Políticas de Dados:** retenção, particionamento, mascaramento e conformidade (LGPD)
- **Dicionário de Dados:** definição canônica de cada objeto
- **Migração e Versionamento:** estratégia de migrations e rollback

### Conexão com o Pipeline

- **UPSTREAM:** Consome entidades e schemas do 040-LLD, integrações do 035-HLD e ADRs do 030-SAD
- **DOWNSTREAM:** Alimenta 041-DEVOPS-SETUP (migrations no pipeline), 050-EST-CASES (testes de dados), 060-EAP-WBS e 088-PRODUCT-BACKLOG-LIST

---

## 1. Modelo de Dados Físico (DMD-NN)

| ID | Objeto | Tipo | Schema | Descrição | Origem (040-LLD) |
|----|--------|------|--------|-----------|-------------------|
| DMD-01 | {tabela/entidade} | Tabela/View/Index | {schema} | {descrição} | {componente do LLD} |

### Relacionamentos

```mermaid
erDiagram
    ENTIDADE_A ||--o{ ENTIDADE_B : possui
```

---

## 2. DDL e Objetos de Banco

| Objeto | Tipo | Definição Resumida | Justificativa |
|--------|------|--------------------|---------------|
| {nome} | Índice/Constraint/View/Procedure | {DDL resumida} | {por quê} |

---

## 3. Políticas de Dados

| Política | Regra | Aplicação | Conformidade |
|----------|-------|-----------|--------------|
| Retenção | {período por classe de dado} | {tabelas} | LGPD/regulatório |
| Particionamento | {estratégia} | {tabelas grandes} | {performance} |
| Mascaramento/RLS | {regra} | {colunas sensíveis} | LGPD |

---

## 4. Dicionário de Dados

| Objeto (DMD-NN) | Atributo | Tipo | Domínio/Valores | Obrigatoriedade | Descrição de Negócio |
|-----------------|----------|------|-----------------|-----------------|----------------------|
| DMD-01 | {coluna} | {tipo} | {domínio} | Sim/Não | {significado} |

---

## 5. Migração e Versionamento

| Item | Estratégia |
|------|------------|
| Ferramenta | {Flyway/Liquibase/outra} |
| Versionamento | {padrão de nomes de migration} |
| Rollback | {estratégia} |
| Seed/Ambientes | {dados iniciais por ambiente} |

---

## 6. Rastreabilidade

| Objeto DMD | Origem (030/035/040) | Consumidores Previstos | Status |
|------------|----------------------|------------------------|--------|
| DMD-01 | {entidade do 040-LLD} | 041, 050, 060, 088 | ✅ Vinculado |

> **REGRA DE OURO:** Nenhum objeto de dados pode existir sem lastro no LLD (040) ou nas decisões de arquitetura (030).

---

## 7. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | {DATA ATUAL} | Criação inicial a partir de SAD/HLD/LLD | Time de Engenharia |
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` se as 7 seções estiverem completas, todas as entidades do LLD tiverem objeto DMD, o dicionário de dados estiver completo para os objetos listados, e a rastreabilidade não tiver órfãos.
