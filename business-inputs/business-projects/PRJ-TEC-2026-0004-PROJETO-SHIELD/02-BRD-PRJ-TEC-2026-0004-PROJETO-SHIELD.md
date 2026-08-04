# Business Requirements Document (BRD): PROJETO SHIELD — Plataforma de Identidade e Segurança
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documento Base** | 01-PROJECT-CHARTER-PRJ-TEC-2026-0004-PROJETO-SHIELD.md |
| **Data de Elaboração** | 03/08/2026 |
| **Versão** | 2.0 — Revisão para Apresentação à Diretoria de Negócios |
| **Metodologia** | WATERFALL |

---

### 1. Requisitos de Negócio

Cada requisito abaixo descreve **o que o produto deve fazer** para entregar valor ao negócio e aos clientes. O *como* será detalhado nos documentos técnicos (Fase 3).

| ID | Necessidade de Negócio | Por que isso importa | Prioridade | Quem precisa disso |
|----|------------------------|---------------------|------------|-------------------|
| REQ-01 | A plataforma deve reconhecer automaticamente qual cliente está acessando (ex: `escola-alfa.com`) e direcioná-lo ao seu ambiente correto | Elimina configuração manual a cada novo cliente. Torna o onboarding escalável | **Alta** | Comercial, Sucesso do Cliente |
| REQ-02 | A plataforma deve garantir que cada cliente opere em um ambiente totalmente isolado — a Escola A jamais pode acessar dados da Escola B | Protege a empresa de riscos de vazamento entre clientes (reputação + LGPD) | **Alta** | Jurídico/Compliance, Diretoria |
| REQ-03 | A plataforma deve proteger as credenciais dos usuários de forma que, mesmo que o navegador do usuário seja comprometido, as senhas e tokens não possam ser roubados | Mitiga os riscos mais comuns de ataque (XSS, CSRF) e fortalece a confiança dos clientes | **Alta** | Segurança da Informação, Clientes |
| REQ-04 | A plataforma deve oferecer funcionalidades padronizadas de autenticação — login, logout, perfil do usuário, renovação de sessão — que possam ser consumidas por qualquer produto do ecossistema FBSO | Acelera o lançamento de novos produtos e reduz custo de desenvolvimento | **Alta** | Product Owner, Times de Produto |
| REQ-05 | A plataforma deve responder em menos de 15 milissegundos ao validar a identidade do usuário | Garante que a experiência de navegação é fluida. Latência perceptível causa abandono | **Média** | Produto, Usuários Finais |
| REQ-06 | A plataforma deve suportar picos de acesso — como o horário de entrada simultânea de milhares de alunos e professores — sem apresentar falhas | As escolas operam em horários concentrados. O sistema não pode cair no momento de maior uso | **Média** | Operações, Clientes |
| REQ-07 | A plataforma deve registrar todas as tentativas de acesso (bem-sucedidas ou não) de forma rastreável, sem armazenar dados sensíveis como senhas | Base para auditoria LGPD, investigação de incidentes e conformidade regulatória | **Média** | Jurídico/Compliance, Segurança |
| REQ-08 | A plataforma deve permitir que um novo cliente (escola) seja ativado em até 4 horas — incluindo criação do ambiente isolado e configuração de domínio | Viabiliza o crescimento do negócio. Hoje o processo leva dias | **Média** | Comercial, Sucesso do Cliente |
| REQ-09 | A plataforma deve ser capaz de se adaptar automaticamente ao aumento de demanda — se mais clientes entrarem ou o uso crescer, a plataforma escala sem intervenção manual | Reduz custo operacional e garante estabilidade em cenários de crescimento | **Baixa** | Operações, Finanças |
| REQ-10 | A plataforma deve garantir que a experiência de login seja a mesma para todos os produtos do ecossistema FBSO | Consistência de marca e experiência do usuário. O cliente não percebe que está mudando de produto | **Média** | Produto, Marketing |

---

### 2. Regras de Negócio

| ID | Regra | Descrição | Vinculado a |
|----|-------|-----------|------------|
| BR-01 | Identificação pelo Domínio | O domínio acessado pelo cliente (ex: `escola-alfa.com`) é o identificador único que determina em qual ambiente ele será autenticado. Um domínio não configurado recebe mensagem padronizada "Domínio não reconhecido" | REQ-01 |
| BR-02 | Isolamento Estrito entre Clientes | Uma sessão iniciada no ambiente da Escola A não pode, sob nenhuma circunstância, acessar recursos ou dados da Escola B. Esta regra é a fundação da segurança da plataforma | REQ-02 |
| BR-03 | Proteção Total das Credenciais | As credenciais de acesso dos usuários jamais são armazenadas ou transmitidas para fora do ambiente seguro de autenticação. O navegador do usuário nunca tem acesso direto a tokens ou senhas | REQ-03 |
| BR-04 | Logout Completo | Quando um usuário sai do sistema, sua sessão é encerrada em todos os níveis — tanto na plataforma de identidade quanto no navegador. Se o servidor de identidade estiver indisponível, a sessão local ainda deve ser limpa | REQ-04 |
| BR-05 | Renovação Transparente de Sessão | O usuário não deve precisar fazer login novamente durante o uso normal do sistema. A sessão é renovada automaticamente em segundo plano, desde que o usuário esteja ativo | REQ-04, REQ-05 |
| BR-06 | Bloqueio Silencioso de Acesso Indevido | Tentativas de acessar dados de outro cliente não devem gerar mensagens de erro que revelem informações sobre a existência ou estrutura dos dados — simplesmente retornam "nada encontrado" | REQ-02 |
| BR-07 | Provisionamento a Partir de Modelo Padrão | Novos clientes são criados a partir de um modelo pré-configurado que já contém as regras de segurança básicas. Customizações específicas são tratadas como solicitações separadas | REQ-08 |
| BR-08 | Invalidação Imediata de Acesso | Quando um cliente é desativado ou suspenso, o acesso de todos os seus usuários deve ser bloqueado imediatamente, sem depender da expiração natural da sessão | REQ-02 |

---

### 3. Restrições de Negócio

| ID | Restrição | Descrição | Impacto no Negócio |
|----|-----------|-----------|-------------------|
| BC-01 | Padrões Tecnológicos Corporativos | A empresa possui um catálogo de tecnologias aprovadas. Qualquer tecnologia fora desse catálogo precisa ser justificada e aprovada pelo comitê de arquitetura | Pode limitar escolhas técnicas, mas garante consistência e reduz custo de manutenção |
| BC-02 | Provedor de Nuvem Exclusivo | A plataforma opera exclusivamente na DigitalOcean. Serviços não disponíveis nesse provedor exigem soluções alternativas | Limita opções de infraestrutura, mas simplifica gestão e negociação |
| BC-03 | Proteção de Borda Obrigatória | Todo tráfego externo passa por camada de segurança (Cloudflare) antes de chegar à plataforma | Adiciona etapa de configuração por cliente, mas é essencial para segurança |
| BC-04 | LGPD — Dados em Nuvem | A DigitalOcean não possui datacenter no Brasil. Isso exige medidas adicionais de conformidade (termo de tratamento, criptografia) | Risco regulatório que precisa ser endereçado com o Jurídico |
| BC-05 | Prazo Máximo de 6 Semanas | O projeto precisa estar em produção em 6 semanas. Atrasos superiores a 20% exigem replanejamento formal | Pressão sobre escopo e qualidade |
| BC-06 | Isolamento Físico Lógico | Cada cliente opera em ambiente logicamente isolado — como se fosse uma instalação dedicada — mas compartilha a mesma infraestrutura física | Modelo otimiza custo sem comprometer segurança |

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

1. O time comercial fecha contrato com uma nova escola
2. O Product Owner solicita a ativação da escola na plataforma
3. A equipe cria o ambiente isolado da nova escola a partir do modelo padrão (em até 4 horas)
4. A equipe configura o domínio da escola na camada de proteção
5. O Product Owner valida o fluxo completo de acesso
6. A escola é liberada para uso

**Fluxo de Bloqueio de Cliente:**

1. O contrato de uma escola é suspenso ou encerrado
2. O time de Sucesso do Cliente notifica o Product Owner
3. O ambiente da escola é marcado como suspenso
4. Todos os acessos ativos daquela escola são imediatamente bloqueados
5. Os dados permanecem armazenados (para eventual reativação ou extração), mas inacessíveis

---

### 5. Partes Interessadas e Suas Necessidades

| Stakeholder | O que precisa | O que espera receber | Impacto no Projeto |
|-------------|--------------|---------------------|-------------------|
| Diretoria de Tecnologia | Plataforma única de identidade para todos os produtos; redução de custo com segurança | Relatórios quinzenais de progresso; zero incidentes de segurança após liberação | **Alta** — aprova orçamento e mudanças de escopo |
| Diretoria de Negócios | Produto que viabilize o crescimento do portfólio sem aumentar risco | Apresentação do produto antes do lançamento; demonstração de isolamento entre clientes | **Alta** — valida se o produto atende à estratégia da empresa |
| Product Owner | Backlog priorizado; critérios de aceite claros; onboarding fluido de novos clientes | Plataforma funcional, documentação completa, processo de ativação definido | **Alta** — define o que entra e em qual ordem |
| Times de Produto (consumidores) | Contrato claro de integração; documentação; ambiente de homologação | Portal de acesso padronizado; tempo de integração previsível | **Média** — consumirão a plataforma |
| Clientes (Escolas/Universidades) | Login rápido, seguro e sem complicações; garantia de isolamento de dados | Experiência de acesso fluida; zero vazamentos | **Alta** — razão de existir do produto |
| Jurídico/Compliance | Conformidade LGPD; rastreabilidade de acessos; isolamento de dados | Registros de auditoria; garantias de proteção de dados pessoais | **Alta** — aprova aspectos regulatórios |
| Segurança da Informação | Proteção contra os principais vetores de ataque; zero credenciais expostas | Relatório de testes de segurança; checklist de proteção aplicado | **Média** — valida a robustez da solução |

---

### 6. Matriz de Rastreabilidade (BRD → Project Charter)

| Requisito de Negócio (BRD) | Objetivo do Projeto (Charter) | Rastreabilidade |
|----------------------------|------------------------------|----------------|
| REQ-01 — Reconhecimento automático do cliente | C1 — Segurança entre Clientes, C6 — Tempo de Ativação | ✅ Vinculado |
| REQ-02 — Isolamento total entre clientes | C1 — Segurança entre Clientes | ✅ Vinculado |
| REQ-03 — Proteção de credenciais contra roubo | C2 — Proteção de Credenciais | ✅ Vinculado |
| REQ-04 — Portal de acesso padronizado para todos os produtos | C3 — Velocidade de Resposta | ✅ Vinculado |
| REQ-05 — Resposta em menos de 15ms | C3 — Velocidade de Resposta | ✅ Vinculado |
| REQ-06 — Suporte a picos de acesso sem falhas | C4 — Capacidade de Atender Picos | ✅ Vinculado |
| REQ-07 — Registro de acessos para auditoria | C8 — Rastreabilidade de Acessos | ✅ Vinculado |
| REQ-08 — Ativação de novo cliente em até 4 horas | C6 — Tempo de Ativação | ✅ Vinculado |
| REQ-09 — Adaptação automática ao crescimento de demanda | C4 — Capacidade de Atender Picos, C7 — Disponibilidade | ✅ Vinculado |
| REQ-10 — Experiência de login consistente em todos os produtos | C1 a C8 — Qualidade geral da plataforma | ✅ Vinculado |

**Cobertura:** 10/10 requisitos de negócio vinculados a critérios de sucesso do Project Charter. 8/8 critérios de sucesso do Charter endereçados por pelo menos um requisito. **Zero órfãos. 100% rastreável.**

---

**[STATUS: SUCESSO]** — Documento revisado para linguagem de negócio. Detalhes técnicos reservados para SRS, SAD, HLD e LLD (Fases 2 e 3).
