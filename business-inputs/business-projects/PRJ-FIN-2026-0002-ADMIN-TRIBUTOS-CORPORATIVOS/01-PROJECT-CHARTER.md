# Termo de Abertura do Projeto (Project Charter - 01-PROJECT-CHARTER.md)
* Projeto Tático: Portal Corporativo de Gestão Tributária — Autonomia do Time de Finanças na Administração de Impostos
* Programa Pai: PRJ-FIN-2026-0001 — Adequação Corporativa à Reforma Tributária Nacional
* Foco: Autonomia operacional, governança fiscal e rastreabilidade societária
* Versão: 1.0 (Visão de Negócios Pura / Alta Gestão)
* Base de Conhecimento: Documentos de suporte técnico-tributário do programa PRJ-FIN-2026-0001, capacitação RTC CFC-RFB, Glossário do Projeto

------------------------------

## 1. Justificativa Estratégica do Projeto

* A Reforma Tributária em curso (EC 132/2023) representa uma transformação sem precedentes na gestão de impostos corporativos. A transição do modelo atual — com 5 tributos sobre o consumo (ICMS, ISS, PIS, COFINS, IPI) — para o IVA Dual (CBS e IBS), somada à introdução do Imposto Seletivo (IS), multiplicará exponencialmente a complexidade da administração tributária da companhia.
* Neste novo cenário, o volume de alíquotas que precisam ser geridas ativamente pelo time de Finanças passará de dezenas para **centenas de combinações**: cada município brasileiro (5.570) poderá definir sua própria alíquota de IBS; cada classificação fiscal (NCM/NBS) terá regras específicas de CBS, IS e regimes especiais; e durante o Período Híbrido (2029–2032), as tabelas dos regimes antigo e novo coexistirão simultaneamente.
* Hoje, a administração dessas tabelas fiscais é realizada de forma **artesanal e dependente da área técnica**: analistas fiscais precisam abrir chamados para cada ajuste de alíquota, aguardar priorização do time de engenharia e confiar que a alteração foi aplicada corretamente — sem visibilidade direta sobre o resultado. Este modelo operacional é insustentável diante do volume e da criticidade das mudanças que se aproximam.
* A dependência do time técnico para ajustes fiscais de rotina gera três riscos de negócio inaceitáveis:
  1. **Morosidade na resposta a mudanças regulatórias:** Entre a publicação de uma nova alíquota por um ente federativo e sua aplicação efetiva nos sistemas de venda e faturamento, transcorrem dias ou semanas — tempo durante o qual a companhia pode estar praticando preços incorretos, com risco de perda de margem ou exposição a autuações.
  2. **Fragilidade na governança fiscal:** Sem uma ferramenta especializada, as alterações em alíquotas são registradas de forma pulverizada (planilhas, e-mails, chamados), impossibilitando auditoria eficiente e comprometendo a rastreabilidade exigida por obrigações societárias (Lei das S.A.) e fiscais.
  3. **Exposição a erros operacionais:** A ausência de validações de negócio automatizadas — como impedir a sobreposição de vigências conflitantes ou a criação de alíquotas duplicadas para o mesmo tributo e região — transfere para o analista fiscal toda a responsabilidade de detectar inconsistências, em um ambiente de complexidade crescente.
* Para mitigar esses riscos e preparar a companhia para a nova realidade tributária, institui-se este projeto tático, desdobrado do programa PRJ-FIN-2026-0001, com o objetivo central de **prover ao time de Finanças um Portal Corporativo de Gestão Tributária** que concentre, valide e rastreie todas as tabelas de impostos da organização, eliminando a dependência técnica nas operações de rotina e estabelecendo governança plena sobre o patrimônio fiscal da companhia.

------------------------------

## 2. Objetivos Principais de Negócio

* Autonomia Operacional do Time de Finanças: Eliminar a dependência do time técnico para ajustes fiscais de rotina, permitindo que analistas fiscais e contadores realizem diretamente — com validações de negócio automáticas e dentro de perímetros de acesso bem definidos — a criação, edição e desativação de alíquotas, regimes, classificações e benefícios fiscais. O tempo médio entre a identificação de uma necessidade de ajuste e sua efetivação deve ser reduzido de dias para minutos.
* Fonte Única da Verdade Fiscal (Single Source of Truth): Consolidar todas as tabelas tributárias da companhia — alíquotas, regimes, isenções, incentivos, classificações de produtos e serviços — em uma única plataforma corporativa, eliminando a pulverização de informações entre planilhas, e-mails e sistemas desconectados que caracteriza o cenário atual.
* Governança e Rastreabilidade Societária: Estabelecer trilha de auditoria completa para toda e qualquer alteração nas tabelas fiscais, registrando: (a) o responsável pela alteração, (b) a data e hora exatas, (c) o valor anterior e o novo valor, e (d) a justificativa de negócio. Este registro atende diretamente aos requisitos de controle interno da Lei das S.A. e de marcos regulatórios como a SOX, além de prover defesa documental em caso de fiscalizações.
* Blindagem contra Erros de Configuração Fiscal: Incorporar na plataforma as validações de negócio que hoje dependem exclusivamente do conhecimento tácito dos analistas — como prevenção de conflitos de vigência, integridade entre tributos relacionados e consistência de regras de transição entre regimes — reduzindo a probabilidade de erros que possam resultar em cálculos incorretos de impostos nas operações de venda e faturamento.
* Prontidão para o Período Híbrido (2029–2032): Estruturar a gestão tributária de forma a comportar simultaneamente as tabelas do regime antigo (ICMS, ISS, PIS, COFINS, IPI) e do novo regime (CBS, IBS, IS), permitindo que o time de Finanças opere com clareza e segurança durante os 4 anos de convivência dos dois modelos.
* Preparação para Expansão Funcional: Estabelecer uma plataforma de gestão tributária que sirva como fundação para futuras capacidades — como dashboards gerenciais de KPIs fiscais, simuladores de impacto de novas alíquotas, e conciliação de créditos — previstas no roadmap do programa PRJ-FIN-2026-0001.

------------------------------

## 3. Escopo de Negócio do Projeto (Frentes de Atuação)

### 🟢 O que está incluído (In-Scope)

#### Frente 1: Portal de Gestão Tributária (Interface do Time de Finanças)

* Disponibilização de uma plataforma web corporativa, acessível a partir dos terminais e tablets do time de Finanças, com os seguintes módulos de negócio:

  **Módulo 1 — Painel de Alíquotas Vigentes:**
  * Visão gerencial consolidada de todas as alíquotas ativas na companhia, com filtros por tributo (CBS, IBS, IS, ICMS, ISS, PIS, COFINS, IPI), por unidade federativa, por município de destino (código IBGE), por período de vigência e por status (vigente, programada para início futuro, expirada).
  * Indicadores visuais de integridade: sinalização de alíquotas prestes a expirar, alíquotas sem substituição programada no Período Híbrido, e conflitos potenciais detectados.

  **Módulo 2 — Cadastro e Manutenção de Alíquotas:**
  * Formulários de negócio para criação, edição e desativação de alíquotas, regimes tributários e classificações fiscais, com validações automáticas de negócio que impeçam:
    - Duas alíquotas conflitantes para o mesmo tributo, mesma região e mesmo período.
    - Criação de alíquotas com data de término anterior à data de início.
    - Associação de alíquota a uma classificação fiscal (NCM/NBS/CClassTrib) inexistente na base.
    - Desativação de alíquota do regime antigo sem a correspondente alíquota de substituição do novo regime durante o Período Híbrido.

  **Módulo 3 — Gestão de Classificações e Regimes:**
  * Cadastro centralizado das classificações fiscais de produtos e serviços (NCM, NBS, CClassTrib, CFOP) que servem como base para a aplicação das alíquotas.
  * Manutenção dos regimes tributários aplicáveis (Lucro Real, Lucro Presumido, Simples Nacional), seus benefícios associados e regras de transição.

  **Módulo 4 — Linha do Tempo e Auditoria:**
  * Visualização cronológica completa de todas as alterações realizadas em qualquer tabela fiscal, permitindo ao gestor fiscal e ao controller navegar pelo histórico e comparar versões sucessivas.
  * Filtros por período, por usuário responsável, por tipo de alteração e por entidade impactada.

  **Módulo 5 — Importação e Exportação em Lote:**
  * Carga massiva de alíquotas a partir de planilhas padronizadas (Excel), permitindo a atualização rápida de tabelas completas — por exemplo, quando da publicação oficial das alíquotas municipais de IBS pelo Comitê Gestor.
  * Exportação das tabelas vigentes para relatórios de conformidade, auditorias externas e prestações de contas a órgãos reguladores.

  **Módulo 6 — Administração de Acessos:**
  * Gestão de usuários do portal e definição de perfis de acesso: Administrador Fiscal (acesso total), Analista Fiscal (criação/edição), Auditor/Controller (apenas leitura e consulta de trilha de auditoria).
  * Rastreabilidade individual: toda alteração é vinculada ao usuário autenticado que a realizou.

#### Frente 2: Governança e Controles Internos

* Implementação de fluxos de aprovação para alterações de alto impacto (ex: alterações em alíquotas que afetem faturamento acima de determinado patamar), com dupla validação antes da efetivação.
* Trilha de auditoria societária completa, contemplando todos os requisitos da Lei das S.A. e melhores práticas de controles internos (framework COSO), com retenção de registros pelo prazo legal.
* Política de segregação de funções: garantia de que o perfil que cria uma alíquota não seja o mesmo que a aprova, e que o perfil de auditoria tenha visibilidade plena mas nenhuma capacidade de alteração.
* Relatórios gerenciais de governança: sumário mensal de alterações realizadas, por tipo de tributo, por responsável e por justificativa, para apresentação ao Comitê Fiscal e à Controladoria.

#### Frente 3: Estruturação da Base de Dados Tributária

* Consolidação e saneamento de todas as tabelas fiscais atualmente dispersas em planilhas, e-mails e sistemas legados, estabelecendo uma base única, íntegra e validada como fundação do portal.
* Carga inicial (seed) das tabelas vigentes no momento do go-live, incluindo:
  - Alíquotas interestaduais e internas de ICMS por UF e por NCM.
  - Alíquotas de ISS por município para os serviços prestados pela companhia (matriz em Santana de Parnaíba/SP).
  - Alíquotas de PIS e COFINS nos regimes cumulativo e não cumulativo.
  - Alíquotas-piloto de CBS (0,9%) e IBS (0,1%) para o ano-calendário 2026.
  - Tabelas de classificação fiscal (NCM, NBS, CFOP) aplicáveis ao portfólio da companhia.
* Preparação da estrutura de dados para o Período Híbrido: mapeamento de correlação entre tributos antigos e novos, permitindo que o time de Finanças visualize e gerencie a transição (ex: "ICMS interestadual 12% será substituído por IBS estadual X% a partir de 2029").

### 🔴 O que NÃO está incluído (Out-of-Scope)

* O cálculo de impostos em si — simulações de checkout, cálculo de CBS/IBS/IS no momento da venda, motor de DIFAL e apuração de créditos. Estas capacidades permanecem sob responsabilidade da calculadora corporativa de impostos já existente e em operação.
* A integração automatizada com sistemas externos de órgãos públicos (Comitê Gestor do IBS, SEFAZ, Receita Federal) para captura de alíquotas oficiais. A obtenção e interpretação das publicações oficiais continuará sendo realizada pelo time fiscal, e os valores serão inseridos no portal.
* A revisão do mérito fiscal das alíquotas praticadas. O portal provê a ferramenta de gestão; a definição das alíquotas corretas, dos regimes aplicáveis e das estratégias de creditamento permanece como atribuição do time de Finanças e do Comitê Fiscal.
* O desenvolvimento de aplicativo mobile nativo (iOS/Android). O portal web será responsivo e acessível via navegadores em tablets, sem necessidade de instalação de aplicativos.
* A gestão de obrigações acessórias (declarações, guias de recolhimento, EFD). O escopo limita-se à administração das tabelas de alíquotas e regras; a geração de declarações fiscais é escopo de outros projetos do programa PRJ-FIN-2026-0001.

------------------------------

## 4. Macro-Cronograma e Rollout Estratégico de Negócio

### 4.1 Ciclo Tático de Implantação

O projeto está organizado em quatro entregas sequenciais, cada uma correspondendo a um bloco de capacidades de negócio que amplia progressivamente a autonomia do time de Finanças:

```text
[ENTREGA 1: Portal — Gestão Básica] ──► [ENTREGA 2: Governança e Auditoria] ──► [ENTREGA 3: Operações em Escala] ──► [ENTREGA 4: Portal Completo]
```

### 4.2 Detalhamento das Entregas

| Fase | Duração Estimada | Descrição de Negócio | Critério de Aceite |
|:---|:---|:---|:---|
| **Fase 0 — Fundamentação** | 3 a 4 semanas | Mapeamento completo das tabelas fiscais atuais (planilhas, sistemas, documentos); levantamento das necessidades do time de Finanças (entrevistas com analistas, controller e gerentes fiscais); definição do modelo conceitual de dados tributários; **design UX/UI do portal em ciclo formal de duas etapas:** _(i)_ criação de wireframes, protótipos e design system, _(ii)_ apresentação e obtenção de aprovação formal dos stakeholders (PO, Comitê Fiscal, usuários-chave); paralelamente, os times técnicos realizam as definições de arquitetura, infraestrutura e segurança necessárias para iniciar o desenvolvimento | Modelo de dados validado pelo time fiscal; design UX/UI aprovado formalmente pelos stakeholders; wireframes e protótipos validados pelos usuários finais; inventário completo das tabelas existentes |
| **Fase 1 — Portal: Gestão Básica de Alíquotas** | 6 a 8 semanas | Implantação do portal com os Módulos 1 (Painel de Alíquotas), 2 (Cadastro e Manutenção de Alíquotas) e 3 (Gestão de Classificações e Regimes). Carga inicial das tabelas vigentes. Treinamento do time de Finanças. | Analistas fiscais realizando operações de rotina diretamente no portal, sem abertura de chamados técnicos. Zero conflitos de vigência detectados em 30 dias de operação. |
| **[ENTREGA 1]** Go-Live Portal — Gestão Básica | — | Liberação do portal para o time de Finanças como ferramenta oficial de administração tributária. A ferramenta anterior utilizada pelo time técnico permanece disponível como contingência durante o período de adaptação. | Portal operacional e adotado como ferramenta primária pelo time fiscal. Redução mensurável no volume de chamados de "ajuste de alíquota". |
| **Fase 2 — Governança e Auditoria Fiscal** | 4 a 6 semanas | Implantação do Módulo 6 (Administração de Acessos e Perfis) e do Módulo 4 (Linha do Tempo e Auditoria), estabelecendo a camada completa de controle de acesso com segregação de funções e a trilha de auditoria imutável. | Perfis de acesso segregados 100% operacionais. Trilha de auditoria registrando todas as alterações. Controller consegue rastrear qualquer mudança em menos de 2 minutos. |
| **[ENTREGA 2]** Consolidação da Governança | — | Portal completo em funcionalidades de governança. Descontinuação do acesso à ferramenta administrativa legada pelo time de Finanças (a ferramenta legada permanece apenas para consulta histórica). | Time fiscal opera exclusivamente via portal. Zero chamados de auditoria ou conformidade relacionados à gestão de alíquotas. |
| **Fase 3 — Operações Fiscais em Escala** | 4 a 6 semanas | Implantação dos fluxos de aprovação em duas etapas para alterações de alto impacto financeiro e do Módulo 5 (Importação e Exportação de Alíquotas em Lote), permitindo carga massiva de publicações oficiais. | Alteração acima do patamar de materialidade bloqueada até aprovação. Carga de planilha com 5.570 municípios processada e validada no mesmo dia útil. |
| **[ENTREGA 3]** Portal — Operações em Escala | — | Portal com capacidade de processar alterações em volume e com governança proporcional ao risco. Time fiscal apto a absorver publicações oficiais de alíquotas em escala nacional. | Carga de IBS municipal simulada com sucesso. Fluxos de aprovação operacionais e auditáveis. |
| **Fase 4 — Expansão Funcional: Inteligência e Analytics** | 6 a 8 semanas | Implantação dos relatórios gerenciais de governança, dashboards de KPIs fiscais e preparação completa da estrutura para suporte ao Período Híbrido (2029–2032). | Relatório mensal gerado automaticamente. Dashboards de KPIs aprovados pelo CFO. Cenário híbrido completo testado e aprovado pelo Comitê Fiscal. |
| **[ENTREGA 4]** Portal Completo | — | Portal como plataforma única e completa de gestão tributária corporativa, com todas as funcionalidades previstas operacionais e time de Finanças plenamente autônomo. | NPS interno ≥ 70, redução ≥ 80% no tempo médio de ajustes fiscais, zero incidentes de erro de alíquota com causa raiz em configuração manual. |

### 4.3 Ciclo de Longo Prazo (Alinhamento com Programa Pai)

| Período | Contexto Regulatório | Marco do Portal |
|:---|:---|:---|
| **Jul–Set 2026** | Alíquotas de teste CBS (0,9%) e IBS (0,1%) vigentes | Entrega 1: Portal com cadastro e painel de alíquotas em produção |
| **Out–Dez 2026** | Preparação para vigência plena da CBS em 2027 | Entrega 2: Governança consolidada — acessos e trilha de auditoria |
| **Jan–Mar 2027** | Extinção de PIS/COFINS, início pleno da CBS | Entrega 3: Operações em escala — fluxos de aprovação e carga em lote |
| **Abr–Jun 2027** | Operação estabilizada no novo regime federal | Entrega 4: Portal completo — dashboards, relatórios e Período Híbrido |
| **2029–2032** | Período Híbrido — convivência ICMS/ISS + IBS | Portal opera com visão dual de regimes, suportando a dupla gestão |
| **2033** | Full IVA Dual (apenas CBS e IBS) | Descontinuação das tabelas legadas no portal; manutenção apenas do IVA Dual |

------------------------------

## 5. Estrutura de Liderança e Partes Interessadas (Stakeholders)

* Patrocinador Executivo (Sponsor): Diretor Financeiro (CFO) — Responsável por aprovar o investimento, validar as diretrizes de governança fiscal e garantir a adoção da plataforma pelo time de Finanças como ferramenta oficial de trabalho.
* Comitê Fiscal e Jurídico de Negócios: Especialistas em direito tributário e controladoria — Responsáveis por ditar as regras de negócio que serão incorporadas como validações automáticas no portal (conflitos de vigência, transição de regimes, integridade entre tributos), validar as alíquotas carregadas e auditar a integridade da base.
* Gerência de Controladoria e Compliance: Responsável por definir os requisitos de trilha de auditoria, aprovar os relatórios de governança e garantir que o portal atenda aos controles internos exigidos pela Lei das S.A. e pelo framework COSO.
* Time de Finanças — Analistas Fiscais e Contadores: Usuários finais do portal. Responsáveis pela operação diária de manutenção das tabelas tributárias, pela validação da usabilidade e completude funcional, e por reportar necessidades de evolução da plataforma.
* Gerência de Tesouraria: Interessada indireta — a precisão das alíquotas geridas no portal impacta diretamente as projeções de fluxo de caixa, especialmente no cenário de split payment onde CBS e IBS são liquidados instantaneamente.
* Lideranças de Canais e Mercado (Diretores e Gerentes Comerciais): Interessados indiretos — as alíquotas administradas no portal determinam os preços finais praticados nos canais de venda. A rapidez na atualização de alíquotas impacta diretamente a competitividade comercial.
* PMO Corporativo (Programa PRJ-FIN-2026-0001): Responsável por garantir que este projeto permaneça alinhado às diretrizes, cronogramas e restrições regulatórias do programa guarda-chuva, e por arbitrar conflitos de priorização com outras iniciativas do programa.

------------------------------

## 6. Premissas e Restrições de Negócio

* Premissas:
  * O time de Finanças terá disponibilidade para participar ativamente das sessões de prototipação, validação de telas (UAT) e treinamentos, alocando ao menos 20% do tempo dos analistas fiscais seniores durante as fases de desenvolvimento.
  * As tabelas fiscais atualmente em uso (planilhas, documentos, sistemas legados) estão suficientemente organizadas para permitir a carga inicial da base do portal com esforço viável de saneamento.
  * O Comitê Gestor do IBS disponibilizará as alíquotas de referência por município em formato que permita a importação em lote para o portal (planilha ou tabela estruturada), dispensando digitação manual de 5.570 municípios.
  * As definições de alíquotas definitivas de CBS, IBS e IS seguirão o calendário constitucional, permitindo que o portal seja abastecido com dados oficiais antes da entrada em vigor de cada fase.
  * A companhia manterá uma calculadora corporativa de impostos dedicada e segregada — o portal de gestão não precisa calcular tributos, apenas administrar as tabelas que alimentam a calculadora.
* Restrições:
  * Calendário Regulatório: O projeto está rigidamente atrelado aos prazos constitucionais de transição tributária. As tabelas de CBS e IBS para o ano-calendário 2026 (alíquotas de teste) precisam estar operacionais no portal o mais rapidamente possível para que o time de Finanças inicie a familiarização com a nova ferramenta antes da vigência plena.
  * Estabilidade das Operações de Venda e Faturamento: Nenhuma atividade de implantação do portal pode interromper ou degradar a capacidade de calcular impostos nas operações comerciais em andamento. A calculadora corporativa é serviço crítico — o portal é ferramenta de suporte.
  * Incerteza de Alíquotas Definitivas: As alíquotas finais de CBS, IBS e IS somente serão conhecidas após a conclusão dos processos legislativo e regulatório, exigindo flexibilidade no portal para ajustes e recargas de tabelas.
  * Capacidade do Time Fiscal: O projeto parte da premissa de que o time de Finanças absorverá a operação do portal como parte de suas atribuições regulares, sem incremento de headcount dedicado à administração da ferramenta.
  * Orçamento Corporativo: O projeto utiliza recursos do orçamento do programa PRJ-FIN-2026-0001; despesas extraordinárias precisam de aprovação específica do CFO.

------------------------------

## 7. Gestão de Riscos do Projeto de Negócios

| Risco de Negócio Identificado | Impacto | Estratégia de Mitigação Corporativa |
|---|---|---|
| **Erro de Alíquota com Impacto em Vendas:** Uma alíquota incorreta cadastrada no portal propagar-se para os sistemas de venda e faturamento, gerando preços errados para clientes e potencial passivo fiscal. | Crítico | Instituir fluxo de dupla validação para alterações em alíquotas que impactem mais de R$ 100 mil em faturamento mensal. Implementar indicadores de "alíquota alterada nas últimas 24h" visível nos dashboards comerciais para detecção rápida de anomalias. |
| **Baixa Adoção pelo Time Fiscal:** Os analistas fiscais resistirem à nova ferramenta e persistirem no uso de planilhas e chamados técnicos, anulando o benefício esperado de autonomia e rastreabilidade. | Alto | Envolver o time fiscal desde a Fase 0 (prototipação e wireframes), incorporar feedback em ciclos curtos (demo quinzenal), realizar treinamento hands-on antes do go-live e estabelecer meta de adoção vinculada à avaliação de desempenho do time. |
| **Divergência entre Base do Portal e Base Ativa:** Durante o período inicial de operação, as alíquotas cadastradas no portal divergirem das alíquotas efetivamente utilizadas nos cálculos, gerando falsa sensação de controle. | Alto | Estabelecer reconciliação diária automatizada entre as tabelas do portal e as tabelas em uso pela calculadora de impostos, com alertas imediatos para qualquer divergência detectada e procedimento de correção documentado. |
| **Sobrecarga do Time Fiscal no Período Híbrido:** Entre 2029 e 2032, a necessidade de manter simultaneamente tabelas dos regimes antigo e novo dobrar a carga de trabalho dos analistas, comprometendo a qualidade dos cadastros. | Alto | Projetar o portal desde já com suporte explícito à gestão dual de regimes: indicadores visuais de "regime em extinção" vs. "regime em vigor", mapeamento automático de correlação entre tributos antigos e novos, e funcionalidade de desativação em lote ao final do período de transição. |
| **Atraso na Disponibilização de Alíquotas Oficiais pelo Comitê Gestor:** O Comitê Gestor do IBS atrasar a publicação das tabelas de alíquotas por município, impedindo a carga completa do portal antes da vigência. | Médio | Desenvolver funcionalidade de carga em lote (Módulo 5) como prioridade da Fase 1, permitindo que tão logo as alíquotas sejam publicadas — mesmo que em formato não ideal — o time fiscal consiga importá-las em horas, não semanas. |
| **Perda de Conhecimento Tácito:** As validações de negócio que hoje residem na experiência dos analistas fiscais seniores não serem adequadamente capturadas nas regras automáticas do portal, resultando em um sistema que não previne todos os erros possíveis. | Alto | Realizar sessões estruturadas de elicitação de conhecimento com os analistas seniores durante a Fase 0, documentando cada regra de validação no formato "Se [condição], então [restrição], porque [impacto fiscal]". Revisar trimestralmente as regras implementadas versus o conhecimento do time. |
| **Obsolescência da Base durante a Transição Regulatória:** Mudanças legislativas de última hora — comuns em reformas tributárias — invalidarem premissas incorporadas ao portal, exigindo retrabalho significativo. | Médio | Projetar o modelo de dados com flexibilidade para acomodar variações regulatórias (ex: novos tributos, alterações de competência, mudanças de alíquotas de referência). Manter o Comitê Fiscal como ponto focal de monitoramento legislativo, com reuniões mensais de alinhamento durante todo o projeto. |

------------------------------

## 8. Critérios Globais de Sucesso do Projeto

* Autonomia Comprovada do Time de Finanças: Após 90 dias do go-live da Entrega 1, o volume de chamados técnicos para ajustes de alíquotas, classificações e regimes fiscais deve ser reduzido em no mínimo 80% em relação à linha de base pré-projeto. Métrica: contagem mensal de tickets abertos pelo time fiscal para a área técnica com assunto "alíquota", "tabela fiscal", "cadastro de imposto" ou similar.
* Conformidade sem Penalidades: Zero notificações, autuações ou apontamentos de auditoria (interna ou externa) relacionados a erros de configuração de alíquotas nos 12 meses seguintes a cada entrega do projeto. Métrica: relatório trimestral de conformidade emitido pela Controladoria.
* Trilha de Auditoria Íntegra: 100% das alterações em tabelas fiscais realizadas através do portal possuem registro completo de auditoria (responsável, data/hora, valor anterior, valor novo, justificativa), e nenhuma alteração pode ser efetuada sem que este registro seja gerado automaticamente. Métrica: auditoria amostral trimestral realizada pela Controladoria com índice de conformidade de 100%.
* Satisfação do Time de Finanças: O portal atinge NPS interno ≥ 70 junto aos usuários do time fiscal (analistas, contadores, controller) após 90 dias de uso, com pontuação SUS (System Usability Scale) ≥ 75. Métrica: pesquisa NPS e SUS aplicada ao final de cada entrega.
* Adoção como Ferramenta Oficial: Após 6 meses do go-live, o portal é a ferramenta primária para consulta e gestão de tabelas fiscais, com a ferramenta anterior acessada apenas para consulta histórica de registros anteriores à implantação. Métrica: proporção de operações de gestão realizadas no portal vs. outros meios.
* Precisão das Tabelas Fiscais: Zero incidentes de erro de precificação ou faturamento com causa raiz atribuída a divergência entre a alíquota configurada no portal e a alíquota oficial vigente publicada pelo ente tributante. Métrica: relatório mensal de reconciliação portal × calculadora × publicações oficiais.
* Prontidão para o Período Híbrido: O portal está apto a gerenciar simultaneamente tabelas dos dois regimes (antigo e novo) até dezembro de 2028, com funcionalidades de mapeamento de transição, visão dual e desativação progressiva operacionais e documentadas. Métrica: teste de carga do cenário híbrido completo executado com sucesso.

------------------------------

## 9. Dependências Externas Críticas

| Dependência | Descrição | Impacto no Projeto |
|:---|:---|:---|
| **Comitê Gestor do IBS** | Disponibilização de alíquotas estaduais e municipais, regras de creditamento e calendários oficiais de transição | Impacto direto no volume e cronograma de carga de dados no portal — sem as alíquotas oficiais, o portal opera com dados parciais |
| **Publicações da Receita Federal (CBS e IS)** | Definição das alíquotas federais de CBS e da lista de produtos sujeitos ao Imposto Seletivo (com respectivas alíquotas) | Necessário para a carga completa das tabelas federais no portal antes de cada fase de vigência |
| **Legislação Complementar Federal** | Publicação das Leis Complementares que regulamentam a reforma tributária, definindo regras de transição, creditamento e obrigações | Define as validações de negócio que o portal deve implementar para prevenir erros de configuração |
| **Time de Finanças (disponibilidade para o projeto)** | Participação ativa dos analistas fiscais seniores nas fases de levantamento, prototipação, validação (UAT) e treinamento | Bloqueante para Fase 0 (levantamento de requisitos) e Fase 1 (validação do portal básico) |
| **Órgãos Estaduais e Municipais (SEFAZ e Secretarias de Fazenda)** | Definição das alíquotas próprias de ICMS (no regime atual) e IBS (no novo regime) em cada ente federativo | Impacto direto na completude e atualidade das tabelas do portal, especialmente para operações interestaduais |
| **Programa PRJ-FIN-2026-0001 (cronograma e prioridades)** | Alinhamento com o cronograma do programa pai, que dita quando cada tributo entra em transição e quais as prioridades de negócio | Impacto no sequenciamento das entregas — o portal precisa estar pronto antes de cada marco regulatório do programa pai |
| **Calculadora Corporativa de Impostos (ferramenta legada de administração)** | Clareza sobre quais informações fiscais são atualmente administradas e como se relacionam com os cálculos | Necessário para o mapeamento completo das tabelas a serem migradas para o portal (Fase 0) |

------------------------------

## 10. Aprovação e Governança

Este Project Charter é o documento de referência para a construção do Portal Corporativo de Gestão Tributária, que proverá ao time de Finanças a autonomia, a rastreabilidade e os controles necessários para administrar o patrimônio fiscal da companhia durante a maior transformação tributária da história recente do país.

O projeto está subordinado ao Program Charter do PRJ-FIN-2026-0001, do qual herda as diretrizes estratégicas, restrições regulatórias e premissas de calendário. Qualquer alteração de escopo, prazo ou recurso deverá ser submetida ao PMO do programa pai e ao Sponsor Executivo (CFO).

* Versão: 1.1
* Data de Criação: 08 de Julho de 2026
* Última Atualização: 10 de Julho de 2026 (revisão da Fase 0 — inclusão do ciclo formal de UX/UI em duas etapas e ajuste de duração para 3 a 4 semanas, por recomendação dos times técnicos acatada pelo time de Finanças)
* Próxima Revisão Programada: Agosto de 2026 (após conclusão da Fase 0 — Fundamentação e Prototipação)
* Programa Pai: [PRJ-FIN-2026-0001 — Adequação Corporativa à Reforma Tributária Nacional](../PRJ-FIN-2026-0001-REFORMA-TRIBUTARIA-2026-CORPORATIVO/01-PROJECT-CHARTER.md)

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: draft-project-charter, brainstorming, stakeholder-analysis.*
