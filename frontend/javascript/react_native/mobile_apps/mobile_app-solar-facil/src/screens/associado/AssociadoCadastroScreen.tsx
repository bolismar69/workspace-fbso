import React from "react";
import {
  KeyboardAvoidingView,
  Platform,
  SafeAreaView,
  ScrollView,
  View,
} from "react-native";
import { useAppTheme } from "@/context/AppThemeContext";
import { FormCadastroAssociado } from "@/components/forms/FormCadastroAssociado";
import { ContatoRodapeCopyRight } from "@/components/ContatoRodapeCopyRight";

export default function AssociadoCadastroScreen() {
  console.log("AssociadoCadastroScreen - Renderizando tela de cadastro de associado");
  const { theme } = useAppTheme();

  const handleSubmit = (data: any) => {
    console.log("AssociadoCadastroScreen - Associado cadastrado com sucesso:", data);
  };

  return (
    <KeyboardAvoidingView
      style={{ flex: 1 }}
      behavior={Platform.OS === "ios" ? "padding" : "height"}
      keyboardVerticalOffset={Platform.OS === "ios" ? 64 : 0}
    >
      <SafeAreaView style={{ flex: 1 }}>
        <ScrollView
          contentContainerStyle={{
            padding: 16,
            flexGrow: 1,
            backgroundColor: theme.screen?.backgroundColor,
          }}
          keyboardShouldPersistTaps="handled"
        >
          <View style={{ gap: 24 }}>
            <FormCadastroAssociado onSubmit={handleSubmit} />
          </View>
        </ScrollView>
        <ContatoRodapeCopyRight />
      </SafeAreaView>
    </KeyboardAvoidingView>
  );
}
