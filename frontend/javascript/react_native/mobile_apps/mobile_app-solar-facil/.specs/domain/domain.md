---
title: "Domínio — mobile_app-solar-facil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["domain", "glossary", "ubiquitous-language", "energia-solar"]
---

# Glossário de Domínio — Solar Fácil

## 1. Introdução

O Solar Fácil é um aplicativo mobile para gestão de associações no ecossistema de energia solar. O domínio modela o relacionamento entre **fornecedores** de energia solar, **beneficiados** (consumidores), e suas **movimentações mensais** de energia, viabilizando o modelo _Solar as a Service_.

O aplicativo é **offline-first**: os dados são armazenados localmente em SQLite, com planos de sincronização futura para backend remoto.

## 2. Termos de Domínio

### 2.1. Entidades Principais

| Termo | Definição | Sinônimos | Código (interface/type/enum) |
|---|---|---|---|
| **Associado** | Pessoa física ou jurídica registrada na plataforma Solar Fácil. Pode atuar como fornecedor, beneficiado, ou ambos (híbrido). | Membro, Usuário | `AssociadoType` — `src/types/AssociadoType.ts` |
| **Fornecedor** | Associado que gera e fornece energia solar excedente. Possui dados de potência instalada e disponibilidade. | Gerador, Produtor | `tipoAssociado: "Fornecedor"`, `FornecedorType` |
| **Beneficiado** | Associado que consome energia solar. Possui dados de concessionária, consumo médio e plano desejado. | Consumidor, Cliente | `tipoAssociado: "Beneficiado"`, `BeneficiadoType` |
| **Movimentação Mensal** | Registro mensal de energia recebida, valores cobrados e economia gerada para um associado. | Fatura, Extrato, Lançamento | `MovimentacaoMensalType` — `src/types/MovimentacaoMensalType.ts` |
| **Plano Comercial** | Oferta de plano de energia solar com condições específicas de economia e tarifa. | Plano, Oferta | `PlanType` — `src/types/PlanType.ts` |
| **Concessionária** | Empresa distribuidora de energia elétrica local (ex: Enel, CPFL, Light). | Distribuidora, Permissionária | `ConcessionariaType` — `src/types/ConcessionariaType.ts` |

### 2.2. Subdomínio: Cadastro & Autenticação

| Termo | Definição | Código |
|---|---|---|
| **Status do Associado** | Estado atual do cadastro: "Em cadastro", "Ativo", "Inativo", "Bloqueado", "Encerrado" | `AssociadoType.status` — CHECK constraint no SQLite |
| **Tipo de Pessoa** | Classificação fiscal: "Pessoa Física" (CPF) ou "Pessoa Jurídica" (CNPJ) | `AssociadoType.tipoPessoa` |
| **Login** | Autenticação local via CPF/CNPJ + senha. Estado mantido em `AuthContext` (memória). | `AuthContext.tsx`, `useQueryAssociadosSearchByCpfCnpjSenha.ts` |
| **Aceite de Termos** | Consentimento explícito do associado: "Sim" ou "Não" | `AssociadoType.aceitaTermos` |

### 2.3. Subdomínio: Energia Solar

| Termo | Definição | Código |
|---|---|---|
| **Consumo Médio** | Média mensal de consumo de energia elétrica (kWh) informada pelo beneficiado. Usada para dimensionar o plano. | `ConsumoMedioType`, `AssociadoType.consumoMedio` |
| **Potência Instalada** | Capacidade de geração do sistema fotovoltaico do fornecedor (kWp). | `AssociadoType.potenciaInstalada` |
| **Disponibilidade** | Indicação de quando o fornecedor tem energia excedente disponível. | `AssociadoType.disponibilidade` |
| **Tipo de Conexão** | Tipo de conexão do sistema de geração à rede elétrica (ex: monofásica, bifásica, trifásica). | `AssociadoType.tipoConexao` |
| **Energia Recebida (kWh)** | Quantidade de energia (em kWh) recebida pelo beneficiado em um mês. | `MovimentacaoMensalType.energiaRecebidaKwh` |
| **Tarifa Unitária (kWh)** | Valor unitário da tarifa de energia (R$/kWh) aplicada no mês. | `MovimentacaoMensalType.tarifaUnitariaKwh` |
| **Valor Economizado** | Diferença entre o valor que seria pago à concessionária e o valor cobrado pelo plano Solar Fácil. | `MovimentacaoMensalType.valorEconomizado` |
| **Percentual Economizado** | Economia percentual em relação ao valor original da concessionária. | `MovimentacaoMensalType.percentualEconomizado` |

### 2.4. Subdomínio: Financeiro

| Termo | Definição | Código |
|---|---|---|
| **Valor Total** | Valor total da movimentação mensal (R$). | `MovimentacaoMensalType.valorTotal` |
| **Valor Cobrado** | Valor efetivamente cobrado do associado no mês (R$). | `MovimentacaoMensalType.valorCobrado` |
| **Valor da Energia Recebida** | Valor monetário correspondente à energia recebida (R$). | `MovimentacaoMensalType.valorEnergiaRecebida` |
| **Status de Pagamento** | Situação do pagamento: "Pago" ou "Pendente". | `MovimentacaoMensalType.statusPagamento` |
| **Data de Vencimento** | Data limite para pagamento da movimentação. | `MovimentacaoMensalType.dataVencimento` |
| **Data de Pagamento** | Data em que o pagamento foi efetuado (null se pendente). | `MovimentacaoMensalType.dataPagamento` |

### 2.5. Subdomínio: Auxiliares

| Termo | Definição | Código |
|---|---|---|
| **FAQ** | Pergunta frequente com resposta. Conteúdo estático mockado. | `FAQType`, `mockFAQs.json` |
| **Estado (UF)** | Unidade federativa brasileira (27 UFs). | `constants/states.ts` |
| **Campo Dinâmico (Field Definition)** | Definição tipada de campo de formulário, usada para renderização dinâmica de inputs. | `FieldDefinitionType` |

## 3. Relações entre Conceitos

```
┌──────────────────────────────────────────────┐
│                  ASSOCIADO                   │
│  (Pessoa Física ou Jurídica)                │
│  tipoAssociado: Fornecedor | Beneficiado     │
│               | Hibrido                      │
└─────────────┬────────────────────────────────┘
              │ 1
              │
              ├──────────┐
              │          │
              ▼ N        ▼ N
┌─────────────────┐  ┌─────────────────────────┐
│  MOVIMENTAÇÃO   │  │  CONCESSIONÁRIA         │
│  MENSAL         │  │  (ex: Enel, CPFL)       │
│  - energia (kWh)│  │  - nome                  │
│  - valores (R$) │  │  - região                │
│  - economia (%) │  └─────────────────────────┘
│  - status (Pago/│
│    Pendente)    │
└─────────────────┘
```

**Regras de integridade:**
- Um Associado pode ter N Movimentações Mensais (`associadoId` FK)
- Um Associado do tipo "Beneficiado" está vinculado a uma Concessionária
- Um Associado pode ser "Hibrido" (atuar como Fornecedor e Beneficiado simultaneamente)
- Movimentações pertencem a exatamente um Associado

## 4. Regras de Negócio Fundamentais

| # | Regra | Tipo |
|---|---|---|
| RN-01 | O CPF/CNPJ do associado deve ser único no sistema | Restrição de unicidade (SQLite UNIQUE) |
| RN-02 | A senha é obrigatória no cadastro e login | Validação (yup schema) |
| RN-03 | O status do associado segue o ciclo: "Em cadastro" → "Ativo" → "Inativo"/"Bloqueado"/"Encerrado" | Máquina de estados |
| RN-04 | A economia é calculada como: `valorEconomizado = (energiaRecebidaKwh × tarifaConcessionaria) - valorCobrado` | Cálculo |
| RN-05 | O percentual economizado = `(valorEconomizado / valorTotal) × 100` | Cálculo |
| RN-06 | O associado deve aceitar os termos de uso ("Sim") para finalizar o cadastro | Validação |
| RN-07 | Pessoas Físicas devem informar CPF + Data de Nascimento; Pessoas Jurídicas devem informar CNPJ + Razão Social + Data de Abertura | Validação condicional |
| RN-08 | Fornecedores devem informar potência instalada e disponibilidade; Beneficiados devem informar concessionária e consumo médio | Validação condicional |

## 5. Estados e Ciclos de Vida

### 5.1. Status do Associado

```
[Em cadastro] ──► [Ativo] ──► [Inativo]
                      │            │
                      ├────────────┤
                      ▼            ▼
                  [Bloqueado]  [Encerrado]
```

### 5.2. Status de Pagamento

```
[Pendente] ──► [Pago]
```

## 6. Ambiguidades Resolvidas

| Ambiguidade | Resolução |
|---|---|
| O projeto original menciona "Redux Toolkit" mas não usa | O gerenciamento de estado é via Context API + React Query. Redux não faz parte da stack. |
| Existem dois backends de persistência (SQLite e AsyncStorage) | A versão canônica é SQLite (`services/database/`). AsyncStorage (`services/storage/`) é legado. |
| "Associado" vs "Usuário" | O termo canônico é **Associado** (usado no código e no domínio de negócio). |

## 7. Questões em Aberto [ASK USER]

| # | Pergunta |
|---|---|
| Q-01 | Qual o backend remoto planejado? Qual URL base da API? |
| Q-02 | A senha deve ser hasheada no cliente ou apenas no backend? |
| Q-03 | O modelo "Hibrido" (fornecedor + beneficiado) já é suportado em produção? |
| Q-04 | Existe integração real com concessionárias (API da ANEEL) ou os dados são apenas mock? |
| Q-05 | Como funciona a monetização? Os planos comerciais são reais ou placeholders? |
