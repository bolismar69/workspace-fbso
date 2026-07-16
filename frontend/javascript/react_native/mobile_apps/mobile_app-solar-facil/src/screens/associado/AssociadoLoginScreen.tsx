// /src/screens/associado/AssociadoLoginScreen.tsx
import { useRouter } from "expo-router";
import { View, Text, TextInput, TouchableOpacity, SafeAreaView, KeyboardAvoidingView, Platform, ScrollView } from "react-native";
import React, { useState } from "react";
import { useAppTheme } from "@/context/AppThemeContext";
import { useAuth } from "@/context/AuthContext";
import { Ionicons } from "@expo/vector-icons";
import { ContatoRodapeCopyRight } from "@/components/ContatoRodapeCopyRight";
import { useAssociadosCopilot } from "@/services/database/useAssociadosCopilot";

export default function AssociadoLoginScreen() {
  console.log("=== INICIO ========================================================================");
  console.log("AssociadoLoginScreen - Iniciando tela de login...");

  const { theme } = useAppTheme();
  const { login, logout, isLoggedIn } = useAuth();
  const router = useRouter();

  const useAssociados = useAssociadosCopilot(); // Call the hook here

  const [identificador, setIdentificador] = useState("");
  const [senha, setSenha] = useState("");
  const [mostrarSenha, setMostrarSenha] = useState(false);
  const [mensagemErro, setMensagemErro] = useState<string | null>(null);
  const [loading, setLoading] = useState(false);

  const handleLogin = async () => {
    console.log("=== INICIO ========================================================================");
    console.log("AssociadoLoginScreen - LOGIN - Iniciando login...");
    setMensagemErro(null);
    setLoading(true);

    try {
      // buscar o associado pelo CPF/CNPJ e senha
      const resultado = await (await useAssociados).searchByCpfCnpjSenha(identificador, senha);
      if (resultado.success === false) {
        console.error("AssociadoLoginScreen - LOGIN - Erro ao buscar associado:", resultado.error);
        setMensagemErro("Erro ao buscar associado. Tente novamente.");
        return;
      }

      if (resultado && resultado.data && resultado.data.length > 0) {
        console.log("AssociadoLoginScreen - LOGIN - Associados encontrados:", resultado.data.length, "associado => ", resultado.data[0]);

        login(resultado.data[0].id, resultado.data[0].nome, resultado.data[0]); // associado vem da resposta do serviço
        router.push({
          pathname: "/cadastro",
        });
      } else {
        setMensagemErro("CPF/CNPJ ou senha inválidos.");
      }
    } catch (error: any) {
      setMensagemErro("Erro ao tentar login. Tente novamente.");
      console.error("AssociadoLoginScreen - Erro no login:", error);
    } finally {
      setLoading(false);
    }
    console.log("=== FIM ========================================================================");
  };

  return (
    <KeyboardAvoidingView
      style={{ flex: 1 }}
      behavior={Platform.OS === "ios" ? "padding" : "height"}
      keyboardVerticalOffset={Platform.OS === "ios" ? 64 : 0}
    >
      <SafeAreaView style={{ flex: 1 }}>
        <ScrollView style={[{ padding: 16, backgroundColor: theme.basicView.backgroundColor }]}>

          {/* Esta parte pedira o login caso não esteja logado */}
          {isLoggedIn === false
            ? (
              // <View style={theme.container}>
              <View style={theme.card}>
                <Text style={[theme.subtitle, { marginBottom: 16 }]}>
                  Informe seu CPF ou CNPJ para entrar
                </Text>

                <View style={[{ flexDirection: "row", alignItems: "center", marginBottom: 12, marginLeft: 32, marginRight: 32 }]}>
                  <TextInput
                    placeholder="CPF ou CNPJ, cadastrado"
                    style={[theme.input, { flex: 1 }]}
                    value={identificador}
                    onChangeText={setIdentificador}
                  // autoCapitalize="none"
                  />
                </View>

                {/* Campo de senha com botão de ver/ocultar */}
                <View style={[{ flexDirection: "row", alignItems: "center", marginBottom: 12, marginLeft: 32, marginRight: 32 }]}>
                  <TextInput
                    placeholder="Senha"
                    value={senha}
                    onChangeText={setSenha}
                    secureTextEntry={!mostrarSenha}
                    style={[theme.input, { flex: 1 }]}
                    placeholderTextColor={theme.placeholder.color}
                  />
                  <TouchableOpacity onPress={() => setMostrarSenha(!mostrarSenha)}>
                    <Ionicons
                      name={mostrarSenha ? "eye" : "eye-off"}
                      size={24}
                      color={theme.placeholder.color}
                    />
                  </TouchableOpacity>
                </View>

                {mensagemErro && (
                  <Text style={{ color: "red", marginBottom: 12, textAlign: "center" }}>
                    {mensagemErro}
                  </Text>
                )}

                <TouchableOpacity
                  onPress={handleLogin}
                  disabled={loading}
                  style={[theme.button, { padding: 16 }]}
                >
                  <Text style={[theme.buttonText, { fontWeight: "600", textAlign: "center" }]}>
                    {loading ? "Validando..." : "ENTRAR"}
                  </Text>
                </TouchableOpacity>

                <Text style={[theme.text, { marginTop: 32, marginBottom: 16 }]}>ou</Text>

                <TouchableOpacity
                  onPress={() => router.push("/cadastro")}
                  style={[theme.button, { padding: 16 }]}
                >
                  <Text style={[theme.buttonText, { fontWeight: "600", textAlign: "center" }]}>
                    Cadastre-se e venha fazer parte da nossa comunidade!
                  </Text>
                </TouchableOpacity>
              </View>
            ) : (
              <View style={theme.container}>
                {/* adicionar botão para fazer logout */}
                <TouchableOpacity
                  onPress={() =>
                    logout()
                  }
                  style={[theme.button, { padding: 16 }]}
                >
                  <Text style={[theme.buttonText, { fontWeight: "600", textAlign: "center" }]}>
                    Clique aqui para fazer logout
                  </Text>
                </TouchableOpacity>
              </View>
            )}
        </ScrollView>
        <ContatoRodapeCopyRight />
      </SafeAreaView>
    </KeyboardAvoidingView >
  );
}
