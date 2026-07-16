// src/screens/movimentacao/MovimentacaoComCardScreen.tsx
import React, { useEffect, useState } from "react";
import { View, Text, FlatList, useColorScheme } from "react-native";
import { useAppTheme } from "@/context/AppThemeContext";
import { useAuth } from "@/context/AuthContext";
import { useMovimentacoesCopilot } from "@/services/database/useMovimentacoesCopilot";
import { MovimentacaoMensalType } from "@/types/MovimentacaoMensalType";
import { FontAwesome5, Octicons } from "@expo/vector-icons";

const getStatusBackgroundColor = (status: string, mode: string) => {
  console.log("getStatusBackgroundColor - status:", status, "mode:", mode);
  const isDark = mode === "dark";
  switch (status) {
    case "Pago":
      return isDark ? "#144d38" : "#e0f7e9"; // verde escuro / verde claro
    case "Pendente":
      return isDark ? "#4b1c1c" : "#fbe9e7"; // vermelho escuro / vermelho claro
    default:
      return isDark ? "#4a4a2e" : "#fffde7"; // amarelo queimado / amarelo claro
  }
};

export default function MovimentacaoComCardScreen() {
  const { theme } = useAppTheme();
  const { isLoggedIn, userID } = useAuth();
  const [movimentacoes, setMovimentacoes] = useState<MovimentacaoMensalType[]>([]);

  const colorScheme = useColorScheme();

  console.log("MovimentacaoScreen - vai chamar o useMovimentacoesCopilot");
  const useMovimentacoes = useMovimentacoesCopilot(); // Call the hook here
  console.log("MovimentacaoScreen - chamou o useMovimentacoesCopilot");

  useEffect(() => {
    const carregarDados = async () => {
      if (userID == null) {
        console.warn("MovimentacaoScreen - carregarDados - userID é nulo, não é possível buscar movimentações.");
        setMovimentacoes([]);
        return;
      }
      const dados = await (await useMovimentacoes).searchByAssociadoId(userID);
      console.log("MovimentacaoScreen - Dados recebidos:", dados);
      if (dados.success === false) {
        console.error("MovimentacaoScreen - Erro ao buscar movimentações:", dados.error);
        setMovimentacoes([]);
        return;
      }
      if (JSON.stringify(dados.data) !== JSON.stringify(movimentacoes)) {
        setMovimentacoes(dados.data ?? []);
      }
      console.log("MovimentacaoScreen - Movimentações carregadas com sucesso.");
    };
    carregarDados();
  }, []);

  const renderItem = ({ item }: { item: MovimentacaoMensalType }) => {
    return (
      <View
        style={
          [theme.cardContent,
          {
            margin: 4,
            borderWidth: 5,
            borderColor: item.dataPagamento ? theme.primary : "#ca0505",
            padding: 4,
          }]
        }
      >
        <Text
          style={
            [theme.cardTitle,
            {
              textAlign: "left",
              color: theme.buttonText.color,
              backgroundColor: item.dataPagamento ? theme.button.backgroundColor : "#ca0505",
              borderRadius: 4,
              padding: 8,
            }]
          }
        >
          Mês/Ano: {String(item.mes).padStart(2, "0")}/{item?.ano}
        </Text>

        <View style={[theme.cardContent, { marginBottom: 0, borderWidth: 0.5, borderColor: theme.primary, padding: 4 }]}>
          <Text style={theme.cardTitle}>🔋 Informações do Consumo</Text>
          <Text style={theme.text}>🔌 Energia Recebida: {item.energiaRecebidaKwh} kWh</Text>
          <Text style={theme.text}>
            ⚡ Tarifa Unitaria: R$ {(item.valorCobrado / item.energiaRecebidaKwh).toFixed(2)} /kWh
          </Text>
          {/* <Text style={theme.text}>💰 Valor Cobrado: R$ {item.valorCobrado.toFixed(2)}</Text> */}
          <Text style={theme.text}>💰 Valor Cobrado: {item.valorCobrado.toLocaleString("ptbr", { style: "currency", currency: "brf" })}</Text>
        </View>

        <View
          style={
            [theme.cardContent,
            {
              marginBottom: 0,
              borderWidth: 0.5,
              borderColor: theme.primary,
              // backgroundColor:
              //   item.statusPagamento === "Pago" ? "#e0f7e9" :
              //     item.statusPagamento === "Pendente" ? "#fbe9e7" :
              // backgroundColor: getStatusBackgroundColor(item.statusPagamento ?? "", colorScheme ?? ""),
              padding: 4
            }]
          }
        >
          <Text style={[theme.cardTitle, { marginTop: 12 }]}>📆 Informações de Pagamento</Text>
          {/* <View style={{ flexDirection: "row", alignItems: "center", gap: 8 }}> */}
          <Text style={theme.text}>
            📅 Vencimento em: {new Date(item.dataVencimento).toLocaleDateString("pt-BR", {
              day: "2-digit",
              month: "2-digit",
              year: "numeric",
            })}
          </Text>
          {/* <FontAwesome5
            name={
              item.statusPagamento === "Pago"
                ? "check-circle"
                : item.statusPagamento === "Pendente"
                  ? "exclamation-circle"
                  : "hourglass-half"
            }
            size={18}
            color={
              item.statusPagamento === "Pago"
                ? "green"
                : item.statusPagamento === "Pendente"
                  ? "red"
                  : "orange"
            }
          />
          <Text style={[theme.cardHighlight, { color: item.statusPagamento === "Pago" ? "green" : item.statusPagamento === "Pendente" ? "red" : "orange" }]}>
            {item.statusPagamento.toUpperCase()}
          </Text> */}
          {/* </View> */}
          <View style={{ flexDirection: "row", alignItems: "center", gap: 8 }}>
            <Text style={theme.text}>
              💸 Pagamento em:
            </Text>
            {item.dataPagamento
              ? (
                <Text style={theme.text}>
                  {new Date(item.dataPagamento).toLocaleDateString("pt-BR", {
                    day: "2-digit",
                    month: "2-digit",
                    year: "numeric",
                  })}
                </Text>
              )
              : (
                <View style={{ flexDirection: "row", alignItems: "center", gap: 8 }}>
                  <Text style={[theme.text, { color: "red", fontWeight: "bold" }]}>
                    {item.statusPagamento}
                  </Text>
                  {(item.statusPagamento === "Pendente")
                    ? <FontAwesome5 name="exclamation-triangle" color={"red"} size={18} solid />
                    : <Octicons name="key-asterisk" color="red" size={18} solid />
                  }
                </View>
              )
            }
          </View>
        </View>

        <View style={[theme.cardContent, { marginBottom: 0, borderWidth: 0.5, borderColor: theme.primary, padding: 4 }]}>
          <Text style={[theme.cardTitle, { marginTop: 12 }]}>📈 Seu Resultado</Text>
          <Text style={[
            theme.text,
            { color: item.valorEnergiaRecebida >= item.valorCobrado ? "#006400" : "#8B0000" },
            { fontWeight: "bold" },
          ]}>
            💹 {item.valorEnergiaRecebida >= item.valorCobrado ? "Ganho" : "Perda"} de R$ {(item.valorEnergiaRecebida - item.valorCobrado).toFixed(2)} (
            {(
              (
                item.valorEnergiaRecebida >= item.valorCobrado
                  ? (1 - item.valorCobrado / item.valorEnergiaRecebida)
                  : (1 - item.valorEnergiaRecebida / item.valorCobrado)
              ) * 100
            ).toFixed(2)}%
            )
          </Text>
        </View>

        {!!item.observacoes && (
          <>
            <Text style={[theme.cardTitle, { marginTop: 12 }]}>📝 Observações:</Text>
            <Text style={theme.text}>{item.observacoes}</Text>
          </>
        )}
      </View>
    );
  };

  // NÃO ESTANDO LOGADO, RETORNA MENSAGEM DE AVISO
  if (!isLoggedIn) {
    console.log("MovimentacaoComCardScreen - Usuário não está logado, exibindo mensagem de aviso.");
    return (
      <View style={[theme.container, { justifyContent: "center", alignItems: "center" }]}>
        <Text style={theme.text}>Você precisa estar logado para visualizar as movimentações.</Text>
      </View>
    );
  }

  // ESTANDO LOGADO, RETORNA A LISTA DE MOVIMENTAÇÕES
  console.log("MovimentacaoComCardScreen - Renderizando lista de movimentações mensais");
  return (
    <View style={[theme.container, { padding: 16 }]}>
      {/* <Text style={theme.title}>Movimentações Mensais</Text> */}
      <FlatList
        data={movimentacoes}
        keyExtractor={(item) => item.id.toString()}
        renderItem={renderItem}
        contentContainerStyle={{ paddingBottom: 32 }}
      />
    </View>
  );
}
