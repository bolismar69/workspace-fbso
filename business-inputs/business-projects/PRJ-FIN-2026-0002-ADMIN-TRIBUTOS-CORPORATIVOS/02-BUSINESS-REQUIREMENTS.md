# Documento de Definição de Requisitos (02-BUSINESS-REQUIREMENTS.md)
* Projeto: Portal Corporativo de Gestão Tributária — Autonomia do Time de Finanças na Administração de Impostos
* Status: Pronto para Desdobramento em Features
* Responsáveis: Product Owner (PO) Corporativo, Comitê Fiscal
* Público-Alvo: Comitê Executivo, Controladoria, Time de Finanças e PMO
* [INDEX] Deriva de: [01-PROJECT-CHARTER.md](./01-PROJECT-CHARTER.md) — Seção 3 (Escopo de Negócio)
* [INDEX] Desdobra em: [03-EPICS.md](./03-EPICS.md)

------------------------------

## 1. Objetivos do Documento

Este documento especifica os Requisitos de Negócio (Business Requirements) necessários para cumprir o objetivo estratégico do Project Charter: prover ao time de Finanças um Portal Corporativo de Gestão Tributária que concentre, valide e rastreie todas as tabelas de impostos da organização, eliminando a dependência técnica nas operações de rotina e estabelecendo governança plena sobre o patrimônio fiscal da companhia.

------------------------------

## 2. Princípios Macroscópicos de Negócio (Guardrails)

* Fonte Única da Verdade Fiscal: Toda e qualquer alíquota, regime, classificação ou benefício fiscal praticado pela companhia deve estar registrado e validado no Portal de Gestão Tributária — nenhuma tabela fiscal pode existir exclusivamente em planilhas, e-mails ou sistemas isolados.
* Autonomia com Governança: O time de Finanças tem autonomia para realizar ajustes fiscais de rotina diretamente no portal, mas dentro de perímetros de acesso e fluxos de aprovação claramente definidos — a autonomia não elimina a segregação de funções nem a rastreabilidade.
* Rastreabilidade como Fundação: Nenhuma alteração em tabela fiscal pode ser efetuada sem que o registro completo de auditoria (quem, quando, valor anterior, valor novo, justificativa) seja gerado automaticamente e preservado de forma imutável.
* Validação Preventiva: O portal deve incorporar as regras de negócio do Comitê Fiscal como validações automáticas, impedindo a criação de configurações fiscais conflitantes ou inconsistentes — a prevenção de erros é preferível à detecção posterior.
* Prontidão para o Período Híbrido: Todas as estruturas de dados, interfaces e validações devem ser projetadas desde o início para suportar a coexistência dos regimes antigo e novo durante o Período Híbrido (2029–2032).

------------------------------

## 3. Requisitos de Negócio (Business Requirements - BR)

As necessidades estão organizadas em 4 blocos, cada um correspondendo a uma entrega do projeto e a um épico. A numeração dos BRs segue a sequência dos blocos (BR-01 a BR-03 = Bloco 1, BR-04 a BR-05 = Bloco 2, etc.).

## 📦 Bloco 1: Motor de Cadastro Fiscal (Alíquotas, Validações e Classificações)

> **Épico:** [01 — Motor de Cadastro Fiscal](./03-EPICS.md#-épico-01-motor-de-cadastro-fiscal--alíquotas-validações-e-classificações)
> **Entrega:** Entrega 1 — Portal: Gestão Básica de Alíquotas

* BR-01: Cadastro Centralizado de Alíquotas
   * Descrição: O portal deve permitir que analistas fiscais cadastrem, editem e desativem alíquotas de todos os tributos aplicáveis à companhia (CBS, IBS, IS, ICMS, ISS, PIS, COFINS, IPI), com campos padronizados de vigência (data de início e fim), abrangência geográfica (nacional, estadual ou municipal com código IBGE) e classificação fiscal associada (NCM, NBS, CClassTrib).
      * Justificativa de Negócio: Eliminar a pulverização de tabelas fiscais em planilhas e e-mails, estabelecendo uma fonte única e validada para todas as alíquotas praticadas pela companhia.

* BR-02: Validação Automática de Conflitos
   * Descrição: O portal deve impedir automaticamente a criação ou edição de alíquotas que resultem em conflitos de vigência (duas alíquotas diferentes para o mesmo tributo, mesma região e mesmo período), conflitos de classificação (alíquota associada a código fiscal inexistente) ou inconsistências de transição (alíquota do regime antigo desativada sem substituição correspondente no novo regime durante o Período Híbrido).
      * Justificativa de Negócio: Transferir as validações que hoje dependem do conhecimento tácito dos analistas fiscais seniores para regras automáticas, reduzindo a probabilidade de erros de configuração com impacto em vendas e faturamento.

* BR-03: Gestão de Classificações Fiscais
   * Descrição: O portal deve centralizar o cadastro e a manutenção das classificações fiscais de produtos e serviços (NCM, NBS, CClassTrib, CFOP) que servem como base para a aplicação das alíquotas, com rastreabilidade de alterações e vinculação às alíquotas aplicáveis.
      * Justificativa de Negócio: Garantir que as alíquotas estejam sempre associadas a classificações fiscais válidas e atualizadas, prevenindo erros de enquadramento tributário.

## 📦 Bloco 2: Controle de Acesso e Rastreabilidade Fiscal

> **Épico:** [02 — Controle de Acesso e Rastreabilidade Fiscal](./03-EPICS.md#-épico-02-controle-de-acesso-e-rastreabilidade-fiscal)
> **Entrega:** Entrega 2 — Governança e Auditoria Fiscal

* BR-04: Administração de Acessos e Perfis
   * Descrição: O portal deve permitir a gestão de usuários com perfis de acesso segregados: Administrador Fiscal (cadastro e aprovação), Analista Fiscal (cadastro), Auditor/Controller (consulta e visualização de trilha de auditoria), sem capacidade de alteração. Toda ação no portal deve ser atribuída ao usuário autenticado que a realizou.
      * Justificativa de Negócio: Atender aos requisitos de segregação de funções exigidos por controles internos (COSO) e Lei das S.A., garantindo que nenhum usuário tenha privilégios incompatíveis.

* BR-05: Trilha de Auditoria Completa
   * Descrição: Toda e qualquer alteração em tabelas fiscais realizada no portal deve gerar automaticamente um registro de auditoria contendo: (a) identificação do usuário responsável, (b) data e hora exatas da alteração, (c) entidade e identificador afetados, (d) valor anterior completo e novo valor completo, (e) justificativa de negócio fornecida pelo usuário. Este registro deve ser imutável e preservado pelo prazo legal exigido para documentos fiscais.
      * Justificativa de Negócio: Prover defesa documental em caso de fiscalizações, atender aos requisitos de controles internos (Lei das S.A., SOX) e permitir auditoria interna e externa sobre a gestão tributária.

## 📦 Bloco 3: Operações Fiscais em Escala

> **Épico:** [03 — Operações Fiscais em Escala](./03-EPICS.md#-épico-03-operações-fiscais-em-escala)
> **Entrega:** Entrega 3 — Portal: Operações em Escala

* BR-06: Fluxo de Aprovação para Alterações de Alto Impacto
   * Descrição: O portal deve suportar fluxos de aprovação em duas etapas para alterações que impactem tributos acima de um patamar de materialidade definido pelo Comitê Fiscal (ex: alteração de alíquota que afete faturamento estimado superior a R$ 100 mil mensais). Nestes casos, a alteração proposta por um analista fiscal deve ser revisada e aprovada por um Administrador Fiscal ou Controller antes de entrar em vigor.
      * Justificativa de Negócio: Adicionar uma camada de governança proporcional ao risco financeiro, sem comprometer a agilidade para ajustes de rotina de baixo impacto.

* BR-07: Importação e Exportação de Alíquotas em Lote
   * Descrição: O portal deve permitir a carga massiva de alíquotas a partir de planilhas padronizadas e a exportação das tabelas vigentes, viabilizando a atualização rápida diante de publicações oficiais (ex: divulgação das alíquotas municipais de IBS pelo Comitê Gestor para 5.570 municípios).
      * Justificativa de Negócio: Eliminar a digitação manual de grandes volumes de alíquotas, reduzindo o tempo de reação a publicações oficiais de semanas para horas e prevenindo erros de transcrição.

## 📦 Bloco 4: Inteligência Fiscal e Analytics

> **Épico:** [04 — Inteligência Fiscal e Analytics](./03-EPICS.md#-épico-04-inteligência-fiscal-e-analytics)
> **Entrega:** Entrega 4 — Portal Completo: Expansão Funcional

* BR-08: Relatórios Gerenciais de Governança
   * Descrição: O portal deve gerar relatórios mensais sumarizando todas as alterações realizadas nas tabelas fiscais, agrupadas por tipo de tributo, por usuário responsável e por justificativa de negócio, para apresentação ao Comitê Fiscal e à Controladoria.
      * Justificativa de Negócio: Prover visibilidade gerencial sobre a atividade de gestão tributária, permitindo que o Comitê Fiscal e a Controladoria monitorem padrões, identifiquem anomalias e demonstrem conformidade em auditorias.

* BR-09: Dashboards Gerenciais de KPIs Fiscais
   * Descrição: O portal deve prover painéis visuais com indicadores-chave da gestão tributária, incluindo: total de alíquotas vigentes por tributo, distribuição geográfica das alíquotas de IBS, alterações realizadas no período, alíquotas com vigência a expirar nos próximos 30 dias e status de completude da base (municípios cobertos vs. total de municípios com operação).
      * Justificativa de Negócio: Permitir que o CFO, o Comitê Fiscal e a Controladoria tenham visibilidade imediata sobre o status do patrimônio fiscal da companhia, sem depender de relatórios manuais do time técnico.

* BR-10: Suporte ao Período Híbrido — Dupla Gestão de Regimes
   * Descrição: O portal deve permitir a gestão simultânea de tabelas dos regimes antigo (ICMS, ISS, PIS, COFINS, IPI) e novo (CBS, IBS, IS), com indicadores visuais claros de qual regime cada alíquota pertence, mapeamento de correlação entre tributos em extinção e seus substitutos, e funcionalidade de desativação progressiva conforme o cronograma constitucional de transição (2029–2032).
      * Justificativa de Negócio: Preparar o time de Finanças para operar com clareza e segurança durante os 4 anos de convivência dos dois modelos tributários, evitando erros de apuração e reduzindo o risco de autuações.

------------------------------

## 4. Matriz de Rastreabilidade

Este mapeamento garante que nenhuma necessidade de negócio fique descoberta na fase de execução e que cada requisito esteja vinculado ao bloco, épico e entrega corretos:

| ID Requisito | Bloco / Épico | Meta Relacionada no Charter | Área de Negócio Madrinha | Entrega |
|---|---|---|---|---|
| BR-01 | Bloco 1 / Épico 01 — Motor de Cadastro Fiscal | Fonte Única da Verdade Fiscal | Controladoria / Comitê Fiscal | Entrega 1 |
| BR-02 | Bloco 1 / Épico 01 — Motor de Cadastro Fiscal | Blindagem contra Erros de Configuração | Comitê Fiscal | Entrega 1 |
| BR-03 | Bloco 1 / Épico 01 — Motor de Cadastro Fiscal | Fonte Única da Verdade Fiscal | Controladoria | Entrega 1 |
| BR-04 | Bloco 2 / Épico 02 — Controle de Acesso e Rastreabilidade | Autonomia com Governança | Controladoria / Compliance | Entrega 2 |
| BR-05 | Bloco 2 / Épico 02 — Controle de Acesso e Rastreabilidade | Governança e Rastreabilidade Societária | Controladoria / Auditoria Interna | Entrega 2 |
| BR-06 | Bloco 3 / Épico 03 — Operações Fiscais em Escala | Governança e Rastreabilidade Societária | Comitê Fiscal / Controladoria | Entrega 3 |
| BR-07 | Bloco 3 / Épico 03 — Operações Fiscais em Escala | Autonomia Operacional do Time de Finanças | Comitê Fiscal | Entrega 3 |
| BR-08 | Bloco 4 / Épico 04 — Inteligência Fiscal e Analytics | Governança e Rastreabilidade Societária | Controladoria / CFO | Entrega 4 |
| BR-09 | Bloco 4 / Épico 04 — Inteligência Fiscal e Analytics | Preparação para Expansão Funcional | CFO / Comitê Fiscal | Entrega 4 |
| BR-10 | Bloco 4 / Épico 04 — Inteligência Fiscal e Analytics | Prontidão para o Período Híbrido | Comitê Fiscal / Controladoria | Entrega 4 |

------------------------------

## 5. Critérios de Homologação de Negócio (UAT)

A aceitação final dos requisitos pelos Business Owners ocorrerá quando:

1. **Motor de Cadastro Fiscal (Bloco 1 / Entrega 1):** Um analista fiscal conseguir cadastrar uma nova alíquota de IBS para um município, ser impedido pelo portal de criar uma alíquota conflitante para o mesmo município e período, e visualizar a alíquota no painel de alíquotas vigentes — tudo sem intervenção do time técnico. Simultaneamente, um código NCM inexistente não pode ser associado a uma alíquota; o portal deve rejeitar a operação com mensagem clara.
2. **Controle de Acesso e Rastreabilidade (Bloco 2 / Entrega 2):** Um usuário com perfil de Auditor/Controller consegue visualizar todas as alíquotas e a trilha de auditoria, mas não consegue editar nenhum dado. O Controller consegue rastrear quem alterou uma alíquota específica, quando, qual era o valor anterior e qual a justificativa registrada.
3. **Operações em Escala (Bloco 3 / Entrega 3):** Uma alteração de alíquota com impacto estimado acima do patamar de materialidade é bloqueada até que um Administrador Fiscal a aprove. O time fiscal consegue carregar uma planilha com alíquotas de IBS para 100 municípios, e o portal processa a carga reportando quais registros foram aceitos e quais foram rejeitados com o motivo específico.
4. **Inteligência Fiscal (Bloco 4 / Entrega 4):** O relatório mensal de governança é gerado automaticamente e reflete exatamente as alterações do período. O dashboard de KPIs permite ao CFO visualizar a cobertura geográfica de alíquotas sem acionar o time técnico. O portal exibe com clareza quais alíquotas pertencem ao regime antigo e quais ao novo, com indicadores de transição entre eles.

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: brd-creation, business-rules-analysis.*
