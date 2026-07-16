// /src/services/database/useAssociadosCopilot.ts
import { AssociadoType } from "@/types/AssociadoType";
import { seedMovimentacoes } from "./seedMovimentacoes";
import { useDatabase } from "@/context/DatabaseContext";

export async function useAssociadosCopilot() {
  console.log("=== INICIO ========================================================================");
  console.log("useAssociadosCopilot - Iniciando o hook useAssociadosCopilot...");

  const {
    getDatabaseConnection,
    isDatabaseConnected,
    initializeDatabaseConnection,
  } = useDatabase();

  async function ensureDbConnected() {
    if (!isDatabaseConnected) {
      await initializeDatabaseConnection();
    }
    const db = getDatabaseConnection();
    if (!db) throw new Error("Banco de dados nao conectado");
    return db;
  }

  async function insertRecord(associado: AssociadoType) {
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
      numberChanges: result?.changes ?? -1,
      rowID: result?.lastInsertRowId ?? -1,
    };
  }

  async function updateRecord(associado: AssociadoType) {
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
      numberChanges: result?.changes ?? -1,
      rowID: associado.id ?? -1,
    };
  }

  async function deleteRecord(id: number) {
    const db = await ensureDbConnected();
    const result = await db.runAsync(`DELETE FROM associados WHERE id = ?;`, [id]);
    return {
      numberChanges: result?.changes ?? -1,
      rowID: id,
    };
  }

  async function searchById(id: number): Promise<AssociadoType[]> {
    const db = await ensureDbConnected();
    const result = await db.getAllAsync(`SELECT * FROM associados WHERE id = ?;`, [id]);
    return result as AssociadoType[] ?? [];
  }

  async function searchByCpfCnpj(cpf_cnpj: string): Promise<AssociadoType[]> {
    const db = await ensureDbConnected();
    const result = await db.getAllAsync(`SELECT * FROM associados WHERE cpf_cnpj = ?;`, [cpf_cnpj]);
    return result as AssociadoType[] ?? [];
  }

  async function searchByCpfCnpjSenha(cpf_cnpj: string, senha: string): Promise<AssociadoType[]> {
    const db = await ensureDbConnected();
    const result = await db.getAllAsync(`SELECT * FROM associados WHERE cpf_cnpj = ? AND senha = ?;`, [cpf_cnpj, senha]);
    return result as AssociadoType[] ?? [];
  }

  async function searchAll(): Promise<AssociadoType[]> {
    const db = await ensureDbConnected();
    const result = await db.getAllAsync(`SELECT * FROM associados;`);
    return result as AssociadoType[] ?? [];
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
