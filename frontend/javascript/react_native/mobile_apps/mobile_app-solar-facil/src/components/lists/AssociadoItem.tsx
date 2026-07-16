import React from "react";
import { View, Text, TouchableOpacity } from "react-native";
import { Ionicons } from "@expo/vector-icons";
import { useRouter } from "expo-router";
import { useAppTheme } from "@/context/AppThemeContext";

interface AssociadoItemProps {
  id: string;
  nome: string;
  tipo: "associado" | "beneficiado" | "fornecedor";
  onDelete?: (id: string) => void;
  onEdit?: (id: string) => void;
}

export function AssociadoItem({ id, nome, tipo, onDelete, onEdit }: AssociadoItemProps) {
  const { theme } = useAppTheme();
  const router = useRouter();

  return (
    <View
      style={{
        backgroundColor: theme.card.backgroundColor,
        padding: 12,
        borderRadius: 10,
        flexDirection: "row",
        alignItems: "center",
        justifyContent: "space-between",
        marginBottom: 10,
      }}
    >
      <View>
        <Text style={[theme.subtitle, { marginBottom: 4 }]}>{nome}</Text>
        <Text style={[theme.text, { fontSize: 12, color: "#6B7280" }]}>
          Tipo: {tipo}
        </Text>
      </View>

      <View style={{ flexDirection: "row", gap: 16 }}>
        <TouchableOpacity onPress={() => router.push(`/associado/dadoscadastrais?id=${id}`)}>
          <Ionicons name="eye-outline" size={24} color="#2563EB" />
        </TouchableOpacity>

        <TouchableOpacity onPress={() => onEdit?.(id)}>
          <Ionicons name="create-outline" size={24} color="#FBBF24" />
        </TouchableOpacity>

        <TouchableOpacity onPress={() => onDelete?.(id)}>
          <Ionicons name="trash-outline" size={24} color="#EF4444" />
        </TouchableOpacity>
      </View>
    </View>
  );
}
