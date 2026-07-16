import { Text, TouchableOpacity } from "react-native";
import { useAppTheme } from "../context/AppThemeContext";

type Props = {
  onPress: () => void;
};

export function ThemedButton({ onPress }: Props) {
  const { theme } = useAppTheme();

  return (
    <TouchableOpacity style={[theme.button]} onPress={onPress}>
      <Text style={theme.buttonText}>Alternar Tema</Text>
    </TouchableOpacity>
  );
}

export default ThemedButton;
