package com.fbso.geolocalidade.controller;

import com.fbso.geolocalidade.dto.MunicipioDTO;
import com.fbso.geolocalidade.dto.PageInfoDTO;
import com.fbso.geolocalidade.dto.PageResponseDTO;
import com.fbso.geolocalidade.dto.ResponseStatusDTO;
import com.fbso.geolocalidade.dto.UfHierarquiaDTO;
import com.fbso.geolocalidade.service.MunicipioService;
import com.fbso.geolocalidade.service.UfService;
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
@RequestMapping("/api/v1/localidades/uf")
public class UfController {

  private final MunicipioService municipioService;
  private final UfService ufService;

  public UfController(MunicipioService municipioService, UfService ufService) {
    this.municipioService = municipioService;
    this.ufService = ufService;
  }

  @GetMapping
  public ResponseEntity<PageResponseDTO<Object>> buscarUfs(
      @PageableDefault(size = 10) Pageable pageable) {

    int requestedSize = pageable.getPageSize();
    int size = Math.min(Math.max(requestedSize, 10), 50);

    var pageRequest = PageRequest.of(
        pageable.getPageNumber(),
        size,
        Sort.by("sigla").ascending());

    List<UfHierarquiaDTO.UfDTO> allUfs = ufService.buscarUfs().stream()
        .map(uf -> new UfHierarquiaDTO.UfDTO(uf.id(), uf.sigla(), uf.nome(), List.of()))
        .toList();

    int total = allUfs.size();
    int fromIndex = Math.min(pageRequest.getPageNumber() * pageRequest.getPageSize(), total);
    int toIndex = Math.min(fromIndex + pageRequest.getPageSize(), total);

    List<UfHierarquiaDTO.UfDTO> slice = (fromIndex >= toIndex)
        ? List.of()
        : List.copyOf(allUfs.subList(fromIndex, toIndex));

    Page<UfHierarquiaDTO.UfDTO> page = new PageImpl<>(slice, pageRequest, total);
    UfHierarquiaDTO hierarquia = UfHierarquiaDTO.fromUfs(page.getContent());

    PageResponseDTO<Object> response = new PageResponseDTO<>(
        List.of((Object) hierarquia),
        ResponseStatusDTO.success(200, "OK"),
        PageInfoDTO.of(page));

    return ResponseEntity.ok(response);
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

    Page<MunicipioDTO> page = municipioService.buscarMunicipiosPorUfSigla(siglaUf, pageRequest);
    UfHierarquiaDTO hierarquia = UfHierarquiaDTO.fromMunicipios(page.getContent());

    PageResponseDTO<Object> response = new PageResponseDTO<>(
      List.of((Object) hierarquia),
      ResponseStatusDTO.success(200, "OK"),
      PageInfoDTO.of(page));

    return ResponseEntity.ok(response);
  }

  @GetMapping("/{sigla_uf}/municipios/ibge/{codigo-ibge}")
  public ResponseEntity<PageResponseDTO<Object>> buscarMunicipioPorUfECodigoIbge(
      @PathVariable("sigla_uf") @NotBlank String siglaUf,
      @PathVariable("codigo-ibge") @NotBlank String codigoIbge) {

    var pageRequest = PageRequest.of(0, 1, Sort.by("id").ascending());
    var municipio = municipioService.buscarMunicipioPorUfSiglaEId(siglaUf, codigoIbge).orElse(null);

    Page<MunicipioDTO> page = (municipio == null)
        ? new PageImpl<>(List.of(), pageRequest, 0)
        : new PageImpl<>(List.of(municipio), pageRequest, 1);

    UfHierarquiaDTO hierarquia = UfHierarquiaDTO.fromMunicipios(page.getContent());
    PageResponseDTO<Object> response = new PageResponseDTO<>(
      List.of((Object) hierarquia),
      ResponseStatusDTO.success(200, "OK"),
      PageInfoDTO.of(page));

    return ResponseEntity.ok(response);
  }

  @GetMapping("/{sigla_uf}/municipios/nome/{nomeMunicipio}")
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

    Page<MunicipioDTO> page = municipioService.buscarMunicipiosPorUfSiglaENomeLike(siglaUf, nomeMunicipio, pageRequest);
    UfHierarquiaDTO hierarquia = UfHierarquiaDTO.fromMunicipios(page.getContent());

    PageResponseDTO<Object> response = new PageResponseDTO<>(
      List.of((Object) hierarquia),
      ResponseStatusDTO.success(200, "OK"),
      PageInfoDTO.of(page));

    return ResponseEntity.ok(response);
  }
}
