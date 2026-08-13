# Software Requirements Specification (SRS): PROJETO SHIELD — Plataforma de Identidade e Segurança
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documentos Base** | 001-PROJECT-CHARTER, 005-BRD, 010-FRD, 015-RTM-FASE-1 |
| **Data de Elaboração** | 08/08/2026 |
| **Versão** | 1.0 |
| **Metodologia** | WATERFALL |

---

## SRS — Software Requirements Specification (Especificação de Requisitos do Sistema)

O **SRS (Software Requirements Specification)** é o documento técnico oficial que traduz as necessidades de negócio do **005-BRD** e as funcionalidades do **010-FRD** em especificações operacionais e computacionais do sistema. Enquanto o FRD define o que o usuário vê e experimenta (funcionalidades, casos de uso, fluxos), o SRS define como o sistema deve se comportar internamente para sustentar essas funcionalidades.

### O Papel do SRS no Pipeline Waterfall

- **Foco no Sistema, não no Usuário:** Detalha comportamentos do sistema, processamento, regras de validação e Requisitos Não-Funcionais (NFRs)
- **Guia para Arquitetos e Engenheiros:** Serve como insumo direto para o desenho da arquitetura macro (**030-SAD**) e detalhada (**040-LLD**)
- **Base para Testes Não-Funcionais:** Orienta o QA na criação de suítes de carga, estresse, resiliência e segurança (**045-EST-PLAN**, **050-EST-CASES**)
- **Ponte para o Gate Upstream (ROM ±50%):** Consolida os NFRs necessários para estimativa macro de viabilidade

### O que contém o SRS

1. **Requisitos Funcionais de Sistema (FR-XX):** Tradução das funcionalidades de negócio (B-FEAT) em comportamentos do sistema
2. **Requisitos Não-Funcionais (NFRs):** Desempenho, Segurança, Disponibilidade, Escalabilidade, Observabilidade — com métricas mensuráveis
3. **Restrições Técnicas:** Diretrizes mandatórias de plataforma e limites operacionais

### Siglas definidas neste documento

| Prefixo | Significado | Origem |
|---------|-------------|--------|
| `B-REQ-` | Requisito de Negócio | 005-BRD |
| `B-RULE-` | Regra de Negócio | 005-BRD |
| `B-LIMIT-` | Restrição de Negócio | 005-BRD |
| `B-FEAT-` | Funcionalidade de Negócio | Este documento (010-FRD) |
| `FUNCTIONAL-REQ-` | Requisitos Funcionais de Sistema (System Functional Requirements) | Este documento (020-SRS) |
| `NO-FUNCTIONAL-REQ-{category}-` | Requisitos Não-Funcionais (Non-Functional Requirements) | Este documento (020-SRS) |
| `TECH-LIMIT-` | Restrições Técnicas de Sistema | Este documento (020-SRS) |

> categorias (`category`) para `NO-FUNCTIONAL-REQ-`:
- `PERFORMANCE`: Performance
- `SECURITY`: Segurança
- `AVAILABLE`: Disponibilidade
- `SCALABILITY`: Escalabilidade
- `OBSERVABILITY`: Observabilidade
- `USABILITY`: Usabilidade

---

### 1. Requisitos Funcionais de Sistema (System Functional Requirements)

Cada FR traduz uma ou mais funcionalidades de negócio (B-FEAT) do FRD em especificações de comportamento do sistema.

| ID | Requisito Funcional | Origem FRD (B-FEAT) | Critério de Aceitação Técnico |
|----|--------------------|---------------------|------------------------------|
| FUNCTIONAL-REQ-01 | O sistema deve extrair o domínio de origem do cliente a partir do cabeçalho da requisição de entrada e utilizá-lo como chave de identificação do ambiente isolado | B-FEAT-01 | Dado um domínio mapeado, o sistema direciona para o ambiente correto; dado um domínio não mapeado, retorna resposta padronizada sem revelar detalhes internos |
| FUNCTIONAL-REQ-02 | O sistema deve manter cache de mapeamento entre domínios e ambientes isolados, com invalidação imediata sob demanda e tempo de vida configurável | B-FEAT-01, B-FEAT-08 | Consulta ao cache responde em < 5ms (p99); atualização de mapeamento se propaga em < 1s |
| FUNCTIONAL-REQ-03 | O sistema deve iniciar o fluxo de autenticação do usuário contra o ambiente isolado correspondente, utilizando protocolo de autorização seguro com desafio criptográfico (PKCE) | B-FEAT-03, B-FEAT-04 | Usuário é direcionado à tela de login do ambiente correto; código de desafio gerado e validado a cada fluxo |
| FUNCTIONAL-REQ-04 | O sistema deve receber o código de autorização, trocá-lo por tokens de acesso e armazenar esses tokens exclusivamente em cookies protegidos no navegador | B-FEAT-03, B-FEAT-04 | Tokens nunca aparecem no corpo da resposta HTTP ou em headers acessíveis via JavaScript; cookies possuem flags HttpOnly, Secure e SameSite=Strict |
| FUNCTIONAL-REQ-05 | O sistema deve fornecer endpoint para consulta do perfil do usuário autenticado, retornando identificação, papéis e permissões sem expor tokens | B-FEAT-04, B-FEAT-05 | Endpoint de perfil responde em < 15ms (p95); resposta não contém tokens nem dados de outros clientes |
| FUNCTIONAL-REQ-06 | O sistema deve permitir a renovação silenciosa da sessão do usuário utilizando token de atualização, sem exigir nova interação de login | B-FEAT-04, B-RULE-16 | Sessão renovada automaticamente enquanto usuário estiver ativo; após expiração do token de atualização, redireciona ao login |
| FUNCTIONAL-REQ-07 | O sistema deve encerrar completamente a sessão do usuário — invalidando-a no ambiente de identidade e removendo os cookies do navegador | B-FEAT-04, B-RULE-15 | Após logout, cookies são removidos e a sessão no ambiente de identidade é invalidada; falha no ambiente remoto não impede limpeza local |
| FUNCTIONAL-REQ-08 | O sistema deve garantir que consultas a dados sejam automaticamente filtradas pelo identificador do cliente corrente, impedindo acesso a dados de outros clientes | B-FEAT-02, B-RULE-12 | Consulta executada com identificador do cliente A retorna zero resultados ao buscar dados do cliente B — sem erro explícito que vaze informação |
| FUNCTIONAL-REQ-09 | O sistema deve registrar todos os eventos de autenticação (login, logout, renovação, falha) com correlation_id, tenant_id e carimbo de tempo, sem incluir dados sensíveis | B-FEAT-07, B-RULE-20 | Cada evento gera registro com correlation_id e tenant_id; logs não contêm senhas, tokens ou dados pessoais |
| FUNCTIONAL-REQ-10 | O sistema deve expor métricas operacionais (taxa de requisições, latência, taxa de erros, sessões ativas por cliente) em formato compatível com ferramentas de monitoramento | B-FEAT-09 | Métricas disponíveis para coleta via endpoint padronizado; dashboards populados automaticamente |
| FUNCTIONAL-REQ-11 | O sistema deve permitir a ativação de um novo cliente mediante criação do ambiente isolado a partir de modelo pré-configurado e associação do domínio | B-FEAT-08, B-RULE-21 | Novo cliente provisionado em < 4 horas a partir da solicitação; ambiente herda todas as configurações de segurança do modelo padrão |
| FUNCTIONAL-REQ-12 | O sistema deve permitir a suspensão imediata de todos os acessos de um cliente específico, independentemente do estado das sessões ativas | B-FEAT-02, B-RULE-13 | Após marcação do cliente como suspenso, qualquer requisição com esse tenant_id é rejeitada em < 1s |
| FUNCTIONAL-REQ-13 | O sistema deve monitorar a latência de validação de identidade e emitir alerta quando a latência p95 ultrapassar 20ms por mais de 2 minutos consecutivos | B-FEAT-05, B-FEAT-06, B-RULE-18 | Alerta emitido automaticamente; serviços não críticos podem ser degradados antes que o login seja afetado |
| FUNCTIONAL-REQ-14 | O sistema deve suportar degradação controlada: se a demanda atingir 90% da capacidade máxima, serviços não essenciais são reduzidos para preservar a função de login | B-FEAT-06, B-RULE-18 | Login permanece operacional; redução de serviços não críticos é registrada em log |
| FUNCTIONAL-REQ-15 | O sistema deve reter registros de auditoria de acesso por no mínimo 6 meses, com política de expurgo automático após esse período | B-FEAT-07, B-RULE-19 | Registros com idade > 6 meses são expurgados conforme política configurada |
| FUNCTIONAL-REQ-16 | O sistema deve permitir a migração individual de cada sistema corporativo para a nova plataforma, com capacidade de reversão (rollback) em até 30 minutos em caso de falha | B-FEAT-10, B-RULE-24 | Rollback concluído em < 30min; sistema retorna ao estado anterior sem perda de dados |
| FUNCTIONAL-REQ-17 | O sistema deve fornecer relatório de acessos filtrável por cliente e período, para fins de investigação de incidentes e auditoria de conformidade | B-FEAT-07, B-RULE-19 | Relatório gerado em < 60s para períodos de até 30 dias; formato inclui data, hora, resultado e correlation_id |

---

### 2. Requisitos Não-Funcionais (Non-Functional Requirements)

Cada NFR estabelece uma meta mensurável e verificável que o sistema deve atender.

| ID | Requisito | Métrica Mensurável | Origem (BRD/Charter) |
|----|-----------|--------------------|----------------------|
| NO-FUNCTIONAL-REQ-PERFORMANCE-01 | Latência de validação de identidade (consulta de perfil) | p95 < 15ms | B-REQ-05, C3 |
| NO-FUNCTIONAL-REQ-PERFORMANCE-02 | Latência de consulta ao cache de mapeamento domínio→ambiente | p99 < 5ms | B-REQ-01 |
| NO-FUNCTIONAL-REQ-PERFORMANCE-03 | Tempo de inicialização de uma nova instância do serviço (cold start) | < 100ms | B-REQ-09 |
| NO-FUNCTIONAL-REQ-SECURITY-01 | Cookies de sessão protegidos contra acesso via JavaScript | 100% dos cookies de autenticação com flag HttpOnly | B-REQ-03, C2 |
| NO-FUNCTIONAL-REQ-SECURITY-02 | Cookies de sessão com restrição de transporte | 100% com flag Secure (HTTPS apenas) | B-REQ-03, C2 |
| NO-FUNCTIONAL-REQ-SECURITY-03 | Cookies de sessão com proteção contra CSRF | 100% com flag SameSite=Strict | B-REQ-03, C5 |
| NO-FUNCTIONAL-REQ-SECURITY-04 | Criptografia de dados em trânsito | TLS 1.3 como versão mínima; cifras fracas desabilitadas | C5 |
| NO-FUNCTIONAL-REQ-SECURITY-05 | Sanitização de logs — zero dados sensíveis | 100% dos logs sem senhas, tokens, CPF, e-mail completo ou PII | B-REQ-07, B-RULE-20 |
| NO-FUNCTIONAL-REQ-SECURITY-06 | Validação de segurança automatizada no pipeline | Scan OWASP ZAP a cada build; vulnerabilidades HIGH bloqueiam deploy | C5 |
| NO-FUNCTIONAL-REQ-AVAILABLE-01 | SLA de uptime do serviço de autenticação | 99.9% (máximo 8.76h de indisponibilidade/ano, excluindo janelas programadas) | C7 |
| NO-FUNCTIONAL-REQ-AVAILABLE-02 | Tempo de recuperação em caso de falha (RTO) | < 5 minutos | C7 |
| NO-FUNCTIONAL-REQ-AVAILABLE-03 | Perda máxima de dados em caso de falha (RPO) | < 1 minuto (dados de sessão); 0 para dados de configuração | C7 |
| NO-FUNCTIONAL-REQ-SCALABILITY-01 | Capacidade de requisições simultâneas de autenticação | Mínimo 5.000 RPS sustentadas | B-REQ-06, C4 |
| NO-FUNCTIONAL-REQ-SCALABILITY-02 | Concorrência máxima de sessões ativas | Mínimo 100.000 sessões ativas simultâneas | B-REQ-09 |
| NO-FUNCTIONAL-REQ-SCALABILITY-03 | Auto-scaling baseado em carga | Novas instâncias provisionadas automaticamente quando uso de CPU > 70% por 2 minutos | B-REQ-09 |
| NO-FUNCTIONAL-REQ-OBSERVABILITY-01 | Tracing distribuído de requisições | 100% das requisições com trace_id propagado entre serviços | — |
| NO-FUNCTIONAL-REQ-OBSERVABILITY-02 | Logs estruturados | 100% dos logs em formato JSON com correlation_id e tenant_id | B-REQ-07 |
| NO-FUNCTIONAL-REQ-OBSERVABILITY-03 | Métricas operacionais expostas | Endpoint de métricas disponível para coleta (Prometheus) | B-REQ-09 |
| NO-FUNCTIONAL-REQ-USABILITY-01 | Consistência de experiência de login | Comportamento de login/logout/erro idêntico para todos os produtos integrados | B-REQ-10 |

---

### 3. Restrições Técnicas de Sistema

| ID | Restrição | Descrição | Origem |
|----|-----------|-----------|--------|
| TECH-LIMIT-01 | Provedor de Nuvem Exclusivo | A plataforma opera exclusivamente na DigitalOcean (DOKS) | 005-BRD BC-02 |
| TECH-LIMIT-02 | Stack de Desenvolvimento | Java 21 + Quarkus + GraalVM Native — conformidade com padrão corporativo FBSO.ORG | 005-BRD BC-01 |
| TECH-LIMIT-03 | Provedor de Identidade | Keycloak como IDP corporativo, suportando Multi-Realm para isolamento de clientes | 005-BRD BC-03 |
| TECH-LIMIT-04 | API Gateway | Kong como ponto único de entrada — toda requisição externa passa pelo Kong | 005-BRD BC-03 |
| TECH-LIMIT-05 | Proteção de Borda | Cloudflare como camada de proteção DDoS e injeção de header de domínio (X-Tenant-Host) | 005-BRD BC-03 |
| TECH-LIMIT-06 | janela de Migração | Migrações de sistemas em produção devem ocorrer fora do horário comercial | 010-FRD B-LIMIT-07 |
| TECH-LIMIT-07 | Tempo Máximo de Rollback | Qualquer migração deve poder ser revertida em até 30 minutos | 010-FRD B-LIMIT-08 |

---

### 4. Matriz de Rastreabilidade (SRS → FRD → BRD)

| Requisito de Sistema (SRS) | Funcionalidade FRD (B-FEAT) | Requisito de Negócio BRD (B-REQ) | Critério Charter |
|:---|:---|:---|:---|
| FUNCTIONAL-REQ-01 — Extração de domínio | B-FEAT-01 | B-REQ-01 | C1 |
| FUNCTIONAL-REQ-02 — Cache de mapeamento | B-FEAT-01, B-FEAT-08 | B-REQ-01, B-REQ-08 | C1, C6 |
| FUNCTIONAL-REQ-03 — Início de autenticação (PKCE) | B-FEAT-03, B-FEAT-04 | B-REQ-03, B-REQ-04 | C2, C5 |
| FUNCTIONAL-REQ-04 — Troca e armazenamento seguro de tokens | B-FEAT-03, B-FEAT-04 | B-REQ-03, B-REQ-04 | C2, C5 |
| FUNCTIONAL-REQ-05 — Consulta de perfil | B-FEAT-04, B-FEAT-05 | B-REQ-04, B-REQ-05 | C3 |
| FUNCTIONAL-REQ-06 — Renovação de sessão | B-FEAT-04 | B-REQ-04 | C2, C5 |
| FUNCTIONAL-REQ-07 — Logout completo | B-FEAT-04 | B-REQ-04 | C2, C5 |
| FUNCTIONAL-REQ-08 — Filtro de isolamento | B-FEAT-02 | B-REQ-02 | C1 |
| FUNCTIONAL-REQ-09 — Registro de eventos de autenticação | B-FEAT-07 | B-REQ-07 | C8 |
| FUNCTIONAL-REQ-10 — Métricas operacionais | B-FEAT-09 | B-REQ-09 | C7 |
| FUNCTIONAL-REQ-11 — Ativação de novo cliente | B-FEAT-08 | B-REQ-08 | C6 |
| FUNCTIONAL-REQ-12 — Suspensão de cliente | B-FEAT-02 | B-REQ-02 | C1 |
| FUNCTIONAL-REQ-13 — Alerta de latência | B-FEAT-05, B-FEAT-06 | B-REQ-05, B-REQ-06 | C3, C4 |
| FUNCTIONAL-REQ-14 — Degradação controlada | B-FEAT-06 | B-REQ-06 | C4 |
| FUNCTIONAL-REQ-15 — Retenção e expurgo de logs | B-FEAT-07 | B-REQ-07 | C8 |
| FUNCTIONAL-REQ-16 — Migração com rollback | B-FEAT-10 | B-REQ-11 | C7 |
| FUNCTIONAL-REQ-17 — Relatório de acessos | B-FEAT-07 | B-REQ-07 | C8 |

**Cobertura:** 17/17 FRs vinculados a B-FEATs e B-REQs. 20 NFRs com métricas mensuráveis. 7 restrições técnicas documentadas. **100% rastreável.**

---

**[STATUS: SUCESSO]** — Documento de especificação de sistema. Arquitetura e design detalhado são reservados para 030-SAD, 035-HLD e 040-LLD.
