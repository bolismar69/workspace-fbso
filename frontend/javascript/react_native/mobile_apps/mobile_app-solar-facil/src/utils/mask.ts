export const maskCEP = (value: string) =>
  value.replace(/\D/g, "").replace(/(\d{5})(\d)/, "$1-$2");

export const maskPhone = (value: string) =>
  value
    .replace(/\D/g, "")
    .replace(/(\d{2})(\d)/, "($1) $2")
    .replace(/(\d{5})(\d)/, "$1-$2");

export function maskCPF(value: string) {
  return value.replace(/(\d{3})(\d{3})(\d{3})(\d{2})/, "$1.$2.$3-$4");
}

export function maskCNPJ(value: string) {
  return value.replace(/(\d{2})(\d{3})(\d{3})(\d{4})(\d{2})/, "$1.$2.$3/$4-$5");
}
