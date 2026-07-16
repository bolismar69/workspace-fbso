// src/screens/associado/AssociadoListaTodosScreen.tsx
import React, { useEffect, useState } from "react";
import { View, Text, ScrollView, TouchableOpacity, Alert, KeyboardAvoidingView, Platform, SafeAreaView } from "react-native";
import { useAppTheme } from "@/context/AppThemeContext";
import { Ionicons } from "@expo/vector-icons";
import { AssociadoType } from "@/types/AssociadoType";
// import { listarAssociados, excluirAssociado } from "@/services/storage/serviceAssociado";
// import { listarBeneficiados, excluirBeneficiado } from "@/services/storage/serviceBeneficiado";
// import { listarFornecedores, excluirFornecedor } from "@/services/storage/serviceFornecedor";
import { useRouter } from "expo-router";
import { limparArmazenamento, clearStorageDocumentDirectory } from "@/services/storage/storageUtils";
import { ContatoRodapeCopyRight } from "@/components/ContatoRodapeCopyRight";
import { useAssociadosCopilot } from "@/services/database/useAssociadosCopilot";

type Categoria = "Associado" | "Beneficiado" | "Fornecedor";

interface ItemProps {
  item: any;
  categoria: Categoria;
  onEditar: () => void;
  onExcluir: () => void;
  onDetalhes: () => void;
}

const ItemComAcoes = ({ item, categoria, onEditar, onExcluir, onDetalhes }: ItemProps) => {
  const { theme } = useAppTheme();

  return (
    <View style={[theme.card, { marginBottom: 12 }]}>
      <Text style={[theme.subtitle, { textAlign: "left" }]}>
        {item.id} - {item.nome}
      </Text>
      <Text style={theme.text}>CPF/CNPJJ: {item.cpf_cnpj}</Text>
      <Text style={theme.text}>Email: {item.email}</Text>
      <Text style={theme.text}>Telefone: {item.telefone}</Text>
      <View style={{ flexDirection: "row", gap: 12, marginTop: 8 }}>
        <TouchableOpacity onPress={onEditar}>
          <Ionicons name="create-outline" size={20} color="#2563EB" />
        </TouchableOpacity>
        <TouchableOpacity onPress={onExcluir}>
          <Ionicons name="trash-outline" size={20} color="#DC2626" />
        </TouchableOpacity>
        <TouchableOpacity onPress={onDetalhes}>
          <Ionicons name="arrow-forward-circle-outline" size={20} color="#4B5563" />
        </TouchableOpacity>
      </View>
    </View>
  );
};

export default function AssociadoListaTodosScreen() {
  console.log("AssociadoListaTodosScreen");
  const { theme } = useAppTheme();
  const router = useRouter();

  const useAssociados = useAssociadosCopilot(); // Call the hook here

  const [associados, setAssociados] = useState<AssociadoType[]>([]);
  // const [beneficiados, setBeneficiados] = useState<any[]>([]);
  // const [fornecedores, setFornecedores] = useState<any[]>([]);

  const carregarDados = async () => {
    console.log("Iniciando o processo de leitura de dados...");

    const associadosList = await (await useAssociados).searchAll();
    if (associadosList.success === false) {
      console.error("Erro ao carregar associados:", associadosList.error);
      Alert.alert("Erro", "Não foi possível carregar os associados.");
      return;
    }
    console.log("Associados carregados:", associadosList.data);

    setAssociados(associadosList.data ?? []);
    // setAssociados(await listarAssociados());
    // setBeneficiados(await listarBeneficiados());
    // setFornecedores(await listarFornecedores());
  };

  useEffect(() => {
    carregarDados();
  }, []);

  const confirmarRemocao = (
    categoria: Categoria,
    item: any,
    remover: () => Promise<void> | null
  ) => {
    Alert.alert(
      `Excluir ${categoria}`,
      `Deseja realmente excluir ${item.nome}?`,
      [
        { text: "Cancelar", style: "cancel" },
        {
          text: "Excluir",
          style: "destructive",
          onPress: async () => {
            await remover();
            carregarDados();
          },
        },
      ]
    );
  };

  const apagarArquivosLocais = async (modo: "DOC_DIR" | "BASE_DIR") => {
    Alert.alert(
      "Limpar Armazenamento",
      "Você tem certeza que deseja apagar todos os arquivos locais?",
      [
        {
          text: "Cancelar",
          style: "cancel",
          onPress: () => { console.log("Operação cancelada") }
        },
        {
          text: "Limpar",
          style: "destructive",
          onPress: async () => {

            console.log("Apagando todos os arquivos locais...");

            console.log("Apagando arquivos locais...");
            if (modo === "BASE_DIR") {
              console.log("BASE_DIR - Apagando arquivos do diretório base...");
              await limparArmazenamento();
              console.log("BASE_DIR - Arquivos do diretório base apagados.");
            } else if (modo === "DOC_DIR") {
              console.log("DOC_DIR - Apagando arquivos do diretório de documentos...");
              await clearStorageDocumentDirectory();
              console.log("DOC_DIR - Arquivos do diretório de documentos apagados.");
            }

            console.log("Recarregando dados após limpeza...");
            carregarDados();
            Alert.alert("Sucesso", "Todos os arquivos locais foram apagados.");
          },
        },
      ]
    );
  };

  return (
    <KeyboardAvoidingView
      style={{ flex: 1 }}
      behavior={Platform.OS === "ios" ? "padding" : "height"}
      keyboardVerticalOffset={Platform.OS === "ios" ? 64 : 0}
    >
      <SafeAreaView style={{ flex: 1 }}>

        <TouchableOpacity onPress={() => { apagarArquivosLocais("DOC_DIR") }}>
          <Text>🗑️ Apagar todos os arquivos locais - DOC_DIR </Text>
        </TouchableOpacity>
        <TouchableOpacity onPress={() => { apagarArquivosLocais("BASE_DIR") }}>
          <Text>🗑️ Apagar todos os arquivos locais - BASE_DIR </Text>
        </TouchableOpacity>

        <ScrollView style={[{ padding: 16, backgroundColor: theme.basicView.backgroundColor }]}>
          {/* <Text style={theme.title}>Todos os Cadastrados</Text> */}

          {[...associados.map((item) => (
            <ItemComAcoes
              key={`assoc-${item.cpf_cnpj}`}
              item={item}
              categoria="Associado"
              onEditar={() => console.log("Editar associado")}
              onExcluir={() =>
                confirmarRemocao("Associado", item, async () => Promise.resolve())
              }
              onDetalhes={() =>
                router.push({
                  pathname: "/cadastro",
                  params: { ...item },
                })
              }
            />
          ))]}

          {/* {[...beneficiados.map((item, index) => (
            <ItemComAcoes
              key={item.cpf}
              item={item}
              categoria="Beneficiado"
              onEditar={() => console.log("Editar beneficiado")}
              onExcluir={() =>
                confirmarRemocao("Beneficiado", item, () => excluirBeneficiado(item.cpf))
              }
              onDetalhes={() => console.log("Detalhes beneficiado")}
            />
          ))]} */}

          {/* {[...fornecedores.map((item, index) => (
            <ItemComAcoes
              key={item.cnpj}
              item={item}
              categoria="Fornecedor"
              onEditar={() => console.log("Editar fornecedor")}
              onExcluir={() =>
                confirmarRemocao("Fornecedor", item, () => excluirFornecedor(item.cnpj))
              }
              onDetalhes={() => console.log("Detalhes fornecedor")}
            />
          ))]} */}
        </ScrollView>
        <ContatoRodapeCopyRight />
      </SafeAreaView>
    </KeyboardAvoidingView>
  );
}
