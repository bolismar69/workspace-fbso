using MsCnpjValidacao;
using MsCnpjValidacao.Dto;

var builder = WebApplication.CreateBuilder(args);

var app = builder.Build();

app.MapGet("/cnpj/validate", (string? value) =>
{
    var normalized = CnpjValidator.Normalize(value);
    var valid = CnpjValidator.IsValid(value);
    return new CnpjValidationResponse(value, normalized, valid);
});

app.MapPost("/cnpj/validate", (CnpjValidationRequest? request) =>
{
    var input = request?.Cnpj;
    var normalized = CnpjValidator.Normalize(input);
    var valid = CnpjValidator.IsValid(input);
    return new CnpjValidationResponse(input, normalized, valid);
});

app.Run();

public partial class Program { }
