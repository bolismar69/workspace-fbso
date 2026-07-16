// src/components/forms/FormBeneficiado.tsx
import React, { useEffect, useState } from "react";
import { FormSection } from "@/components/forms/FormSection";
import { BeneficiadoType } from "@/types/BeneficiadoType";
import { FieldDefinitionType } from "@/types/FieldDefinitionType";
import { fetchPlanOptions } from "@/services/servicePlans";
import { fetchConcessionariasOptions } from "@/services/serviceConcessionarias";
import { fetchConsumoMedioOptions } from "@/services/serviceConsumoMedio";
import { validateFormData } from "@/utils/validates/validateFormData";
import { isValidCPF, formatCPF } from "@/utils/validators/validatorCPF";
import { brazilianStates } from "@/constants/states";
import { isValidRG, formatRG } from "@/utils/validators/validatorRG";
// import { salvarBeneficiadoMock } from "@/services/mock/serviceBeneficiadoMock";
import { salvarBeneficiado } from "@/services/storage/serviceBeneficiado";

interface Props {
  onSubmit?: (data: BeneficiadoType) => void;
}

export function FormBeneficiado({ onSubmit }: Props) {
  const [planOptions, setPlanOptions] = useState<{ label: string; value: number }[]>([]);
  const [concessionariasOptions, setConcessionariasOptions] = useState<{ label: string; value: number }[]>([]);
  const [consumoMedioOptions, setConsumoMedioOptions] = useState<{ label: string; value: number }[]>([]);

  useEffect(() => {
    fetchPlanOptions().then(setPlanOptions);
    fetchConcessionariasOptions().then(setConcessionariasOptions);
    fetchConsumoMedioOptions().then(setConsumoMedioOptions);
  }, []);

  const fields: FieldDefinitionType<BeneficiadoType>[] = [
    {
      name: "nome",
      label: "Nome completo",
      placeholder: "João Silva",
      type: "text",
      required: true,
      keyboardType: "default",
    },
    {
      name: "cpf",
      label: "CPF",
      placeholder: "000.000.000-00",
      type: "number",
      required: true,
      keyboardType: "numeric",
      validation: (value: string) => (isValidCPF(value) ? null : "CPF inválido"),
      errorMessage: "Formato esperado: 000.000.000-00",
      formattedValue: formatCPF,
    },
    {
      name: "documentoIdentidade",
      label: "Documento de Identidade (RG)",
      type: "number",
      placeholder: "12345678-9",
      required: false,
      keyboardType: "numeric",
      validation: (value: string) => (isValidRG(value) ? null : "RG inválido"),
      formattedValue: formatRG,
      errorMessage: "Formato esperado: 12.345.678-9",
    },
    {
      name: "dataNascimento",
      label: "Data de Nascimento",
      type: "date",
      required: true,
      placeholder: "DD/MM/AAAA",
    },
    {
      name: "email",
      label: "Email",
      placeholder: "email@email.com",
      type: "email",
      required: true,
      keyboardType: "email-address",
    },
    {
      name: "telefone",
      label: "Telefone",
      placeholder: "(11) 99999-0000",
      type: "number",
      required: true,
      keyboardType: "phone-pad",
      validation: (value: string) =>
        value?.length >= 10 ? null : "Telefone inválido",
    },
    {
      name: "endereco",
      label: "Endereço",
      placeholder: "Rua Exemplo, 123",
      type: "text",
      required: true,
      keyboardType: "default",
    },
    {
      name: "cidade",
      label: "Cidade",
      placeholder: "São Paulo",
      type: "text",
      required: true,
      keyboardType: "default",
      toUpperCase: true
    },
    {
      name: "estado",
      label: "Estado",
      placeholder: "SP",
      type: "select",
      required: true,
      keyboardType: "default",
      options: brazilianStates,
    },
    {
      name: "nomeConcessionaria",
      label: "Informe a concessionaria",
      type: "select",
      required: true,
      placeholder: "Selecione a concessionária",
      options: concessionariasOptions.map(option => ({
        ...option,
        value: String(option.value),
      })),
    },
    {
      name: "consumoMedio",
      label: "Consumo médio mensal (kWh)",
      type: "select",
      required: true,
      placeholder: "Selecione o consumo médio",
      options: consumoMedioOptions.map(option => ({
        ...option,
        value: String(option.value),
      })),
    },
    {
      name: "planoDesejado",
      label: "Plano desejado",
      placeholder: "Selecione um plano",
      type: "select",
      required: true,
      options: planOptions.map(option => ({
        ...option,
        value: String(option.value),
      })),
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
  ];

  const initialValues: BeneficiadoType = {
    nome: "",
    cpf: "",
    documentoIdentidade: "",
    dataNascimento: "",
    email: "",
    telefone: "",
    endereco: "",
    cidade: "",
    estado: "",
    consumoMedio: "",
    planoDesejado: "",
    aceitaTermos: false,
    observacoes: "",
    nomeConcessionaria: "",
  };

  return (
    <FormSection
      title="Cadastro de Beneficiado"
      fields={fields}
      initialValues={initialValues}
      onSubmit={(data) => {
        const errors = validateFormData(data, fields);
        if (Object.keys(errors).length > 0) {
          console.warn("Erros de validação:", errors);
          return;
        }
        console.log("📤 Enviando dados do beneficiado:", data);
        salvarBeneficiado(data)
          .then(() => {
            console.log("✅ Beneficiado salvo com sucesso!");
          })
          .catch((error) => {
            console.error("❌ Erro ao salvar beneficiado:", error);
          });

        // Chama a função onSubmit passada como prop, se existir
        if (onSubmit) {
          onSubmit?.(data);
        }
      }}
    />
  );
}
