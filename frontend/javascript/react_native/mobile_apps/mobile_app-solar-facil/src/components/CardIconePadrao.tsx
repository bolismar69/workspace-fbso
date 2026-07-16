import React from "react";
import { View, Text, ColorValue } from "react-native";
import { MaterialCommunityIcons } from "@expo/vector-icons";

interface CardIconePadraoProps {
  iconName: keyof typeof MaterialCommunityIcons.glyphMap; // Nome do ícone
  iconColor: ColorValue | undefined; // Cor do ícone (ajustado para aceitar valores dinâmicos)
  iconBackground: string; // Cor de fundo do ícone
  title: string; // Título do card
  description: string; // Descrição do card
  theme: any; // Tema para estilos
}

const CardIconePadrao: React.FC<CardIconePadraoProps> = ({
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
          <MaterialCommunityIcons name={iconName} size={24} color={iconColor} />
        </View>
        <Text style={theme.highlightText}>{title}</Text>
      </View>
      <Text style={[theme.text, { marginTop: 8 }]}>{description}</Text>
    </View>
  );
};

export default CardIconePadrao;
