// src/utils/validatorCNPJ.ts
// função que valida CNPJ
export function isValidCNPJ(cnpj: string): boolean {
  console.log("Validando CNPJ:", cnpj);
  cnpj = cnpj.replace(/[^\d]/g, ""); // Remove pontos, traços e barras

  if (cnpj.length !== 14 || /^(\d)\1+$/.test(cnpj)) {
    console.log("CNPJ inválido");
    return false;
  }

  // Cálculo do primeiro dígito verificador
  let sum = 0;
  for (let i = 0; i < 12; i++) {
    const weight = i < 4 ? 5 - i : 13 - i;
    sum += parseInt(cnpj.charAt(i)) * weight;
  }

  let remainder = sum % 11;
  let digit1 = remainder < 2 ? 0 : 11 - remainder;

  if (digit1 !== parseInt(cnpj.charAt(12))) {
    console.log("CNPJ inválido");
    return false;
  }

  // Cálculo do segundo dígito verificador
  sum = 0;
  for (let i = 0; i < 13; i++) {
    const weight = i < 5 ? 6 - i : 14 - i;
    sum += parseInt(cnpj.charAt(i)) * weight;
  }

  remainder = sum % 11;
  let digit2 = remainder < 2 ? 0 : 11 - remainder;

  if (digit2 !== parseInt(cnpj.charAt(13))) {
    console.log("CNPJ inválido");
    return false;
  }

  console.log("CNPJ válido");
  return true;
}

export function CNPJEhValido(cnpj: string): string | null {
  if (!isValidCNPJ(cnpj)) return "CNPJ inválido";
  return null;
}

export function formatCNPJ(cnpj: string): string {
    cnpj = cnpj.replace(/\D/g, ""); // Remove tudo que não for número
  if (cnpj.length !== 14) return cnpj;
  return cnpj.replace(/(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})/, "$1.$2.$3/$4-$5");
}
