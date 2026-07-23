# Definition of Done (DoD)

- **Projeto:** FBSO Platform — Portal Administrativo SaaS
- **Código:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Versão:** 1.0
- **Atualizado:** 2026-07-13
- **Objetivo:** Estabelecer critérios objetivos, verificáveis e não-ambíguos para determinar quando uma user story, feature ou entrega está concluída. Este documento é o contrato compartilhado entre time de negócios, Product Owner e time administrativo da FBSO.ORG.

---

## 1. Princípios

1. **Binário, não subjetivo.** Todo critério é verificável por checklist objetivo. Nenhum critério depende de "parece bom" ou "acho que está pronto."
2. **Acumulativo.** Uma user story só está DONE quando TODOS os critérios aplicáveis são satisfeitos — não quando "a maioria" está.
3. **Focado em resultado de negócio.** Os critérios medem o valor entregue ao time administrativo da FBSO.ORG e aos clientes do portal, não métricas técnicas de implementação.
4. **Entregas têm DoD cumulativo.** A DoD de uma feature inclui a DoD de todas as suas user stories. A DoD de uma entrega inclui a DoD de todas as suas features.
5. **Validação pelo usuário final.** Nenhuma funcionalidade está concluída até que o usuário-alvo (time interno FBSO.ORG ou cliente do portal) a valide em ambiente de homologação.

---

## 2. Níveis de DoD

```
DoD de USER STORY  →  menor unidade de entrega (1 necessidade de negócio)
DoD de FEATURE     →  conjunto de user stories relacionadas (1 módulo do portal)
DoD de ENTREGA     →  conjunto de features (1 marco do projeto: D1 a D7)
```

Cada nível herda os critérios do nível anterior e adiciona os seus próprios.

---

## 3. DoD de USER STORY

Uma user story individual está **DONE** quando TODOS os critérios abaixo são satisfeitos.

### 3.1 Funcionalidade

| # | Critério | Como verificar | Evidência |
|---|----------|----------------|-----------|
| F1 | A funcionalidade entrega exatamente o descrito na user story — nem mais, nem menos | Demonstração para o PO comparando com a especificação no FEATURES.md | Ata de demo com aprovação registrada |
| F2 | A funcionalidade está disponível no portal e acessível ao perfil de usuário correto | Login com cada perfil aplicável e verificação de visibilidade e permissões conforme RN10-01 | Checklist de perfis × funcionalidades verificados |
| F3 | Mensagens de erro e validações são apresentadas em linguagem clara para o usuário de negócio (não linguagem técnica) | Provocar cada condição de erro e verificar a clareza da mensagem | Registro de testes de erro com avaliação de clareza |

### 3.2 Validações de Negócio

| # | Critério | Como verificar | Evidência |
|---|----------|----------------|-----------|
| V1 | As regras de negócio descritas na user story estão implementadas e impedem ações inválidas (ex: CNPJ duplicado ativo, transição de status inválida, usuário sem unidade vinculada) | Executar cada cenário de validação descrito nas RNs do FEATURES.md e verificar bloqueio | Lista de cenários testados com resultado |
| V2 | As regras de validação foram revisadas e aprovadas pelo Dono do Produto (quando envolverem regras de negócio complexas) | Aprovação registrada do PO | Registro de aprovação |
| V3 | Casos de borda identificados na user story (ex: tenant sem assinatura, unidade desativada com produtos vinculados, revogação de acesso com usuário logado) são tratados adequadamente | Testar cada caso de borda | Lista de casos de borda testados |

### 3.3 Trilha de Auditoria

| # | Critério | Como verificar | Evidência |
|---|----------|----------------|-----------|
| A1 | Toda ação administrativa que altera dados sensíveis (ativação/suspensão de tenant, mudança de plano, alteração de permissões) gera registro de auditoria com: responsável, data/hora, entidade, valor anterior, valor novo | Realizar uma alteração e verificar o registro gerado no Histórico de Auditoria (F02-05) | Print do registro de auditoria com todos os campos |
| A2 | O registro de auditoria está acessível ao perfil de Administrador FBSO.ORG e ao Administrador do Tenant (quando aplicável) | Login com cada perfil e consulta ao histórico | Confirmação de visibilidade por perfil |

### 3.4 Aceitação do Usuário

| # | Critério | Como verificar | Evidência |
|---|----------|----------------|-----------|
| U1 | Pelo menos um representante do público-alvo (time interno FBSO.ORG ou cliente early adopter) testou a funcionalidade em ambiente de homologação e aprovou | Sessão de teste com usuário designado | Feedback registrado e aprovação assinada |
| U2 | O PO validou que a funcionalidade atende aos critérios de aceite descritos na user story (FEATURES.md) | Revisão do PO contra a especificação | Critérios de aceite marcados como ✅ |
| U3 | A documentação de uso da funcionalidade está disponível (guia rápido, tooltip ou help integrado ao portal) para funcionalidades do Portal do Cliente (EP-04) | Acessar a documentação no portal ou no repositório de materiais | Documentação publicada |

---

## 4. DoD de FEATURE

Uma feature (conjunto de user stories relacionadas, correspondendo a uma linha do FEATURES.md) está **DONE** quando, ALÉM de todas as suas user stories atenderem à DoD de USER STORY:

| # | Critério adicional de FEATURE | Como verificar | Evidência |
|---|-------------------------------|----------------|-----------|
| F1 | Todas as user stories da feature estão DONE (DoD de USER STORY satisfeita para cada uma) | Checklist contra o backlog da feature no FEATURES.md | Todas as user stories marcadas como concluídas |
| F2 | O fluxo completo da feature funciona de ponta a ponta sem interrupções | Executar o cenário principal da feature do início ao fim (ex: criar tenant → ativar → vincular plano → cliente acessa portal) | Roteiro de teste ponta a ponta executado com sucesso |
| F3 | Nenhuma regressão: funcionalidades de features já entregues continuam operando normalmente | Verificação das features já homologadas após a implantação da nova feature | Checklist de regressão executado |
| F4 | O stakeholder principal da feature (conforme STAKEHOLDER-MAP.md) validou a feature completa em ambiente de homologação | Sessão formal de homologação com o responsável designado | Termo de aceite da feature assinado |
| F5 | Os KPIs afetados por esta feature (conforme MATRIZ-KPI.md) têm baseline definida para medição pós-implantação | Verificar se os KPIs aplicáveis estão identificados e com valor de referência pré-implantação registrado | Baseline de KPIs documentada |

---

## 5. DoD de ENTREGA

Uma entrega do projeto (D1 a D7, conforme Project Charter) está **DONE** quando, ALÉM de todas as suas features atenderem à DoD de FEATURE:

| # | Critério adicional de ENTREGA | Como verificar | Evidência |
|---|-------------------------------|----------------|-----------|
| E1 | Todas as features planejadas para a entrega estão DONE | Checklist contra o escopo da entrega no Project Charter | 100% das features concluídas |
| E2 | Os critérios de homologação de negócio (UAT) descritos no 02-BRD-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md para esta entrega foram satisfeitos | Execução dos cenários UAT descritos na Seção 6 do BRD | Cenários UAT executados e aprovados |
| E3 | Os KPIs da entrega (conforme MATRIZ-KPI.md) foram medidos e estão dentro da meta ou com plano de ação para desvios | Medição dos KPIs aplicáveis à entrega | Dashboard de KPIs atualizado |
| E4 | O Sponsor (Diretoria FBSO.ORG) ou seu representante foi informado da conclusão da entrega e dos resultados dos KPIs | Comunicação formal ao Sponsor com sumário executivo | Status report executivo enviado e acknowledged |
| E5 | A Matriz de Riscos do projeto (Project Charter, Seção 8) foi revisada: riscos mitigados foram fechados, novos riscos foram identificados | Revisão da matriz de riscos em reunião de governança | Matriz de riscos atualizada |
| E6 | As lições aprendidas da entrega foram documentadas para aplicação nas fases seguintes | Sessão de retrospectiva com time de negócios e PO | Documento de lições aprendidas |

---

## 6. O que NÃO faz parte da DoD

Estes itens são importantes mas NÃO são critério de DONE para uma user story ou feature:

| Item | Onde é tratado |
|------|----------------|
| Desempenho e capacidade de resposta do portal | Monitoramento contínuo pela área de operações — critério de serviço, não de entrega. Nota: embora não seja critério de DoD de negócio, a validação técnica de performance é coberta pelo TEST_PLAN.md do time de engenharia |
| Treinamento de novos usuários que ingressarem após o go-live | Programa de onboarding do cliente |
| Evoluções e melhorias além do escopo original da user story | Backlog de produto — tratado como novas user stories |
| Funcionalidades dos módulos-produto (Tributali-Engine, Storekeeper Portal) | Fases futuras do programa FBSO Platform (fora do escopo deste projeto) |
| Disponibilidade 24/7 e recuperação de desastres | Plano de continuidade de negócios corporativo. Nota: embora não seja critério de DoD de negócio, os NFRs do BRD (BR-NFR05 a BR-NFR08) exigem validação técnica de disponibilidade |

---

## 7. Uso com Agentes de IA

### 7.1 Como o agente usa a DoD

Esta seção descreve o processo de uso da DoD por agentes de IA e humanos. É uma referência de processo, não um critério de aceitação. O agente de IA deve:
1. Carregar este arquivo junto com os documentos de negócio do projeto no início da execução
2. Para cada user story, percorrer a checklist da Seção 3 e marcar cada item como ✅ ou ❌
3. Registrar o resultado em documento de evidência
4. Se algum item falhar, a user story NÃO está DONE — corrigir e re-verificar

### 7.2 Como o humano usa a DoD

- **Product Owner:** audita as user stories contra esta DoD — itens ❌ são bloqueantes para aceitação
- **Líder Administrativo:** verifica F1–F3 e U1–U3 (funcionalidade e aceitação do usuário interno)
- **Líder Comercial:** verifica aderência às regras de plano e assinatura
- **Dono do Produto:** verifica V1–V3 (validações de negócio)
- **Diretoria / PMO:** verifica E1–E6 (conclusão da entrega)

---

## 8. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | 2026-07-13 | Criação inicial: DoD de USER STORY (11 critérios), FEATURE (5 critérios), ENTREGA (6 critérios) com foco em resultado de negócio | Time de Negócios |
| 1.1 | 2026-07-15 | Revisão caveman: correção da contagem de critérios (12→11), adição de nota técnica na Seção 6, atualização da Seção 7 | Caveman Review |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: agile-ba-practices, acceptance-criteria.*
🔍 *Revisado pelo skill caveman-review em 15/07/2026. Ajustes aplicados: contagem de critérios corrigida (12→11), Seção 6 com ressalva de validação técnica (NFRs), Seção 7 contextualizada como referência de processo.*
