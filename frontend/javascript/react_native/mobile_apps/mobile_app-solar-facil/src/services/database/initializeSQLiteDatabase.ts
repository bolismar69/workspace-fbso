// src/services/database/initializeSQLiteDatabase.ts
import * as SQLite from "expo-sqlite";

let dbInstance: SQLite.SQLiteDatabase | null = null;

const initializeSQLiteDatabase = async (): Promise<SQLite.SQLiteDatabase> => {
  if (dbInstance) return dbInstance; // já inicializado

  dbInstance = SQLite.openDatabaseSync("solarfacil.db");

  try {
    await dbInstance.withExclusiveTransactionAsync(async (tx) => {
      // Tabela de associados
      await tx.execAsync(`
        CREATE TABLE IF NOT EXISTS associados (
          id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL,
          dataCadastro TEXT NOT NULL,
          dataAtualizacao TEXT NOT NULL,
          senha TEXT NOT NULL,
          status TEXT CHECK(status IN ('Em cadastro', 'Ativo', 'Inativo', 'Bloqueado', 'Encerrado')) NOT NULL,
          tipoAssociado TEXT CHECK(tipoAssociado IN ('Fornecedor', 'Beneficiado', 'Hibrido')) NOT NULL,
          
          tipoPessoa TEXT CHECK(tipoPessoa IN ('Pessoa Física', 'Pessoa Jurídica')) NOT NULL,
          cpf_cnpj TEXT UNIQUE NOT NULL,
          nome TEXT NOT NULL,
          email TEXT NOT NULL,
          telefone TEXT NOT NULL,

          cep TEXT,
          endereco TEXT,
          numero TEXT,
          bairro TEXT,
          cidade TEXT,
          estado TEXT,
          complemento TEXT,

          aceitaTermos TEXT CHECK(aceitaTermos IN ('Sim', 'Não')),
          observacoes TEXT,

          dataNascimento TEXT,
          nomeSocial TEXT,

          dataAbertura TEXT,
          razaoSocial TEXT,
          nomeFantasia TEXT,

          nomeConcessionaria TEXT,
          consumoMedio TEXT,
          planoDesejado TEXT,

          potenciaInstalada TEXT,
          disponibilidade TEXT,
          tipoConexao TEXT
        );
      `);

      // Tabela de movimentações
      await tx.execAsync(`
        CREATE TABLE IF NOT EXISTS movimentacoes (
          id INTEGER PRIMARY KEY AUTOINCREMENT NOT NULL, 
          dataCadastro TEXT NOT NULL,
          dataAtualizacao TEXT NOT NULL,
          associadoId INTEGER NOT NULL,
          mes INTEGER,
          ano INTEGER,
          --
          valorTotal REAL,
          dataVencimento TEXT,
          dataPagamento TEXT,
          statusPagamento TEXT,
          observacoes TEXT,
          ---
          energiaRecebidaKwh REAL,
          valorEnergiaRecebida REAL,
          tarifaUnitariaKwh REAL,
          valorCobrado REAL,
          valorEconomizado REAL,
          percentualEconomizado REAL
        );
      `);
    });

    console.log("✅ Banco SQLite e tabelas criados com sucesso");
  } catch (error) {
    console.error("❌ Erro ao criar tabelas no banco SQLite:", error);
    throw error;
  }
  return dbInstance;
};

// Retorna a conexão ativa do banco
// Retorna conexão já inicializada (ou dispara erro claro)
const getDatabaseConnection = (): SQLite.SQLiteDatabase => {
  if (!dbInstance) {
    throw new Error(
      "Banco de dados não inicializado. Execute 'initializeDatabase()' primeiro."
    );
  }
  return dbInstance;
};

export {
  initializeSQLiteDatabase as initializeDatabase,
  getDatabaseConnection,
};
