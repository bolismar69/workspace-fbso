# Lista de Épicos do Projeto
* **Projeto:** Portal Corporativo de Gestão Tributária — Autonomia do Time de Finanças na Administração de Impostos
* **Código:** PRJ-FIN-2026-0002-ADMIN-TRIBUTOS-CORPORATIVOS
* **Status:** Pronto para Refinamento Funcional
* **Responsáveis:** Product Owner (PO) — Portal de Gestão Tributária, Comitê Fiscal
* **Referência:** [02-BUSINESS-REQUIREMENTS.md](./02-BUSINESS-REQUIREMENTS.md)
* [INDEX] Deriva de: [02-BUSINESS-REQUIREMENTS.md](./02-BUSINESS-REQUIREMENTS.md) — Seção 3 (Requisitos de Negócio)

------------------------------
## 1. Objetivo do Documento

Este documento define os Épicos (Grandes Blocos de Capacidades de Negócio) necessários para prover ao time de Finanças o Portal Corporativo de Gestão Tributária, atendendo aos 10 Requisitos de Negócio (BR-01 a BR-10) estabelecidos no [02-BUSINESS-REQUIREMENTS.md](./02-BUSINESS-REQUIREMENTS.md).

Cada épico corresponde a exatamente uma **entrega** do projeto, em sequência lógica de construção: primeiro estabelece-se a base de cadastro fiscal (Entrega 1), depois os controles de acesso e rastreabilidade (Entrega 2), em seguida as operações de escala (Entrega 3), e por fim a camada de inteligência analítica (Entrega 4). A numeração sequencial (`01` a `04`) reflete esta ordem.

------------------------------
## 2. Visão Geral da Jornada do Projeto

```text
┌──────────────────────────────────────────────────────────────────────────────────────────────┐
│                         ÉPICO 01 / ENTREGA 1: MOTOR DE CADASTRO FISCAL                        │
│  [ ALÍQUOTAS ] ──────► [ VALIDAÇÕES ] ──────► [ CLASSIFICAÇÕES E REGIMES ]                    │
│      BR-01                 BR-02                      BR-03                                   │
│  Cadastro centralizado  Regras automáticas       NCM, NBS, CClassTrib,                       │
│  de todos os tributos   de conflitos             CFOP e regimes fiscais                       │
│                                                                                               │
│  Módulos: M1 (Painel), M2 (Cadastro), M3 (Classificações e Regimes)                           │
└──────────────────────────────────────────────────────────────────────────────────────────────┘

┌──────────────────────────────────────────────────────────────────────────────────────────────┐
│                   ÉPICO 02 / ENTREGA 2: CONTROLE DE ACESSO E RASTREABILIDADE                  │
│  [ PERFIS E ACESSOS ] ──────────────────────► [ TRILHA DE AUDITORIA ]                         │
│          BR-04                                          BR-05                                 │
│  Administração de usuários e                Registro imutável de toda                        
│  segregação de funções                      alteração em tabela fiscal                        │
│                                                                                               │
│  Módulos: M6 (Administração de Acessos), M4 (Linha do Tempo e Auditoria)                      │
└──────────────────────────────────────────────────────────────────────────────────────────────┘

┌──────────────────────────────────────────────────────────────────────────────────────────────┐
│                       ÉPICO 03 / ENTREGA 3: OPERAÇÕES FISCAIS EM ESCALA                       │
│  [ FLUXOS DE APROVAÇÃO ] ──────────────────► [ CARGA E EXPORTAÇÃO EM LOTE ]                   │
│           BR-06                                          BR-07                                │
│  Dupla validação para alterações            Importação massiva de alíquotas                   
│  de alto impacto financeiro                 e exportação de tabelas vigentes                  
│                                                                                               │
│  Módulos: — (funcionalidades transversais), M5 (Importação/Exportação em Lote)                 │
└──────────────────────────────────────────────────────────────────────────────────────────────┘

┌──────────────────────────────────────────────────────────────────────────────────────────────┐
│                     ÉPICO 04 / ENTREGA 4: INTELIGÊNCIA FISCAL E ANALYTICS                     │
│  [ RELATÓRIOS ] ──────► [ DASHBOARDS ] ──────► [ PERÍODO HÍBRIDO ]                            │
│      BR-08                 BR-09                      BR-10                                   │
│  Relatórios gerenciais  KPIs visuais para         Dupla gestão de regimes                    
│  de governança          CFO e Comitê Fiscal       antigo e novo (2029–2032)                    │
│                                                                                               │
│  Módulos: — (funcionalidades transversais)                                                    │
└──────────────────────────────────────────────────────────────────────────────────────────────┘
```

------------------------------
## 3. Épicos

### 🔢 ÉPICO 01: Motor de Cadastro Fiscal — Alíquotas, Validações e Classificações

> **Entrega:** Entrega 1 — Portal: Gestão Básica de Alíquotas
> **Duração estimada:** 6 a 8 semanas (Fase 1 do Project Charter)
> **Partes Interessadas Primárias:** Analistas Fiscais (operação diária), Comitê Fiscal (validação de regras), Gerente Fiscal (supervisão)

* **Domínio de Negócio:** Fundação do Portal — o cadastro das tabelas fiscais é a capacidade central sem a qual nenhuma outra funcionalidade do portal tem propósito.

* **Descrição de Alto Nível:** Estabelecer a base de dados tributária centralizada da companhia, permitindo que o time de Finanças cadastre, edite e desative alíquotas de todos os tributos (CBS, IBS, IS, ICMS, ISS, PIS, COFINS, IPI) com validações automáticas de negócio que impeçam configurações conflitantes. Inclui também o cadastro das classificações fiscais (NCM, NBS, CClassTrib, CFOP) e dos regimes tributários (Lucro Real, Lucro Presumido, Simples Nacional) que servem como base para aplicação das alíquotas.

* **Justificativa de Negócio:** Hoje as tabelas fiscais estão pulverizadas em planilhas, e-mails e conhecimento tácito dos analistas. Sem uma base centralizada e validada, qualquer funcionalidade de governança ou inteligência fiscal construída sobre dados inconsistentes herdará esses problemas. Este épico entrega a "Fonte Única da Verdade Fiscal" — pré-requisito para todos os demais épicos.

* **Requisitos Vinculados ([02-BUSINESS-REQUIREMENTS.md](./02-BUSINESS-REQUIREMENTS.md)):** BR-01, BR-02, BR-03.

* **Capacidades Esperadas (Features do Produto):**

| Feature | Descrição de Negócio | Módulo |
|:---|:---|:---|
| **01.1 — Painel de Alíquotas Vigentes** | Visão consolidada de todas as alíquotas ativas, com filtros por tributo, UF, município, período e status. Indicadores visuais de integridade: alíquotas prestes a expirar, conflitos potenciais, ausência de substituição no Período Híbrido. | M1 |
| **01.2 — Cadastro e Manutenção de Alíquotas** | Criação, edição e desativação de alíquotas com validações automáticas de conflito (sobreposição de vigência, referência fiscal inexistente, transição inconsistente entre regimes). | M2 |
| **01.3 — Gestão de Classificações e Regimes** | Cadastro centralizado de NCM, NBS, CClassTrib, CFOP e dos regimes tributários aplicáveis (Lucro Real, Lucro Presumido, Simples Nacional), com vinculação às alíquotas. | M3 |

* **Critérios de Sucesso do Épico:**
  1. Um analista fiscal consegue cadastrar uma nova alíquota de IBS para qualquer município brasileiro, com validação automática de conflitos, sem intervenção técnica.
  2. O painel reflete em tempo real todas as alíquotas vigentes, com sinalização clara de status.
  3. Nenhuma classificação fiscal (NCM/NBS) utilizada nas operações da companhia permanece fora do portal após a conclusão deste épico.

---

### 🔐 ÉPICO 02: Controle de Acesso e Rastreabilidade Fiscal

> **Entrega:** Entrega 2 — Governança e Auditoria Fiscal
> **Duração estimada:** 4 a 6 semanas (Fase 2 do Project Charter)
> **Partes Interessadas Primárias:** Controladoria e Compliance (requisitos de trilha e segregação), Gerente Fiscal (gestão de acessos), Auditoria Interna (consulta à trilha)

* **Domínio de Negócio:** Governança — quem pode fazer o quê no portal e como cada ação é registrada para fins de auditoria e conformidade.

* **Descrição de Alto Nível:** Implementar a camada de controle de acesso com segregação de funções (Administrador Fiscal, Analista Fiscal, Auditor/Controller) e a trilha de auditoria completa e imutável que registra toda alteração em tabela fiscal — contemplando identificação do usuário, data e hora, valor anterior, valor novo e justificativa de negócio. Estas duas capacidades formam o núcleo de governança do portal: a primeira define quem pode agir, a segunda registra tudo o que foi feito.

* **Justificativa de Negócio:** A autonomia operacional do time de Finanças — objetivo central do projeto — só é aceitável sob governança. Sem perfis de acesso segregados, qualquer analista poderia alterar qualquer alíquota sem supervisão. Sem trilha de auditoria, a companhia fica exposta a riscos de conformidade (Lei das S.A., COSO, SOX) e sem defesa documental em fiscalizações. Estas duas capacidades são agrupadas na mesma entrega porque uma sem a outra gera valor incompleto: acessos sem trilha não provê rastreabilidade; trilha sem acessos não garante que as ações registradas foram realizadas por pessoas autorizadas.

* **Requisitos Vinculados ([02-BUSINESS-REQUIREMENTS.md](./02-BUSINESS-REQUIREMENTS.md)):** BR-04, BR-05.

* **Capacidades Esperadas (Features do Produto):**

| Feature | Descrição de Negócio | Módulo |
|:---|:---|:---|
| **02.1 — Administração de Acessos e Perfis** | Gestão de usuários do portal com três perfis de acesso segregados: Administrador Fiscal, Analista Fiscal e Auditor/Controller. Cada perfil tem privilégios claramente delimitados e incompatíveis entre si. | M6 |
| **02.2 — Trilha de Auditoria e Linha do Tempo** | Registro imutável de toda alteração em tabela fiscal com: usuário, data/hora, entidade, valor anterior, valor novo e justificativa. Visualização cronológica com comparação entre versões sucessivas. | M4 |

* **Critérios de Sucesso do Épico:**
  1. Nenhum usuário consegue realizar ações fora de seu perfil de acesso (ex: Auditor não consegue alterar alíquotas; Analista Fiscal não consegue aprovar as próprias alterações de alto impacto).
  2. 100% das alterações em tabelas fiscais possuem registro de auditoria automático e imutável — meta verificada por auditoria amostral mensal da Controladoria.
  3. O Controller consegue rastrear, em menos de 2 minutos, quem alterou uma alíquota específica, quando, qual era o valor anterior e qual a justificativa registrada.

---

### 📊 ÉPICO 03: Operações Fiscais em Escala

> **Entrega:** Entrega 3 — Portal: Operações em Escala
> **Duração estimada:** 4 a 6 semanas (Fase 3 do Project Charter)
> **Partes Interessadas Primárias:** Gerente Fiscal e Controller (aprovação de alto impacto), Comitê Fiscal (definição de patamares de materialidade), Analistas Fiscais (execução de cargas em lote)

* **Domínio de Negócio:** Eficiência operacional — capacidades que permitem ao time de Finanças operar com volume e agilidade, seja aprovando alterações de alto impacto com a governança adequada, seja processando milhares de alíquotas de uma só vez.

* **Descrição de Alto Nível:** Prover dois mecanismos complementares de escala: (a) o fluxo de aprovação em duas etapas para alterações de alto impacto financeiro, garantindo que decisões relevantes tenham dupla validação sem travar ajustes rotineiros; e (b) a importação e exportação de alíquotas em lote via planilhas padronizadas, permitindo que publicações oficiais — como a divulgação de alíquotas de IBS para 5.570 municípios pelo Comitê Gestor — sejam absorvidas em horas, não semanas.

* **Justificativa de Negócio:** O fluxo de aprovação (BR-06) é a resposta de governança ao risco de erro humano em alterações de grande impacto financeiro — ele adiciona controle onde o risco é maior, sem burocratizar ajustes rotineiros. A carga em lote (BR-07) é a resposta operacional ao volume: com 5.570 municípios potencialmente definindo alíquotas próprias de IBS, a digitação manual é inviável e arriscada. Estas duas capacidades, embora distintas, compartilham o propósito de permitir que o time de Finanças opere com segurança em cenários de alto volume e alto impacto.

* **Requisitos Vinculados ([02-BUSINESS-REQUIREMENTS.md](./02-BUSINESS-REQUIREMENTS.md)):** BR-06, BR-07.

* **Capacidades Esperadas (Features do Produto):**

| Feature | Descrição de Negócio | Módulo |
|:---|:---|:---|
| **03.1 — Fluxos de Aprovação para Alto Impacto** | Mecanismo de dupla validação para alterações acima do patamar de materialidade definido pelo Comitê Fiscal. O Analista Fiscal propõe, o Administrador Fiscal ou Controller aprova. Alterações abaixo do patamar seguem o fluxo normal. | — (transversal) |
| **03.2 — Importação e Exportação em Lote** | Carga massiva de alíquotas a partir de planilhas padronizadas com validação automática linha a linha, e exportação de tabelas vigentes para relatórios de conformidade e auditorias externas. | M5 |

* **Critérios de Sucesso do Épico:**
  1. Uma alteração de alíquota com impacto estimado acima do patamar de materialidade é bloqueada até que um Administrador Fiscal a aprove — e a aprovação fica registrada na trilha de auditoria.
  2. O time fiscal consegue carregar uma planilha com alíquotas de IBS para 100 municípios, e o portal processa a carga com validação automática, reportando quais registros foram aceitos e quais foram rejeitados (com o motivo específico).
  3. Uma planilha com 5.570 municípios é processada em tempo hábil para que o time fiscal confira e aprove os resultados dentro do mesmo dia útil.

---

### 📈 ÉPICO 04: Inteligência Fiscal e Analytics

> **Entrega:** Entrega 4 — Portal Completo: Expansão Funcional
> **Duração estimada:** 6 a 8 semanas (Fase 4 do Project Charter)
> **Partes Interessadas Primárias:** CFO e Comitê Fiscal (dashboards e KPIs), Controladoria (relatórios de governança), Gerente Fiscal e Analistas (operação no Período Híbrido)

* **Domínio de Negócio:** Visibilidade gerencial — capacidades que transformam os dados do portal em informação acionável para o CFO, Comitê Fiscal e Controladoria, e preparam a plataforma para o longo prazo.

* **Descrição de Alto Nível:** Prover três capacidades de inteligência: (a) relatórios gerenciais mensais de governança que sumarizam todas as alterações do período para o Comitê Fiscal e Controladoria; (b) dashboards visuais de KPIs fiscais que dão ao CFO visibilidade imediata sobre o status do patrimônio fiscal da companhia; e (c) suporte completo ao Período Híbrido (2029–2032), permitindo a gestão simultânea de tabelas dos regimes antigo e novo com indicadores visuais claros, mapeamento de transição e desativação progressiva.

* **Justificativa de Negócio:** Os relatórios (BR-08) substituem a compilação manual de dados de governança que hoje consome horas do time de Controladoria. Os dashboards (BR-09) eliminam a dependência do CFO de relatórios ad hoc do time técnico para responder perguntas básicas como "quantas alíquotas de IBS estão vigentes hoje?" e "quais municípios onde operamos ainda não têm alíquota cadastrada?". O suporte ao Período Híbrido (BR-10) não é apenas uma funcionalidade — é a garantia de que o investimento no portal permanecerá válido durante os 4 anos mais complexos da transição tributária (2029–2032).

* **Requisitos Vinculados ([02-BUSINESS-REQUIREMENTS.md](./02-BUSINESS-REQUIREMENTS.md)):** BR-08, BR-09, BR-10.

* **Capacidades Esperadas (Features do Produto):**

| Feature | Descrição de Negócio | Módulo |
|:---|:---|:---|
| **04.1 — Relatórios Gerenciais de Governança** | Sumário mensal de todas as alterações em tabelas fiscais, agrupadas por tributo, responsável e justificativa, para apresentação ao Comitê Fiscal e Controladoria. | — (transversal) |
| **04.2 — Dashboards de KPIs Fiscais** | Painéis visuais com indicadores-chave: total de alíquotas vigentes por tributo, cobertura geográfica, alterações no período, alíquotas a expirar, status de completude da base. | — (transversal) |
| **04.3 — Suporte ao Período Híbrido** | Gestão simultânea de tabelas dos regimes antigo (ICMS, ISS, PIS, COFINS, IPI) e novo (CBS, IBS, IS) com indicadores visuais de regime, mapeamento de correlação entre tributos, e funcionalidade de desativação progressiva conforme cronograma constitucional. | — (transversal) |

* **Critérios de Sucesso do Épico:**
  1. O relatório mensal de governança é gerado automaticamente e está na mesa do Comitê Fiscal até o 5º dia útil do mês seguinte, sem intervenção manual.
  2. O CFO acessa o dashboard e responde "quantas alíquotas de IBS estão vigentes e quantos municípios cobrimos?" em menos de 1 minuto, sem acionar o time técnico.
  3. Um analista fiscal consegue visualizar simultaneamente as alíquotas de ICMS (regime antigo) e IBS (novo regime) para o estado de São Paulo, com indicador claro de qual substitui qual e em que data.

------------------------------
## 4. Critérios de Sucesso Transversais do Projeto

Além dos critérios específicos de cada épico, o projeto como um todo será considerado concluído quando:

1. **Autonomia Comprovada:** O time de Finanças realiza ≥ 95% das operações de manutenção de tabelas fiscais através do portal, sem abertura de chamados técnicos (KPI A1). [INDEX] → [MATRIZ-KPI.md](./MATRIZ-KPI.md)
2. **Conformidade Plena:** Zero apontamentos de auditoria interna ou externa relacionados a erros de configuração de alíquotas nos 12 meses seguintes ao go-live da Entrega 4 (KPI G3). [INDEX] → [MATRIZ-KPI.md](./MATRIZ-KPI.md)
3. **Satisfação do Usuário:** NPS interno ≥ 70 junto ao time de Finanças após 90 dias de uso do portal completo (KPI E2). [INDEX] → [MATRIZ-KPI.md](./MATRIZ-KPI.md)
4. **Prontidão Híbrida:** O portal está apto a gerenciar simultaneamente tabelas dos dois regimes até dezembro de 2028, com teste de cenário completo executado e aprovado pelo Comitê Fiscal. [INDEX] → [01-PROJECT-CHARTER.md](./01-PROJECT-CHARTER.md) — Seção 8

------------------------------
## 5. Matriz de Rastreabilidade dos Épicos

| Épico | Entrega | Requisitos Vinculados | Features |
|:---|:---|:---|:---|
| 01 — Motor de Cadastro Fiscal | Entrega 1 | BR-01, BR-02, BR-03 | 01.1, 01.2, 01.3 (3 features) |
| 02 — Controle de Acesso e Rastreabilidade | Entrega 2 | BR-04, BR-05 | 02.1, 02.2 (2 features) |
| 03 — Operações Fiscais em Escala | Entrega 3 | BR-06, BR-07 | 03.1, 03.2 (2 features) |
| 04 — Inteligência Fiscal e Analytics | Entrega 4 | BR-08, BR-09, BR-10 | 04.1, 04.2, 04.3 (3 features) |

------------------------------
## 6. Mapa de Entregas × Features

```text
ENTREGA 1                ENTREGA 2                ENTREGA 3                ENTREGA 4
Portal: Gestão Básica    Governança e Auditoria   Operações em Escala      Portal Completo
─────────────────────    ─────────────────────    ────────────────────     ────────────────────

 01.1 Painel (M1)         02.1 Acessos (M6)        03.1 Fluxos Aprovação    04.1 Relatórios
 01.2 Cadastro (M2)       02.2 Trilha (M4)         03.2 Import/Export (M5)   04.2 Dashboards KPIs
 01.3 Classif. (M3)                                                          04.3 Período Híbrido

 3 features               2 features               2 features                3 features
```

------------------------------

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: 012-agile-epic, agile-ba-practices.*
