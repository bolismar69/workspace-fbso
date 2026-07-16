// src/context/AppThemeContext.tsx
import { darkTheme } from "@/styles/darkTheme";
import { lightTheme } from "@/styles/lightTheme";
import { AppThemeStyles } from "@/types/ThemeTypes";
import React, { createContext, ReactNode, useContext, useEffect, useState } from "react";
import { useColorScheme } from "react-native";

const AppThemeContext = createContext<{
  theme: AppThemeStyles;
  toggleTheme: () => void;
} | null>(null);

export const AppThemeProvider = ({ children }: { children: ReactNode }) => {
  const systemTheme = useColorScheme(); // Detecta o tema do sistema operacional
  const [isDark, setIsDark] = useState(systemTheme === "dark"); // Define o tema inicial com base no sistema
  const theme = isDark ? darkTheme : lightTheme;

  const toggleTheme = () => {
    setIsDark((prev) => !prev);
  };

  useEffect(() => {
    // Atualiza o tema quando o sistema muda
    setIsDark(systemTheme === "dark");
  }, [systemTheme]);

  return (
    <AppThemeContext.Provider value={{ theme, toggleTheme }}>
      {children}
    </AppThemeContext.Provider>
  );
};

export const useAppTheme = () => {
  const context = useContext(AppThemeContext);
  if (!context) {
    throw new Error("useAppTheme must be used within an AppThemeProvider");
  }
  return context;
};
