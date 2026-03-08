using MsCnpjValidacao;

namespace MsCnpjValidacao.Tests;

public class CnpjValidatorTests
{
    [Fact]
    public void ShouldValidateKnownValidCnpj()
    {
        Assert.True(CnpjValidator.IsValid("04.252.011/0001-10"));
        Assert.True(CnpjValidator.IsValid("40.688.134/0001-61"));
    }

    [Fact]
    public void ShouldRejectInvalidCnpj()
    {
        Assert.False(CnpjValidator.IsValid("04.252.011/0001-11"));
        Assert.False(CnpjValidator.IsValid("00000000000000"));
        Assert.False(CnpjValidator.IsValid("11111111111111"));
        Assert.False(CnpjValidator.IsValid(""));
        Assert.False(CnpjValidator.IsValid(null));
    }

    [Fact]
    public void ShouldNormalizeDigits()
    {
        Assert.Equal("04252011000110", CnpjValidator.Normalize("04.252.011/0001-10"));
        Assert.Equal("123", CnpjValidator.Normalize("  1-2-3  "));
        Assert.Null(CnpjValidator.Normalize(""));
        Assert.Null(CnpjValidator.Normalize("   "));
        Assert.Null(CnpjValidator.Normalize(null));
    }
}
