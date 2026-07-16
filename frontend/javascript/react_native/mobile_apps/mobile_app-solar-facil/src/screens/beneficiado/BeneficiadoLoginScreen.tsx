// src/screens/beneficiado/BeneficiadoScreen.tsx
import React, { useState } from "react";
import {
  KeyboardAvoidingView,
  Platform,
  SafeAreaView,
  ScrollView,
  View,
  Text,
  TextInput,
  TouchableOpacity,
} from "react-native";
import { useAppTheme } from "@/context/AppThemeContext";
import { FormBeneficiado } from "@/components/forms/FormBeneficiado";

export default function BeneficiadoLoginScreen() {
  const { theme } = useAppTheme();
  const [showForm, setShowForm] = useState(false);
  const [cpf, setCpf] = useState("");
  const [senha, setSenha] = useState("");

  const handleLogin = () => {
    // Simulação de login simples
    if (cpf && senha) {
      console.log("Login simulado com CPF:", cpf);
      // Aqui você poderia redirecionar ou buscar dados
    } else {
      alert("Preencha CPF e senha");
    }
  };

  const handleSubmit = (data: any) => {
    console.log("Beneficiado cadastrado com sucesso:", data);
    // Aqui pode integrar com API, salvar localmente, etc.
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
            {!showForm ? (
              <>
                <Text style={theme.title}>Login do Beneficiário</Text>

                <View style={{ gap: 16 }}>
                  <View>
                    <Text style={theme.label}>CPF</Text>
                    <TextInput
                      style={theme.input}
                      value={cpf}
                      onChangeText={setCpf}
                      keyboardType="numeric"
                      placeholder="000.000.000-00"
                    />
                  </View>

                  <View>
                    <Text style={theme.label}>Senha</Text>
                    <TextInput
                      style={theme.input}
                      value={senha}
                      onChangeText={setSenha}
                      secureTextEntry
                      placeholder="Digite sua senha"
                    />
                  </View>

                  <TouchableOpacity
                    style={theme.button}
                    onPress={handleLogin}
                  >
                    <Text style={theme.buttonText}>Entrar</Text>
                  </TouchableOpacity>

                  <TouchableOpacity
                    onPress={() => setShowForm(true)}
                    style={{ marginTop: 12 }}
                  >
                    <Text style={[theme.text, { textAlign: "center" }]}>
                      Ainda não sou beneficiado. Quero me cadastrar.
                    </Text>
                  </TouchableOpacity>
                </View>
              </>
            ) : (
              <FormBeneficiado onSubmit={handleSubmit} />
            )}
          </View>
        </ScrollView>
      </SafeAreaView>
    </KeyboardAvoidingView>
  );
}
