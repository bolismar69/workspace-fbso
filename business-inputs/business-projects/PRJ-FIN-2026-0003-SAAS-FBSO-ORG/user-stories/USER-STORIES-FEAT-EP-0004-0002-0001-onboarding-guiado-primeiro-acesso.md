# User Stories: Onboarding Guiado de Primeiro Acesso

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** FEAT-EP-0004-0002 — Onboarding Guiado de Primeiro Acesso
- **Épico:** EP-0004 — Experiência do Cliente e Autoatendimento
- **Prioridade:** Must Have
- **Data-Alvo:** 30/09/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

---

## Objetivo de Negócio

Conduzir o cliente por um fluxo simples e guiado no primeiro acesso, garantindo que ele configure o essencial para começar a usar a plataforma sem precisar de ajuda do suporte.

---

## User Stories

### US-040 — Fluxo Guiado de Onboarding

**Como** Cliente no primeiro acesso, **quero** ser recebido por um fluxo guiado de onboarding que me conduza passo a passo pelas configurações iniciais obrigatórias.

**Critérios de Aceitação:**
- Ao detectar primeiro login, sistema inicia automaticamente o onboarding
- Barra de progresso visível (Passo 1 de 4, Passo 2 de 4...)
- Todas as 4 etapas são obrigatórias no primeiro acesso. Cliente avança linearmente.
- Cliente pode salvar e continuar depois. Progresso é preservado por 7 dias. Após esse período, onboarding reinicia do Passo 1.

### US-041 — Confirmação de Dados Cadastrais

**Como** Cliente no onboarding, **quero** confirmar e complementar meus dados cadastrais (razão social, nome fantasia, segmento) **para** garantir que as informações estão corretas.

**Critérios de Aceitação:**
- Dados pré-preenchidos com informações fornecidas pelo time FBSO.ORG
- Cliente confirma ou edita cada campo
- Avançar para o próximo passo salva os dados

### US-042 — Cadastro da Primeira Unidade de Negócio

**Como** Cliente no onboarding, **quero** cadastrar minha primeira Unidade de Negócio (CNPJ matriz, regime tributário, endereço) **para** começar a operar na plataforma.

**Critérios de Aceitação:**
- Formulário de cadastro de Unidade de Negócio integrado ao fluxo
- Campos: CNPJ, razão social, regime tributário (Simples, Lucro Real, Lucro Presumido), endereço
- Validação de CNPJ verifica formato e dígitos verificadores (algoritmo da Receita Federal). Validação considera apenas formato e dígitos — não consulta base da Receita Federal online.

### US-043 — Resumo do Plano Contratado

**Como** Cliente no onboarding, **quero** visualizar um resumo do meu plano contratado (nome do plano, módulos incluídos, valor) **para** entender o que está disponível para mim.

**Critérios de Aceitação:**
- Card com informações do plano: nome, módulos incluídos (ícones e nomes), valor mensal
- Informação de que novos módulos podem ser contratados futuramente
- Botão "Começar a usar" para finalizar o onboarding

### US-044 — Tela de Boas-Vindas Pós-Onboarding

**Como** Cliente, **quero** ser recebido com uma tela de boas-vindas após concluir o onboarding, com orientações sobre os próximos passos.

**Critérios de Aceitação:**
- Tela de boas-vindas com: saudação personalizada, resumo do que foi configurado
- Sugestões de próximos passos: "Convide seu time", "Cadastre seus produtos"
- Botão "Ir para o Portal" que leva ao Seletor de Módulo (FEAT-EP-0004-0004). Caso o Dashboard do Cliente (FEAT-EP-0004-0003, Should Have) esteja implementado, o destino será o dashboard.

---

## Regras de Negócio

| ID | Regra |
|----|-------|
| **RN-FEAT-EP-0004-0002-0001** | Onboarding é obrigatório no primeiro acesso; não pode ser pulado |
| **RN-FEAT-EP-0004-0002-0002** | Primeira Unidade de Negócio cadastrada no onboarding é automaticamente definida como Matriz |
| **RN-FEAT-EP-0004-0002-0003** | Onboarding só é considerado concluído quando todos os passos obrigatórios são finalizados |
| **RN-FEAT-EP-0004-0002-0004** | Tenant só muda para status "Ativo" após conclusão do onboarding |

---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | Fluxo completo de 4 passos executado sem erros | Percurso completo com barra de progresso |
| F2 | Cliente consegue salvar e retomar onboarding | Interrupção no passo 2 e retomada no passo 2 |
| F3 | Tenant muda para "Ativo" somente após conclusão total | Verificação de status antes e depois do onboarding |
| F4 | Primeira Unidade de Negócio definida como Matriz automaticamente | Verificação do vínculo hierárquico |

### Casos de Borda

- Cliente que abandona o onboarding e faz login novamente é redirecionado automaticamente para o passo onde parou (desde que dentro do prazo de 7 dias de preservação do progresso).


---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*

---
👷 *Revisão técnica realizada pelo Agente: CaveMan em 15/07/2026, conforme User Story Review. Foram utilizados os skills: caveman-review.*
