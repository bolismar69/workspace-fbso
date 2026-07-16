import React, { useState } from "react";
import { Text, TextInput, TextInputProps, View, TouchableOpacity } from "react-native";
import { Picker } from "@react-native-picker/picker";
import { Ionicons } from "@expo/vector-icons";
import { useAppTheme } from "@/context/AppThemeContext";

interface InputTextProps extends TextInputProps {
  label?: string;
  error?: string;
  isButton?: boolean;
  onPress?: () => void;
  options?: { label: string; value: string }[];
  value?: string;
  onChangeText?: (value: string) => void;
  type?: "text" | "password" | "email" | "number";
  editable?: boolean;
}

export const InputText: React.FC<InputTextProps> = ({
  label,
  error,
  isButton,
  options,
  value,
  onChangeText,
  type = "text",
  editable = true, // Default value for editable
  ...props
}) => {
  const { theme } = useAppTheme();
  const [secureText, setSecureText] = useState(type === "password");

  // Para campos SELECT
  if (options) {
    return (
      <View style={{ marginBottom: 16 }}>
        {label && <Text style={theme.inputLabel}>{label}</Text>}
        <View style={[theme.pickerContainer, error && theme.pickerError]}>
          <Picker
            selectedValue={value}
            onValueChange={(itemValue) => onChangeText?.(itemValue)}
            enabled={editable}
          >
            {options.map((option) => (
              <Picker.Item key={option.value} label={option.label} value={option.value} />
            ))}
          </Picker>
        </View>
        {error && <Text style={theme.inputError}>{error}</Text>}
      </View>
    );
  }

  return (
    <View style={{ marginBottom: 16 }}>
      {label && <Text style={theme.inputLabel}>{label}</Text>}
      <View style={{ position: "relative" }}>
        <TextInput
          style={[
            theme.input,
            { paddingRight: type === "password" ? 40 : 12 },
            error && { borderColor: "red", borderWidth: 1 },
          ]}
          placeholderTextColor={theme.placeholder.color}
          value={value}
          onChangeText={onChangeText}
          secureTextEntry={secureText}
          editable={editable}
          {...props}
        />

        {type === "password" && (
          <TouchableOpacity
            style={{ position: "absolute", right: 10, top: 12 }}
            onPress={() => setSecureText((prev) => !prev)}
          >
            <Ionicons
              name={secureText ? "eye-off" : "eye"}
              size={20}
              color={theme.placeholder.color}
            />
          </TouchableOpacity>
        )}
      </View>

      {error && <Text style={theme.inputError}>{error}</Text>}
    </View>
  );
};
