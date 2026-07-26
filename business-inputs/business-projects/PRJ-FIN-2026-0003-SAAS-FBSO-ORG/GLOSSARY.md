# Glossário do Projeto

- **Projeto:** FBSO Platform — Portal Administrativo SaaS
- **Código:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Versão:** 1.0
- **Atualizado:** 2026-07-13
- **Objetivo:** Fonte única de verdade para a terminologia específica deste projeto. Este glossário complementa o Glossário de Termos de Negócio contido no [Project Charter](./01-PROJECT-CHARTER-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md) (Seção Glossário).

---

## Nota Importante

Este glossário contém os termos **específicos deste projeto**.

---

## 1. Conceitos da Plataforma

### FBSO Platform

Plataforma SaaS multi-produto da FBSO.ORG. Opera no modelo de Suíte com módulos ativáveis por plano contratado. O Core administrativo (escopo deste projeto) gerencia contas, planos, usuários e permissões para todos os módulos.

- **Escopo deste projeto:** Portal administrativo interno + Portal do cliente + Cadastros base (Unidades de Negócio e Catálogo de Produtos).
- **O que NÃO é:** Os módulos-produto (Tributali-Engine e Storekeeper Portal) que serão desenvolvidos em fases futuras.
- **Sinônimos no contexto do projeto:** "a plataforma", "o portal", "o SaaS", "FBSO Platform".

### Core Administrativo

A camada fundamental e compartilhada da plataforma — gerencia contas de clientes (Tenants), planos comerciais, assinaturas, usuários e permissões de acesso (RBAC). É o alicerce sobre o qual todos os módulos-produto futuros serão acoplados.

- **Objetivo:** Prover uma fundação única para gestão administrativa, eliminando a necessidade de cada módulo-produto construir sua própria camada de administração.
- **Sinônimos:** "Core", "fundação", "camada base".

### Modelo "Lego" (Multi-Product Suite)

Estratégia de produto onde o SaaS é uma plataforma única com módulos independentes que podem ser ativados ou desativados por cliente conforme o plano contratado. Um supermercado que contrata o Storekeeper Portal pode, no futuro, ativar o Tributali-Engine com um simples upgrade de plano — sem migração de dados, sem nova integração, sem troca de sistema.

---

## 2. Entidades de Negócio

### Tenant

Conta corporativa de um cliente na plataforma. Representa a empresa que contratou o SaaS da FBSO.ORG.

- **Atributos principais:** razão social, nome fantasia, segmento de mercado, status (PENDING_ONBOARDING / Pendente Onboarding, ACTIVE / Ativo, SUSPENDED / Suspenso, INACTIVE / Inativo).
- **Ciclo de vida:** Criado pelo time FBSO.ORG → Cliente realiza onboarding → Status muda para ACTIVE → Pode ser SUSPENDED ou INACTIVE.
- **Sinônimos:** "conta", "cliente", "empresa contratante".
- **Convenção:** Documentos de negócio usam termos em português. Documentos técnicos (TECHNICAL-PLAN.md, API-CONTRACTS.md, código) usam enums em inglês.

### Unidade de Negócio (Business Unit)

Um CNPJ ou filial vinculada a um Tenant. Cada unidade pode ter configurações próprias de regime tributário e catálogo de produtos. É a entidade central de isolamento de dados — um usuário com acesso a uma unidade não pode ver dados de outra.

- **Atributos principais:** CNPJ, razão social, regime tributário (Simples Nacional, Lucro Real, Lucro Presumido), endereço, vínculo hierárquico (Matriz/Filial).
- **Hierarquia:** Matriz (primeira unidade cadastrada no onboarding) → Filiais (níveis recursivos).
- **Sinônimos:** "BU", "filial", "CNPJ", "empresa".

### Plano (Plan)

Pacote comercial oferecido pela FBSO.ORG. Define quais módulos/produtos estão disponíveis, o valor e a recorrência de cobrança.

- **Atributos principais:** nome, descrição, valor, recorrências disponíveis (MONTHLY / mensal, QUARTERLY / trimestral, YEARLY / anual), lista de módulos incluídos, status (ACTIVE / Ativo, DISCONTINUED / Descontinuado).
- **Exemplos:** Plano Básico (apenas 1 módulo), Plano Core (módulos essenciais), Plano Full Suite (todos os módulos).
- **Sinônimos:** "pacote", "oferta comercial".

### Assinatura (Subscription)

Vínculo ativo entre um Tenant e um Plano, com data de início, vigência e status. A assinatura determina quais módulos o cliente pode acessar.

- **Atributos principais:** plano vinculado, data de início, data de término (opcional), status (Ativa, Suspensa, Cancelada).
- **Regra de negócio:** Um tenant pode ter apenas uma assinatura ativa por vez. Upgrade/downgrade encerra a assinatura anterior e cria uma nova.
- **Sinônimos:** "contratação", "vínculo comercial".

---

## 3. Módulos e Produtos

### App Switcher (Seletor de Aplicativos)

Componente central de navegação do portal, posicionado no topo da interface. Permite que o usuário alterne entre os diferentes módulos/produtos da FBSO Platform sem sair do portal.

- **Comportamento:** Exibe apenas os módulos que o usuário tem permissão para acessar (interseção entre módulos do plano e permissões do usuário). Ao trocar de módulo, o menu lateral e o conteúdo da tela se adaptam.
- **Fase 0 (este projeto):** App Switcher é visível com um módulo placeholder "FBSO Platform". A estrutura está preparada para receber os módulos Tributali-Engine e Storekeeper Portal quando forem desenvolvidos.

### Tributali-Engine

Módulo-produto futuro da plataforma, focado em gestão tributária da Reforma Tributária (IBS/CBS). Realizará cálculos fiscais, Split Payment e automação de obrigações acessórias.

- **Status:** Fora do escopo deste projeto. Será desenvolvido como módulo ativável sobre o Core.

### Storekeeper Portal

Módulo-produto futuro da plataforma, focado em varejo e gestão comercial. Incluirá PDV (frente de caixa), controle de estoque, painel de vendas e integrações comerciais para supermercados, farmácias e lojistas.

- **Status:** Fora do escopo deste projeto. Será desenvolvido como módulo ativável sobre o Core.

---

## 4. Perfis de Acesso e Governança

### Administrador do Tenant (Admin Tenant)

Perfil de acesso com privilégios máximos dentro de um Tenant. Pode acessar todas as Unidades de Negócio, gerenciar usuários e permissões, e visualizar o plano contratado.

- **Atribuído a:** Dono da conta ou contador master.
- **Restrição:** Não pode acessar dados de outros Tenants.

### Gerente de Unidade (Manager BU)

Perfil de acesso com privilégios de gestão dentro de Unidades de Negócio específicas. Pode gerenciar produtos, catálogo e operações, mas apenas nas unidades vinculadas ao seu perfil.

- **Restrição:** Não pode alterar configurações de outros usuários nem acessar unidades não vinculadas.

### Operador de Unidade (Operator BU)

Perfil de acesso operacional. Executa tarefas do dia a dia (cadastros, consultas) nas unidades vinculadas, sem acesso a configurações sensíveis ou gestão de usuários.

- **Restrição:** Não pode criar ou editar regras de negócio, nem gerenciar permissões.

### Auditor

Perfil de acesso exclusivamente consultivo. Pode visualizar dados das unidades autorizadas, incluindo a trilha de auditoria completa, mas não pode criar, editar ou excluir nada.

- **Atribuído a:** Auditores internos ou externos, controllers.

### RBAC (Role-Based Access Control)

Modelo de controle de acesso baseado em papéis. Cada usuário recebe um papel (Admin, Gerente, Operador, Auditor) que determina automaticamente o que ele pode ver e fazer na plataforma. As permissões são definidas por papel, não por usuário individual.

- **Matriz de permissões:** Documentada em FEATURES.md, RN10-01.
- **Sinônimos:** "controle de acesso", "perfis de acesso".

### Segregação de Funções

Princípio de controle interno pelo qual um Operador de Unidade não pode ter acesso a configurações fiscais que são de responsabilidade do Administrador do Tenant, e um Auditor não pode executar ações que audita.

---

## 5. Fluxos de Negócio

### Onboarding

Fluxo guiado de primeiro acesso do cliente ao portal. Conduz o novo cliente por 4 passos: (1) confirmação de dados cadastrais, (2) cadastro da primeira Unidade de Negócio (Matriz), (3) orientação sobre o plano e módulos disponíveis, (4) boas-vindas e liberação do portal.

- **Duração esperada:** ≤ 10 minutos.
- **Obrigatoriedade:** Todo novo cliente deve concluir o onboarding. O status do Tenant só muda para "Ativo" após a conclusão.
- **Sinônimos:** "primeiro acesso", "ativação guiada".

### Soft Delete (Deleção Lógica)

Política de negócio onde registros não são removidos fisicamente do sistema. Em vez disso, são marcados com data de exclusão (campo `deleted_dt`), tornando-se invisíveis para o usuário comum mas preservados para auditoria e conformidade.

- **Aplica-se a:** Unidades de Negócio (libera CNPJ para reúso), Produtos do Catálogo, Usuários.
- **Impacto na unicidade:** CNPJ de unidade desativada pode ser reutilizado por novo cadastro ativo.

---

## 6. Fases do Projeto

### Fase 0 — Portal Administrativo SaaS (Este Projeto)

Escopo atual: construção do Core administrativo da FBSO Platform. Inclui o portal administrativo interno, gestão de contas e planos, RBAC, portal do cliente com onboarding e App Switcher, cadastro de Unidades de Negócio e Catálogo de Produtos.

- **Duração:** 14 semanas (24/07 a 30/10/2026).
- **Entregas:** D1 a D7, organizadas em 7 marcos (M1 a M7).

### Fases Futuras — Módulos-Produto

Desenvolvimento dos módulos Tributali-Engine e Storekeeper Portal, que serão acoplados ao Core construído nesta Fase 0. Fora do escopo deste projeto.

---

## 7. Indicadores e Métricas

### KPIs de Adoção (A)

Indicadores que medem a adoção do portal pelos clientes: A1 (Percentual de onboarding autônomo), A2 (Tempo médio de onboarding) e A3 (Taxa de abandono de onboarding).

### KPIs de Operação (O)

Indicadores que medem a eficiência operacional do time interno: O1 (Tempo de ativação de conta), O2 (Cobertura de auditoria administrativa), O3 (Incidentes de vazamento entre Unidades de Negócio) e O4 (Tempo de bloqueio de acesso).

### KPIs de Satisfação (S)

Indicadores que medem a percepção de qualidade: S1 (NPS do time interno), S2 (Satisfação do cliente com onboarding) e S3 (Chamados de suporte por cliente novo).

### KPIs de Prontidão (P)

Indicadores que medem a preparação para fases futuras: P1 (Cobertura de catálogo pré-mapeamento) e P2 (Tempo para ativar primeiro módulo-produto).

> **Nota:** A definição completa de cada KPI (fórmula, meta, frequência de reporte) está documentada em [MATRIZ-KPI.md](./MATRIZ-KPI.md).

---

## 8. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | 2026-07-13 | Criação inicial: conceitos da plataforma, entidades de negócio, perfis de acesso, fluxos, fases | Time de Negócios |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: domain-modeling, agile-ba-practices.*
🔍 *Revisado pelo skill caveman-review em 15/07/2026. Ajustes aplicados: typo "Componento"→"Componente", referência do charter corrigida, enums bilíngues (PT/EN), lista de KPIs completa.*
