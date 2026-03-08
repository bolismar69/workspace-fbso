using System.Net;
using System.Net.Http.Json;
using Microsoft.AspNetCore.Mvc.Testing;
using MsCnpjValidacao.Dto;

namespace MsCnpjValidacao.Tests;

public class CnpjResourceTests : IClassFixture<WebApplicationFactory<Program>>
{
    private readonly HttpClient _client;

    public CnpjResourceTests(WebApplicationFactory<Program> factory)
    {
        _client = factory.CreateClient();
    }

    [Fact]
    public async Task ValidateGetShouldReturnValidTrue()
    {
        var res = await _client.GetAsync("/cnpj/validate?value=04.252.011/0001-10");
        Assert.Equal(HttpStatusCode.OK, res.StatusCode);

        var body = await res.Content.ReadFromJsonAsync<CnpjValidationResponse>();
        Assert.NotNull(body);
        Assert.True(body!.Valid);
        Assert.Equal("04252011000110", body.Normalized);
    }

    [Fact]
    public async Task ValidatePostShouldReturnValidFalse()
    {
        var res = await _client.PostAsJsonAsync("/cnpj/validate", new CnpjValidationRequest("04.252.011/0001-11"));
        Assert.Equal(HttpStatusCode.OK, res.StatusCode);

        var body = await res.Content.ReadFromJsonAsync<CnpjValidationResponse>();
        Assert.NotNull(body);
        Assert.False(body!.Valid);
        Assert.Equal("04252011000111", body.Normalized);
    }
}
