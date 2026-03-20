package com.fbso.geolocalidade.controller;

import com.fbso.geolocalidade.dto.AwesomeCepDTO;
import com.fbso.geolocalidade.dto.PageResponseDTO;
import com.fbso.geolocalidade.exception.AwesomeApiException;
import com.fbso.geolocalidade.service.AwesomeCepService;
import jakarta.validation.constraints.NotBlank;
import java.util.List;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.PageRequest;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@Validated
@RestController
@RequestMapping("/api/v1/localidades")
public class CepController {

  private final AwesomeCepService awesomeCepService;

  public CepController(AwesomeCepService awesomeCepService) {
    this.awesomeCepService = awesomeCepService;
  }

  @GetMapping("/cep/{cep}")
  public ResponseEntity<PageResponseDTO<Object>> buscarCepAwesome(@PathVariable @NotBlank String cep) {
    AwesomeCepDTO dto = awesomeCepService.obterCoordenadas(cep);

    var content = (dto == null) ? List.<AwesomeCepDTO>of() : List.of(dto);
    var page = new PageImpl<>(content, PageRequest.of(0, 1), content.size());
    return ResponseEntity.ok(PageResponseDTO.ofObject(page));
  }
}
