import { AppThemeStyles } from "@/types/ThemeTypes";

export const lightTheme: AppThemeStyles = {
  primary: "#1E5631",
  secondary: "#A5C9CA",
  links: "#1E90FF",
  backgroundColor: "#ffffbf",
  textColor: "#000",

  basicView: {
    backgroundColor: "#fff",
  },
  basicText: {
    color: "#000",
  },

  container: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: "#fff",
    padding: 16,
  },

  title: {
    fontSize: 24,
    fontWeight: "bold",
    color: "#1E5631",
    marginTop: 16,
    marginBottom: 16,
    textAlign: "center",
    backgroundColor: "#ffffbf", // Fundo amarelo claro para destaque
  },
  
  subtitle: {
    fontSize: 20,
    fontWeight: "bold",
    color: "#444",
    marginTop: 24,
    marginBottom: 10,
    textAlign: "center",
  },

  text: {
    fontSize: 16,
    color: "#000",
    marginBottom: 8,
    textAlign: "left",
    lineHeight: 24,
  },

  card: {
    backgroundColor: "#fff",
    borderRadius: 8,
    padding: 16,
    boxShadow: "0px 2px 4px rgba(0, 0, 0, 0.1)", // Substituído por boxShadow
    elevation: 2, // Mantido para compatibilidade com Android
    marginBottom: 16,
  },
  cardContent: {
    backgroundColor: "#ffffff",
    padding: 16,
    borderRadius: 12,
    marginVertical: 8,
    shadowColor: "#000", // Cor da sombra
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.3,
    shadowRadius: 4,
    elevation: 2,
  },
  cardTitle: {
    fontSize: 16,
    fontWeight: "600",
    marginBottom: 4,
    color: "#333333",
  },
  cardHighlight: {
    fontWeight: "700",
    fontSize: 14,
  },
  statusBackground: {
    pago: "#e0f7e9",
    atrasado: "#fbe9e7",
    pendente: "#fffde7"
  },

  imagePreview: {
    width: "100%",
    height: 200,
    borderRadius: 8,
    marginBottom: 16,
  },

  highlightText: {
    fontWeight: "bold",
    fontSize: 16,
    color: "#1E5631",
    marginBottom: 8,
    textAlign: "center",
  },

  button: {
    backgroundColor: "#1E5631", // Verde escuro
    paddingVertical: 10,
    paddingHorizontal: 24,
    borderRadius: 8,
    marginBottom: 8,
    alignSelf: "center",
  },
  buttonLow: {
    backgroundColor: "#1E5631", // Verde escuro
    // paddingVertical: 10,
    // paddingHorizontal: 24,
    borderRadius: 8,
    // marginBottom: 8,
    alignSelf: "center",
    height: 32, // Altura do botão
    width: "25%", // Largura total do botão
    justifyContent: "center", // Centraliza o texto verticalmente
    alignItems: "center", // Centraliza o texto horizontalmente
    shadowColor: "#000",
    shadowOffset: {
      width: 0,   // Deslocamento horizontal da sombra
      height: 2,  // Deslocamento vertical da sombra
    },
    shadowOpacity: 0.25, // Opacidade da sombra
    shadowRadius: 3.84, // Raio de desfoque da sombra
    elevation: 5, // Elevação para Android
  },
  
  secondaryButton: {
    backgroundColor: "#A5C9CA",
    paddingVertical: 10,
    paddingHorizontal: 24,
    borderRadius: 8,
    marginBottom: 8,
    alignSelf: "center",
  },

  buttonText: {
    color: "#fff",
    fontWeight: "bold",
    fontSize: 16,
    textAlign: "center",
  },
  
  iconContainer: {
    justifyContent: "center",
    alignItems: "center",
    borderRadius: 28, // Metade do tamanho do ícone para criar um círculo
    width: 56, // Largura igual ao tamanho do ícone
    height: 56, // Altura igual ao tamanho do ícone
  },

  inputText: {
    borderWidth: 1,
    borderColor: "#ccc",
    backgroundColor: "#fff",
    color: "#000",
    borderRadius: 8,
    paddingVertical: 10,
    paddingHorizontal: 12,
    fontSize: 16,
    marginBottom: 12,
  },
  inputContainer: {
    marginBottom: 16,
    width: "100%",
  },
  inputLabel: {
    fontSize: 16,
    color: "#444",
    marginBottom: 4,
    fontWeight: "bold",
  },
  inputError: {
    borderColor: "#ff0000", // Vermelho para bordas de erro
    borderWidth: 0.25,
    backgroundColor: "#e8e9eb", // cinza claro
    color: "#ff0000", // Vermelho para texto de erro
    borderRadius: 8,
    paddingVertical: 10,
    paddingHorizontal: 12,
    fontSize: 16,
    marginBottom: 12,
  },

  input: {
    borderWidth: 1,
    borderColor: "#ccc",
    borderRadius: 8,
    padding: 12,
    color: "#111",
    backgroundColor: "#fff",
  },
  label: {
    fontSize: 14,
    color: "#333",
    fontWeight: "500",
  },
  placeholder: {
    color: "#999",
  },

  tab: {
    backgroundColor: "#c3c3c3", //"#fff",
    activeColor: "#1E90FF",
    inactiveColor: "#888",
    indicatorColor: "#1E90FF",
  },
  screen: {
    backgroundColor: "#f0f0f0",
    flex: 1,
    padding: 16,
  },

  pickerContainer: {
    borderWidth: 1,
    borderColor: "#ccc",
    borderRadius: 8,
    overflow: "hidden",
  },
  pickerError: {
    borderColor: "red",
  },

  inputBorder: {
    borderWidth: 1,
    borderColor: "#ccc",
    borderRadius: 8,
    backgroundColor: "#1f2937", // cinza escuro
    color: "#f9fafb",
    paddingVertical: 10,
    paddingHorizontal: 12,
    fontSize: 16,
    marginBottom: 12,
  },

  safe: {
    flex: 1,
    backgroundColor: "#f0f0f0",
  },
  screenNoFlex: {
    backgroundColor: "#f0f0f0",
    padding: 16,
  },

  linkButton: {
    backgroundColor: "#1E90FF",
    paddingVertical: 10,
    paddingHorizontal: 24,
    borderRadius: 8,
    marginBottom: 8,
    alignSelf: "center",  
  },
  linkText: {
    color: "#fff",
    fontWeight: "bold",
    fontSize: 16,
    textAlign: "center",
  },

  // Estilos para o rodapé
  footer: {
    backgroundColor: "#f0f0f0", // Fundo cinza claro para o rodapé
    textColor: "#888888", // Cinza claro para o texto do rodapé
  },
  footerText: {
    // #7a7a7a#6c6c6c #5f5f5f #515151 #444444
    backgroundColor: "#f0f0f0", // Fundo cinza claro para o rodapé
    color: "#888888", // Cinza claro para o texto do rodapé
    fontSize: 14,
    textAlign: "center",
    marginTop: 0,
    padding: 0,
  },

  header: {
    backgroundColor: "#fff",
    textColor: "#000",
  },
};
