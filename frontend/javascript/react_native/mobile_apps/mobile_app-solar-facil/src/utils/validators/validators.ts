import * as yup from "yup";

export const beneficiarioSchema = yup.object({
  nome: yup.string().required("Nome é obrigatório"),
  cep: yup.string().required("CEP é obrigatório"),
  email: yup.string().email("Email inválido").required("Email é obrigatório"),
});

export const fornecedorSchema = yup.object({
  nome: yup.string().required("Nome é obrigatório"),
  cpf: yup.string().required("CPF é obrigatório"),
  excedente: yup.number().required("Informe o excedente de energia"),
});
  