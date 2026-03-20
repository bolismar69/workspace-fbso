package com.fbso.geolocalidade.controller;

import com.fbso.geolocalidade.dto.MunicipioDTO;
import com.fbso.geolocalidade.dto.PageResponseDTO;
import com.fbso.geolocalidade.service.MunicipioService;
import jakarta.validation.constraints.NotBlank;
import java.util.List;
import org.springframework.data.domain.PageRequest;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
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
@RequestMapping("/api/v1/localidades/municipios")
public class MunicipioController {

  private final MunicipioService municipioService;

  public MunicipioController(MunicipioService municipioService) {
    this.municipioService = municipioService;
  }

  @GetMapping
  public ResponseEntity<PageResponseDTO<Object>> buscarMunicipios(@PageableDefault(size = 10) Pageable pageable) {
    int requestedSize = pageable.getPageSize();
    int size = Math.min(Math.max(requestedSize, 1), 50);

    var pageRequest = PageRequest.of(
        pageable.getPageNumber(),
        size,
        Sort.by("id").ascending());

    return ResponseEntity.ok(PageResponseDTO.ofObject(municipioService.buscarMunicipios(pageRequest)));
  }

  @GetMapping("/id/{id:\\d+}")
  public ResponseEntity<PageResponseDTO<Object>> buscarMunicipioPorId(@PathVariable @NotBlank String id) {
    var pageRequest = PageRequest.of(0, 1, Sort.by("id").ascending());
    var municipio = municipioService.buscarMunicipioPorId(id).orElse(null);

    Page<MunicipioDTO> page = (municipio == null)
        ? new PageImpl<>(List.of(), pageRequest, 0)
        : new PageImpl<>(List.of(municipio), pageRequest, 1);

    return ResponseEntity.ok(PageResponseDTO.ofObject(page));
  }

  @GetMapping("/nome/{nomeMunicipio:[^0-9].*}")
  public ResponseEntity<PageResponseDTO<Object>> buscarMunicipiosPorNome(
      @PathVariable @NotBlank String nomeMunicipio,
      @PageableDefault(size = 10) Pageable pageable) {
    int requestedSize = pageable.getPageSize();
    int size = Math.min(Math.max(requestedSize, 1), 50);

    var pageRequest = PageRequest.of(
        pageable.getPageNumber(),
        size,
        Sort.by("id").ascending());

    return ResponseEntity.ok(PageResponseDTO.ofObject(
        municipioService.buscarMunicipiosPorNome(nomeMunicipio, pageRequest)));
  }
}
