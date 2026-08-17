# Histórias de Usuário (User Stories) — Feature 03.1 (Onda 2)
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Módulo: ERP SAP MM — Compras e Suprimentos (Onda 2)
- Feature Relacionada: 02.03.1 — Auditoria Fiscal de Entrada e Bloqueio de Créditos (Procure-to-Pay) [INDEX]
- Status: Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Validação de Regularidade Fiscal do Fornecedor no Ato da Entrada de Mercadoria

### 1. Descrição da História (Visão de Negócio)

Como Analista de Suprimentos e Compliance Fiscal,
Quero que o SAP valide automaticamente a regularidade fiscal do fornecedor e a legitimidade do crédito de CBS/IBS no momento da entrada da mercadoria,
Para que a empresa só se aproprie de créditos tributários legítimos e evite riscos de autuação por aproveitamento indevido no Lucro Real [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Checklist de Validação do Fornecedor): No recebimento fiscal da mercadoria, o SAP deve verificar automaticamente: (a) se o fornecedor está com cadastro ativo e qualificado no programa de compliance fiscal da empresa, (b) se o regime tributário do fornecedor permite a geração de créditos de CBS/IBS para o adquirente, (c) se a nota fiscal de entrada possui os campos de CBS/IBS devidamente destacados e calculados.
* RN02 (Bloqueio de Crédito por Inconformidade): Se qualquer item do checklist falhar, o SAP deve bloquear o lançamento do crédito na conta "Impostos a Recuperar" e direcionar o valor para uma conta transitória de "Créditos em Análise — Fornecedor Não Qualificado".
* RN03 (Notificação ao Fornecedor): Quando um crédito for bloqueado, o SAP deve gerar automaticamente uma comunicação para o fornecedor informando o motivo do bloqueio e as ações necessárias para regularização (ex: "Fornecedor: regularize seu cadastro no programa de compliance fiscal para liberação dos créditos de CBS/IBS ao adquirente").
* RN04 (Reanálise Periódica de Créditos Bloqueados): Créditos na conta transitória devem ser reanalisados a cada 30 dias. Se o fornecedor se regularizar no período, o crédito é transferido para "Impostos a Recuperar". Caso contrário, permanece bloqueado.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Fornecedor qualificado — crédito liberado automaticamente

* Dado que uma nota fiscal de entrada (NF-e #9012) foi emitida por um fornecedor qualificado no programa de compliance fiscal;
* E a NF-e destaca CBS de R$ 4.400,00 e IBS de R$ 8.850,00 corretamente calculados;
* E o fornecedor é optante pelo Lucro Real (regime que gera crédito integral para o adquirente);
* Quando o SAP MM processar o recebimento fiscal da mercadoria;
* Então o sistema deve validar todos os itens do checklist com sucesso;
* E lançar os créditos automaticamente:
  - Débito: Impostos a Recuperar — CBS: R$ 4.400,00
  - Débito: Impostos a Recuperar — IBS: R$ 8.850,00
  - Crédito: Fornecedores a Pagar: valor líquido
* E registrar no log: "NF-e #9012 — Fornecedor qualificado. Créditos liberados automaticamente" [INDEX].

#### Cenário 2: Fornecedor do Simples Nacional — crédito bloqueado

* Dado que uma NF-e de entrada foi emitida por um fornecedor optante pelo Simples Nacional;
* E o Simples Nacional não gera crédito integral de CBS/IBS para o adquirente no Lucro Real;
* Quando o SAP processar o recebimento fiscal;
* Então o sistema deve detectar que o regime do fornecedor é incompatível com a geração de crédito integral;
* E bloquear o lançamento em "Impostos a Recuperar";
* E direcionar o valor para a conta transitória "Créditos em Análise — Fornecedor Não Qualificado";
* E gerar uma notificação para o time de Suprimentos: "Fornecedor XYZ — Simples Nacional. Créditos de CBS/IBS bloqueados. Avalie a viabilidade de migração do fornecedor para o Lucro Real ou renegociação de preço." [INDEX].

#### Cenário 3: Fornecedor se regulariza em 30 dias — crédito liberado

* Dado que o fornecedor ABC teve créditos bloqueados há 25 dias;
* E o fornecedor concluiu a migração para o Lucro Real e atualizou seu cadastro no programa de compliance;
* Quando o job de reanálise periódica (30 dias) executar;
* Então o SAP deve detectar a regularização do fornecedor;
* E transferir os créditos da conta transitória para "Impostos a Recuperar";
* E registrar: "Fornecedor ABC regularizado. Créditos transferidos: CBS R$ X, IBS R$ Y" [INDEX].

------------------------------
## 📝 US-02: Due Diligence Fiscal Pré-Contratação de Novos Fornecedores

### 1. Descrição da História (Visão de Negócio)

Como Gerente de Suprimentos e Compras,
Quero que o SAP execute uma avaliação fiscal automática de potenciais novos fornecedores antes da homologação no cadastro,
Para que a empresa condicione novas relações comerciais à capacidade do fornecedor de gerar créditos tributários de CBS/IBS [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Score Fiscal do Fornecedor): O SAP deve calcular um score fiscal de 0 a 100 para cada potencial fornecedor com base em: regime tributário (Lucro Real = 100, Presumido = 50, Simples = 20), certidões fiscais válidas, histórico de compliance e capacidade de emitir NF-e com destaque de CBS/IBS.
* RN02 (Nota de Corte para Homologação): Fornecedores com score abaixo de 50 devem ser bloqueados para homologação automática e exigir aprovação do Gerente de Suprimentos com justificativa de negócio documentada.
* RN03 (Reavaliação Anual): Fornecedores ativos devem ser reavaliados anualmente, e a queda do score abaixo de 50 deve disparar um plano de ação corretiva ou descredenciamento.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Novo fornecedor com score alto — homologado automaticamente

* Dado que o time de Compras está cadastrando um novo fornecedor que opera no Lucro Real, com certidões fiscais válidas e NF-e com destaque de CBS/IBS;
* Quando o SAP calcular o score fiscal;
* Então o score deve ser ≥ 80;
* E o fornecedor deve ser homologado automaticamente;
* E o status "Qualificado para Créditos CBS/IBS" deve ser atribuído ao cadastro [INDEX].

#### Cenário 2: Fornecedor com score baixo — bloqueado e exige aprovação gerencial

* Dado que um potencial fornecedor opera no Simples Nacional e não emite NF-e com destaque de CBS/IBS;
* Quando o SAP calcular o score fiscal e retornar 20;
* Então o sistema deve bloquear a homologação automática;
* E exibir: "Fornecedor não atinge o score mínimo de compliance fiscal (20/100). Homologação requer aprovação do Gerente de Suprimentos.";
* E gerar uma tarefa no workflow do Gerente com os dados do fornecedor e a justificativa necessária [INDEX].
