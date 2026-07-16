# DOCS-OTHERS-CAVEMAN-REVIEW.md — Revisão Cruzada de Documentos de Apoio

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Data da Revisão:** 15 de Julho de 2026
- **Revisor:** Agente Caveman Review + Analista de Negócios/Claude
- **Documentos Revisados:** README.md, GLOSSARY.md, ARCHITECTURE.md, TECHNICAL-PLAN.md, DEFINITION_OF_DONE.md, MATRIZ-KPI.md, STAKEHOLDER-MAP.md, TECHNICAL-TEAM-MAP.md, API-CONTRACTS.md, INTEGRATION-MAP.md
- **Método:** Estilo caveman — uma linha por achado: localização, problema, correção. Severidade: 🔴 bug | 🟡 risk | 🔵 nit | ❓ q

---

## Resumo

| Documento | Achados | 🔴 bugs | 🟡 risks | 🔵 nits | ❓ q |
|:---|:---:|:---:|:---:|:---:|:---:|
| README.md | 6 | 0 | 3 | 2 | 1 |
| GLOSSARY.md | 5 | 0 | 3 | 2 | 0 |
| ARCHITECTURE.md | 7 | 1 | 4 | 2 | 0 |
| TECHNICAL-PLAN.md | 6 | 0 | 4 | 2 | 0 |
| DEFINITION_OF_DONE.md | 4 | 0 | 2 | 1 | 1 |
| MATRIZ-KPI.md | 5 | 0 | 3 | 2 | 0 |
| STAKEHOLDER-MAP.md | 4 | 0 | 2 | 2 | 0 |
| TECHNICAL-TEAM-MAP.md | 5 | 0 | 3 | 2 | 0 |
| API-CONTRACTS.md | 6 | 0 | 4 | 2 | 0 |
| INTEGRATION-MAP.md | 6 | 1 | 3 | 2 | 0 |
| **Total** | **54** | **2** | **31** | **19** | **2** |

---

## 1. README.md

**Resumo:** Índice bem estruturado e completo. Principais problemas: omissão de documentos técnicos (API-CONTRACTS, INTEGRATION-MAP, ARCHITECTURE) no índice, inconsistência no cabeçalho com o próprio princípio de "Zero citações técnicas", e discrepância na contagem de regras de negócio.

**Achados:**

README.md:L203 🟡 risk: Convenção "Zero citações técnicas nos documentos de negócio (01 a 04)" mas o próprio README referencia TECHNICAL-TEAM-MAP.md e ferramentas como mermaid. Adicionar ressalva "exceto neste índice de navegação".

README.md:L141-142 🟡 risk: O índice (§3.5 e §3.6) omite API-CONTRACTS.md, INTEGRATION-MAP.md e ARCHITECTURE.md — documentos que já existem no diretório. Adicionar seção "3.7 Nível Técnico" com esses 3 documentos.

README.md:L178 🟡 risk: Contagem "18 Regras de Negócio (RN01-01 a RN18-04)" — mas FEATURES.md lista 18 features, não 18 RNs. Verificar se são 18 RNs ou se o número real é diferente. A numeração RN18-04 sugere feature 18, não regra 18.

README.md:L113 🔵 nit: Diz "7 entregas (D1-D7)" mas D2 e D3 são entregues no mesmo marco (M3), comportando-se como 6 ciclos de entrega. Esclarecer que são 7 entregas lógicas em 6 marcos de release.

README.md:L208 🔵 nit: MoSCoW "Should Have (2 features)" — quais? Listar F01-03 e F04-06 (as 2 confirmadas como Should Have) para rastreabilidade.

README.md:L24-25 ❓ q: O mermaid flowchart mostra 18 features no total, mas a seção §5 mapeia EP-04a (4 features) e EP-04b (2 features) = 6 features para EP-04. Confere com FEATURES.md?

---

## 2. GLOSSARY.md

**Resumo:** Boa cobertura terminológica. Falhas: typo em "Componento", referência quebrada ao Project Charter (código PC-* não usado nos arquivos reais), lista incompleta de KPIs e status de Tenant em português vs. enum técnico em inglês.

**Achados:**

GLOSSARY.md:L80 🔵 nit: `Componento central de navegação` → typo. Corrigir para `Componente`.

GLOSSARY.md:L7 🟡 risk: Referencia "PC-FBSO-PLATFORM-2026-001" como código do Project Charter. Nenhum arquivo no projeto usa esse código. Atualizar para "01-PROJECT-CHARTER-FBSO-PLATFORM.md".

GLOSSARY.md:L48 🟡 risk: Status do Tenant usa português ("Pendente Onboarding, Ativo, Suspenso, Inativo") enquanto TECHNICAL-PLAN.md usa inglês (PENDING_ONBOARDING, ACTIVE, SUSPENDED, INACTIVE). Definir idioma canônico para enums e aplicá-lo consistentemente em todos os docs.

GLOSSARY.md:L175-181 🟡 risk: Seção 7 lista apenas KPIs A1, A2, O1, O2, S1, S2. Omite A3 (Abandono de Onboarding), O3 (Vazamento entre BUs), O4 (Tempo de Bloqueio), S3 (Chamados de Suporte), P1, P2. Completar a lista ou remover a seção e referenciar MATRIZ-KPI.md.

GLOSSARY.md:L62 🔵 nit: Recorrências do Plano em português ("mensal, trimestral, anual") vs TECHNICAL-PLAN.md em inglês (MONTHLY, QUARTERLY, YEARLY). Mesmo problema de idioma dos enums.

---

## 3. ARCHITECTURE.md

**Resumo:** Visão arquitetural sólida com C4, ADRs e estrutura de pacotes. Problemas críticos: exemplo de SQL com concatenação de string (risco de injeção), JWT de exemplo contém módulos que não existem na Fase 0, divergência no número de ADRs vs TECHNICAL-PLAN.md.

**Achados:**

ARCHITECTURE.md:L337 🔴 bug: `"AND tenant_id = 't-12345'"` — concatenação de string em exemplo de SQL ensina padrão inseguro. Substituir por `"AND tenant_id = ?"` com placeholder parametrizado.

ARCHITECTURE.md:L164 🟡 risk: JWT de exemplo inclui `"modules": ["TRIBUTALI_ENGINE", "STOREKEEPER_PORTAL"]` — módulos que não existem na Fase 0. Adicionar `"FBSO_PLATFORM"` como módulo da Fase 0 ou deixar array vazio com comentário "// preenchido na Fase 0 apenas com FBSO_PLATFORM".

ARCHITECTURE.md:L183-192 🟡 risk: Lista 8 ADRs (ADR-01 a ADR-08). TECHNICAL-PLAN.md §2.3 lista apenas 6 ADRs. Harmonizar: remover duplicação ou fazer TECHNICAL-PLAN referenciar ARCHITECTURE como fonte canônica dos ADRs.

ARCHITECTURE.md:L56 ✅ CONFIRMADO: Stack do "Backend" lista "Java 25 LTS + Spring Boot + GraalVM" — versão correta (Oracle GraalVM 25.0.3+9.1). Documento complementado com GraalVM na coluna Stack.

ARCHITECTURE.md:L136 🔵 nit: JWT payload usa `business_unit_ids[]` como array separado mas TECHNICAL-PLAN.md ERD mostra USER_PERMISSION como tabela ponte com múltiplos registros. A claim `business_unit_ids[]` não aparece no mapeamento de claims do Keycloak em INTEGRATION-MAP.md L134. Alinhar representação.

ARCHITECTURE.md:L199-263 🔵 nit: Estrutura de pacotes usa `com.fbso.platform.admin/` mas o projeto no TECHNICAL-PLAN.md é nomeado `ms-fbso-platform-admin`. Padronizar: ou o pacote reflete o nome do serviço ou o serviço reflete o pacote.

ARCHITECTURE.md:L261 🟡 risk: `BaseEntity.java` como superclasse — Spring Data JDBC (ADR-02) não tem `@MappedSuperclass`. Se for JPA, ADR-02 está errado. Verificar e alinhar.

---

## 4. TECHNICAL-PLAN.md

**Resumo:** Documento técnico mais completo do projeto. Inclui stack, ERD, matriz de artefatos e sequenciamento. Stack confirmada: Java 25 LTS (Oracle GraalVM 25.0.3+9.1).

**Achados:**

TECHNICAL-PLAN.md:L70 ✅ CONFIRMADO: "Java 25 (versão LTS mais recente disponível)" — Java 25 é LTS (Setembro 2025). Versão exata: Oracle GraalVM 25.0.3+9.1 (build 25.0.3+9-LTS-jvmci-b01). Documento complementado com a versão exata.

TECHNICAL-PLAN.md:L49 🟡 risk: "Backend: Java 25 + Spring Boot (Microserviços)" — mas toda a arquitetura descreve um backend monolítico com package-by-layer, banco único compartilhado, sem service discovery, sem API gateway. Substituir "Microserviços" por "Aplicação Monolítica Modular" ou "Backend REST".

TECHNICAL-PLAN.md:L116-121 🟡 risk: Lista 6 ADRs (ADR-01 a ADR-06). ARCHITECTURE.md lista 8 ADRs (ADR-01 a ADR-08). ADR-07 (JWT Stateless) e ADR-08 (Docker+K8s) estão ausentes do TECHNICAL-PLAN. Adicionar os 2 ADRs faltantes ou remover do ARCHITECTURE.md.

TECHNICAL-PLAN.md:L73 🟡 risk: Menciona "GraalVM Native Image para inicialização rápida". Spring Boot 3.x/4.x + GraalVM Native Image tem restrições conhecidas (reflection, AOT compilation). Adicionar ressalva sobre complexidade ou mover para "futuro".

TECHNICAL-PLAN.md:L477-491 🔵 nit: Lista 11 grupos de endpoints mas §2.4.5 mapeia apenas 10 entidades para API. RESOURCE_ACTION e ROLE_RESOURCE são marcados como "sem CRUD externo" mas não têm endpoint correspondente listado. Consistência: ou lista os endpoints internos ou remove as entidades do mapeamento.

TECHNICAL-PLAN.md:L608 🔵 nit: 8 perguntas em aberto na §9. Nenhuma foi respondida desde 13/07/2026 (3 dias). Adicionar prazo para resposta ou status "em discussão" para cada pergunta.

---

## 5. DEFINITION_OF_DONE.md

**Resumo:** DoD bem estruturada com 3 níveis e critérios binários. Bom foco em resultado de negócio. Problemas: contagem de critérios errada, ausência de validação técnica, seção de IA auto-referencial.

**Achados:**

DEFINITION_OF_DONE.md:L136 🟡 risk: Cabeçalho diz "DoD de USER STORY (12 critérios)" mas a soma real é 11: F1-F3 (3) + V1-V3 (3) + A1-A2 (2) + U1-U3 (3) = 11. Corrigir para 11 ou adicionar 1 critério faltante (ex: critério de desempenho ou segurança).

DEFINITION_OF_DONE.md:L115-121 🟡 risk: Seção 7 "Uso com Agentes de IA" é auto-referencial (o documento dizendo como o agente deve lê-lo) e não adiciona valor de negócio. Mover para guia de onboarding do time ou documento de engenharia.

DEFINITION_OF_DONE.md:L99-110 🔵 nit: Seção 6 "O que NÃO faz parte da DoD" lista "Desempenho e capacidade" e "Disponibilidade 24/7" como fora do escopo, mas TECHNICAL-PLAN.md §8 lista riscos de performance e os NFRs do BRD exigem disponibilidade. Adicionar ressalva: "não são critério de DoD NO NÍVEL DE NEGÓCIO, mas são validados pelo time técnico".

❓ q: A DoD não menciona critérios técnicos como cobertura de testes, aprovação de code review, ou deploy em staging. Isso é intencional (visão pura de negócio) ou deve ser complementado por DoD técnico separado?

---

## 6. MATRIZ-KPI.md

**Resumo:** KPIs relevantes e bem vinculados aos critérios de sucesso do Charter. Problemas: conflito de métricas no S1 (NPS vs escala 1-5), dashboard executivo incompleto, referência quebrada a objetivo do BRD.

**Achados:**

MATRIZ-KPI.md:L65-72 🟡 risk: S1 define duas métricas conflitantes: "NPS ≥ 50" e "Nota média ≥ 4,0 em escala de 1-5". São métricas diferentes com escalas e métodos de cálculo distintos. Escolher uma: NPS ou Likert 1-5. Não usar as duas como equivalentes.

MATRIZ-KPI.md:L29 🟡 risk: A3 referencia "Objetivo de Negócio O2 do BRD" — mas o BRD (02-BUSINESS-REQUIREMENTS.md) não usa nomenclatura "O1/O2". Os requisitos são nomeados BR-A01 a BR-B05 e BR-NFR01 a BR-NFR08. Corrigir referência para o BR correto.

MATRIZ-KPI.md:L107-116 🟡 risk: Template de Dashboard Executivo exibe apenas 7 dos 10 KPIs. Ausentes: A2 (Tempo Médio de Onboarding), O2 (Cobertura de Auditoria), O4 (Tempo de Bloqueio), S3 (Chamados de Suporte), P2 (Tempo para Ativar Módulo). Adicionar linhas faltantes ou justificar exclusão (ex: "P2 medido apenas na fase seguinte").

MATRIZ-KPI.md:L107 🔵 nit: Coluna "Impacto no Negócio" no dashboard template tem textos descritivos mas não quantifica o impacto financeiro ou operacional. Ex: "Redução de custo de suporte" — qual custo? qual redução esperada?

MATRIZ-KPI.md:L83-86 🔵 nit: S3 "Chamados de Suporte por Cliente Novo" não especifica classificação dos chamados (N1/N2/N3). Chamado N1 de "esqueci a senha" tem peso diferente de N3 "bug crítico". Adicionar nota metodológica.

---

## 7. STAKEHOLDER-MAP.md

**Resumo:** Matriz RACI bem detalhada por entrega. Canais de comunicação e escalation path claros. Problema principal: documento está 100% não preenchido (template vazio), sem canais técnicos.

**Achados:**

STAKEHOLDER-MAP.md:L15-L47 🟡 risk: Todos os 8 stakeholders estão com `<nome>` e `<email>` como placeholders. Documento de 13/07/2026 ainda é um template, não um artefato populado. Preencher com nomes reais ou marcar status como "DRAFT — Aguardando Nomeações" consistente com L7.

STAKEHOLDER-MAP.md:L121-154 🟡 risk: Escalation path cobre apenas impedimentos de negócio, operacional e comercial. Não há caminho de escalação para incidentes técnicos (infra fora do ar, vazamento de dados, ataque de segurança). Adicionar "IMPEDIMENTO TÉCNICO" com caminho: DevOps → Arquiteto → Coordenador → Diretoria.

STAKEHOLDER-MAP.md:L111-115 🔵 nit: Canais de comunicação listam 5 fóruns, todos de negócio. Nenhum canal para sincronização técnica (daily do time de dev, revisão de arquitetura). Adicionar "Daily Técnica" e "Revisão de Arquitetura" como fóruns.

STAKEHOLDER-MAP.md:L45-L46 🔵 nit: Dois Early Adopters listados, mas PROJECT-CHARTER menciona apenas "1-2 clientes parceiros" como early adopters. Confirmar se são 2 clientes distintos ou 2 contatos do mesmo cliente.

---

## 8. TECHNICAL-TEAM-MAP.md

**Resumo:** Estrutura ideal de time bem dimensionada (10-12 pessoas). Bom alerta sobre time reduzido atual. Problemas: falta de skills específicas da stack, confusão entre "vagas" e "papéis", nenhum profissional designado.

**Achados:**

TECHNICAL-TEAM-MAP.md:L59 🟡 risk: Competências do Full-Stack listam apenas skills genéricas ("Desenvolvimento Web Full-Stack, APIs REST"). Não menciona as tecnologias específicas da stack: Java 25, Spring Boot, React, Next.js, Keycloak. Adicionar stack concreta às competências requeridas.

TECHNICAL-TEAM-MAP.md:L157 🟡 risk: Headcount "Desenvolvimento: 3" — mas §3 lista 3 papéis distintos (Full-Stack, Frontend, Backend). Se o time real tiver 1-2 pessoas, elas acumulam papéis? Adicionar coluna "Acúmulo de Papéis" para cenário de time reduzido.

TECHNICAL-TEAM-MAP.md:L63 🟡 risk: Desenvolvedor Frontend não menciona Tailwind CSS, Next.js App Router, ou MSW mock nas competências requeridas. Adicionar.

TECHNICAL-TEAM-MAP.md:L7 🔵 nit: Versão 1.1 mas cabeçalho diz "Data de Criação: 13 de Julho de 2026" sem campo "Última Atualização" separado. Adicionar campo de atualização como nos outros docs.

TECHNICAL-TEAM-MAP.md:L203 🔵 nit: Referência "risco R2 do Project Charter" — o Charter (01-PROJECT-CHARTER-FBSO-PLATFORM.md) lista riscos com IDs diferentes. Verificar o ID correto e atualizar a referência cruzada.

---

## 9. API-CONTRACTS.md

**Resumo:** Contratos de API bem definidos com 11 recursos, schemas de exemplo e matriz RBAC. Problemas: PUT para operações parciais (deveria ser PATCH), divergência de versionamento com INTEGRATION-MAP, onboarding com verbos HTTP inconsistentes.

**Achados:**

API-CONTRACTS.md:L286-L287 🟡 risk: `PUT /users/{userId}/permissions` e `PUT .../permissions/role` para operações de atualização parcial. PUT implica substituição completa do recurso; PATCH é o verbo correto para atualização parcial. Trocar por PATCH.

API-CONTRACTS.md:L481-L483 🟡 risk: Onboarding step-1 usa PATCH, step-2 usa POST, step-3 é GET (leitura). Inconsistência semântica: todos os passos que alteram estado deveriam usar o mesmo verbo. Padronizar: POST para ações de avanço de passo + GET para consulta de status.

API-CONTRACTS.md:L26 🔵 nit: Base URL `/api/v1` declarada como fixa. INTEGRATION-MAP.md L89 diz "a confirmar". Resolver divergência: ou confirmar `/api/v1` ou alinhar ambos os docs para "a definir".

API-CONTRACTS.md:L554-569 🟡 risk: Matriz RBAC — "Manager BU" pode "Criar, Editar, Ver" Products, mas §R-07 define POST /products como "Admin Tenant / Manager BU". A matriz e o detalhamento conferem, mas "Operator BU" e "Auditor" têm "Ver" para Business Units enquanto §R-06 diz "Manager BU (apenas suas)" — falta granularidade: Manager BU vê apenas suas BUs, mas a matriz mostra "Ver + Editar (sua)" — OK. Verificar se a matriz de 1 linha é suficiente para capturar a nuance "apenas suas".

API-CONTRACTS.md:L132 🟡 risk: INTEGRATION-MAP.md L132 lista os papéis do Keycloak como 4 roles (ADMIN_TENANT, MANAGER_BU, OPERATOR_BU, AUDITOR) mas a matriz RBAC do API-CONTRACTS inclui "Admin FBSO" — que não é um role do Keycloak. O "Admin FBSO" é o time interno usando o realm `master` ou uma role separada? Documentar.

API-CONTRACTS.md:L108 🔵 nit: Tenant endpoint `POST /tenants/{id}/resend-invite` — nome do endpoint em inglês vs ações irmãs `activate`, `suspend`, `reactivate`. Padronizar idioma das ações: todas em inglês ou todas em português.

---

## 10. INTEGRATION-MAP.md

**Resumo:** Diagrama de integrações claro com 8 integrações mapeadas, fluxos detalhados e Docker Compose funcional. Problemas: senha hardcoded, conflito de portas, role Admin FBSO ausente do Keycloak.

**Achados:**

INTEGRATION-MAP.md:L291 🔴 bug: `KEYCLOAK_ADMIN_PASSWORD: admin` — senha hardcoded em documentação publicada. Substituir por placeholder `KEYCLOAK_ADMIN_PASSWORD: ${KEYCLOAK_ADMIN_PASSWORD:-changeme}` com nota de que deve ser sobrescrito por variável de ambiente.

INTEGRATION-MAP.md:L305 🔴 bug: Keycloak mapeado na porta `8080` — Spring Boot também usa `8080` como porta padrão. Backend está na `8081`, mas em ambientes sem Docker Compose (IDE), o conflito aparece. Adicionar nota ou trocar Keycloak para `8443`.

INTEGRATION-MAP.md:L132 🟡 risk: Roles do Keycloak listam 4 perfis (ADMIN_TENANT, MANAGER_BU, OPERATOR_BU, AUDITOR). O papel "Admin FBSO" (time interno) não está mapeado. Ele autentica via realm `master` do Keycloak? Via role `ROLE_ADMIN` separada? Documentar.

INTEGRATION-MAP.md:L367-L373 🟡 risk: Próximos passos listam "Definir schema SQL completo" e "Especificar OpenAPI YAML" como pendentes, mas TECHNICAL-PLAN.md §2.4 já define o ERD e API-CONTRACTS.md já esboça os endpoints. Atualizar próximos passos para refletir o que já foi feito: "Refinar schema SQL (Flyway/Liquibase)" e "Gerar OpenAPI YAML a partir do API-CONTRACTS.md".

INTEGRATION-MAP.md:L169 🟡 risk: `Isolation level: READ_COMMITTED` — para operações financeiras (assinaturas, planos), READ_COMMITTED permite phantom reads e non-repeatable reads. Adicionar nota de que operações de alteração de assinatura devem usar `@Transactional(isolation = REPEATABLE_READ)` ou `SELECT ... FOR UPDATE`.

INTEGRATION-MAP.md:L285 🔵 nit: Docker Compose usa `postgres:17` e `keycloak:26` — versões sem tag patch (ex: `17.4`). Em produção, pinar versão exata para reprodutibilidade. Adicionar comentário "use tag exata em produção".

---

## 11. Problemas Transversais (Multi-Documento)

Estes achados afetam 2 ou mais documentos simultaneamente.

### 11.1 Inconsistência de Idiomas (Português vs Inglês)

🟡 risk: Enums e status alternam entre português e inglês em 5 documentos:

| Conceito | GLOSSARY.md | TECHNICAL-PLAN.md | API-CONTRACTS.md | ARCHITECTURE.md |
|:---|:---|:---|:---|:---|
| Status do Tenant | Ativo, Suspenso | ACTIVE, SUSPENDED | ACTIVE | — |
| Recorrência | mensal, trimestral | MONTHLY, QUARTERLY | MONTHLY, YEARLY | — |
| Roles | Gerente de Unidade | — | MANAGER_BU | MANAGER_BU |

**Decisão necessária:** Definir idioma canônico para valores técnicos (recomendação: inglês para enums em código, português para documentos de negócio). Documentar no GLOSSARY.md.

### 11.2 Divergência ADR Count (TECHNICAL-PLAN vs ARCHITECTURE)

🟡 risk: TECHNICAL-PLAN.md §2.3 lista 6 ADRs. ARCHITECTURE.md §4 lista 8 ADRs. ADR-07 (JWT Stateless) e ADR-08 (Docker+K8s) só existem no ARCHITECTURE.md. Solução: remover a lista de ADRs do TECHNICAL-PLAN, substituir por referência "Ver ARCHITECTURE.md §4 para decisões arquiteturais completas".

### 11.3 Java 25 LTS — Confirmado (NÃO é bug)

✅ CONFIRMADO: Java 25 é versão LTS, lançada em Setembro de 2025. Versão exata em uso: Oracle GraalVM 25.0.3+9.1 (build 25.0.3+9-LTS-jvmci-b01). A referência a "Java 25 LTS" em TECHNICAL-PLAN.md e ARCHITECTURE.md está CORRETA. Este falso positivo foi removido da lista de bugs na revisão final. Os documentos foram complementados com a versão exata do GraalVM.

### 11.4 Module Gap na Fase 0

🟡 risk: JWT payload (ARCHITECTURE.md:L164) e resposta de user (API-CONTRACTS.md:L270) referenciam `modules: ["TRIBUTALI_ENGINE", "STOREKEEPER_PORTAL"]`. A Fase 0 não entrega esses módulos. Adicionar `"FBSO_PLATFORM"` como módulo base e marcar os demais como `"// fases futuras"`.

### 11.5 Keycloak — Role "Admin FBSO" Não Mapeada

🟡 risk: API-CONTRACTS.md matriz RBAC e INTEGRATION-MAP.md mencionam "Admin FBSO" como perfil com acesso total. Mas INTEGRATION-MAP.md:L132 lista apenas 4 roles no Keycloak, nenhuma delas "Admin FBSO". Definir: é uma role no realm `fbso-platform`? É acesso via realm `master`? É controle puramente backend (fora do Keycloak)?

---

## 12. Recomendações Prioritárias

Ordem sugerida de correção (impacto × severidade):

| # | Ação | Docs Afetados | Severidade |
|:---|:---|:---|:---|
| 1 | Remover senha hardcoded do Docker Compose | INTEGRATION-MAP | 🔴 bug |
| 2 | Corrigir exemplo de SQL injection no ARCHITECTURE.md | ARCHITECTURE | 🔴 bug |
| 3 | Definir idioma canônico para enums (PT vs EN) | GLOSSARY, TECHNICAL-PLAN, API-CONTRACTS | 🟡 risk |
| 4 | Harmonizar ADRs (6 vs 8) entre TECHNICAL-PLAN e ARCHITECTURE | TECHNICAL-PLAN, ARCHITECTURE | 🟡 risk |
| 5 | Mapear role "Admin FBSO" no Keycloak | API-CONTRACTS, INTEGRATION-MAP | 🟡 risk |
| 6 | Remover módulos inexistentes do JWT de exemplo | ARCHITECTURE, API-CONTRACTS | 🟡 risk |
| 7 | Corrigir contagem de critérios DoD (11 vs 12) | DEFINITION_OF_DONE | 🟡 risk |
| 8 | Corrigir métrica conflitante S1 (NPS vs Likert) | MATRIZ-KPI | 🟡 risk |
| 9 | Preencher stakeholders com nomes reais | STAKEHOLDER-MAP | 🟡 risk |
| 10 | Corrigir "Microserviços" → "Monolítico Modular" no TECHNICAL-PLAN | TECHNICAL-PLAN | 🟡 risk |

---

## 13. Nota sobre Documentos Não Revisados

Os 4 documentos da cadeia principal de negócios NÃO foram revisados nesta análise (revisados anteriormente em DOCS-USER-STORY-REVIEW.md):

- `01-PROJECT-CHARTER-FBSO-PLATFORM.md`
- `02-BUSINESS-REQUIREMENTS.md`
- `03-EPICS.md`
- `04-FEATURES.md`

Os 18 arquivos de User Story (`05-USER-STORYS-*.md`) também foram revisados separadamente.

---

## 14. Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 15/07/2026 | Revisão inicial: 54 achados em 10 documentos, 12 problemas transversais, 10 recomendações prioritárias | Caveman Review + Analista de Negócios/Claude |
| 1.1 | 15/07/2026 | Correção pós-revisão: Java 25 confirmado como LTS (Oracle GraalVM 25.0.3+9.1) — falso positivo removido. Ajustes aplicados em todos os 10 documentos + DOCS-OTHERS-CAVEMAN-REVIEW.md | Caveman Review |

------------------------------

---
🤖 *Documento gerado pelo skill caveman-review. Estilo: terse, um achado por linha. Severidade: 🔴 bug | 🟡 risk | 🔵 nit | ❓ question.*
