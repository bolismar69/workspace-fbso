# Histórias de Usuário (User Stories) — Feature 03.1
* **Projeto:** Portal Corporativo de Gestão Tributária
* **Módulo:** Funcionalidade Transversal — Fluxos de Aprovação (Entrega 3)
* **Feature Relacionada:** [03.1 — Fluxos de Aprovação para Alterações de Alto Impacto](./04-FEATURES.md#feature-031-fluxos-de-aprovação-para-alterações-de-alto-impacto) [INDEX]
* **Status:** Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Disparo Automático do Fluxo de Aprovação por Patamar de Materialidade

### 1. Descrição da História (Visão de Negócio)

Como Analista Fiscal,
Quero que o portal identifique automaticamente quando uma alteração de alíquota ultrapassa o patamar de materialidade definido pelo Comitê Fiscal e a encaminhe para o fluxo de aprovação em duas etapas,
Para que eu possa realizar meu trabalho sem precisar decidir subjetivamente quais alterações precisam de aprovação adicional [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Cálculo do Impacto): No momento de salvar uma criação ou edição de alíquota, o portal estima o faturamento mensal impactado com base no histórico de operações na região geográfica e classificação fiscal afetadas. Se o valor estimado for igual ou superior ao patamar de materialidade, o fluxo de aprovação é acionado automaticamente.
* RN02 (Patamar Parametrizável): O valor do patamar de materialidade é definido pelo Comitê Fiscal e pode ser ajustado por um Administrador Fiscal no módulo de parametrização. O valor padrão inicial é R$ 100.000 de faturamento mensal estimado impactado.
* RN03 (Indicador Visual no Formulário): Durante o preenchimento do formulário de cadastro/edição, o portal exibe um indicador em tempo real: "Impacto estimado: R$ [X]/mês — [Fluxo Normal / Requer Aprovação]".

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Alteração abaixo do patamar segue fluxo normal

* Dado que o patamar de materialidade está definido em R$ 100.000/mês;
* Quando o Analista Fiscal cadastra uma alíquota de IBS para um município com faturamento mensal estimado de R$ 45.000;
* Então o sistema exibe "Impacto estimado: R$ 45.000/mês — Fluxo Normal";
* E a alíquota entra em vigor imediatamente ao ser salva, sem necessidade de aprovação [INDEX].

#### Cenário 2: Alteração acima do patamar dispara fluxo de aprovação

* Dado que o patamar de materialidade está definido em R$ 100.000/mês;
* Quando o Analista Fiscal edita uma alíquota de CBS nacional (impacto estimado: R$ 2.500.000/mês);
* Então o sistema exibe "Impacto estimado: R$ 2.500.000/mês — Requer Aprovação";
* E ao salvar, a alíquota entra no status "Pendente de Aprovação" em vez de "Vigente";
* E um alerta é enviado aos Administradores Fiscais: "Nova alteração de alto impacto aguardando aprovação" [INDEX].

------------------------------
## 📝 US-02: Aprovação e Rejeição de Alterações de Alto Impacto

### 1. Descrição da História (Visão de Negócio)

Como Administrador Fiscal,
Quero revisar as alterações de alto impacto pendentes de aprovação, com visibilidade completa do impacto estimado e da justificativa do proponente,
Para aprovar ou rejeitar cada alteração com base em critérios de negócio, registrando formalmente a decisão na trilha de auditoria [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN04 (Fila de Aprovação): O portal mantém uma fila de alterações pendentes de aprovação, acessível apenas aos perfis Administrador Fiscal e Controller. A fila é ordenada por data da proposta (mais antiga primeiro) e destaca itens com mais de 3 dias úteis aguardando.
* RN05 (Aprovação com Registro): Ao aprovar, o Administrador Fiscal pode adicionar um comentário (opcional). A aprovação é registrada na trilha de auditoria como evento distinto, vinculado ao evento da proposta original.
* RN06 (Rejeição com Justificativa Obrigatória): Ao rejeitar, o Administrador Fiscal deve preencher obrigatoriamente a justificativa da rejeição. A alíquota retorna ao status "Rascunho" para que o Analista Fiscal proponente a revise.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Aprovação de alteração pendente

* Dado que existe uma alteração de alíquota de CBS pendente de aprovação desde 14/07/2026;
* Quando o Administrador Fiscal "Ana" (diferente do proponente) acessa a fila, revisa o impacto (R$ 2.500.000/mês) e a justificativa ("Adequação à alíquota cheia"), e clica em "Aprovar";
* Então a alíquota entra em vigor imediatamente;
* E a trilha de auditoria registra: "Alteração aprovada por Ana em [data/hora]. Comentário: [comentário, se houver]." [INDEX].

#### Cenário 2: Rejeição com justificativa

* Dado que existe uma alteração de alíquota pendente de aprovação;
* Quando o Controller "Roberto" revisa e identifica que a alíquota proposta conflita com um benefício fiscal vigente, preenche a justificativa "Conflito com benefício fiscal de Santana de Parnaíba — necessário revisar enquadramento" e clica em "Rejeitar";
* Então a alíquota retorna ao status "Rascunho";
* E o Analista Fiscal proponente recebe uma notificação: "Sua proposta de alteração de [alíquota] foi rejeitada. Justificativa: Conflito com benefício fiscal de Santana de Parnaíba — necessário revisar enquadramento." [INDEX].

------------------------------
## 📝 US-03: Gestão do Patamar de Materialidade

### 1. Descrição da História (Visão de Negócio)

Como Comitê Fiscal (representado pelo Administrador Fiscal),
Quero ajustar o valor do patamar de materialidade que dispara o fluxo de aprovação,
Para refletir mudanças no perfil de risco da companhia ou orientações da Controladoria, sem depender de intervenção técnica [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN07 (Acesso Restrito à Parametrização): Apenas o perfil Administrador Fiscal pode acessar a tela de parametrização do patamar de materialidade.
* RN08 (Registro de Alteração do Patamar): Toda alteração do valor do patamar é registrada na trilha de auditoria com: usuário, data/hora, valor anterior, novo valor e justificativa obrigatória.
* RN09 (Valor Mínimo): O patamar não pode ser inferior a R$ 10.000. Tentativas de definir valor abaixo deste piso são rejeitadas.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Ajuste do patamar com justificativa

* Dado que o patamar atual é R$ 100.000;
* Quando o Administrador Fiscal acessa "Parametrização > Patamar de Materialidade", altera o valor para R$ 50.000, preenche a justificativa "Redução por recomendação da Controladoria — maior controle durante a transição para alíquotas cheias de IBS" e clica em "Salvar";
* Então o sistema deve confirmar: "Patamar de materialidade alterado para R$ 50.000/mês.";
* E todas as novas alterações passam a usar o novo patamar imediatamente [INDEX].

#### Cenário 2: Rejeição de valor abaixo do mínimo

* Dado que o Administrador Fiscal tenta definir o patamar como R$ 5.000;
* Quando clica em "Salvar";
* Então o sistema deve exibir: "O patamar de materialidade não pode ser inferior a R$ 10.000." [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
