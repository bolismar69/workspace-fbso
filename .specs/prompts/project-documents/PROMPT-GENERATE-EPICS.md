## Arquivo Atualizado: PROMPT-GENERATE-EPICS.md

# PROMPT: GERADOR DE ÉPICOS ÁGEIS DE ALTA DENSIDADE (EPICS SPECIFICATION)
## Arquivo: PROMPT-GENERATE-EPICS.md
## Versão: 4.0 — Estrutura Modular com Arquivos Individuais por Épico

Atue como um Product Manager Sênior e Especialista em Engenharia de Backlog Ágil, aplicando as melhores práticas de `breakdown-epic-pm`, `agile-ba-practices`, `program-manager`, `senior-pm`, `delivery-manager`, `agile-coach`, `scrum-master` e os filtros pragmáticos corporativos de `caveman` e `caveman-review`. Sua missão é ler os documentos de negócio congelados (Project Charter e BRD) e gerar uma **Especificação de Épicos (Epics Specification)** exaustiva e corporativa, estruturada de forma modular: um arquivo índice central + arquivos individuais por épico na pasta `epics/`.

### 🛑 DIRETRIZES CRÍTICAS DE CONTEXTO E ANATOMIA:
1. **Estrutura Modular (NOVO — v4.0):** O documento de Épicos passa a ser entregue em dois artefatos complementares:
   - **Arquivo Índice:** `03-EPICS-{PROJECT_ID_NAME}.md` — resumo executivo com visão consolidada, cronograma, mapa de dependências e matriz de rastreabilidade BRD×Épico×Jornada.
   - **Arquivos Individuais:** `epics/EP-NNNN-{nome-slugificado}.md` — um arquivo por épico com as 8 seções de detalhamento completo e matriz BRD×Épico×Jornada específica.
2. **Rastreabilidade e Herança Absoluta:** Mantenha a nomenclatura de IDs sincronizada. Os Épicos devem utilizar a taxonomia `EP-0001`, `EP-0002`, etc. (4 dígitos, zero-padded), mapeando explicitamente quais IDs de entregas do Charter (`D1`, `D2`) e requisitos do BRD (`BR-XX`) eles cobrem.
3. **Estruturas Visuais Exigidas (no Índice):** O arquivo índice deve conter obrigatoriamente uma Visão Geral dos Épicos com links para os arquivos individuais, datas-alvo que reflitam a janela total do projeto, um Cronograma detalhado de amarração com os Marcos (`M1, M2...`) e um Mapa de Linha de Tempo em formato de diagrama ASCII (baseado nos marcos cronológicos do Charter).
4. **Detalhamento Cirúrgico por Épico (nos Arquivos Individuais):** Cada Épico deve ser quebrado em 8 subseções corporativas: 1. Nome; 2. Objetivo detalhado (Problema, Solução, Impacto); 3. Tabela de Personas; 4. Jornadas de Usuário de Alto Nível com referências ao BRD; 5. Requisitos de Negócio (Funcionais e Não-Funcionais); 6. Métricas de Sucesso (KPI vs Meta); 7. Fora do Escopo específico do Épico; 8. Matriz de Avaliação de Valor de Negócio.
5. **Blindagem Contra Termos Técnicos de TI:** Descreva o backlog sob a perspectiva de uso e valor comercial. Não cite códigos, bancos de dados específicos, infraestrutura de CI/CD ou linguagens de programação. 
6. **Análise de Conformidade Interna (Caveman Review):** Aplique um filtro mental pragmático durante a geração. Elimine qualquer termo rebuscado de tecnologia que mascare a falta de uma regra de negócio real.
7. **Links Cruzados Obrigatórios:** 
   - O índice deve linkar para cada arquivo individual: `[EP-0001](epics/EP-0001-nome-do-epico.md)`
   - Cada arquivo individual deve linkar de volta para o índice: `[Índice de Épicos](../03-EPICS-{PROJECT_ID_NAME}.md)`
   - Cada arquivo individual deve referenciar os requisitos BRD: `[BR-XX](../02-BRD-{PROJECT_ID_NAME}.md)`
   - Cada jornada deve referenciar o requisito BRD que atende: `🏷️ Atende [BR-XX](../02-BRD-{PROJECT_ID_NAME}.md)`

---
### INSTRUÇÕES DE EXECUÇÃO:
1. **Inputs de Entrada:** Consuma o arquivo `02-BRD-{PROJECT_ID_NAME}.md` (Fase 2) e o `01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` (Fase 1) fornecidos pelo usuário.
2. **Criar Estrutura de Diretórios:** Execute `mkdir -p {PROJECT_COMPLETE_PATH_NAME}/epics/` para garantir que a pasta de artefatos modulares exista.
3. **Processamento:** Agrupe os requisitos correlacionados em jornadas de valor lógicas e sequenciais cronológicas de acordo com o escopo do projeto recebido. Atribua códigos sequenciais `EP-0001`, `EP-0002`... `EP-NNNN` a cada épico identificado.
4. **Geração dos Arquivos Individuais:** Para cada épico identificado, crie um arquivo na pasta `epics/` com o nome `EP-NNNN-{nome-slugificado}.md` seguindo o template de 8 seções + matriz de rastreabilidade específica (detalhado abaixo).
5. **Geração do Arquivo Índice:** Crie o arquivo `03-EPICS-{PROJECT_ID_NAME}.md` na raiz do projeto com o conteúdo consolidado (template detalhado abaixo).
6. **Formatação:** Retorne os documentos finais estruturados exatamente com os layouts de saída detalhados abaixo, preenchendo as variáveis com os dados do projeto do usuário.

---
### ESTRUTURA E LAYOUT DO ARQUIVO ÍNDICE (`03-EPICS-{PROJECT_ID_NAME}.md`):

```
# Épicos do Projeto: [Inserir Nome do Projeto]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documento** | EPICS-{PROJECT_ID_NAME} |
| **Versão** | 1.0 — Documento Inicial de Épicos (Estrutura Modular) |
| **Data** | [Data Atual] |
| **Origem** | `02-BRD-{PROJECT_ID_NAME}.md` v[X.X] e `01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` v[X.X] |
| **Status** | Em Revisão / Aguardando Validação |

---

## Visão Geral dos Épicos

[Construir tabela cruzando os Épicos identificados: ID com link para o arquivo individual, Épico, Objetivo de Negócio, Quantidade de Funcionalidades Estimadas, Prioridade MoSCoW, Data-Alvo/Prazo Estimado].

| ID | Épico | Objetivo de Negócio | Func. | Prioridade | Data-Alvo |
|----|-------|---------------------|-------|------------|-----------|
| **EP-0001** | [Nome do Épico 1] | [Objetivo resumido] | N | Must Have | **DD/MM/AAAA** |
| **EP-0002** | [Nome do Épico 2] | [Objetivo resumido] | N | Must Have | **DD/MM/AAAA** |

> 📄 **Detalhamento completo** de cada épico disponível na pasta [`epics/`](epics/):
> - [EP-0001 — Nome do Épico 1](epics/EP-0001-nome-do-epico-1.md)
> - [EP-0002 — Nome do Épico 2](epics/EP-0002-nome-do-epico-2.md)

### Cronograma de Épicos

[Construir tabela vinculando a linha do tempo do projeto: Data-Alvo, Marco (M1, M2...), Épicos e códigos de Funcionalidades relacionados].

### Mapa de Dependências entre Épicos

[Desenhar um mapa de linha de tempo em bloco ASCII customizado para os prazos do usuário, demarcando os Marcos (M1 a M7) e blocos horizontais indicando a evolução sequencial e dependências dos Épicos (EP-0001, EP-0002...)].

> **Nota:** [Descrever dependências sequenciais entre épicos conforme milestones do Project Charter].

---

## Matriz de Rastreabilidade BRD → Épicos

Todo requisito funcional do BRD deve estar coberto por pelo menos um épico. A matriz abaixo audita essa cobertura e vincula cada BR às jornadas de usuário que o realizam.

| BRD | Requisito Funcional | Épico | Jornada(s) que Realizam |
|:---|:---|:---|:---|
| **BR-01** | [Descrição do requisito] | [**EP-0001** — Nome do Épico](epics/EP-0001-nome-do-epico-1.md) | J1: [Nome da jornada] · J2: [Nome da jornada] |
| **BR-02** | [Descrição do requisito] | [**EP-0002** — Nome do Épico](epics/EP-0002-nome-do-epico-2.md) | J1: [Nome da jornada] |

### Auditoria de Cobertura

| Métrica | Resultado |
|---|---|
| BRs cobertos por épicos | N/N (100%) |
| BRs cobertos por jornadas | N/N |
| Épicos com rastreabilidade BRD | N/N (100%) |
| Épicos com BRs novos (não mapeados) | 0 |
| Requisitos órfãos (sem épico) | 0 |

> 💡 [Notas sobre cobertura, ex: BRs sem jornada narrativa explícita mas cobertos via requisitos funcionais].

---

## Sumário de Cobertura do Escopo

[Construir tabela cruzando explicitamente a "Entrega do Project Charter (D1, D2...)" contra os "Épico(s) que cobrem (EP-0001, EP-0002...)" para fins de auditoria mecânica de rastreabilidade vertical].

| Entrega do Project Charter | Épico(s) que cobrem |
|---------------------------|---------------------|
| D1 — [Nome da Entrega] | EP-0001 |
| D2 — [Nome da Entrega] | EP-0002 |

---

> **Este documento é um índice resumido.** O detalhamento completo de cada épico — incluindo objetivo, personas, jornadas, requisitos de negócio, métricas de sucesso, escopo excluído e valor de negócio — está nos arquivos individuais da pasta [`epics/`](epics/).

---
🤖 *Documentação gerada de forma automatizada pelo Agente: Analista de Negócios/Claude. Foram utilizados os skills: breakdown-epic-pm, agile-ba-practices. Estrutura modular v4.0.*
```

---
### ESTRUTURA E LAYOUT DE CADA ARQUIVO INDIVIDUAL (`epics/EP-NNNN-{nome-slugificado}.md`):

> **Nota:** O slug do nome deve ser derivado do título do épico, em kebab-case, preservando os caracteres ASCII. Ex: "Portal Administrativo Interno" → `portal-administrativo-interno`.

```
# EP-NNNN: [Inserir Nome do Épico]

| Campo | Detalhe |
|-------|---------|
| **Épico** | EP-NNNN — [Nome do Épico] |
| **Projeto** | {PROJECT_ID_NAME} |
| **Documento** | EPICS-{PROJECT_ID_NAME} |
| **Versão** | 1.0 — Documento Inicial de Épicos |
| **Data** | [Data Atual] |
| **Origem** | `02-BRD-{PROJECT_ID_NAME}.md` e `01-PROJECT-CHARTER-{PROJECT_ID_NAME}.md` |
| **Status** | Em Revisão / Aguardando Validação |

> 📄 **Índice de Épicos:** [`03-EPICS-{PROJECT_ID_NAME}.md`](../03-EPICS-{PROJECT_ID_NAME}.md) | **Próximo:** [EP-NNNN — Nome](../EP-NNNN-proximo-epico.md) | **Anterior:** [EP-NNNN — Nome](../EP-NNNN-epico-anterior.md)

---

## 1. Nome do Épico
**[Inserir Nome Longo e Descritivo do Épico]**

**Requisitos BRD Vinculados:** [BR-XX](../02-BRD-{PROJECT_ID_NAME}.md) — [Nome do Requisito], [BR-YY](../02-BRD-{PROJECT_ID_NAME}.md) — [Nome do Requisito]

## 2. Objetivo (Goal)
- **Problema:** [A dor operacional real que a organização/usuário sofre hoje na ausência desta capacidade específica do projeto].
- **Solução:** [Como este bloco de entrega resolve centralizadamente o problema descrito].
- **Impacto:** [Os ganhos de velocidade, governança, receita ou autonomia para a empresa].

## 3. Personas de Usuário (User Personas)
[Construir tabela com as personas mapeadas para o ecossistema do usuário: Persona, Descrição e Necessidades específicas de negócio dentro deste Épico].

## 4. Jornadas de Usuário de Alto Nível (High-Level User Journeys)

**Jornada 1: [Nome da Jornada]**
1. [Passo 1 da jornada]
2. [Passo 2 da jornada]
...
> 🏷️ Atende [BR-XX](../02-BRD-{PROJECT_ID_NAME}.md)

**Jornada 2: [Nome da Jornada]**
1. [Passo 1 da jornada]
2. [Passo 2 da jornada]
...
> 🏷️ Atende [BR-YY](../02-BRD-{PROJECT_ID_NAME}.md)

## 5. Requisitos de Negócio (Business Requirements)

### Requisitos Funcionais
- [Lista em tópicos detalhando as capacidades do sistema mapeadas para este épico, herdadas estritamente do BRD].

### Requisitos Não-Funcionais
- [Desempenho de negócio, restrições operacionais corporativas e limites de acesso comercial].

## 6. Métricas de Sucesso (Success Metrics)
[Construir tabela correlacionando os indicadores de sucesso deste bloco: KPI e Meta Clara de Negócio].

| KPI | Meta |
|-----|------|
| [Indicador 1] | [Meta 1] |
| [Indicador 2] | [Meta 2] |

## 7. Fora do Escopo (Out of Scope)
- [Listar explicitamente o que não será feito dentro deste bloco específico para blindar o time contra desvios de escopo].

## 8. Valor de Negócio (Business Value)
[Construir tabela de priorização de produto contendo: Critério, Avaliação (Baixo/Médio/Alto/Crítico) e Justificativa de negócio].

| Critério | Avaliação | Justificativa |
|----------|-----------|---------------|
| Valor de Negócio | **[Alto/Crítico/Médio]** | [Justificativa] |

---

## Matriz de Rastreabilidade BRD → Este Épico

| BRD | Requisito Funcional | Este Épico | Jornada(s) que Realizam |
|:---|:---|:---|:---|
| **BR-XX** | [Descrição do requisito] | **EP-NNNN** — [Nome do Épico] | J1: [Nome da jornada] · J2: [Nome da jornada] |

---

> 📄 **Índice de Épicos:** [`03-EPICS-{PROJECT_ID_NAME}.md`](../03-EPICS-{PROJECT_ID_NAME}.md)
```

---
### PÓS-GERAÇÃO:

Após criar todos os arquivos individuais e o índice, execute uma verificação de integridade:
1. Confirme que todos os links no índice apontam para arquivos existentes na pasta `epics/`
2. Confirme que todos os arquivos individuais têm link de volta para o índice
3. Confirme que cada arquivo individual referencia corretamente os requisitos BRD aplicáveis
4. Confirme que a matriz consolidada no índice cobre 100% dos BRs
5. Confirme que não há BRs órfãos e nenhum épico criou escopo extra não mapeado no BRD

Ao final de cada arquivo gerado (índice e individuais), insira a tag:
`[STATUS: SUCESSO - ENVIADO PARA RE-AUDITORIA DE ÉPICOS]`
