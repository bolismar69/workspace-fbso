import 'package:flutter/services.dart';

class Cnpj {
  static String? normalize(String? input) {
    if (input == null || input.trim().isEmpty) {
      return null;
    }

    final digits = _extractDigits(input);
    return digits.isEmpty ? null : digits;
  }

  static bool isValid(String? normalizedDigits) {
    if (normalizedDigits == null || normalizedDigits.trim().isEmpty) {
      return false;
    }

    final digits = _extractDigits(normalizedDigits);
    if (digits.length != 14) {
      return false;
    }

    // reject repeated digits: 000... / 111... etc
    final first = digits[0];
    var allSame = true;
    for (var i = 1; i < digits.length; i++) {
      if (digits[i] != first) {
        allSame = false;
        break;
      }
    }
    if (allSame) {
      return false;
    }

    final dv1 = _calcDigit(digits.substring(0, 12), const [5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2]);
    final dv2 = _calcDigit(digits.substring(0, 13), const [6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2]);

    return digits[12] == dv1 && digits[13] == dv2;
  }

  static String formatMasked(String? input) {
    final digits = _extractDigits(input ?? '');
    if (digits.isEmpty) {
      return '';
    }

    final truncated = digits.length > 14 ? digits.substring(0, 14) : digits;
    final buffer = StringBuffer();
    for (var i = 0; i < truncated.length; i++) {
      if (i == 2 || i == 5) buffer.write('.');
      if (i == 8) buffer.write('/');
      if (i == 12) buffer.write('-');
      buffer.write(truncated[i]);
    }
    return buffer.toString();
  }

  static String _calcDigit(String base, List<int> weights) {
    var sum = 0;
    for (var i = 0; i < weights.length; i++) {
      sum += int.parse(base[i]) * weights[i];
    }
    final mod = sum % 11;
    final digit = mod < 2 ? 0 : 11 - mod;
    return digit.toString();
  }

  static String _extractDigits(String input) {
    final buffer = StringBuffer();
    for (final codeUnit in input.codeUnits) {
      if (codeUnit >= 48 && codeUnit <= 57) {
        buffer.writeCharCode(codeUnit);
      }
    }
    return buffer.toString();
  }
}

class CnpjMaskFormatter extends TextInputFormatter {
  const CnpjMaskFormatter();

  @override
  TextEditingValue formatEditUpdate(TextEditingValue oldValue, TextEditingValue newValue) {
    final masked = Cnpj.formatMasked(newValue.text);

    return TextEditingValue(
      text: masked,
      selection: TextSelection.collapsed(offset: masked.length),
      composing: TextRange.empty,
    );
  }
}
