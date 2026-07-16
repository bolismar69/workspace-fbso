// /src/screens/associado/AssociadosGerenciarScreen.tsx
import React from "react";
import { useRouter } from "expo-router";
import { View, Text, FlatList, Alert, TouchableOpacity, ScrollView } from "react-native";
import { useAppTheme } from "@/context/AppThemeContext";
import { Ionicons, MaterialCommunityIcons } from "@expo/vector-icons";
import { useQueryAssociadosSearchAll } from "@/hooks/queries/useQueryAssociadosSearchAll";
import { useMutationAssociadoDeleteRecord } from "@/hooks/mutations";
import { AssociadoType } from "@/types/AssociadoType";
import { useNavigation } from "@react-navigation/native";

interface ItemProps {
  item: AssociadoType;
  onEditar: () => void;
  onExcluir: () => void;
  onDetalhes: () => void;
  onNovo: () => void;
  onMovimento: () => void;
  onCancelar: () => void;
  onAtivar: () => void;
  onBloquear: () => void;
}

const ItemComAcoes = ({ item, onEditar, onExcluir, onDetalhes, onNovo, onMovimento, onCancelar, onAtivar, onBloquear }: ItemProps) => {
  const { theme } = useAppTheme();
  return (
    <View style={[theme.card, { marginBottom: 12 }]}>
      <Text style={[theme.subtitle, { textAlign: "left" }]}>{item.id} - {item.nome}</Text>
      <Text style={theme.text}>CPF/CNPJ: {item.cpf_cnpj}</Text>
      <Text style={theme.text}>Email: {item.email}</Text>
      <Text style={theme.text}>Telefone: {item.telefone}</Text>
      <View style={{ flexDirection: "row", flexWrap: "wrap", gap: 12, marginTop: 8 }}>
        <TouchableOpacity onPress={onDetalhes}><MaterialCommunityIcons name="account-details" size={20} color="#4B5563" /></TouchableOpacity>
        <TouchableOpacity onPress={onEditar}><MaterialCommunityIcons name="account-edit" size={20} color="#2563EB" /></TouchableOpacity>
        <TouchableOpacity onPress={onExcluir}><Ionicons name="trash-outline" size={20} color="#DC2626" /></TouchableOpacity>
        <TouchableOpacity onPress={onNovo}><Ionicons name="person-add" size={20} color="#16A34A" /></TouchableOpacity>
        <TouchableOpacity onPress={onMovimento}><MaterialCommunityIcons name="account-cash" size={20} color="#FACC15" /></TouchableOpacity>
        <TouchableOpacity onPress={onCancelar}><MaterialCommunityIcons name="account-off" size={20} color="#DC2626" /></TouchableOpacity>
        <TouchableOpacity onPress={onAtivar}><MaterialCommunityIcons name="account-check" size={24} color="green" /></TouchableOpacity>
        <TouchableOpacity onPress={onBloquear}><MaterialCommunityIcons name="account-lock" size={24} color="red" /></TouchableOpacity>
        <TouchableOpacity onPress={onAtivar}><MaterialCommunityIcons name="account-lock-open" size={24} color="blue" /></TouchableOpacity>
      </View>
    </View>
  );
};

export default function AssociadosGerenciarScreen() {
  const { data, isLoading } = useQueryAssociadosSearchAll();
  const mutationDelete = useMutationAssociadoDeleteRecord();
  const navigation = useNavigation();
  const router = useRouter();

  const handleExcluir = (id: number) => {
    Alert.alert("Confirmar Exclusão", "Deseja realmente excluir este associado?", [
      { text: "Cancelar", style: "cancel" },
      {
        text: "Confirmar",
        onPress: () => mutationDelete.mutate(id),
        style: "destructive",
      },
    ]);
  };

  return (
    <ScrollView contentContainerStyle={{ padding: 16 }}>
      <Text style={{ fontSize: 22, fontWeight: "bold", marginBottom: 16 }}>Gerenciar Associados</Text>
      {isLoading ? <Text>Carregando...</Text> : (
        data?.data?.map((item) => (
          <ItemComAcoes
            key={item.id}
            item={item}

            onEditar={() =>
              router.push({
                pathname: "/cadastro",
                params: {
                  associado: JSON.stringify(item),
                  editable: "true"
                },
              })
            }

            onDetalhes={() =>
              router.push({
                pathname: "/cadastro",
                params: {
                  associado: JSON.stringify(item),
                  editable: "false"
                },
              })
            }

            onExcluir={() => handleExcluir(item.id)}

            onNovo={() =>
              router.push({
                pathname: "/cadastro",
                params: {
                  associado: JSON.stringify(item),
                  editable: "true"
                },
              })
            }

            // onMovimento={() => navigation.navigate("FormMovimentacao", { associadoId: item.id })}
            onMovimento={() =>
              router.push({
                pathname: "/movimentacao",
                params: {
                  associado: JSON.stringify(item),
                },
              })
            }

            onCancelar={() => { }}
            onAtivar={() => { }}
            onBloquear={() => { }}
          />
        ))
      )}
    </ScrollView>
  );
}
