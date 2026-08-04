# Plano de Comunicação: PROJETO SHIELD
## [STATUS: COMPLIANCE]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | PRJ-TEC-2026-0004-PROJETO-SHIELD |
| **Documento Base** | 01-PROJECT-CHARTER |
| **Data** | 03/08/2026 | **Versão** | 1.0 | **Metodologia** | WATERFALL |

---

## 1. Matriz de Comunicação

| Público | Mensagem | Frequência | Canal | Responsável |
|---------|----------|-----------|------|-------------|
| Diretoria de Tecnologia (Sponsor) | Progresso, riscos, decisões de escopo/orçamento | Quinzenal (sexta) | Relatório executivo + reunião 30min | PM |
| Diretoria de Negócios | Visão do produto, demonstração, impacto no negócio | Marco M5 (pré-Go-Live) | Apresentação presencial/virtual + demo | PO |
| Product Owner | Status das entregas, bloqueios, mudanças de escopo | Diário (async) + Planning semanal | Slack + Planning Session | PM |
| Equipe do Projeto (9 pessoas) | Tarefas do dia, bloqueios, dependências | Diário (15min, 9h) | Reunião virtual + Slack #projeto-shield | PM |
| Times Consumidores (outros produtos) | Contrato de integração, documentação, cronograma | Quinzenal | E-mail + demo técnica | PO + Tech Lead |
| Clientes Piloto | Disponibilidade, mudanças, feedback | Sob demanda | Via Product Owner | PO |
| Compliance/Jurídico | Conformidade LGPD, isolamento de dados | Marco M5 (pré-Go-Live) | Relatório de auditoria de segurança | Tech Lead + IAM |

## 2. Fluxo de Escalação

| Nível | Gatilho | Responsável | Ação |
|-------|--------|-------------|------|
| **N1 — Equipe** | Bloqueio técnico, dependência não resolvida | Tech Lead | Resolução em até 4h |
| **N2 — Tech Lead / PM** | Impedimento persiste >4h, risco de atraso no marco | PM + Tech Lead | Plano de ação em até 24h |
| **N3 — Comitê de Projeto** | Atraso >20% do cronograma, mudança de escopo >10% | Diretoria + PO + Tech Lead | Reunião extraordinária em até 48h |
| **N4 — Emergência** | Incidente de segurança, vazamento de dados | Diretoria | Comunicação imediata; war room |

## 3. Calendário de Reuniões e Rituais

| Reunião | Frequência | Participantes | Duração | Dia/Hora |
|---------|-----------|--------------|---------|---------|
| Daily Standup | Diário | Equipe completa | 15min | 9h00 |
| Planning Semanal | Semanal (segunda) | Equipe + PO | 1h | 10h00 |
| Review / Demo | Quinzenal (sexta) | Equipe + PO + Stakeholders | 45min | 15h00 |
| Retrospectiva | Quinzenal (sexta) | Equipe | 30min | 16h00 |
| Steering Committee | Quinzenal | Diretoria + PO + Tech Lead + PM | 30min | Sexta 14h |

## 4. Repositório de Documentos

| Tipo de Documento | Localização | Acesso |
|------------------|-------------|--------|
| Documentos do Projeto (Charter, BRD, etc.) | `business-projects/PRJ-TEC-2026-0004-PROJETO-SHIELD/` | Equipe + Diretoria |
| Código Fonte | `backend/java/spring/microservices/ms-shield-identity-auth/` | Equipe técnica |
| Infraestrutura como Código | Repositório GitOps | DevOps + Tech Lead |
| Documentação de API | OpenAPI publicado no portal interno | Todos os times consumidores |
| Relatórios de Status | Pasta compartilhada no Google Drive | Diretoria + PO + PM |

## 5. Plano de Comunicação em Crise

| Cenário | Quem Comunicar | Quando | Canal |
|---------|---------------|--------|------|
| Incidente de segurança (vazamento, breach) | Diretoria + Jurídico + Compliance | Imediato (< 1h) | Telefone + War Room |
| Indisponibilidade da plataforma (> 30min) | Times consumidores + PO | < 30min | Slack + E-mail |
| Atraso em marco (> 1 semana) | Diretoria + PO | < 24h | Relatório de impacto |
| Saída de membro crítico da equipe | PM + Diretoria | < 24h | Reunião de replanejamento |

---

**[STATUS: SUCESSO]** — Plano com 7 públicos, 4 níveis de escalação, 5 rituais, 5 tipos de documentos, 4 cenários de crise.
