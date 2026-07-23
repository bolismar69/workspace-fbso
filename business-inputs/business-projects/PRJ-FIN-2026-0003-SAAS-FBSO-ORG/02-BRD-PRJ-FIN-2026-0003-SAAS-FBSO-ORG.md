# Business Requirements Document (BRD): FBSO Platform — Portal Administrativo SaaS

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Documento** | BRD-FBSO-PLATFORM-2026-001 |
| **Versão** | 1.1 — Revisada conforme Docs Review (15/07/2026) |
| **Data** | 13 de julho de 2026 |
| **Autor** | Time de Produto FBSO.ORG |
| **Status** | Aprovado |

---

## 1. Sumário Executivo (Executive Summary)

### O Problema de Negócio

A FBSO.ORG possui expertise consolidada em soluções fiscais e de varejo, mas não dispõe de uma plataforma SaaS unificada para gerenciar clientes, planos de assinatura e controle de acesso. Atualmente, qualquer operação de ativação de cliente, configuração de plano ou gestão de permissões é inexistente ou realizada de forma manual e desconectada, impedindo a escalabilidade do negócio.

### A Solução Proposta

Construir o **Portal Administrativo da FBSO Platform** — a fundação do futuro SaaS multi-produto. Este portal será a camada Core que permitirá ao time interno gerenciar contas de clientes, planos comerciais e permissões de acesso, ao mesmo tempo que oferece aos clientes uma experiência de autoatendimento para onboarding, cadastro de unidades de negócio e gestão de seu portfólio de produtos.

### Benefícios Esperados

- **Operação comercial estruturada:** Time interno gerencia clientes e planos de forma centralizada, reduzindo processos manuais e retrabalho
- **Base para produtos futuros:** Módulos como Tributali-Engine e Storekeeper Portal poderão ser acoplados ao Core sem reestruturação
- **Experiência profissional do cliente:** Portal de autoatendimento com onboarding guiado, App Switcher e menus adaptados ao plano contratado
- **Governança e segurança:** Controle granular de acesso por permissões, garantindo que cada usuário veja apenas o que seu papel permite

### Principais Métricas de Sucesso

| Métrica | Situação Atual | Meta |
|---------|---------------|------|
| Processo de ativação de cliente | Manual (baseline: ~2 dias úteis) | Ativação em até 5 minutos via portal |
| Clientes que completam onboarding sozinhos | 0% | ≥ 80% |
| Incidentes de vazamento de dados entre filiais | — | Zero incidentes |
| Satisfação do time interno com ferramentas administrativas | — | Nota ≥ 4,0 / 5,0 |

---

## 2. Objetivos de Negócio (Business Objectives)

### Objetivo Primário

Estabelecer a plataforma SaaS administrativa da FBSO.ORG como fundação operacional até o final do projeto, permitindo que o time interno gerencie clientes, planos e acessos de forma centralizada e que os clientes realizem autoatendimento básico.

### Objetivos Secundários

| # | Objetivo | Alinhamento Estratégico |
|---|----------|------------------------|
| O1 | Estruturar a operação comercial do SaaS com gestão de contas, planos e assinaturas | Viabiliza a comercialização futura de módulos-produto |
| O2 | Oferecer portal de autoatendimento ao cliente com onboarding guiado | Reduz custo de aquisição e suporte a clientes |
| O3 | Implementar governança de acessos baseada em papéis (RBAC) | Garante segurança e isolamento de dados entre filiais |
| O4 | Preparar arquitetura de produto para acoplamento de módulos futuros | Habilita cross-selling e upselling sem reestruturação |
| O5 | Criar cadastro de unidades de negócio e catálogo de produtos | Prepara a base de dados comercial para os módulos fiscais e de varejo |

### Critérios de Sucesso (SMART)

| # | Critério | Específico | Mensurável | Temporal |
|---|----------|-----------|------------|----------|
| C1 | Portal administrativo funcional | 100% das funcionalidades D1-D7 entregues | Quantidade de funcionalidades validadas | Ao final do projeto |
| C2 | Onboarding autônomo | Cliente cria conta, cadastra primeira Unidade de Negócio e acessa o portal sem ajuda | ≥ 80% dos clientes | Em até 3 meses pós-lançamento |
| C3 | Tempo de ativação de conta | Da solicitação à conta ativa | ≤ 5 minutos (baseline manual: ~2 dias úteis) | Desde o lançamento |
| C4 | Satisfação interna | Pesquisa NPS com time administrativo e comercial | Nota média ≥ 4,0/5,0 | Em até 1 mês pós-lançamento |
| C5 | Zero vazamento entre filiais | Nenhum usuário acessa dados de Unidade de Negócio não autorizada | 0 incidentes | Desde o lançamento |

---

## 3. Contexto e Antecedentes (Background & Context)

### Situação Atual

A FBSO.ORG está em processo de estruturação de seu portfólio de produtos SaaS. A empresa identificou duas grandes oportunidades de mercado:

1. **Tributali-Engine:** Solução fiscal para a Reforma Tributária (IBS/CBS), atendendo empresas e escritórios de contabilidade que precisam se adequar ao Split Payment e às novas regras de arrecadação.
2. **Storekeeper Portal:** Solução de varejo (PDV, estoque, gestão comercial) para supermercados, farmácias e lojistas.

Ambos os produtos compartilham a necessidade de uma plataforma comum para gerenciamento de clientes, planos, permissões e catálogo de produtos. Sem essa fundação, cada produto precisaria construir sua própria camada administrativa, gerando duplicação de esforço, custos multiplicados e experiência fragmentada para clientes que eventualmente contratem mais de um módulo.

### Visão de Produto (Modelo "Lego")

A FBSO Platform foi concebida como uma **Suíte Multi-Produto** operando no modelo de módulos ativáveis por plano:

- Um cliente contrata o plano desejado e os módulos correspondentes são liberados em sua conta
- O App Switcher (seletor de aplicativos) permite alternar entre módulos sem sair da plataforma
- As permissões de acesso controlam quais usuários veem quais módulos e funcionalidades
- Menus e interfaces se adaptam dinamicamente ao plano contratado e ao módulo ativo

### Direcionadores de Mercado

- **Reforma Tributária (IBS/CBS):** Cronograma de implementação pelo governo federal, criando demanda por soluções fiscais automatizadas
- **Digitalização do varejo:** Pequenos e médios lojistas buscando soluções integradas de gestão
- **Consolidação de plataformas:** Tendência de mercado para suítes multi-produto em vez de sistemas isolados

---

## 4. Análise de Partes Interessadas (Stakeholder Analysis)

| Grupo | Representantes | Interesse / Preocupação | Requisitos-Chave |
|-------|---------------|------------------------|------------------|
| **Diretoria FBSO.ORG** | Sócios / Diretores | Retorno sobre investimento, time-to-market, visão estratégica | Portal viabilize comercialização futura; custos controlados |
| **Time Administrativo** | Equipe de operações | Gerenciar clientes, ativar/suspender contas, controlar inadimplência | Interface simples; ações rápidas; histórico de auditoria |
| **Time Comercial** | Vendedores, Sales Ops | Gerir plano de clientes, propor upgrades, acompanhar contas | Visão clara do status de cada cliente; registro de plano contratado |
| **Time de Produto** | Product Owner, Analista de Negócios | Escopo bem definido, backlog priorizado, critérios de aceitação claros | Requisitos documentados; métricas de sucesso mensuráveis |
| **Cliente Final (futuro)** | Empresas e contadores | Portal funcional, onboarding simples, segurança de dados | Autoatendimento; acesso restrito por usuário; experiência intuitiva |
| **Time Técnico (futuro)** | Desenvolvedores | Clareza de requisitos de negócio; decisões de produto antes da codificação | Sem citações técnicas neste documento |

---

## 5. Definição de Escopo (Scope Definition)

### 5.1 Dentro do Escopo

**Bloco A — Operações Internas (Time FBSO.ORG)**

| Área | Funcionalidades |
|------|----------------|
| **Dashboard Administrativo** | Métricas de contas ativas; planos contratados; status de tenants; visão consolidada da base |
| **Gestão de Contas (Tenants)** | Ativação, suspensão e reativação de contas; acompanhamento de status; histórico de ações |
| **Gestão de Planos** | Cadastro de planos comerciais; definição de módulos por plano; valores e recorrências |
| **Gestão de Assinaturas** | Vinculação cliente-plano; data de vigência; status da assinatura |
| **Gestão de Usuários (RBAC)** | Cadastro de usuários; definição de papéis; vinculação por Unidade de Negócio; controle de visibilidade |

**Bloco B — Experiência do Cliente (Auto-Serviço)**

| Área | Funcionalidades |
|------|----------------|
| **Portal do Cliente** | Login; dashboard pessoal; área de perfil |
| **Onboarding Guiado** | Fluxo passo a passo de primeiro acesso; criação da primeira Unidade de Negócio |
| **App Switcher** | Seletor de módulos (estrutura preparada para produtos futuros) |
| **Unidades de Negócio** | Cadastro de CNPJs/filiais; estrutura hierárquica Matriz/Filial; dados de regime tributário |
| **Catálogo de Produtos/Serviços** | Cadastro do portfólio comercial do cliente (nome, tipo, classificação); estrutura preparada para mapeamento fiscal futuro |

### 5.2 Fora do Escopo

| Área | Exclusões |
|------|-----------|
| **Módulo Tributali-Engine** | Cálculos de IBS/CBS; Split Payment; mapeamento fiscal (NCM/NBS/CNAE); geração de guias |
| **Módulo Storekeeper Portal** | PDV; controle de estoque; painel de vendas; integrações com adquirentes |
| **Comercialização** | Nenhum módulo-produto disponível para venda ou ativação por clientes; sem faturamento real |
| **Integrações Externas** | ERP de clientes; gateways de pagamento; sistemas bancários ou governamentais |
| **Funcionalidades Operacionais** | Emissão de pedidos com cálculos fiscais; workflow Quote→Order→Invoice→Payment; conciliação contábil |

---

## 6. Requisitos de Negócio (Business Requirements)

### 6.1 Requisitos Funcionais (Alto Nível)

#### Bloco A: Operações Internas

**BR-A01: Dashboard Administrativo**
- **Requisito:** O sistema deve apresentar um painel de controle para o time FBSO.ORG com métricas operacionais do SaaS
- **Valor de Negócio:** Visibilidade em tempo real da operação; tomada de decisão baseada em dados
- **Prioridade:** Must Have
- **Critério de Aceitação:** Dashboard exibe contas ativas, distribuição por plano, status de tenants e permite filtro por período

**BR-A02: Ativação e Gestão de Contas de Clientes**
- **Requisito:** O time administrativo deve conseguir ativar, suspender e reativar contas de clientes (Tenants)
- **Valor de Negócio:** Controle operacional sobre a base de clientes; gestão de inadimplência
- **Prioridade:** Must Have
- **Critério de Aceitação:** Ação de ativação/suspensão concluída em até 3 cliques; registro de auditoria gerado automaticamente

**BR-A03: Configuração de Planos Comerciais**
- **Requisito:** O time de produto deve cadastrar e configurar planos comerciais, definindo nome, valor, recorrência e quais módulos cada plano libera
- **Valor de Negócio:** Flexibilidade para criar e ajustar ofertas comerciais sem depender de desenvolvimento
- **Prioridade:** Must Have
- **Critério de Aceitação:** Plano cadastrado fica disponível para vinculação a clientes em tempo real

**BR-A04: Vinculação de Assinaturas**
- **Requisito:** Cada cliente deve ser vinculado a um plano, com data de início, vigência e status
- **Valor de Negócio:** Base para faturamento futuro e controle de acesso a módulos
- **Prioridade:** Must Have
- **Critério de Aceitação:** Cliente com assinatura ativa acessa o portal; cliente com assinatura suspensa tem acesso bloqueado

**BR-A05: Gestão de Usuários e Permissões**
- **Requisito:** O administrador do tenant deve convidar usuários, atribuir papéis (MVP: Admin, Gerente, Operador — papel "Auditor" previsto para fase futura) e restringir acesso por Unidade de Negócio
- **Valor de Negócio:** Segurança; isolamento de dados entre filiais; conformidade com governança corporativa
- **Prioridade:** Must Have
- **Critério de Aceitação:** Usuário sem permissão para uma Unidade de Negócio não visualiza seus dados; menus e botões respeitam o papel do usuário

#### Bloco B: Experiência do Cliente

**BR-B01: Portal do Cliente com Autenticação**
- **Requisito:** O cliente deve acessar um portal próprio após autenticação segura, visualizando dashboard, perfil e menus adaptados ao seu plano
- **Valor de Negócio:** Experiência profissional; percepção de produto completo; base para autoatendimento
- **Prioridade:** Must Have
- **Critério de Aceitação:** Login funcional; menus renderizados conforme permissões do usuário

**BR-B02: Onboarding Guiado de Primeiro Acesso**
- **Requisito:** No primeiro acesso, o cliente deve ser conduzido por um fluxo passo a passo: validação de dados, criação da primeira Unidade de Negócio e boas-vindas ao portal
- **Valor de Negócio:** Redução de atrito na ativação; diminuição de chamados de suporte; primeira experiência positiva
- **Prioridade:** Must Have
- **Critério de Aceitação:** ≥ 80% dos clientes concluem o fluxo sem intervenção do time FBSO.ORG

**BR-B03: App Switcher (Seletor de Aplicativos)**
- **Requisito:** O portal deve exibir um seletor que permita ao usuário alternar entre os módulos disponíveis no seu plano
- **Valor de Negócio:** Fundação para o modelo multi-produto; experiência de suíte integrada; preparação para cross-selling
- **Prioridade:** Must Have
- **Critério de Aceitação:** App Switcher visível no topo do portal; menus laterais se adaptam ao módulo selecionado

**BR-B04: Cadastro de Unidades de Negócio**
- **Requisito:** O cliente deve cadastrar múltiplos CNPJs/filiais, com estrutura hierárquica (Matriz/Filial), regime tributário e dados cadastrais
- **Valor de Negócio:** Base para operação multi-company dos módulos futuros; pré-requisito para Tributali-Engine e Storekeeper
- **Prioridade:** Must Have
- **Critério de Aceitação:** CNPJs duplicados detectados e bloqueados para o mesmo tenant ativo; hierarquia Matriz/Filial visível na interface

**BR-B05: Catálogo de Produtos/Serviços**
- **Requisito:** O cliente deve cadastrar itens do seu portfólio comercial com nome, SKU, tipo (Produto, Serviço) e status
- **Valor de Negócio:** Prepara a base de dados para o futuro mapeamento fiscal; reduz esforço na ativação do Tributali-Engine
- **Prioridade:** Must Have
- **Critério de Aceitação:** Cadastro de produto funcional; tabela `product_tax_mapping` com schema definido e contrato de interface documentado para acoplamento do Tributali-Engine, sem ativar funcionalidades tributárias.

### 6.2 Requisitos Não-Funcionais (Alto Nível)

| ID | Categoria | Requisito | Prioridade |
|----|-----------|-----------|------------|
| BR-NFR01 | **Disponibilidade** | Portal deve estar disponível 99,5% do tempo em horário comercial (6h-23h) | Must Have |
| BR-NFR02 | **Segurança** | Autenticação obrigatória para qualquer acesso a dados; isolamento total entre tenants e entre unidades de negócio | Must Have |
| BR-NFR03 | **Auditabilidade** | Toda ação administrativa (ativação, suspensão, alteração de permissão) deve gerar registro de auditoria com identificação do responsável e data/hora | Must Have |
| BR-NFR04 | **Usabilidade** | Portal do cliente deve ser intuitivo a ponto de não exigir treinamento; time interno necessita de no máximo 2 horas de capacitação | Should Have |
| BR-NFR05 | **Performance** | Tempo de carregamento de telas principais ≤ 3 segundos em conexão padrão de internet | Should Have |
| BR-NFR06 | **Escalabilidade** | Estrutura deve suportar acoplamento de novos módulos sem reestruturação do Core | Must Have |
| BR-NFR07 | **Acessibilidade** | Portal do cliente deve atender critérios mínimos de acessibilidade (contraste, navegação por teclado, textos alternativos) | Could Have |
| BR-NFR08 | **Idioma** | Interface em português (Brasil); estrutura preparada para multi-idioma no futuro | Must Have (PT-BR) / Could Have (multi-idioma) |

---

## 7. Premissas (Assumptions)

| # | Premissa | Impacto se Inválida |
|---|----------|---------------------|
| A1 | A visão de produto multi-módulo (Tributali-Engine, Storekeeper Portal) possui direção estratégica definida, mas reconhece-se que requisitos detalhados ainda estão em maturação. Contratos de interface entre Core e módulos serão congelados até M2 para isolar o Core de oscilações na visão de produto. | Retrabalho estrutural no Core se os módulos exigirem mudanças de arquitetura — mitigado pelo congelamento dos contratos de interface |
| A2 | O time técnico, embora reduzido, possui competência para entregar o escopo com qualidade | Atrasos ou necessidade de reforço externo |
| A3 | Stakeholders designados (Comercial, Administrativo) terão disponibilidade para validações periódicas | Atraso nas validações; retrabalho por desalinhamento |
| A4 | Clientes potenciais participarão como early adopters para validação do Portal do Cliente | Portal validado apenas internamente; risco de não atender necessidades reais |
| A5 | A Reforma Tributária (IBS/CBS) seguirá o cronograma de implementação previsto | Pressão para antecipar o Tributali-Engine antes do Core estar pronto |
| A6 | O modelo de negócio baseado em planos com módulos ativáveis será mantido | Redesenho da lógica comercial e de acessos |

---

## 8. Restrições (Constraints)

| # | Categoria | Restrição |
|---|-----------|-----------|
| C1 | **Recurso** | Time técnico reduzido — velocidade de entrega limitada pela capacidade atual |
| C2 | **Orçamento** | Sem previsão de contratações significativas ou investimentos em infraestrutura de grande porte |
| C3 | **Negócio** | Nenhum módulo-produto pode ser comercializado antes da conclusão do Core administrativo |
| C4 | **Sequenciamento** | Portal do Cliente deve estar funcional antes do início do desenvolvimento de qualquer módulo-produto |
| C5 | **Neutralidade** | Decisões de negócio não devem ser restringidas por escolhas técnicas |

---

## 9. Dependências (Dependencies)

| # | Dependência | Descrição | Responsável | Status | Risco |
|---|-------------|-----------|-------------|--------|------|
| D1 | Definição do portfólio de planos | Nomes, preços e módulos de cada plano comercial precisam ser definidos pelo time comercial | Líder Comercial | Pendente | Alto |
| D2 | Papéis e permissões (RBAC) | Mapeamento final de papéis × funcionalidades precisa ser validado | Dono do Produto | Pendente | Médio |
| D3 | Early adopters para validação | Clientes parceiros precisam ser selecionados e engajados para testes do Portal do Cliente | Líder Comercial | Pendente | Médio |
| D4 | Conteúdo do onboarding | Textos, passos e mensagens do fluxo de onboarding precisam ser definidos | Analista de Negócios | Pendente | Baixo |
| D5 | Aprovação do Project Charter | Charter aprovado pela diretoria conforme PC-FBSO-PLATFORM-2026-001 | Coordenador do Projeto | Concluído | Alto |

---

## 10. Riscos e Mitigações (Risks & Mitigation)

| # | Risco | Impacto | Probabilidade | Mitigação |
|---|-------|---------|---------------|-----------|
| R1 | Expansão prematura do escopo com funcionalidades dos módulos Tributali-Engine ou Storekeeper | Alto | Alta | Reforçar o Out of Scope a cada revisão; funcionalidades de módulos viram backlog futuro |
| R2 | Equipe reduzida não entrega no prazo esperado | Alto | Alta | Priorização MoSCoW; foco no MVP de cada entrega; avaliar reforço pontual |
| R3 | Requisitos de produto ainda imaturos e mudam durante o projeto | Alto | Média | Congelar contratos de interface do Core; mudanças afetam apenas módulos futuros |
| R4 | Portal validado sem feedback real de clientes | Alto | Média | Engajar 2-3 early adopters para validações quinzenais do Portal do Cliente |
| R5 | RBAC fica complexo demais para a fase atual | Médio | Média | Começar com 3 papéis essenciais (Admin, Gerente, Operador). Papel "Auditor" documentado e com schema previsto, mas não implementado nesta fase. |
| R6 | Desalinhamento de expectativa: "produto pronto para vender" | Alto | Baixa | Comunicação explícita de que esta fase entrega a fundação, não um produto comercializável |

---

## 11. Métricas de Sucesso e KPIs (Success Metrics & KPIs)

### Métricas de Adoção

| Métrica | Meta | Prazo |
|---------|------|-------|
| Clientes que completam onboarding sem ajuda | ≥ 80% | 3 meses pós-lançamento |
| Tempo médio de onboarding (do login ao portal liberado) | ≤ 10 minutos | Desde o lançamento |
| Usuários ativos no portal do cliente (mensal) | ≥ 90% dos usuários cadastrados | 3 meses pós-lançamento |

### Métricas de Impacto no Negócio

| Métrica | Meta | Prazo |
|---------|------|-------|
| Tempo de ativação de nova conta (solicitação → ativa) | Imediato (sem intervenção manual) | Desde o lançamento |
| Tempo para ativar primeiro módulo-produto após Core pronto | ≤ 1 sprint | Na fase seguinte |
| Satisfação do time interno (NPS) | ≥ 4,0 / 5,0 | 1 mês pós-lançamento |

### Métricas de Qualidade

| Métrica | Meta | Prazo |
|---------|------|-------|
| Incidentes de vazamento de dados entre filiais | Zero | Desde o lançamento |
| Disponibilidade do portal | 99,5% | Desde o lançamento |
| Registros de auditoria sem falhas | 100% das ações cobertas (ver BR-NFR03) | Desde o lançamento |

---

## 12. Timeline e Marcos (Timeline & Milestones)

| Fase | Marco | Duração Estimada | Data-Alvo |
|------|-------|-----------------|-----------|
| **Planejamento** | Aprovação do BRD, Epics e Features | — | 31/07/2026 |
| **Fundação** | Portal Admin — Versão Inicial com dashboard e visão de contas | — | 15/08/2026 |
| **Operação** | Gestão de Contas e Planos funcional | — | 31/08/2026 |
| **Governança** | Usuários e Permissões (RBAC) completo | — | 15/09/2026 |
| **Experiência** | Portal do Cliente, Onboarding e App Switcher | — | 30/09/2026 |
| **Catálogo** | Unidades de Negócio e Catálogo de Produtos | — | 15/10/2026 |
| **Homologação** | Aceite final com todas as entregas validadas | — | 30/10/2026 |

> **Nota:** Datas definidas e alinhadas com o Project Charter (PC-FBSO-PLATFORM-2026-001, §7). Cadência quinzenal. Cronograma completo: 14 semanas (24/07 a 30/10/2026).

---

## 13. Orçamento (Budget Estimate)

| Categoria | Descrição | Estimativa (R$) |
|-----------|-----------|------------|
| Time de Desenvolvimento | Equipe técnica alocada ao projeto | 100.000,00 |
| Ferramentas e Serviços | Infraestrutura, autenticação, monitoramento | 25.000,00 |
| Validação com Clientes | Sessões com early adopters, coleta de feedback | 10.000,00 |
| Reserva Técnica | Contingência para riscos | 18.000,00 |
| **Total** | | **R$ 153.000,00** (estimativa preliminar — vide planilha orçamentária complementar) |

> **Nota:** Valores acima são placeholders estimados. O orçamento detalhado será elaborado após o dimensionamento do backlog e definição da equipe alocada (TECHNICAL-TEAM-MAP.md). Enquanto o time técnico não for definido, os valores de cada rubrica permanecem como estimativa preliminar.

---

## 14. Ambiente de Homologação

As validações pelos stakeholders e early adopters serão realizadas em ambiente de homologação dedicado, separado do ambiente de desenvolvimento:

- **Staging environment** provisionado junto com a entrega do Portal Admin (M2)
- **Uptime esperado:** ≥ 95% durante o período de validação
- **Acesso:** Time interno FBSO.ORG e early adopters convidados
- **Cadência de deploy:** Alinhada com entregas quinzenais. Cada entrega sobe para homologação na data-alvo.

---

## 15. Plano de Comunicação

| O Quê | Quem Reporta | Para Quem | Quando | Canal |
|-------|-------------|-----------|--------|-------|
| Status Report do Projeto | Coordenador do Projeto | Diretoria, PO, Stakeholders | Quinzenal (após cada checkpoint) | Reunião de 1h com pauta objetiva |
| Atualização de Riscos | Coordenador do Projeto | Diretoria, PO | Quinzenal (no Status Report) | Mesma reunião |
| Validação de Entregas | Coordenador do Projeto | Líder Administrativo, Líder Comercial | A cada entrega (datas-alvo) | Sessão de demonstração em homologação |
| Feedback de Early Adopters | Líder Comercial | Coordenador do Projeto, PO | Quinzenal | Relatório de feedback consolidado |
| Comunicação de Lançamento Interno | Coordenador do Projeto + Diretoria | Toda a FBSO.ORG | M1 (Kickoff) e M7 (Aceite Final) | Email formal + apresentação |

---

## 16. Definição de Pronto (Definition of Done)

Uma entrega (D1-D7) é considerada concluída quando:

1. **Código:** Funcionalidade implementada e revisada por par
2. **Testes:** Cobertura de testes automatizados ≥ 70% para módulos Core
3. **Implantação:** Deploy realizado no ambiente de homologação
4. **Validação:** Stakeholder designado aprovou a entrega em sessão de demonstração
5. **Documentação:** Endpoints documentados no OpenAPI spec (quando aplicável); ADRs atualizados
6. **Auditoria:** Ações administrativas registram trilha de auditoria (quem, quando, o quê)

---

## 17. Aprovação e Assinatura (Approval & Sign-off)

| Papel | Nome | Assinatura | Data |
|-------|------|-----------|------|
| **Patrocinador** — Diretoria FBSO.ORG | | | |
| **Líder do Projeto** — Coordenador | | | |
| **Dono do Produto** — Product Owner | | | |
| **Líder Comercial** | | | |
| **Líder Administrativo** | | | |

---

## Glossário de Termos de Negócio

| Termo | Definição |
|-------|-----------|
| **Tenant** | Conta corporativa de um cliente na plataforma. Representa a empresa que contratou o SaaS. |
| **Unidade de Negócio (Business Unit)** | Um CNPJ ou filial vinculada a um Tenant. Cada unidade pode ter configurações próprias. |
| **Plano (Plan)** | Pacote comercial contratado pelo cliente, definindo módulos disponíveis e condições de cobrança. |
| **Assinatura (Subscription)** | Vínculo ativo entre um Tenant e um Plano, com vigência e status definidos. |
| **App Switcher** | Seletor de aplicativos no portal que permite alternar entre módulos sem sair da plataforma. |
| **RBAC** | Controle de Acesso Baseado em Papéis — define o que cada usuário pode ver e fazer. |
| **Onboarding** | Fluxo guiado de primeiro acesso do cliente ao portal. |
| **Core** | Camada fundamental da plataforma — gerencia contas, planos, usuários e permissões. |
| **Módulo / Produto** | Solução de negócio da plataforma (ex: Tributali-Engine, Storekeeper Portal). |
| **MoSCoW** | Técnica de priorização: Must Have, Should Have, Could Have, Won't Have. |

---

> **Este documento é estritamente de negócio.** Detalhamentos técnicos, escolhas de arquitetura e definições de implementação serão tratados em documentos complementares (FRS, ADRs, Especificações Técnicas).

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: brd-creation, agile-ba-practices, brainstorming. Revisão 1.1 baseada no Docs Review (15/07/2026) — skills: caveman, caveman-review.*
