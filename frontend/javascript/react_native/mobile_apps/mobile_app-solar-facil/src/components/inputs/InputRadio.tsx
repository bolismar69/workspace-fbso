import React from "react";
import { View, Text, TouchableOpacity } from "react-native";
import { useAppTheme } from "@/context/AppThemeContext";

interface Option {
  label: string;
  value: string;
}

interface InputRadioProps {
  label: string;
  value: string;
  onChangeText: (val: string) => void;
  onBlur?: () => void;
  error?: string;
  options: Option[];
  editable?: boolean;
}

export function InputRadio({
  label,
  value,
  onChangeText,
  onBlur,
  error,
  options,
  editable = true, // Default value for editable
}: InputRadioProps) {
  const { theme } = useAppTheme();

  return (
    <View style={{ marginBottom: 16 }}>
      <Text style={[theme.label, { marginBottom: 8 }]}>{label}</Text>

      {options.map((option) => {
        const isSelected = value === option.value;

        if (!editable && !isSelected) {
          return null;
        }

        return (
          <TouchableOpacity
            key={option.value}
            onPress={() => {
              onChangeText(option.value);
              onBlur?.();
            }}
            style={{
              flexDirection: "row",
              alignItems: "center",
              marginBottom: 8,
            }}
          >
            <View
              style={{
                height: 20,
                width: 20,
                borderRadius: 10,
                borderWidth: 2,
                borderColor: isSelected ? theme.primary : "#aaa",
                alignItems: "center",
                justifyContent: "center",
                marginRight: 8,
              }}
            >
              {isSelected && (
                <View
                  style={{
                    height: 10,
                    width: 10,
                    borderRadius: 5,
                    backgroundColor: theme.primary,
                  }}
                />
              )}
            </View>
            <Text style={theme.text}>{option.label}</Text>
          </TouchableOpacity>
        );
      })}

      {error && <Text style={theme.inputError}>{error}</Text>}
    </View>
  );
}
