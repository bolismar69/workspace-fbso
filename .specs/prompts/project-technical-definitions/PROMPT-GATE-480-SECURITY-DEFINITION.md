# PROMPT-GATE-480-SECURITY-DEFINITION

## Contexto

Este prompt implementa o **Gate de Validação da Definição de Segurança** para o artefato `480-SECURITY-DEFINITION.md`. Verifica se as regras de segurança do projeto estão alinhadas com o GLOBAL-SECURITY.md e cobrem todas as soluções.

**Princípio fundamental:** As 3 Regras de Ouro do GLOBAL-SECURITY.md são inegociáveis. Qualquer violação resulta em NÃO COMPLIANCE automático.

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

**Arquivos gerados pelo GENERATE:** `480-SECURITY-DEFINITION.md`

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `480-SECURITY-DEFINITION.md`, GLOBAL-SECURITY.md, Architecture Definition, Catálogo de Soluções.

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Conformidade com GLOBAL-SECURITY.md (INEGOCIÁVEL)
| # | Verificação | Critério |
|---|---|---|
| 1.1 | Regra de Ouro #1: Menor Privilégio | Controles de acesso cross-solution documentados |
| 1.2 | Regra de Ouro #2: Zero Hardcoded Secrets | Estratégia de secrets management definida |
| 1.3 | Regra de Ouro #3: Não Confiar no Input | Validação cross-solution documentada |
| 1.4 | Checklist SDD | 4 áreas cobertas para o projeto |

#### Dimensão 2: Cobertura de Segurança
| # | Verificação | Critério |
|---|---|---|
| 2.1 | Threat model macro | Atores, superfícies e riscos documentados |
| 2.2 | IAM cross-solution | Estratégia de autenticação/autorização entre soluções |
| 2.3 | Criptografia em trânsito | TLS/mTLS entre soluções definido |

#### Dimensão 3: Completude Operacional
| # | Verificação | Critério |
|---|---|---|
| 3.1 | Pipeline DevSecOps | SAST, SCA, Secret Scanning configurados |
| 3.2 | Matriz de compliance | LGPD, PCI, SOC2 mapeados por solução |
| 3.3 | Controles acionáveis | Cada controle: o quê, quem, como verificar |

### Passo 3 — Emitir Veredito

**Regra especial:** Violação de qualquer Regra de Ouro (Dimensão 1) = NÃO COMPLIANCE automático, independentemente das outras dimensões.

---

## FORMATO OBRIGATÓRIO DE SAÍDA (O RELATÓRIO DO GATE)

### 🚨 CENÁRIO A: SE FOREM ENCONTRADOS DESVIOS OU VIOLAÇÕES DE REGRAS DE OURO (NÃO COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE SEGURANÇA: [Nome do Projeto]

##### 🔍 Pontos Conflitantes Identificados:
- **[ID-CONFLITO-SEC-01] - [Título Curto]:**
  - **O que foi gerado:** [Descrever o trecho problemático]
  - **O que o GLOBAL-SECURITY.md/Architecture determinava:** [Descrever a referência]
  - **Impacto:** [O risco de segurança ou não-conformidade]
  - **Sugestão de tratativa:** [O que poderia ser feito para corrigir]

> ⚠️ **Atenção especial para violações de Regras de Ouro (Dimensão 1):** Estas são inegociáveis e devem ser corrigidas obrigatoriamente antes do COMPLIANCE.

##### ❓ Perguntas de Alinhamento para o Usuário:
Para que possamos corrigir a definição de segurança, por favor, responda:
1. Quanto ao **[ID-CONFLITO-SEC-01]**, qual é a definição correta a ser aplicada?
2. [Perguntas diretas para sanar os desvios encontrados]

---
### 🛑 STATUS DO GATE: [NÃO COMPLIANCE]
*(Instrução: O processo pausa aqui. Assim que o humano responder, injete este relatório + respostas no PROMPT-FIX-480-SECURITY-DEFINITION.md)*

---

### ✅ CENÁRIO B: SE A SEGURANÇA ESTIVER 100% CONFORME (PRÉ-COMPLIANCE)

#### 📊 RELATÓRIO DE AUDITORIA DE SEGURANÇA: [Nome do Projeto]

### 🛑 STATUS DO GATE: [PRÉ-COMPLIANCE INTERNO - AGUARDANDO VALIDAÇÃO HUMANA]

- **DOCUMENTO:** `480-SECURITY-DEFINITION.md` gerado conforme GLOBAL-SECURITY.md e arquitetura.
- **AUDITORIA DA IA:** Conformidade verificada. 3 Regras de Ouro em conformidade. Threat model macro documentado. IAM cross-solution definido. Pipeline DevSecOps configurado. Matriz de compliance preenchida.
- **DIRETRIZ:** Peço que leia a definição de segurança para verificar se as regras e controles atendem às suas expectativas.

Por favor, responda às seguintes perguntas para podermos prosseguir ou reajustar:

1. A definição de segurança está em compliance com o GLOBAL-SECURITY.md e atende às necessidades do projeto?
2. Deseja enviar mais documentos/arquivos para enriquecer a definição de segurança?
3. Deseja enviar mais informações ou novos direcionamentos via input de texto neste momento?

*(Instrução de Orquestração: Se "Sim, Não, Não" → [STATUS: COMPLIANCE] e Fase 9 (DATA-ARCHITECTURE-DEFINITION). Se novos inputs → retrocede ao PROMPT-GENERATE).*

---

## Skills Utilizados

| Ordem | Skill | Propósito | Categoria |
|---|---|---|---|
| 1 | `security-auditor` | Auditar conformidade com GLOBAL-SECURITY.md | Auditoria |
| 2 | `threat-modeling-expert` | Validar threat model | Threat Model |
| 3 | `senior-security` | Validação sênior | Estratégia |
| 4 | `security-reviewer` | Revisão de cobertura | Revisão |
| 5 | `gap-analysis` | Identificar gaps de segurança | Análise |

> **🔄 Flexibilidade:** Substituir skills conforme aderência.

---

## Registro de Alterações do Documento

| Versão | Data | Alteração | Autor |
|:---|:---|:---|:---|
| 1.0 | 25/07/2026 | Criação inicial: gate de validação da definição de segurança | Time de Arquitetura |
| 2.0 | 28/07/2026 | Refatoração: adoção do padrão HITL com 3 perguntas obrigatórias e veredito binário | Time de Arquitetura |
| 3.0 | 30/07/2026 | Atualização F4→F8: orquestração redirecionada para Fase 9 (DATA-ARCHITECTURE-DEFINITION) | Time de Arquitetura |

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
