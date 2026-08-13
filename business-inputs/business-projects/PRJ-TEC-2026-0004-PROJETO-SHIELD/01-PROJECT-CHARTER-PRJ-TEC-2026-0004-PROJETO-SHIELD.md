# Project Charter: PROJETO SHIELD — Plataforma de Identidade e Segurança
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Produto** | Plataforma Shield — Identidade e Acesso Centralizada |
| **Data de Elaboração** | 03/08/2026 |
| **Versão** | 2.0 — Revisão para Apresentação à Diretoria |
| **Patrocinador** | Diretoria de Tecnologia — FBSO.ORG |
| **Metodologia** | WATERFALL |
| **Status** | Em análise |

**Documentos Relacionados:**
- `BRIEFING.md` — Briefing de Produto Plataforma Shield
- `02-BRD-PRJ-TEC-2026-0004-PROJETO-SHIELD.md` — Requisitos de Negócio (Próxima Fase)

--- 

### 1. O Problema de Negócio (Business Problem)

A FBSO.ORG opera um portfólio de produtos digitais para o setor educacional — Gestão Escolar, Comunidades de Ensino, Portal da Reforma. Cada produto gerencia seus próprios usuários e senhas de forma independente. Isso gera:

- **Risco de vazamento de dados entre clientes:** sem uma camada centralizada de segurança, os dados de uma escola podem ser acessados indevidamente por outra. Um incidente desse tipo compromete a reputação da empresa e pode gerar sanções regulatórias (LGPD).
- **Experiência fragmentada para os clientes:** cada escola que contrata nossos produtos passa por um processo manual e demorado de configuração de acessos, sem um portal unificado.
- **Custo crescente de manutenção:** cada produto duplica funcionalidades de login, recuperação de senha e controle de permissões, consumindo tempo de desenvolvimento que deveria estar focado no negócio.
- **Impossibilidade de crescer o portfólio:** lançar um novo produto hoje significa implementar autenticação do zero novamente — o que atrasa o time-to-market e encarece cada iniciativa.

**O custo de não resolver:** risco regulatório (LGPD), perda de competitividade, e inviabilidade de escalar o ecossistema de produtos para dezenas ou centenas de escolas.

---

### 2. A Solução Proposta (Product Purpose)

O **Projeto Shield** cria a **Plataforma de Identidade e Segurança da FBSO.ORG** — um produto interno que será a porta de entrada única para todos os sistemas da empresa.

O que ele entrega para o negócio:

- **Reconhecimento automático do cliente:** um diretor acessa o sistema da sua escola e a plataforma já sabe quem ele é e para onde direcioná-lo, sem qualquer configuração manual.
- **Isolamento total entre clientes:** os dados da Escola A são completamente inacessíveis para a Escola B. Cada cliente opera em seu próprio ambiente seguro — como se tivesse uma instalação dedicada.
- **Login protegido contra ataques:** as credenciais dos usuários jamais ficam expostas no navegador, eliminando os riscos mais comuns de roubo de sessão e vazamento de dados.
- **Velocidade de resposta:** a validação de identidade acontece em menos de 15 milissegundos — imperceptível para o usuário final.
- **Porta única para todos os produtos:** qualquer novo sistema da FBSO.ORG usará a mesma plataforma de identidade. Lançar um novo produto deixa de incluir o custo de reconstruir autenticação.

#### 2.1 Visão de Longo Prazo

- **Onboarding self-service:** uma nova escola poderá ser provisionada na plataforma em até 4 horas (hoje são dias).
- **Login unificado (SSO):** um usuário fará login uma vez e acessará todos os produtos FBSO — Gestão Escolar, Comunidades, Portal da Reforma — sem novas senhas.
- **Futuro:** integração com identidades que as escolas já usam (Google for Education, Microsoft 365 for Education, GOV.BR).

---

### 3. Escopo (Scope)

#### 3.1 Dentro do Escopo (In Scope)

1. **Plataforma de Identidade Centralizada** — O módulo central que reconhece o cliente pelo domínio, gerencia o fluxo de login/logout e protege as credenciais.
2. **Ambiente Isolado por Cliente** — Cada escola ou universidade opera em seu próprio ambiente isolado, sem possibilidade de acessar dados de outro cliente.
3. **Portal de Acesso Padronizado** — Conjunto de funcionalidades de autenticação (`login`, `logout`, `troca de senha`, `perfil do usuário`) disponível para todos os produtos do ecossistema.
4. **Infraestrutura de Segurança** — Camada de proteção que garante que credenciais e dados pessoais jamais fiquem expostos ou sejam registrados em logs.
5. **Base para Auditoria de Conformidade** — Registro de eventos de acesso que permitirá auditoria LGPD e futura certificação ISO 27001.
6. **Observabilidade do Negócio** — Métricas e dashboards sobre quem está acessando, de onde, e se há tentativas de acesso indevido.
7. **Suíte de Validação de Segurança** — Testes que comprovam que os dados de clientes estão isolados e protegidos.
8. **Documentação do Produto** — Guias para os times que consumirão a plataforma e manuais para a equipe de operações.

#### 3.2 Fora do Escopo (Out of Scope)

- Desenvolvimento de funcionalidades de negócio dos produtos que consumirão a plataforma (Gestão Escolar, etc.)
- Migração de bases de usuários já existentes para a nova plataforma
- Integração com provedores externos de identidade (Google, Microsoft, GOV.BR) — fase futura
- Desenvolvimento de interfaces de usuário das aplicações cliente
- Certificação ISO 27001 completa — este projeto entrega a base técnica para a certificação

---

### 4. Entregas e Critérios de Aceitação

| # | Entrega | Critérios de Aceitação de Negócio | Data-Alvo |
|---|---|---|---|
| D1 | Ambiente de Produção da Plataforma | Infraestrutura provisionada, segura e operacional. Acesso restrito e controlado | Semana 2 |
| D2 | Motor de Identidade por Cliente | Cada cliente possui seu ambiente isolado. Login de um cliente não acessa dados de outro — comprovado por teste | Semana 2 |
| D3 | Portal de Acesso (Login, Logout, Perfil) | Funcionalidades de autenticação operacionais para consumo por qualquer produto FBSO. Tempo de resposta abaixo de 15ms | Semana 4 |
| D4 | Camada de Isolamento de Dados | Dados de clientes isolados em nível de armazenamento. Tentativa de acesso cruzado não retorna dados — comprovado por auditoria | Semana 4 |
| D5 | Monitoramento e Alertas | Dashboards mostrando acessos ativos, tentativas bloqueadas e saúde da plataforma | Semana 5 |
| D6 | Homologação de Segurança | Testes confirmam: (1) dados isolados entre clientes, (2) credenciais jamais expostas, (3) plataforma suporta picos de acesso | Semana 6 |
| D7 | Liberação para Uso (Go-Live) | Plataforma em produção, documentação publicada, termo de aceite assinado | Semana 6 |

---

### 5. Partes Interessadas e Matriz RACI

| Parte Interessada | Papel | D1 Infra | D2 Motor ID | D3 Portal | D4 Isolamento | D5 Monitor | D6 Homolog | D7 Go-Live |
|---|---|---|---|---|---|---|---|---|
| Diretoria de Tecnologia | Patrocinador | A | A | A | A | A | A | A |
| Diretoria de Negócios | Cliente Interno | I | I | C | I | I | I | C |
| Product Owner (PO) | Definição de Produto | C | C | R | C | I | R | R |
| Project Manager (PM) | Coordenação | R | R | R | R | R | R | R |
| Tech Lead / Arquiteto | Liderança Técnica | R | R | R | R | R | C | C |
| Especialista em Segurança | Proteção de Dados | C | R | C | C | I | R | I |
| Equipe de Desenvolvimento | Implementação | I | C | R | C | C | C | I |
| Equipe de Infraestrutura | Operação | R | C | C | C | R | C | I |
| Equipe de Qualidade (QA) | Validação | I | I | C | C | I | R | C |

---

### 6. Critérios de Sucesso

| # | Critério | Como Medimos | Meta |
|---|---|---|---|
| C1 | Segurança entre Clientes | Tentativas de acessar dados de outro cliente | 100% bloqueadas |
| C2 | Proteção de Credenciais | Verificação de que credenciais não aparecem em lugar nenhum fora do ambiente seguro | 100% dos acessos protegidos |
| C3 | Velocidade de Resposta | Tempo para validar identidade do usuário | Abaixo de 15ms |
| C4 | Capacidade de Atender Picos | Simular o horário de entrada de todas as escolas simultaneamente | Sem falhas no sistema |
| C5 | Cobertura de Testes de Segurança | Principais categorias de ataque cobertas | 100% dos cenários |
| C6 | Tempo para Adicionar Novo Cliente | Processo completo de ativação de uma nova escola | Abaixo de 4 horas |
| C7 | Disponibilidade da Plataforma | Tempo em que o sistema de login fica fora do ar | Máximo 0.1% do tempo (99.9%) |
| C8 | Rastreabilidade de Acessos | Registro de cada tentativa de acesso (bem-sucedida ou não) | 100% dos eventos |

---

### 7. Premissas

1. A infraestrutura de nuvem contratada (DigitalOcean) estará disponível e provisionada conforme plano corporativo.
2. O ambiente de identidade por cliente (Keycloak) será configurado no modelo de isolamento por organização.
3. As aplicações cliente (frontends dos produtos FBSO) serão adaptadas para consumir a nova plataforma de login.
4. A equipe de infraestrutura terá acesso necessário para configurar domínios e proteção de borda (Cloudflare).
5. Os padrões tecnológicos corporativos (documento de arquitetura) são vinculantes — desvios precisam ser aprovados.
6. As aplicações de negócio existentes delegarão a validação de identidade à nova plataforma.
7. A plataforma será otimizada para ocupar pouca memória e iniciar rapidamente, viabilizando escala sob demanda.

---

### 8. Restrições

| # | Restrição | Descrição |
|---|---|---|
| R1 | Prazo | 6 semanas (30 dias úteis) — duração máxima do projeto |
| R2 | Provedor de Nuvem | Exclusivamente DigitalOcean |
| R3 | Proteção de Borda | Todo tráfego externo passa por camada de segurança (Cloudflare) |
| R4 | Padrões Corporativos | Tecnologias fora do padrão exigem justificativa e aprovação do comitê |
| R5 | Proteção de Dados (LGPD) | Dados pessoais de alunos e professores não podem aparecer em logs; ambientes de clientes devem ser totalmente isolados |
| R6 | Orçamento | A definir pelo patrocinador; primeira projeção de custos na Semana 2 |
| R7 | Equipe | 9 profissionais alocados (alguns em dedicação parcial). Mudanças na equipe impactam o cronograma |

---

### 9. Riscos de Alto Nível

| Risco | Probabilidade | Impacto | Mitigação |
|---|---|---|---|
| Curva de aprendizado em tecnologia de compilação nativa | Alta | Alto | Buffer de 24h reservado; time recebe mentoria do Arquiteto |
| Complexidade de configuração do ambiente multi-cliente | Média | Alto | Especialista dedicado com 104h alocadas; prova de conceito antes do desenvolvimento |
| Performance do isolamento de dados em cenários complexos | Média | Médio | Especialista de banco de dados com 64h; testes de carga específicos |
| Indisponibilidade de serviços de nuvem gerenciados | Baixa | Alto | Configuração de alta disponibilidade; backups automáticos; procedimento de recuperação documentado |
| Vazamento de dados entre clientes | Média | Crítico | Testes dedicados de isolamento; revisão de segurança antes da liberação |
| Atraso na adaptação dos produtos consumidores | Média | Médio | Desenvolvedor frontend alocado desde o início; contrato de integração definido cedo |
| Exigências excessivas de customização visual por cliente | Baixa | Médio | Escopo inicial limita a tema padrão; customizações tratadas como projeto separado |

---

### 10. Marcos do Projeto

| Marco | Data | Critério de Conclusão |
|---|---|---|
| M1: Kickoff | 04/08/2026 | Project Charter aprovado, equipe mobilizada |
| M2: Base Pronta | 15/08/2026 | Infraestrutura e ambientes por cliente operacionais |
| M3: Plataforma Desenvolvida | 29/08/2026 | Portal de acesso funcional, isolamento de dados ativo e verificado |
| M4: Monitoramento Ativo | 05/09/2026 | Dashboards, logs e alertas 100% operacionais |
| M5: Homologação Concluída | 12/09/2026 | Testes de segurança e isolamento aprovados, zero pendências críticas |
| M6: Go-Live | 15/09/2026 | Plataforma em produção, termo de aceite assinado |

---

### 11. Orçamento Estimado

| Categoria | Descrição | Estimativa |
|---|---|---|
| **Equipe (992 horas)** | 9 profissionais (Arquiteto, Desenvolvedores, Infraestrutura, Segurança, Banco de Dados, Qualidade, PM, PO) | 992 horas/homem |
| **Nuvem (6 semanas)** | Servidores, banco de dados gerenciado, cache, balanceadores | ~R$ 7.000-10.000/mês |
| **Ferramentas** | Proteção de borda (Cloudflare), automação (GitHub Actions) | ~R$ 100-1.000/mês |
| **Reserva de Contingência (15%)** | Buffer para riscos de curva de aprendizado e complexidade | ~150 horas adicionais |
| **Total Estimado** | | **~1.140 horas + custos de nuvem** |

---

### 12. Plano de Comunicação

| Público | Frequência | Canal | Responsável |
|---|---|---|---|
| Diretoria de Tecnologia (Patrocinador) | Quinzenal | Relatório executivo + reunião 30min | Project Manager |
| Diretoria de Negócios | Marco M5 (pré-Go-Live) | Apresentação de produto + demonstração | Product Owner |
| Product Owner | Diário (assíncrono) + Planejamento semanal | Slack + Planning | Project Manager |
| Equipe do Projeto | Reunião diária (15min) | Virtual + Slack | Project Manager |
| Times Consumidores (outros produtos FBSO) | Quinzenal | E-mail de status + Demo técnica | Product Owner + Tech Lead |
| Clientes Piloto | Sob demanda | Via Product Owner | Product Owner |

---

### 13. Governança

- **Comitê de Projeto:** Diretoria de Tecnologia + Product Owner + Tech Lead. Reúne-se quinzenalmente para aprovar mudanças de escopo, orçamento e riscos.
- **Liderança Técnica:** Responsável por todas as decisões de arquitetura. Tecnologias fora do padrão corporativo exigem justificativa documentada e aprovação.
- **Gestão de Projeto:** Responsável por cronograma, dependências, riscos, custos e relatórios de status. Remove bloqueios da equipe.
- **Produto:** Responsável por definir e priorizar funcionalidades, validar critérios de aceite, e ser a voz do cliente interno.
- **Controle de Mudanças:** Alterações de escopo, prazo ou orçamento acima de 10% exigem aprovação formal do Comitê.
- **Portal de Segurança:** Nenhuma liberação para produção sem revisão de segurança aprovada pelo Tech Lead e Especialista em Segurança.

---

### 14. Aprovações

| Nome | Papel | Data | Assinatura |
|---|---|---|---|
| [Diretor de Tecnologia] | Patrocinador | __/__/____ | ______________ |
| [Diretor de Negócios] | Cliente Interno | __/__/____ | ______________ |
| [Product Owner] | Definição de Produto | __/__/____ | ______________ |
| [Tech Lead / Arquiteto] | Aprovação Técnica | __/__/____ | ______________ |
| [Project Manager] | Aprovação de Planejamento | __/__/____ | ______________ |

---

**[STATUS: SUCESSO]** — Documento revisado para linguagem de negócio. Detalhes técnicos reservados para os documentos de arquitetura (Fase 3).
