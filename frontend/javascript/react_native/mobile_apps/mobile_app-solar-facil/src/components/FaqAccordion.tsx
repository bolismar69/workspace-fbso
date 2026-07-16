// src/components/FaqAccordion.tsx
import React, { useState } from "react";
import { View, Text, TouchableOpacity, StyleSheet, LayoutAnimation, Platform, UIManager } from "react-native";
import { Ionicons } from "@expo/vector-icons";
import { useAppTheme } from "@/context/AppThemeContext";

if (Platform.OS === "android" && UIManager.setLayoutAnimationEnabledExperimental) {
  UIManager.setLayoutAnimationEnabledExperimental(true);
}

type FaqAccordionProps = {
  pergunta: string;
  resposta: string;
};

export const FaqAccordion = ({ pergunta, resposta }: FaqAccordionProps) => {
  const { theme } = useAppTheme();
  const [expanded, setExpanded] = useState(false);

  const toggle = () => {
    LayoutAnimation.configureNext(LayoutAnimation.Presets.easeInEaseOut);
    setExpanded(!expanded);
  };

  return (
    <View style={[styles.container, { backgroundColor: theme.card.backgroundColor }]}>
      <TouchableOpacity onPress={toggle} style={styles.header}>
        <Text style={[theme.subtitle, { flex: 1 }]}>{pergunta}</Text>
        <Ionicons
          name={expanded ? "chevron-up-outline" : "chevron-down-outline"}
          size={20}
          color={theme.text.color}
        />
      </TouchableOpacity>
      {expanded && (
        <View style={styles.body}>
          <Text style={theme.text}>{resposta}</Text>
        </View>
      )}
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    borderRadius: 8,
    padding: 12,
    marginBottom: 12,
    shadowColor: "#000",
    shadowOpacity: 0.05,
    shadowRadius: 3,
    elevation: 1,
  },
  header: {
    flexDirection: "row",
    alignItems: "center",
  },
  body: {
    marginTop: 8,
  },
});
