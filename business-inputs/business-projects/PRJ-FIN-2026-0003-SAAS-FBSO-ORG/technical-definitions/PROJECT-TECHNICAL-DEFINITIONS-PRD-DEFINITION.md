# PROJECT-TECHNICAL-DEFINITIONS-PRD-DEFINITION — Baseline de Requisitos de Produto

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Programa:** FBSO Platform — Portal Administrativo SaaS
- **Versão:** 1.0
- **Data de Criação:** 26 de Julho de 2026
- **Última Atualização:** 26 de Julho de 2026
- **Status:** CREATED — Aguardando validação humana (Gate → COMPLIANCE)
- **Documentos Complementares:** [SOLUTIONS-CATALOG](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md) · [STACK-MATRIX](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md) · [TEAM-MAP](./PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md)

---

## 1. Visão do Produto

### 1.1 Propósito

A **FBSO Platform** é o portal administrativo SaaS multi-produto da FBSO.ORG — a camada fundamental ("Core") sobre a qual todos os produtos futuros (Tributali-Engine fiscal, Storekeeper Portal varejo) serão acoplados como módulos ativáveis por plano.

### 1.2 Declaração de Produto

> Para **a FBSO.ORG**, que precisa de uma plataforma unificada para gerenciar clientes, planos, permissões e acesso aos seus produtos SaaS, a **FBSO Platform** é um **portal administrativo multi-tenant com RBAC** que permite ativar, suspender e gerenciar contas de clientes, configurar planos comerciais com módulos ativáveis e controlar acesso granular por usuário e unidade de negócio. Diferentemente de construir cada produto com sua própria estrutura administrativa, a FBSO Platform oferece **experiência unificada para o cliente** — quando um cliente do Storekeeper Portal precisar do Split Payment da Reforma Tributária, o upgrade será uma simples ativação de módulo na mesma plataforma.

### 1.3 Personas

| Persona | Descrição | Área de Acesso |
|:---|:---|:---|
| **Administrador FBSO** | Time interno que gerencia a operação do SaaS | Portal Admin Interno |
| **Líder Comercial** | Acompanha crescimento, métricas e tendências | Portal Admin Interno |
| **Diretoria** | Visão estratégica da evolução da base | Portal Admin Interno |
| **Administrador do Tenant** | Cliente que gerencia sua própria conta, usuários e unidades de negócio | Portal do Cliente |
| **Gerente de Unidade** | Cliente que gerencia uma unidade de negócio específica | Portal do Cliente |
| **Operador** | Usuário operacional do cliente com permissões restritas | Portal do Cliente |

### 1.4 Proposta de Valor por Persona

| Persona | Dor Atual | Proposta de Valor |
|:---|:---|:---|
| Administrador FBSO | Sem visão centralizada dos clientes e planos | Dashboard com métricas em tempo real, busca e filtros |
| Líder Comercial | Sem dados de crescimento e conversão | Gráficos de evolução, filtros por período |
| Administrador do Tenant | Sem autonomia para gerenciar seus usuários | Portal self-service, onboarding guiado, RBAC |
| Operador | Acesso confuso a múltiplos sistemas | App Switcher unificado, menus por permissão |

---

## 2. Matriz Requisito → Solução

### 2.1 Épicos × Soluções

| Épico | Nº Features | Soluções Primárias | Soluções de Suporte |
|:---|:---:|:---|:---|
| **EP-01** Portal Admin Interno | 3 | S01 (Backend), S02 (Frontend) | S03, S05, S06, S09, S11, S14 |
| **EP-02** Clientes e Assinaturas | 5 | S01 (Backend), S02 (Frontend) | S03, S05, S06, S11, S14 |
| **EP-03** RBAC | 4 | S01 (Backend), S02 (Frontend), S04 (Keycloak) | S03, S05, S06, S07, S11, S14 |
| **EP-04** Portal do Cliente | 6 | S01 (Backend), S02 (Frontend), S04 (Keycloak) | S03, S05, S06, S07, S11, S13, S14 |

### 2.2 Features × Soluções (Matriz Completa)

| ID | Feature | Prioridade | S01 BE | S02 FE | S04 IAM | Data-Alvo |
|:---|:---|:---:|:---:|:---:|:---:|:---|
| F01-01 | Dashboard de Métricas Operacionais | Must | ✅ API métricas | ✅ Gráficos | — | 15/08/2026 |
| F01-02 | Visão de Contas com Filtros | Must | ✅ API busca | ✅ Tabela+filtros | — | 15/08/2026 |
| F01-03 | Alertas e Indicadores de Atenção | Should | ✅ API alertas | ✅ Badges | — | 15/08/2026 |
| F02-01 | Cadastro e Ativação de Contas | Must | ✅ CRUD tenants | ✅ Formulários | — | 31/08/2026 |
| F02-02 | Gestão de Status do Tenant | Must | ✅ Status workflow | ✅ Indicadores | — | 31/08/2026 |
| F02-03 | Configuração de Planos Comerciais | Must | ✅ CRUD plans | ✅ Configurador | — | 31/08/2026 |
| F02-04 | Vinculação e Gestão de Assinaturas | Must | ✅ Subscriptions | ✅ Vinculação | — | 31/08/2026 |
| F02-05 | Histórico de Auditoria Administrativa | Must | ✅ Audit tables | ✅ Timeline | — | 31/08/2026 |
| F03-01 | Cadastro e Convite de Usuários | Must | ✅ CRUD users | ✅ Convite | ✅ JWT+roles | 15/09/2026 |
| F03-02 | Definição de Papéis e Permissões (RBAC) | Must | ✅ RBAC engine | ✅ Matriz | ✅ Realms | 15/09/2026 |
| F03-03 | Vinculação Usuário × Unidade × Módulo | Must | ✅ Vínculos | ✅ Interface | ✅ Claims | 15/09/2026 |
| F03-04 | Controle de Visibilidade de Menus e Ações | Must | ✅ Permissions | ✅ Menus dinâmicos | ✅ Roles | 15/09/2026 |
| F04-01 | Autenticação e Recuperação de Senha | Must | ✅ Auth endpoints | ✅ Login OIDC | ✅ OIDC+PKCE | 30/09/2026 |
| F04-02 | Onboarding Guiado de Primeiro Acesso | Must | ✅ Onboarding API | ✅ Wizard 4 passos | ✅ Realm setup | 30/09/2026 |
| F04-03 | Dashboard do Cliente | Should | ✅ Métricas tenant | ✅ Dashboard cliente | — | 30/09/2026 |
| F04-04 | App Switcher (Seletor de Módulos) | Must | ✅ Módulos API | ✅ Switcher UI | ✅ Permissions | 30/09/2026 |
| F04-05 | Gestão de Unidades de Negócio | Must | ✅ CRUD BUs | ✅ Hierarquia | ✅ BU claims | 15/10/2026 |
| F04-06 | Catálogo de Produtos e Serviços | Must | ✅ CRUD catálogo | ✅ Cards/tabela | ✅ Módulo ativo | 15/10/2026 |

### 2.3 User Stories por Feature

| Feature | US# | Título Resumido |
|:---|:---|:---|
| F01-01 | US-001 | Dashboard com indicadores principais (contas ativas, status, planos) |
| F01-01 | US-002 | Filtros de período (7, 30, 90 dias, mês atual, ano atual) |
| F01-01 | US-003 | Gráfico de evolução da base de clientes |
| F01-02 | US-004 | Lista completa de contas com paginação (25 por página) |
| F01-02 | US-005 | Busca textual por nome/razão social (a partir de 3 caracteres) |
| F01-03 | US-006 | Destaque visual de tenants com pendências |
| F01-03 | US-007 | Indicadores de atenção (onboarding incompleto, atraso pagamento) |
| F02-01 | US-008 | Cadastro de nova conta com dados do cliente |
| F02-01 | US-009 | Ativação de conta com envio de credenciais |
| F02-01 | US-010 | Validação de CNPJ no cadastro de conta |
| F02-01 | US-011 | Criação automática da primeira Unidade de Negócio |
| F02-02 | US-012 | Alteração de status (ativo, suspenso, inativo) com registro de auditoria |
| F02-02 | US-013 | Visualização do histórico de status do tenant |
| F02-02 | US-014 | Notificação de suspensão para o tenant |
| F02-03 | US-015 | Cadastro de plano com nome, descrição, valor e módulos |
| F02-03 | US-016 | Definição de vigência e recorrência do plano |
| F02-03 | US-017 | Tabela associativa `plan_modules` (quais módulos cada plano libera) |
| F02-03 | US-018 | Visualização da matriz plano×módulo |
| F02-04 | US-019 | Vinculação de tenant a plano com data de início |
| F02-04 | US-020 | Troca de plano (upgrade/downgrade) com efeito imediato |
| F02-04 | US-021 | Histórico de assinaturas do tenant |
| F02-05 | US-022 | Registro automático de ações administrativas (quem, quando, qual ação) |
| F02-05 | US-023 | Consulta de auditoria por tenant, período e tipo de ação |
| F03-01 | US-024 | Convite de usuário por email com link de ativação |
| F03-01 | US-025 | Cadastro de usuário com perfil e Unidade de Negócio |
| F03-01 | US-026 | Lista de usuários com filtros por tenant e status |
| F03-02 | US-027 | Definição de papéis padrão (Admin Tenant, Gerente, Operador) |
| F03-02 | US-028 | Criação de papel customizado pelo Admin do Tenant |
| F03-02 | US-029 | Atribuição de permissões granulares por papel |
| F03-02 | US-030 | Visualização da matriz papel×permissão |
| F03-03 | US-031 | Vinculação usuário a uma ou mais Unidades de Negócio |
| F03-03 | US-032 | Ativação/desativação de módulos por usuário |
| F03-03 | US-033 | Troca de contexto entre Unidades de Negócio (menu dropdown) |
| F03-04 | US-034 | Menu lateral renderizado conforme permissões do usuário |
| F03-04 | US-035 | Ocultação de ações (botões, links) sem permissão |
| F03-04 | US-036 | Mensagem amigável ao acessar rota sem permissão |
| F04-01 | US-037 | Tela de login via Keycloak com personalização por tenant (logo, cores) |
| F04-01 | US-038 | Recuperação de senha via email |
| F04-01 | US-039 | Logout com redirect para tela de login |
| F04-02 | US-040 | Boas-vindas pós-primeiro login |
| F04-02 | US-041 | Wizard de 4 passos (perfil, unidade, catálogo, confirmação) |
| F04-02 | US-042 | Validação de CNPJ da Unidade de Negócio no wizard |
| F04-02 | US-043 | Configuração de preferências iniciais (fuso horário, idioma) |
| F04-02 | US-044 | Tour guiado pelos menus e funcionalidades |
| F04-03 | US-045 | Dashboard com métricas do tenant (unidades, usuários, módulos ativos) |
| F04-03 | US-046 | Atalhos para ações frequentes (nova unidade, convidar usuário) |
| F04-04 | US-047 | App Switcher no topo com módulos contratados |
| F04-04 | US-048 | Indicador visual do módulo ativo |
| F04-04 | US-049 | Preparação para expansão: estrutura suporta novos módulos |
| F04-05 | US-050 | Cadastro de Unidade de Negócio (CNPJ, regime tributário) |
| F04-05 | US-051 | Estrutura hierárquica Matriz/Filial |
| F04-05 | US-052 | Lista de Unidades com busca e filtros |
| F04-05 | US-053 | Visualização detalhada da Unidade com usuários vinculados |
| F04-05 | US-054 | Edição e desativação de Unidade |
| F04-06 | US-055 | Cadastro de produto/serviço (nome, tipo, classificação) |
| F04-06 | US-056 | Lista de produtos com busca e filtros |
| F04-06 | US-057 | Tabela `product_tax_mapping` com schema definido |
| F04-06 | US-058 | Contrato de interface para acoplamento futuro do Tributali-Engine |

---

## 3. MVP Global

### 3.1 Definição de MVP

O MVP da FBSO Platform é o conjunto mínimo de funcionalidades que permite à FBSO.ORG operar comercialmente o SaaS — cadastrar clientes, configurar planos e controlar acesso — sem os módulos de produto final (Tributali-Engine, Storekeeper Portal).

### 3.2 Escopo MVP (Must Have)

| Épico | Features MVP | User Stories | Data-Alvo |
|:---|:---|:---:|:---|
| EP-01 | F01-01, F01-02 | 5 US | 15/08/2026 (M2) |
| EP-02 | F02-01 a F02-05 | 14 US | 31/08/2026 (M3) |
| EP-03 | F03-01 a F03-04 | 13 US | 15/09/2026 (M4) |
| EP-04a | F04-01, F04-02, F04-04 | 11 US | 30/09/2026 (M5) |
| EP-04b | F04-05, F04-06 | 9 US | 15/10/2026 (M6) |
| **Total MVP** | **16 features** | **52 US** | **15/10/2026** |

### 3.3 Fora do MVP (Should Have / Pós-MVP)

| Feature | Prioridade | Justificativa |
|:---|:---|:---|
| F01-03 (Alertas) | Should | Valor percebido é menor que Dashboard + Lista. Pode ser entregue no M2 se houver capacidade. |
| F04-03 (Dashboard Cliente) | Should | Cliente navega pelo menu. Dashboard é bônus se tempo permitir. |

### 3.4 Entregas (D1-D7) × Marcos (M1-M7)

| Entrega | Descrição | Marco | Data |
|:---|:---|:---:|:---|
| **D1** | Infraestrutura e setup inicial (Docker, DB, Keycloak, Kong, CI/CD) | M1 | 15/07/2026 ✅ |
| **D2** | Portal Admin — Dashboard e visão de contas (EP-01) | M2 | 15/08/2026 |
| **D3** | Gestão de clientes, planos e assinaturas (EP-02) | M3 | 31/08/2026 |
| **D4** | Gestão de usuários e permissões RBAC (EP-03) | M4 | 15/09/2026 |
| **D5** | Portal do Cliente — autenticação, onboarding, app switcher (EP-04a) | M5 | 30/09/2026 |
| **D6** | Unidades de negócio e catálogo (EP-04b) | M6 | 15/10/2026 |
| **D7** | Homologação final e ajustes pré-produção | M7 | 30/10/2026 |

---

## 4. Requisitos Funcionais Cross-Solution

### 4.1 Fluxos Multi-Solução

#### FCS-01: Onboarding de Novo Cliente (End-to-End)

```
1. [S02] Administrador FBSO acessa Portal Admin → cadastra tenant (US-008)
2. [S01] Backend cria tenant, gera credenciais, persiste no schema fbso_portal
3. [S04] Keycloak cria Realm para o tenant com configurações padrão
4. [S01] Backend dispara email de boas-vindas → [S07] MailHog captura (dev)
5. [S02] Administrador do Tenant recebe credenciais, acessa Portal do Cliente
6. [S04] Keycloak redireciona para tela de login OIDC com marca do tenant
7. [S14] Kong valida JWT, injeta X-Tenant-ID
8. [S01] Backend recebe headers, aplica SET app.current_tenant_id
9. [S02] Cliente passa pelo wizard de onboarding (US-040 a US-044)
10. [S01] Primeira Unidade de Negócio criada automaticamente
```

#### FCS-02: Controle de Acesso Granular (RBAC End-to-End)

```
1. [S02] Admin do Tenant define papel customizado (US-028)
2. [S01] Backend persiste papel e permissões no schema fbso_portal
3. [S04] Keycloak mapeia permissões como claims no JWT via Protocol Mappers
4. [S02] Admin vincula usuário ao papel e Unidade de Negócio (US-031)
5. [S14] Kong valida JWT, extrai claims → headers: X-User-Permissions
6. [S01] Backend aplica filtro de permissões nas queries
7. [S02] Frontend renderiza menu lateral conforme permissões (US-034)
8. [S02] Botões/ações sem permissão são ocultados (US-035)
```

#### FCS-03: White-Label por Domínio (Multi-Tenant Visual)

```
1. [S13] Cloudflare recebe requisição em cliente.com
2. [S13] Cloudflare aplica WAF, valida SSL, encaminha para DigitalOcean
3. [S02] Frontend extrai request.headers['host'] → "cliente.com"
4. [S01] API valida domínio → resolve tenant_id → retorna configuração do tenant
5. [S02] Frontend aplica tema customizado (logo, cores, fontes do tenant)
6. [S04] Se for tela de login, Keycloak já renderizou com a marca do tenant (Realm config)
```

### 4.2 Regras de Negócio Cross-Solution

| ID | Regra | Escopo | Soluções Afetadas |
|:---|:---|:---|:---|
| **RN-CS-01** | Todo tenant criado gera automaticamente um Realm no Keycloak | Global | S01, S04 |
| **RN-CS-02** | Toda query SQL DEVE passar pelo RLS — `SET app.current_tenant_id` antes da transação | Global | S01, S03 |
| **RN-CS-03** | Nenhum dado de tenant pode vazar para outro — isolamento validado em 100% dos testes | Global | S01, S03 |
| **RN-CS-04** | Soft Delete é obrigatório — `deleted_dt` preenchido, nunca DELETE físico | Global | S01, S03, S06 |
| **RN-CS-05** | Todo JWT emitido pelo Keycloak DEVE conter `tenant_id` e `permissions` como claims | Global | S04, S14 |
| **RN-CS-06** | Kong DEVE validar JWT antes de qualquer requisição atingir o backend | Global | S14, S01 |
| **RN-CS-07** | Logs NUNCA devem conter dados sensíveis (senhas, tokens, CPF) | Global | S01, S08 |
| **RN-CS-08** | CORS restrito aos domínios autorizados — nunca `Origin: *` em produção | Global | S01, S14 |

---

## 5. Requisitos Não-Funcionais Globais

### 5.1 Performance

| NFR | Métrica | Alvo | Solução |
|:---|:---|:---|:---|
| **NFR-P01** | Tempo de carregamento do dashboard | ≤ 3 segundos | S01, S02, S03 |
| **NFR-P02** | Latência de validação JWT (Kong) | ≤ 50ms | S14 |
| **NFR-P03** | Tempo de resposta API (p95) | ≤ 500ms | S01 |
| **NFR-P04** | Tempo de build + deploy (CI/CD) | ≤ 15 minutos | S11 |

### 5.2 Disponibilidade

| NFR | Métrica | Alvo |
|:---|:---|:---|
| **NFR-A01** | Disponibilidade do backend | 99.5% (exceto janelas de manutenção) |
| **NFR-A02** | Disponibilidade do Keycloak | 99.9% — autenticação é ponto crítico |
| **NFR-A03** | RTO (Recovery Time Objective) | ≤ 4 horas |
| **NFR-A04** | RPO (Recovery Point Objective) | ≤ 1 hora (backup PostgreSQL) |

### 5.3 Escalabilidade

| NFR | Métrica | Alvo |
|:---|:---|:---|
| **NFR-E01** | Tenants simultâneos suportados (MVP) | 50 |
| **NFR-E02** | Usuários por tenant | Até 500 |
| **NFR-E03** | Conexões simultâneas PostgreSQL | 100 (configurável) |

### 5.4 Segurança

| NFR | Requisito | Referência |
|:---|:---|:---|
| **NFR-S01** | Zero Hardcoded Secrets | GLOBAL-SECURITY.md §1.2 |
| **NFR-S02** | Menor Privilégio — endpoints privados por padrão | GLOBAL-SECURITY.md §1.1 |
| **NFR-S03** | Sanitização de inputs — XSS, SQL Injection | GLOBAL-SECURITY.md §1.3 |
| **NFR-S04** | Rate Limiting por endpoint | GLOBAL-SECURITY.md §2.4 |
| **NFR-S05** | Criptografia em repouso (senhas: bcrypt/argon2) | GLOBAL-SECURITY.md §2.2 |
| **NFR-S06** | RBAC granular — permissão por ação e módulo | S01, S04 |
| **NFR-S07** | Multi-Tenant Isolation — RLS com FORCE | S03 |

### 5.5 Observabilidade

| NFR | Requisito | Solução |
|:---|:---|:---|
| **NFR-O01** | Logs estruturados (JSON) com trace_id | S01, S08 |
| **NFR-O02** | Métricas de negócio exportadas (contas ativas, conversão) | S01, S09 |
| **NFR-O03** | Health checks em todos os serviços | S01, S03, S04, S14 |
| **NFR-O04** | Tracing distribuído em fluxos cross-solution | S08 |

---

## 6. Restrições de Produto

### 6.1 Fora do Escopo (Out of Scope)

| Área | O que NÃO será entregue | Justificativa |
|:---|:---|:---|
| **Tributali-Engine** | Cálculos IBS/CBS, Split Payment, NCM/NBS, guias | Projeto separado (PRJ-FIN-2026-0001) |
| **Storekeeper Portal** | PDV, estoque, vendas, pagamentos | Projeto separado futuro |
| **Faturamento Real** | Cobrança de clientes via plataforma | Apenas estrutura de plano cadastrada |
| **Integrações ERP** | Totvs, SAP, Omie | Fora do escopo do Core |
| **Gateways Pagamento** | Transações financeiras reais | Futuro |
| **Migração de Dados** | Importação de base de clientes legado | Fora do escopo |
| **Expurgo Definitivo** | DELETE físico de registros | Soft Delete apenas (ADR-004) |
| **RabbitMQ** | Mensageria entre módulos | Futuro (quando houver 2+ módulos) |

### 6.2 Premissas

| Premissa | Impacto se inválida |
|:---|:---|
| Time técnico disponível conforme TEAM-CAPACITY | Atraso nas entregas |
| Keycloak 26 compatível com OIDC Authorization Code + PKCE | Revisão de protocolo |
| PostgreSQL 17 suporta RLS com FORCE | Revisão de estratégia multi-tenant |
| GraalVM Native Image compatível com Spring Boot 3.5.14 | Fallback para JVM |
| Cloudflare Custom Hostnames disponível no plano contratado | Revisão de CDN |

### 6.3 Restrições Técnicas

| Restrição | Detalhe |
|:---|:---|
| **Banco único compartilhado** | Isolamento lógico via `tenant_id` + RLS. Sem schema-per-tenant. |
| **Monolítico Modular** | Backend único (não microserviços). Package-by-layer. |
| **Java 25 LTS** | Versão de runtime fixa. GraalVM Oracle. |
| **PostgreSQL 17** | Versão de banco fixa. |
| **Keycloak 26** | Versão de IAM fixa. |

---

## 7. Glossário de Domínio Unificado

| Termo | Definição Canônica | Sinônimos Proibidos |
|:---|:---|:---|
| **Tenant** | Conta de cliente na plataforma. Agrupa usuários, unidades de negócio e assinaturas. | Cliente, Account, Org |
| **Plano** | Configuração comercial que define módulos ativos, valor e recorrência. | Pacote, Tier, Membership |
| **Assinatura** | Vínculo entre um Tenant e um Plano com data de início e vigência. | Subscription, Contrato |
| **Unidade de Negócio** | CNPJ, filial ou empresa vinculada a um Tenant. Estrutura hierárquica Matriz/Filial. | Business Unit, Filial, Estabelecimento |
| **Papel (Role)** | Conjunto nomeado de permissões atribuído a um usuário. | Role, Perfil, Grupo |
| **Permissão** | Ação granular que um usuário pode executar em um módulo. | Permission, Right, Privilege |
| **Módulo** | Produto ou funcionalidade ativável por plano (ex: Tributali-Engine, Storekeeper). | App, Produto, Feature Flag |
| **Realm** | Domínio de isolamento no Keycloak. Um por tenant, com configurações de login customizadas. | — |
| **App Switcher** | Componente UI que permite alternar entre módulos ativos do tenant. | Seletor de Apps, Module Picker |
| **Onboarding** | Fluxo guiado de primeiro acesso do Administrador do Tenant. | Setup Wizard, Primeiro Acesso |
| **Soft Delete** | Marcação lógica de exclusão (`deleted_dt`). Registro permanece no banco. | Exclusão Lógica, Inativação |
| **RLS (Row-Level Security)** | Mecanismo PostgreSQL que força filtro `tenant_id` em todas as queries. | — |

---

## 8. Critérios de Aceitação Cross-Solution

### 8.1 Cenários End-to-End

#### CA-01: Isolamento Multi-Tenant

```
DADO QUE existem 2 tenants (Tenant A e Tenant B)
E o Tenant A tem 3 Unidades de Negócio e 10 usuários
QUANDO o Administrador do Tenant A acessa o Portal do Cliente
ENTÃO ele vê APENAS dados do Tenant A
E NUNCA vê dados do Tenant B
E o log de auditoria registra tenant_id = <tenant_a_uuid>
E nenhum erro de RLS é retornado
```

#### CA-02: RBAC Granular

```
DADO QUE o Tenant A tem 3 usuários:
  - Admin (acesso total)
  - Gerente Unidade X (acesso à Unidade X apenas)
  - Operador (acesso somente leitura ao catálogo)
QUANDO cada usuário acessa o Portal do Cliente
ENTÃO o Admin vê todos os menus e ações
E o Gerente vê apenas menus da Unidade X e ações de gestão
E o Operador vê apenas o menu Catálogo e nenhum botão de edição
E tentativas de acessar rotas sem permissão retornam 403
```

#### CA-03: Fluxo Completo de Onboarding

```
DADO QUE a FBSO.ORG cadastrou um novo tenant "Supermercado ABC"
QUANDO o Administrador do Supermercado ABC recebe credenciais
E acessa o Portal do Cliente pela primeira vez
ENTÃO ele é redirecionado ao Keycloak (tela com logo do Supermercado ABC)
E após login OIDC, o Kong valida o JWT
E o frontend exibe o wizard de onboarding (4 passos)
E ao concluir o wizard, a primeira Unidade de Negócio está criada
E o menu lateral reflete os módulos do plano contratado
E o App Switcher mostra apenas os módulos ativos
```

#### CA-04: White-Label por Domínio

```
DADO QUE o tenant "Supermercado ABC" configurou domínio "portal.superabc.com"
QUANDO um usuário acessa https://portal.superabc.com
ENTÃO a Cloudflare valida o domínio e encaminha via SSL
E o frontend identifica tenant_id pelo header 'host'
E a página de login do Keycloak exibe logo e cores do Supermercado ABC
E após login, o portal exibe a identidade visual do Supermercado ABC
```

---

## 9. Referências

| Documento | Relação |
|:---|:---|
| [01-PROJECT-CHARTER](../01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | Escopo, entregas D1-D7, marcos M1-M7 |
| [02-BRD](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | 10 requisitos de negócio, 8 NFRs |
| [03-EPICS](../03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | 4 épicos com jornadas de usuário |
| [04-FEATURES](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) | 18 features, 58 user stories |
| [05-MATRIZ-RTM](../05-MATRIZ-RASTREABILIDADE-RTM.md) | Rastreabilidade D→EP→F→US |
| [user-stories/](../user-stories/) | 18 arquivos de user stories |
| [TECHNICAL-PLAN.md](../TECHNICAL-PLAN.md) | Stack tecnológica e decisões |
| [GLOBAL-SECURITY.md](../../../.specs/security/GLOBAL-SECURITY.md) | Política de segurança |
| [SOLUTIONS-CATALOG](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-CATALOG.md) | 14 soluções técnicas |
| [STACK-MATRIX](./PROJECT-TECHNICAL-DEFINITIONS-SOLUTIONS-STACK-MATRIX.md) | Stacks por solução |
| [TEAM-MAP](./PROJECT-TECHNICAL-DEFINITIONS-TEAM-MAP.md) | Skills do time |

---

## Histórico de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 26/07/2026 | Criação inicial: PRD Definition com visão do produto, matriz requisito→solução (18 features, 58 US), MVP global (52 US), 3 fluxos cross-solution, 8 regras de negócio cross-solution, 18 NFRs, glossário unificado, 4 cenários end-to-end. | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Resultado da Fase 4 do Roadmap de Definições Técnicas — Pipeline: Generate → Gate → Fix → COMPLIANCE.*
