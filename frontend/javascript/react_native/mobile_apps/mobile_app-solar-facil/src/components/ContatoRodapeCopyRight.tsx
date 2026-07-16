// src/components/ContatoRodape.tsx
import React from "react";
import { Text, View } from "react-native";
import { useAppTheme } from "@/context/AppThemeContext";

interface ContatoRodapeProps {
  label?: string; // Ensure 'label' is defined
}

export const ContatoRodapeCopyRight: React.FC<ContatoRodapeProps> = ({ label }) => {
  const { theme } = useAppTheme();

  return (
    // <View style={{ flex: 1 }}>
    <View style={theme.footer}>
      <Text style={[theme.footerText, { textAlign: "center", marginTop: 8, marginBottom: 8 }]}>
        {/* Use the label prop if provided, otherwise default to the static text */}
        {label || "© 2025 Solar Fácil - Todos os direitos reservados."}
      </Text>
    </View >
    // </View>
  );
};
