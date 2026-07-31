# PROMPT-GATE-SOURCING-FACTORY-BIDDING-FACTORY-NOTIFICATION
## Contexto
Este prompt implementa o **GATE de Validação do FACTORY-NOTIFICATION** — Fase 7. O GATE audita criticamente o artefato, verificando critérios de qualidade e conformidade DTA.
**Postura do GATE:** Cético e rigoroso. Cada NC deve ser específica, localizada e acionável.
**Propósito:** Todas fábricas notificadas com arquivos em `notifications/`. Padrão: `FACTORY-NOTIFICATION-{NOME}.md`. **Nomes de arquivo NUNCA revelam status** (selecionada/rejeitada) — informação confidencial consta apenas no conteúdo.
## Dimensões de Validação
| 1.1 | 100% Notificadas | Todas fábricas da F3 têm notificação com Data Envio, Canal e E-mail preenchidos |
| 1.2 | Tom Adequado | Vencedora: positivo+next steps. Rejeitadas: respeitoso+feedback construtivo |
| 1.3 | Motivo Específico | Cada rejeitada tem motivo técnico exato da F5 |
| 2.1 | Próximos Passos | Kickoff, contrato, acesso |
| 2.2 | Recomendações | Para futuras participações |
## Formato de Saída
### 🚨 NÃO COMPLIANCE — Para cada NC: ID-CONFLITO, Localização, Problema, Impacto, Sugestão
### ✅ PRÉ-COMPLIANCE — 3 perguntas obrigatórias; se Sim/Não/Não → COMPLIANCE
## Skills
| 1 | `gap-analysis` | Detecção de gaps | 2 | `requirements-validation` | Validação de critérios |
🤖 *Fase 7 GATE*
