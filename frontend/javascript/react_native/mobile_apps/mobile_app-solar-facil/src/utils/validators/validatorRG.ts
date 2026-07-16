export function isValidRG(rg: string): boolean {
  console.log("Validando RG:", rg);
  rg = rg.replace(/[^\d]/g, ""); // Remove pontos, traços e espaços

  if (rg.length !== 9 || /^(\d)\1+$/.test(rg)) {
    console.log("RG inválido");
    return false;
  }

  const digits = rg.slice(0, -1); // Pega os 8 primeiros dígitos
  console.log("Dígitos do RG:", digits);
  const checkDigit = parseInt(rg.slice(-1), 10); // Último dígito como dígito verificador
  console.log("Dígito verificador:", checkDigit);
  let sum = 0;

  for (let i = 0; i < digits.length; i++) {
    sum += parseInt(digits[i], 10) * (9 - i); // Multiplica os dígitos pela sequência decrescente de 9 a 2
    console.log(`Soma parcial após o dígito ${i + 1}:`, sum);
  }

  const calculatedCheckDigit = sum % 11;
  console.log("Dígito verificador calculado:", calculatedCheckDigit);

  // RG com dígito verificador '0' quando o cálculo resulta em 10
  let ehvalido = false;
  if (calculatedCheckDigit === 10) {
    ehvalido =  checkDigit === 0 ? true : false;
  } else {
    ehvalido =   checkDigit === calculatedCheckDigit ? true : false;
  }
  console.log("RG válido:", ehvalido);
  return ehvalido;
}
export function RGEhValido(rg: string): string | null {
  if (!isValidRG(rg)) return "RG inválido";
  return null;
}
export function formatRG(rg: string): string {
  rg = rg.replace(/[^\d]/g, ""); // Remove tudo que não for número
  if (rg.length < 8 || rg.length > 10) return rg; // Retorna sem formatação se o tamanho for inválido

  // Formata o RG para o padrão 12.345.678-9
  if (rg.length === 9) {
    return rg.replace(/(\d{2})(\d{3})(\d{3})(\d)/, "$1.$2.$3-$4");
  } else {
    return rg.replace(/(\d{2})(\d{3})(\d{3})/, "$1.$2.$3");
  }
}
