
# MICROSERVICE:
- Linguagem: JAVA-21 com GraalVM

## Objetivo: construir microserviço que recebera um CEP, consultará em API publica (https://cep.awesomeapi.com.br/) recebera lat/long desse CEP, e após isso fara cosulta na mesma API em outro operção que recebera lat/long e distancia para descobrir todas as localidades dentro do raio informado. Assim podera cruzar com tabelas ja existentes que possuem o cadastro de municipios, distritos, sub-distritos brasileiros para retornar as cidades vizinhas que estão dentro do raio solicitado.

## Estrutura de pastas do projeto -- estrutura proposta

```text
.
├── src/main/java/com/fbso/ms-geolocalidade/
│   ├── config/
│   │   ├── ClientConfig.java        # Configuração do RestClient (Java 21) e beans de HTTP.
│   │   ├── BatchConfig.java         # Definição de Jobs e Steps do Spring Batch.
│   │   └── IngestaoHints.java       # RuntimeHints para compatibilidade com GraalVM AOT.
│   ├── controller/
│   │   └── LocalidadeController.java # Endpoints REST para consulta de CEP e vizinhança. 
│   ├── dto/                         # Records Java 21 para transporte de dados (imutáveis).
│   │   ├── AwesomeCepDTO.java       # Mapeamento da API externa.
│   │   └── RespostaCompletaDTO.java # Resposta consolidada para o cliente final.
│   ├── entity/                      # Entidades JPA com suporte a Hibernate Spatial.
│   │   ├── Municipio.java           # Mapeamento da tabela de municípios (IBGE).
│   │   └── Subdistrito.java         # Mapeamento da tabela de subdistritos (IBGE).
│   ├── repository/                  # Interfaces Spring Data JPA para persistência.
│   ├── service/                     # Regras de negócio e orquestração.
│   │   ├── AwesomeCepService.java   # Integração segura (Token/Key) com a API externa.
│   │   ├── LocalidadeOrchestrator.java # Lógica de união (API + Dados Locais IBGE).
│   │   └── CepUtils.java            # Utilitário com Pattern pré-compilado para limpeza de CEP.
│   └── GeoApplication.java          # Classe principal (Entry point).
├── src/main/resources/
│   └── application.yml              # Configurações de ambiente, tokens e caminhos de volumes.
├── Dockerfile                       # Build multi-stage para gerar imagem nativa GraalVM.
├── pom.xml                          # Gerenciamento de dependências e plugin nativo do Spring.
└── k8s/                             # Manifestos de Deployment e PVC para produção.
```

### 


### Vamos Utilizar a API publica **AwesomeAPI** realmente simplifica drasticamente a arquitetura do seu MVP, pois ela resolve o maior desafio técnico que tínhamos: a **geolocalização (Lat/Lng)** e a **busca por raio** de forma externa e pronta para uso.

Aqui está como fica o design da solução, muito mais ágil:

### 1. Nova Arquitetura de Fluxo (Simplificada)

O microserviço passa a ser um **Orquestrador de APIs** com cache local:

1.  **Chamada Inicial:** Recebe o CEP e o Raio.
2.  **Step 1 (Geocodificação):** Consulta `https://cep.awesomeapi.com.br/json/{cep}` para obter a Latitude e Longitude do ponto central.
3.  **Step 2 (Busca Espacial):** Usa as coordenadas obtidas para consultar `https://cep.awesomeapi.com.br/search?lat={lat}&lng={lng}&d={raio}`.
4.  **Step 3 (Enriquecimento):** Cruza os códigos IBGE retornados com sua base local (os arquivos DTB que já mapeamos) para retornar os nomes oficiais de Municípios, Distritos e Subdistritos.

---

### 2. Implementação do novo DTO e Service

#### DTO para a AwesomeAPI
```java
public record AwesomeCepDTO(
    String cep,
    String address_type,
    String address_name,
    String address,
    String state,
    String city,
    String city_ibge,
    String lat,
    String lng,
    String district // Bairro
) {}
```

#### Service de Consulta Integrado
```java
@Service
public class AwesomeCepService {

    private final RestTemplate restTemplate;
    private final String BASE_URL = "https://cep.awesomeapi.com.br/json/";
    private final String SEARCH_URL = "https://cep.awesomeapi.com.br/search";

    public AwesomeCepService(RestTemplate restTemplate) {
        this.restTemplate = restTemplate;
    }

    // Obtém coordenadas do CEP de origem
    public AwesomeCepDTO obterCoordenadas(String cep) {
        return restTemplate.getForObject(BASE_URL + cep.replaceAll("\\D", ""), AwesomeCepDTO.class);
    }

    // Busca vizinhos num raio
    public List<AwesomeCepDTO> buscarVizinhosNoRaio(String lat, String lng, Double raio) {
        String url = String.format("%s?lat=%s&lng=%s&d=%s", SEARCH_URL, lat, lng, raio);
        AwesomeCepDTO[] response = restTemplate.getForObject(url, AwesomeCepDTO[].class);
        return response != null ? Arrays.asList(response) : Collections.emptyList();
    }
}
```

---

### 3. Ajuste no Endpoint (Controller)

O seu Controller agora retorna uma visão consolidada que mistura os dados da API com os dados oficiais do IBGE que você tem nos arquivos CSV.

```java
@GetMapping("/vizinhas-agil")
public ResponseEntity<ResponseConsolidadaDTO> buscarVizinhasAwesome(
        @RequestParam String cep,
        @RequestParam(defaultValue = "5") Double raio) {
    
    // 1. Pega Lat/Lng do CEP informado
    AwesomeCepDTO origem = awesomeService.obterCoordenadas(cep);
    
    // 2. Busca endereços/CEPs vizinhos no raio
    List<AwesomeCepDTO> vizinhos = awesomeService.buscarVizinhosNoRaio(origem.lat(), origem.lng(), raio);
    
    // 3. (Opcional) Enriquece com sua base DTB local para trazer nomes de distritos/subdistritos
    // Aqui você usaria o código IBGE retornado pela API para buscar na sua tabela 'municipios'
    
    return ResponseEntity.ok(new ResponseConsolidadaDTO(origem, vizinhos));
}
```

---

### 4. Vantagens e Pontos de Atenção desta Solução

**Vantagens:**
* **Velocidade de Entrega:** Você não precisa fazer o ETL dos arquivos de coordenadas (`11.csv`, `35.csv`, etc.), o que poupa semanas de trabalho e infraestrutura.
* **Simplicidade:** O banco de dados pode ser um SQL comum (MySQL/H2/Postgres simples) sem necessidade de extensões espaciais complexas para o MVP.
* **Manutenção:** A atualização das coordenadas fica a cargo da API.

**Pontos de Atenção:**
* **Limites de Taxa (Rate Limit):** Verifique se o volume de requisições do seu microserviço não excederá os limites da AwesomeAPI.
* **Dependência Externa:** Se o serviço da AwesomeAPI oscilar, seu microserviço para de funcionar. Recomendo implementar um **Cache (Redis ou Caffeine)** para os CEPs já consultados.
* **Granularidade:** A busca por raio da API retorna "Endereços/CEPs" próximos. Para retornar "Cidades" próximas, você terá que filtrar os resultados para remover duplicidade de códigos IBGE.


### O **OrchestratorService** funcionará como um motor de enriquecimento: ele recebe as coordenadas, descobre quais CEPs/localidades estão próximas e, para cada resultado, "mergulha" na sua base local para trazer o nome correto do Distrito e Subdistrito.

### 1. DTO de Resposta Consolidada
Este DTO refletirá exatamente a estrutura que você solicitou, unindo os dados externos e internos.

```java
public record RespostaCompletaDTO(
    AwesomeCepDTO zipcodeInfo, // Payload original (usaremos o da AwesomeAPI que é similar)
    LocalidadeDetalhadaDTO localidade, // Dados do CEP de origem na base local
    List<VizinhoEnriquecidoDTO> cidadesProximas // Lista de vizinhos com nomes do IBGE
) {}

public record VizinhoEnriquecidoDTO(
    String cep,
    String cidade,
    String ibge,
    String distrito,
    String subdistrito,
    Double distanciaKm
) {}
```

---

### 2. Implementação do OrchestratorService

Este serviço orquestra as duas chamadas à API e faz o "Join" em memória com os Repositories do Spring Data JPA.

```java
@Service
public class LocalidadeOrchestratorService {

    @Autowired
    private AwesomeCepService awesomeService;

    @Autowired
    private MunicipioRepository municipioRepository;

    @Autowired
    private SubdistritoRepository subdistritoRepository;

    public RespostaCompletaDTO processarBuscaPorCep(String cep, Double raioKm) {
        // 1. Obtém Lat/Lng do CEP de origem
        AwesomeCepDTO origemApi = awesomeService.obterCoordenadas(cep);
        
        // 2. Busca vizinhos via AwesomeAPI (Busca Espacial Externa)
        List<AwesomeCepDTO> vizinhosApi = awesomeService.buscarVizinhosNoRaio(
            origemApi.lat(), origemApi.lng(), raioKm);

        // 3. Enriquecimento: Para cada vizinho, buscamos nomes oficiais na base local
        List<VizinhoEnriquecidoDTO> vizinhosEnriquecidos = vizinhosApi.stream()
            .map(this::enriquecerComDadosIBGE)
            .toList();

        // 4. Detalha a localidade de origem (Busca por código IBGE na base DTB local)
        LocalidadeDetalhadaDTO localidadeInfo = municipioRepository.findById(origemApi.city_ibge())
            .map(m -> new LocalidadeDetalhadaDTO(
                m.getCodigoIbge7(), m.getNomeMunicipio(), m.getUfSigla(),
                origemApi.lat(), origemApi.lng()
            )).orElse(null);

        return new RespostaCompletaDTO(origemApi, localidadeInfo, vizinhosEnriquecidos);
    }

    private VizinhoEnriquecidoDTO enriquecerComDadosIBGE(AwesomeCepDTO v) {
        // Tentamos buscar se o código IBGE retornado possui um Subdistrito mapeado
        // A AwesomeAPI retorna o city_ibge (7 dígitos). 
        // Se precisarmos de subdistrito, cruzamos com a tabela local de subdistritos.
        String nomeSubdistrito = subdistritoRepository.findNomeByCodigo(v.city_ibge())
            .orElse("Sede / Não Informado");

        return new VizinhoEnriquecidoDTO(
            v.cep(),
            v.city(),
            v.city_ibge(),
            v.district(), // O campo 'district' da API costuma ser o Bairro/Distrito
            nomeSubdistrito,
            null // A distância vem no campo 'd' da API se disponível
        );
    }
}
```

---

### Mesmo usando a AwesomeAPI para as coordenadas, o processo de carga dos arquivos `DTB_Municípios`, `DTB_Distritos` e `DTB_Subdistritos` continua sendo vital por três motivos:

1.  **Padronização:** A API pode retornar "S. Paulo", enquanto sua base IBGE terá "São Paulo". Para relatórios oficiais, o dado do IBGE é o que vale.
2.  **Hierarquia:** A API raramente sabe qual é a "Região Geográfica Imediata" ou o "Código do Subdistrito" (11 dígitos). Sua base local provê esse detalhamento.
3.  **Fallback e Validação:** Você pode validar se o `city_ibge` retornado pela API realmente existe e está ativo (cruzando com o arquivo de `Distritos Novos e Extintos`).

### 4. Ajuste no Repository para Busca Rápida
Como faremos muitas buscas por código IBGE durante o enriquecimento da lista de vizinhos, certifique-se de que os índices estão criados:

```java
@Repository
public interface SubdistritoRepository extends JpaRepository<Subdistrito, String> {
    // Busca rápida pelo prefixo do código IBGE ou código completo
    @Query("SELECT s.nomeSubdistrito FROM Subdistrito s WHERE s.codigoSubdistrito11 LIKE :codigoIBGE%")
    Optional<String> findNomeByCodigo(String codigoIBGE);
}
```

---

## AUTENTICACAO

### 1. Configuração de Propriedades (`application.yml`)
Mantenha suas credenciais seguras fora do código.

```yaml
awesomeapi:
  token: ${AWESOME_API_TOKEN:seu_token_aqui}
  key: ${AWESOME_API_KEY:sua_key_aqui}
  base-url: https://cep.awesomeapi.com.br
```

### 2. Implementação do Service com Autenticação
Vamos ajustar o `AwesomeCepService` para aplicar o token na URL no primeiro passo e a API Key no Header no segundo passo.

```java
@Service
public class AwesomeCepService {

    @Value("${awesomeapi.token}")
    private String apiToken;

    @Value("${awesomeapi.key}")
    private String apiKey;

    @Value("${awesomeapi.base-url}")
    private String baseUrl;

    private final RestTemplate restTemplate;

    public AwesomeCepService(RestTemplate restTemplate) {
        this.restTemplate = restTemplate;
    }

    // Chamada 1: CEP -> Lat/Lng (Usando Token na URL)
    public AwesomeCepDTO obterCoordenadas(String cep) {
        String url = String.format("%s/json/%s?token=%s", 
            baseUrl, cep.replaceAll("\\D", ""), apiToken);
            
        return restTemplate.getForObject(url, AwesomeCepDTO.class);
    }

    // Chamada 2: Busca por Raio (Usando x-api-key no Header)
    public List<AwesomeCepDTO> buscarVizinhosNoRaio(String lat, String lng, Double raio) {
        String url = String.format("%s/search?lat=%s&lng=%s&d=%s", 
            baseUrl, lat, lng, raio);

        // Configuração dos Headers de Autenticação
        HttpHeaders headers = new HttpHeaders();
        headers.set("x-api-key", apiKey);
        HttpEntity<String> entity = new HttpEntity<>(headers);

        ResponseEntity<AwesomeCepDTO[]> response = restTemplate.exchange(
            url, 
            HttpMethod.GET, 
            entity, 
            AwesomeCepDTO[].class
        );

        return response.getBody() != null ? Arrays.asList(response.getBody()) : Collections.emptyList();
    }
}
```

### 3. Ajuste no OrchestratorService (Refinado)
O orquestrador agora utiliza essas chamadas seguras e faz o cruzamento com sua base local de subdistritos para enriquecer os dados.

```java
@Service
public class LocalidadeOrchestratorService {

    @Autowired
    private AwesomeCepService awesomeService;

    @Autowired
    private SubdistritoRepository subdistritoRepository;

    public RespostaCompletaDTO processarBuscaPorCep(String cep, Double raioKm) {
        // 1. Geocodificação segura
        AwesomeCepDTO origem = awesomeService.obterCoordenadas(cep);

        // 2. Busca por raio segura (com API Key no header)
        List<AwesomeCepDTO> vizinhosApi = awesomeService.buscarVizinhosNoRaio(
            origem.lat(), origem.lng(), raioKm);

        // 3. Cruzamento com dados oficiais IBGE carregados via Spring Batch
        List<VizinhoEnriquecidoDTO> vizinhosEnriquecidos = vizinhosApi.stream()
            .map(v -> {
                // Busca o nome oficial do subdistrito na sua base local
                String subdistritoOficial = subdistritoRepository
                    .findNomeByCodigo(v.city_ibge())
                    .orElse("Sede");

                return new VizinhoEnriquecidoDTO(
                    v.cep(), v.city(), v.city_ibge(), 
                    v.district(), subdistritoOficial, null
                );
            }).toList();

        return new RespostaCompletaDTO(origem, null, vizinhosEnriquecidos);
    }
}
```

---

# README.md

## GeoLocalidade API 2026

Microserviço de alta performance para consulta de geolocalização, CEP e vizinhança, integrando dados oficiais do **IBGE (DTB 2024)** com geocodificação em tempo real via **AwesomeAPI**.

Otimizado para **Java 21** e compilação nativa com **Oracle GraalVM**, garantindo baixíssimo consumo de memória e inicialização instantânea em ambientes **Kubernetes**.

## 🚀 Arquitetura e Tecnologias

- **Java 21 (LTS)**: Utilizando `Records` para DTOs e `RestClient` para chamadas fluídas.
- **Oracle GraalVM 21**: Compilação nativa (*Ahead-of-Time*) para performance extrema.
- **Spring Boot 3.4+**: Com suporte nativo a AOT.
- **Spring Batch**: Processamento robusto dos arquivos da Divisão Territorial Brasileira (DTB).
- **PostgreSQL + Hibernate Spatial**: Armazenamento e validação de dados geográficos.
- **AwesomeAPI**: Provedor externo de geocodificação e busca por raio.



## 📂 Estrutura de Dados (Ingestão)

A aplicação consome os dados oficiais do IBGE através de arquivos CSV montados em volumes externos.

### Hierarquia Suportada:
1. **Municípios**: `RELATORIO_DTB_BRASIL_2024_MUNICIPIOS.csv`
2. **Distritos**: `RELATORIO_DTB_BRASIL_2024_DISTRITOS.csv`
3. **Subdistritos**: `RELATORIO_DTB_BRASIL_2024_SUBDISTRITOS.csv`

O processo de ingestão é executado via Spring Batch, ignorando cabeçalhos de metadados e garantindo a integridade referencial entre as tabelas.

## 🛠️ Configuração e Instalação

### Pré-requisitos
- Docker & Kubernetes (ou Minikube)
- GraalVM 21 (se desejar compilar localmente)
- Maven 3.9+

### Variáveis de Ambiente Necessárias
| Variável | Descrição |
| :--- | :--- |
| `AWESOME_API_TOKEN` | Token para consulta de CEP (passado via URL) |
| `AWESOME_API_KEY` | API Key para busca por raio (passada via Header) |
| `APP_IMPORT_PATH` | Caminho do volume onde os CSVs estão montados (ex: `/mnt/data/ibge`) |

### Compilação Nativa
Para gerar o binário nativo localmente:
```bash
./mvnw native:compile -Pnative -DskipTests
```

## 🐳 Docker e Kubernetes

O projeto utiliza um `Dockerfile` multi-stage para gerar imagens mínimas.

```bash
# Build da imagem
docker build -t geo-service-api:latest .
```

No **Kubernetes**, certifique-se de que o volume contendo os arquivos CSV está corretamente mapeado para o caminho definido em `APP_IMPORT_PATH`.

## 📍 Endpoints Principais

### Busca de Vizinhas por CEP
`GET /api/v1/localidades/vizinhas-agil?cep=01001000&raio=5`

**Resposta:**
```json
{
  "zipcodeInfo": {
    "cep": "01001000",
    "address_type": "Praça",
    "address_name": "da Sé",
    "address": "Praça da Sé",
    "city": "São Paulo",
    "city_ibge": "3550308",
    "lat": "-23.5502784",
    "lng": "-46.6342179",
    "district": "Sé",
    "distanciaKm": null

  },
  "localidade": {
    "codigoIbge": "3550308",
    "municipio": "São Paulo",
    "uf": "SP"
  },
  "cidadesProximas": [
    {
      "cep": "01002000",
      "cidade": "São Paulo",
      "ibge": "3550308",
      "distrito": "Sé",
      "subdistrito": "Sede",
      "distanciaKm": 0.5
    }
  ]
}
```

## ⚙️ Otimizações Implementadas

- **Regex Performance**: Uso de `Pattern` pré-compilado com `[^0-9]` para limpeza de strings, evitando recompilações em tempo de execução.
- **AOT Hints**: Registro manual de reflexão para Records e Entidades JPA para compatibilidade total com GraalVM.
- **Volume Mapping**: Preparado para Cloud Native, desacoplando os dados estáticos (CSV) do binário da aplicação.

## 📄 Licença
Distribuído sob a licença MIT. Veja `LICENSE` para mais informações.

---
**Desenvolvido como solução de alta performance para sistemas de logística e geolocalização.**


### README.md - parte 2

Para completar o seu `README.md` com o rigor técnico que um projeto em **Java 21** e **GraalVM** exige, adicionei a seção de estrutura de pastas. Esta estrutura segue o padrão **Maven** aliado às necessidades de configuração do **Kubernetes (K8S)** e do processamento de arquivos externos.

Adicione o seguinte conteúdo ao seu arquivo:

---

## 📂 Estrutura do Projeto

Abaixo está a árvore de diretórios com a função de cada componente principal:

```text
.
├── src/main/java/com/fbso/ms-geolocalidade/
│   ├── config/
│   │   ├── ClientConfig.java        # Configuração do RestClient (Java 21) e beans de HTTP.
│   │   ├── BatchConfig.java         # Definição de Jobs e Steps do Spring Batch.
│   │   └── IngestaoHints.java       # RuntimeHints para compatibilidade com GraalVM AOT.
│   ├── controller/
│   │   └── LocalidadeController.java # Endpoints REST para consulta de CEP e vizinhança.
│   ├── dto/                         # Records Java 21 para transporte de dados (imutáveis).
│   │   ├── AwesomeCepDTO.java       # Mapeamento da API externa.
│   │   └── RespostaCompletaDTO.java # Resposta consolidada para o cliente final.
│   ├── entity/                      # Entidades JPA com suporte a Hibernate Spatial.
│   │   ├── Municipio.java           # Mapeamento da tabela de municípios (IBGE).
│   │   └── Subdistrito.java         # Mapeamento da tabela de subdistritos (IBGE).
│   ├── repository/                  # Interfaces Spring Data JPA para persistência.
│   ├── service/                     # Regras de negócio e orquestração.
│   │   ├── AwesomeCepService.java   # Integração segura (Token/Key) com a API externa.
│   │   ├── LocalidadeOrchestrator.java # Lógica de união (API + Dados Locais IBGE).
│   │   └── CepUtils.java            # Utilitário com Pattern pré-compilado para limpeza de CEP.
│   └── GeoApplication.java          # Classe principal (Entry point).
├── src/main/resources/
│   └── application.yml              # Configurações de ambiente, tokens e caminhos de volumes.
├── Dockerfile                       # Build multi-stage para gerar imagem nativa GraalVM.
├── pom.xml                          # Gerenciamento de dependências e plugin nativo do Spring.
└── k8s/                             # Manifestos de Deployment e PVC para produção.
```

### Detalhes dos Componentes Chave:

* **`CepUtils.java`**: Implementa a limpeza de strings utilizando o `Pattern` pré-compilado `[^0-9]`, garantindo que apenas dígitos cheguem às APIs, otimizando o uso de CPU.
* **`IngestaoHints.java`**: Arquivo vital para o **GraalVM**. Ele informa ao compilador nativo quais classes serão acessadas via reflexão pelo Jackson (JSON) e pelo Hibernate, evitando erros em tempo de execução na imagem nativa.
* **`application.yml`**: Centraliza a configuração do `APP_IMPORT_PATH`, que aponta para o volume montado no K8S, permitindo que o Spring Batch localize os CSVs do IBGE sem que eles estejam embutidos no binário.
* **`RestClient`**: Utiliza a nova interface fluida do Spring Framework 6+ para chamadas síncronas, substituindo o RestTemplate com melhor suporte a tipos genéricos e Records.

---

