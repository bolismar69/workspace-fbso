// src/context/DatabaseContext.tsx
import React, { createContext, useContext, useState, useEffect, ReactNode } from "react";
import * as SQLite from "expo-sqlite";
import { initializeDatabase } from "./../services/database/initializeSQLiteDatabase";

// Tipagem para o contexto
interface DatabaseContextType {
  isDatabaseConnected: boolean;
  dbInstance: SQLite.SQLiteDatabase | null;
  initializeDatabaseConnection: () => Promise<void>;
  finalizeDatabaseConnection: () => Promise<void>;
  getDatabaseConnection: () => SQLite.SQLiteDatabase | null;
}

const DatabaseContext = createContext<DatabaseContextType | undefined>(undefined);

// Props para o DatabaseProvider
interface DatabaseProviderProps {
  children: ReactNode;
  autoInitialize?: boolean; // Define se a conexão deve ser inicializada automaticamente
}

// Provider
export const DatabaseProvider = ({ children, autoInitialize = false }: DatabaseProviderProps) => {
  console.log("📦 DatabaseContext - DatabaseProvider - iniciado.");
  const [isDatabaseConnected, setIsDatabaseConnected] = useState(false);
  const [dbInstance, setDbInstance] = useState<SQLite.SQLiteDatabase | null>(null);

  const initializeDatabaseConnection = async () => {
    console.log("🔗 DatabaseContext - initializeDatabaseConnection - Inicializando conexão com o banco SQLite...");
    if (!isDatabaseConnected) {
      try {
        const db = await initializeDatabase(); // Cria e retorna a instância
        setDbInstance(db);
        setIsDatabaseConnected(true);
        console.log("✅ DatabaseContext - initializeDatabaseConnection - Conexão com o banco SQLite inicializada.");
      } catch (error) {
        console.error("❌ DatabaseContext - initializeDatabaseConnection - Erro ao inicializar o banco:", error);
      }
    } else {
      console.log("🔗 DatabaseContext - initializeDatabaseConnection - Conexão já está ativa.");
    }
  };

  const finalizeDatabaseConnection = async () => {
    console.log("🔗 DatabaseContext - finalizeDatabaseConnection - Finalizando conexão com o banco SQLite...");
    if (dbInstance) {
      try {
        dbInstance.closeAsync(); // ⬅️ Se estiver usando expo-sqlite/next
        setIsDatabaseConnected(false);
        setDbInstance(null);
        console.log("📴 DatabaseContext - finalizeDatabaseConnection - Conexão com banco encerrada.");
      } catch (error) {
        console.error("DatabaseContext - finalizeDatabaseConnection - Erro ao fechar o banco:", error);
      }
    }
  };

  const getDatabaseConnection = () => {
    console.log("🔗 DatabaseContext - getDatabaseConnection - Obtendo conexão com o banco SQLite...");
    return dbInstance;
  };

  // Inicializa automaticamente a conexão se autoInitialize for true
  useEffect(() => {
    if (autoInitialize) {
      console.log("📦 DatabaseContext - autoInitialize está habilitado, inicializando conexão...");
      initializeDatabaseConnection();
    }
  }, [autoInitialize]);

  return (
    <DatabaseContext.Provider
      value={{
        isDatabaseConnected,
        dbInstance,
        initializeDatabaseConnection,
        finalizeDatabaseConnection,
        getDatabaseConnection,
      }}
    >
      {children}
    </DatabaseContext.Provider>
  );
};

// Hook customizado para consumir o contexto
export const useDatabase = (): DatabaseContextType => {
  console.log("🔗 DatabaseContext - useDatabase - Obtendo o contexto do banco de dados...");
  const context = useContext(DatabaseContext);
  if (!context) {
    throw new Error("DatabaseContext - useDatabase deve ser usado dentro do DatabaseProvider");
  }
  return context;
};
