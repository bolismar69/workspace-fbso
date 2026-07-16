# Definition of Done (DoD)

- **Projeto:** Portal Corporativo de Gestão Tributária — Autonomia do Time de Finanças na Administração de Impostos
- **Código:** PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS
- **Versão:** 1.0
- **Atualizado:** 2026-07-08
- **Objetivo:** Estabelecer critérios objetivos, verificáveis e não-ambíguos para determinar quando uma user story, feature ou entrega está concluída. Este documento é o contrato compartilhado entre time de negócios, Product Owner e time de Finanças (usuários finais).

---

## 1. Princípios

1. **Binário, não subjetivo.** Todo critério é verificável por checklist objetivo. Nenhum critério depende de "parece bom" ou "acho que está pronto."
2. **Acumulativo.** Uma user story só está DONE quando TODOS os critérios aplicáveis são satisfeitos — não quando "a maioria" está.
3. **Focado em resultado de negócio.** Os critérios medem o valor entregue ao time de Finanças, não métricas técnicas de implementação.
4. **Entregas têm DoD cumulativo.** A DoD de uma feature inclui a DoD de todas as suas user stories. A DoD de uma entrega inclui a DoD de todas as suas features.
5. **Validação pelo usuário final.** Nenhuma funcionalidade está concluída até que o time de Finanças (ou seu representante designado) a valide em ambiente de homologação.

---

## 2. Níveis de DoD

```
DoD de USER STORY  →  menor unidade de entrega (1 necessidade do time fiscal)
DoD de FEATURE     →  conjunto de user stories relacionadas (1 módulo do portal)
DoD de ENTREGA     →  conjunto de features (1 marco do projeto: Gestão Básica, Governança, Operações em Escala, Portal Completo)
```

Cada nível herda os critérios do nível anterior e adiciona os seus próprios.

---

## 3. DoD de USER STORY

Uma user story individual está **DONE** quando TODOS os critérios abaixo são satisfeitos.

### 3.1 Funcionalidade

| # | Critério | Como verificar | Evidência |
|---|----------|----------------|-----------|
| F1 | A funcionalidade entrega exatamente o descrito na user story — nem mais, nem menos | Demonstração para o PO e Gerente Fiscal comparando com a especificação | Ata de demo com aprovação registrada |
| F2 | A funcionalidade está disponível no portal e acessível ao perfil de usuário correto | Login com cada perfil aplicável e verificação de visibilidade e permissões | Checklist de perfis × funcionalidades verificados |
| F3 | Mensagens de erro e validações são apresentadas em linguagem clara para o usuário de negócio (não linguagem técnica) | Provocar cada condição de erro e verificar a clareza da mensagem | Registro de testes de erro com avaliação de clareza |

### 3.2 Validações de Negócio

| # | Critério | Como verificar | Evidência |
|---|----------|----------------|-----------|
| V1 | As regras de validação de negócio descritas na user story estão implementadas e impedem ações inválidas | Executar cada cenário de validação (ex: tentar criar alíquota conflitante) e verificar bloqueio | Lista de cenários testados com resultado |
| V2 | As regras de validação foram revisadas e aprovadas pelo Comitê Fiscal (quando aplicável) | Aprovação registrada do Comitê Fiscal | Registro de aprovação |
| V3 | Casos de borda identificados na user story (ex: datas limite, valores zero, municípios sem alíquota) são tratados adequadamente | Testar cada caso de borda | Lista de casos de borda testados |

### 3.3 Trilha de Auditoria

| # | Critério | Como verificar | Evidência |
|---|----------|----------------|-----------|
| A1 | Toda operação que altera dados fiscais gera registro de auditoria com: usuário, data/hora, entidade, valor anterior, valor novo | Realizar uma alteração e verificar o registro gerado na Linha do Tempo | Print do registro de auditoria com todos os campos |
| A2 | O registro de auditoria está acessível ao perfil de Auditor/Controller | Login com perfil de Auditor e consulta à Linha do Tempo | Confirmação de visibilidade |

### 3.4 Aceitação do Usuário

| # | Critério | Como verificar | Evidência |
|---|----------|----------------|-----------|
| U1 | Pelo menos um analista fiscal (usuário real) testou a funcionalidade em ambiente de homologação e aprovou | Sessão de teste com analista fiscal designado | Feedback registrado e aprovação assinada |
| U2 | O PO validou que a funcionalidade atende aos critérios de aceite descritos na user story | Revisão do PO contra a especificação | Critérios de aceite marcados como ✅ |
| U3 | A documentação de uso da funcionalidade está disponível para o time de Finanças (guia rápido, tooltip ou help integrado ao portal) | Acessar a documentação no portal ou no repositório de materiais de treinamento | Documentação publicada |

---

## 4. DoD de FEATURE

Uma feature (conjunto de user stories relacionadas, tipicamente correspondendo a um Módulo do portal) está **DONE** quando, ALÉM de todas as suas user stories atenderem à DoD de USER STORY:

| # | Critério adicional de FEATURE | Como verificar | Evidência |
|---|-------------------------------|----------------|-----------|
| F1 | Todas as user stories da feature estão DONE (DoD de USER STORY satisfeita para cada uma) | Checklist contra o backlog da feature | Todas as user stories marcadas como concluídas |
| F2 | O fluxo completo do módulo funciona de ponta a ponta sem interrupções | Executar o cenário principal do módulo do início ao fim (ex: cadastrar alíquota → visualizar no painel → ver na trilha de auditoria) | Roteiro de teste ponta a ponta executado com sucesso |
| F3 | Nenhuma regressão: funcionalidades de módulos já entregues continuam operando normalmente | Verificação dos módulos já homologados após a implantação do novo módulo | Checklist de regressão executado |
| F4 | O Gerente Fiscal ou Controller validou a feature completa em ambiente de homologação | Sessão formal de homologação com o responsável designado | Termo de aceite da feature assinado |
| F5 | O time de Finanças recebeu treinamento ou demonstração do módulo completo | Sessão de treinamento registrada com lista de presença | Lista de presença e material de treinamento |
| F6 | Os KPIs afetados por esta feature (conforme MATRIZ-KPI.md) têm baseline definida para medição pós-implantação | Verificar se os KPIs aplicáveis estão identificados e com valor de referência pré-implantação registrado | Baseline de KPIs documentada |

---

## 5. DoD de ENTREGA

Uma entrega do projeto está **DONE** quando, ALÉM de todas as suas features atenderem à DoD de FEATURE:

| # | Critério adicional de ENTREGA | Como verificar | Evidência |
|---|-------------------------------|----------------|-----------|
| E1 | Todas as features planejadas para a entrega estão DONE | Checklist contra o escopo da entrega no Project Charter | 100% das features concluídas |
| E2 | Os critérios de homologação de negócio (UAT) descritos no 02-BUSINESS-REQUIREMENTS.md para esta entrega foram satisfeitos | Execução dos cenários UAT descritos na Seção 5 do documento de requisitos | Cenários UAT executados e aprovados |
| E3 | Os KPIs da entrega (conforme MATRIZ-KPI.md) foram medidos e estão dentro da meta ou com plano de ação para desvios | Medição dos KPIs aplicáveis à entrega | Dashboard de KPIs atualizado |
| E4 | O Sponsor (CFO) ou seu representante foi informado da conclusão da entrega e dos resultados dos KPIs | Comunicação formal ao CFO com sumário executivo | Status report executivo enviado e acknowledged |
| E5 | A Matriz de Riscos do projeto (Project Charter, Seção 7) foi revisada: riscos mitigados foram fechados, novos riscos foram identificados | Revisão da matriz de riscos em reunião de governança | Matriz de riscos atualizada |
| E6 | As lições aprendidas da entrega foram documentadas para aplicação nas fases seguintes | Sessão de retrospectiva com time de negócios e PO | Documento de lições aprendidas |

---

## 6. O que NÃO faz parte da DoD

Estes itens são importantes mas NÃO são critério de DONE para uma user story ou feature:

| Item | Onde é tratado |
|------|----------------|
| Desempenho e capacidade de resposta do portal | Monitoramento contínuo pela área de operações — critério de serviço, não de entrega |
| Treinamento de novos usuários que ingressarem após o go-live | Programa de onboarding do time de Finanças |
| Evoluções e melhorias além do escopo original da user story | Backlog de produto — tratado como novas user stories |
| Integração com sistemas externos não previstos no escopo | Fase de expansão futura (pós-Entrega 4) |
| Disponibilidade 24/7 e recuperação de desastres | Plano de continuidade de negócios corporativo |

---

## 7. Uso com Agentes de IA

### 7.1 Como o agente usa a DoD

O agente de IA deve:
1. Carregar este arquivo junto com os documentos de negócio do projeto no início da execução
2. Para cada user story, percorrer a checklist da Seção 3 e marcar cada item como ✅ ou ❌
3. Registrar o resultado em documento de evidência
4. Se algum item falhar, a user story NÃO está DONE — corrigir e re-verificar

### 7.2 Como o humano usa a DoD

- **Product Owner:** audita as user stories contra esta DoD — itens ❌ são bloqueantes para aceitação
- **Gerente Fiscal:** verifica F1–F3 e U1–U3 (funcionalidade e aceitação do usuário)
- **Controladoria:** verifica A1–A2 (trilha de auditoria)
- **Comitê Fiscal:** verifica V2 (validações de negócio aprovadas)
- **CFO / PMO:** verifica E1–E6 (conclusão da entrega)

---

## 8. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | 2026-07-08 | Criação inicial: DoD de USER STORY (9 critérios), FEATURE (6 critérios), ENTREGA (6 critérios) com foco em resultado de negócio | Time de Negócios |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: agile-ba-practices, acceptance-criteria.*
