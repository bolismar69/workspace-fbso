# User Stories: Gestão de Unidades de Negócio

- **Projeto:** PRJ-FIN-2026-0003-SAAS-FBSO-ORG
- **Feature:** FEAT-EP-0004-0005 — Gestão de Unidades de Negócio
- **Épico:** EP-0004 — Experiência do Cliente e Autoatendimento
- **Prioridade:** Must Have
- **Data-Alvo:** 15/10/2026
- **Versão:** 1.1 — Revisada conforme User Story Review (15/07/2026)
- **Origem:** [04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md](../04-FEATURES-PRJ-FIN-2026-0003-SAAS-FBSO-ORG.md)

---

## Objetivo de Negócio

Permitir que o cliente cadastre e gerencie suas filiais e CNPJs de forma autônoma, com estrutura hierárquica clara (Matriz/Filial).

---

## User Stories

### US-050 — Lista Hierárquica de Unidades de Negócio

**Como** Cliente, **quero** visualizar a lista das minhas Unidades de Negócio organizadas hierarquicamente (Matriz no topo, filiais recuadas abaixo) **para** entender a estrutura da minha empresa.

**Critérios de Aceitação:**
- Visualização em cards com recuo visual para filiais
- Cada card exibe: razão social, CNPJ (mascarado), regime tributário, status
- Indicador visual de Matriz vs. Filial

### US-051 — Cadastro de Nova Unidade de Negócio (Filial)

**Como** Cliente, **quero** cadastrar uma nova Unidade de Negócio (filial) informando CNPJ, razão social, regime tributário e definindo a qual unidade ela se vincula (Matriz ou outra filial).

**Critérios de Aceitação:**
- Formulário com campos obrigatórios: CNPJ, razão social, regime tributário
- Seletor de unidade pai (Matriz ou filial existente)
- Validação de duplicidade de CNPJ ativo para o mesmo tenant

### US-052 — Edição de Dados da Unidade de Negócio

**Como** Cliente, **quero** editar os dados de uma Unidade de Negócio (razão social, regime tributário, endereço) **para** manter as informações atualizadas.

**Critérios de Aceitação:**
- Tela de edição acessível a partir do card da unidade
- Campos editáveis: razão social, regime tributário, endereço
- CNPJ não pode ser alterado após o cadastro

### US-053 — Desativação de Unidade de Negócio

**Como** Cliente, **quero** desativar uma Unidade de Negócio que não está mais em operação, mantendo seu histórico no sistema.

**Critérios de Aceitação:**
- Botão "Desativar" no card da unidade
- Confirmação exigida antes da desativação
- Unidade desativada não aparece nos seletores para novos cadastros
- Dados históricos permanecem acessíveis para consulta

### Casos de Borda

- Tentativa de desativar a Matriz quando há filiais ativas exibe alerta: 'Existem filiais ativas vinculadas a esta unidade. Reatribua ou desative as filiais antes de desativar a Matriz.'


### US-054 — Seletor de Unidade de Negócio no Topo do Portal

**Como** Cliente, **quero** usar o seletor de Unidade de Negócio no topo do portal para alternar entre minhas filiais e visualizar os dados específicos de cada uma.

**Critérios de Aceitação:**
- Seletor dropdown no topo do portal, ao lado do Seletor de Módulo
- Exibe apenas unidades que o usuário tem permissão para acessar
- Ao trocar de unidade, dados exibidos nas telas são filtrados automaticamente

---

## Regras de Negócio

| ID | Regra |
|----|-------|
| **RN-FEAT-EP-0004-0005-0001** | CNPJ deve ser único entre Unidades de Negócio ativas do mesmo tenant (desativação lógica libera o CNPJ para reúso) |
| **RN-FEAT-EP-0004-0005-0002** | Uma unidade desativada não pode ser definida como "pai" de novas filiais |
| **RN-FEAT-EP-0004-0005-0003** | A primeira unidade cadastrada (durante onboarding) é automaticamente a Matriz |
| **RN-FEAT-EP-0004-0005-0004** | Não há limite de níveis hierárquicos (Matriz → Filial → Sub-filial) |
| **RN-FEAT-EP-0004-0005-0005** | Seletor de Unidade de Negócio reflete apenas as unidades que o usuário tem permissão |
| **RN-FEAT-EP-0004-0005-0006** | Não é permitido criar ciclos na hierarquia (ex: Unidade A → Unidade B → Unidade A). Sistema valida a cadeia hierárquica antes de salvar. |

### Casos de Borda (US-053)

- Unidade desativada não pode ser definida como 'pai' de novas unidades. Filiais existentes de unidade desativada permanecem vinculadas mas exibem indicador 'Matriz Inativa'.


---

## Critérios de Aceitação da Feature

| # | Critério | Evidência |
|---|----------|-----------|
| F1 | Lista hierárquica com Matriz e filiais visualmente distintas | Print com estrutura Matriz + 3 filiais |
| F2 | Cadastro de filial com validação de CNPJ duplicado ativo | Tentativa de cadastro com CNPJ já existente |
| F3 | Seletor de Unidade de Negócio filtra dados em tempo real | Troca de unidade e verificação dos dados exibidos |
| F4 | Desativação preserva histórico e libera CNPJ para reúso | Recadastro do mesmo CNPJ após desativação |

---

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices. Revisão 1.1 baseada no User Story Review (15/07/2026) — skills: caveman, caveman-review.*

---
👷 *Revisão técnica realizada pelo Agente: CaveMan em 15/07/2026, conforme User Story Review. Foram utilizados os skills: caveman-review.*
