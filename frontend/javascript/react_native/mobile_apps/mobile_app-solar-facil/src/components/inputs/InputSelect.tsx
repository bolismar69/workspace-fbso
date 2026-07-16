import React, { useState } from "react";
import { View, Text, TouchableOpacity, Modal, FlatList } from "react-native";
import { useAppTheme } from "@/context/AppThemeContext";

interface Option {
  label: string;
  value: string;
}

interface InputSelectProps {
  label: string;
  value: string;
  onChangeText: (val: string) => void;
  onBlur?: () => void;
  error?: string;
  options: Option[];
  placeholder?: string;
  editable?: boolean;
}

export function InputSelect({
  label,
  value,
  onChangeText,
  onBlur,
  error,
  options,
  placeholder,
  editable = true, // Default value for editable
}: InputSelectProps) {
  const { theme } = useAppTheme();
  const [modalVisible, setModalVisible] = useState(false);

  const selectedLabel = options.find((opt) => opt.value === value)?.label || placeholder || "Selecionar...";

  return (
    <View style={{ marginBottom: 16 }}>
      <Text style={[theme.label, { marginBottom: 4 }]}>{label}</Text>

      <TouchableOpacity
        onPress={() => setModalVisible(true && editable)}
        style={[
          theme.inputContainer,
          {
            padding: 12,
            borderWidth: 1,
            borderColor: error ? theme.inputError.color : "#ccc",
            borderRadius: 8,
          },
        ]}
      >
        <Text style={[theme.inputText, { color: value ? theme.text.color : "#999" }]}>
          {selectedLabel}
        </Text>
      </TouchableOpacity>

      <Modal
        visible={modalVisible}
        animationType="slide"
        transparent={true}
        onRequestClose={() => setModalVisible(false)}
      >
        <TouchableOpacity
          style={{
            flex: 1,
            backgroundColor: "rgba(0,0,0,0.3)",
            justifyContent: "flex-end",
          }}
          onPressOut={() => setModalVisible(false)}
          activeOpacity={1}
        >
          <View
            style={{
              backgroundColor: theme.screen.backgroundColor,
              padding: 20,
              borderTopLeftRadius: 16,
              borderTopRightRadius: 16,
              maxHeight: "50%",
            }}
          >
            <Text style={[theme.subtitle, { marginBottom: 12 }]}>Escolha uma opção</Text>
            <FlatList
              data={options}
              keyExtractor={(item) => item.value}
              renderItem={({ item }) => (
                <TouchableOpacity
                  onPress={() => {
                    onChangeText(item.value);
                    onBlur?.();
                    setModalVisible(false);
                  }}
                  style={{ paddingVertical: 12 }}
                >
                  <Text style={theme.text}>{item.label}</Text>
                </TouchableOpacity>
              )}
            />
          </View>
        </TouchableOpacity>
      </Modal>

      {error && <Text style={theme.inputError}>{error}</Text>}
    </View>
  );
}
