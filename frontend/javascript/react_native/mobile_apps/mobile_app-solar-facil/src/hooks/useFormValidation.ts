import { yupResolver } from "@hookform/resolvers/yup";
import { useForm } from "react-hook-form";
import * as yup from "yup";

export function useFormValidation(schema: yup.AnyObjectSchema) {
  const {
    control,
    handleSubmit,
    formState: { errors },
  } = useForm({ resolver: yupResolver(schema) });

  return { control, handleSubmit, errors };
}

// Exemplo de Serviço: api.ts
export const api = {
  sendBeneficiarioData: async (data: any) => {
    return new Promise((resolve) => {
      setTimeout(() => resolve({ success: true, data }), 1000);
    });
  },
  sendFornecedorData: async (data: any) => {
    return new Promise((resolve) => {
      setTimeout(() => resolve({ success: true, data }), 1000);
    });
  },
};
