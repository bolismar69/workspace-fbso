using System.Text;

namespace WebAppCnpjValidacao.Domain;

public static class CnpjValidator
{
    public static string? Normalize(string? input)
    {
        if (string.IsNullOrWhiteSpace(input))
        {
            return null;
        }

        var digits = ExtractDigits(input);
        return digits.Length == 0 ? null : digits;
    }

    public static string FormatMasked(string? input)
    {
        var digits = ExtractDigits(input ?? string.Empty);
        if (digits.Length == 0)
        {
            return string.Empty;
        }

        // Apply mask: 99.999.999/9999-99
        var sb = new StringBuilder();
        for (var i = 0; i < digits.Length && i < 14; i++)
        {
            if (i == 2 || i == 5)
            {
                sb.Append('.');
            }
            else if (i == 8)
            {
                sb.Append('/');
            }
            else if (i == 12)
            {
                sb.Append('-');
            }

            sb.Append(digits[i]);
        }

        return sb.ToString();
    }

    public static bool IsValid(string? normalizedDigits)
    {
        if (string.IsNullOrWhiteSpace(normalizedDigits))
        {
            return false;
        }

        var digits = ExtractDigits(normalizedDigits);
        if (digits.Length != 14)
        {
            return false;
        }

        // Reject sequences like 000... or 111...
        var allSame = true;
        for (var i = 1; i < digits.Length; i++)
        {
            if (digits[i] != digits[0])
            {
                allSame = false;
                break;
            }
        }

        if (allSame)
        {
            return false;
        }

        var dv1 = CalculateDigit(digits[..12], new[] { 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2 });
        var dv2 = CalculateDigit(digits[..13], new[] { 6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2 });

        return digits[12] == dv1 && digits[13] == dv2;
    }

    private static char CalculateDigit(string digits, int[] weights)
    {
        var sum = 0;
        for (var i = 0; i < weights.Length; i++)
        {
            sum += (digits[i] - '0') * weights[i];
        }

        var mod = sum % 11;
        var result = mod < 2 ? 0 : 11 - mod;
        return (char)('0' + result);
    }

    private static string ExtractDigits(string input)
    {
        var sb = new StringBuilder(input.Length);
        foreach (var ch in input)
        {
            if (ch >= '0' && ch <= '9')
            {
                sb.Append(ch);
            }
        }

        return sb.ToString();
    }
}
