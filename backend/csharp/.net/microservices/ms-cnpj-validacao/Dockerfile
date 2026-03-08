# syntax=docker/dockerfile:1

FROM mcr.microsoft.com/dotnet/sdk:8.0 AS build
WORKDIR /src

COPY MsCnpjValidacao.csproj ./
RUN dotnet restore ./MsCnpjValidacao.csproj

COPY . ./
RUN dotnet publish ./MsCnpjValidacao.csproj -c Release -o /app/out --no-restore

FROM mcr.microsoft.com/dotnet/sdk:8.0 AS test
WORKDIR /src

COPY MsCnpjValidacao.csproj ./
COPY MsCnpjValidacao.Tests/MsCnpjValidacao.Tests.csproj ./MsCnpjValidacao.Tests/
RUN dotnet restore ./MsCnpjValidacao.Tests/MsCnpjValidacao.Tests.csproj

COPY . ./
WORKDIR /src/MsCnpjValidacao.Tests
ENTRYPOINT ["dotnet", "test", "-c", "Release", "--no-restore"]

FROM mcr.microsoft.com/dotnet/aspnet:8.0
WORKDIR /app

COPY --from=build /app/out ./

ENV ASPNETCORE_URLS=http://0.0.0.0:8080
EXPOSE 8080

ENTRYPOINT ["dotnet", "MsCnpjValidacao.dll"]
