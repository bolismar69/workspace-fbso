import { openDatabaseSync, SQLiteDatabase } from "expo-sqlite";

export async function initializeDatabase() {

  let db: SQLiteDatabase | undefined;

  console.log("=== INICIO ========================================================================");
  console.log("initializeDatabase - Iniciando o processo de inicialização do banco de dados...");
  try {

    // const db = SQLite.openDatabase('solarfacil.db');
    db = openDatabaseSync('solarfacil.db');

    // Verifica se o banco de dados já está aberto
    if (!db) {
      console.error("initializeDatabase - Erro ao abrir o banco de dados.");
      return;
    }

    // Cria as tabelas necessárias
    // await createAssociadoTable(db);
    // await createMovimentacaoTable(db);
    await db.execAsync(    
        `CREATE TABLE IF NOT EXISTS associados (
          id TEXT PRIMARY KEY NOT NULL,
          nome TEXT,
          email TEXT,
          telefone TEXT,
          tipoPessoa TEXT,
          cpf_cnpj TEXT UNIQUE,
          senha TEXT,
          tipoAssociado TEXT,
          status TEXT,
          dataCadastro TEXT,
          dataAtualizacao TEXT
        );
        CREATE TABLE IF NOT EXISTS movimentacoes (
          id TEXT PRIMARY KEY NOT NULL,
          associadoId TEXT,
          mes INTEGER,
          ano INTEGER,
          valorTotal REAL,
          dataCadastro TEXT
        );
        CREATE TABLE IF NOT EXISTS movimentacoes_detalhes (
          id TEXT PRIMARY KEY NOT NULL,
          movimentacaoId TEXT,
          energiaRecebidaKwh REAL,
          valorEnergiaRecebida REAL,
          tarifaUnitariaKwh REAL,
          valorCobrado REAL,
          dataVencimento TEXT,
          dataPagamento TEXT,
          statusPagamento TEXT,
          observacoes TEXT,
          criadoEm TEXT,
          atualizadoEm TEXT
        );`
      ).then(() => {
        console.log("initializeDatabase - Tabelas criadas com sucesso.");
      }).catch((error) => {
        console.error("initializeDatabase - Erro ao criar tabelas:", error);
        throw error;
      }).finally(() => {
        console.log("initializeDatabase - Finalizando o processo de inicialização do banco de dados e criação das tabelas.");
      });

    console.log("initializeDatabase - Banco de dados inicializado com sucesso.");
  } catch (error) {
    console.error("initializeDatabase - Erro ao inicializar o banco de dados:", error);
  }
  console.log("=== TERMINO =======================================================================");
  return;
}
