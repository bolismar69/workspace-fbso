namespace MsCnpjValidacao.Dto;

public sealed record CnpjValidationResponse(
    string? Input,
    string? Normalized,
    bool Valid
);
