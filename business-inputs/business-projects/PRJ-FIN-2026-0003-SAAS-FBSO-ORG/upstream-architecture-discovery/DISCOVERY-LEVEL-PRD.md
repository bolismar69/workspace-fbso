# DISCOVERY-LEVEL-PRD.md — PRD Discovery-Level
## Fase 1 — Bloco 0: Product Definition Discovery-Level

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Produto** | FBSO Platform — Portal Administrativo SaaS |
| **Documento** | DISCOVERY-LEVEL-PRD-v1.0 |
| **Versão** | 1.0 — Discovery-Level (Análise de Viabilidade) |
| **Data** | 02 de agosto de 2026 |
| **Tipo** | Documento de Negócio — briefing executivo do PM/PO/Analista de Negócios para o time de TI |
| **Status** | [STATUS: COMPLIANCE] — Aprovado em 02/08/2026 |

**Documentos Vinculados:**
- [`01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Project Charter
- [`02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Business Requirements Document
- [`03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md`](../03-EPICS-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) — Índice de Épicos
- [`epics/EP-0001-portal-administrativo-interno.md`](../epics/EP-0001-portal-administrativo-interno.md)
- [`epics/EP-0002-gestao-de-clientes-e-assinaturas.md`](../epics/EP-0002-gestao-de-clientes-e-assinaturas.md)
- [`epics/EP-0003-governanca-de-acessos-e-permissoes.md`](../epics/EP-0003-governanca-de-acessos-e-permissoes.md)
- [`epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md`](../epics/EP-0004-experiencia-do-cliente-e-autoatendimento.md)

---

## 1. Visão do Produto (Resumo Executivo para o Time de TI)

A **FBSO Platform** é uma suíte SaaS multi-produto que servirá como a fundação operacional e comercial da FBSO.ORG. Seu propósito é ser o ponto de entrada unificado para todos os produtos futuros da empresa — soluções fiscais (Tributali-Engine), varejo (Storekeeper Portal) e demais módulos — operando no modelo de módulos ativáveis por plano contratado.

O escopo **desta fase** é construir o **Portal Administrativo Core**: a camada fundamental que estrutura a operação comercial do SaaS. Isso significa que o time interno da FBSO.ORG terá ferramentas para gerenciar clientes, planos de assinatura e permissões de acesso, enquanto os clientes terão um portal de autoatendimento para onboarding, cadastro de suas filiais e gestão de seu portfólio de produtos.

**Por que isso importa agora:** Sem essa fundação, cada produto futuro (Tributali-Engine, Storekeeper Portal) precisaria construir sua própria camada administrativa, gerando retrabalho, custos duplicados e uma experiência fragmentada para clientes que contratem mais de um módulo. O Core resolve isso de uma vez, permitindo que novos produtos sejam acoplados como módulos ativáveis por plano, sem reestruturação.

**Visão de longo prazo:** Quando um cliente do Storekeeper Portal precisar se adequar ao Split Payment da Reforma Tributária, o upgrade será uma simples ativação de módulo na mesma plataforma — sem migração de dados, sem nova integração, sem troca de sistema.

---

## 2. Épicos × Soluções (Matriz Macro de Cobertura)

Cada épico de negócio demanda uma ou mais soluções para sua realização. A matriz abaixo estabelece o vínculo entre os épicos definidos pelo Negócio e as soluções previstas, servindo como mapa de rastreabilidade para o time de arquitetura.

| Épico | Objetivo de Negócio | Funcionalidades | Soluções Demandadas | Prioridade | Data-Alvo |
|-------|---------------------|-----------------|---------------------|------------|-----------|
| **EP-0001** — Portal Administrativo Interno | Time FBSO.ORG gerencia a operação SaaS com visibilidade em tempo real | 3 (Dashboard, métricas, visão consolidada) | Portal Administrativo Interno | Must Have | 15/08/2026 |
| **EP-0002** — Gestão de Clientes e Assinaturas | Estruturar a operação comercial: contas, planos e ciclo de vida do cliente | 5 (Tenants, planos, assinaturas, ativação, suspensão) | Portal Administrativo Interno | Must Have | 31/08/2026 |
| **EP-0003** — Governança de Acessos e Permissões | Garantir segurança e isolamento de dados entre clientes e entre filiais | 4 (RBAC, papéis, permissões, isolamento multi-tenant) | Portal Administrativo Interno + Portal do Cliente | Must Have | 15/09/2026 |
| **EP-0004** — Experiência do Cliente e Autoatendimento | Cliente realiza onboarding, gerencia seus dados e navega entre módulos | 6 (Portal, onboarding, App Switcher, Unidades de Negócio, Catálogo) | Portal do Cliente | Must Have | 15/10/2026 |

### Dependências entre Épicos

```
EP-0001 (Portal Admin Interno)
  └──▶ EP-0002 (Clientes e Assinaturas) — depende de tenants criados
         └──▶ EP-0003 (Governança e Permissões) — depende de tenants e usuários
                └──▶ EP-0004 (Portal do Cliente e Autoatendimento) — depende de toda camada de governança
```

EP-0001 é pré-requisito para EP-0002. EP-0003 depende da existência de Tenants e Usuários. EP-0004 depende de toda a camada de governança estabelecida. EP-0004 é entregue em duas etapas: Portal do Cliente + Onboarding (M5) e Unidades de Negócio + Catálogo (M6).

---

## 3. MVP Macro — O Que é Essencial para o Primeiro Delivery

O MVP Macro define o subconjunto mínimo de funcionalidades que entrega valor de negócio e permite validar a fundação da plataforma. Tudo o que está listado abaixo é **Must Have** para o primeiro ciclo de delivery.

### Bloco 1: Operação Interna Estruturada (Semanas 1-6)

| # | Funcionalidade | Épico | Por que é essencial |
|---|---|---|---|
| M1 | Dashboard administrativo com métricas da base de clientes | EP-0001 | Time interno precisa de visibilidade operacional desde o dia 1 |
| M2 | Cadastro, ativação, suspensão e reativação de contas de clientes (Tenants) | EP-0002 | Sem gestão de clientes não há operação SaaS |
| M3 | Cadastro de planos comerciais com definição de módulos incluídos | EP-0002 | Planos são a unidade básica da oferta comercial |
| M4 | Vinculação de clientes a planos (assinaturas) com data de vigência | EP-0002 | Conecta o cliente à sua oferta contratada |
| M5 | Registro de auditoria para ações administrativas | EP-0002 | Requisito de compliance e governança |

### Bloco 2: Segurança e Isolamento (Semanas 7-10)

| # | Funcionalidade | Épico | Por que é essencial |
|---|---|---|---|
| M6 | Cadastro de usuários com papéis (Admin, Gerente, Operador) | EP-0003 | Sem permissões, qualquer usuário acessa qualquer dado |
| M7 | Vinculação de usuários a Unidades de Negócio específicas | EP-0003 | Isolamento de dados entre filiais é inegociável |
| M8 | Controle de acesso: usuário vê apenas módulos e unidades autorizados | EP-0003 | Segurança multi-tenant e compliance |

### Bloco 3: Experiência do Cliente (Semanas 11-16)

| # | Funcionalidade | Épico | Por que é essencial |
|---|---|---|---|
| M9 | Portal do cliente com tela de boas-vindas pós-login | EP-0004 | Cliente precisa de um ponto de entrada profissional |
| M10 | Fluxo de onboarding guiado (confirmação de dados, cadastro da primeira filial) | EP-0004 | Reduz carga manual do time interno; meta: ≥80% autônomo |
| M11 | Navegação entre módulos (App Switcher / menu adaptado ao plano) | EP-0004 | Base para o modelo multi-produto; menus condicionais por plano |
| M12 | Cadastro de Unidades de Negócio com hierarquia Matriz/Filial | EP-0004 | Estrutura fundamental de dados comerciais do cliente |
| M13 | Cadastro de catálogo de produtos/serviços do cliente | EP-0004 | Prepara base de dados para futuros módulos fiscais e de varejo |

### Cronograma Macro do MVP

| Período | Marco | Entregas | Épicos |
|---------|-------|----------|--------|
| Semanas 1-3 | M1 | Kickoff + Portal Admin Interno | EP-0001 |
| Semanas 4-6 | M2-M5 | Gestão de Clientes, Planos e Assinaturas | EP-0002 |
| Semanas 7-10 | M6-M8 | Governança de Acessos e Permissões (RBAC) | EP-0003 |
| Semanas 11-14 | M9-M11 | Portal do Cliente, Onboarding, App Switcher | EP-0004 (parte 1) |
| Semanas 15-16 | M12-M13 | Unidades de Negócio e Catálogo de Produtos | EP-0004 (parte 2) |

---

## 4. Restrições Conhecidas (Constraints)

### 4.1 Fora do Escopo Desta Fase (Out of Scope)

As seguintes capacidades **não fazem parte** desta fase do projeto. Elas estão documentadas para garantir alinhamento de expectativas e evitar que o time de arquitetura dimensione esforço para itens que serão tratados futuramente.

| Item | Justificativa | Quando será tratado |
|------|---------------|---------------------|
| Módulo Tributali-Engine (cálculos fiscais IBS/CBS, Split Payment) | Fora do escopo do Core; é produto futuro | Fase pós-Core |
| Módulo Storekeeper Portal (PDV, estoque, varejo) | Fora do escopo do Core; é produto futuro | Fase pós-Core |
| Processamento real de cobranças e faturamento | Estrutura de dados preparada, sem execução financeira | Fase de comercialização |
| Integração com gateways de pagamento | Depende de contrato comercial e definição de parceiro | Fase de comercialização |
| Renovação automática de assinaturas | Será manual nesta fase | Fase pós-MVP |
| Período de trial gratuito automático | Pode ser simulado manualmente | Fase pós-MVP |
| Exportação de relatórios em PDF/Excel | Funcionalidade complementar | Fase pós-MVP |
| Dashboards customizáveis por usuário | MVP: dashboard único para todos os administradores | Fase pós-MVP |

### 4.2 Restrições de Negócio

| Restrição | Descrição | Impacto |
|-----------|-----------|---------|
| **R1 — Multi-Tenant Lógico** | Cada cliente (Tenant) tem seus dados completamente isolados. Nenhum cliente pode, sob qualquer hipótese, acessar dados de outro cliente | Arquitetura de dados e permissões deve garantir isolamento desde o primeiro dia |
| **R2 — RBAC Obrigatório** | Controle de acesso por papel é inegociável. MVP: 3 papéis (Admin, Gerente, Operador). Papel "Auditor" previsto para fase futura | Toda funcionalidade deve considerar permissões antes da implementação |
| **R3 — Auditoria Imutável** | Toda ação administrativa (ativação, suspensão, alteração de plano) deve ser registrada com identificação de quem fez, o que fez e quando fez | Exige trilha de auditoria desde o primeiro épico entregue |
| **R4 — Modelo de Módulos Ativáveis** | Produtos futuros são acoplados como módulos ativáveis por plano. A estrutura de planos deve suportar módulos associados (tabela `plan_modules`) | Flexibilidade na definição de planos; arquitetura extensível |
| **R5 — Onboarding Autônomo** | Meta de ≥80% dos clientes completarem onboarding sem ajuda humana | Fluxo de onboarding deve ser intuitivo e guiado |
| **R6 — Tempo de Bloqueio** | Suspensão de conta deve bloquear acesso do cliente em até 5 minutos | Requer mecanismo de propagação rápida de mudança de status |
| **R7 — Isolamento entre Filiais** | Usuário vinculado a Unidade de Negócio específica não pode acessar dados de outra filial do mesmo cliente | Multi-tenancy interno (intra-tenant); segunda camada de isolamento |

### 4.3 Restrições de Mercado e Regulatórias

| Restrição | Descrição |
|-----------|-----------|
| **LGPD** | Dados cadastrais de clientes e usuários são dados pessoais. Armazenamento, processamento e exclusão devem seguir a lei brasileira |
| **Reforma Tributária** | O cronograma de implementação da Reforma (IBS/CBS) pelo governo federal cria janela de oportunidade para o Tributali-Engine. O Core precisa estar pronto antes que a demanda fiscal se materialize |
| **Concorrência** | Mercado de plataformas SaaS para gestão fiscal e varejo possui players estabelecidos. Diferencial FBSO: suíte integrada multi-produto com experiência unificada |

---

## 5. Glossário Inicial — Termos Canônicos do Domínio

Este glossário estabelece a linguagem comum entre Negócio e TI. Todos os termos abaixo são extraídos dos documentos de negócio vinculados (Charter, BRD, Épicos) e devem ser usados consistentemente em todos os artefatos técnicos.

| Termo | Definição | Fonte |
|-------|-----------|-------|
| **Tenant** | Uma conta de cliente na plataforma SaaS. Representa uma empresa/organização que contratou a FBSO Platform. Cada tenant tem seus dados completamente isolados dos demais | Charter §3.1, BRD §5.1 |
| **Unidade de Negócio (BU)** | Uma filial, matriz ou estabelecimento vinculado a um Tenant. Possui CNPJ próprio e regime tributário. Estrutura hierárquica: Matriz → Filiais | Charter §3.1, EP-0004 |
| **Plano** | Oferta comercial que define quais módulos o cliente pode acessar e em quais condições (valor, recorrência). Ex: Básico, Core, Full Suite | Charter §3.1, EP-0002 |
| **Assinatura** | Vínculo entre um Tenant e um Plano, com data de início, vigência e status. Um cliente tem exatamente uma assinatura ativa por vez | EP-0002 |
| **Módulo** | Um produto ou funcionalidade ativável por plano. Ex: Tributali-Engine (futuro), Storekeeper Portal (futuro). No Core atual, os módulos são as áreas funcionais do portal | BRD §3, Charter §2 |
| **App Switcher** | Componente de navegação que permite ao cliente alternar entre os diferentes módulos/produtos que seu plano libera, sem sair da plataforma | Charter §3.1, EP-0004 |
| **RBAC (Role-Based Access Control)** | Modelo de controle de acesso baseado em papéis. Cada usuário recebe um papel (Admin, Gerente, Operador) que determina o que ele pode ver e fazer | EP-0003 |
| **Papel (Role)** | Função de um usuário no sistema que define seu nível de permissão. MVP: Administrador do Tenant, Gerente de Unidade, Operador | EP-0003 |
| **Onboarding** | Fluxo guiado de primeiro acesso do cliente, incluindo confirmação de dados cadastrais, criação da primeira Unidade de Negócio e orientação sobre os módulos disponíveis | EP-0004 |
| **Audit Trail (Trilha de Auditoria)** | Registro imutável de todas as ações administrativas realizadas na plataforma, contendo identificação do autor, ação realizada, data/hora e objeto afetado | EP-0002 |
| **Soft Delete** | Prática de marcar um registro como removido sem excluí-lo fisicamente do banco de dados, preservando o histórico e permitindo recuperação | BRD §5.1 |
| **FBSO Platform** | Nome do produto SaaS multi-produto da FBSO.ORG. A Plataforma é o ecossistema completo; o Portal Administrativo é sua interface de gestão | Charter §1 |
| **Core** | A camada fundamental da FBSO Platform — o Portal Administrativo com gestão de clientes, planos, permissões e autoatendimento. É o alicerce sobre o qual produtos futuros serão acoplados | Charter §2 |
| **Tributali-Engine** | Módulo futuro de solução fiscal (cálculos IBS/CBS, Split Payment, regras de retenção). Fora do escopo desta fase | Charter §3.2 |
| **Storekeeper Portal** | Módulo futuro de solução de varejo (PDV, estoque, gestão comercial). Fora do escopo desta fase | Charter §3.2 |
| **product_tax_mapping** | Estrutura de dados preparada para acoplamento futuro do Tributali-Engine. Schema definido, sem ativação de regras de tributação nesta fase. Contrato de interface documentado | Charter §3.1 |

---

## 6. Público-Alvo e Personas

| Persona | Descrição | Principais Necessidades | Épicos Relacionados |
|---------|-----------|------------------------|---------------------|
| **Administrador FBSO.ORG** | Colaborador do time interno que gerencia a operação SaaS | Visão rápida da base; ativar/suspender contas; auditar ações | EP-0001, EP-0002 |
| **Gestor de Produto FBSO.ORG** | Define e ajusta a oferta comercial de planos | Criar/alterar planos; definir módulos incluídos; versionar ofertas | EP-0002 |
| **Líder Comercial FBSO.ORG** | Responsável pela estratégia de vendas e relacionamento | Acompanhar adoção de planos; identificar upgrades; visão por status | EP-0001, EP-0002 |
| **Cliente Administrador** | Dono ou responsável pela conta da empresa no SaaS | Fazer onboarding rápido; cadastrar filiais; gerenciar catálogo; convidar usuários | EP-0004 |
| **Cliente Operador** | Funcionário que usa o portal no dia a dia | Acessar funcionalidades da sua unidade; realizar tarefas operacionais | EP-0003, EP-0004 |
| **Diretoria FBSO.ORG** | Sócios e diretores | Visão macro da operação; indicadores de crescimento; ROI | EP-0001 |

---

## 7. Métricas de Sucesso do Negócio

| Métrica | Situação Atual | Meta | Épico Relacionado |
|---------|---------------|------|-------------------|
| Tempo de ativação de cliente | ~2 dias úteis (manual) | ≤ 5 minutos | EP-0002 |
| Clientes que completam onboarding sozinhos | 0% | ≥ 80% | EP-0004 |
| Incidentes de vazamento de dados entre filiais | — | Zero incidentes | EP-0003 |
| Satisfação do time interno com ferramentas administrativas | — | Nota ≥ 4,0 / 5,0 | EP-0001 |
| Cobertura de auditoria das ações administrativas | 0% | 100% | EP-0002 |
| Tempo para suspender conta e bloquear acesso | — | ≤ 5 minutos | EP-0002 |

---

## Registro de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 02/08/2026 | Criação inicial: PRD Discovery-Level para análise de viabilidade. Cobre 4 épicos, 13 funcionalidades de MVP, 8 restrições de escopo, 7 restrições de negócio, 16 termos de glossário | Product Owner / Analista de Negócios |

---

🤖 *Documentação gerada como parte do Upstream Architecture Discovery — Fase 1 (Bloco 0). Este documento tem viés 100% negócio e é o briefing executivo para o time de TI iniciar a análise de viabilidade técnica.*
