// src/components/forms/FormFornecedor.tsx
import React from "react";
import { FormSection } from "./FormSection";
import { FieldDefinitionType } from "@/types/FieldDefinitionType";
import { FornecedorType } from "@/types/FornecedorType";
import { validateFormData } from "@/utils/validates/validateFormData";
import { isValidCNPJ, formatCNPJ } from "@/utils/validators/validatorCNPJ";
import { brazilianStates } from "@/constants/states";
import { salvarFornecedorMock } from "@/services/mock/serviceFornecedorMock";

interface Props {
  onSubmit?: (data: FornecedorType) => void;
}

export function FormFornecedor({ onSubmit }: Props) {
  console.log("FormFornecedor rendered");

  const fields: FieldDefinitionType<FornecedorType>[] = [
    {
      name: "nome",
      label: "Nome completo",
      placeholder: "Digite seu nome",
      keyboardType: "default",
      required: true,
      type: "text",
    },
    {
      name: "email",
      label: "E-mail",
      placeholder: "exemplo@email.com",
      keyboardType: "email-address",
      required: true,
      type: "email",
    },
    {
      name: "telefone",
      label: "Telefone",
      placeholder: "(11) 91234-5678",
      keyboardType: "phone-pad",
      required: true,
      type: "text",
    },
    {
      name: "cnpj",
      label: "CNPJ",
      placeholder: "12.345.678/0001-90",
      type: "number",
      required: true,
      keyboardType: "numeric",
      validation: (value: string) => (isValidCNPJ(value) ? null : "CNPJ inválido"),
      errorMessage: "Formato inválido. Use: 00.000.000/0000-00",
      formattedValue: formatCNPJ,
    },
    {
      name: "dataFundacao",
      label: "Data de Fundação",
      type: "date",
      required: true,
      placeholder: "DD/MM/AAAA",
    },
    {
      name: "endereco",
      label: "Endereço",
      placeholder: "Rua, número, bairro",
      keyboardType: "default",
      required: true,
      type: "text",
    },
    {
      name: "cidade",
      label: "Cidade",
      placeholder: "São Paulo",
      keyboardType: "default",
      required: true,
      type: "text",
    },
    {
      name: "estado",
      label: "Estado",
      placeholder: "Selecione seu estado",
      keyboardType: "default",
      required: true,
      type: "select",
      options: brazilianStates,
    },
    {
      name: "potenciaInstalada",
      label: "Potência Instalada (kW)",
      placeholder: "Ex: 10",
      keyboardType: "numeric",
      required: true,
      type: "number",
    },
    {
      name: "disponibilidade",
      label: "Disponibilidade de Energia (kWh)",
      placeholder: "Ex: 150",
      keyboardType: "numeric",
      required: true,
      type: "number",
    },
    {
      name: "aceitaTermos",
      label: "Aceito os termos de uso",
      type: "switch",
      required: true,
      placeholder: "",
    },
    {
      name: "observacoes",
      label: "Observações",
      placeholder: "Descreva aqui informações adicionais que queira nos passar...",
      type: "textarea",
      required: false
    },
    {
      name: "tipoConexao",
      label: "Tipo de Conexão",
      placeholder: "Selecione o tipo",
      type: "select",
      required: true,
      options: [
        { label: "Monofásica", value: "monofasica" },
        { label: "Bifásica", value: "bifasica" },
        { label: "Trifásica", value: "trifasica" },
      ],
    },
  ];
  const initialValues: FornecedorType = {
    nome: "",
    email: "",
    telefone: "",
    cnpj: "",
    dataFundacao: "",
    endereco: "",
    cidade: "",
    estado: "",
    potenciaInstalada: "",
    disponibilidade: "",
    aceitaTermos: false,
    observacoes: "",
    tipoConexao: "",
  };

  return (
    <FormSection
      title="Cadastro de Fornecedor"
      fields={fields}
      initialValues={initialValues}
      onSubmit={(data) => {
        const errors = validateFormData(data, fields);
        if (Object.keys(errors).length > 0) {
          console.warn("Erros de validação:", errors);
          return;
        }
        console.log("Dados do fornecedor:", data);
        salvarFornecedorMock(data)
          .then((response) => {
            console.log("✅ Dados salvos com sucesso - salvarFornecedorMock:", response);
            alert("Fornecedor cadastrado com sucesso (salvarFornecedorMock)!");
          })
          .catch((error) => {
            console.error("❌ Erro ao salvar fornecedor - salvarFornecedorMock:", error);
            alert("Erro ao cadastrar fornecedor. Tente novamente mais tarde (salvarFornecedorMock).");
          });
        if (onSubmit) {
          onSubmit?.(data);
        }
      }}
    />
  );
}
