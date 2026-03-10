package br.com.estudo.cnpj.api;

import static org.junit.jupiter.api.Assertions.*;

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

import org.junit.jupiter.api.Test;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.boot.test.web.server.LocalServerPort;

@SpringBootTest(webEnvironment = SpringBootTest.WebEnvironment.RANDOM_PORT)
class CnpjResourceTest {

    @LocalServerPort
    private int port;

    @Test
    void validateGetShouldReturnValidTrue() throws Exception {
        HttpClient client = HttpClient.newHttpClient();
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create("http://localhost:" + port + "/cnpj/validate?value=04.252.011/0001-10"))
                .GET()
                .build();

        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());

        assertEquals(200, response.statusCode());
        assertTrue(response.body().contains("\"valid\":true"));
        assertTrue(response.body().contains("\"normalized\":\"04252011000110\""));
    }

    @Test
    void validatePostShouldReturnValidFalse() throws Exception {
        HttpClient client = HttpClient.newHttpClient();
        String json = "{\"cnpj\":\"04.252.011/0001-11\"}";
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create("http://localhost:" + port + "/cnpj/validate"))
                .header("Content-Type", "application/json")
                .POST(HttpRequest.BodyPublishers.ofString(json))
                .build();

        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());

        assertEquals(200, response.statusCode());
        assertTrue(response.body().contains("\"valid\":false"));
        assertTrue(response.body().contains("\"normalized\":\"04252011000111\""));
    }
}
