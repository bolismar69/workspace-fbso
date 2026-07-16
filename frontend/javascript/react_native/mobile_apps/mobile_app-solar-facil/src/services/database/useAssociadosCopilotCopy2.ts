// src/services/database/useAssociadosCopilot.ts
import { AssociadoType } from "@/types/AssociadoType";
import {getDatabaseConnection} from "./initializeSQLiteDatabase"

export const useAssociadosCopilotCopy2 = () => {
  const insertAssociado = async (associado: AssociadoType) => {
    const db = getDatabaseConnection();

    const {
      dataCadastro, dataAtualizacao, senha, status, tipoAssociado,
      tipoPessoa, cpf_cnpj, nome, email, telefone,
      cep, endereco, numero, bairro, cidade, estado, complemento,
      aceitaTermos, observacoes,
      dataNascimento, nomeSocial,
      dataAbertura, razaoSocial, nomeFantasia,
      nomeConcessionaria, consumoMedio, planoDesejado,
      potenciaInstalada, disponibilidade, tipoConexao
    } = associado;

    await db.runAsync(
      `INSERT INTO associados (
        dataCadastro, dataAtualizacao, senha, status, tipoAssociado,
        tipoPessoa, cpf_cnpj, nome, email, telefone,
        cep, endereco, numero, bairro, cidade, estado, complemento,
        aceitaTermos, observacoes,
        dataNascimento, nomeSocial,
        dataAbertura, razaoSocial, nomeFantasia,
        nomeConcessionaria, consumoMedio, planoDesejado,
        potenciaInstalada, disponibilidade, tipoConexao
      ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
      [
        dataCadastro, dataAtualizacao, senha, status, tipoAssociado,
        tipoPessoa, cpf_cnpj, nome, email, telefone,
        cep, endereco, numero, bairro, cidade, estado, complemento,
        aceitaTermos, observacoes,
        dataNascimento, nomeSocial,
        dataAbertura, razaoSocial, nomeFantasia,
        nomeConcessionaria, consumoMedio, planoDesejado,
        potenciaInstalada, disponibilidade, tipoConexao
      ]
    );
  };

  const getAssociadoPorCpfCnpj = async (cpfCnpj: string): Promise<AssociadoType | null> => {
    const db = getDatabaseConnection();
    const result = await db.getFirstAsync<AssociadoType>(
      `SELECT * FROM associados WHERE cpf_cnpj = ?`,
      [cpfCnpj]
    );
    return result ?? null;
  };

  const getAssociadoPorCpfCnpjSenha = async (cpfCnpj: string, senha: string): Promise<AssociadoType | null> => {
    const db = getDatabaseConnection();
    const result = await db.getFirstAsync<AssociadoType>(
      `SELECT * FROM associados WHERE cpf_cnpj = ? AND senha = ?`,
      [cpfCnpj, senha]
    );
    return result ?? null;
  };

  const listarAssociados = async (): Promise<AssociadoType[]> => {
    const db = getDatabaseConnection();
    const result = await db.getAllAsync<AssociadoType>(`SELECT * FROM associados`);
    return result ?? [];
  };

  const atualizarAssociado = async (associado: AssociadoType) => {
    const db = getDatabaseConnection();
    await db.runAsync(
      `UPDATE associados SET
        dataAtualizacao = ?, senha = ?, status = ?, tipoAssociado = ?,
        tipoPessoa = ?, nome = ?, email = ?, telefone = ?,
        cep = ?, endereco = ?, numero = ?, bairro = ?, cidade = ?, estado = ?, complemento = ?,
        aceitaTermos = ?, observacoes = ?, dataNascimento = ?, nomeSocial = ?,
        dataAbertura = ?, razaoSocial = ?, nomeFantasia = ?, nomeConcessionaria = ?,
        consumoMedio = ?, planoDesejado = ?, potenciaInstalada = ?, disponibilidade = ?, tipoConexao = ?
      WHERE cpf_cnpj = ?`,
      [
        associado.dataAtualizacao, associado.senha, associado.status, associado.tipoAssociado,
        associado.tipoPessoa, associado.nome, associado.email, associado.telefone,
        associado.cep, associado.endereco, associado.numero, associado.bairro, associado.cidade, associado.estado, associado.complemento,
        associado.aceitaTermos, associado.observacoes, associado.dataNascimento, associado.nomeSocial,
        associado.dataAbertura, associado.razaoSocial, associado.nomeFantasia, associado.nomeConcessionaria,
        associado.consumoMedio, associado.planoDesejado, associado.potenciaInstalada, associado.disponibilidade, associado.tipoConexao,
        associado.cpf_cnpj
      ]
    );
  };

  return {
    insertAssociado,
    listarAssociados,
    getAssociadoPorCpfCnpj,
    getAssociadoPorCpfCnpjSenha,
    atualizarAssociado,
  };
};
