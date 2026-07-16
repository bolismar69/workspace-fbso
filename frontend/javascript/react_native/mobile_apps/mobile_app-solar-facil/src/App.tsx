import { AuthProvider } from "./context/AuthContext";
import { AppThemeProvider } from "./context/AppThemeContext";
import { Slot } from "expo-router";

export default function App() {
  return (
    <AppThemeProvider>
      <AuthProvider>
        <Slot />
      </AuthProvider>
    </AppThemeProvider>
  );
}
