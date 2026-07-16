// src/utils/validateFormData.ts
import { FieldDefinitionType } from "@/types/FieldDefinitionType";

export function validateFormData<T extends Record<string, any>>(
  formData: T,
  fields: FieldDefinitionType<T>[]
): Partial<Record<keyof T, string>> {
  console.log("validateFormData");
  const errors: Partial<Record<keyof T, string>> = {};

  for (const field of fields) {
    const value = formData[field.name];
    const fieldName = field.name;

    // Validação obrigatória
    if (field.required && (value === undefined || value === null || value === "")) {
      errors[fieldName] = `${field.label} é obrigatório`;
      continue;
    }

    // Validação pattern (Regex)
    if (field.pattern && typeof value === "string" && !field.pattern.test(value)) {
      errors[fieldName] = field.errorMessage || `Formato inválido para ${field.label}`;
      continue;
    }

    // Validação customizada com (value, allValues)
    if (field.validation) {
      const result = field.validation(value, formData);
      if (result !== null && result !== undefined) {
        errors[fieldName] = result;
        continue;
      }
    }

    // Validação de opções
    if (field.options && !field.options.find((opt) => opt.value === value)) {
      errors[fieldName] = `${field.label} deve ser uma das opções válidas`;
      continue;
    }

    // Validação numérica
    if (field.type === "number" && value && !/^\d+(\.\d+)?$/.test(String(value))) {
      errors[fieldName] = `${field.label} deve conter apenas números`;
    }
  }

  console.log("errors", errors);
  return errors;
}
