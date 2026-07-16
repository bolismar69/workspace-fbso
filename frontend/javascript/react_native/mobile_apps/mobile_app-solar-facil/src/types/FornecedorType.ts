// /src/types/FornecedorType.ts
export type FornecedorType = {
  nome: string;
  email: string;
  telefone: string;
  cnpj: string;
  dataFundacao: string;
  endereco: string;
  cidade: string;
  estado: string;
  potenciaInstalada: string;
  disponibilidade: string;
  aceitaTermos: boolean;
  observacoes: string;
  tipoConexao: string; // Adicionando tipo de conexão
};
