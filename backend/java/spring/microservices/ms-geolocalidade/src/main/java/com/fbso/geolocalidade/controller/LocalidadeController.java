package com.fbso.geolocalidade.controller;

import com.fbso.geolocalidade.dto.PageResponseDTO;
import com.fbso.geolocalidade.dto.RespostaCompletaDTO;
import com.fbso.geolocalidade.service.LocalidadeService;
import jakarta.validation.constraints.NotBlank;
import jakarta.validation.constraints.Positive;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@Validated
@RestController
@RequestMapping("/api/v1/localidades")
public class LocalidadeController {

  private final LocalidadeService localidadeService;

  public LocalidadeController(LocalidadeService localidadeService) {
    this.localidadeService = localidadeService;
  }

  @GetMapping("/ceps-proximos")
  public ResponseEntity<PageResponseDTO<Object>> buscarCepsProximosAwesome(
      @RequestParam @NotBlank String cep,
      @RequestParam(defaultValue = "5") @Positive Double raio) {
    RespostaCompletaDTO dto = localidadeService.processarBuscaPorCepsProximos(cep, raio);
    return ResponseEntity.ok(PageResponseDTO.success(dto));
  }
}
