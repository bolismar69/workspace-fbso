# Business Requirements Document (BRD): PROJETO SHIELD — Plataforma de Identidade e Segurança
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documentos Base** | 001-PROJECT-CHARTER, 002-STAKEHOLDER-MAP, 003-PERSONAS-JORNADAS, 004-MAPEAMENTO-AS-IS-TO-BE |
| **Data de Elaboração** | 07/08/2026 |
| **Versão** | 1.5 — Aprovação humana (19/08/2026, P1=SIM/P2–P4=NÃO — atalho OK) — documento congelado em COMPLIANCE |
| **Metodologia** | WATERFALL |

---

## BRD — Business Requirements Document (Documento de Requisitos de Negócio)

O **BRD (Business Requirements Document)** é o documento formal que define os objetivos, o valor gerado e as necessidades estratégicas e operacionais de uma empresa para um novo projeto. Ele responde estritamente ao **"o quê"** a empresa precisa e ao **"porquê"**, servindo como o contrato primário de negócio antes de qualquer detalhamento funcional (**010-FRD**) ou especificação técnica (**020-SRS**).

### Principais Elementos do BRD

- **Resumo Executivo:** Visão geral da iniciativa, justificativa de negócio e problema a ser resolvido.
- **Objetivos de Negócio:** Metas quantificáveis e qualitativas que a empresa busca atingir (ex: redução de custos, conformidade regulatória, aumento de receita).
- **Escopo Declarado (In-Scope e Out-of-Scope):** Delimitação clara das fronteiras do projeto para evitar o crescimento desordenado do escopo (*scope creep*).
- **Requisitos de Negócio (`B-REQ-NN`):** Lista das necessidades de alto nível e regras de negócio macro que a solução deve atender para considerar o projeto um sucesso.
- **Cenário Atual vs. Futuro (As-Is / To-Be):** Descrição do processo operacional atual e do impacto esperado com a nova solução.
- **Premissas, Restrições e Riscos de Negócio:** Prazos limite, orçamento, políticas operacionais internas e riscos do negócio.
- **Matriz de Stakeholders:** Identificação dos patrocinadores (*Sponsors*), aprovadores do projeto e áreas impactadas.

### Para que serve no Pipeline Waterfall

- **Alinhamento Estratégico:** Garante que a diretoria, as áreas de negócio e a TI tenham exatamente a mesma expectativa sobre os entregáveis do projeto.
- **Insumo Direto para o FRD (010):** Serve como pai conceitual do **010-FRD**. Cada requisito do BRD (`B-REQ-NN`) dará origem a funcionalidades (`B-FEAT-NN`), regras de negócio (`B-RULE-NN`) e casos de uso (`B-UC-NN`).
- **Sustentação do Gate Upstream (ROM ±50%):** Fornece as justificativas financeiras e de escopo para que a TI elabore a arquitetura macro e submeta a primeira análise de viabilidade ao Comitê de Governança.
- **Âncora de Alterações (CCR):** Funciona como a primeira linha de defesa contra mudanças de escopo. Se uma nova demanda não apoiar nenhum objetivo do BRD, ela é sumariamente rejeitada ou exige revisão contratual.

### Siglas definidas no documento
- **B-REQ-** _(Business Constraint)_: Requisito de Negócios ambito Geral/Genérico (identificador padrão para qualquer `requisito de negócio` de ambito geral/macro).
- **B-RULE-** _(Business Rules)_: Regra de Negócio (restrições ou limitações que o negócio impõe).
- **B-LIMIT-** _(Business Constraint)_: Restrição de Negócio (limitações ou restrições de negócio, usado para rastrear, categorizar e referenciar limitações impostas pela empresa que afetam o escopo ou a entrega do projeto)
- **B-FEAT-** _(Funcionalidade de Negócio)_: funcionalidades e recursos de negócio, no nível macro
- **B-UC** _(Caso de Uso de Negócio)_: descreve um processo ou uma macro-atividade de negócio da organização, no nível macro, sob a ótica de valor para o negócio ou para o cliente fim do produto.
- **B-REQ-SECURITY-** _(Business Requirement — Security)_: Requisito de Segurança e Compliance no nível de negócio (obrigações regulatórias/LGPD), derivado dos B-REQ da Seção 1 (ver subseção 1.3).
- **B-PERSONA-** _(Business Persona)_: Persona de negócio — origem 003-PERSONAS-JORNADAS
- **B-JOURNEY-** _(Business Journey)_: Jornada de negócio — origem 003-PERSONAS-JORNADAS
- **B-PROCESS-** _(Business Process)_: Processo de negócio (AS-IS/TO-BE) — origem 004-MAPEAMENTO-AS-IS-TO-BE
- **B-GAP-ANALYSIS-** _(Business Gap Analysis)_: Gap AS-IS → TO-BE — origem 004-MAPEAMENTO-AS-IS-TO-BE

---

### 1. Requisitos de Negócio (Business Requeriments)

Cada requisito abaixo descreve **o que o produto deve fazer** para entregar valor ao negócio e aos clientes. O *como* será detalhado nos documentos técnicos (Fase 2 e Fase 3).

| ID | Necessidade de Negócio | Por que isso importa | Prioridade | Quem precisa disso |
|----|------------------------|---------------------|------------|-------------------|
| B-REQ-01 | A plataforma deve reconhecer automaticamente qual cliente está acessando (ex: `escola-alfa.com`) e direcioná-lo ao seu ambiente correto | Elimina configuração manual a cada novo cliente. Torna o onboarding escalável | **Alta** | Gerência Comercial |
| B-REQ-02 | A plataforma deve garantir que cada cliente opere em um ambiente totalmente isolado — a Escola A jamais pode acessar dados da Escola B | Protege a empresa de riscos de vazamento entre clientes (reputação + LGPD) | **Alta** | Diretoria de Tecnologia |
| B-REQ-03 | A plataforma deve proteger as credenciais dos usuários de forma que, mesmo que o navegador do usuário seja comprometido, as senhas e tokens não possam ser roubados | Mitiga os riscos mais comuns de ataque e fortalece a confiança dos clientes | **Alta** | Gerência de Tecnologia |
| B-REQ-04 | A plataforma deve oferecer funcionalidades padronizadas de autenticação — login, logout, perfil do usuário, renovação de sessão — que possam ser consumidas por qualquer produto do ecossistema FBSO | Acelera o lançamento de novos produtos e reduz custo de desenvolvimento | **Alta** | Product Owner |
| B-REQ-05 | A plataforma deve responder em menos de 15 milissegundos ao validar a identidade do usuário | Garante que a experiência de navegação é fluida. Latência perceptível causa abandono | **Média** | Gerência Comercial |
| B-REQ-06 | A plataforma deve suportar picos de acesso — como o horário de entrada simultânea de milhares de alunos e professores — sem apresentar falhas | As escolas operam em horários concentrados. O sistema não pode cair no momento de maior uso | **Média** | Gerência Comercial |
| B-REQ-07 | A plataforma deve registrar todas as tentativas de acesso (bem-sucedidas ou não) de forma rastreável, sem armazenar dados sensíveis como senhas | Base para auditoria LGPD, investigação de incidentes e conformidade regulatória | **Média** | Gerência de Tecnologia |
| B-REQ-08 | A plataforma deve permitir que um novo cliente (escola) seja ativado em até 4 horas — incluindo criação do ambiente isolado e configuração de domínio | Viabiliza o crescimento do negócio. Hoje o processo leva dias | **Média** | Gerência Comercial |
| B-REQ-09 | A plataforma deve ser capaz de se adaptar automaticamente ao aumento de demanda — se mais clientes entrarem ou o uso crescer, a plataforma escala sem intervenção manual | Reduz custo operacional e garante estabilidade em cenários de crescimento | **Baixa** | Gerência de Finanças |
| B-REQ-10 | A plataforma deve garantir que a experiência de login seja a mesma para todos os produtos do ecossistema FBSO | Consistência de marca e experiência do usuário. O cliente não percebe que está mudando de produto | **Média** | Gerência Comercial |
| B-REQ-11 | A transição dos sistemas atuais para a nova plataforma de acesso deve ocorrer sem interrupção perceptível para os usuários finais | Os sistemas das escolas já estão em operação. Uma parada ou falha durante a migração impacta diretamente a confiança dos clientes e a operação das instituições | **Alta** | Gerência Comercial, Gerência de Tecnologia |

#### 1.1 Requisitos de Dados

| Entidade | Descrição | Requisito Vinculado |
|----------|-----------|---------------------|
| Cliente (Escola) | Identificação única de cada escola, com ambiente isolado e domínio próprio | B-REQ-01, B-REQ-02, B-REQ-08 |
| Domínio | Endereço de acesso da escola — identificador único do ambiente | B-REQ-01 |
| Usuário | Perfil de acesso de professores, coordenadores e alunos por escola | B-REQ-03, B-REQ-04 |
| Sessão | Sessão de acesso protegida, com renovação transparente | B-REQ-04, B-REQ-05 |
| Registro de Auditoria | Registro de todas as tentativas de acesso (bem-sucedidas ou não), sem dados sensíveis | B-REQ-07 |

#### 1.2 Requisitos de Interface e Integração

| Interface | Descrição | Tipo | Requisito Vinculado |
|-----------|-----------|------|---------------------|
| Produtos do ecossistema FBSO | Todos os produtos da empresa consomem a plataforma única de acesso | Interface de negócio | B-REQ-04, B-REQ-10 |
| Fluxo de login/logout | Padronização de entrada, saída e renovação de sessão para todos os produtos | Interface de negócio | B-REQ-04, B-REQ-05 |
| Migração dos sistemas atuais | Cada sistema é integrado à nova plataforma individualmente, com plano de contingência | Interface de negócio | B-REQ-11 |

#### 1.3 Requisitos de Segurança e Compliance

| ID | Requisito | Regulação/Política | Requisito Vinculado |
|----|-----------|-------------------|---------------------|
| B-REQ-SECURITY-01 | Isolamento estrito entre ambientes de clientes | LGPD (segurança dos dados), política interna de segurança | B-REQ-02 |
| B-REQ-SECURITY-02 | Credenciais jamais expostas fora do ambiente seguro de autenticação | Política interna de segurança | B-REQ-03 |
| B-REQ-SECURITY-03 | Registro de auditoria de acessos sem dados sensíveis, com retenção mínima de 6 meses | LGPD (responsabilização e auditoria) | B-REQ-07 |
| B-REQ-SECURITY-04 | Conformidade com LGPD para dados processados em nuvem sem datacenter no Brasil | LGPD — diretrizes do Jurídico | B-LIMIT-04 |

---

### 2. Regras de Negócio (Business Rules)

| ID | Regra | Descrição | Vinculado a |
|----|-------|-----------|------------|
| B-RULE-01 | Identificação pelo Domínio | O domínio acessado pelo cliente (ex: `escola-alfa.com`) é o identificador único que determina em qual ambiente ele será autenticado. Um domínio não configurado recebe mensagem padronizada "Domínio não reconhecido" | B-REQ-01 |
| B-RULE-02 | Isolamento Estrito entre Clientes | Uma sessão iniciada no ambiente da Escola A não pode, sob nenhuma circunstância, acessar recursos ou dados da Escola B. Esta regra é a fundação da segurança da plataforma | B-REQ-02 |
| B-RULE-03 | Proteção Total das Credenciais | As credenciais de acesso dos usuários jamais são armazenadas ou transmitidas para fora do ambiente seguro de autenticação. O navegador do usuário nunca tem acesso direto a tokens ou senhas | B-REQ-03 |
| B-RULE-04 | Logout Completo | Quando um usuário sai do sistema, sua sessão é encerrada em todos os níveis — tanto na plataforma de identidade quanto no navegador | B-REQ-04 |
| B-RULE-05 | Renovação Transparente de Sessão | O usuário não deve precisar fazer login novamente durante o uso normal do sistema. A sessão é renovada automaticamente em segundo plano, desde que o usuário esteja ativo | B-REQ-04, B-REQ-05 |
| B-RULE-06 | Bloqueio Silencioso de Acesso Indevido | Tentativas de acessar dados de outro cliente não devem gerar mensagens de erro que revelem informações sobre a existência ou estrutura dos dados — simplesmente retornam "nada encontrado" | B-REQ-02 |
| B-RULE-07 | Provisionamento a Partir de Modelo Padrão | Novos clientes são criados a partir de um modelo pré-configurado que já contém as regras de segurança básicas. Customizações específicas são tratadas como solicitações separadas | B-REQ-08 |
| B-RULE-08 | Invalidação Imediata de Acesso | Quando um cliente é desativado ou suspenso, o acesso de todos os seus usuários deve ser bloqueado imediatamente, sem depender da expiração natural da sessão | B-REQ-02 |
| B-RULE-09 | Migração Gradual e Segura | Nenhum sistema em produção pode ser migrado para a nova plataforma sem que um plano de contingência e retorno ao estado anterior tenha sido aprovado pelo Comitê de Projeto. Cada sistema é migrado individualmente e validado antes do próximo | B-REQ-11 |

---

### 3. Restrições de Negócio (Business Constraint/Limitation)

| ID | Restrição | Descrição | Impacto no Negócio |
|----|-----------|-----------|-------------------|
| B-LIMIT-01 | Padrões Tecnológicos Corporativos | A empresa possui um catálogo de tecnologias aprovadas. Qualquer tecnologia fora desse catálogo precisa ser justificada e aprovada pelo comitê de arquitetura | Pode limitar escolhas técnicas, mas garante consistência e reduz custo de manutenção |
| B-LIMIT-02 | Provedor de Nuvem Exclusivo | A plataforma opera exclusivamente no provedor de nuvem corporativo aprovado. Serviços não disponíveis nesse provedor exigem soluções alternativas | Limita opções de infraestrutura, mas simplifica gestão e negociação |
| B-LIMIT-03 | Proteção de Borda Obrigatória | Todo tráfego externo passa por camada de segurança antes de chegar à plataforma | Adiciona etapa de configuração por cliente, mas é essencial para segurança |
| B-LIMIT-04 | LGPD — Dados em Nuvem | O provedor de nuvem corporativo aprovado não possui datacenter no Brasil. Isso exige medidas adicionais de conformidade com o Jurídico | Risco regulatório que precisa ser endereçado |
| B-LIMIT-05 | Prazo Máximo de 6 Semanas | O projeto precisa estar em produção em 6 semanas. Atrasos superiores a 20% exigem replanejamento formal | Pressão sobre escopo e qualidade |
| B-LIMIT-06 | Isolamento Lógico de Clientes | Cada cliente opera em ambiente logicamente isolado — como se fosse uma instalação dedicada — mas compartilha a mesma infraestrutura física | Modelo otimiza custo sem comprometer segurança |

---

### 4. Fluxos de Negócio

**Fluxo Principal — Acesso do Cliente:**

1. O usuário (professor, coordenador, aluno) acessa o endereço da sua escola no navegador
2. A plataforma identifica automaticamente qual é a escola pelo endereço acessado
3. Se a escola é reconhecida, o usuário é direcionado para a tela de login da sua instituição
4. O usuário informa suas credenciais
5. Após autenticação bem-sucedida, o usuário acessa o sistema com sua sessão protegida
6. Durante o uso, a sessão é renovada automaticamente — o usuário não percebe
7. Ao sair, a sessão é encerrada em todos os níveis

**Fluxo Alternativo — Escola não Reconhecida:**

1-2. Idêntico ao fluxo principal
3. O endereço acessado não corresponde a nenhum cliente ativo
4. O sistema exibe mensagem padronizada informando que o domínio não está configurado, sem revelar detalhes internos

**Fluxo de Ativação de Novo Cliente:**

1. A Gerência Comercial fecha contrato com uma nova escola
2. O Product Owner solicita a ativação da escola na plataforma
3. A equipe cria o ambiente isolado da nova escola a partir do modelo padrão (em até 4 horas)
4. A equipe configura o domínio da escola na camada de proteção
5. O Product Owner valida o fluxo completo de acesso
6. A escola é liberada para uso

**Fluxo de Bloqueio de Cliente:**

1. O contrato de uma escola é suspenso ou encerrado
2. A Gerência Comercial notifica o Product Owner
3. O ambiente da escola é marcado como suspenso
4. Todos os acessos ativos daquela escola são imediatamente bloqueados
5. Os dados permanecem armazenados (para eventual reativação ou extração), mas inacessíveis

---

### 5. Partes Interessadas e Suas Necessidades

> **📌 Documento completo:** `002-STAKEHOLDER-MAP-PRJ-TEC-2026-0004-PROJETO-SHIELD.md`

| Stakeholder | O que precisa | O que espera receber | Impacto no Projeto |
|-------------|--------------|---------------------|-------------------|
| Diretoria de Tecnologia | Plataforma única de identidade para todos os produtos; redução de custo com segurança | Relatórios quinzenais de progresso; zero incidentes de segurança após liberação | **Alta** — aprova orçamento e mudanças de escopo |
| Gerência Comercial | Produto que viabilize o crescimento do portfólio sem aumentar risco; onboarding rápido de novas escolas | Apresentação do produto antes do lançamento; demonstração de isolamento entre clientes | **Alta** — valida se o produto atende à estratégia da empresa |
| Gerência de Tecnologia | Arquitetura robusta e segura; conformidade com padrões corporativos | Equipe técnica ter clareza do escopo antes de iniciar codificação | **Alta** — responsável pela entrega |
| Gerência de Finanças | Budget controlado; sem surpresas de custo | Visibilidade de custos recorrentes antes do Go-Live | **Média** — libera pagamentos |
| Product Owner | Prioridades definidas; critérios de aceite claros; onboarding fluido de novos clientes | Plataforma funcional, documentação completa, processo de ativação definido | **Alta** — define o que entra e em qual ordem |
| PMO Corporativo | Cronograma realista; marcos mensuráveis | Alinhamento com portfólio corporativo; sem conflitos de recurso | **Média** — coordena com portfólio |
| Clientes (Escolas/Universidades) | Login rápido, seguro e sem complicações; garantia de isolamento de dados | Experiência de acesso fluida; zero vazamentos | **Alta** — razão de existir do produto |

---

### 6. Matriz de Rastreabilidade (BRD → Project Charter)

| Requisito de Negócio (BRD) | Critério de Sucesso (Charter) | Justificativa |
|----------------------------|------------------------------|----------------|
| B-REQ-01 — Reconhecimento automático do cliente | C1 — Segurança entre Clientes | Reconhecer o cliente evita direcionamento ao ambiente errado |
| B-REQ-02 — Isolamento total entre clientes | C1 — Segurança entre Clientes | Isolamento estrito = fundação da segurança |
| B-REQ-03 — Proteção de credenciais contra roubo | C2 — Proteção de Credenciais, C5 — Cobertura a Ataques Cibernéticos | Proteger credenciais combate diretamente os principais vetores de ataque |
| B-REQ-04 — Portal de acesso padronizado para todos os produtos | C6 — Tempo para Adicionar Novo Cliente | Portal padronizado acelera a integração de novos produtos e escolas |
| B-REQ-05 — Resposta em menos de 15ms | C3 — Velocidade de Resposta | 15ms é a meta de latência |
| B-REQ-06 — Suporte a picos de acesso sem falhas | C4 — Capacidade de Atender Picos | Picos de acesso = cenário de entrada simultânea |
| B-REQ-07 — Registro de acessos para auditoria | C8 — Rastreabilidade de Acessos | Registro de tentativas = base da rastreabilidade |
| B-REQ-08 — Ativação de novo cliente em até 4 horas | C6 — Tempo para Adicionar Novo Cliente | Meta de 4 horas para ativação |
| B-REQ-09 — Adaptação automática ao crescimento de demanda | C4 — Capacidade de Atender Picos, C7 — Disponibilidade da Plataforma | Escalar automaticamente garante capacidade em picos e disponibilidade contínua |
| B-REQ-10 — Experiência de login consistente em todos os produtos | C5 — Cobertura a Ataques Cibernéticos | Consistência entre produtos evita brechas de segurança por implementações divergentes |
| B-REQ-11 — Transição sem interrupção para os usuários finais | C7 — Disponibilidade da Plataforma, Premissas 1 e 2 (Charter Seção 7) | A migração não pode causar indisponibilidade ou impacto operacional nos clientes |

**Cobertura:** 11/11 requisitos de negócio vinculados a critérios de sucesso do Project Charter. 8/8 critérios de sucesso do Charter endereçados. Premissas do Charter cobertas. **Zero órfãos. 100% rastreável.**

---

## 7. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | 07/08/2026 | Criação inicial a partir do Project Charter (001) | Time de Negócios / Orquestrador WATERFALL |
| 1.1 | 19/08/2026 | Revisão de atualização: Documentos Base incluem 003/004; subseções 1.1–1.3 (Dados, Integração, Segurança) adicionadas; linguagem de negócio | Time de Negócios / skill waterfall-business-documents |
| 1.2 | 19/08/2026 | Correção cirúrgica (review FASE 1): marcador residual de status removido do rodapé — o status oficial permanece no cabeçalho | Time de Negócios / skill waterfall-business-documents |
| 1.3 | 19/08/2026 | Aprovação humana (P1=SIM, P2/P3/P4=NÃO — atalho OK) — documento congelado em COMPLIANCE | Orquestrador / skill waterfall-business-documents |
| 1.4 | 19/08/2026 | Correção cirúrgica (update pós-selo, F2/F3/F4/F6): B-REQ-11 passa a citar Premissas 1 e 2 (Seção 6); fluxos harmonizados ('Gerência Comercial' nos Fluxos de Ativação e Bloqueio); siglas B-FEAT/B-UC com glosa em português; campo Versão do cabeçalho alinhado | Time de Negócios / skill waterfall-business-documents |
| 1.5 | 19/08/2026 | Aprovação humana (P1=SIM, P2/P3/P4=NÃO — atalho OK) — correções do update pós-selo aprovadas; documento congelado em COMPLIANCE | Orquestrador / skill waterfall-business-documents |
