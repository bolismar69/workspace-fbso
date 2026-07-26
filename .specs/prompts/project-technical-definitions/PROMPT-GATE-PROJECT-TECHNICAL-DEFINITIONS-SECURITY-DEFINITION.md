# PROMPT-GATE-PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION

## Contexto

Este prompt implementa o **Gate de Validação da Definição de Segurança** para o artefato `PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md`. Verifica se as regras de segurança do projeto estão alinhadas com o GLOBAL-SECURITY.md e cobrem todas as soluções.

**Princípio fundamental:** As 3 Regras de Ouro do GLOBAL-SECURITY.md são inegociáveis. Violação de qualquer uma = REPROVAÇÃO automática.

---

## Parâmetros de Entrada

| Parâmetro | Descrição |
|---|---|
| `{PROJECT_PATH}` | Caminho base dos projetos de negócio |
| `{PROJECT_ID_NAME}` | Identificador completo do projeto |
| `{TECHNICAL_DEFINITIONS_PATH}` | Caminho da pasta technical-definitions |
| `{SECURITY_GLOBAL}` | Caminho do GLOBAL-SECURITY.md |

---

## Fluxo de Execução

### Passo 1 — Carregar Documentos Base
Ler `PROJECT-TECHNICAL-DEFINITIONS-SECURITY-DEFINITION.md`, GLOBAL-SECURITY.md, Architecture Definition, Catálogo de Soluções.

### Passo 2 — Executar Dimensões de Validação

#### Dimensão 1: Conformidade com GLOBAL-SECURITY.md
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

### Passo 3 — Calcular Veredito
Regra especial: violação de qualquer Regra de Ouro = REPROVADO automático.

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

---

🤖 *Documentação gerada de forma automatizada pelo Agente: Arquiteto de Soluções/Claude.*
