import { AssociadoType } from "@/types/AssociadoType";
import React, { createContext, useContext, useState, ReactNode } from "react";

type AuthContextType = {
  isLoggedIn: boolean;
  login: (_userID: number, _userName: string, _associado?: AssociadoType) => void;
  logout: () => void;
  updatelogin: (_associado: AssociadoType) => void;
  userID: number | null;
  userName?: string | null;
  associado?: AssociadoType; // Placeholder for AssociadoType
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export const AuthProvider = ({ children }: { children: ReactNode }) => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [userID, setUserID] = useState<number | null>(null);
  const [userName, setUserName] = useState<string | null>(null);
  const [associado, setAssociado] = useState<AssociadoType>(); // Placeholder for AssociadoType

  const login = (_userID: number,
    _userName?: string,
    _associado?: AssociadoType
  ) => {
    console.log("login");
    setIsLoggedIn(true);
    setUserID(_userID || null);
    setUserName(_userName || _userID.toString() || null);
    setAssociado(_associado);
  };

  const logout = () => {
    console.log("logout");
    setIsLoggedIn(false);
    setUserID(null);
    setUserName(null);
    setAssociado(undefined); // Reset associado on logout
  }

  // atualizado informações do associado
  const updatelogin = (
    _associado: AssociadoType
  ) => {
    console.log("Atualizando associado:", _associado);
    // Aqui você pode implementar a lógica para atualizar o associado no backend
    // Por exemplo, fazer uma chamada API para atualizar os dados do associado
    setAssociado(_associado);
  };

  return (
    <AuthContext.Provider value={{ isLoggedIn, login, logout, userID, userName, associado, updatelogin }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) throw new Error("useAuth must be used within an AuthProvider");
  return context;
};
