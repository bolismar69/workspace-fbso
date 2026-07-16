// /src/types/AssociadoType.ts
export type AssociadoType = {
  // dados basicos
  id: number;
  dataCadastro: string;
  dataAtualizacao: string;
  senha: string;
  status: "Em cadastro" | "Ativo" | "Inativo" | "Bloqueado" | "Encerrado";
  tipoAssociado: "Fornecedor" | "Beneficiado" | "Hibrido";

  // dados do associado
  nome: string;
  email: string;
  telefone: string;
  tipoPessoa: "Pessoa Física" | "Pessoa Jurídica";
  cpf_cnpj: string;

  // endereco
  cep: string;
  endereco: string;
  numero: string;
  bairro: string;
  cidade: string;
  estado: string;
  complemento: string;

  // dados adicionais
  aceitaTermos: "Sim" | "Não";
  observacoes: string;

  // pessoa fisica
  dataNascimento: string;
  nomeSocial: string;

  // pessoa juridica
  dataAbertura: string;
  razaoSocial: string;
  nomeFantasia: string;

  // dados especificos de beneficiado
  nomeConcessionaria: string;
  consumoMedio: string;
  planoDesejado: string;

  // dados especificos de fornecedor
  potenciaInstalada: string;
  disponibilidade: string;
  tipoConexao: string; // Adicionando tipo de conexão
};
