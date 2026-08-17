# PROMPT-GENERATE-450-TEAM-SKILLS-MAP

## Contexto

Este prompt gera o artefato `450-TEAM-SKILLS-MAP.md` — a **matriz de skills** do time técnico do projeto. Diferente do `460-TEAM-CAPACITY.md` (que responde QUEM está disponível e quantas horas/dia), este documento foca em O QUE cada perfil sabe fazer e seu nível de proficiência técnica.

**Discovery Team (Upstream Architecture):** Este documento é o ponto de partida do Bloco A (People & Solutions). Ele define os papéis do time de **Discovery Técnico** — a equipe de arquitetura upstream responsável por analisar o escopo do projeto, definir a arquitetura, stacks, segurança e especificações técnicas antes do início do desenvolvimento. Os integrantes deste time executarão as fases subsequentes do Bloco A e B.

O artefato serve como:
- **Referência de alocação:** Tech Lead consulta a matriz para designar a pessoa certa para cada tarefa
- **Detector de gaps:** Identifica competências necessárias que o time não possui
- **Plano de capacitação:** Base para definir treinamentos e contratações

**Inputs upstream (Bloco 0 → Bloco A):**
- **PRODUCT-BACKLOG-LIST.md (F3):** Backlog priorizado para entender o escopo do projeto
- **PRD-DEFINITION.md (F4):** PRD de Negócio (congelado) — visão do produto, MVP e requisitos de negócio
- **Documentos de Negócio:** Charter, BRD, Epics, Features para contexto adicional

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{TECHNICAL_SOLUTION_PATH}` | Caminho base das soluções técnicas |
| `{TECHNICAL_SOLUTION_NAMES}` | Lista de nomes das soluções técnicas do projeto |
| `{ARCHITECTURE_GLOBAL}` | Caminho para a pasta de arquitetura global (ADRs, blueprints) |
| `{SECURITY_GLOBAL}` | Caminho para o documento de segurança global (GLOBAL-SECURITY.md) |
| `{PROJECT_DOCUMENTS_INPUTS}` | (Opcional) Lista de caminhos para documentos brutos de entrada adicionais |
| `{PROJECT_PROMPT_INPUTS}` | (Opcional) Lista de caminhos para prompts auxiliares ou contextos adicionais |

---

## Fluxo de Execução

### Passo 0 — Validação de Parâmetros
Verificar se TODOS os parâmetros foram informados.

### Passo 1 — Questionar o Humano sobre Integrantes do Discovery Team

Antes de gerar o arquivo, apresentar os papéis fixos do **Discovery Team** e perguntar ao usuário:

> **👥 Discovery Team — Upstream Architecture**
>
> Os seguintes papéis compõem o time de Discovery Técnico que executará a análise de arquitetura upstream:
>
> | # | Papel | Responsabilidade no Discovery |
> |---|-------|-------------------------------|
> | 1 | **Engenheiro de Sistemas** | Visão sistêmica, integração entre componentes, requisitos não-funcionais |
> | 2 | **Arquiteto de Soluções** | Definição da arquitetura macro, C4, ADRs, padrões cross-solution |
> | 3 | **Arquiteto de Banco de Dados** | Modelagem de dados, multi-tenant, schema design, migrações |
> | 4 | **Arquiteto de DevOps/SRE** | Infraestrutura, CI/CD, containerização, observabilidade, deploy |
> | 5 | **Arquiteto de Segurança** | Threat model, IAM, secrets management, compliance (LGPD, PCI, SOC2) |
> | 6 | **Arquiteto/Especialista de Testes** | Estratégia de testes, tipos de teste, ferramentas, qualidade |
> | 7 | **Arquiteto/Especialista de Cloud-Provider** | Serviços cloud, networking, custos, multi-cloud |
> | 8 | **Líder Técnico / Tech Lead** | Coordenação técnica, code review, mentoria, decisões de design |
> | 9 | **Especialista em Integrações/APIs** | Contratos de API, mensageria, integração entre soluções |
> | — | **Principal Architecture (transversal)** | Supervisão de arquitetura cross-team, padrões corporativos, governança técnica — atua em múltiplos projetos simultaneamente |
>
> Deseja informar os nomes e dados dos integrantes para esses papéis?
> - **SIM** → Forneça para cada papel: Nome, Contato (e-mail), Nível (★☆☆ a ★★★), e tecnologias que domina
> - **NÃO** → O arquivo será criado com os papéis e a matriz de skills vazia. Os nomes e contatos serão preenchidos posteriormente.

### Passo 2 — Carregar Documentos Base
Ler documentos de negócio (Charter, BRD, Epics, Features) + PRODUCT-BACKLOG-LIST.md (F3) + PRD-DEFINITION.md (F4) para entender as demandas técnicas, o escopo do projeto e as entregas de negócio definidas.

### Passo 3 — Invocar Skills Especializadas
Invocar skills para mapear competências necessárias vs. disponíveis, gerar matriz de skills com níveis de proficiência e identificar gaps.

### Passo 4 — Gerar o Artefato
Gerar `{TECHNICAL_DEFINITIONS_PATH}/450-TEAM-SKILLS-MAP.md` seguindo o modelo da seção "Modelo do Arquivo" abaixo, preenchendo:
- Contexto técnico do projeto (stack, escopo)
- Matriz de skills: Papel × Tecnologia × Nível (★☆☆ a ★★★) para cada integrante informado
- Skills por categoria (Linguagens, Frameworks, Bancos, Cloud, DevOps, Segurança, Frontend, Mobile, Integrações)
- Gap analysis: skills necessárias vs. disponíveis
- Recomendações de contratação/capacitação
- Referência ao TEAM-CAPACITY.md para nomes e contatos

Se o humano NÃO forneceu dados dos integrantes no Passo 1, gerar o documento com os 9 papéis listados, a estrutura completa de seções, mas com a matriz de skills vazia (colunas de profissionais sem nomes, aguardando preenchimento).

### Passo 5 — Validação Pós-Geração
Verificar: arquivo no caminho correto, papéis do Discovery Team listados (incluindo Principal Architecture como transversal), matriz de skills estruturada, gaps documentados, estrutura de seções completa.

---

## Modelo do Arquivo

O arquivo `450-TEAM-SKILLS-MAP.md` deve seguir a estrutura abaixo:

```markdown
# 450-TEAM-SKILLS-MAP — Matriz de Skills do Time Técnico

- **Projeto:** {PROJECT_ID_NAME}
- **Versão:** 1.0
- **Data de Criação:** [DATA ATUAL]
- **Última Atualização:** [DATA ATUAL]
- **Status:** Em Revisão / Aguardando Validação
- **Documento Complementar:** [460-TEAM-CAPACITY.md](./460-TEAM-CAPACITY.md) (nomes, contatos, horas/semana)

---

## 1. Objetivo

Este documento apresenta a **matriz de competências técnicas** do time executor do projeto — o **Discovery Team** responsável pela fase de Upstream Architecture. Enquanto o [TEAM-CAPACITY](./460-TEAM-CAPACITY.md) responde **QUEM** está disponível e **quantas horas/dia**, este documento foca em **O QUE** cada perfil sabe fazer e seu **nível de proficiência**.

Ele serve como:
- **Referência de alocação:** Tech Lead consulta a matriz para designar a pessoa certa para cada tarefa
- **Detector de gaps:** Identifica competências necessárias que o time não possui
- **Plano de capacitação:** Base para definir treinamentos e contratações

---

## 2. Discovery Team — Papéis da Upstream Architecture

Os papéis abaixo compõem o time fixo de Discovery Técnico. Cada papel representa uma competência essencial para a fase de análise e definição da arquitetura do projeto.

| # | Papel | Responsabilidade no Discovery | Nome | Contato | Nível |
|---|-------|-------------------------------|------|---------|-------|
| 1 | Engenheiro de Sistemas | Visão sistêmica, integração entre componentes, requisitos não-funcionais | (a designar) | (a designar) | — |
| 2 | Arquiteto de Soluções | Definição da arquitetura macro, C4, ADRs, padrões cross-solution | (a designar) | (a designar) | — |
| 3 | Arquiteto de Banco de Dados | Modelagem de dados, multi-tenant, schema design, migrações | (a designar) | (a designar) | — |
| 4 | Arquiteto de DevOps/SRE | Infraestrutura, CI/CD, containerização, observabilidade, deploy | (a designar) | (a designar) | — |
| 5 | Arquiteto de Segurança | Threat model, IAM, secrets management, compliance (LGPD, PCI, SOC2) | (a designar) | (a designar) | — |
| 6 | Arquiteto/Especialista de Testes | Estratégia de testes, tipos de teste, ferramentas, qualidade | (a designar) | (a designar) | — |
| 7 | Arquiteto/Especialista de Cloud-Provider | Serviços cloud, networking, custos, multi-cloud | (a designar) | (a designar) | — |
| 8 | Líder Técnico / Tech Lead | Coordenação técnica, code review, mentoria, decisões de design | (a designar) | (a designar) | — |
| 9 | Especialista em Integrações/APIs | Contratos de API, mensageria, integração entre soluções | (a designar) | (a designar) | — |
| — | **Principal Architecture (transversal)** | Supervisão de arquitetura cross-team, padrões corporativos, governança técnica | (a designar) | (a designar) | — |

> ⚠️ **Status:** [Se nomes não foram fornecidos: "Aguardando definição dos integrantes. Os campos Nome, Contato e Nível devem ser preenchidos tão logo os profissionais sejam designados."]
> **Papéis Transversais:** Principal Architecture e Tech Lead são papéis transversais que podem atuar em múltiplos projetos simultaneamente. Sua capacidade deve ser rateada conforme necessário.

---

## 3. Contexto Técnico do Projeto

### 3.1 Stack Tecnológica Definida

| Camada | Tecnologia | Relevância |
|:---|:---|:---|
| [preencher conforme documentos de negócio e TECHNICAL-PLAN.md] |

### 3.2 Escopo Técnico

| Épico | Funcionalidades | Complexidade Técnica |
|:---|:---|:---|
| [preencher conforme documentos de negócio] |

---

## 4. Matriz de Competências

### 4.1 Matriz Papel × Tecnologia × Proficiência

**Legenda de Proficiência:**
- ★★★ Senior/Especialista — Autônomo, define padrões, mentoriza
- ★★☆ Pleno/Intermediário — Produtivo, segue padrões, resolve problemas típicos
- ★☆☆ Junior/Básico — Assistido, executa tarefas bem definidas
- `—` Sem exposição conhecida

| Tecnologia / Domínio | [Papel 1] | [Papel 2] | ... | [Papel 9] |
|:---|:---:|:---:|:---:|:---:|
| **Linguagens** | | | | |
| [preencher conforme stack do projeto] | | | | |
| **Frameworks Backend** | | | | |
| [preencher conforme stack] | | | | |
| **Frameworks Frontend** | | | | |
| [preencher conforme stack] | | | | |
| **Banco de Dados** | | | | |
| [preencher conforme stack] | | | | |
| **IAM & Segurança** | | | | |
| [preencher conforme stack] | | | | |
| **DevOps & Infra** | | | | |
| [preencher conforme stack] | | | | |
| **Cloud & Hosting** | | | | |
| [preencher conforme stack] | | | | |
| **Qualidade & Testes** | | | | |
| [preencher conforme stack] | | | | |
| **Observabilidade** | | | | |
| [preencher conforme stack] | | | | |
| **Integrações & APIs** | | | | |
| [preencher conforme stack] | | | | |

### 4.2 Resumo por Papel

| Papel | Profissional | Tecnologias Core (★★★) | Carga |
|:---|:---|:---|:---:|
| [preencher para cada papel] | [Nome ou "a designar"] | [Lista] | [h/d] |

---

## 5. Skills por Categoria

[Para cada categoria de tecnologia, criar tabela com: Skill, Cobertura, Profissionais ★★★, Gaps]

---

## 6. Gap Analysis: Competências Necessárias vs. Disponíveis

### 6.1 Gaps Críticos (🔴 — Risco Alto)

| Gap | Severidade | Impacto | Mitigação |
|:---|:---:|:---|:---|
| [preencher conforme análise] |

### 6.2 Gaps Moderados (🟡 — Risco Médio)

| Gap | Severidade | Impacto | Mitigação |
|:---|:---:|:---|:---|

### 6.3 Gaps Menores (🟢 — Risco Baixo)

| Gap | Severidade | Impacto | Mitigação |
|:---|:---:|:---|:---|

---

## 7. Recomendações

### 7.1 Contratação / Realocação

| # | Recomendação | Prazo | Prioridade |
|:---|:---|:---|:---:|

### 7.2 Capacitação

| # | Treinamento | Público | Quando | Duração |
|:---|:---|:---|:---|:---:|

### 7.3 Ações Imediatas (Sprint 0)

[Lista de ações por papel]

---

## 8. Referências

| Documento | Relação |
|:---|:---|
| [460-TEAM-CAPACITY.md](./460-TEAM-CAPACITY.md) | Nomes, contatos, horas/semana de cada profissional |
| [PRODUCT-BACKLOG-LIST.md](../../business/project-documents/PRODUCT-BACKLOG-LIST.md) | Backlog priorizado do produto (F3 — Bloco 0) |
| [PRD-DEFINITION.md](./440-PRD-DEFINITION.md) | PRD de Negócio congelado (F4 — Bloco 0) |
| [Documentos de Negócio] | Charter, BRD, Epics, Features |

---

## Histórico de Alterações

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | [DATA ATUAL] | Criação inicial: matriz de skills do Discovery Team (9 papéis de Upstream Architecture) | Time de Arquitetura |
```

---

## Skills Utilizados

> **📌 Nota sobre Skills:** Skills recomendados para esta fase. O agente tem autonomia para selecionar outros mais aderentes.

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `team-composition-analysis` | Analisar composição do Discovery Team e identificar gaps | People |
| 2 | `gap-analysis` | Análise de gaps de competências vs. necessidades | Análise |
| 3 | `skill-audit` | Auditoria de skills existentes no time | Discovery |
| 4 | `engineering-skills` | Validar skills de engenharia necessárias | Engenharia |
| 5 | `senior-architect` | Validar skills de arquitetura necessárias | Arquitetura |
| 6 | `senior-pm` | Validar skills de gestão necessárias | PM |
| 7 | `documentation-writer` | Redigir o TEAM-SKILLS-MAP.md consolidado | Documentação |

> **🔄 Flexibilidade:** Substituir skills conforme aderência e justificar no changelog do artefato.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: prompt gerador da matriz de skills do time técnico | Time de Arquitetura |
| 2.0 | 28/07/2026 | Refatoração: renomeado para TEAM-SKILLS-MAP; adicionado Discovery Team com 9 papéis fixos de Upstream Architecture; incluído modelo completo do arquivo; questionamento ao humano sobre dados dos integrantes | Time de Arquitetura |
| 3.0 | 30/07/2026 | Atualização Bloco A (F5): adicionado Principal Architecture como papel transversal; adicionados inputs PRODUCT-BACKLOG-LIST (F3) e PRD-DEFINITION (F4) do Bloco 0 | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude. Skills de referência listados acima. Outros skills podem ser utilizados conforme aderência.*
