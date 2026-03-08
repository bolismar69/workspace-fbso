function normalizeCnpj(input: string | null | undefined): string | null {
  if (input == null) return null;
  var digitsOnly = input.replace(/\D/g, "");
  if (digitsOnly.trim() === "") return null;
  return digitsOnly;
}

function isValidCnpj(input: string | null | undefined): boolean {
  var normalized = normalizeCnpj(input);
  if (normalized == null) return false;
  if (normalized.length !== 14) return false;
  if (allDigitsSame(normalized)) return false;

  var d1 = calculateCheckDigit(normalized, 12);
  var d2 = calculateCheckDigit(normalized, 13);

  return (
    normalized.charAt(12) === String.fromCharCode("0".charCodeAt(0) + d1) &&
    normalized.charAt(13) === String.fromCharCode("0".charCodeAt(0) + d2)
  );
}

function formatCnpjMasked(input: string | null | undefined): string {
  var digits = (input || "").replace(/\D/g, "").slice(0, 14);

  // 99.999.999/9999-99
  var p1 = digits.slice(0, 2);
  var p2 = digits.slice(2, 5);
  var p3 = digits.slice(5, 8);
  var p4 = digits.slice(8, 12);
  var p5 = digits.slice(12, 14);

  var out = "";
  if (p1) out += p1;
  if (p2) out += "." + p2;
  if (p3) out += "." + p3;
  if (p4) out += "/" + p4;
  if (p5) out += "-" + p5;
  return out;
}

function allDigitsSame(digits: string): boolean {
  var first = digits.charAt(0);
  for (var i = 1; i < digits.length; i++) {
    if (digits.charAt(i) !== first) return false;
  }
  return true;
}

function calculateCheckDigit(digits: string, length: number): number {
  var weights =
    length === 12
      ? [5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2]
      : [6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2];

  var sum = 0;
  for (var i = 0; i < length; i++) {
    sum += (digits.charCodeAt(i) - 48) * weights[i];
  }

  var mod = sum % 11;
  return mod < 2 ? 0 : 11 - mod;
}
