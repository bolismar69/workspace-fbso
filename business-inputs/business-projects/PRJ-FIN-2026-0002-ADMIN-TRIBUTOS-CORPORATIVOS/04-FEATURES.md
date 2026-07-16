# Detalhamento de Features do Projeto
* **Projeto:** Portal Corporativo de Gestão Tributária — Autonomia do Time de Finanças na Administração de Impostos
* **Código:** PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS
* **Status:** Pronto para Detalhamento de User Stories
* **Responsáveis:** Product Owner (PO) — Portal de Gestão Tributária, Comitê Fiscal
* **Referência:** [03-EPICS.md](./03-EPICS.md)
* [INDEX] Deriva de: [03-EPICS.md](./03-EPICS.md) — Seção 3 (Épicos)

------------------------------
## 1. Objetivo do Documento

Este documento detalha as Features (capacidades funcionais de produto) derivadas dos Épicos definidos em [03-EPICS.md](./03-EPICS.md). Cada Feature está vinculada a um Épico e descreve o objetivo de negócio, o comportamento funcional esperado e as regras de negócio aplicáveis.

A numeração adota o padrão **Épico.Feature** (`01.1` = Épico 1, Feature 1), garantindo rastreabilidade completa da diretriz estratégica até a User Story.

**Features por Entrega:** 3 (Entrega 1) + 2 (Entrega 2) + 2 (Entrega 3) + 3 (Entrega 4) = **10 Features** no total.

------------------------------
## 2. Entrega 1 — Portal: Gestão Básica de Alíquotas

* **Fase:** Fundação do Portal — Cadastro e Consulta de Tabelas Fiscais
* **Épico:** 01 — Motor de Cadastro Fiscal (Alíquotas, Validações e Classificações)
* **Módulos do Portal:** M1 (Painel de Alíquotas Vigentes), M2 (Cadastro e Manutenção de Alíquotas), M3 (Gestão de Classificações e Regimes)
* **Requisitos Vinculados:** BR-01, BR-02, BR-03
* **Partes Interessadas Primárias:** Analistas Fiscais, Comitê Fiscal, Gerente Fiscal

### 🔢 ÉPICO 01: Motor de Cadastro Fiscal — Alíquotas, Validações e Classificações

#### Feature 01.1: Painel de Alíquotas Vigentes

* **Objetivo de Negócio:** Prover ao time de Finanças uma visão gerencial consolidada e imediata de todas as alíquotas ativas na companhia, eliminando a necessidade de consultar planilhas dispersas ou abrir chamados técnicos para responder perguntas básicas sobre o status do patrimônio fiscal [INDEX].
* **Descrição Funcional:** Tela inicial do portal apresentando uma visão tabular e filtrável de todas as alíquotas vigentes. O painel permite filtrar por tributo (CBS, IBS, IS, ICMS, ISS, PIS, COFINS, IPI), por unidade federativa, por município de destino (código IBGE), por período de vigência e por status (vigente, programada para início futuro, expirada). Inclui indicadores visuais de integridade que sinalizam alíquotas prestes a expirar nos próximos 30 dias, alíquotas com potenciais conflitos detectados e alíquotas do regime antigo sem substituição correspondente no novo regime durante o Período Híbrido.
* **Regras de Negócio e Restrições:**
  * O painel é acessível a todos os perfis de usuário (Administrador Fiscal, Analista Fiscal, Auditor/Controller), sem restrição de visualização.
  * O indicador de "alíquotas a expirar" deve considerar a data de término de vigência e alertar com 30 dias de antecedência.
  * O indicador de "conflito potencial" é calculado automaticamente: duas alíquotas para o mesmo tributo, mesma região geográfica e períodos de vigência com interseção são sinalizadas em amarelo (sobreposição parcial) ou vermelho (sobreposição total).
  * A ordenação padrão exibe primeiro as alíquotas com alertas ativos, depois por tributo e data de vigência.

#### Feature 01.2: Cadastro e Manutenção de Alíquotas

* **Objetivo de Negócio:** Permitir que o time de Finanças realize com autonomia e segurança a criação, edição e desativação de alíquotas de todos os tributos aplicáveis, com validações automáticas que impeçam configurações fiscais inválidas e substituam o conhecimento tácito dos analistas seniores por regras formais [INDEX].
* **Descrição Funcional:** Conjunto de formulários de negócio para cadastro de novas alíquotas e edição de alíquotas existentes. Cada formulário contém campos padronizados: tributo (seleção por lista), abrangência geográfica (nacional, estadual com UF, ou municipal com código IBGE), classificação fiscal associada (NCM, NBS ou CClassTrib), data de início e fim de vigência, valor da alíquota (percentual), regime tributário aplicável e campo de observações para justificativa de negócio. A desativação de uma alíquota não a remove fisicamente — apenas encerra sua vigência, preservando o histórico.
* **Regras de Negócio e Restrições:**
  * **RN-01 — Conflito de Vigência:** O sistema impede a criação ou edição de uma alíquota se já existir outra alíquota para o mesmo tributo, mesma região geográfica e com período de vigência que se sobreponha total ou parcialmente à nova alíquota. Mensagem: "Já existe uma alíquota de [tributo] vigente para [região] no período [data início] a [data fim]. Ajuste as datas ou revise a alíquota existente."
  * **RN-02 — Integridade de Classificação:** O sistema impede a associação de uma alíquota a um código de classificação fiscal (NCM, NBS, CClassTrib) que não esteja previamente cadastrado no Módulo 3. Mensagem: "O código [NCM/NBS/CClassTrib] [código] não está cadastrado. Cadastre a classificação fiscal antes de associar uma alíquota."
  * **RN-03 — Consistência de Transição (Período Híbrido):** Durante o Período Híbrido (2029–2032), ao desativar uma alíquota de um tributo do regime antigo (ICMS, ISS, PIS, COFINS, IPI), o sistema verifica se existe ao menos uma alíquota do tributo correspondente do novo regime (IBS para ICMS/ISS, CBS para PIS/COFINS/IPI) vigente para a mesma região. Se não existir, emite um alerta — mas não bloqueia a operação, pois pode haver cenários legítimos de extinção sem substituição.
  * **RN-04 — Datas de Vigência:** A data de fim de vigência, quando preenchida, deve ser estritamente posterior à data de início. O sistema impede datas retrógradas. Mensagem: "A data de término da vigência deve ser posterior à data de início."
  * **RN-05 — Justificativa Obrigatória:** Toda desativação de alíquota (encerramento de vigência) exige o preenchimento do campo de justificativa de negócio. Mensagem: "É obrigatório informar a justificativa para encerramento da vigência desta alíquota."

#### Feature 01.3: Gestão de Classificações e Regimes

* **Objetivo de Negócio:** Centralizar o cadastro e a manutenção das classificações fiscais de produtos e serviços e dos regimes tributários que servem como base para a aplicação das alíquotas, garantindo que as alíquotas estejam sempre associadas a classificações válidas e atualizadas [INDEX].
* **Descrição Funcional:** Módulo de cadastro centralizado das classificações fiscais (NCM — Nomenclatura Comum do Mercosul, NBS — Nomenclatura Brasileira de Serviços, CClassTrib — Código de Classificação Tributária, CFOP — Código Fiscal de Operações e Prestações) e dos regimes tributários (Lucro Real, Lucro Presumido, Simples Nacional). Permite a criação de novas classificações, edição de descrições e desativação de códigos obsoletos. Exibe, para cada classificação, a lista de alíquotas a ela associadas.
* **Regras de Negócio e Restrições:**
  * **RN-06 — Vinculação Protegida:** Uma classificação fiscal não pode ser desativada se houver alíquotas vigentes a ela associadas. O sistema lista as alíquotas vinculadas e exige a desativação ou reassociação delas antes de prosseguir. Mensagem: "Esta classificação possui [N] alíquotas vigentes associadas. Reassocie ou desative as alíquotas antes de desativar a classificação."
  * **RN-07 — Código Único:** Não é permitido cadastrar duas classificações com o mesmo código e tipo (ex: dois NCM com o mesmo código numérico). O sistema impede a duplicação. Mensagem: "O código [código] já está cadastrado como [tipo] — [descrição existente]."
  * **RN-08 — Formato de Código:** O sistema valida o formato do código conforme o tipo de classificação: NCM com 8 dígitos, CFOP com 4 dígitos. Códigos em formato inválido são rejeitados com mensagem explicativa.
  * **RN-09 — Regime Padrão:** O sistema exige que exatamente um regime tributário esteja marcado como "padrão" para novos cadastros de alíquotas que não especifiquem um regime. Esta marcação é gerenciada pelo Administrador Fiscal.

------------------------------
## 3. Entrega 2 — Governança e Auditoria Fiscal

* **Fase:** Controle de Acesso e Rastreabilidade — Quem pode fazer o quê e como cada ação é registrada
* **Épico:** 02 — Controle de Acesso e Rastreabilidade Fiscal
* **Módulos do Portal:** M6 (Administração de Acessos e Perfis), M4 (Linha do Tempo e Auditoria)
* **Requisitos Vinculados:** BR-04, BR-05
* **Partes Interessadas Primárias:** Controladoria e Compliance, Gerente Fiscal, Auditoria Interna

### 🔐 ÉPICO 02: Controle de Acesso e Rastreabilidade Fiscal

#### Feature 02.1: Administração de Acessos e Perfis

* **Objetivo de Negócio:** Garantir a segregação de funções no portal, assegurando que cada usuário tenha apenas os privilégios compatíveis com seu papel e que nenhum usuário acumule permissões conflitantes, em conformidade com os controles internos exigidos pela Lei das S.A. e pelo framework COSO [INDEX].
* **Descrição Funcional:** Módulo de gestão de usuários acessível exclusivamente ao perfil de Administrador Fiscal. Permite cadastrar novos usuários, associar cada usuário a um dos três perfis de acesso (Administrador Fiscal, Analista Fiscal, Auditor/Controller), desativar usuários e redefinir perfis. Cada perfil possui um conjunto fixo e não sobreponível de permissões.
* **Regras de Negócio e Restrições:**
  * **RN-10 — Perfis e Permissões:**
    | Ação | Administrador Fiscal | Analista Fiscal | Auditor/Controller |
    |:---|:---:|:---:|:---:|
    | Consultar alíquotas, classificações e regimes | ✅ | ✅ | ✅ |
    | Criar e editar alíquotas e classificações | ✅ | ✅ | ❌ |
    | Aprovar alterações de alto impacto (fluxo de duas etapas) | ✅ | ❌ | ❌ |
    | Gerenciar usuários e perfis de acesso | ✅ | ❌ | ❌ |
    | Visualizar trilha de auditoria | ✅ | ✅ | ✅ |
    | Exportar tabelas e relatórios | ✅ | ✅ | ✅ |
  * **RN-11 — Segregação de Funções:** Um Administrador Fiscal não pode aprovar uma alteração de alto impacto que ele próprio criou. O sistema deve identificar o criador da alteração e, se for o mesmo usuário que tenta aprovar, bloquear a ação. Mensagem: "Você não pode aprovar uma alteração que você mesmo propôs. A aprovação deve ser realizada por outro Administrador Fiscal."
  * **RN-12 — Auditor/Controller sem Escrita:** O perfil de Auditor/Controller não possui nenhuma capacidade de alteração no portal. Todos os botões e links de criação, edição ou desativação são ocultados ou desabilitados para este perfil.
  * **RN-13 — Rastreabilidade de Acesso:** Toda ação de gerenciamento de usuários (criação, alteração de perfil, desativação) é registrada na trilha de auditoria com identificação do Administrador Fiscal que a executou.

#### Feature 02.2: Trilha de Auditoria e Linha do Tempo

* **Objetivo de Negócio:** Prover rastreabilidade completa e imutável de todas as alterações realizadas nas tabelas fiscais do portal, permitindo que a Controladoria, a Auditoria Interna e órgãos externos reconstituam o histórico exato de qualquer alíquota, classificação ou regime — quem alterou, quando, o que foi alterado e por quê [INDEX].
* **Descrição Funcional:** Módulo de consulta que exibe a linha do tempo completa de alterações, com filtros por período, por usuário responsável, por tipo de entidade afetada (alíquota, classificação, regime) e por tipo de operação (criação, edição, desativação). Para cada evento, é possível expandir e visualizar a comparação lado a lado entre o estado anterior e o novo estado da entidade (diff visual). A trilha de auditoria é gerada automaticamente pelo portal — nenhuma ação do usuário é necessária para ativá-la, e nenhum usuário (incluindo Administrador Fiscal) pode desativá-la ou alterar registros já gerados.
* **Regras de Negócio e Restrições:**
  * **RN-14 — Automaticidade e Imutabilidade:** O registro de auditoria é gerado automaticamente no momento em que a alteração é confirmada no portal. Uma vez gerado, não pode ser modificado, excluído ou ocultado por nenhum perfil de usuário. Tentativas de adulteração devem ser detectadas e registradas como evento de segurança.
  * **RN-15 — Conteúdo do Registro:** Cada registro de auditoria contém obrigatoriamente: (a) identificação do usuário autenticado que realizou a ação, (b) data e hora exatas (timestamp), (c) tipo de entidade afetada e seu identificador único, (d) snapshot completo do estado anterior (JSON com todos os campos), (e) snapshot completo do novo estado, (f) tipo de operação (CRIAÇÃO, EDIÇÃO, DESATIVAÇÃO) e (g) justificativa de negócio fornecida pelo usuário.
  * **RN-16 — Retenção:** Os registros de auditoria são preservados pelo prazo legal aplicável a documentos fiscais (mínimo de 5 anos para tributos federais, podendo se estender a 10 anos conforme legislação específica). O portal deve garantir a integridade e acessibilidade dos registros durante todo o período de retenção.
  * **RN-17 — Acesso à Trilha:** A Linha do Tempo é acessível a todos os perfis de usuário, mas apenas para consulta. O perfil de Auditor/Controller tem a Linha do Tempo como sua tela principal de trabalho.

------------------------------
## 4. Entrega 3 — Portal: Operações em Escala

* **Fase:** Eficiência Operacional — Governança proporcional ao risco e processamento em volume
* **Épico:** 03 — Operações Fiscais em Escala
* **Módulos do Portal:** M5 (Importação e Exportação em Lote); Fluxos de Aprovação (funcionalidade transversal)
* **Requisitos Vinculados:** BR-06, BR-07
* **Partes Interessadas Primárias:** Gerente Fiscal e Controller, Comitê Fiscal, Analistas Fiscais

### 📊 ÉPICO 03: Operações Fiscais em Escala

#### Feature 03.1: Fluxos de Aprovação para Alterações de Alto Impacto

* **Objetivo de Negócio:** Adicionar uma camada de governança proporcional ao risco financeiro, garantindo que alterações de alíquotas com impacto material relevante sejam revisadas e aprovadas por um segundo par antes de entrar em vigor, sem burocratizar os ajustes rotineiros de baixo impacto [INDEX].
* **Descrição Funcional:** Mecanismo de fluxo de aprovação em duas etapas que é acionado automaticamente quando uma alteração de alíquota atinge ou ultrapassa o patamar de materialidade definido pelo Comitê Fiscal. Quando um Analista Fiscal propõe uma alteração de alto impacto, ela entra no status "Pendente de Aprovação" e não entra em vigor. Um Administrador Fiscal (diferente do proponente) ou Controller deve revisar e aprovar explicitamente a alteração para que ela seja efetivada. Alterações abaixo do patamar seguem o fluxo normal — o Analista Fiscal as realiza diretamente, sem aprovação adicional.
* **Regras de Negócio e Restrições:**
  * **RN-18 — Patamar de Materialidade:** O patamar é definido e revisado periodicamente pelo Comitê Fiscal, expresso em valor de faturamento mensal estimado impactado pela alteração (ex: ≥ R$ 100.000/mês). O portal permite que o Comitê Fiscal ajuste este valor através de parametrização acessível ao perfil de Administrador Fiscal.
  * **RN-19 — Bloqueio até Aprovação:** Enquanto uma alteração estiver no status "Pendente de Aprovação", a alíquota anterior permanece vigente e a nova alíquota não produz efeitos. O sistema exibe a alteração pendente no Painel de Alíquotas com um indicador visual específico (ex: ícone de relógio e cor laranja).
  * **RN-20 — Segregação Proponente × Aprovador:** O sistema impede que o mesmo usuário que propôs a alteração a aprove. Se houver apenas um Administrador Fiscal cadastrado, o sistema alerta sobre a necessidade de cadastrar um segundo administrador ou designar um Controller como aprovador.
  * **RN-21 — Registro na Trilha:** Tanto a proposta quanto a aprovação (ou rejeição) são registradas na trilha de auditoria como eventos distintos, permitindo rastrear o ciclo completo da decisão. Se rejeitada, a justificativa da rejeição é obrigatória.
  * **RN-22 — Prazo de Aprovação:** Se uma alteração de alto impacto permanecer "Pendente de Aprovação" por mais de 5 dias úteis, o sistema envia um alerta ao Gerente Fiscal e ao Comitê Fiscal. Este prazo é parametrizável.

#### Feature 03.2: Importação e Exportação de Alíquotas em Lote

* **Objetivo de Negócio:** Permitir que o time de Finanças responda com agilidade a publicações oficiais de grande volume — como a divulgação das alíquotas municipais de IBS pelo Comitê Gestor para todos os 5.570 municípios — processando milhares de alíquotas em horas em vez de semanas, e com validação automática que previne erros de transcrição [INDEX].
* **Descrição Funcional:** Funcionalidade de carga massiva de alíquotas a partir de planilhas padronizadas (formato Excel ou CSV). O usuário faz o upload do arquivo e o portal processa cada linha, aplicando as mesmas validações de negócio do cadastro manual (conflitos de vigência, integridade de classificações, consistência de transição). Ao final do processamento, o portal exibe um relatório detalhado: total de linhas processadas, quantas foram aceitas e efetivadas, quantas foram rejeitadas e o motivo específico de cada rejeição. O usuário pode então corrigir a planilha e reenviar apenas as linhas rejeitadas. A exportação permite extrair as tabelas vigentes em formato de planilha para relatórios de conformidade e auditorias externas.
* **Regras de Negócio e Restrições:**
  * **RN-23 — Template Obrigatório:** O portal disponibiliza um arquivo template padronizado com as colunas obrigatórias e a formatação esperada. O upload de arquivos que não seguem o template é rejeitado com mensagem indicando as divergências de estrutura.
  * **RN-24 — Processamento com Validação:** Cada linha do arquivo é validada individualmente contra as mesmas regras de negócio aplicadas ao cadastro manual (RN-01 a RN-05). Linhas que violem regras são rejeitadas individualmente, sem interromper o processamento das demais. O relatório final lista cada linha rejeitada com o número da linha no arquivo original, o conteúdo e o motivo da rejeição.
  * **RN-25 — Atomicidade por Linha:** Cada linha aceita gera um registro de auditoria individual. O processamento em lote não é uma transação atômica — linhas aceitas são efetivadas, linhas rejeitadas não.
  * **RN-26 — Limite de Volume:** O portal suporta o processamento de arquivos com até 10.000 linhas por upload. Acima deste limite, o sistema orienta o usuário a dividir o arquivo em lotes menores.
  * **RN-27 — Histórico de Cargas:** O portal mantém um histórico de todas as operações de importação em lote, com data, usuário, arquivo original e relatório de processamento, acessível na Linha do Tempo e filtrável por tipo de operação "Importação em Lote".

------------------------------
## 5. Entrega 4 — Portal Completo: Expansão Funcional

* **Fase:** Inteligência Fiscal e Analytics — Informação acionável para Alta Gestão e Preparação para o Longo Prazo
* **Épico:** 04 — Inteligência Fiscal e Analytics
* **Módulos do Portal:** Relatórios, Dashboards e Suporte ao Período Híbrido (funcionalidades transversais)
* **Requisitos Vinculados:** BR-08, BR-09, BR-10
* **Partes Interessadas Primárias:** CFO e Comitê Fiscal, Controladoria, Gerente Fiscal e Analistas

### 📈 ÉPICO 04: Inteligência Fiscal e Analytics

#### Feature 04.1: Relatórios Gerenciais de Governança

* **Objetivo de Negócio:** Substituir a compilação manual de dados de governança — que hoje consome horas do time de Controladoria — por relatórios gerenciais mensais gerados automaticamente, sumarizando toda a atividade de gestão tributária do período para apresentação ao Comitê Fiscal e à Controladoria [INDEX].
* **Descrição Funcional:** Funcionalidade de geração de relatórios gerenciais mensais que sumarizam todas as alterações realizadas nas tabelas fiscais do portal. O relatório é gerado automaticamente no primeiro dia útil de cada mês, cobrindo o mês anterior, e fica disponível para download e consulta online. Ele agrupa as alterações por tipo de tributo, por usuário responsável e por justificativa de negócio, e inclui seções específicas para: (a) alterações de alto impacto aprovadas no período, (b) alterações de alto impacto rejeitadas, (c) conflitos de vigência detectados e (d) resumo estatístico (total de criações, edições e desativações).
* **Regras de Negócio e Restrições:**
  * **RN-28 — Geração Automática:** O relatório é gerado automaticamente no 1º dia útil de cada mês, cobrindo o período do mês anterior. O Administrador Fiscal pode solicitar a geração de relatórios ad hoc para períodos específicos.
  * **RN-29 — Disponibilidade:** O relatório fica disponível no portal por 24 meses. Relatórios mais antigos são arquivados e podem ser solicitados ao Administrador Fiscal.
  * **RN-30 — Envio Automático:** O portal envia o relatório por e-mail para a lista de distribuição definida pelo Comitê Fiscal (Controladoria, Gerente Fiscal, Controller) no momento da geração.
  * **RN-31 — Formato:** O relatório está disponível em formato PDF (para apresentação e arquivamento) e Excel (para análises adicionais pela Controladoria).

#### Feature 04.2: Dashboards Gerenciais de KPIs Fiscais

* **Objetivo de Negócio:** Eliminar a dependência do CFO e do Comitê Fiscal de relatórios ad hoc do time técnico para responder perguntas básicas sobre o status do patrimônio fiscal da companhia, provendo painéis visuais com indicadores-chave atualizados e acessíveis a qualquer momento [INDEX].
* **Descrição Funcional:** Conjunto de painéis visuais (dashboards) integrados ao portal, acessíveis aos perfis de Administrador Fiscal e Auditor/Controller, que exibem em tempo real os indicadores-chave da gestão tributária definidos na MATRIZ-KPI.md. Cada dashboard é específico para um conjunto de indicadores:
  * **Dashboard de Cobertura Fiscal:** Mapa do Brasil com indicação visual (escala de cores) dos municípios com alíquotas de IBS cadastradas vs. municípios onde a companhia opera sem alíquota cadastrada. Indicador numérico de cobertura geográfica (KPI E1).
  * **Dashboard de Atividade:** Gráfico de barras com volume de alterações por mês, segmentado por tipo de operação (criação, edição, desativação) e por tributo. Indicador de tendência de adoção do portal (KPI E3).
  * **Dashboard de Governança:** Indicadores de cobertura de trilha de auditoria (KPI G1), conflitos prevenidos no período (KPI G2) e status das aprovações de alto impacto pendentes.
  * **Dashboard de Vencimentos:** Linha do tempo das próximas 12 semanas com as alíquotas que expirarão no período, agrupadas por semana e por tributo.
* **Regras de Negócio e Restrições:**
  * **RN-32 — Atualização:** Os dados dos dashboards são atualizados em tempo real, refletindo as alterações realizadas no portal imediatamente.
  * **RN-33 — Acesso:** Os dashboards são acessíveis aos perfis de Administrador Fiscal e Auditor/Controller. O perfil de Analista Fiscal tem acesso apenas ao Dashboard de Vencimentos (relevante para sua operação diária).
  * **RN-34 — Drill-Down:** Ao clicar em um elemento do dashboard (ex: uma barra do gráfico de alterações, um município no mapa), o usuário é direcionado para a visão detalhada correspondente (ex: lista de alterações do período, alíquotas do município selecionado).
  * **RN-35 — Exportação:** Os dashboards podem ser exportados para PDF (formato de apresentação executiva) para inclusão em reportes ao CFO e ao Conselho.

#### Feature 04.3: Suporte ao Período Híbrido — Dupla Gestão de Regimes

* **Objetivo de Negócio:** Garantir que o investimento no portal permaneça válido durante os 4 anos mais complexos da transição tributária nacional (2029–2032), permitindo ao time de Finanças gerenciar com clareza e segurança as tabelas dos regimes antigo e novo simultaneamente, com indicadores visuais de transição e funcionalidades de desativação progressiva [INDEX].
* **Descrição Funcional:** Conjunto de funcionalidades que preparam o portal para operar durante o Período Híbrido:
  * **Visão Dual de Regimes:** Cada alíquota exibe um indicador visual de regime: "LEGADO" (para ICMS, ISS, PIS, COFINS, IPI), "IVA DUAL" (para CBS, IBS, IS) ou "TRANSITÓRIO" (para alíquotas criadas especificamente para o período de convivência). O Painel de Alíquotas Vigentes (Feature 01.1) permite filtrar por regime.
  * **Mapeamento de Correlação:** O portal permite definir e visualizar relações de substituição entre tributos: "ICMS interestadual 12% para SP → MG será substituído por IBS estadual [X]% a partir de [data]". Estas relações são usadas pela validação de transição (RN-03) para alertar sobre substituições ausentes.
  * **Desativação Progressiva:** Funcionalidade que permite desativar em lote todas as alíquotas de um tributo do regime legado a partir de uma data específica, seguindo o cronograma constitucional de transição. A desativação em lote exige aprovação de um Administrador Fiscal (fluxo de duas etapas, Feature 03.1) e gera registros de auditoria individuais para cada alíquota desativada.
  * **Indicador de Prontidão:** O Dashboard de Cobertura Fiscal (Feature 04.2) inclui uma visão específica do Período Híbrido: para cada tributo em extinção, mostra o percentual de alíquotas que já possuem substituição mapeada no novo regime.
* **Regras de Negócio e Restrições:**
  * **RN-36 — Cronograma Constitucional:** O portal deve permitir o cadastro e a manutenção do cronograma de transição (datas de redução progressiva de cada tributo antigo e aumento correspondente do novo), que serve como base para os alertas de transição e para a funcionalidade de desativação progressiva.
  * **RN-37 — Alertas de Transição:** Com 90 dias de antecedência de cada marco do cronograma de transição, o portal gera um alerta no Painel de Alíquotas e envia notificação ao Gerente Fiscal e ao Comitê Fiscal, listando as alíquotas do regime legado que precisarão de substituição.
  * **RN-38 — Dupla Exibição no Checkout (Informativo):** Embora o portal não realize cálculos de impostos, ele provê uma visão informativa de como uma mesma operação seria tributada nos dois regimes, permitindo ao time de Finanças e ao Comitê Fiscal avaliar impactos comparativos durante o planejamento da transição.

------------------------------
## 6. Matriz de Rastreabilidade das Features

| Feature | Épico | Entrega | BRs Vinculados | Módulo | User Stories |
|:---|:---|:---|:---|:---|:---|
| 01.1 — Painel de Alíquotas Vigentes | 01 — Motor de Cadastro Fiscal | Entrega 1 | BR-01 | M1 | 🔜 Pendente |
| 01.2 — Cadastro e Manutenção de Alíquotas | 01 — Motor de Cadastro Fiscal | Entrega 1 | BR-01, BR-02 | M2 | 🔜 Pendente |
| 01.3 — Gestão de Classificações e Regimes | 01 — Motor de Cadastro Fiscal | Entrega 1 | BR-03 | M3 | 🔜 Pendente |
| 02.1 — Administração de Acessos e Perfis | 02 — Controle de Acesso e Rastreabilidade | Entrega 2 | BR-04 | M6 | 🔜 Pendente |
| 02.2 — Trilha de Auditoria e Linha do Tempo | 02 — Controle de Acesso e Rastreabilidade | Entrega 2 | BR-05 | M4 | 🔜 Pendente |
| 03.1 — Fluxos de Aprovação para Alto Impacto | 03 — Operações Fiscais em Escala | Entrega 3 | BR-06 | Transversal | 🔜 Pendente |
| 03.2 — Importação e Exportação em Lote | 03 — Operações Fiscais em Escala | Entrega 3 | BR-07 | M5 | 🔜 Pendente |
| 04.1 — Relatórios Gerenciais de Governança | 04 — Inteligência Fiscal e Analytics | Entrega 4 | BR-08 | Transversal | 🔜 Pendente |
| 04.2 — Dashboards Gerenciais de KPIs Fiscais | 04 — Inteligência Fiscal e Analytics | Entrega 4 | BR-09 | Transversal | 🔜 Pendente |
| 04.3 — Suporte ao Período Híbrido | 04 — Inteligência Fiscal e Analytics | Entrega 4 | BR-10 | Transversal | 🔜 Pendente |

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 013-agile-feature, agile-ba-practices.*
