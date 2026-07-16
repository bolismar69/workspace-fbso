import React from "react";
import { View, Text } from "react-native";
import { Ionicons } from "@expo/vector-icons";

interface CardIconeAmareloProps {
  iconName: keyof typeof Ionicons.glyphMap; // Nome do ícone do Ionicons
  iconColor: string; // Cor do ícone
  iconBackground: string; // Cor de fundo do ícone
  title: string; // Título do card
  description: string; // Descrição do card
  theme: any; // Tema para estilos
}

const CardIconeAmarelo: React.FC<CardIconeAmareloProps> = ({
  iconName,
  iconColor,
  iconBackground,
  title,
  description,
  theme,
}) => {
  return (
    <View style={theme.card}>
      <View style={[{ flexDirection: "row", alignItems: "center" }]}>
        <View style={[theme.iconContainer, { backgroundColor: iconBackground, marginRight: 8 }]}>
          <Ionicons name={iconName} size={56} color={iconColor} />
        </View>
        <Text style={[theme.highlightText, { marginLeft: 4 }]}>{title}</Text>
      </View>
      <Text style={[theme.text, { marginTop: 8 }]}>{description}</Text>
    </View>
  );
};

export default CardIconeAmarelo;
