export function normalizeCnpj(input: string | null | undefined): string | null {
  if (!input) return null;
  const digits = input.replace(/\D/g, '');
  return digits.length === 0 ? null : digits;
}

export function formatCnpjMasked(input: string | null | undefined): string {
  const digits = (input ?? '').replace(/\D/g, '').slice(0, 14);
  if (digits.length === 0) return '';

  let out = '';
  for (let i = 0; i < digits.length; i++) {
    if (i === 2 || i === 5) out += '.';
    if (i === 8) out += '/';
    if (i === 12) out += '-';
    out += digits[i];
  }
  return out;
}

export function isValidCnpj(normalizedDigits: string | null | undefined): boolean {
  if (!normalizedDigits) return false;

  const digits = normalizedDigits.replace(/\D/g, '');
  if (digits.length !== 14) return false;

  // reject repeated digits
  if (/^(\d)\1{13}$/.test(digits)) return false;

  const dv1 = calcDigit(digits.slice(0, 12), [5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2]);
  const dv2 = calcDigit(digits.slice(0, 13), [6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2]);

  return digits[12] === dv1 && digits[13] === dv2;
}

function calcDigit(base: string, weights: number[]): string {
  let sum = 0;
  for (let i = 0; i < weights.length; i++) {
    sum += Number(base[i]) * weights[i];
  }
  const mod = sum % 11;
  const digit = mod < 2 ? 0 : 11 - mod;
  return String(digit);
}
