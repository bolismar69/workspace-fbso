package com.fbso.geolocalidade.controller;

import com.fbso.geolocalidade.dto.RespostaCompletaDTO;
import com.fbso.geolocalidade.service.LocalidadeOrchestrator;
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

  private final LocalidadeOrchestrator orchestrator;

  public LocalidadeController(LocalidadeOrchestrator orchestrator) {
    this.orchestrator = orchestrator;
  }

  @GetMapping("/vizinhas-agil")
  public ResponseEntity<RespostaCompletaDTO> buscarVizinhasAwesome(
      @RequestParam @NotBlank String cep,
      @RequestParam(defaultValue = "5") @Positive Double raio) {
    return ResponseEntity.ok(orchestrator.processarBuscaPorCep(cep, raio));
  }
}
