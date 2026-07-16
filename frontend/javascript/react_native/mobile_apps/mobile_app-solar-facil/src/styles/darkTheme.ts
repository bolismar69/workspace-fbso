import { AppThemeStyles } from "@/types/ThemeTypes";

export const darkTheme: AppThemeStyles = {
  primary: "#A5C9CA", // Azul claro
  secondary: "#1E5631", // Verde escuro
  links: "#bb86fc", // Roxo para links
  backgroundColor: "#1e1e1e", // Fundo escuro -- "#121212"
  textColor: "#ffffff", // Texto branco

  basicView: {
    backgroundColor: "#121212",
  },
  basicText: {
    color: "#ffffff",
  },

  container: {
    flex: 1,
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: "#121212",
    padding: 16,
  },

  title: {
    fontSize: 24,
    fontWeight: "bold",
    color: "#ffffff",
    marginTop: 16,
    marginBottom: 16,
    textAlign: "center",
    backgroundColor: "#1e1e1e", // Fundo escuro para destaque
  },

  subtitle: {
    fontSize: 16,
    color: "#bbbbbb",
    marginBottom: 32,
    textAlign: "center",
  },

  text: {
    fontSize: 16,
    color: "#ffffff",
    marginBottom: 8,
    textAlign: "left",
    lineHeight: 24,
  },

  card: {
      backgroundColor: "#1e1e1e",
      borderRadius: 8,
      padding: 16,
      boxShadow: "0px 2px 4px rgba(0, 0, 0, 0.1)", // Substituído por boxShadow
      elevation: 2, // Mantido para compatibilidade com Android
      marginBottom: 16,
  },
  cardContent: {
    backgroundColor: "#1e1e1e",
    padding: 16,
    borderRadius: 12,
    marginVertical: 8,
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.3,
    shadowRadius: 4,
    elevation: 2,
  },
  cardTitle: {
    fontSize: 16,
    fontWeight: "600",
    marginBottom: 4,
    color: "#f0f0f0",
  },
  cardHighlight: {
    fontWeight: "700",
    fontSize: 14,
    color: "#ffd700", // amarelo-ouro para destaque
  },
  statusBackground: {
    pago: "#144d38",
    atrasado: "#4b1c1c",
    pendente: "#4a4a2e"
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
    color: "#A5C9CA",
    marginBottom: 8,
    textAlign: "center",
  },

  button: {
    backgroundColor: "#bb86fc",
    paddingVertical: 10,
    paddingHorizontal: 24,
    borderRadius: 8,
    marginBottom: 8,
    alignSelf: "center",
  },
  buttonLow: {
    backgroundColor: "#bb86fc",
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
    backgroundColor: "#1E5631",
    paddingVertical: 10,
    paddingHorizontal: 24,
    borderRadius: 8,
    marginBottom: 8,
    alignSelf: "center",
  },

  buttonText: {
    color: "#000000",
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
    borderColor: "#444",
    backgroundColor: "#1f2937", // cinza escuro
    color: "#f9fafb",
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
    color: "#f9fafb",
    marginBottom: 4,
    fontWeight: "bold",
  },
  inputError: {
    color: "#ff0000", // Vermelho para erros
    borderColor: "#ff0000", // Vermelho para bordas de erro
    borderWidth: 1,
    backgroundColor: "#1f2937", // cinza escuro
    borderRadius: 8,
    paddingVertical: 10,
    paddingHorizontal: 12,
    fontSize: 16,
    marginBottom: 12,
  },

  input: {
    borderWidth: 1,
    borderColor: "#555",
    borderRadius: 8,
    padding: 12,
    color: "#fff",
    backgroundColor: "#1f1f1f",
  },
  label: {
    fontSize: 14,
    color: "#eee",
    fontWeight: "500",
  },
  placeholder: {
    color: "#888",
  },

  tab: {
    backgroundColor: "#1e1e1e",
    activeColor: "#bb86fc", // Roxo para o texto ativo
    inactiveColor: "#888888", // Cinza para o texto inativo
    indicatorColor: "#bb86fc", // Roxo para o indicador
  },
  screen: {
    backgroundColor: "#121212",
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
    borderColor: "#444",
    backgroundColor: "#1f2937", // cinza escuro
    color: "#f9fafb",
    borderRadius: 8,
    paddingVertical: 10,
    paddingHorizontal: 12,
    fontSize: 16,
    marginBottom: 12,
  },

  safe: {
    flex: 1,
    backgroundColor: "#121212",
  },
  screenNoFlex: {
    backgroundColor: "#121212",
    padding: 16,
  },

  linkButton: {
    backgroundColor: "transparent",
    paddingVertical: 10,
    paddingHorizontal: 24,
    borderRadius: 8,
    marginBottom: 8,
    alignSelf: "center",
  },
  linkText: {
    color: "#bb86fc", // Roxo para links
    fontWeight: "bold",
    fontSize: 16,
    textAlign: "center",
  },

  footer: {
    backgroundColor: "#1e1e1e", // Fundo escuro para o rodapé
    textColor: "#888888", // Cinza claro para o texto do rodapé
  },
  footerText: {
    color: "#888888", // Cinza claro para o texto do rodapé
    fontSize: 14,
    textAlign: "center",
    marginTop: 16,
    padding: 0, // Removido padding para rodapé
  },

  header: {
    backgroundColor: "#333",
    textColor: "#fff",
  },
};
