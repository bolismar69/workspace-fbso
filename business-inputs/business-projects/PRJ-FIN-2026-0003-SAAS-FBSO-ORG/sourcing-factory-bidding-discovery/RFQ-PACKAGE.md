# RFQ-PACKAGE.md — Request for Quotation
## Sourcing & Factory Bidding — Fase 1 — Bloco A

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Produto** | FBSO Platform — Portal Administrativo SaaS (Core) |
| **Documento** | RFQ-PACKAGE-v1.0 |
| **Versão** | 1.0 — Discovery-Level |
| **Data** | 03 de agosto de 2026 |
| **Modo** | `discovery` (±50% ROM, baseado em Épicos) |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em 03/08/2026 |

**Artefatos Técnicos Vinculados:**
- [`DISCOVERY-LEVEL-PRD.md`](../upstream-architecture-discovery/DISCOVERY-LEVEL-PRD.md)
- [`DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md`](../upstream-architecture-discovery/DISCOVERY-LEVEL-ARCHITECTURE-DEFINITION.md)
- [`DISCOVERY-LEVEL-SECURITY-DEFINITION.md`](../upstream-architecture-discovery/DISCOVERY-LEVEL-SECURITY-DEFINITION.md)
- [`DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md`](../upstream-architecture-discovery/DISCOVERY-LEVEL-DATA-ARCHITECTURE-DEFINITION.md)
- [`DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md`](../upstream-architecture-discovery/DISCOVERY-LEVEL-DEVOPS-SRE-DEFINITION.md)
- [`DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md`](../upstream-architecture-discovery/DISCOVERY-LEVEL-TEST-STRATEGY-DEFINITION.md)
- [`DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md`](../upstream-architecture-discovery/DISCOVERY-LEVEL-INFRA-CLOUD-DEFINITION.md)
- [`DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md`](../upstream-architecture-discovery/DISCOVERY-LEVEL-SOLUTIONS-CATALOG.md)
- [`DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md`](../upstream-architecture-discovery/DISCOVERY-LEVEL-SOLUTIONS-MATRIX.md)
- [`DISCOVERY-LEVEL-SPECS.md`](../upstream-architecture-discovery/DISCOVERY-LEVEL-SPECS.md)
- [`DISCOVERY-LEVEL-ROM-ESTIMATE.md`](../upstream-architecture-discovery/DISCOVERY-LEVEL-ROM-ESTIMATE.md)

---

## 1. Visão do Projeto

A **FBSO Platform** é uma suíte SaaS multi-produto. Esta RFQ cobre a construção do **Portal Administrativo Core** — a fundação operacional que permitirá à FBSO.ORG gerenciar clientes, planos de assinatura, permissões de acesso e oferecer autoatendimento aos clientes.

O Core é um **monólito modular** (backend Java/Spring Boot + frontend React/Next.js) sobre infraestrutura **DigitalOcean DOKS** com **Kong↔Keycloak** como trust boundary, **PostgreSQL RLS** multi-tenant e stack completa de observabilidade.

**Objetivo desta RFQ:** Solicitar estimativas de esforço (horas), prazo (semanas), time (pessoas) e valor para implementação do escopo abaixo, conforme template padronizado.

---

## 2. Escopo — Épicos e Funcionalidades

### 2.1 Épicos para Estimativa

| ID | Épico | Objetivo de Negócio | Funcionalidades |
|----|-------|---------------------|:---------------:|
| **EP-0001** | Portal Administrativo Interno | Time FBSO.ORG gerencia a operação SaaS com visibilidade em tempo real | 3 |
| **EP-0002** | Gestão de Clientes e Assinaturas | Estruturar a operação comercial: contas, planos e ciclo de vida do cliente | 5 |
| **EP-0003** | Governança de Acessos e Permissões | Garantir segurança e isolamento de dados entre clientes e entre filiais | 4 |
| **EP-0004** | Experiência do Cliente e Autoatendimento | Cliente realiza onboarding, gerencia seus dados e navega entre módulos | 6 |

### 2.2 Detalhamento por Épico

#### EP-0001 — Portal Administrativo Interno (3 funcionalidades)

| Cód. | Funcionalidade | Descrição |
|:-----|---------------|-----------|
| F01-01 | Dashboard com métricas operacionais | Contas ativas, novas contas, distribuição por plano, status de tenants, filtro por período |
| F01-02 | Visão consolidada da base | Lista de clientes com busca por nome, status e plano; indicadores visuais de alerta |
| F01-03 | Gráfico de evolução | Evolução da base de clientes ao longo do tempo; estrutura preparada para métricas financeiras |

#### EP-0002 — Gestão de Clientes e Assinaturas (5 funcionalidades)

| Cód. | Funcionalidade | Descrição |
|:-----|---------------|-----------|
| F02-01 | Cadastro de Tenants | Criação de conta com dados corporativos (razão social, nome fantasia, segmento) |
| F02-02 | Ativação, suspensão e reativação | Ciclo de vida do tenant com registro de motivo e data |
| F02-03 | Cadastro de Planos Comerciais | Nome, descrição, valor, recorrência, módulos incluídos (tabela `plan_modules`) |
| F02-04 | Vinculação de Assinaturas | Tenant ↔ Plano com data de início, vigência e status |
| F02-05 | Histórico de Auditoria | Registro imutável de todas as ações administrativas (quem, o quê, quando) |

#### EP-0003 — Governança de Acessos e Permissões (4 funcionalidades)

| Cód. | Funcionalidade | Descrição |
|:-----|---------------|-----------|
| F03-01 | Cadastro de Usuários | Convite por e-mail, definição de senha no primeiro acesso |
| F03-02 | Definição de Papéis (RBAC) | MVP: 3 papéis (Admin Tenant, Gerente BU, Operador). Papel "Auditor" futuro |
| F03-03 | Vinculação a Unidades de Negócio | Usuário acessa apenas BUs autorizadas; escopo de permissão por módulo |
| F03-04 | Controle de Acesso Granular | Usuário vê apenas módulos e unidades autorizados; menus adaptados ao plano |

#### EP-0004 — Experiência do Cliente e Autoatendimento (6 funcionalidades · 23 US)

| Cód. | Funcionalidade | Descrição |
|:-----|---------------|-----------|
| F04-01 | Portal do Cliente | Tela de boas-vindas pós-login; área de perfil e configurações |
| F04-02 | Fluxo de Onboarding Guiado | Confirmação de dados, cadastro da primeira BU, orientação sobre módulos |
| F04-03 | App Switcher | Navegação entre módulos; menus adaptados ao plano contratado |
| F04-04 | Cadastro de Unidades de Negócio | CNPJ, estrutura hierárquica Matriz/Filial, regime tributário |
| F04-05 | Catálogo de Produtos/Serviços | Cadastro do portfólio comercial do cliente com classificação por tipo |
| F04-06 | Autenticação e Perfil do Cliente | Login OIDC, recuperação de senha, gestão de dados cadastrais |

### 2.3 Dependências entre Épicos

```
EP-0001 (Portal Admin Interno)
  └──▶ EP-0002 (Clientes e Assinaturas)
         └──▶ EP-0003 (Governança e Permissões)
                └──▶ EP-0004 (Portal do Cliente e Autoatendimento)
```

### 2.4 Fora do Escopo

- Módulo Tributali-Engine (cálculos fiscais IBS/CBS, Split Payment)
- Módulo Storekeeper Portal (PDV, estoque, varejo)
- Processamento real de cobranças e faturamento
- Integração com gateways de pagamento
- Renovação automática de assinaturas
- Exportação de relatórios em PDF/Excel

---

## 3. Contexto Técnico

### 3.1 Stack Tecnológica (Padrões Corporativos FBSO)

| Camada | Tecnologia |
|--------|-----------|
| **Backend** | Java 21 LTS + Spring Boot 3.x + GraalVM Native Image AOT |
| **Frontend** | React + Next.js + Tailwind CSS |
| **Banco de Dados** | PostgreSQL 17 (DO Managed) — Multi-Tenant via RLS + Soft Delete |
| **Cache** | Redis (DO Managed) |
| **IAM** | Keycloak 26 (OIDC + SAML 2.0) |
| **API Gateway** | Kong (Service-ID/Token-ID via Keycloak) |
| **Edge/CDN/WAF** | Cloudflare |
| **Cloud** | DigitalOcean (DOKS — Kubernetes) |
| **Service Mesh** | Istio (mTLS STRICT) |
| **Autoscaling** | Keda (pods) + Karpenter (nodes) |
| **IaC** | Terraform + Ansible |
| **CI/CD** | GitHub Actions |
| **Observabilidade** | Prometheus + Loki + Jaeger + Grafana + Elastic Stack |
| **Mensageria (futuro)** | RabbitMQ |

### 3.2 Arquitetura Macro

- **Modelo:** Monólito Modular com 9 módulos internos e boundaries de domínio explícitas
- **API Gateway:** Kong como trust boundary exclusiva — nenhum serviço backend recebe tráfego direto
- **Multi-Tenant:** Isolamento lógico via PostgreSQL RLS (coluna `tenant_id` em todas as tabelas)
- **Auditoria:** Trilha imutável (apenas INSERT) para todas as ações administrativas
- **Comunicação:** REST/HTTPS (síncrono); RabbitMQ planejado para fase futura (assíncrono)

### 3.3 Soluções Técnicas

| Solução | Tipo | Status |
|---------|------|--------|
| **ms-fbso-platform-admin** | Backend (Java/Spring Boot) | Existente |
| **web-app-fbso-platform-portal** | Frontend (React/Next.js) | Novo |

---

## 4. Instruções para Estimativa

### 4.1 Template Padronizado

As fábricas devem preencher o template CSV `ESTIMATION-SCHEMA.csv` (fornecido pela FBSO.ORG na Fase 2), com as seguintes colunas:

| Coluna | Descrição | Formato Discovery | Obrig. |
|--------|-----------|-------------------|:------:|
| `fabrica` | Nome da fábrica | Texto livre | ✅ |
| `id_epico` | Código do épico | `EP-0001` a `EP-0004` | ✅ |
| `titulo` | Título do épico | Conforme RFQ §2.2 | ✅ |
| `horas_dev` | Horas de desenvolvimento | Número inteiro | ✅ |
| `horas_qa` | Horas de QA/testes | Número inteiro | ✅ |
| `horas_arch` | Horas de arquitetura/SRE | Número inteiro | ✅ |
| `horas_devops` | Horas de DevOps/infra | Número inteiro | ✅ |
| `horas_gestao` | Horas de gestão/PM | Número inteiro | ✅ |
| `total_horas` | Soma das 5 colunas de horas | `=horas_dev+horas_qa+horas_arch+horas_devops+horas_gestao` | ✅ |
| `prazo_entrega_meses` | Prazo de entrega do épico em meses | Número decimal (ex: `2`, `3`) | ✅ |
| `time_estimado_pessoas` | Tamanho do time alocado para o épico | Número inteiro (ex: `13`) | ✅ |
| `valor_estimado` | Valor total do épico em R$ | Número (ex: `400000`) | ✅ |
| `complexidade` | Complexidade do épico | `Media` ou `Alta` | ✅ |
| `stack_aderencia` | Tecnologias da stack utilizadas no épico | Lista separada por `/ ` (ex: `DigitalOcean/ PostgreSQL/ Keycloak/ Kong`) | ✅ |
| `premissas` | Premissas do time, risco e pricing | Texto descritivo (ex: "Time de 13 pessoas (5 backend + 3 FE + 3 QA + 1 DevOps + 1 TL). Premium de 40%...") | ✅ |
| `comentarios` | Racional técnico detalhado da estimativa | Texto descritivo com justificativas por épico | ✅ |

> **Nota:** O modo Discovery utiliza 16 colunas. As colunas `features_codigos`, `qtd_features`, `user_stories_codigos` e `qtd_user_stories` são exclusivas do modo Full e **não** devem ser preenchidas neste template.

### 4.2 Regras de Qualidade

A FBSO.ORG validará as estimativas recebidas contra os seguintes critérios:

| Critério | Regra | Limite |
|----------|-------|--------|
| **QA Balanceado** | QA deve ser ≥ 25% do total de horas (global) e ≥ 20% por épico | QA < 10% → ❌ REJEITADA |
| **Arquitetura/SRE** | Arch deve ser ≥ 5% do total | Arch < 2% → ❌ REJEITADA |
| **Consistência Prazo×Horas** | Divergência entre prazo declarado e calculado | > 50% → ❌ REJEITADA |
| **Outliers** | Total de horas fora da faixa cross-fábrica | > 2× ou < 0.5× da mediana → ⚠️ Alerta |
| **PIB (Baseline Interna)** | Comparação com estimativa de referência FBSO.ORG | Usado na matriz comparativa |

### 4.3 Premissas para Estimativa

| # | Premissa |
|---|----------|
| P1 | Time sênior com 100% de dedicação |
| P2 | Stack 100% aderente aos padrões corporativos FBSO |
| P3 | 3 ambientes: Dev, Staging, Prod |
| P4 | Infraestrutura: DigitalOcean DOKS (São Paulo) + Cloudflare |
| P5 | Escopo limitado às 18 funcionalidades listadas (4 épicos, 62 US) |
| P6 | Sem dependências externas bloqueantes |

---

## 5. Critérios de Avaliação das Propostas

| Critério | Peso | Descrição |
|----------|:----:|-----------|
| **Custo Total** | 25-30% | Valor estimado (R$) |
| **Prazo** | 20-25% | Prazo de entrega (meses) |
| **Qualidade Técnica** | 20% | QA ≥ 25%, Arch ≥ 5% |
| **PIB — Proximidade à Baseline** | 15% | Comparação com estimativa de referência FBSO.ORG |
| **Consistência Prazo×Horas** | 15% | Divergência ≤ 30% |

---

## 6. Cronograma do Processo

| Marco | Descrição | Data Prevista |
|-------|-----------|:-------------:|
| Envio da RFQ | Distribuição do pacote às fábricas | A definir |
| Prazo para Q&A | Período de perguntas e esclarecimentos | A definir |
| Entrega das estimativas | Deadline para envio do CSV preenchido | A definir |
| Validação DTA | Análise de conformidade das estimativas | A definir |
| Seleção | Comunicação da fábrica selecionada | A definir |

---

## 7. Instruções de Submissão

1. Preencher o template `ESTIMATION-SCHEMA.csv` conforme Seção 4
2. Nomear o arquivo como `ESTIMATION-SCHEMA-{NOME-DA-FABRICA}.csv`
3. Enviar para o canal indicado pela FBSO.ORG
4. Prazo de entrega conforme cronograma (Seção 6)

**Contato para dúvidas:** Time de Arquitetura FBSO.ORG

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 03/08/2026 | Criação inicial: RFQ Package discovery-level. 4 épicos, 17 funcionalidades, stack corporativa, instruções de estimativa, 5 critérios de avaliação | PMO / Tech Lead |

---

🤖 *Sourcing & Factory Bidding — Fase 1. Pacote RFQ para distribuição às fábricas de software.*
