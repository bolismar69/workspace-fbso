import { openDatabaseSync, SQLiteDatabase } from "expo-sqlite";
// import { initializeDatabase } from "./initializeDatabase";
import {AssociadoType} from "../../types/AssociadoType";
import {getDatabaseConnection, initializeDatabase as InitDB } from "./initializeSQLiteDatabase";

export async function useAssociados() { 

  // inicializa o banco de dados
  await InitDB();

  // pega conexao com o banco de dados
  let database = await getDatabaseConnection();

  async function initialize() {
  try {
    // Inicia o processo de inicialização do banco de dados
    await initializeDatabase();
  } catch (error) {
    console.error("useAssociados - Erro ao inicializar o banco de dados:", error);
    throw error;
  }
  }

  // recebe a entidade AssociadoType para inserir na tabela de associados
  async function insertRecord( associado: AssociadoType ) {
    await initialize();

    if (!database) {
      console.error("useAssociados - Banco de dados não inicializado.");
      return;
    }
    database = openDatabaseSync('solarfacil.db');
    if (!database) {
      console.error("useAssociados - Erro ao abrir o banco de dados.");
      return;
    }

    const statement = database.prepareAsync(`
      INSERT INTO associados (id, nome, email, telefone, tipoPessoa, cpf_cnpj, senha, tipoAssociado, status, dataCadastro, dataAtualizacao)
      VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);
    `);
    const values = [
      associado.id,
      associado.nome,
      associado.email,
      associado.telefone,
      associado.tipoPessoa,
      associado.cpf_cnpj,
      associado.senha,
      associado.tipoAssociado,
      associado.status,
      associado.dataCadastro,
      associado.dataAtualizacao,
    ];
    try {
      const result =  await (await statement).executeAsync(values);
      console.log("useAssociados - Associado inserido com sucesso:", result);
    } catch (error) {
      console.log("useAssociados - Banco de dados fechado após erro.");
      throw error
    }
    finally {
      if(statement){
        // Fechar a declaração após a execução
        await (await statement).finalizeAsync();
        console.log("useAssociados - statement finalizado.");
      }
      if(database){
        // Fechar o banco de dados após a inserção
        await database.closeAsync();
        console.log("useAssociados - Banco de dados fechado após inserção.");
      }
    }
  }

  return{ insertRecord , initialize, database};
}
