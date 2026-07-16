// /src/types/ThemeTypes.ts
import { ImageStyle, TextStyle, ViewStyle } from "react-native";

export interface AppThemeStyles {
  header: {
    backgroundColor: string;
    textColor: string;
  }
  footer: {
    backgroundColor: string;
    textColor: string;
  },
  footerText: TextStyle
  linkText: TextStyle;
  linkButton: ViewStyle;
  
  safe: ViewStyle;

  inputBorder: any;
  primary: string;
  secondary: string;
  links: string;
  backgroundColor: string;
  textColor: string;

  basicView: ViewStyle;
  basicText: TextStyle;

  container: ViewStyle;
  title: TextStyle;
  subtitle: TextStyle;
  text: TextStyle;
  highlightText: TextStyle;

  button: ViewStyle;
  buttonLow: ViewStyle;
  secondaryButton: ViewStyle;
  buttonText: TextStyle;

  card: ViewStyle;
  cardContent: ViewStyle;
  cardTitle: TextStyle;
  cardHighlight: TextStyle;
  statusBackground: {
    pago: string;
    atrasado: string;
    pendente: string;
  };
  
  imagePreview: ImageStyle;

  iconContainer: ViewStyle;

  inputText: TextStyle;
  inputContainer: ViewStyle;
  inputLabel: TextStyle;
  inputError: TextStyle;

  placeholder: TextStyle;
  label: TextStyle;
  input: TextStyle;

  tab: {
    backgroundColor: string;
    activeColor: string;  
    inactiveColor: string;
    indicatorColor: string;
  };
  screen: {
    backgroundColor: string;
    flex: number;
    padding: number;
  };
  screenNoFlex: {
    backgroundColor: string;
    padding: number;
  };

  pickerContainer: ViewStyle;
  pickerError: ViewStyle;

}
