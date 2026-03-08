package br.com.estudo.cnpj;

import java.util.Objects;

public final class CnpjValidator {

    private CnpjValidator() {
    }

    public static boolean isValid(String input) {
        String normalized = normalize(input);
        if (normalized == null) {
            return false;
        }

        if (normalized.length() != 14) {
            return false;
        }

        if (allDigitsSame(normalized)) {
            return false;
        }

        int d1 = calculateCheckDigit(normalized, 12);
        int d2 = calculateCheckDigit(normalized, 13);

        return normalized.charAt(12) == (char) ('0' + d1)
                && normalized.charAt(13) == (char) ('0' + d2);
    }

    public static String normalize(String input) {
        if (input == null) {
            return null;
        }
        String digitsOnly = input.replaceAll("\\D", "");
        if (digitsOnly.isBlank()) {
            return null;
        }
        return digitsOnly;
    }

    private static boolean allDigitsSame(String digits) {
        char first = digits.charAt(0);
        for (int i = 1; i < digits.length(); i++) {
            if (digits.charAt(i) != first) {
                return false;
            }
        }
        return true;
    }

    private static int calculateCheckDigit(String digits, int length) {
        Objects.checkFromIndexSize(0, length, digits.length());

        int[] weights = (length == 12)
                ? new int[]{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
                : new int[]{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2};

        int sum = 0;
        for (int i = 0; i < length; i++) {
            int digit = digits.charAt(i) - '0';
            sum += digit * weights[i];
        }

        int mod = sum % 11;
        return (mod < 2) ? 0 : (11 - mod);
    }
}
