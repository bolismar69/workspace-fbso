# Mapa do Time Técnico — Alocação e Competências

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Programa:** FBSO Platform — Portal Administrativo SaaS
- **Data de Criação:** 13 de Julho de 2026
- **Versão:** 1.2
- **Data de Criação:** 13 de Julho de 2026
- **Última Atualização:** 15 de Julho de 2026 (revisão caveman: skills específicas, nota de acúmulo de papéis)
- **Status:** ⚠️ TODO — Aguardando definição do time técnico e preenchimento de nomes, contatos e alocações

---

## 1. Objetivo

Este documento é a planilha de alocação do time técnico. Ele relaciona todos os papéis necessários ao projeto, as competências requeridas, a fase em que cada profissional atua e a carga horária prevista.

As colunas **Nome** e **Contato** devem ser preenchidas pelo Coordenador do Projeto ou Tech Lead responsável tão logo os profissionais sejam designados.

> ⚠️ **ATENÇÃO:** Neste momento, a FBSO.ORG possui um time técnico reduzido. As vagas listadas abaixo representam a estrutura ideal necessária para entregar o escopo completo do projeto (D1-D7) no cronograma de 14 semanas. O preenchimento real dependerá de contratações, realocação de recursos ou ajuste de escopo e prazo.

---

## 2. Legenda

| Campo | Descrição |
|:---|:---|
| **Nome** | Nome completo do profissional designado (a preencher) |
| **Contato** | E-mail ou canal de comunicação corporativo (a preencher) |
| **Papel** | Função no projeto |
| **Skill Primária** | Principal competência técnica exigida para o papel |
| **Nível** | Proficiência esperada: ★★★ Avançado/Autônomo, ★★☆ Intermediário/Produtivo, ★☆☆ Básico/Assistido |
| **Fases** | Em quais fases do projeto o profissional atua: F0 = Fundação (Sprint 0), F1 = Desenvolvimento EP-01/EP-02, F2 = Desenvolvimento EP-03/EP-04, F3 = Homologação e Go-Live |
| **Horas/dia** | Carga horária diária prevista com base no percentual de alocação × jornada padrão de 8h |
| **Horas Totais** | Estimativa de horas totais no projeto (a recalcular quando as durações dos sprints forem definidas) |

### 2.1 Correspondência de Fases

| Código | Fase | Sprints | Marcos | Duração Estimada |
|:---|:---|:---|:---|:---|
| **F0** | Fundação e Setup | Sprint 0 | Kickoff → Setup do ambiente | 1-2 semanas (24/07 a 07/08) |
| **F1** | Desenvolvimento EP-01 e EP-02 | Sprints 1-2 | M2 (Portal Admin) + M3 (Contas e Planos) | 4 semanas (07/08 a 31/08) |
| **F2** | Desenvolvimento EP-03 e EP-04 | Sprints 3-5 | M4 (RBAC) + M5 (Portal Cliente) + M6 (BUs e Catálogo) | 6 semanas (01/09 a 15/10) |
| **F3** | Homologação e Go-Live | Sprint 6+ | M7 (Aceite Final) | 2 semanas (15/10 a 30/10) |

> **Nota:** As durações acima são estimativas baseadas no cronograma do Project Charter. As horas totais serão recalculadas quando o time for definido.

---

## 3. Time de Desenvolvimento

> **Nota sobre senioridade:** Os papéis abaixo são genéricos (Full-Stack, Frontend, Backend) pois o time ainda não foi definido. Quando os profissionais forem designados, recomenda-se detalhar os níveis de senioridade (Tech Lead, Sênior, Pleno) conforme o padrão do projeto de referência PRJ-FIN-2026-0002, que estratifica cada papel em três níveis com competências e alocações específicas por senioridade.

### 3.1 Desenvolvedor Full-Stack (Backend + Frontend)

| Nome | Contato | Papel | Skill Primária | Nível | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---:|:---:|
| `<TODO>` | `<TODO>` | Desenvolvedor Full-Stack | Desenvolvimento Web Full-Stack | ★★★ | F0, F1, F2, F3 | A definir | ⬜ |

**Competências requeridas:** Java 25 LTS + Spring Boot (backend), React + Next.js + Tailwind CSS (frontend), APIs REST, PostgreSQL, Keycloak (autenticação/autorização JWT + SAML), testes automatizados, Docker, controle de versão (Git).

### 3.2 Desenvolvedor Frontend / UX

| Nome | Contato | Papel | Skill Primária | Nível | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---:|:---:|
| `<TODO>` | `<TODO>` | Desenvolvedor Frontend | Desenvolvimento Frontend | ★★★ | F0, F1, F2, F3 | A definir | ⬜ |

**Competências requeridas:** React + Next.js (App Router) + Tailwind CSS, componentização com design systems, MSW mock para desenvolvimento paralelo, usabilidade e acessibilidade (WCAG AA), consumo de APIs REST com JWT, testes de interface (Playwright/Cypress).

### 3.3 Desenvolvedor Backend

| Nome | Contato | Papel | Skill Primária | Nível | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---:|:---:|
| `<TODO>` | `<TODO>` | Desenvolvedor Backend | Desenvolvimento Backend / APIs | ★★★ | F0, F1, F2, F3 | A definir | ⬜ |

**Competências requeridas:** Desenvolvimento backend, APIs REST, modelagem de dados, autenticação/autorização (RBAC), migrações de banco de dados, testes automatizados.

---

## 4. Time de Qualidade

### 4.1 QA / Test Engineer (1-2 vagas)

| Nome | Contato | Papel | Skill Primária | Nível | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---:|:---:|
| `<TODO>` | `<TODO>` | QA / Test Engineer | Testes Automatizados E2E | ★★★ | F1, F2, F3 | A definir | ⬜ |
| `<TODO>` | `<TODO>` | QA / Test Engineer | Testes Manuais / Exploratórios | ★★★ | F1, F2, F3 | A definir | ⬜ |

**Competências requeridas:** Testes automatizados E2E (★★★), Testes manuais/exploratórios (★★★), Testes de regressão (★★★), Homologação/UAT (★★☆), Análise de requisitos/critérios de aceite (★★☆), Metodologias ágeis (★★☆), Domínio de negócio SaaS (★☆☆).

### 4.2 Time de Homologação — Negócio (1-2 vagas)

| Nome | Contato | Papel | Skill Primária | Nível | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---:|:---:|
| `<TODO>` | `<TODO>` | Analista de Homologação (Negócio) | Domínio de Negócio SaaS | ★★★ | F1, F2, F3 | A definir | ⬜ |
| `<TODO>` | `<TODO>` | Analista de Homologação (Negócio) | Domínio de Negócio SaaS / Portal | ★★★ | F1, F2, F3 | A definir | ⬜ |

**Competências requeridas:** Domínio de negócio SaaS/portais (★★★), Testes manuais/exploratórios (★★★), Homologação/UAT (★★★), Testes de regressão (★★★), Análise de requisitos/critérios de aceite (★★★), Testes automatizados E2E (★☆☆).

> **Nota:** O time de homologação é idealmente composto por usuários-chave do time administrativo e comercial da FBSO.ORG, com dedicação mais intensa na Fase 3 (Homologação e Go-Live).

---

## 5. Times Transversais

### 5.1 Arquiteto de Solução

| Nome | Contato | Papel | Skill Primária | Nível | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---:|:---:|
| `<TODO>` | `<TODO>` | Arquiteto de Solução | Arquitetura de Software | ★★★ | F0, F1, F2 | A definir | ⬜ |

**Competências requeridas:** Arquitetura de software, desenho de APIs, modelagem de dados, segurança de aplicações, definição de stacks tecnológicas.

> **Nota:** Alocação mais intensa em F0 e F1 (definições e setup). Em F2 atua com carga reduzida (estimativa: 20-30%) para orientação arquitetural e revisão de decisões técnicas durante o desenvolvimento de EP-03 e EP-04.

### 5.2 DevOps / Infraestrutura

| Nome | Contato | Papel | Skill Primária | Nível | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---:|:---:|
| `<TODO>` | `<TODO>` | DevOps Engineer | Infraestrutura e CI/CD | ★★★ | F0, F1, F3 | A definir | ⬜ |

**Competências requeridas:** Infraestrutura como código, pipelines de CI/CD, gestão de ambientes, monitoramento, gestão de segredos e configurações.

---

## 6. Gestão

### 6.1 Product Manager — PM (1 vaga)

| Nome | Contato | Papel | Skill Primária | Nível | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---:|:---:|
| `<TODO>` | `<TODO>` | Product Manager (PM) | Gestão de Projetos Ágeis | ★★★ | F0, F1, F2, F3 | A definir | ⬜ |

**Competências requeridas:** Gestão de projetos ágeis Scrum/Kanban (★★★), Gestão de stakeholders (★★★), Comunicação e facilitação (★★★), Domínio de negócio SaaS (★★☆), Análise de requisitos (★★☆).

### 6.2 Product Owner (PO)

| Nome | Contato | Papel | Skill Primária | Nível | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---:|:---:|
| `<TODO>` | `<TODO>` | Product Owner (PO) | Domínio de Negócio SaaS | ★★★ | F0, F1, F2, F3 | A definir | ⬜ |

**Competências requeridas:** Domínio de negócio SaaS, gestão de stakeholders, análise de requisitos/critérios de aceite, comunicação e facilitação, homologação/UAT.

### 6.3 Coordenador do Projeto

| Nome | Contato | Papel | Skill Primária | Nível | Fases | Horas/dia | Horas Totais |
|:---|:---|:---|:---|:---:|:---|:---:|:---:|
| `<TODO>` | `<TODO>` | Coordenador do Projeto | Gestão de Projetos | ★★★ | F0, F1, F2, F3 | A definir | ⬜ |

**Competências requeridas:** Gestão de projetos, gestão de stakeholders, comunicação e facilitação, gestão de riscos, metodologias ágeis.

---

## 7. Tabela Resumo

### 7.1 Headcount por Área

| Área | Papéis | Vagas (Ideais) | Vagas (Preenchidas) | Status |
|:---|:---|:---:|:---:|:---|
| **Desenvolvimento** | Full-Stack, Frontend, Backend | 3 vagas (3 papéis — podem ser acumulados em 1-2 pessoas no cenário de time reduzido) | 0 | ⚠️ TODO |
| **Qualidade** | QA/Test (1-2), Homologação Negócio (1-2) | 2-4 | 0 | ⚠️ TODO |
| **Transversais** | Arquiteto, DevOps | 2 | 0 | ⚠️ TODO |
| **Gestão** | PM, PO, Coordenador | 3 | 0 | ⚠️ TODO |
| **Total** | | **10-12** | **0** | ⚠️ TODO |

### 7.2 Alocação Estimada por Fase (a confirmar)

| Fase | Desenvolvimento (3p) | Qualidade (2-4p) | Transversais (2p) | Gestão (3p) | Total Ativo |
|:---|:---:|:---:|:---:|:---:|:---:|
| **F0 — Fundação** | 3,0 (100%) | 0,5 (15%) | 1,0 (50%) | 2,5 (83%) | **7,0** |
| **F1 — EP-01 e EP-02** | 3,0 (100%) | 1,0 (30%) | 0,5 (25%) | 2,5 (83%) | **7,0** |
| **F2 — EP-03 e EP-04** | 3,0 (100%) | 1,5 (45%) | 0,5 (25%) | 2,5 (83%) | **7,5** |
| **F3 — Homologação** | 1,5 (50%) | 3,5 (100%) | 0,5 (25%) | 3,0 (100%) | **8,5** |

> **Nota:** Os números acima são projeções baseadas na estrutura ideal de 10-12 profissionais. A Fase 3 (Homologação) concentra a maior alocação, com o time de qualidade em carga máxima e gestão em 100%. Com o time técnico reduzido atual, os percentuais de alocação e os prazos precisarão ser recalibrados assim que a equipe real for definida.

### 7.3 Marcos de Alocação

| Marco | Fase | O que acontece |
|:---|:---|:---|
| **Pré-Sprint 0** | — | Coordenador, PM e Arquiteto agendam sessões de definição com os times para o Sprint 0 |
| **Sprint 0 (F0)** | Fundação | Todos os times mobilizados. Desenvolvimento em carga máxima (setup do ambiente, definição arquitetural). UX/Frontend inicia prototipação. Qualidade com carga mínima (planejamento de testes). Arquiteto e DevOps em 50%. |
| **Sprint 1-2 (F1)** | EP-01 e EP-02 | **Pico do Desenvolvimento.** Portal Admin (EP-01) e Gestão de Clientes/Planos (EP-02) construídos em paralelo. Times transversais reduzem carga (suporte). QA inicia preparação de casos de teste. Homologação de Negócio inicia revisão de critérios de aceite. |
| **Sprints 3-5 (F2)** | EP-03 e EP-04 | **Fase mais longa do projeto** (6 semanas). RBAC (EP-03) e Portal do Cliente + BUs + Catálogo (EP-04) desenvolvidos sequencialmente. Qualidade aumenta progressivamente. Arquiteto em carga reduzida para orientação. PM e PO ajustam backlog conforme feedback dos early adopters. |
| **Sprint 6+ (F3)** | Homologação e Go-Live | **Pico de Qualidade e Gestão.** Desenvolvimento reduz para 50% (correção de bugs). QA em 100% (regressão completa). Homologação de Negócio em 100% (UAT). PM, PO e Coordenador em 100% coordenando aceites, revisão de riscos e encerramento. |

---

## 8. Instruções de Preenchimento

1. **Nome e Contato:** Preencher tão logo o profissional seja designado. O Coordenador do Projeto é responsável por manter esta planilha atualizada.
2. **Horas Totais:** A coluna será preenchida quando as durações dos sprints forem definidas. Fórmula: `Horas/dia × dias úteis na fase × semanas da fase`.
3. **Nível:** A proficiência indicada é a **esperada** para o papel. Se o profissional designado tiver nível diferente, revisar o plano de capacitação ou mentoria.
4. **Substituições:** Em caso de substituição de profissional, atualizar a linha correspondente e registrar a alteração no histórico abaixo.
5. **Time Reduzido:** Caso a equipe real seja menor que a estrutura ideal (10-12 profissionais), revisitar o cronograma (Project Charter §7) e a matriz de riscos (R2 — Equipe reduzida) para avaliar impacto em prazos e escopo.

### 8.1 Histórico de Alterações

| Data | Alteração | Responsável |
|:---|:---|:---|
| 13/07/2026 | Criação do documento (v1.0) — estrutura ideal, sem profissionais designados | Time de Negócios |
| 13/07/2026 | Revisão (v1.1) — Adição de PM (1 vaga), QA/Test Engineer expandido para 1-2 vagas, Time de Homologação — Negócio (1-2 vagas), Arquiteto com atuação estendida para F2, nota de senioridade no §3, coluna Sprints na tabela de fases, seção Marcos de Alocação (§7.3), correção de headcount total para 10-12 profissionais | Time de Negócios |

---

> ⚠️ **TODO:** Este documento depende de decisões organizacionais sobre contratação e alocação de recursos técnicos. Enquanto o time não for definido, o risco R2 do Project Charter ("Equipe reduzida não consegue entregar no prazo esperado") permanece com severidade Crítica.

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: agile-ba-practices, stakeholder-analysis.*
🔍 *Revisado pelo skill caveman-review em 15/07/2026. Ajustes aplicados: skills específicas da stack (Java 25 LTS, Spring Boot, React, Next.js, Keycloak) nas competências, headcount com nota de acúmulo de papéis, campo "Última Atualização".*
