package com.fbso.geolocalidade.controller;

import com.fbso.geolocalidade.dto.AwesomeCepDTO;
import com.fbso.geolocalidade.dto.AwesomeCepResponseDTO;
import com.fbso.geolocalidade.dto.CepInfoDTO;
import com.fbso.geolocalidade.dto.PageInfoDTO;
import com.fbso.geolocalidade.dto.SingleResponseDTO;
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
  public ResponseEntity<SingleResponseDTO<Object>> buscarCepAwesome(@PathVariable @NotBlank String cep) {
    AwesomeCepDTO dto = awesomeCepService.obterCoordenadas(cep);

    var content = (dto == null) ? List.<AwesomeCepDTO>of() : List.of(dto);
    var page = new PageImpl<>(content, PageRequest.of(0, 1), content.size());

    CepInfoDTO payload = (dto == null) ? null : new CepInfoDTO(AwesomeCepResponseDTO.from(dto));
    return ResponseEntity.ok(SingleResponseDTO.of(payload, PageInfoDTO.of(page)));
  }
}
