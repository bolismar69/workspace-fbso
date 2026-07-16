// /src/services/database/useAssociadosCopilot.ts
import { AssociadoType } from "@/types/AssociadoType";
import { seedMovimentacoes } from "./seedMovimentacoes";
import { useDatabase } from "@/context/DatabaseContext";
import { DBResponse } from "@/types/DBResponse";

export function useAssociadosCopilot() {
  const {
    getDatabaseConnection,
    isDatabaseConnected,
    initializeDatabaseConnection,
    dbInstance,
  } = useDatabase();

  console.log("useAssociadosCopilot - Iniciando o hook useDatabase(): ", {
    isDatabaseConnected,
    dbInstance });

  async function ensureDbConnected() {
    console.log("🔗 useAssociadosCopilot - ensureDbConnected - Verificando conexão com o banco de dados...");
    if (!isDatabaseConnected) {
      console.log("🔗 useAssociadosCopilot - ensureDbConnected - Inicializando conexão com o banco de dados...");
      await initializeDatabaseConnection();
    } else {
      console.log("✅ useAssociadosCopilot - ensureDbConnected - Conexão já está ativa.");
    }
    const db = getDatabaseConnection();
    if (!db) throw new Error("❌ Banco de dados não conectado");
    return db;
  }

  async function insertRecord(associado: AssociadoType): Promise<DBResponse<{ numberChanges: number; rowID: number }>> {
    try {
      const db = await ensureDbConnected();

      associado.id = parseInt(associado.cpf_cnpj.replace(/\D/g, ""), 10);

      const result = await db.runAsync(
        `INSERT INTO associados (id, nome, email, telefone, tipoPessoa, cpf_cnpj, senha, tipoAssociado, status, dataCadastro, dataAtualizacao)
         VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`,
        [
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
        ]
      );

      await seedMovimentacoes(associado.id, db);

      return {
        success: true,
        data: {
          numberChanges: result?.changes ?? -1,
          rowID: result?.lastInsertRowId ?? -1,
        },
      };
    } catch (error) {
      console.error("Erro ao inserir associado:", error);
      return { success: false, error: "Erro ao inserir associado" };
    }
  }

  async function updateRecord(associado: AssociadoType): Promise<DBResponse<{ numberChanges: number; rowID: number }>> {
    try {
      const db = await ensureDbConnected();

      const result = await db.runAsync(
        `UPDATE associados SET 
          nome = ?, email = ?, telefone = ?, tipoPessoa = ?, cpf_cnpj = ?, senha = ?, tipoAssociado = ?, status = ?,
          dataCadastro = ?, dataAtualizacao = ?, cep = ?, endereco = ?, numero = ?, bairro = ?, cidade = ?, estado = ?,
          complemento = ?, aceitaTermos = ?, observacoes = ?, dataNascimento = ?, nomeSocial = ?, dataAbertura = ?,
          razaoSocial = ?, nomeFantasia = ?, nomeConcessionaria = ?, consumoMedio = ?, planoDesejado = ?,
          potenciaInstalada = ?, disponibilidade = ?, tipoConexao = ?
         WHERE id = ?;`,
        [
          associado.nome, associado.email, associado.telefone, associado.tipoPessoa, associado.cpf_cnpj, associado.senha,
          associado.tipoAssociado, associado.status, associado.dataCadastro, associado.dataAtualizacao, associado.cep,
          associado.endereco, associado.numero, associado.bairro, associado.cidade, associado.estado, associado.complemento,
          associado.aceitaTermos, associado.observacoes, associado.dataNascimento, associado.nomeSocial, associado.dataAbertura,
          associado.razaoSocial, associado.nomeFantasia, associado.nomeConcessionaria, associado.consumoMedio,
          associado.planoDesejado, associado.potenciaInstalada, associado.disponibilidade, associado.tipoConexao,
          associado.id
        ]
      );

      return {
        success: true,
        data: {
          numberChanges: result?.changes ?? -1,
          rowID: associado.id ?? -1,
        },
      };
    } catch (error) {
      console.error("Erro ao atualizar associado:", error);
      return { success: false, error: "Erro ao atualizar associado" };
    }
  }

  async function deleteRecord(id: number): Promise<DBResponse<{ numberChanges: number; rowID: number }>> {
    try {
      const db = await ensureDbConnected();
      const result = await db.runAsync(`DELETE FROM associados WHERE id = ?;`, [id]);
      return {
        success: true,
        data: {
          numberChanges: result?.changes ?? -1,
          rowID: id,
        },
      };
    } catch (error) {
      console.error("Erro ao deletar associado:", error);
      return { success: false, error: "Erro ao deletar associado" };
    }
  }

  async function searchById(id: number): Promise<DBResponse<AssociadoType[]>> {
    try {
      const db = await ensureDbConnected();
      const result = await db.getAllAsync(`SELECT * FROM associados WHERE id = ?;`, [id]);
      return { success: true, data: (result ?? []) as AssociadoType[] };
    } catch (error) {
      console.error("Erro ao buscar associado por ID:", error);
      return { success: false, error: "Erro ao buscar associado por ID" };
    }
  }

  async function searchByCpfCnpj(cpf_cnpj: string): Promise<DBResponse<AssociadoType[]>> {
    try {
      const db = await ensureDbConnected();
      const result = await db.getAllAsync(`SELECT * FROM associados WHERE cpf_cnpj = ?;`, [cpf_cnpj]);
      return { success: true, data: (result ?? []) as AssociadoType[] };
    } catch (error) {
      console.error("❌ Erro ao buscar associado por CPF/CNPJ:", error);
      return { success: false, error: "Erro ao buscar associado por CPF/CNPJ" };
    }
  }

  async function searchByCpfCnpjSenha(cpf_cnpj: string, senha: string): Promise<DBResponse<AssociadoType[]>> {
    console.log("useAssociadosCopilot - searchByCpfCnpjSenha - Buscando associado por CPF/CNPJ e Senha...", cpf_cnpj, senha);
    try {
      const db = await ensureDbConnected();

      console.log("🔍 useAssociadosCopilot - searchByCpfCnpjSenha - Executando consulta no banco de dados...");
      const result = await db.getAllAsync(
        `SELECT * FROM associados WHERE cpf_cnpj = ? AND senha = ?;`,
        [cpf_cnpj, senha]
      );
      console.log("✅ useAssociadosCopilot - searchByCpfCnpjSenha - Consulta executada com sucesso. Result:", result);
      return { success: true, data: (result ?? []) as AssociadoType[] };
    } catch (error) {
      console.error("❌ Erro ao buscar associado por CPF/CNPJ e Senha:", error);
      return { success: false, error: "Erro ao buscar associado com CPF/CNPJ e Senha" };
    }
  }

  async function searchAll(): Promise<DBResponse<AssociadoType[]>> {
    console.log("useAssociadosCopilot - searchAll - Buscando todos os associados...");
    try {
      const db = await ensureDbConnected();
      const result = await db.getAllAsync(`SELECT * FROM associados;`);
      return { success: true, data: (result ?? []) as AssociadoType[] };
    } catch (error) {
      console.error("useAssociadosCopilot - searchAll - Erro ao buscar todos os associados:", error);
      return { success: false, error: "Erro ao buscar todos os associados" };
    }
  }

  return {
    insertRecord,
    updateRecord,
    deleteRecord,
    searchById,
    searchByCpfCnpj,
    searchByCpfCnpjSenha,
    searchAll,
  };
}
