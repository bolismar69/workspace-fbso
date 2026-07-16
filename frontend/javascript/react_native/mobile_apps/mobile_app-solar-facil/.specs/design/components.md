---
title: "Catálogo de Componentes — Solar Fácil"
version: "1.0"
date_created: "2026-07-08"
last_updated: "2026-07-08"
---

# Catálogo de Componentes — Solar Fácil

## 1. Componentes de Formulário

### FormSection
- **Caminho:** `src/components/forms/FormSection.tsx`
- **Propósito:** Wrapper de seção de formulário com título e children
- **Props:** `title: string`, `children: ReactNode`

### FormBeneficiado
- **Caminho:** `src/components/forms/FormBeneficiado.tsx`
- **Propósito:** Formulário completo de cadastro de beneficiado
- **Dependências:** `InputText`, `InputSelect`, `InputDate`, `useFormValidation`

### FormCadastroAssociado
- **Caminho:** `src/components/forms/FormCadastroAssociado.tsx`
- **Propósito:** Formulário completo de cadastro de associado
- **Variantes:** Pessoa Física / Pessoa Jurídica
- **Dependências:** `InputText`, `InputSelect`, `InputDate`, `InputRadio`, `InputSwitch`

### FormCadastroAssociadoStorage
- **Caminho:** `src/components/forms/FormCadastroAssociadoStorage.tsx`
- **Propósito:** Versão com persistência AsyncStorage
- **Diferença:** Dados salvos em AsyncStorage em vez de SQLite

### FormDadosCadastraisAssociado
- **Caminho:** `src/components/forms/FormDadosCadastraisAssociado.tsx`
- **Propósito:** Visualização e edição de dados cadastrais existentes

### FormFornecedor
- **Caminho:** `src/components/forms/FormFornecedor.tsx`
- **Propósito:** Formulário completo de cadastro de fornecedor

## 2. Componentes de Input

### InputText
- **Caminho:** `src/components/inputs/InputText.tsx`
- **Props:** `name`, `control`, `errors`, `label`, `placeholder`, `keyboardType?`
- **Estados:** default, error (borda vermelha), focused

### InputDate
- **Caminho:** `src/components/inputs/InputDate.tsx`
- **Props:** `name`, `control`, `errors`, `label`
- **Nativo:** `@react-native-community/datetimepicker`

### InputSelect
- **Caminho:** `src/components/inputs/InputSelect.tsx`
- **Props:** `name`, `control`, `errors`, `label`, `options: {label, value}[]`
- **Nativo:** `@react-native-picker/picker`

### InputRadio
- **Caminho:** `src/components/inputs/InputRadio.tsx`
- **Props:** `name`, `control`, `errors`, `label`, `options`

### InputSwitch
- **Caminho:** `src/components/inputs/InputSwitch.tsx`
- **Props:** `name`, `control`, `label`
- **Nativo:** React Native `Switch`

### InputPasswordWithToggle
- **Caminho:** `src/components/inputs/InputPasswordWithToggle.tsx`
- **Props:** `name`, `control`, `errors`, `label`
- **Comportamento:** Toggle de visibilidade (olho) + ícone de senha

### InputTextarea
- **Caminho:** `src/components/inputs/InputTextarea.tsx`
- **Props:** `name`, `control`, `errors`, `label`, `numberOfLines?`

### DynamicInput
- **Caminho:** `src/components/inputs/DynamicInput.tsx`
- **Props:** `field: FieldDefinitionType`, `control`, `errors`
- **Comportamento:** Renderiza o input correto baseado em `field.type`

### KeyboardSafeScreen
- **Caminho:** `src/components/inputs/KeyboardSafeScreen.tsx`
- **Props:** `children: ReactNode`
- **Comportamento:** KeyboardAvoidingView + ScrollView + SafeAreaView

## 3. Componentes de Card

### CardIconeAmarelo
- **Caminho:** `src/components/CardIconeAmarelo.tsx`
- **Props:** `icon: string`, `title: string`, `description?: string`
- **Estilo:** Fundo amarelo (`#ffffbf`), ícone + título + descrição

### CardIconePadrao
- **Caminho:** `src/components/CardIconePadrao.tsx`
- **Props:** `icon: string`, `title: string`, `description?: string`
- **Estilo:** Fundo branco, layout similar ao CardIconeAmarelo

### CardPlan
- **Caminho:** `src/components/CardPlan.tsx`
- **Props:** `plan: PlanType`
- **Estilo:** Card temático de plano comercial

## 4. Componentes de Lista

### AssociadoItem
- **Caminho:** `src/components/lists/AssociadoItem.tsx`
- **Props:** `associado: AssociadoType`, `onPress?: () => void`
- **Estilo:** Card com nome, CPF/CNPJ, status e botão de ação

## 5. Componentes de UI

### ThemedButton
- **Caminho:** `src/components/ThemedButton.tsx`
- **Props:** `title: string`, `onPress: () => void`, `variant?: "primary" | "secondary" | "link" | "low"`
- **Variantes:** primary (`#1E5631`), secondary (`#A5C9CA`), link (`#1E90FF`), low (height 32px)
- **Estados:** default, pressed (TODO)

### TabIcon
- **Caminho:** `src/components/ui/TabIcon.tsx`
- **Props:** `icon: string`, `color: string`, `size: number`, `focused: boolean`

### SolarFacilIconeLogo
- **Caminho:** `src/components/SolarFacilIconeLogo.tsx`
- **Props:** `size?: number`
- **Estilo:** Logotipo do Solar Fácil

### FaqAccordion
- **Caminho:** `src/components/FaqAccordion.tsx`
- **Props:** `faqs: FAQType[]`
- **Comportamento:** Accordion expand/colapsa com título (pergunta) e conteúdo (resposta)

### ContatoRodapeCopyRight
- **Caminho:** `src/components/ContatoRodapeCopyRight.tsx`
- **Estilo:** Texto de copyright no rodapé

### ContatoRodapeIconesContato
- **Caminho:** `src/components/ContatoRodapeIconesContato.tsx`
- **Estilo:** Ícones de contato (WhatsApp, Email, etc.) no rodapé

## 6. Platform Specifics

| Componente | iOS | Android | Diferença |
|---|---|---|---|
| InputDate | DateTimePicker (spinner) | DateTimePicker (calendar) | Estilo nativo diferente |
| InputSelect | Picker (wheel) | Picker (dropdown) | Estilo nativo diferente |
| KeyboardSafeScreen | KeyboardAvoidingView (padding) | KeyboardAvoidingView (height) | Comportamento diferente |
| SafeAreaView | Notch/Dynamic Island/Home Indicator | StatusBar/NavigationBar | Áreas seguras diferentes |
