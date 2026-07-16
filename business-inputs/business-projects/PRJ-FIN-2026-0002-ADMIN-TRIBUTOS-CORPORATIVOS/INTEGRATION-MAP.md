# Mapa de Integrações — Portal de Gestão Tributária

- **Projeto:** PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS
- **Programa Pai:** PRJ-FIN-2026-0001 — Adequação Corporativa à Reforma Tributária Nacional
- **Data de Criação:** 11 de Julho de 2026
- **Versão:** 1.0
- **Status:** Definição Inicial
- **Referências:**
  - [01-PROJECT-CHARTER.md](./01-PROJECT-CHARTER.md) — escopo e entregas de negócio
  - [API-CONTRACTS.md](./API-CONTRACTS.md) — contrato de API entre frontend e backend
  - [../backend/go/fiber/microservices/ms-billing-engine-tax-rates/.specs/architecture/erd.md](../backend/go/fiber/microservices/ms-billing-engine-tax-rates/.specs/architecture/erd.md) — modelo de dados existente (15 tabelas)
  - [../backend/go/fiber/microservices/ms-billing-engine-tax-rates/.specs/architecture/data-dictionary.md](../backend/go/fiber/microservices/ms-billing-engine-tax-rates/.specs/architecture/data-dictionary.md) — dicionário de dados

---

## 1. Objetivo

Este documento define o **mapa completo de integrações** entre todos os componentes do sistema de gestão tributária, abrangendo:

- A comunicação entre o **novo portal** (frontend React), o **novo microserviço de administração** (backend Java/Spring — DT-1) e o **motor de cálculo existente** (Go/Fiber — DT-3 após refatoração)
- O relacionamento com o **banco de dados compartilhado** (PostgreSQL, schema `billing_tax_rates`), com as **novas tabelas** a serem criadas pelo DT-1 e as **colunas de multi-tenancy** a serem adicionadas às tabelas existentes
- A integração com o **provedor de identidade corporativo** (Keycloak / SAML 2.0)
- As dependências de **infraestrutura** (Kubernetes, Secrets Manager)
- Os requisitos de **segurança** por canal de comunicação

---

## 2. Diagrama de Contexto (C4 — Nível 1)

```
┌─────────────────────────────────────────────────────────────────────────────────────────┐
│                               USUÁRIOS DO NEGÓCIO                                        │
│                                                                                          │
│   ┌──────────────┐   ┌──────────────┐   ┌──────────────┐   ┌──────────────────┐        │
│   │  Analista    │   │ Administrador│   │   Auditor /  │   │ Comitê Fiscal /  │        │
│   │  Fiscal      │   │   Fiscal     │   │  Controller  │   │      CFO         │        │
│   └──────┬───────┘   └──────┬───────┘   └──────┬───────┘   └────────┬─────────┘        │
│          │                  │                  │                    │                   │
│          └──────────────────┼──────────────────┼────────────────────┘                   │
│                             │                  │                                        │
└─────────────────────────────┼──────────────────┼────────────────────────────────────────┘
                              │   HTTPS (443)    │
                              ▼                  ▼
┌─────────────────────────────────────────────────────────────────────────────────────────┐
│                         SISTEMA DE GESTÃO TRIBUTÁRIA                                      │
│                                                                                          │
│  ┌─────────────────────────────────────┐    ┌──────────────────────────────────┐        │
│  │        DT-2: Portal Frontend        │    │   Provedor de Identidade          │        │
│  │        React 19 + Vite              │◄───│   Keycloak (SAML 2.0)             │        │
│  │        web_app-billing-admin-       │    │   Autenticação + Autorização      │        │
│  │        tax-rates                    │    │   (RBAC: 3 perfis)                │        │
│  └────────────┬────────────────────────┘    └──────────────────────────────────┘        │
│               │                                                                          │
│               │  HTTPS (443) — REST/JSON                                                 │
│               │  Header: Authorization: Bearer <SAML-token>                              │
│               ▼                                                                          │
│  ┌─────────────────────────────────────┐                                                 │
│  │     DT-1: Backend Admin             │                                                 │
│  │     Java 21 + Spring Boot 4.0.1     │                                                 │
│  │     ms-billing-admin-tax-rates      │                                                 │
│  │                                      │                                                 │
│  │  ┌────────────────────────────────┐ │                                                 │
│  │  │ Recursos (API REST):           │ │                                                 │
│  │  │ • /api/v1/aliquotas            │ │                                                 │
│  │  │ • /api/v1/classificacoes       │ │                                                 │
│  │  │ • /api/v1/regimes              │ │                                                 │
│  │  │ • /api/v1/usuarios             │ │                                                 │
│  │  │ • /api/v1/auditoria            │ │                                                 │
│  │  │ • /api/v1/lotes                │ │                                                 │
│  │  │ • /api/v1/aprovacoes           │ │                                                 │
│  │  │ • /api/v1/relatorios           │ │                                                 │
│  │  │ • /api/v1/empresas             │ │                                                 │
│  │  └────────────────────────────────┘ │                                                 │
│  └──────┬───────────────┬──────────────┘                                                 │
│         │               │                                                                │
│         │ TCP/5432      │ HTTPS (443) — REST/JSON                                        │
│         │ (pgJDBC)      │ (leitura após DT-3 refatoração)                                │
│         ▼               ▼                                                                │
│  ┌──────────────┐  ┌──────────────────────────────────────┐                              │
│  │ PostgreSQL   │  │  DT-3: Motor de Cálculo (Existente)  │                              │
│  │ Schema:      │  │  Go + Fiber                          │                              │
│  │ billing_tax_ │  │  ms-billing-engine-tax-rates         │                              │
│  │ rates        │  │                                      │                              │
│  │              │  │  Após refatoração:                   │                              │
│  │ [Tabelas     │  │  • Remove endpoints admin            │                              │
│  │  existentes] │  │  • Mantém SELECT nas tabelas         │                              │
│  │ [Novas       │  │    administrativas                   │                              │
│  │  tabelas     │  │  • Motor permanece apenas            │                              │
│  │  DT-1]       │  │    para cálculo de impostos          │                              │
│  └──────────────┘  └──────────────────────────────────────┘                              │
│                                                                                          │
└─────────────────────────────────────────────────────────────────────────────────────────┘
```

---

## 3. Inventário de Componentes

| # | Componente | Tipo | Stack | Responsabilidade | Fase de Entrada |
|:---|:---|:---|:---|:---|:---|
| **C1** | Portal Frontend (DT-2) | Aplicação Web | React 19 + Vite | Interface do time de Finanças para gestão de tabelas fiscais | Fase 2 (Sprint 2) |
| **C2** | Backend Admin (DT-1) | Microserviço REST | Java 21 + Spring Boot 4.0.1 | CRUD de alíquotas, classificações, regimes, usuários, lotes, aprovações, relatórios e auditoria | Fase 2 (Sprint 2) |
| **C3** | Motor de Cálculo (DT-3) | Microserviço REST | Go + Fiber | Cálculo de impostos em operações de venda/faturamento. Após refatoração: somente leitura das tabelas administrativas | Existente; refatoração na Fase 3 |
| **C4** | PostgreSQL | Banco de Dados | PostgreSQL (via pgJDBC / pgx) | Schema `billing_tax_rates` — 15 tabelas existentes + novas tabelas do DT-1 | Fase 0 (Sprint 0 — definições de acesso) |
| **C5** | Keycloak | Provedor de Identidade | Keycloak Corporativo | Autenticação SAML 2.0, autorização RBAC (3 perfis), emissão de tokens | Existente; integração configurada na Fase 1 |
| **C6** | Redis | Cache | Redis | Cache de regras fiscais para o motor de cálculo (DT-3). Não utilizado pelo DT-1 diretamente | Existente |

---

## 4. Canais de Integração

### 4.1 Portal Frontend (DT-2) → Backend Admin (DT-1)

| Atributo | Valor |
|:---|:---|
| **Protocolo** | HTTPS / REST JSON |
| **Direção** | Frontend → Backend (requisições); Backend → Frontend (respostas) |
| **Autenticação** | Token SAML 2.0 via header `Authorization: Bearer <token>`. Token obtido do Keycloak (C5) no login do portal |
| **Autorização** | Baseada no perfil do usuário (RBAC): `ADMINISTRADOR_FISCAL`, `ANALISTA_FISCAL`, `AUDITOR_CONTROLLER` |
| **Contrato** | [API-CONTRACTS.md](./API-CONTRACTS.md) — 9 grupos de recursos, ~35 endpoints |
| **Formato de erros** | JSON padronizado com `codigo`, `mensagem`, `detalhes`, `timestamp`, `trace_id` |
| **Timeout** | 30 segundos (operações CRUD); 120 segundos (carga em lote e geração de relatórios) |
| **Rate Limiting** | 100 requisições por minuto por usuário (configurável) |
| **Mock** | MSW (Mock Service Worker) local no frontend durante Fase 0-2, baseado no OpenAPI YAML |

### 4.2 Backend Admin (DT-1) → PostgreSQL (C4)

| Atributo | Valor |
|:---|:---|
| **Protocolo** | TCP/5432 — PostgreSQL wire protocol |
| **Driver** | pgJDBC (JDBC 4.3) |
| **Schema** | `billing_tax_rates` (compartilhado com DT-3) |
| **Credenciais** | Obtidas via Secrets Manager corporativo (Vault). **Não armazenar** em `application.yml` ou variáveis de ambiente planas |
| **Pool de Conexões** | HikariCP (gerenciado pelo Spring Boot). Mínimo 5, máximo 20 conexões |
| **Transações** | Read committed (padrão PostgreSQL). Operações de carga em lote usam transações por lote |
| **Permissões DT-1** | `SELECT`, `INSERT`, `UPDATE`, `DELETE` nas tabelas administrativas (ver Seção 5) |
| **Permissões DT-3 (pós-refatoração)** | Apenas `SELECT` nas tabelas administrativas; mantém `SELECT`, `INSERT`, `UPDATE` nas tabelas operacionais e de cálculo |

### 4.3 Portal Frontend (DT-2) e Backend Admin (DT-1) → Keycloak (C5)

| Atributo | Valor |
|:---|:---|
| **Protocolo** | SAML 2.0 |
| **Fluxo** | 1. Usuário acessa o portal → redirecionado ao Keycloak para login; 2. Keycloak emite SAML assertion; 3. Portal armazena token; 4. Toda requisição ao backend inclui o token no header `Authorization` |
| **Perfis (RBAC)** | `ADMINISTRADOR_FISCAL`, `ANALISTA_FISCAL`, `AUDITOR_CONTROLLER` — mapeados como roles no Keycloak |
| **Validação no Backend** | O backend valida o token a cada requisição via introspecção no Keycloak. Middleware extrai perfil e injeta no contexto da requisição |
| **SLO (Single Logout)** | Suportado — logout no portal invalida a sessão SAML no Keycloak |

### 4.4 Backend Admin (DT-1) → Motor de Cálculo (DT-3) (pós-refatoração)

| Atributo | Valor |
|:---|:---|
| **Protocolo** | HTTPS / REST JSON (leitura) |
| **Direção** | DT-1 consulta DT-3 para verificar consistência entre alíquotas cadastradas e alíquotas em uso nos cálculos |
| **Endpoint DT-3** | `GET /v1/admin/tax-rates/iva-dual` (existente, mantido como leitura) |
| **Frequência** | Sob demanda (reconciliação manual) e programada (job diário de verificação de consistência — ver Seção 8) |
| **Autenticação** | Service-to-service: client credentials grant via Keycloak (client ID + secret) |
| **Timeout** | 15 segundos |
| **Fallback** | Se DT-3 estiver indisponível, a reconciliação é reagendada. O portal (DT-1 + DT-2) não depende do DT-3 para operações normais |

### 4.5 Motor de Cálculo (DT-3) → PostgreSQL (C4) (pós-refatoração)

| Atributo | Valor |
|:---|:---|
| **Permissões mantidas** | `SELECT` em todas as tabelas (administrativas e operacionais); `INSERT`, `UPDATE`, `DELETE` apenas nas tabelas operacionais/de cálculo (`tax_tokens`, `fornecedor_fiscal`) |
| **Permissões revogadas** | `INSERT`, `UPDATE`, `DELETE` nas tabelas administrativas (alíquotas, classificações, regimes, lotes, usuários, auditoria) — ver Seção 9 para o procedimento de revogação |
| **Cache** | Redis (C6) — cache de regras fiscais com TTL 24h, chaves `tax:iva:<ncm>:<uf>:<municipio>`, `tax:icms:<orig>:<dest>`, `tax:federal:<regime>:<cstPis>:<cstCofins>` |

---

## 5. Modelo de Dados — Tabelas Existentes e Novas

### 5.1 Schema `billing_tax_rates` — Visão Completa

O schema é **compartilhado** entre DT-1 (admin) e DT-3 (motor de cálculo). As tabelas estão organizadas em quatro categorias:

| Categoria | Tabelas | Dono | Descrição |
|:---|:---|:---|:---|
| **Regime Atual (7 tab.)** | `icms_rules`, `federal_tax_rules`, `product_tax_exceptions`, `tax_equivalence`, `simples_nacional_rates`, `ipi_regras`, `iss_rates` | DT-1 (admin) + DT-3 (leitura) | Regras dos tributos pré-reforma |
| **Reforma Tributária (6 tab.)** | `iva_dual_rules`, `iva_dual_rules_log`, `reforma_tributaria_rules` (legado), `ncm_seletivo`, `cbs_rates`, `cst_reforma` | DT-1 (admin) + DT-3 (leitura) | Regras do IVA Dual (CBS/IBS/IS) |
| **Operacional (2 tab.)** | `tax_tokens`, `fornecedor_fiscal` | DT-3 (leitura/escrita) | Suporte a cálculo — **não administradas pelo DT-1** |
| **Novas Tabelas DT-1 (6 tab.)** | `empresas`, `tenants`, `fornecedores`, `lotes_carga`, `lotes_carga_itens`, `auditoria_log` | DT-1 (admin) + DT-3 (leitura) | Infraestrutura multi-tenancy, carga em lote, auditoria unificada |

### 5.2 Novas Tabelas — DT-1

#### 5.2.1 `empresas`

**Propósito:** Raiz do multi-tenancy. Cada empresa do grupo econômico é cadastrada aqui. As tabelas de alíquotas e classificações são segmentadas por `empresa_id`.

| Coluna | Tipo | Descrição |
|:---|:---|:---|
| `id` | serial PK | Identificador único |
| `cnpj_raiz` | varchar(8) | CNPJ raiz (8 primeiros dígitos) |
| `razao_social` | varchar(200) | Razão social completa |
| `nome_fantasia` | varchar(100) | Nome fantasia |
| `status` | varchar(20) | `ATIVA`, `INATIVA` |
| `criado_em` | timestamp | Data de criação |
| `atualizado_em` | timestamp | Data de atualização |

#### 5.2.2 `tenants`

**Propósito:** Estabelecimentos (filiais, unidades) dentro de uma empresa. Permite granularidade de alíquotas por CNPJ completo quando necessário.

| Coluna | Tipo | Descrição |
|:---|:---|:---|
| `id` | serial PK | Identificador único |
| `empresa_id` | integer FK → `empresas.id` | Empresa à qual o tenant pertence |
| `cnpj_completo` | varchar(14) | CNPJ completo (14 dígitos) |
| `nome` | varchar(200) | Nome do estabelecimento |
| `uf` | varchar(2) | UF do estabelecimento |
| `municipio_ibge` | varchar(7) | Código IBGE do município |
| `status` | varchar(20) | `ATIVO`, `INATIVO` |

#### 5.2.3 `fornecedores`

**Propósito:** Cadastro corporativo de fornecedores com qualificação fiscal expandida. Complementa a tabela `fornecedor_fiscal` existente (que é focada em crédito), adicionando dimensões de negócio.

| Coluna | Tipo | Descrição |
|:---|:---|:---|
| `id` | serial PK | Identificador único |
| `empresa_id` | integer FK → `empresas.id` | Empresa do grupo |
| `cnpj` | varchar(14) | CNPJ do fornecedor |
| `razao_social` | varchar(200) | Razão social |
| `regime_tributario` | varchar(30) | Regime (Lucro Real, Lucro Presumido, Simples Nacional) |
| `cnae_principal` | varchar(7) | CNAE fiscal principal |
| `status` | varchar(20) | `ATIVO`, `PENDENTE`, `BLOQUEADO` |
| `criado_em` / `atualizado_em` | timestamp | Datas de criação e atualização |

> **Nota:** A tabela `fornecedor_fiscal` existente (schema `billing_tax_rates`) permanece como catálogo operacional do motor de cálculo (DT-3). A nova tabela `fornecedores` é o cadastro mestre de negócio, gerenciado pelo DT-1. A reconciliação entre ambas é tratada na Seção 8.

#### 5.2.4 `lotes_carga`

**Propósito:** Cabeçalho dos lotes de importação de alíquotas. Cada arquivo enviado pelo time fiscal gera um registro aqui. Os dados só são efetivados nas tabelas finais após **aprovação**.

| Coluna | Tipo | Descrição |
|:---|:---|:---|
| `id` | serial PK | Identificador único do lote |
| `empresa_id` | integer FK → `empresas.id` | Empresa destinatária |
| `tributo` | varchar(10) | Tributo das alíquotas no lote (IBS, CBS, ICMS, etc.) |
| `nome_arquivo` | varchar(255) | Nome original do arquivo enviado |
| `status` | varchar(25) | `EM_VALIDACAO` → `AGUARDANDO_APROVACAO` → `APROVADO` ou `REJEITADO` |
| `total_linhas` | integer | Total de linhas no arquivo |
| `linhas_aceitas` | integer | Linhas validadas com sucesso |
| `linhas_rejeitadas` | integer | Linhas com erro de validação |
| `linhas_com_alertas` | integer | Linhas aceitas mas com alertas (ex: RN-03 no Período Híbrido) |
| `enviado_por` | varchar(100) | Usuário que enviou o arquivo |
| `aprovado_por` | varchar(100) | Administrador Fiscal que aprovou (nulo até aprovação) |
| `justificativa` | text | Justificativa fornecida no envio |
| `enviado_em` | timestamp | Data/hora do envio |
| `aprovado_em` | timestamp | Data/hora da aprovação (nulo até lá) |

#### 5.2.5 `lotes_carga_itens`

**Propósito:** Linhas individuais do arquivo de carga. Cada linha é validada contra as regras de negócio (RN-01 a RN-05) e classificada como `ACEITO`, `REJEITADO` ou `COM_ALERTA`.

| Coluna | Tipo | Descrição |
|:---|:---|:---|
| `id` | serial PK | Identificador único do item |
| `lote_id` | integer FK → `lotes_carga.id` | Lote ao qual o item pertence |
| `numero_linha` | integer | Número da linha no arquivo original |
| `conteudo_original` | jsonb | Conteúdo completo da linha (para reprocessamento e auditoria) |
| `status` | varchar(20) | `ACEITO`, `REJEITADO`, `COM_ALERTA` |
| `motivo_rejeicao` | text | Motivo da rejeição (RN violada), nulo se aceito |
| `entidade_criada_tipo` | varchar(30) | Tipo da entidade criada após aprovação (ex: `ALIQUOTA_IBS`) |
| `entidade_criada_id` | integer | ID da entidade criada na tabela final (nulo até aprovação) |

#### 5.2.6 `auditoria_log`

**Propósito:** Trilha de auditoria unificada para **todas** as entidades gerenciadas pelo DT-1. Substitui o modelo fragmentado atual (onde apenas `iva_dual_rules` possui `iva_dual_rules_log`).

| Coluna | Tipo | Descrição |
|:---|:---|:---|
| `id` | bigserial PK | Identificador único do evento |
| `entidade_tipo` | varchar(30) | Tipo da entidade: `ALIQUOTA`, `CLASSIFICACAO`, `REGIME`, `USUARIO`, `LOTE`, `EMPRESA`, `FORNECEDOR` |
| `entidade_id` | integer | ID da entidade afetada |
| `operacao` | varchar(15) | `CRIACAO`, `EDICAO`, `DESATIVACAO`, `APROVACAO`, `REJEICAO` |
| `usuario_id` | integer | ID do usuário autenticado |
| `usuario_nome` | varchar(100) | Nome do usuário (desnormalizado para consulta) |
| `usuario_perfil` | varchar(30) | Perfil do usuário no momento da ação |
| `estado_anterior` | jsonb | Snapshot completo do estado anterior (null em CRIACAO) |
| `estado_novo` | jsonb | Snapshot completo do novo estado (null em DESATIVACAO sem substituição) |
| `justificativa` | text | Justificativa fornecida pelo usuário |
| `ip_origem` | varchar(45) | Endereço IP de origem da requisição |
| `data_hora` | timestamp | Timestamp exato da ação (RN-15) |

**Regras de imutabilidade (RN-14):**
- Nenhum perfil de usuário pode alterar ou excluir registros desta tabela
- Triggers de banco impedem `UPDATE` e `DELETE` na tabela (`REVOKE UPDATE, DELETE ON auditoria_log FROM <todos>`)
- Retenção: mínimo de 5 anos (RN-16); partição por ano para gerenciamento de volume

### 5.3 Colunas de Multi-Tenancy e Rastreabilidade — Tabelas Existentes

As seguintes colunas devem ser **adicionadas** às tabelas de regras fiscais existentes para suportar multi-tenancy e rastreabilidade de origem:

| Coluna | Tipo | Adicionada a | Descrição |
|:---|:---|:---|:---|
| `empresa_id` | integer FK → `empresas.id` | `icms_rules`, `federal_tax_rules`, `product_tax_exceptions`, `iss_rates`, `ipi_regras`, `iva_dual_rules`, `ncm_seletivo`, `cbs_rates` | Empresa do grupo econômico à qual a regra se aplica |
| `tenant_id` | integer FK → `tenants.id` (nullable) | Mesmas tabelas acima | Estabelecimento específico (null = aplica a todos os estabelecimentos da empresa) |
| `origem_cadastro` | varchar(10) DEFAULT 'MANUAL' | Mesmas tabelas acima | `MANUAL` (cadastro via formulário) ou `LOTE` (carga em lote) |
| `lote_origem_id` | integer FK → `lotes_carga.id` (nullable) | Mesmas tabelas acima | Se origem = LOTE, aponta para o lote |
| `lote_item_origem_id` | integer FK → `lotes_carga_itens.id` (nullable) | Mesmas tabelas acima | Se origem = LOTE, aponta para o item específico no lote |

**Estratégia de migração para tabelas existentes:**
- Colunas são adicionadas como **nullable** inicialmente
- Registros existentes recebem `empresa_id` = 1 (empresa padrão) e `origem_cadastro` = `MANUAL`
- Após a migração e validação, as colunas `empresa_id` passam a ser **NOT NULL** para novos registros

---

## 6. Requisitos de Segurança por Canal

| Canal | Autenticação | Criptografia | Autorização | Auditoria |
|:---|:---|:---|:---|:---|
| **DT-2 → DT-1** (portal → backend) | SAML 2.0 Bearer token | HTTPS/TLS 1.3 | RBAC (3 perfis) via Keycloak roles | Token validado a cada requisição; IP de origem registrado na auditoria |
| **DT-1 → PostgreSQL** | User/password via Secrets Manager | TLS (pgJDBC SSL mode: verify-full) | Usuário de banco com grants específicos por operação | Trigger de auditoria em todas as tabelas administrativas |
| **DT-2 / DT-1 → Keycloak** | SAML 2.0 assertion | HTTPS/TLS 1.3 | Roles mapeadas nos clients do Keycloak | Logs de autenticação no Keycloak |
| **DT-1 → DT-3** (reconciliação) | Client credentials grant (OAuth2) | HTTPS/TLS 1.3 | Client ID + secret via Secrets Manager | Requisições registradas no log de integração |
| **DT-3 → PostgreSQL** (pós-refatoração) | User/password via Secrets Manager | TLS | Apenas SELECT nas tabelas admin; INSERT/UPDATE/DELETE revogados | Logs de acesso no PostgreSQL |

---

## 7. Catálogo de Endpoints por Componente

### 7.1 DT-1 — Backend Admin (a construir)

Ver [API-CONTRACTS.md](./API-CONTRACTS.md) Seção 6 para a matriz completa de endpoints × perfis.

**Resumo:**
- `GET/POST/PUT/PATCH /api/v1/aliquotas[/{id}][/desativar][/historico]` — CRUD de alíquotas
- `GET/POST/PUT/PATCH /api/v1/classificacoes[/{id}][/desativar]` — CRUD de classificações
- `GET/POST/PUT/PATCH /api/v1/regimes[/{id}]` — CRUD de regimes
- `GET/POST/PUT/PATCH /api/v1/usuarios[/{id}][/desativar]` — Gestão de usuários
- `GET /api/v1/auditoria[/{id}]` — Consulta de trilha de auditoria
- `POST/GET /api/v1/lotes[/{id}][/itens][/aprovar][/rejeitar]` — Carga em lote
- `GET/POST /api/v1/aprovacoes/pendentes[/{id}]/(aprovar|rejeitar)` — Fluxos de aprovação
- `GET /api/v1/relatorios/*` — Relatórios e dashboards
- `GET/POST/PUT/PATCH /api/v1/empresas[/{id}][/desativar][/tenants]` — Gestão corporativa

### 7.2 DT-3 — Motor de Cálculo (existente, a refatorar)

| Método | Path | Status Pós-Refatoração |
|:---|:---|:---|
| `POST` | `/v1/calculate` | ✅ Mantido — cálculo de impostos |
| `POST` | `/v1/simulate` | ✅ Mantido — simulação |
| `POST` | `/v1/token/generate` | ✅ Mantido — congelamento de alíquotas |
| `GET` | `/v1/token/{id}` | ✅ Mantido — validação de token |
| `POST` | `/v1/credit/calculate` | ✅ Mantido — cálculo de créditos |
| `POST` | `/v1/supplier/validate` | ✅ Mantido — validação de fornecedor |
| `GET` | `/v1/supplier/{cnpj}` | ✅ Mantido — consulta de fornecedor |
| `GET` | `/v1/admin/tax-rates/iva-dual` | ⚠️ Mantido apenas como leitura (consulta) |
| `POST/PUT/DELETE` | `/v1/admin/*` (demais admin) | 🔴 **Removidos** — migrados para DT-1 |

---

## 8. Fluxos de Integração entre Componentes

### 8.1 Fluxo: Reconciliação Portal × Motor de Cálculo

```
[Job diário — 02:00 AM]                     [DT-1 Backend]
        │                                        │
        │  1. Consulta todas as alíquotas       │
        │     vigentes no PostgreSQL             │
        │                                        │
        │  2. Para cada par (NCM, UF, mun):     │
        │     GET /v1/admin/tax-rates/iva-dual   │──► [DT-3 Motor]
        │     ?ncm={ncm}&uf={uf}&mun={mun}       │
        │                                        │
        │  3. Compara alíquotas:                 │
        │     Portal vs. Motor                   │
        │                                        │
        │  4. Se divergência:                    │
        │     • Gera alerta no painel            │
        │     • Notifica Gerente Fiscal          │
        │     • Registra em auditoria_log        │
```

### 8.2 Fluxo: Login e Autorização

```
[Analista Fiscal]──►[Portal DT-2]──►[Keycloak]  (redireciona para login SAML)
                                          │
[Analista Fiscal]◄──[Keycloak]  (formulário de login)
                                          │
[Analista Fiscal]──►[Keycloak]  (credenciais)
                                          │
[Portal DT-2]◄──[Keycloak]  (SAML assertion + token)
                                          │
[Portal DT-2]──►[DT-1 Backend]  (GET /api/v1/aliquotas + Bearer token)
                                          │
[DT-1 Backend]──►[Keycloak]  (introspect token)
                                          │
[DT-1 Backend]──►[Portal DT-2]  (200 + dados de alíquotas)
```

### 8.3 Fluxo: Carga em Lote com Aprovação

Ver [API-CONTRACTS.md](./API-CONTRACTS.md) Seção 5.2 para o fluxo detalhado.

---

## 9. Procedimento de Revogação de Acessos — DT-3

Após o deploy do DT-1 em staging e a migração das operações administrativas para o novo portal, o acesso do motor de cálculo (DT-3) às tabelas administrativas será reduzido ao perfil **somente leitura**. O procedimento abaixo deve ser executado pelos times de DBA, DevOps e DevSecOps.

### 9.1 Grants a Revogar (usuário do DT-3)

| Tabela | Grants Revogados | Grant Mantido |
|:---|:---|:---|
| `icms_rules` | `INSERT`, `UPDATE`, `DELETE` | `SELECT` |
| `federal_tax_rules` | `INSERT`, `UPDATE`, `DELETE` | `SELECT` |
| `product_tax_exceptions` | `INSERT`, `UPDATE`, `DELETE` | `SELECT` |
| `tax_equivalence` | `INSERT`, `UPDATE`, `DELETE` | `SELECT` |
| `simples_nacional_rates` | `INSERT`, `UPDATE`, `DELETE` | `SELECT` |
| `ipi_regras` | `INSERT`, `UPDATE`, `DELETE` | `SELECT` |
| `iss_rates` | `INSERT`, `UPDATE`, `DELETE` | `SELECT` |
| `iva_dual_rules` | `INSERT`, `UPDATE`, `DELETE` | `SELECT` |
| `ncm_seletivo` | `INSERT`, `UPDATE`, `DELETE` | `SELECT` |
| `cbs_rates` | `INSERT`, `UPDATE`, `DELETE` | `SELECT` |
| `cst_reforma` | `INSERT`, `UPDATE`, `DELETE` | `SELECT` |
| `empresas` | `INSERT`, `UPDATE`, `DELETE` | `SELECT` |
| `tenants` | `INSERT`, `UPDATE`, `DELETE` | `SELECT` |
| `fornecedores` | `INSERT`, `UPDATE`, `DELETE` | `SELECT` |
| `lotes_carga` | `INSERT`, `UPDATE`, `DELETE` | `SELECT` |
| `lotes_carga_itens` | `INSERT`, `UPDATE`, `DELETE` | `SELECT` |
| `auditoria_log` | Qualquer acesso | `SELECT` |

### 9.2 Grants Mantidos (sem alteração)

| Tabela | Grants | Justificativa |
|:---|:---|:---|
| `tax_tokens` | `SELECT`, `INSERT`, `UPDATE`, `DELETE` | Tabela operacional do motor — gere tokens de congelamento de alíquotas |
| `fornecedor_fiscal` | `SELECT`, `INSERT`, `UPDATE` | Tabela operacional do motor — qualifica fornecedores para crédito |
| `iva_dual_rules_log` | `SELECT` | Histórico de alterações (leitura) |

### 9.3 Roteiro de Execução

| Passo | Responsável | Ação |
|:---|:---|:---|
| 1 | DBA | Executar script de revogação em ambiente de staging |
| 2 | DevOps | Atualizar secrets do DT-3 no K8S (remover credenciais de escrita) |
| 3 | DevSecOps | Validar que o DT-3 não consegue mais executar INSERT/UPDATE/DELETE nas tabelas administrativas |
| 4 | QA | Executar bateria de regressão no motor de cálculo — garantir que cálculos continuam funcionando com acesso somente leitura |
| 5 | DBA + DevSecOps | Repetir passos 1-3 em produção durante janela de manutenção |
| 6 | DevSecOps | Auditar logs de acesso PostgreSQL por 7 dias pós-revogação para detectar tentativas de escrita |
| 7 | DBA + DevSecOps | Atualizar documentação de acessos e conformidade |

---

## 10. Requisitos de Infraestrutura

| Componente | Plataforma | Escalabilidade | Observabilidade |
|:---|:---|:---|:---|
| **DT-1 Backend** | Kubernetes (cluster corporativo) | Horizontal Pod Autoscaler (HPA): min 2, max 6 pods. Métrica: CPU > 70% | Métricas Prometheus + Grafana. Tracing OpenTelemetry. Logs centralizados (Elasticsearch/Kibana) |
| **DT-2 Frontend** | CDN + storage estático (S3/equivalente) | Distribuição global via CDN corporativo | Google Analytics + Sentry (erros frontend) |
| **DT-3 Motor** | Kubernetes (existente — sem alteração) | Configuração atual mantida | Observabilidade existente mantida |
| **PostgreSQL** | Serviço gerenciado (RDS/equivalente) | Atual: instância existente. Monitorar crescimento com as novas tabelas | Métricas de conexão, slow queries, volume de dados |
| **Keycloak** | Serviço corporativo (existente) | Gerenciado pelo time de Infraestrutura | Logs de autenticação e autorização |
| **Redis** | Serviço gerenciado (ElastiCache/equivalente) | Configuração atual mantida | Métricas de hit/miss ratio |

---

## 11. Premissas e Restrições de Integração

1. **Banco de dados compartilhado:** DT-1 e DT-3 compartilham o mesmo schema `billing_tax_rates`. Não há replicação ou sincronização entre bases — a consistência é garantida pelo acesso ao mesmo PostgreSQL
2. **Isolamento de responsabilidades:** DT-1 é o único componente com permissão de escrita nas tabelas administrativas. DT-3 mantém escrita apenas nas tabelas operacionais (`tax_tokens`, `fornecedor_fiscal`)
3. **Segredos:** Nenhuma credencial de banco, API key ou token é armazenada em código-fonte, arquivos de configuração ou variáveis de ambiente. Todas as credenciais são obtidas via Secrets Manager (Vault) em runtime
4. **Rede:** A comunicação entre DT-1, DT-3 e PostgreSQL ocorre dentro da VPC corporativa. O portal (DT-2) é o único componente exposto publicamente (via CDN)
5. **Período de convivência:** Durante a Fase 3 (Integração), DT-1 e os endpoints admin do DT-3 coexistem. O cutover (desativação dos endpoints admin do DT-3) ocorre na Fase 4 (Go-Live), após validação de que todas as operações administrativas são realizadas exclusivamente via portal
6. **Carga em lote:** Arquivos de carga são armazenados temporariamente para processamento e removidos após 30 dias da aprovação. O conteúdo original de cada linha é preservado em `lotes_carga_itens.conteudo_original` (jsonb) para auditoria

---

🤖 *Documento gerado com apoio de Claude Code (Anthropic), em 11 de Julho de 2026.*
