package br.com.estudo.cnpj.api;

import br.com.estudo.cnpj.CnpjValidator;
import br.com.estudo.cnpj.api.dto.CnpjValidationRequest;
import br.com.estudo.cnpj.api.dto.CnpjValidationResponse;
import jakarta.ws.rs.Consumes;
import jakarta.ws.rs.GET;
import jakarta.ws.rs.POST;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.QueryParam;
import jakarta.ws.rs.core.MediaType;

@Path("/cnpj")
@Produces(MediaType.APPLICATION_JSON)
@Consumes(MediaType.APPLICATION_JSON)
public class CnpjResource {

    @GET
    @Path("/validate")
    public CnpjValidationResponse validateGet(@QueryParam("value") String value) {
        String normalized = CnpjValidator.normalize(value);
        boolean valid = CnpjValidator.isValid(value);
        return new CnpjValidationResponse(value, normalized, valid);
    }

    @POST
    @Path("/validate")
    public CnpjValidationResponse validatePost(CnpjValidationRequest request) {
        String input = (request == null) ? null : request.cnpj();
        String normalized = CnpjValidator.normalize(input);
        boolean valid = CnpjValidator.isValid(input);
        return new CnpjValidationResponse(input, normalized, valid);
    }
}
