// src/components/inputs/InputPasswordWithToggle.tsx
import React, { useState } from "react";
import { View, TextInput, Text, TouchableOpacity } from "react-native";
import { Ionicons } from "@expo/vector-icons";
import { useAppTheme } from "@/context/AppThemeContext";

interface InputPasswordProps {
  label?: string;
  placeholder?: string;
  value?: string;
  onChangeText?: (text: string) => void;
  error?: string;
  editable?: boolean; // Default value is true
}

export const InputPasswordWithToggle: React.FC<InputPasswordProps> = ({
  label,
  placeholder,
  value = "",
  onChangeText,
  error,
  editable = true, // Default value for editable
}) => {
  const [hidePassword, setHidePassword] = useState(true);
  const { theme } = useAppTheme();

  const getStrength = (password: string): { label: string; color: string } => {
    if (!password) return { label: "", color: "transparent" };
    if (password.length < 6) return { label: "Fraca", color: "#DC2626" }; // vermelho
    if (password.match(/[a-z]/) && password.match(/[A-Z]/) && password.match(/\d/))
      return { label: "Forte", color: "#16A34A" }; // verde
    return { label: "Média", color: "#F59E0B" }; // laranja
  };

  const strength = getStrength(value);

  return (
    <View style={{ marginBottom: 16 }}>
      {label && <Text style={theme.inputLabel}>{label}</Text>}

      <View
        style={[
          {
            flexDirection: "row",
            alignItems: "center",
            marginTop: 4,
            paddingRight: 12,
          },
          error && { borderColor: "red", borderWidth: 1 },
        ]}
      >
        <TextInput
          placeholder={placeholder}
          placeholderTextColor={theme.placeholder.color}
          secureTextEntry={hidePassword}
          style={[theme.input, { flex: 1, color: theme.text.color }]}
          value={value}
          onChangeText={onChangeText}
          editable={editable}
        />
        <TouchableOpacity onPress={() => setHidePassword(!hidePassword)}>
          <Ionicons
            name={hidePassword ? "eye-off" : "eye"}
            size={22}
            color={theme.text.color}
          />
        </TouchableOpacity>
      </View>

      {!!strength.label && (
        <Text style={{ color: strength.color, marginTop: 4 }}>
          Força da senha: {strength.label}
        </Text>
      )}

      {error && <Text style={theme.inputError}>{error}</Text>}
    </View>
  );
};
