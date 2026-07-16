# Histórias de Usuário (User Stories) — Feature 01.2
* **Projeto:** Portal Corporativo de Gestão Tributária
* **Módulo:** M2 — Cadastro e Manutenção de Alíquotas (Entrega 1)
* **Feature Relacionada:** [01.2 — Cadastro e Manutenção de Alíquotas](./04-FEATURES.md#feature-012-cadastro-e-manutenção-de-alíquotas) [INDEX]
* **Status:** Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Criação de Nova Alíquota com Validação Automática

### 1. Descrição da História (Visão de Negócio)

Como Analista Fiscal,
Quero cadastrar uma nova alíquota no portal preenchendo um formulário padronizado com validações automáticas que impeçam erros de configuração,
Para que a alíquota entre em operação imediatamente, sem depender de revisão técnica e com garantia de que não há conflitos com alíquotas existentes [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Campos Obrigatórios): Tributo, abrangência geográfica, data de início de vigência e valor da alíquota são campos obrigatórios. O sistema não permite salvar sem o preenchimento destes campos.
* RN02 (Conflito de Vigência): O sistema varre a base antes de confirmar o cadastro. Se existir alíquota para o mesmo tributo e mesma região geográfica com interseção de períodos, o sistema bloqueia o cadastro e exibe a alíquota conflitante.
* RN03 (Integridade de Classificação): Se o campo de classificação fiscal (NCM, NBS, CClassTrib) for preenchido, o sistema valida que o código existe na base do Módulo 3 antes de confirmar o cadastro.
* RN04 (Datas de Vigência): A data de fim, quando preenchida, deve ser posterior à data de início. O sistema rejeita datas retrógradas.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Cadastro bem-sucedido de alíquota de IBS municipal

* Dado que o Analista Fiscal está na tela de "Nova Alíquota";
* Quando preencher Tributo = "IBS", UF = "SP", Município = "Santana de Parnaíba" (3547304), Data Início = "01/01/2027", Valor = "5,0%" e clicar em "Salvar";
* E não existir nenhuma alíquota de IBS para o mesmo município no mesmo período;
* Então o sistema deve confirmar o cadastro e exibir a mensagem "Alíquota cadastrada com sucesso. Vigência: 01/01/2027.";
* E a nova alíquota deve aparecer imediatamente no Painel de Alíquotas Vigentes [INDEX].

#### Cenário 2: Bloqueio por conflito de vigência

* Dado que já existe uma alíquota de ICMS para SP com vigência de 01/01/2026 a 31/12/2026;
* Quando o Analista Fiscal tentar cadastrar outra alíquota de ICMS para SP com vigência de 01/07/2026 a 30/06/2027;
* Então o sistema deve bloquear o cadastro e exibir: "Já existe uma alíquota de ICMS vigente para SP no período 01/01/2026 a 31/12/2026. Ajuste as datas ou revise a alíquota existente.";
* E exibir um link para a alíquota conflitante [INDEX].

#### Cenário 3: Rejeição por classificação fiscal inexistente

* Dado que o Analista Fiscal está cadastrando uma alíquota de IS (Imposto Seletivo);
* Quando preencher o campo NCM com "9999.99.99" e clicar em "Salvar";
* E o código "9999.99.99" não está cadastrado no Módulo 3;
* Então o sistema deve exibir: "O código NCM 9999.99.99 não está cadastrado. Cadastre a classificação fiscal antes de associar uma alíquota." [INDEX].

------------------------------
## 📝 US-02: Edição de Alíquota Existente com Preservação de Histórico

### 1. Descrição da História (Visão de Negócio)

Como Analista Fiscal,
Quero editar uma alíquota existente para ajustar seu valor, período de vigência ou classificação associada,
Para corrigir imprecisões ou refletir mudanças regulatórias, com a garantia de que o valor anterior fique registrado na trilha de auditoria [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN05 (Preservação do Histórico): A edição de uma alíquota não sobrescreve o registro original — o sistema cria uma nova versão da alíquota e preserva a versão anterior na trilha de auditoria.
* RN06 (Revalidação Completa): A edição dispara as mesmas validações do cadastro (RN02, RN03, RN04). Se a edição criar um conflito de vigência que não existia antes, o sistema bloqueia.
* RN07 (Justificativa para Alteração de Valor): Se o valor da alíquota for alterado em mais de 0,5 ponto percentual, o sistema exige o preenchimento obrigatório do campo de justificativa de negócio.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Edição simples de valor de alíquota

* Dado que existe uma alíquota de CBS com valor de 0,9%;
* Quando o Analista Fiscal editar o valor para 9,0% e preencher a justificativa "Adequação à alíquota cheia conforme LC 214/2025";
* Então o sistema deve confirmar a edição e exibir "Alíquota atualizada com sucesso.";
* E o novo valor 9,0% deve aparecer no Painel;
* E a trilha de auditoria deve registrar: valor anterior 0,9%, novo valor 9,0%, justificativa "Adequação à alíquota cheia conforme LC 214/2025" [INDEX].

#### Cenário 2: Edição que gera conflito bloqueada

* Dado que existe uma alíquota A de PIS para nacional vigente de 01/01/2026 a 30/06/2026;
* E existe uma alíquota B de PIS para nacional vigente de 01/07/2026 a 31/12/2026;
* Quando o Analista editar a data fim da alíquota A para 31/12/2026;
* Então o sistema deve bloquear a edição com a mesma mensagem de conflito do cadastro (RN02) [INDEX].

------------------------------
## 📝 US-03: Desativação de Alíquota com Justificativa Obrigatória

### 1. Descrição da História (Visão de Negócio)

Como Analista Fiscal,
Quero desativar uma alíquota que não é mais aplicável (por expiração legal, revogação ou substituição por nova regra),
Para que a alíquota seja removida da base ativa mas seu histórico seja preservado para consultas e auditorias futuras [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN08 (Desativação Lógica): A desativação não remove fisicamente o registro da alíquota — ela encerra a vigência com a data atual e marca o status como "Expirada".
* RN09 (Justificativa Obrigatória): Toda desativação exige o preenchimento do campo de justificativa de negócio, independentemente do valor ou impacto da alíquota.
* RN10 (Alerta de Transição no Período Híbrido): Durante o Período Híbrido (2029–2032), ao desativar uma alíquota do regime antigo, o sistema verifica se existe alíquota substituta no novo regime. Se não existir, exibe um alerta (não bloqueante).

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Desativação bem-sucedida com justificativa

* Dado que existe uma alíquota de ISS para Santana de Parnaíba/SP vigente;
* Quando o Analista Fiscal clicar em "Desativar", preencher a justificativa "ISS extinto para este serviço conforme transição para IBS — vigência encerrada em 31/12/2028" e confirmar;
* Então o sistema deve encerrar a vigência da alíquota e exibir "Alíquota desativada com sucesso.";
* E a alíquota deve sair do Painel de Alíquotas Vigentes (filtro padrão = vigentes);
* E deve permanecer visível quando o filtro "Status = Expirada" for aplicado [INDEX].

#### Cenário 2: Tentativa de desativação sem justificativa

* Dado que o Analista Fiscal clicou em "Desativar" em uma alíquota;
* Quando tentar confirmar a desativação sem preencher o campo de justificativa;
* Então o sistema deve impedir a confirmação e exibir "É obrigatório informar a justificativa para encerramento da vigência desta alíquota." [INDEX].

#### Cenário 3: Alerta de transição no Período Híbrido

* Dado que estamos em 2030 (Período Híbrido) e o Analista Fiscal desativa uma alíquota de ICMS para MG;
* E não existe alíquota de IBS cadastrada para MG como substituta;
* Quando o Analista confirmar a desativação;
* Então o sistema deve exibir o alerta "⚠️ Não foi encontrada alíquota de IBS substituta para esta alíquota de ICMS em MG. Consulte o Comitê Fiscal antes de prosseguir.";
* E permitir que o Analista prossiga ou cancele a operação [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
