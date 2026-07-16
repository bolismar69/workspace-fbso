import React, { useState } from "react";
import {
  View,
  Text,
  TouchableOpacity,
  Platform,
  Modal,
  Pressable,
} from "react-native";
import DateTimePicker from "@react-native-community/datetimepicker";
import { useAppTheme } from "@/context/AppThemeContext";

interface InputDateProps {
  label: string;
  value: Date | undefined;
  onChangeText: (val: Date) => void;
  onBlur?: () => void;
  error?: string;
  editable?: boolean;
}

export function InputDate({
  label,
  value,
  onChangeText,
  onBlur,
  error,
  editable = true, // Default value for editable
}: InputDateProps) {
  const { theme } = useAppTheme();
  const [showPicker, setShowPicker] = useState(false);
  const [selectedDate, setSelectedDate] = useState<Date | undefined>(value); // Estado para armazenar a data selecionada

  const handleDateChange = (event: any, newDate?: Date) => {
    setShowPicker(false);
    if (newDate) {
      setSelectedDate(newDate); // Atualiza o estado local com a data selecionada
      onChangeText(newDate); // Chama a função de callback para atualizar o estado externo
    }
    onBlur?.();
  };

  const formattedDate =
    selectedDate instanceof Date
      ? selectedDate.toLocaleDateString("pt-BR")
      : "Selecionar data";
  console.log("InputDate - Selected Date:", formattedDate); // Log para verificar a data selecionada

  return (
    <View style={{ marginBottom: 16 }}>
      <Text style={[theme.label, { marginBottom: 4 }]}>{label}</Text>

      <TouchableOpacity
        style={[
          theme.inputContainer,
          { paddingVertical: 12, paddingHorizontal: 10 },
        ]}
        onPress={() => setShowPicker(true)}
      >
        <Text style={theme.inputText}>{formattedDate}</Text>
      </TouchableOpacity>

      {error && <Text style={theme.inputError}>{error}</Text>}

      {showPicker && Platform.OS === "ios" ? (
        <Modal transparent animationType="fade">
          <Pressable
            style={{
              flex: 1,
              justifyContent: "center",
              alignItems: "center",
              backgroundColor: "rgba(0,0,0,0.5)",
            }}
            onPress={() => setShowPicker(false)}
          >
            <View
              style={{
                backgroundColor: theme.card.backgroundColor,
                padding: 20,
                borderRadius: 10,
              }}
            >
              <DateTimePicker
                value={selectedDate instanceof Date ? selectedDate : new Date()} // Usa o estado local
                mode="date"
                display="spinner"
                onChange={handleDateChange}
                disabled={!editable}
              />
            </View>
          </Pressable>
        </Modal>
      ) : (
        showPicker && (
          <DateTimePicker
            value={selectedDate instanceof Date ? selectedDate : new Date()} // Usa o estado local
            mode="date"
            display="default"
            onChange={handleDateChange}
            disabled={!editable}
          />
        )
      )}
    </View>
  );
}
