# Personas e Jornadas de Negócio: PROJETO SHIELD — Plataforma de Identidade e Segurança
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documentos Base** | 001-PROJECT-CHARTER, 002-STAKEHOLDER-MAP |
| **Data de Elaboração** | 19/08/2026 |
| **Versão** | 1.5 — Aprovação humana (19/08/2026, P1=SIM/P2–P4=NÃO — atalho OK) — documento congelado em COMPLIANCE |
| **Metodologia** | WATERFALL |

---

### Siglas definidas no documento

- **B-PERSONA-** _(Business Persona)_: Persona de negócio (prefixo de identificação P→B-PERSONA)
- **B-JOURNEY-** _(Business Journey)_: Jornada de negócio (prefixo de identificação J→B-JOURNEY)

---

## Personas e Jornadas de Negócio

O **documento de Personas e Jornadas** formaliza QUEM usa o produto (personas) e COMO cada perfil percorre o negócio hoje (jornadas). É o primeiro documento derivado do Charter e do Stakeholder Map e fundamenta os requisitos de negócio (005-BRD) e os casos de uso (010-FRD) com evidência de usuário.

### O que contém

- **Personas (B-PERSONA-NN):** perfis de usuário com objetivos, dores, contexto e nível de influência — sempre derivados dos stakeholders documentados (002-STAKEHOLDER-MAP ou 005-BRD Seção 5)
- **Jornadas (B-JOURNEY-NN):** caminho de cada persona por etapas, ações, pontos de contato, dores e oportunidades
- **Matriz Persona × Jornada × Ponto de Contato:** cruzamento que evidencia quais perfis são atendidos em cada interação
- **Rastreabilidade:** cada persona/jornada aponta o stakeholder de origem (002) e o objetivo de negócio (001)

### Conexão com o Pipeline

- **UPSTREAM:** Consome objetivos de negócio do 001-PROJECT-CHARTER e partes interessadas do 002-STAKEHOLDER-MAP
- **DOWNSTREAM:** Alimenta 004-MAPEAMENTO-AS-IS-TO-BE (processos por perfil), 005-BRD (requisitos fundamentados em usuário), 010-FRD (casos de uso por persona), 016-PROTOTIPOS-UX-UI (design por persona), 088-PRODUCT-BACKLOG-LIST e 097-MANUAIS-USUARIO

---

## 1. Personas (B-PERSONA-NN)

| ID | Nome | Perfil/Função | Objetivos | Dores | Contexto de Uso | Stakeholder de Origem (005-BRD Seção 5) |
|----|------|---------------|-----------|-------|-----------------|------------------------------|
| B-PERSONA-01 | Diretora Escolar | Diretora de escola cliente | Acessar os sistemas da sua escola rapidamente; garantir que alunos e professores entrem sem dificuldade; saber que os dados da escola estão isolados | Processo de acesso fragmentado entre produtos; medo de vazamento de dados entre escolas; dependência de configuração manual | Início do dia letivo e picos (entrada/saída de aulas); usa de qualquer navegador | Clientes (Escolas/Universidades) — 005-BRD Seção 5 |
| B-PERSONA-02 | Professora/Coordenadora | Professora ou coordenadora pedagógica | Entrar no sistema da escola em segundos e continuar de onde parou; não repetir login durante o dia | Senhas por produto; login repetido em horários de pico; medo de conta bloqueada por erro alheio | Sala de aula, sala dos professores, casa (horários alternativos) | Clientes (Escolas/Universidades) — 005-BRD Seção 5 |
| B-PERSONA-03 | Aluno | Estudante da escola cliente | Acessar o conteúdo da sua turma sem barreiras | Esquecer senha; não entender mensagens de erro; lentidão na entrada simultânea | Horário de entrada das aulas (pico simultâneo) | Clientes (Escolas/Universidades) — 005-BRD Seção 5 |
| B-PERSONA-04 | Especialista de Integração de Sistemas | Papel operacional interno responsável pela migração dos sistemas atuais para a nova Plataforma Única de Acesso | Garantir a equivalência entre os sistemas atuais e a nova Plataforma Única de Acesso — acessos e permissões equivalentes para todos os usuários, sem interrupção perceptível durante a transição; preservar a continuidade das jornadas existentes (B-JOURNEY-01/02/03) | Ausência de processo padronizado de migração; risco de interrupção para usuários ativos; divergência de acessos/permissões entre o sistema atual e a nova plataforma | Operação interna — janelas de baixo uso, aprovações do Comitê de Projeto, execução de contingência com retorno ao estado anterior | Gerência de Tecnologia — 002 Seção 1.2 / 005-BRD Seção 5 |

> **Legenda:** B-PERSONA-01 a B-PERSONA-03 são personas usuárias (uso final da plataforma). B-PERSONA-04 é uma persona executora interna, derivada da Gerência de Tecnologia — sua missão é garantir a continuidade das jornadas existentes durante a transição.

> **REGRA:** Toda persona deve derivar de pelo menos um stakeholder documentado no 002-STAKEHOLDER-MAP ou no registro de partes interessadas do 005-BRD (Seção 5). Personas sem origem documentada são proibidas (gold-plating).

---

## 2. Jornadas de Negócio (B-JOURNEY-NN)

Cada jornada descreve o caminho de uma persona por um fluxo de negócio relevante ao projeto.

### B-JOURNEY-01: Acesso Diário aos Sistemas da Escola

| Campo | Detalhe |
|-------|---------|
| **ID** | B-JOURNEY-01 |
| **Persona** | B-PERSONA-01, B-PERSONA-02, B-PERSONA-03 |
| **Objetivo da Jornada** | Autenticar e acessar os sistemas da escola sem atrito, com sessão segura e sem repetição de login |
| **Objetivo de Negócio Relacionado (001)** | C1 (Segurança entre Clientes), C2 (Proteção de Credenciais), C3 (Velocidade), C4 (Picos) |

**Etapas:**

| Etapa | Ação do Usuário | Ponto de Contato | Dor | Oportunidade |
|-------|-----------------|------------------|-----|--------------|
| 1. Chegada | Acessa o endereço da escola no navegador | Página de entrada da escola | Produtos diferentes exigem endereços diferentes | Reconhecimento automático pelo domínio → candidato a REQ (B-REQ-01) |
| 2. Identificação | Informa credenciais | Tela de login | Repetir login várias vezes ao dia; medo de senha roubada | Sessão única protegida e renovada automaticamente → B-REQ-03/04/05 |
| 3. Uso | Navega pelos sistemas da escola | Sistemas da instituição | Lentidão no pico de entrada | Validação <15ms mesmo em pico → B-REQ-05/06 |
| 4. Saída | Encerra a sessão | Logout | Dúvida se a sessão foi realmente encerrada | Logout completo em todos os níveis → B-RULE-04 |

### B-JOURNEY-02: Ativação de uma Nova Escola na Plataforma

| Campo | Detalhe |
|-------|---------|
| **ID** | B-JOURNEY-02 |
| **Persona** | B-PERSONA-01 |
| **Objetivo da Jornada** | Ter a escola pronta para uso com ambiente isolado e domínio configurado, sem espera de dias |
| **Objetivo de Negócio Relacionado (001)** | C6 (Tempo para Adicionar Novo Cliente) |

**Etapas:**

| Etapa | Ação do Usuário | Ponto de Contato | Dor | Oportunidade |
|-------|-----------------|------------------|-----|--------------|
| 1. Contratação | Escola fecha contrato com a FBSO | Gerência Comercial | Nenhuma visibilidade do prazo de ativação | Processo padronizado com SLA de 4 horas → B-REQ-08 |
| 2. Configuração | Equipe FBSO cria o ambiente da escola | Ambiente isolado da escola | Dias de configuração manual | Ativação a partir de modelo padrão → B-RULE-07 |
| 3. Primeiro acesso | Diretora testa o acesso | Tela de login da escola | Configuração errada descoberta só no primeiro uso | Validação completa do fluxo antes da liberação → B-UC-02 |

### B-JOURNEY-03: Suspensão de Acesso de uma Escola

| Campo | Detalhe |
|-------|---------|
| **ID** | B-JOURNEY-03 |
| **Persona** | B-PERSONA-01 |
| **Objetivo da Jornada** | Garantir que, ao encerrar o contrato, nenhum acesso da escola permaneça ativo |
| **Objetivo de Negócio Relacionado (001)** | C1 (Segurança entre Clientes), C8 (Rastreabilidade de Acessos) |

**Etapas:**

| Etapa | Ação do Usuário | Ponto de Contato | Dor | Oportunidade |
|-------|-----------------|------------------|-----|--------------|
| 1. Encerramento | Contrato é suspenso ou encerrado | Gerência Comercial → PO | Bloqueio dependia de processo manual e demorado | Bloqueio imediato de todos os acessos → B-RULE-08/B-RULE-13 |
| 2. Verificação | PO confirma o bloqueio | Painel de gestão | Sem evidência de que todos os acessos caíram | Registro de auditoria de todas as tentativas → B-REQ-07 |

### B-JOURNEY-04: Migração dos Sistemas Atuais para a Plataforma Única de Acesso

| Campo | Detalhe |
|-------|---------|
| **ID** | B-JOURNEY-04 |
| **Persona** | B-PERSONA-04 |
| **Objetivo da Jornada** | Migrar cada sistema atual para a nova plataforma de acesso sem interrupção perceptível para os usuários finais, com equivalência de acessos e permissões |
| **Objetivo de Negócio Relacionado (001)** | C7 (Disponibilidade da Plataforma), Premissas 1 e 2 (Charter Seção 7) |

**Etapas:**

| Etapa | Ação do Responsável | Ponto de Contato | Dor | Oportunidade |
|-------|---------------------|------------------|-----|--------------|
| 1. Preparação | Comitê de Projeto aprova o plano de contingência do sistema alvo | Comitê de Projeto | Migração sem processo padronizado; risco de impacto em usuários ativos | Processo individual por sistema com contingência aprovada → B-REQ-11 |
| 2. Execução | Equipe executa a migração em janela de baixo uso | Janela de migração fora do horário comercial | Interrupção perceptível se a janela falhar | Migração individual por sistema sem afetar os demais → B-RULE-23 |
| 3. Validação | Product Owner valida o funcionamento com a nova plataforma | Sistema migrado em operação | Divergência de acessos/permissões entre sistemas | Equivalência de acessos e permissões validada antes da confirmação → B-REQ-11 |
| 4. Contingência | Em caso de falha, retorno ao estado anterior em até 30 minutos | Plano de contingência | Sem plano, usuários finais percebem a interrupção | Rollback em até 30 minutos e nova janela agendada → B-RULE-24 |

> **REGRA:** Cada persona identificada na Seção 1 deve ter pelo menos uma jornada. As oportunidades de cada etapa são candidatas a requisitos de negócio (REQ-NN) e serão validadas no 005-BRD.

---

## 3. Matriz Persona × Jornada × Ponto de Contato

| Persona | Jornada | Pontos de Contato | Etapas Cobertas |
|---------|---------|-------------------|-----------------|
| B-PERSONA-01 | B-JOURNEY-01 | Página de entrada, tela de login, sistemas da escola | 1, 2, 3, 4 |
| B-PERSONA-02 | B-JOURNEY-01 | Página de entrada, tela de login, sistemas da escola | 1, 2, 3, 4 |
| B-PERSONA-03 | B-JOURNEY-01 | Página de entrada, tela de login, sistemas da escola | 1, 2, 3, 4 |
| B-PERSONA-01 | B-JOURNEY-02 | Gerência Comercial, ambiente isolado, tela de login | 1, 2, 3 |
| B-PERSONA-01 | B-JOURNEY-03 | Gerência Comercial, painel de gestão | 1, 2 |
| B-PERSONA-04 | B-JOURNEY-04 | Comitê de Projeto, janela de migração, sistema migrado, plano de contingência | 1, 2, 3, 4 |

---

## 4. Rastreabilidade

| Item | Origem (001/002) | Consumidores Previstos | Status |
|------|------------------|------------------------|--------|
| B-PERSONA-01 | Clientes (Escolas/Universidades) — 005-BRD Seção 5 | 004, 005, 010, 016 | ✅ Vinculado |
| B-PERSONA-02 | Clientes (Escolas/Universidades) — 005-BRD Seção 5 | 004, 005, 010, 016 | ✅ Vinculado |
| B-PERSONA-03 | Clientes (Escolas/Universidades) — 005-BRD Seção 5 | 004, 005, 010, 016 | ✅ Vinculado |
| B-PERSONA-04 | Gerência de Tecnologia — 002 Seção 1.2 / 005-BRD Seção 5 | 004, 005, 010, 016 | ✅ Vinculado |
| B-JOURNEY-01 | C1–C4 (001 Seção 6) | 004, 005, 010 | ✅ Vinculado |
| B-JOURNEY-02 | C6 (001 Seção 6) | 004, 005, 010 | ✅ Vinculado |
| B-JOURNEY-03 | C1, C8 (001 Seção 6) | 004, 005, 010 | ✅ Vinculado |
| B-JOURNEY-04 | C7, Premissas 1 e 2 (001 Seções 6/7) | 004, 005, 010 | ✅ Vinculado |

> **REGRA DE OURO:** Nenhuma persona ou jornada pode existir sem lastro no Charter (001), no Stakeholder Map (002) ou no registro de partes interessadas do BRD (005-BRD Seção 5). A RTM-FASE-1 (015) validará esta rastreabilidade formalmente.

---

## 5. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | 19/08/2026 | Criação inicial a partir do Charter e Stakeholder Map | Time de Negócios / skill waterfall-business-documents |
| 1.1 | 19/08/2026 | Correção cirúrgica (review FASE 1): stakeholder de origem das personas corrigido — "Clientes (Escolas/Universidades)" consta no registro de partes interessadas do 005-BRD (Seção 5), não no 002 | Time de Negócios / skill waterfall-business-documents |
| 1.2 | 19/08/2026 | Correção cirúrgica (review FASE 1, F1): B-PERSONA-04 (Especialista de Integração de Sistemas) e B-JOURNEY-04 (Migração dos Sistemas Atuais) adicionadas — cadeia de origem do B-REQ-11 completa | Time de Negócios / skill waterfall-business-documents |
| 1.3 | 19/08/2026 | Aprovação humana (P1=SIM, P2/P3/P4=NÃO — atalho OK) — documento congelado em COMPLIANCE | Orquestrador / skill waterfall-business-documents |
| 1.4 | 19/08/2026 | Correção cirúrgica (update pós-selo, F2/F3/F4/F6): B-REQ-11 passa a citar Premissas 1 e 2 (J-04 e Seção 4); pontos de contato e matriz harmonizados ('Gerência Comercial'); título com subtítulo do produto; campo Versão do cabeçalho alinhado | Time de Negócios / skill waterfall-business-documents |
| 1.5 | 19/08/2026 | Aprovação humana (P1=SIM, P2/P3/P4=NÃO — atalho OK) — correções do update pós-selo aprovadas; documento congelado em COMPLIANCE | Orquestrador / skill waterfall-business-documents |
