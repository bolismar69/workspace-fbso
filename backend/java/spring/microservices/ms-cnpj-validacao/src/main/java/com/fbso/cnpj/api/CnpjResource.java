package br.com.estudo.cnpj.api;

import br.com.estudo.cnpj.CnpjValidator;
import br.com.estudo.cnpj.api.dto.CnpjValidationRequest;
import br.com.estudo.cnpj.api.dto.CnpjValidationResponse;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/cnpj")
public class CnpjResource {

    @GetMapping("/validate")
    public CnpjValidationResponse validateGet(@RequestParam("value") String value) {
        String normalized = CnpjValidator.normalize(value);
        boolean valid = CnpjValidator.isValid(value);
        return new CnpjValidationResponse(value, normalized, valid);
    }

    @PostMapping("/validate")
    public CnpjValidationResponse validatePost(@RequestBody(required = false) CnpjValidationRequest request) {
        String input = (request == null) ? null : request.cnpj();
        String normalized = CnpjValidator.normalize(input);
        boolean valid = CnpjValidator.isValid(input);
        return new CnpjValidationResponse(input, normalized, valid);
    }
}
