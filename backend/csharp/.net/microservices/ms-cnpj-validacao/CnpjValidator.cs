using System.Text.RegularExpressions;

namespace MsCnpjValidacao;

public static class CnpjValidator
{
    private static readonly Regex NonDigits = new(@"\D", RegexOptions.Compiled);

    public static bool IsValid(string? input)
    {
        var normalized = Normalize(input);
        if (normalized is null)
            return false;

        if (normalized.Length != 14)
            return false;

        if (AllDigitsSame(normalized))
            return false;

        var d1 = CalculateCheckDigit(normalized, 12);
        var d2 = CalculateCheckDigit(normalized, 13);

        return normalized[12] == (char)('0' + d1)
            && normalized[13] == (char)('0' + d2);
    }

    public static string? Normalize(string? input)
    {
        if (input is null)
            return null;

        var digitsOnly = NonDigits.Replace(input, "");
        if (string.IsNullOrWhiteSpace(digitsOnly))
            return null;

        return digitsOnly;
    }

    private static bool AllDigitsSame(string digits)
    {
        var first = digits[0];
        for (var i = 1; i < digits.Length; i++)
        {
            if (digits[i] != first)
                return false;
        }
        return true;
    }

    private static int CalculateCheckDigit(string digits, int length)
    {
        int[] weights = length == 12
            ? new[] { 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2 }
            : new[] { 6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2 };

        var sum = 0;
        for (var i = 0; i < length; i++)
        {
            var digit = digits[i] - '0';
            sum += digit * weights[i];
        }

        var mod = sum % 11;
        return mod < 2 ? 0 : 11 - mod;
    }
}
