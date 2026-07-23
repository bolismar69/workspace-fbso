# Mapa do Time Técnico — Resumo de Capacidade de Trabalho

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Programa:** FBSO Platform — Portal Administrativo SaaS
- **Versão:** 1.5
- **Data de Criação:** 13 de Julho de 2026
- **Última Atualização:** 22 de Julho de 2026 (adição de Tech Lead — gap analysis operacional: liderança técnica diária)
- **Status:** ⚠️ TODO — Aguardando definição dos integrantes técnicos que executarão as soluções

---

## Objetivo

Este documento apresenta a tabela de capacidade do time técnico — os profissionais que efetivamente executarão as tarefas de implementação das soluções técnicas do projeto. Ele serve como referência rápida para o Tech Lead e Coordenador alocarem trabalho nos sprints.

As colunas **Nome** e **Contato** devem ser preenchidas tão logo os profissionais sejam designados.

> ⚠️ **ATENÇÃO:** A FBSO.ORG possui um time técnico reduzido. Os papéis abaixo representam a estrutura necessária para entregar o escopo técnico. O preenchimento real dependerá de contratações, realocação ou ajuste de escopo e prazo.

---

## Time Técnico x Capacidade de Trabalho

| Papel                             | Nível | Data Inicial | Data Final | Horas/dia | Dias da semana | Nome     | Contato |
|-----------------------------------|-------|--------------|------------|-----------|----------------|----------|---------|
| Tech Lead / Líder Técnico         | ★★★   |              |            | 8         | 1,2,3,4,5      | FRANCISCO OLIVEIRA | francisco.oliveira.fbso.org@gmail.com |
| Desenvolvedor Full-Stack          | ★★★   |              |            | 8         | 1,2,3,4,5      | BOLISMAR OLIVEIRA | bolismar.oliveira.fbso.org@gmail.com |
| Desenvolvedor Frontend            | ★★★   | 01/11/2026   |            | 8         | 1,2,3,4,5      | TOM SANTOS | tom.santos.fbso.org@gmail.com |
| Desenvolvedor Backend             | ★★★   |              |            | 8         | 1,2,3,4,5      | MARIA MADALENA | maria.madalena.fbso.org@gmail.com |
| QA / Test Engineer                | ★★★   |              |            | 8         | 1,2,3,4,5      | FELIPE CANEDAS | felipe.canedas.fbso.org@gmail.com |
| Analista de Homologação (Negócio) | ★★★   |              |            | 8         | 1,2,3,4,5      | MAURO HANASHIRO | mauro.hanashiro.fbso.org@gmail.com |
| Desenvolvedor Banco de Dados      | ★★★   |              |            | 8         | 1,2,3,4,5      | CARLOS CALDAS | carlos.caldas.fbso.org@gmail.com |
| Arquiteto de Solução              | ★★★   |              |            | 4         | 1,2,3,4,5      | BRUNO GRATTO | bruno.gratto.fbso.org@gmail.com |
| DevOps Engineer                   | ★★★   |              |            | 4         | 1,2,3,4,5      | DAVI SILVA | davi.silva.fbso.org@gmail.com |
| Especialista IAM / Keycloak       | ★★★   |              |            | 8         | 1,2,3,4,5      | GERTRUDES PAIVA | gertrudes.paiva.fbso.org@gmail.com |

---

## Legenda

| Campo | Descrição |
|:---|:---|
| **Nome** | Nome completo do profissional designado (a preencher) |
| **Contato** | E-mail ou canal de comunicação corporativo (a preencher) |
| **Papel** | Função no projeto |
| **Nível** | Proficiência esperada: ★★★ Senior/Especialista/Avançado/Autônomo, ★★☆ Pleno/Intermediário/Produtivo, ★☆☆ Junior/Básico/Assistido |
| **Horas/dia** | Carga horária diária prevista (8h = dedicação integral, 4h = dedicação parcial) |
| **Dias da semana** | 1-Segunda, 2-Terça, 3-Quarta, 4-Quinta, 5-Sexta, 6-Sábado, 7-Domingo |

> **Nota sobre papéis em carga parcial (4h/dia):** Arquiteto, DevOps e Especialista IAM/Keycloak atuam com dedicação parcial pois seu foco é concentrado em fases específicas: mais intenso em F0 (Fundação/Setup) e reduzido em F1-F2 (desenvolvimento), conforme definido no [PROJECT-TEAM-MAP.md](./PROJECT-TEAM-MAP.md).

---

## Instruções de Preenchimento

1. **Nome e Contato:** Preencher tão logo o profissional seja designado. O Coordenador do Projeto é responsável por manter esta planilha atualizada.
2. **Carga horária:** Tech Lead e todos os papéis de desenvolvimento, qualidade e homologação operam em 8h/dia (dedicação integral). Arquiteto, DevOps e Especialista IAM/Keycloak em 4h/dia (dedicação parcial, com picos em fases específicas).
3. **Substituições:** Em caso de substituição de profissional, atualizar a linha correspondente e registrar abaixo.

### Histórico de Alterações

| Data | Alteração | Responsável |
|:---|:---|:---|
| 13/07/2026 | Criação do documento original (v1.0) como parte do PROJECT-TEAM-MAP.md | Time de Negócios |
| 22/07/2026 | Refatoração (v1.3) — Documento reduzido ao escopo da seção 7.4 (time técnico executor). Adicionados Arquiteto de Solução e DevOps Engineer (oriundos das seções 5.1 e 5.2 do PROJECT-TEAM-MAP.md) com carga de 4h/dia. | Time de Negócios |
| 22/07/2026 | Gap analysis (v1.4) — Adicionado Especialista IAM/Keycloak (4h/dia) para cobrir Keycloak 26 como plataforma IAM dedicada. Nota de carga parcial estendida para incluir o novo perfil. | Time de Negócios |
| 22/07/2026 | Gap analysis operacional (v1.5) — Adicionado Tech Lead / Líder Técnico (8h/dia) como perfil dedicado de liderança técnica diária (code review, mentoria, decisões técnicas, gestão de dívida técnica). Papel já era referenciado no documento mas não estava formalizado na tabela. | Time de Negócios |

---

> ⚠️ **TODO:** Este documento depende de decisões organizacionais sobre contratação e alocação de recursos técnicos. Enquanto o time não for definido, o risco R2 do Project Charter ("Equipe reduzida não consegue entregar no prazo esperado") permanece com severidade Crítica.

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude.*
🔍 *Refatorado em 22/07/2026 — escopo reduzido para contemplar exclusivamente a tabela de capacidade do time técnico executor (§7.4).*
