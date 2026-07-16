// /src/services/database/initializeDatabaseCopilot.ts
import * as SQLite from "expo-sqlite";

let db: SQLite.SQLiteDatabase | null;

export async function initializeDatabaseCopilot(): Promise<SQLite.SQLiteDatabase | null> {
  console.log("=== INICIO ========================================================================");
  console.log("initializeDatabase - Iniciando o processo de inicialização do banco de dados...");

  try {
    if (!db) {
      db = await SQLite.openDatabaseAsync("solarfacil.db");
    }
    if (!db) {
      throw new Error("initializeDatabase - Erro ao abrir o banco de dados.");
    }
    console.log("initializeDatabase - Banco de dados aberto com sucesso.");

    // ****************************************************************************************************************
    // testar se a conexão com o banco de dados está funcionando
    console.log("initializeDatabase - Validando se a conexão com o banco de dados esta funcionando ...");
    try {
      await db.getAllAsync("SELECT 1").then(() => {
        console.log("initializeDatabase - Conexão com o banco de dados testada com sucesso.");
      });
    } catch (error) {
      console.log("=== INICIO ========================================================================");
      console.info("initializeDatabase - Erro ao testar o funcionamento da conexão com o banco de dados:", error);
      console.group("initializeDatabase - FORÇANDO fechar o banco de dados e reabri-lo...");

      // **************************************************************************************************************
      // aqui vamos forçar o fechamento do banco de dados para abri-lo novamente
      try {
        await db.closeAsync();
        console.info("initializeDatabase - Banco de dados fechado com sucesso.");
      } catch (error) {
        console.info("initializeDatabase - Erro ao fechar o banco de dados:", error);
        
      } finally {
        console.info("initializeDatabase - Setando db=null...");
        db = null;
      }
      console.info("initializeDatabase - Garantindo que o banco de dados foi/esta fechado com sucesso.");
      console.info("initializeDatabase - Setando db=null...");
      db = null;

      // **************************************************************************************************************
      // aqui vamos tentar abrir o banco de dados novamente
      console.info("initializeDatabase - FORÇANDO reabrir o banco de dados...");
      db = await SQLite.openDatabaseAsync("solarfacil.db");
      if (!db) {
        console.error("initializeDatabase - Erro ao tentar reabrir o banco de dados.");
        throw new Error("initializeDatabase - Erro ao tentar reabrir o banco de dados.");
      }
      console.info("initializeDatabase - Banco de dados reaberto com sucesso.");
      console.groupEnd();
      console.log("=== TERMINO ========================================================================");
    }

    // ****************************************************************************************************************
    // força dropar as tabelas
      // db?.execAsync(`DROP TABLE IF EXISTS associados;`);
      // console.info("initializeDatabase - DROP Tabela 'associados'.");
      // db?.execAsync(`DROP TABLE IF EXISTS movimentacoes;`);
      // console.info("initializeDatabase - DROP Tabela 'movimentacoes'.");

    // ****************************************************************************************************************
    // Verifica se as tabelas já existem antes de criá-las
    console.group("initializeDatabase - Verificando se as tabelas já existem...");
    const tables = await db.getAllAsync("SELECT name FROM sqlite_master WHERE type='table'");

    console.log("initializeDatabase - Tabelas encontradas:", tables);
    const tableNames = tables.map((table: any) => table.name);

    // se tablesNames não contiver as tabelas, cria as tabelas
    console.log("initializeDatabase - Verificando se a tabela 'associados' já existe...");
    if (!tableNames.includes("associados")) {
      console.info("initializeDatabase - Tabela 'associados' ira ser criada.");
      await new Promise<void>((resolve, reject) => {
        db?.execAsync(`
                        CREATE TABLE IF NOT EXISTS associados (
                          id INTEGER PRIMARY KEY NOT NULL,
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

                          cep TEXTL,
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
                        );`
          )
          .then(() => {
            console.log("initializeDatabase - Tabela 'associados' criada com sucesso.");
            resolve();
          })
          .catch((error) => {
            console.error("initializeDatabase - Erro ao criar tabela 'associados':", error);
            reject(error);
          });
      });
    } else {
      console.info("initializeDatabase - Tabela 'associados' já existe.");
    }

    console.log("initializeDatabase - Verificando se a tabela 'movimentacoes' já existe...");
    if (!tableNames.includes("movimentacoes")) {
      console.info("initializeDatabase - Tabela 'movimentacoes' ira ser criada.");
      await new Promise<void>((resolve, reject) => {
        db?.execAsync(`
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
            `)
          .then(() => {
            console.log("initializeDatabase - Tabela 'movimentacoes' criada com sucesso.");
            resolve();
          })
          .catch((error) => {
            console.error("initializeDatabase - Erro ao criar tabela 'movimentacoes':", error);
            reject(error);
          });
      });
    } else {
      console.info("initializeDatabase - Tabela 'movimentacoes' já existe.");
    }

    // console.log("initializeDatabase - Verificando se a tabela 'movimentacoes_detalhes' já existe...");
    // if (!tableNames.includes("movimentacoes_detalhes")) {
    //   console.info("initializeDatabase - Tabela 'movimentacoes_detalhes' ira ser criada.");
      // await new Promise<void>((resolve, reject) => {
      //   db?.execAsync(
      //       `CREATE TABLE IF NOT EXISTS movimentacoes_detalhes (
      //         id INTEGER PRIMARY KEY NOT NULL,
      //         dataCadastro TEXT NOT NULL,
      //         dataAtualizacao TEXT NOT NULL,
      //         movimentacaoId INTEGER NOT NULL,
      //         energiaRecebidaKwh REAL,
      //         valorEnergiaRecebida REAL,
      //         tarifaUnitariaKwh REAL,
      //         valorCobrado REAL,
      //         valorEconomizado REAL,
      //         percentualEconomizado REAL
      //       );`)
      //     .then(() => {
      //       console.log("initializeDatabase - Tabela 'movimentacoes_detalhes' criada com sucesso.");
      //       resolve();
      //     })
      //     .catch((error) => {
      //       console.error("initializeDatabase - Erro ao criar tabela 'movimentacoes_detalhes':", error);
      //       reject(error);
      //     });
      // });
    // } else {
    //   // db?.execAsync(`DROP TABLE IF EXISTS movimentacoes_detalhes;`);
    //   // console.info("initializeDatabase - DROP Tabela 'movimentacoes_detalhes'.");
    //   console.info("initializeDatabase - Tabela 'movimentacoes_detalhes' já existe.");
    // }
    console.groupEnd();

    console.log("initializeDatabase - Banco de dados inicializado com sucesso.");
    console.log("=== TERMINO =======================================================================");
    return db;
  } catch (error) {
    console.error("initializeDatabase - Erro ao inicializar o banco de dados:", error);
    console.log("=== TERMINO =======================================================================");
    throw error;
  }
}
