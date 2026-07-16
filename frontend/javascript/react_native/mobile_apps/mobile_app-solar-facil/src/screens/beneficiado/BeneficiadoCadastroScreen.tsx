// src/screens/beneficiado/BeneficiadoCadastroScreen.tsx
import React from "react";
import {
  KeyboardAvoidingView,
  Platform,
  SafeAreaView,
  ScrollView,
  View,
} from "react-native";
import { useAppTheme } from "@/context/AppThemeContext";
import { FormBeneficiado } from "@/components/forms/FormBeneficiado";

export default function BeneficiadoCadastroScreen() {
  console.log("BeneficiadoScreen rendered");
  const { theme } = useAppTheme();

  const handleSubmit = (data: any) => {
    console.log("Beneficiado cadastrado com sucesso:", data);
    // Aqui pode integrar com API, salvar localmente, etc.
  };

  return (
    <KeyboardAvoidingView
      style={{ flex: 1 }}
      behavior={Platform.OS === "ios" ? "padding" : "height"}
      // ou ajustar conforme cabeçalho
      keyboardVerticalOffset={Platform.OS === "ios" ? 64 : 0}
    >
      <SafeAreaView style={[{ flex: 1 }]}>
        <ScrollView
          contentContainerStyle={{
            padding: 16,
            flexGrow: 1,
            backgroundColor: theme.screen?.backgroundColor,
          }}
          keyboardShouldPersistTaps="handled"
        >
          <View style={{ gap: 24 }}>
            <FormBeneficiado onSubmit={handleSubmit} />
          </View>
        </ScrollView>
      </SafeAreaView>
    </KeyboardAvoidingView>
  );
}
