package br.com.estudo.cnpj.api;

import static io.restassured.RestAssured.given;
import static org.hamcrest.Matchers.*;

import io.quarkus.test.junit.QuarkusTest;
import org.junit.jupiter.api.Test;

@QuarkusTest
class CnpjResourceTest {

    @Test
    void validateGetShouldReturnValidTrue() {
        given()
                .queryParam("value", "04.252.011/0001-10")
        .when()
                .get("/cnpj/validate")
        .then()
                .statusCode(200)
                .body("valid", is(true))
                .body("normalized", is("04252011000110"));
    }

    @Test
    void validatePostShouldReturnValidFalse() {
        given()
                .contentType("application/json")
                .body("{\"cnpj\":\"04.252.011/0001-11\"}")
        .when()
                .post("/cnpj/validate")
        .then()
                .statusCode(200)
                .body("valid", is(false))
                .body("normalized", is("04252011000111"));
    }
}
