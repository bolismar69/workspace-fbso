package com.fbso.geolocalidade.controller;

import com.fbso.geolocalidade.dto.PageResponseDTO;
import com.fbso.geolocalidade.service.MunicipioService;
import com.fbso.geolocalidade.service.UfService;
import jakarta.validation.constraints.NotBlank;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Pageable;
import org.springframework.data.domain.Sort;
import org.springframework.data.web.PageableDefault;
import org.springframework.http.ResponseEntity;
import org.springframework.validation.annotation.Validated;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@Validated
@RestController
@RequestMapping("/api/v1/localidades/uf")
public class UfController {

  private final MunicipioService municipioService;
  private final UfService ufService;

  public UfController(MunicipioService municipioService, UfService ufService) {
    this.municipioService = municipioService;
    this.ufService = ufService;
  }

  @GetMapping
  public ResponseEntity<PageResponseDTO<Object>> buscarUfs() {
    return ResponseEntity.ok(PageResponseDTO.successList(ufService.buscarUfs()));
  }

  @GetMapping("/{sigla_uf}/municipios")
  public ResponseEntity<PageResponseDTO<Object>> buscarMunicipiosPorUf(
      @PathVariable("sigla_uf") @NotBlank String siglaUf,
      @PageableDefault(size = 10) Pageable pageable) {

    int requestedSize = pageable.getPageSize();
    int size = Math.min(Math.max(requestedSize, 10), 50);

    var pageRequest = PageRequest.of(
        pageable.getPageNumber(),
        size,
        Sort.by("id").ascending());

    return ResponseEntity.ok(PageResponseDTO.ofObject(
      municipioService.buscarMunicipiosPorUfSigla(siglaUf, pageRequest)));
  }

  @GetMapping("/{sigla_uf}/municipios/{nomeMunicipio}")
  public ResponseEntity<PageResponseDTO<Object>> buscarMunicipiosPorUfENomeLike(
      @PathVariable("sigla_uf") @NotBlank String siglaUf,
      @PathVariable @NotBlank String nomeMunicipio,
      @PageableDefault(size = 10) Pageable pageable) {

    int requestedSize = pageable.getPageSize();
    int size = Math.min(Math.max(requestedSize, 10), 50);

    var pageRequest = PageRequest.of(
        pageable.getPageNumber(),
        size,
        Sort.by("id").ascending());

    return ResponseEntity.ok(PageResponseDTO.ofObject(
      municipioService.buscarMunicipiosPorUfSiglaENomeLike(siglaUf, nomeMunicipio, pageRequest)));
  }
}
