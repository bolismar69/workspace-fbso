// src/utils/validatorCPF.ts
export function isValidCPF(cpf: string): boolean {
  console.log("Validando CPF:", cpf);
  cpf = cpf.replace(/[^\d]/g, ""); // remove pontos e traços

  if (cpf.length !== 11 || /^(\d)\1+$/.test(cpf)) {
    console.log("CPF inválido");
    return false;
  }

  let sum = 0;
  for (let i = 0; i < 9; i++) {
    sum += parseInt(cpf.charAt(i)) * (10 - i);
  }

  let remainder = (sum * 10) % 11;
  if (remainder === 10 || remainder === 11) remainder = 0;
  if (remainder !== parseInt(cpf.charAt(9))) {
    console.log("CPF inválido");
    return false;
  }

  sum = 0;
  for (let i = 0; i < 10; i++) {
    sum += parseInt(cpf.charAt(i)) * (11 - i);
  }

  remainder = (sum * 10) % 11;
  if (remainder === 10 || remainder === 11) remainder = 0;
  if (remainder !== parseInt(cpf.charAt(10))) {
    console.log("CPF inválido");
    return false;
  }
  return true;
}
export function CPFEhValido(cpf: string): string | null {
  if (!isValidCPF(cpf)) return "CPF inválido";
  return null;
}
export function formatCPF(cpf: string): string {
  cpf = cpf.replace(/\D/g, ""); // Remove tudo que não for número
  if (cpf.length !== 11) return cpf;
  return cpf.replace(/(\d{3})(\d{3})(\d{3})(\d{2})/, "$1.$2.$3-$4");
}
// function formatCPF(value: string): string {
//   return value
//     .replace(/\D/g, "")
//     .replace(/(\d{3})(\d)/, "$1.$2")
//     .replace(/(\d{3})(\d)/, "$1.$2")
//     .replace(/(\d{3})(\d{1,2})$/, "$1-$2");
// }
