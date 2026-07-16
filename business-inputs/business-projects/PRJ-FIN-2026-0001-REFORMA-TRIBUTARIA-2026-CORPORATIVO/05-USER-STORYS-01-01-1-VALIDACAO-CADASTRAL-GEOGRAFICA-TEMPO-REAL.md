# Histórias de Usuário (User Stories) — Feature 01.1
- Programa: Adequação Corporativa à Reforma Tributária Nacional
- Módulo: CRM / Onboarding de Clientes (Onda 1)
- Feature Relacionada: 01.01.1 — Validação Cadastral Geográfica em Tempo Real [INDEX]
- Status: Pronto para Desenvolvimento Técnico

------------------------------
## 📝 US-01: Autopreenchimento Geográfico via CEP no Cadastro de Clientes

### 1. Descrição da História (Visão de Negócio)

Como Analista de Cadastro e Onboarding de Clientes,
Quero que o sistema preencha automaticamente Estado, Município e Código IBGE a partir do CEP informado no cadastro,
Para eliminar erros manuais de digitação e garantir que o local de consumo esteja correto para o cálculo do IBS baseado no destino [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Disparador da Validação): A consulta geográfica deve ser acionada automaticamente no momento em que o usuário preencher o campo CEP e sair do campo (evento onBlur).
* RN02 (Campos Bloqueados Após Autopreenchimento): Os campos de Estado, Município e Código IBGE preenchidos automaticamente pelo sistema devem ser bloqueados para edição manual, exibindo o ícone de "validado automaticamente" ao lado.
* RN03 (CEP Inválido ou Não Encontrado): Se a base de CEPs retornar erro ou não localizar o CEP informado, o sistema deve exibir uma mensagem de alerta e liberar os campos para preenchimento manual guiado (veja US-02).

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Cadastro de cliente com CEP válido e encontrado na base

* Dado que o analista de cadastro está na tela de onboarding de um novo cliente PJ;
* Quando o analista digitar o CEP "06543-001" (Santana de Parnaíba - SP) no campo de endereço e sair do campo;
* Então o sistema deve consultar a base integrada de CEPs e preencher automaticamente os campos: Estado = "SP", Município = "Santana de Parnaíba", Código IBGE = "3547304";
* E exibir o ícone "✓ Validado" ao lado de cada campo autopreenchido;
* E bloquear a digitação manual nesses campos [INDEX].

#### Cenário 2: Cadastro com CEP não localizado na base

* Dado que o analista está cadastrando um cliente em uma região de expansão urbana recente;
* Quando o analista digitar um CEP que não consta na base de consulta e sair do campo;
* Então o sistema deve exibir a mensagem: "⚠️ CEP não localizado. Preencha manualmente os dados geográficos do destino.";
* E liberar os campos Estado, Município e Código IBGE para preenchimento manual guiado por listas fechadas [INDEX].

------------------------------
## 📝 US-02: Seleção Manual Guiada por Lista Fechada de Municípios (Zonas Rurais e Exceções)

### 1. Descrição da História (Visão de Negócio)

Como Analista de Cadastro atuando em regiões de difícil localização (zonas rurais, distritos industriais remotos),
Quero selecionar o município e código IBGE a partir de listas oficiais fechadas do IBGE,
Para garantir que mesmo cadastros sem CEP válido possuam o local de destino 100% qualificado para o IVA Dual [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Lista Fechada do IBGE): Os campos de Estado e Município devem utilizar componentes de autocomplete (typeahead) alimentados exclusivamente pela tabela oficial de municípios do IBGE, sem permitir digitação de texto livre.
* RN02 (Hierarquia UF → Município): A seleção do Estado deve filtrar dinamicamente a lista de municípios disponíveis, impedindo combinações inválidas (ex: município de MG com UF = SP).
* RN03 (Preenchimento Obrigatório do IBGE): O campo Código IBGE deve ser preenchido automaticamente ao selecionar o município na lista guiada, sem necessidade de digitação adicional.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Cadastro manual de endereço em zona rural sem CEP

* Dado que o analista está cadastrando um cliente cuja operação fica em zona rural sem CEP específico;
* E o sistema liberou os campos para preenchimento manual (conforme US-01);
* Quando o analista selecionar o Estado "Mato Grosso" no componente de autocomplete;
* Então o campo de Município deve exibir apenas os municípios do Mato Grosso na lista de sugestões;
* E ao selecionar "Cuiabá", o campo Código IBGE deve ser preenchido automaticamente com "5103403";
* E o sistema deve marcar o cadastro como "Dados Geográficos Validados" [INDEX].

#### Cenário 2: Tentativa de combinação inválida UF/Município

* Dado que o analista selecionou o Estado "São Paulo" no campo de UF;
* Quando o analista tentar digitar manualmente "Belo Horizonte" no campo de Município;
* Então o sistema não deve retornar resultados na lista de sugestões;
* E deve exibir a mensagem: "Nenhum município encontrado para o Estado selecionado" [INDEX].

------------------------------
## 📝 US-03: Indicador de Prontidão Fiscal e Flag "Apto para Venda"

### 1. Descrição da História (Visão de Negócio)

Como Gerente de Onboarding e Compliance Fiscal,
Quero que o sistema atribua automaticamente um selo de "Apto para Venda" quando o cadastro do cliente atingir 100% de conformidade geográfica,
Para que a força de vendas possa identificar visualmente quais clientes estão liberados para geração de propostas sem risco de erro fiscal [INDEX].

### 2. Regras de Negócio (Business Rules)

* RN01 (Condição para o Selo "Apto para Venda"): O status "Apto para Venda" só pode ser atribuído quando todos os campos obrigatórios estiverem preenchidos e validados: CEP, Estado, Município e Código IBGE.
* RN02 (Revogação Automática do Selo): Qualquer alteração nos campos geográficos do cadastro deve revogar imediatamente o selo "Apto para Venda" e reiniciar o processo de validação.
* RN03 (Auditoria de Mudanças): Toda alteração de dados geográficos deve ser registrada em log de auditoria com: ID do usuário, timestamp, campos alterados, valores antigos e novos.

### 3. Critérios de Aceite no Padrão BDD

#### Cenário 1: Cadastro atinge conformidade e recebe selo automaticamente

* Dado que um cadastro de cliente estava com o campo Código IBGE pendente;
* Quando o analista concluir o preenchimento do CEP e o sistema validar todos os campos geográficos automaticamente;
* Então o sistema deve alterar o status do cliente para "Apto para Venda — IVA Dual";
* E exibir um selo verde visível na ficha do cliente no CRM;
* E registrar no log de auditoria: "Cliente [ID] qualificado geograficamente para o IVA Dual por [Analista] em [Data/Hora]" [INDEX].

#### Cenário 2: Alteração cadastral revoga selo de conformidade

* Dado que um cliente possui o selo "Apto para Venda";
* Quando um analista alterar o endereço de entrega principal do cliente;
* Então o sistema deve remover imediatamente o selo "Apto para Venda";
* E exibir o status "⚠️ Dados Geográficos Pendentes de Revalidação";
* E registrar a revogação no log de auditoria, preservando os valores anteriores e os novos [INDEX].

#### Cenário 3: Consulta em lote de clientes não aptos

* Dado que o Gerente de Compliance Fiscal acessa o painel de governança cadastral;
* Quando o gerente aplicar o filtro "Clientes sem selo Apto para Venda";
* Então o sistema deve listar todos os clientes ativos que não possuem os 4 campos geográficos validados;
* E permitir exportação da lista para priorização do esforço de higienização cadastral [INDEX].

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 014-agile-user-story, agile-ba-practices.*
