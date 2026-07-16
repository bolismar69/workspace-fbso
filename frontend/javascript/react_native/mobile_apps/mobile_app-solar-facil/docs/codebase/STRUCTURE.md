---
title: "Estrutura do Projeto — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
owner: "Time de Engenharia"
tags: ["structure", "diretórios", "organização", "mobile"]
---

# Estrutura do Projeto — Solar Fácil

## 1. Visão Geral

O projeto segue uma arquitetura modular com separação clara de responsabilidades. A raiz do projeto contém arquivos de configuração e os diretórios `src/` (código fonte) e `.specs/` (documentação).

```
mobile_app-solar-facil/
├── app.json                    ← Configuração Expo (scheme, plugins, splash, updates)
├── eas.json                    ← Configuração EAS Build (development/preview/production)
├── package.json                ← Dependências npm e scripts
├── tsconfig.json               ← Configuração TypeScript (strict, path aliases)
├── tailwind.config.js          ← Configuração NativeWind (cores, conteúdo)
├── eslint.config.js            ← Configuração ESLint
├── .nvmrc                      ← Versão do Node.js
├── .gitignore                  ← Arquivos ignorados pelo git
├── README.md                   ← Guia de instalação e uso
├── README-ARQUITETURA.md       ← Visão arquitetural
├── README-POLITICA-PRIVACIDADE.md ← Política de privacidade
│
├── .specs/                     ← Documentação técnica (SPEC-MINING)
├── docs/codebase/              ← Documentação da codebase (STACK, STRUCTURE, etc.)
├── .github/                    ← Configurações GitHub (CI/CD)
├── nativewind/                 ← Configuração extra NativeWind
│
├── scripts/
│   └── reset-project.js        ← Script de reset do projeto
│
├── src/
│   ├── App.tsx                 ← Ponto de entrada raiz
│   ├── app/                    ← Expo Router (file-based routing)
│   ├── components/             ← Componentes reutilizáveis
│   ├── screens/                ← Telas organizadas por domínio
│   ├── context/                ← React Context Providers
│   ├── hooks/                  ← Custom hooks (queries, mutations)
│   ├── services/               ← Serviços (database, mock, storage)
│   ├── types/                  ← Definições de tipos TypeScript
│   ├── styles/                 ← Temas (light, dark)
│   ├── utils/                  ← Utilitários (cn, mask, validators)
│   ├── constants/              ← Constantes
│   ├── assets/                 ← Imagens, fontes, ícones
│   └── mocks/                  ← Dados mock (.json)
│
└── .expo/                      ← Cache do Expo (gerado, não versionado)
```

## 2. Entry Points

| Arquivo | Função |
|---|---|
| `src/App.tsx` | Ponto de entrada principal da aplicação |
| `src/app/_layout.tsx` | Root layout do Expo Router — providers e tabs |
| `src/app/index.tsx` | Tela inicial (Home) |
| `package.json` → `"main": "expo-router/entry"` | Entry point do Expo Router |

## 3. Rotas (Expo Router — File-based)

| Arquivo | Rota | Título | Tab |
|---|---|---|---|
| `src/app/_layout.tsx` | `/` | Root Layout | — |
| `src/app/index.tsx` | `/` | Bem vindo ao Solar Fácil | Solar |
| `src/app/planos.tsx` | `/planos` | Planos Comerciais | Plano |
| `src/app/saibamais.tsx` | `/saibamais` | Saiba Mais | Saiba |
| `src/app/faq.tsx` | `/faq` | Perguntas Frequentes | FAQ |
| `src/app/login.tsx` | `/login` | Faça seu Login | Login |
| `src/app/cadastro.tsx` | `/cadastro` | Cadastro | User |
| `src/app/movimentacao.tsx` | `/movimentacao` | Movimentações Mensais | ... |
| `src/app/listatodos.tsx` | `/listatodos` | Lista | Lista |

## 4. Árvore de Providers (Hierarquia)

```
ReactQueryProvider
  └── AppThemeProvider (tema claro/escuro)
        └── AuthProvider (autenticação)
              └── SafeAreaProvider
                    └── StatusBar
                          └── DatabaseProvider (SQLite, autoInitialize=true)
                                └── GestureHandlerRootView
                                      └── Tabs (8 rotas)
```

## 5. Componentes

### 5.1. Componentes de UI

| Componente | Caminho | Propósito |
|---|---|---|
| CardIconeAmarelo | `components/CardIconeAmarelo.tsx` | Card com ícone e fundo amarelo |
| CardIconePadrao | `components/CardIconePadrao.tsx` | Card com ícone padrão |
| CardPlan | `components/CardPlan.tsx` | Card de plano comercial |
| FaqAccordion | `components/FaqAccordion.tsx` | Accordion para FAQ |
| SolarFacilIconeLogo | `components/SolarFacilIconeLogo.tsx` | Logotipo do Solar Fácil |
| ThemedButton | `components/ThemedButton.tsx` | Botão temático |
| TabIcon | `components/ui/TabIcon.tsx` | Ícone de tab |
| ContatoRodapeCopyRight | `components/ContatoRodapeCopyRight.tsx` | Copyright no rodapé |
| ContatoRodapeIconesContato | `components/ContatoRodapeIconesContato.tsx` | Ícones de contato no rodapé |

### 5.2. Componentes de Formulário

| Componente | Caminho | Propósito |
|---|---|---|
| FormBeneficiado | `components/forms/FormBeneficiado.tsx` | Formulário de beneficiado |
| FormCadastroAssociado | `components/forms/FormCadastroAssociado.tsx` | Formulário de cadastro de associado |
| FormCadastroAssociadoStorage | `components/forms/FormCadastroAssociadoStorage.tsx` | Cadastro com persistência em storage |
| FormDadosCadastraisAssociado | `components/forms/FormDadosCadastraisAssociado.tsx` | Dados cadastrais do associado |
| FormFornecedor | `components/forms/FormFornecedor.tsx` | Formulário de fornecedor |
| FormSection | `components/forms/FormSection.tsx` | Seção de formulário |

### 5.3. Componentes de Input

| Componente | Caminho | Propósito |
|---|---|---|
| DynamicInput | `components/inputs/DynamicInput.tsx` | Input dinâmico baseado em tipo |
| InputDate | `components/inputs/InputDate.tsx` | Input de data |
| InputPasswordWithToggle | `components/inputs/InputPasswordWithToggle.tsx` | Input de senha com toggle |
| InputRadio | `components/inputs/InputRadio.tsx` | Radio button |
| InputSelect | `components/inputs/InputSelect.tsx` | Select/dropdown |
| InputSwitch | `components/inputs/InputSwitch.tsx` | Switch/toggle |
| InputText | `components/inputs/InputText.tsx` | Input de texto |
| InputTextarea | `components/inputs/InputTextarea.tsx` | Textarea |
| KeyboardSafeScreen | `components/inputs/KeyboardSafeScreen.tsx` | Tela com keyboard avoidance |

### 5.4. Componentes de Lista

| Componente | Caminho | Propósito |
|---|---|---|
| AssociadoItem | `components/lists/AssociadoItem.tsx` | Item de lista de associado |

## 6. Telas (Screens)

### 6.1. Associado

| Tela | Caminho | Propósito |
|---|---|---|
| AssociadoCadastroScreen | `screens/associado/AssociadoCadastroScreen.tsx` | Cadastro de associado |
| AssociadoDadosCadastraisScreen | `screens/associado/AssociadoDadosCadastraisScreen.tsx` | Visualização/edição de dados cadastrais |
| AssociadoListaTodosScreen | `screens/associado/AssociadoListaTodosScreen.tsx` | Lista de todos os associados |
| AssociadoListaTodosScreenStorage | `screens/associado/AssociadoListaTodosScreenStorage.tsx` | Lista com persistência storage |
| AssociadoLoginScreen | `screens/associado/AssociadoLoginScreen.tsx` | Login de associado |
| AssociadosGerenciarScreen | `screens/associado/AssociadosGerenciarScreen.tsx` | Gerenciamento de associados |

### 6.2. Beneficiado

| Tela | Caminho |
|---|---|
| BeneficiadoCadastroScreen | `screens/beneficiado/BeneficiadoCadastroScreen.tsx` |
| BeneficiadoLoginScreen | `screens/beneficiado/BeneficiadoLoginScreen.tsx` |

### 6.3. Fornecedor

| Tela | Caminho |
|---|---|
| FornecedorCadastroScreen | `screens/fornecedor/FornecedorCadastroScreen.tsx` |

### 6.4. Geral

| Tela | Caminho |
|---|---|
| HomeScreen | `screens/general/HomeScreen.tsx` |
| PlanosScreen | `screens/general/PlanosScreen.tsx` |
| SaibaMaisScreen | `screens/general/SaibaMaisScreen.tsx` |
| FAQScreen | `screens/general/FAQScreen.tsx` |

### 6.5. Movimentação

| Tela | Caminho |
|---|---|
| MovimentacaoComCardScreen | `screens/movimentacao/MovimentacaoComCardScreen.tsx` |

## 7. Serviços

### 7.1. Database (SQLite)

| Arquivo | Propósito |
|---|---|
| `services/database/initializeSQLiteDatabase.ts` | Inicialização do SQLite, criação de tabelas |
| `services/database/useAssociados.ts` | CRUD de associados no SQLite |
| `services/database/seedMovimentacoes.ts` | Seed de dados de movimentações |

### 7.2. Storage (AsyncStorage)

| Arquivo | Propósito |
|---|---|
| `services/storage/serviceAssociado.ts` | CRUD de associados via AsyncStorage |
| `services/storage/serviceBeneficiado.ts` | CRUD de beneficiados via AsyncStorage |
| `services/storage/serviceFornecedor.ts` | CRUD de fornecedores via AsyncStorage |
| `services/storage/serviceMovimentacaoMensal.ts` | CRUD de movimentações via AsyncStorage |
| `services/storage/storageUtils.ts` | Utilitários de storage |

### 7.3. Mock Services

| Arquivo | Propósito |
|---|---|
| `services/mock/serviceBeneficiadoMock.ts` | Mock de serviço de beneficiado |
| `services/mock/serviceFornecedorMock.ts` | Mock de serviço de fornecedor |
| `services/serviceConcessionarias.ts` | Lista de concessionárias de energia |
| `services/serviceConsumoMedio.ts` | Dados de consumo médio |
| `services/serviceFAQs.ts` | FAQs |
| `services/servicePlans.ts` | Planos comerciais |

## 8. Tipos TypeScript

| Arquivo | Tipo Principal | Campos Notáveis |
|---|---|---|
| `types/AssociadoType.ts` | `AssociadoType` | 30 campos — dados pessoais, endereço, PJ/PF, fornecedor/beneficiado |
| `types/BeneficiadoType.ts` | `BeneficiadoType` | Dados de beneficiado |
| `types/ConcessionariaType.ts` | `ConcessionariaType` | Concessionária de energia |
| `types/ConsumoMedioType.ts` | `ConsumoMedioType` | Consumo médio de energia |
| `types/DBResponse.ts` | `DBResponse` | Resposta de operações de banco |
| `types/FAQType.ts` | `FAQType` | Pergunta frequente |
| `types/FieldDefinitionType.ts` | `FieldDefinitionType` | Definição de campo de formulário |
| `types/FornecedorType.ts` | `FornecedorType` | Dados de fornecedor |
| `types/MovimentacaoMensalType.ts` | `MovimentacaoMensalType` | 18 campos — movimentação mensal de energia |
| `types/PlanType.ts` | `PlanType` | Plano comercial |
| `types/ThemeTypes.ts` | `AppThemeStyles` | Estilos do tema (~50 propriedades) |

## 9. Contextos

| Contexto | Caminho | Estado Gerenciado |
|---|---|---|
| AppThemeContext | `context/AppThemeContext.tsx` | Tema claro/escuro (segue preferência do sistema) |
| AuthContext | `context/AuthContext.tsx` | Autenticação (isLoggedIn, userID, userName, associado) |
| DatabaseContext | `context/DatabaseContext.tsx` | Conexão SQLite (isDatabaseConnected, dbInstance) |
| ReactQueryProvider | `context/ReactQueryProvider.tsx` | Provider do TanStack React Query |

## 10. Hooks

### 10.1. Queries (React Query)

| Hook | Caminho | Query |
|---|---|---|
| useQueryAssociadosSearchAll | `hooks/queries/useQueryAssociadosSearchAll.ts` | Buscar todos associados |
| useQueryAssociadosSearchByCpfCnpjSenha | `hooks/queries/useQueryAssociadosSearchByCpfCnpjSenha.ts` | Buscar por CPF/CNPJ + senha |
| useQueryAssociadosSearchById | `hooks/queries/useQueryAssociadosSearchById.ts` | Buscar por ID |
| useQueryMovimentacoesSearchAll | `hooks/queries/useQueryMovimentacoesSearchAll.ts` | Buscar todas movimentações |
| useQueryMovimentacoesSearchByAssociadoId | `hooks/queries/useQueryMovimentacoesSearchByAssociadoId.ts` | Buscar movimentações por associado |

### 10.2. Mutations (React Query)

| Hook | Caminho | Mutation |
|---|---|---|
| useMutationAssociadoInsertRecord | `hooks/mutations/useMutationAssociadoInsertRecord.ts` | Inserir associado |
| useMutationAssociadoUpdateRecord | `hooks/mutations/useMutationAssociadoUpdateRecord.ts` | Atualizar associado |
| useMutationAssociadoDeleteRecord | `hooks/mutations/useMutationAssociadoDeleteRecord.ts` | Deletar associado |
| useMutationMovimentacoesInsertRecord | `hooks/mutations/useMutationMovimentacoesInsertRecord.ts` | Inserir movimentação |
| useMutationMovimentacoesUpdateRecord | `hooks/mutations/useMutationMovimentacoesUpdateRecord.ts` | Atualizar movimentação |
| useMutationMovimentacoesDeleteRecord | `hooks/mutations/useMutationMovimentacoesDeleteRecord.ts` | Deletar movimentação |
| useMutationSeedMovimentacoes | `hooks/mutations/useMutationSeedMovimentacoes.ts` | Seed de movimentações |

### 10.3. Custom Hooks

| Hook | Caminho | Propósito |
|---|---|---|
| useFormValidation | `hooks/useFormValidation.ts` | Validação de formulários (react-hook-form + yup) |

## 11. Estilos

| Arquivo | Propósito |
|---|---|
| `styles/lightTheme.ts` | Tema claro (~280 linhas, 50+ estilos) |
| `styles/darkTheme.ts` | Tema escuro |
| `tailwind.config.js` | Cores do Tailwind (primary, secondary, accent, background, neutral) |

## 12. Utilitários

| Arquivo | Propósito |
|---|---|
| `utils/cn.ts` | Utilitário de classes (similar a clsx) |
| `utils/mask.ts` | Máscaras de input (CPF, CNPJ, telefone, etc.) |
| `utils/validateRG.ts` | Validação de RG |
| `utils/validates/validateFormData.ts` | Validação de dados de formulário |
| `utils/validators/validatorCNPJ.ts` | Validador de CNPJ |
| `utils/validators/validatorCPF.ts` | Validador de CPF |
| `utils/validators/validatorRG.ts` | Validador de RG |
| `utils/validators/validators.ts` | Validadores gerais |
