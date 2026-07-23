# Project Charter: FBSO Platform — Portal Administrativo SaaS
## Document ID: PC-FBSO-PLATFORM-2026-001

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-FIN-2026-0003-SAAS-FBSO-ORG |
| **Produto** | FBSO Platform (Portal Administrativo) |
| **Data de Elaboração** | 13 de julho de 2026 |
| **Versão** | 1.1 — Revisada conforme Caveman Review (15/07/2026) |
| **Patrocinador** | Diretoria FBSO.ORG |
| **Metodologia** | Híbrida (Ágil com marcos clássicos) |
| **Status** | Aprovado |

**Documentos Relacionados:**
- `BUSINESS-REQUIREMENTS.md` (BRD) — Requisitos de negócio detalhados, seção 12: cronograma
- `.specs/architecture/` — ADRs, modelagem de dados, diagramas (documentos técnicos complementares)
- `PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/` — Projeto da Reforma Tributária (módulos futuros)

---

### 1. Declaração do Problema (Problem Statement)

A FBSO.ORG possui know-how em soluções fiscais e de varejo, mas não dispõe de uma plataforma SaaS unificada para comercializar esses produtos de forma integrada. Hoje, não existe um portal centralizado que permita gerenciar clientes, planos de assinatura e controle de acesso — etapa fundamental para estruturar a operação comercial antes que qualquer produto seja oferecido ao mercado.

Sem essa fundação, o lançamento futuro dos módulos Tributali-Engine (fiscal/tributário) e Storekeeper Portal (varejo/PDV) será inviável. Cada produto exigiria estrutura administrativa própria. Resultado: retrabalho, custos duplicados e experiência fragmentada para o cliente.

### 2. Propósito do Projeto (Project Purpose)

Construir o **Portal Administrativo da FBSO Platform** — a camada fundamental ("Core") do futuro SaaS multi-produto da FBSO.ORG. Este portal será a base operacional que permitirá à empresa gerenciar contas de clientes, planos de assinatura, permissões de usuários e cadastros de unidades de negócio de forma unificada. Trata-se do alicerce sobre o qual todos os produtos futuros (Tributali-Engine, Storekeeper Portal e demais módulos) serão acoplados como módulos ativáveis por plano, seguindo o modelo de "Suíte de Produtos" com experiência de uso unificada para o cliente.

#### 2.1 Visão de Longo Prazo

Quando um cliente do Storekeeper Portal precisar, no futuro, adequar-se ao Split Payment da Reforma Tributária, o upgrade será uma simples ativação de módulo na mesma plataforma — sem migração de dados, sem nova integração, sem troca de sistema.

### 3. Escopo (Scope)

#### 3.1 Dentro do Escopo (In Scope)

1. **Portal Administrativo Interno (Painel de Controle FBSO.ORG)**
   - Dashboard com métricas operacionais do SaaS: contas ativas, planos contratados, status de tenants, taxa de conversão de onboarding, distribuição por plano
   - Visão consolidada da base de clientes e suas respectivas configurações

2. **Gestão de Contas de Clientes (Tenants)**
   - Ativação, suspensão e reativação de contas de clientes
   - Acompanhamento do status de cada conta (ativo, inativo, em onboarding)
   - Histórico de ações administrativas sobre cada conta com registro de auditoria (quem fez, quando, qual ação)

3. **Gestão de Planos e Assinaturas**
   - Cadastro e configuração de planos comerciais (ex: Básico, Core, Full Suite)
   - Definição de quais módulos/produtos cada plano libera via tabela associativa `plan_modules`
   - Vinculação de clientes aos planos contratados
   - Suporte a diferentes modelos de cobrança (recorrência, vigência)

4. **Gestão de Usuários e Permissões (RBAC)**
   - Cadastro e administração de usuários do ecossistema
   - Definição de papéis de acesso. MVP: 3 papéis essenciais (Administrador do Tenant, Gerente de Unidade, Operador). Papel "Auditor" disponível para expansão futura.
   - Vinculação de usuários a Unidades de Negócio específicas
   - Controle granular de acesso: quem pode ver e fazer o quê, em qual unidade

5. **Portal do Cliente (Auto-Serviço)**
   - Tela de boas-vindas e onboarding guiado para novos clientes
   - Fluxo de cadastro autônomo com criação da primeira Unidade de Negócio
   - Área de perfil e configurações básicas da conta do cliente
   - App Switcher (seletor de aplicativos/módulos) — estrutura preparada para produtos futuros. Na versão atual, exibido como menu de navegação entre módulos; expande para switcher visual quando houver 2+ produtos ativos.

6. **Cadastro de Unidades de Negócio (Business Units)**
   - Registro de CNPJs, filiais e empresas vinculadas a um mesmo cliente
   - Estrutura hierárquica (Matriz/Filial)
   - Informações de regime tributário e dados cadastrais por unidade

7. **Estrutura-base de Catálogo de Produtos/Serviços**
   - Cadastro básico do portfólio comercial do cliente (nome, tipo, classificação)
   - Tabela `product_tax_mapping` criada com schema definido, sem ativar regras de tributação. Contrato de interface documentado para acoplamento futuro do Tributali-Engine.

#### 3.2 Fora do Escopo (Out of Scope)

1. **Módulo Tributali-Engine (Fiscal/Tributário)**
   - Cálculos de IBS/CBS, regras de Split Payment, mapeamento fiscal (NCM/NBS/CNAE)
   - Qualquer funcionalidade de retenção ou apuração de impostos
   - Engine de cálculo tributário e geração de guias de arrecadação

2. **Módulo Storekeeper Portal (Varejo/PDV)**
   - Frente de caixa (PDV), controle de estoque, painel de vendas
   - Integrações comerciais com adquirentes e gateways de pagamento

3. **Comercialização de Produtos Finais**
   - Nenhum módulo-produto será disponibilizado para venda ou ativação por clientes nesta fase
   - Não haverá faturamento de clientes via plataforma (apenas estrutura de plano cadastrada)

4. **Integrações Externas**
   - Integração com ERPs de clientes (Totvs, SAP, Omie)
   - Integração com gateways de pagamento para transações financeiras reais
   - Integração com sistemas bancários ou de arrecadação governamental

5. **Funcionalidades Operacionais Avançadas dos Módulos**
   - Emissão de pedidos e faturas com cálculos fiscais
   - Workflow de vendas (Quote → Order → Invoice → Payment)
   - Conciliação contábil e relatórios fiscais

6. **Migração de Dados de Clientes Existentes**
   - A FBSO.ORG não possui base de clientes pré-existente em plataforma digital. Não há migração de dados prevista para esta fase.

7. **Treinamento do Time Interno**
   - Capacitação dos operadores do Portal Administrativo será tratada como atividade paralela conduzida pelo Líder Administrativo, fora do escopo de desenvolvimento. Sessões de demonstração ocorrerão nos checkpoints quinzenais.

### 4. Entregas (Deliverables)

| # | Entrega | Critérios de Aceitação | Data-Alvo |
|---|---------|------------------------|-----------|
| D1 | **Portal Administrativo Interno** | Dashboard funcional exibindo: contas ativas, plano contratado, status de tenants, taxa de conversão de onboarding, distribuição por plano. Time interno da FBSO.ORG consegue visualizar e filtrar a base completa de clientes. | 15/08/2026 |
| D2 | **Módulo de Gestão de Contas** | Time administrativo consegue ativar, suspender e reativar contas de clientes. Cada ação gera registro de auditoria (quem fez, quando, qual ação). Histórico de ações consultável via painel administrativo. Cliente recebe notificação de ativação da conta. | 31/08/2026 |
| D3 | **Módulo de Planos e Assinaturas** | Planos comerciais cadastráveis com definição de módulos incluídos e valores. Cliente vinculado a plano com data de vigência e status. Tabela associativa `plan_modules` implementada como contrato de interface para acoplamento de módulos futuros. | 31/08/2026 |
| D4 | **Módulo de Usuários e Permissões** | Administrador do tenant consegue convidar usuários, definir papéis (MVP: Admin, Gerente, Operador) e restringir acesso por Unidade de Negócio. Usuário sem permissão não enxerga funcionalidades não autorizadas. Papel "Auditor" documentado para fase posterior. | 15/09/2026 |
| D5 | **Portal do Cliente** | Cliente consegue realizar login, passar por fluxo guiado de onboarding (criação da primeira Unidade de Negócio), acessar área de perfil e visualizar menus adaptados ao seu plano. Menu de navegação entre módulos visível (expande para App Switcher visual quando houver 2+ produtos ativos). | 30/09/2026 |
| D6 | **Cadastro de Unidades de Negócio** | Cliente cadastra múltiplos CNPJs/filiais com estrutura hierárquica (Matriz/Filial). Informações de regime tributário e dados cadastrais armazenados por unidade. Validação de duplicidade de CNPJ ativo. | 15/10/2026 |
| D7 | **Catálogo de Produtos/Serviços** | Cliente cadastra itens do seu portfólio comercial (nome, SKU, tipo). Tabela `product_tax_mapping` criada com schema definido e contrato de interface documentado para acoplamento do Tributali-Engine, sem ativar funcionalidades tributárias. | 15/10/2026 |

> **Nota:** Datas definidas em conjunto com o time do projeto. Cadência quinzenal de entregas. Alinhado com o cronograma da seção 12 do BRD (BUSINESS-REQUIREMENTS.md). D6 e D7 podem parcialmente paralelizar com D5 — dependência é mais frouxa do que o diagrama de marcos sugere.

### 5. Partes Interessadas e Matriz RACI (Stakeholders & RACI)

| Parte Interessada | Papel no Projeto | D1 Portal Admin | D2 Gestão Contas | D3 Planos | D4 Usuários | D5 Portal Cliente | D6 Unid. Negócio | D7 Catálogo |
|-------------------|------------------|----|----|----|----|----|----|----|
| **Diretoria FBSO.ORG** | Patrocinador (Sponsor) | A | A | A | A | A | A | A |
| **Coordenador do Projeto** | Líder do Projeto | R | R | R | R | R | R | R |
| **Dono do Produto** | Product Owner — gestor, não executor direto | C | C | C | C | C | C | C |
| **Analista de Negócios** | Requisitos e Regras | C | R | R | R | C | R | R |
| **Líder Comercial** | Visão de Mercado | I | C | R | I | C | C | R |
| **Líder Administrativo** | Operação do Portal | C | R | C | R | I | C | I |
| **Time de Vendas** | Uso Futuro do Portal | I | I | C | I | I | I | I |
| **Cliente Final (futuro)** | Usuário Final | I | I | I | I | C | C | C |

*R=Responsável (executa), A=Autoridade (aprova — apenas um por entrega), C=Consultado (fornece insumos), I=Informado (recebe atualizações)*

> **Nota:** O Dono do Produto atua como gestor e ponto de decisão de produto, não como executor direto. Por isso figura como Consultado em todas as entregas. O Coordenador do Projeto concentra a responsabilidade de execução (R).

### 6. Critérios de Sucesso (Success Criteria)

| # | Critério | Indicador | Meta |
|---|----------|-----------|------|
| C1 | Portal administrativo operacional para o time interno | Percentual de funcionalidades administrativas disponíveis vs. planejadas | 100% das entregas D1-D4 concluídas e validadas |
| C2 | Tempo de ativação de nova conta de cliente | Tempo entre solicitação e conta ativa | Redução de processo manual (baseline: ~2 dias úteis) para ativação em até 5 minutos via portal |
| C3 | Cliente realiza onboarding autônomo | Percentual de clientes que completam o fluxo de onboarding sem intervenção do time FBSO.ORG | ≥ 80% dos clientes concluem o onboarding sozinhos |
| C4 | Estrutura preparada para acoplar novos módulos | Tempo estimado para ativar o primeiro módulo-produto (ex: Tributali-Engine) após conclusão do Core | Menos de 1 sprint (2 semanas) de desenvolvimento para iniciar a integração de um novo módulo |
| C5 | Isolamento de acesso entre Unidades de Negócio | Número de incidentes de vazamento de dados entre filiais | Zero incidentes reportados |
| C6 | Satisfação do time interno com o portal | Pesquisa de satisfação com time administrativo e comercial | Nota média ≥ 4,0 em escala de 1-5 |
| C7 | Execução dentro do orçamento aprovado | Desvio entre custo real e custo planejado | Desvio ≤ 10% do orçamento total da fase (ver seção 12) |
| C8 | Qualidade das entregas | Cobertura de testes automatizados e disponibilidade do ambiente de homologação | Cobertura de testes ≥ 70% para módulos Core; uptime do ambiente de homologação ≥ 95% durante período de validação |

### 7. Marcos do Projeto (Milestones)

| Marco | Descrição | Data-Alvo | Dependências |
|-------|-----------|-----------|--------------|
| **M1: Kickoff** | Alinhamento formal com stakeholders, validação do escopo, definição do time e comunicação formal sobre natureza do projeto (fundação administrativa, não produto comercializável) | 24/07/2026 | Aprovação deste Project Charter |
| **M2: Portal Admin — Versão Inicial** | Dashboard interno com visão de contas e métricas básicas operacional. Contratos de interface (Core↔Módulos) documentados em repositório compartilhado. | 15/08/2026 | M1 concluído |
| **M3: Gestão de Contas e Planos** | Módulos de ativação/suspensão de tenants e configuração de planos funcionais | 31/08/2026 | M2 concluído |
| **M4: Usuários e Permissões** | Estrutura de RBAC completa: 3 papéis (MVP), vinculação por Unidade de Negócio, controle de acesso | 15/09/2026 | M3 concluído |
| **M5: Portal do Cliente** | Onboarding guiado, menu de navegação entre módulos e área do cliente disponíveis | 30/09/2026 | M4 concluído |
| **M6: Unidades de Negócio e Catálogo** | Cadastro de filiais/CNPJs e estrutura de portfólio comercial. Início pode paralelizar com final de M5. | 15/10/2026 | M5 concluído (parcial — D6/D7 podem iniciar durante validação de D5) |
| **M7: Aceite Final** | Homologação completa de D1-D7, validação do patrocinador e encerramento do projeto | 30/10/2026 | M6 concluído |

> **Nota:** Datas definidas em conjunto com o time do projeto seguindo cadência quinzenal. Cronograma completo: 14 semanas (24/07 a 30/10/2026). Alinhado com a seção 12 do BRD. Atenção: janela de homologação (M6→M7) é de apenas 2 semanas para 7 entregas. Recomenda-se validação contínua a cada entrega em ambiente de homologação dedicado, não concentrada no final.

### 8. Registro de Riscos (Risk Register)

| ID | Risco | Categoria | Probabilidade | Impacto | Severidade | Mitigação | Responsável |
|----|-------|-----------|---------------|---------|------------|-----------|-------------|
| R1 | **Expansão prematura do escopo** — stakeholders pressionam para incluir funcionalidades dos módulos Tributali-Engine ou Storekeeper ainda nesta fase | Escopo | Alta | Alta | **Crítica** | Checklist de scope creep com critérios de rejeição automática aplicado a cada revisão. Qualquer funcionalidade de produto-módulo deve ser registrada como backlog futuro, não como mudança de escopo. | Coordenador do Projeto |
| R2 | **Equipe reduzida não consegue entregar no prazo esperado** — time técnico pequeno limita a velocidade de desenvolvimento | Recurso | Alta | Alta | **Crítica** | Priorização rigorosa das entregas (MoSCoW). Foco no mínimo viável para cada entrega antes de refinamentos. Plano de contingência: D7 (Catálogo) é primeira candidata a escopo reduzido; D5 (Portal do Cliente) pode ter funcionalidades secundárias postergadas. Avaliar contratação de reforço pontual se necessário. | Coordenador do Projeto |
| R3 | **Requisitos de produto ainda imaturos** — visão dos módulos futuros (Tributali-Engine, Storekeeper) muda durante o projeto e impacta o design do Core | Escopo | Média | Alta | **Alta** | Definir contratos de interface (o que o Core oferece aos módulos) e documentá-los em repositório compartilhado até M2. Congelar contratos após M2. Mudanças na visão de produto não devem reestruturar o Core, apenas os módulos futuros. | Dono do Produto |
| R4 | **Falta de validação com clientes reais** — portal é construído sem feedback do mercado, resultando em retrabalho futuro | Mercado | Média | Alta | **Alta** | Convidar 5 clientes ou parceiros estratégicos até M1; meta: 3 aceitarem participar como "early adopters" para validações periódicas do Portal do Cliente (D5). Critérios de seleção: porte compatível, interesse em SaaS fiscal, disponibilidade para feedback quinzenal. Feedbacks coletados a cada 2 semanas. | Líder Comercial |
| R5 | **Complexidade do RBAC vs. simplicidade do MVP** — sistema de permissões fica complexo demais para o momento atual, atrasando entregas | Escopo | Média | Média | **Média** | MVP entregue com 3 papéis essenciais (Admin, Gerente, Operador). Papel "Auditor" documentado e com schema previsto, mas não implementado nesta fase. Granularidade por tela/ação postergada. | Dono do Produto |
| R6 | **Falta de engajamento dos stakeholders operacionais** — time administrativo e comercial não prioriza tempo para validar o portal | Recurso | Média | Média | **Média** | Agendar checkpoints fixos (quinzenais) com pauta objetiva e duração máxima de 1 hora. Envolver liderança na cobrança de participação. | Coordenador do Projeto |
| R7 | **Expectativa desalinhada sobre "produto pronto para vender"** — diretoria ou time comercial entende que o portal administrativo já é um produto comercializável | Comunicação | Baixa | Alta | **Média** | Comunicação formal no M1 explicitando que esta fase entrega a fundação administrativa, não um produto final para o mercado. Mensagem de lançamento interno vs. lançamento comercial documentada e aprovada pela Diretoria. Reforço a cada marco. | Coordenador do Projeto |
| R8 | **Mudanças regulatórias na Reforma Tributária** — alterações nas regras de IBS/CBS durante o projeto impactam premissas do desenho futuro dos módulos | Externo | Baixa | Alta | **Média** | Monitorar calendário de regulamentação da Reforma Tributária. Estrutura do Core deve ser genérica o suficiente para acomodar ajustes regulatórios sem reestruturação. | Dono do Produto |

*Probabilidade/Impacto: Baixo, Médio, Alto*
*Severidade = Probabilidade × Impacto*

### 9. Premissas e Restrições (Assumptions and Constraints)

#### 9.1 Premissas (Assumptions)

- **A1:** O time técnico, embora reduzido, possui competência para desenvolver a plataforma com qualidade e dentro de prazos razoáveis.
- **A2:** A visão de produto multi-módulo (Tributali-Engine, Storekeeper Portal) possui direção estratégica definida, mas reconhece-se que requisitos detalhados ainda estão em maturação (ver R3). Contratos de interface entre Core e módulos serão congelados até M2 para isolar o Core de oscilações na visão de produto.
- **A3:** Os stakeholders designados (Comercial, Administrativo) terão disponibilidade para validações periódicas conforme o cronograma.
- **A4:** O Coordenador do Projeto convidará 5 clientes potenciais até M1 para atuarem como early adopters. Meta: ao menos 3 aceitarem participar das validações do Portal do Cliente (D5).
- **A5:** A Reforma Tributária (IBS/CBS) seguirá o cronograma de implementação previsto pelo governo, sem antecipações que pressionem o lançamento do Tributali-Engine antes do Core estar pronto.
- **A6:** O modelo de negócio baseado em planos com módulos ativáveis (Básico / Core / Full Suite) será mantido como estratégia comercial da FBSO Platform.

#### 9.2 Restrições (Constraints)

- **C1:** Time técnico reduzido — a velocidade de entrega é limitada pela capacidade atual da equipe.
- **C2:** Orçamento limitado para esta fase — não estão previstas contratações significativas ou investimentos em infraestrutura de grande porte. Ver seção 12 para detalhamento.
- **C3:** O Portal do Cliente (D5) deve estar funcional antes que qualquer módulo-produto comece a ser desenvolvido.
- **C4:** O projeto deve respeitar a neutralidade tecnológica nas definições de negócio — decisões técnicas não devem restringir opções comerciais futuras.
- **C5:** Nenhum produto-módulo pode ser comercializado antes da conclusão e estabilização do Core administrativo (esta é uma regra de negócio, mantida como restrição operacional).

### 10. Ambiente de Homologação

As validações pelos stakeholders (C6) e early adopters (R4) serão realizadas em ambiente de homologação dedicado, separado do ambiente de desenvolvimento. Características:

- **Staging environment** provisionado junto com D1 (Portal Admin)
- **Uptime esperado:** ≥ 95% durante o período de validação (alinhado com C8)
- **Acesso:** Time interno FBSO.ORG e early adopters convidados
- **Cadência de deploy:** Alinhada com entregas quinzenais (D1-D7). Cada entrega sobe para homologação na data-alvo.

### 11. Plano de Comunicação

| O Quê | Quem Reporta | Para Quem | Quando | Canal |
|-------|-------------|-----------|--------|-------|
| Status Report do Projeto | Coordenador do Projeto | Diretoria, PO, Stakeholders | Quinzenal (após cada checkpoint) | Reunião de 1h com pauta objetiva |
| Atualização de Riscos | Coordenador do Projeto | Diretoria, PO | Quinzenal (no Status Report) | Mesma reunião |
| Validação de Entregas (D1-D7) | Coordenador do Projeto | Líder Administrativo, Líder Comercial | A cada entrega (datas-alvo) | Sessão de demonstração em homologação |
| Feedback de Early Adopters | Líder Comercial | Coordenador do Projeto, PO | Quinzenal | Relatório de feedback consolidado |
| Comunicação de Lançamento Interno | Coordenador do Projeto + Diretoria | Toda a FBSO.ORG | M1 (Kickoff) e M7 (Aceite Final) | Email formal + apresentação |

### 12. Orçamento e Recursos

| Categoria | Descrição | Estimativa |
|-----------|-----------|------------|
| Time de Desenvolvimento | Equipe técnica reduzida (headcount existente) | Custo interno — sem contratação adicional prevista |
| Infraestrutura | Ambiente de homologação + produção inicial | Cloud — tier inicial de consumo |
| Ferramentas | Licenças de desenvolvimento, repositório, CI/CD | Ferramentas já licenciadas |
| Early Adopters | Eventual incentivo para participação | Simbólico — desconto futuro ou acesso antecipado |
| Contingência | Reserva para imprevistos | 10% do total da fase |

> **Nota:** Valores detalhados em planilha orçamentária complementar (documento financeiro interno). O orçamento total da fase é considerado "limitado" (C2) e não prevê contratações significativas. Meta C7: desvio ≤ 10%.

### 13. Definição de Pronto (Definition of Done)

Uma entrega (D1-D7) é considerada concluída quando:

1. **Código:** Funcionalidade implementada e revisada por par
2. **Testes:** Cobertura de testes automatizados ≥ 70% para módulos Core (C8)
3. **Implantação:** Deploy realizado no ambiente de homologação
4. **Validação:** Stakeholder designado aprovou a entrega em sessão de demonstração
5. **Documentação:** Endpoints documentados no OpenAPI spec (quando aplicável); ADRs atualizados
6. **Auditoria:** Ações administrativas registram trilha de auditoria (quem, quando, o quê)

### 14. Aprovação (Approval)

| Papel | Nome | Data |
|-------|------|------|
| **Patrocinador (Sponsor)** — Diretoria FBSO.ORG | | |
| **Líder do Projeto** — Coordenador do Projeto | | |
| **Dono do Produto** — Product Owner | | |

---

### Glossário de Termos de Negócio

| Termo | Definição |
|-------|-----------|
| **Tenant** | Conta corporativa de um cliente na plataforma. Representa a empresa que contratou o SaaS. |
| **Unidade de Negócio (Business Unit)** | Um CNPJ ou filial vinculada a um Tenant. Cada unidade pode ter configurações próprias de regime tributário e catálogo de produtos. |
| **Plano (Plan)** | Pacote comercial contratado pelo cliente, definindo quais módulos estão disponíveis e as condições de cobrança. |
| **Assinatura (Subscription)** | Vínculo ativo entre um Tenant e um Plano, com vigência e status definidos. |
| **App Switcher** | Seletor de aplicativos/módulos no portal. Na versão atual (Core), exibido como menu de navegação entre módulos. Evolui para switcher visual quando houver 2+ produtos ativos. |
| **RBAC** | Controle de Acesso Baseado em Papéis (Role-Based Access Control) — define o que cada usuário pode ver e fazer com base no papel atribuído. MVP: 3 papéis (Admin, Gerente, Operador). |
| **Onboarding** | Fluxo guiado de primeiro acesso, onde o cliente configura sua conta, cadastra sua primeira Unidade de Negócio e aprende a usar o portal. |
| **Core** | A camada fundamental e compartilhada da plataforma — gerencia contas, planos, usuários e permissões para todos os módulos. |
| **Módulo / Produto** | Cada solução de negócio oferecida dentro da plataforma (ex: Tributali-Engine para gestão fiscal, Storekeeper Portal para varejo). |
| **Tributali-Engine** | Módulo futuro da plataforma focado em gestão tributária — cálculos de IBS/CBS, Split Payment e automação fiscal da Reforma Tributária. |
| **Storekeeper Portal** | Módulo futuro da plataforma focado em varejo — PDV (frente de caixa), controle de estoque e gestão comercial para lojistas e supermercados. |
| **Contrato de Interface** | Definição documentada de como o Core expõe serviços aos módulos (APIs, tabelas associativas, eventos). Congelado em M2 para isolar o Core de mudanças na visão de produto. |

---

> **Este documento é estritamente de negócio.** Detalhamentos técnicos, escolhas de arquitetura, modelagem de banco de dados e definições de infraestrutura serão tratados em documentos técnicos complementares (ADRs, especificações de engenharia) e não fazem parte do escopo deste Project Charter.

------------------------------
---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Revisão 1.1 baseada no Caveman Review (15/07/2026).*
