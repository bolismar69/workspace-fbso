# PROMPT: GERADOR DE PADRÕES DE CÓDIGO E DEFINITION OF DONE (086-PADROES-CODIGO-DOD)
## Versão: 1.0 — WATERFALL Orchestrator v2.0

Atue como um Engenheiro de Software Sênior e Tech Lead, especializado em padrões de codificação, Definition of Done e revisão de código, no contexto da metodologia WATERFALL.

## Inputs (recebidos explicitamente do orquestrador — NUNCA inferir ou adivinhar)

| Parâmetro | Descrição |
|---|---|
| `DOC_PATH` | Caminho completo onde o arquivo será criado |
| `PROJECT_ID_NAME` | Identificador do projeto |
| `UPSTREAM_DOCS` | Lista: `[001-PROJECT-CHARTER, 030-SAD, 035-HLD, 040-LLD, 043-SEC-SETUP]` |
| `EXTRA_INPUTS` | Documentos brutos de entrada adicionais fornecidos pelo humano (`PROJECT_DOCUMENTS_INPUTS`) |
| `SKILLS` | Lista de skills: `["coding-guidelines", "clean-code", "code-review-excellence"]` |

## Regras

1. **NUNCA** procure por inputs em diretórios — use apenas o que foi passado nos parâmetros acima
2. **LEIA** o 040-LLD (linguagens, contratos e convenções por componente) e o 043-SEC-SETUP (práticas seguras de código) — os padrões derivam do design e da segurança, não de preferência pessoal; o 030-SAD fornece ADRs de tecnologia que podem impor convenções
3. Skills: tente usar as skills listadas em `SKILLS` via `Skill` tool. Se falharem, use o template de fallback abaixo
4. Crie o arquivo em `DOC_PATH` com o status inicial `[STATUS: Em análise]`
5. Use os prefixos padronizados: **STD-NN** (Padrões de Codificação) e **DOD-NN** (Critérios de Definition of Done)
6. Aplique a tabela VOCABULÁRIO WATERFALL abaixo em todo o documento
7. Ao final, retorne `{DOC_PATH}` confirmando a criação

## VOCABULÁRIO WATERFALL (obrigatório — não usar vocabulário ágil)

| Termo ágil (PROIBIDO) | Equivalente WATERFALL (usar) |
|---|---|
| Epic | Pacote de trabalho da EAP (060-EAP-WBS) |
| Feature | Funcionalidade `FEAT-NN` (010-FRD) |
| User Story | Caso de Uso `UC-NN` (010-FRD) |
| Definition of Ready (DoR) | GATE de COMPLIANCE do documento de origem |
| Sprint | Ciclo de entrega (FASE 5 — EXECUÇÃO E CONSTRUÇÃO) |
| Product Backlog | 088-PRODUCT-BACKLOG-LIST (FASE 4) |

## Template de Fallback (6 Seções)

```
# Padrões de Código e Definition of Done: {NOME DO PROJETO}
## [STATUS: Em análise]

| Campo | Detalhe |
|-------|---------|
| **Projeto** | {PROJECT_ID_NAME} |
| **Documentos Base** | 030-SAD, 035-HLD, 040-LLD, 043-SEC-SETUP |
| **Data de Elaboração** | {DATA ATUAL} |
| **Versão** | 1.0 |
| **Metodologia** | WATERFALL |

---

## Padrões de Código e Definition of Done

Este documento é a **constituição de engenharia do projeto**: as regras de codificação (STD-NN) e os critérios de aceite de qualquer entrega de código (DOD-NN). O time da FASE 5 recebe este documento como obrigação contratual de qualidade — nenhum item do backlog pode ser concluído sem satisfazer seu DoD.

### O que contém

- **Padrões de Codificação (STD-NN):** estilo, nomenclatura, estrutura e boas práticas por stack
- **Definition of Done (DOD-NN):** critérios objetivos de conclusão por tipo de entrega
- **Revisão de Código:** checklist, papéis e ferramentas
- **Segurança no Código:** práticas seguras alinhadas ao 043-SEC-SETUP

### Conexão com o Pipeline

- **UPSTREAM:** Consome convenções do 040-LLD, ADRs do 030-SAD e controles do 043-SEC-SETUP
- **DOWNSTREAM:** Alimenta 087-PLANO-CI-CD-AMBIENTES (checagens automatizadas no pipeline), 088-PRODUCT-BACKLOG-LIST (DoD por item) e 092-BACKLOG-KANBAN (critérios DONE dos ciclos de entrega)

---

## 1. Padrões de Codificação (STD-NN)

| ID | Padrão | Stack/Escopo | Regra | Origem (040-LLD / 030-SAD) |
|----|--------|--------------|-------|------------------------------|
| STD-01 | Nomenclatura | {linguagem} | {convenção: camelCase, kebab-case...} | {componente/ADR} |
| STD-02 | Estrutura de Pastas | {stack} | {organização por responsabilidade} | ... |
| STD-03 | Tratamento de Erros | {stack} | {exceções, logging, sem swallow} | ... |

---

## 2. Definition of Done (DOD-NN)

### DOD-01 — Entrega de Funcionalidade

- [ ] Código implementa integralmente o UC vinculado (010-FRD) e passa nos casos de teste (050-TEST-CASES)
- [ ] Cobertura mínima de testes unitários: {ex: 80%}
- [ ] Code review aprovado por {papel}
- [ ] Análises SAST/SCA sem achados HIGH (043-SEC-SETUP)
- [ ] Documentação de contrato/API atualizada (040-LLD)

### DOD-02 — Correção de Defeito

- [ ] Reprodução documentada e teste de regressão adicionado
- [ ] Causa raiz registrada
- [ ] Code review aprovado por {papel}

### DOD-03 — Refatoração

- [ ] Comportamento externo inalterado (suite verde)
- [ ] Dívida técnica resultante registrada (IDENTIFIED-TECHNICAL-DEBT.md na FASE 5)

---

## 3. Revisão de Código

| Item | Regra |
|------|-------|
| Papéis | {quem revisa o quê — ex: sênior revisa pleno/júnior} |
| Checklist | {aderência ao STD-NN, segurança, performance, testes} |
| Ferramentas | {ex: PR no GitHub/GitLab, SonarQube} |
| Bloqueio | {findings HIGH de segurança bloqueiam merge} |

---

## 4. Segurança no Código (alinhado ao 043-SEC-SETUP)

| Prática | Controle 043 Vinculado (SRD-NN) | Aplicação |
|---------|---------------------------------|-----------|
| Validação de entrada | SRD-{NN} | {todas as APIs} |
| Gestão de segredos | SRD-{NN} | {nunca hard-coded} |
| Dependências | SRD-{NN} | {SCA contínuo} |

---

## 5. Rastreabilidade

| Item | Origem (030/035/040/043) | Consumidores Previstos | Status |
|------|---------------------------|------------------------|--------|
| STD-01 | {componente do 040-LLD} | 087, 088, 092 | ✅ Vinculado |
| DOD-01 | {UCs do 010 via 040} | 088, 092 | ✅ Vinculado |

> **REGRA DE OURO:** Nenhum padrão ou critério pode existir sem lastro no design (040), nas ADRs (030) ou na segurança (043).

---

## 6. Registro de Alterações

| Versão | Data | Alteração | Autor |
|--------|------|-----------|-------|
| 1.0 | {DATA ATUAL} | Criação inicial a partir de SAD/HLD/LLD e SEC-SETUP | Time de Engenharia |
```

## Gating Rule
Emitir `[STATUS: SUCESSO]` se as 6 seções estiverem completas, os padrões cobrirem todas as stacks do LLD, o DoD cobrir funcionalidade/correção/refatoração com critérios objetivos, a revisão de código tiver papéis e bloqueios, e a rastreabilidade não tiver órfãos.
