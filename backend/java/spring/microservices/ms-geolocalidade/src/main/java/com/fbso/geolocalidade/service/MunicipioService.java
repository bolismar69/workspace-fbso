package com.fbso.geolocalidade.service;

import com.fbso.geolocalidade.dto.MunicipioDTO;
import com.fbso.geolocalidade.dto.MunicipioDTO.DistritoDTO;
import com.fbso.geolocalidade.dto.MunicipioDTO.SubdistritoDTO;
import com.fbso.geolocalidade.entity.Distrito;
import com.fbso.geolocalidade.entity.Municipio;
import com.fbso.geolocalidade.entity.RegiaoImediata;
import com.fbso.geolocalidade.entity.RegiaoIntermediaria;
import com.fbso.geolocalidade.entity.Subdistrito;
import com.fbso.geolocalidade.entity.Uf;
import com.fbso.geolocalidade.repository.DistritoRepository;
import com.fbso.geolocalidade.repository.MunicipioRepository;
import com.fbso.geolocalidade.repository.SubdistritoRepository;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.PageImpl;
import org.springframework.data.domain.Pageable;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

@Service
public class MunicipioService {

  private final MunicipioRepository municipioRepository;
  private final DistritoRepository distritoRepository;
  private final SubdistritoRepository subdistritoRepository;

  public MunicipioService(
      MunicipioRepository municipioRepository,
      DistritoRepository distritoRepository,
      SubdistritoRepository subdistritoRepository) {
    this.municipioRepository = municipioRepository;
    this.distritoRepository = distritoRepository;
    this.subdistritoRepository = subdistritoRepository;
  }

  @Transactional(readOnly = true)
  public Page<MunicipioDTO> buscarMunicipios(Pageable pageable) {
    Page<Municipio> page = municipioRepository.findAllByOrderByIdAsc(pageable);
    List<MunicipioDTO> result = mapMunicipios(page.getContent());
    return new PageImpl<>(result, pageable, page.getTotalElements());
  }

  @Transactional(readOnly = true)
  public Page<MunicipioDTO> buscarMunicipiosPorNome(String nome, Pageable pageable) {
    Page<Municipio> page = municipioRepository.findByNomeContainingIgnoreCaseOrderByIdAsc(nome, pageable);
    List<MunicipioDTO> result = mapMunicipios(page.getContent());
    return new PageImpl<>(result, pageable, page.getTotalElements());
  }

  @Transactional(readOnly = true)
  public Page<MunicipioDTO> buscarMunicipiosPorUfSigla(String siglaUf, Pageable pageable) {
    Page<Municipio> page = municipioRepository
        .findByRegiaoImediata_RegiaoIntermediaria_Uf_SiglaIgnoreCaseOrderByIdAsc(siglaUf, pageable);
    List<MunicipioDTO> result = mapMunicipios(page.getContent());
    return new PageImpl<>(result, pageable, page.getTotalElements());
  }

  @Transactional(readOnly = true)
  public Page<MunicipioDTO> buscarMunicipiosPorUfSiglaENomeLike(String siglaUf, String nome, Pageable pageable) {
    Page<Municipio> page = municipioRepository
        .findByRegiaoImediata_RegiaoIntermediaria_Uf_SiglaIgnoreCaseAndNomeContainingIgnoreCaseOrderByIdAsc(
            siglaUf,
            nome,
            pageable);
    List<MunicipioDTO> result = mapMunicipios(page.getContent());
    return new PageImpl<>(result, pageable, page.getTotalElements());
  }

  @Transactional(readOnly = true)
  public Optional<MunicipioDTO> buscarMunicipioPorId(String id) {
    return municipioRepository.findById(id)
        .map(m -> mapMunicipios(List.of(m)))
        .filter(list -> !list.isEmpty())
        .map(list -> list.get(0));
  }

  @Transactional(readOnly = true)
  public Optional<MunicipioDTO> buscarMunicipioPorUfSiglaEId(String siglaUf, String id) {
    return municipioRepository.findByIdAndRegiaoImediata_RegiaoIntermediaria_Uf_SiglaIgnoreCase(id, siglaUf)
        .map(m -> mapMunicipios(List.of(m)))
        .filter(list -> !list.isEmpty())
        .map(list -> list.get(0));
  }

  private List<MunicipioDTO> mapMunicipios(List<Municipio> municipios) {
    List<String> municipioIds = municipios.stream().map(Municipio::getId).toList();

    Map<String, List<Distrito>> distritosPorMunicipioId = new HashMap<>();
    List<Distrito> distritos = municipioIds.isEmpty()
        ? List.of()
        : distritoRepository.findByMunicipioIdInOrderByIdAsc(municipioIds);

    for (Distrito distrito : distritos) {
      String municipioId = distrito.getMunicipio().getId();
      distritosPorMunicipioId.computeIfAbsent(municipioId, k -> new ArrayList<>()).add(distrito);
    }

    List<String> distritoIds = distritos.stream().map(Distrito::getId).toList();

    Map<String, List<Subdistrito>> subdistritosPorDistritoId = new HashMap<>();
    List<Subdistrito> subdistritos = distritoIds.isEmpty()
        ? List.of()
        : subdistritoRepository.findByDistritoIdInOrderByIdAsc(distritoIds);

    for (Subdistrito subdistrito : subdistritos) {
      String distritoId = subdistrito.getDistrito().getId();
      subdistritosPorDistritoId.computeIfAbsent(distritoId, k -> new ArrayList<>()).add(subdistrito);
    }

    List<MunicipioDTO> result = new ArrayList<>(municipios.size());
    for (Municipio municipio : municipios) {
      List<DistritoDTO> distritoDTOs = distritosPorMunicipioId.getOrDefault(municipio.getId(), List.of()).stream()
        .map(d -> {
        List<SubdistritoDTO> subdistritoDTOs = subdistritosPorDistritoId.getOrDefault(d.getId(), List.of()).stream()
          .map(s -> new SubdistritoDTO(s.getId(), s.getCodigo(), s.getNome()))
          .toList();

        return new DistritoDTO(d.getId(), d.getCodigo(), d.getNome(), subdistritoDTOs);
        })
        .toList();

        Uf uf = getUf(municipio);
        RegiaoIntermediaria regiaoIntermediaria = getRegiaoIntermediaria(municipio);
        RegiaoImediata regiaoImediata = municipio.getRegiaoImediata();

      result.add(new MunicipioDTO(
        municipio.getId(),
        municipio.getNome(),
        municipio.getCodigo(),
        uf == null ? null : uf.getId(),
        uf == null ? null : uf.getNome(),
        uf == null ? null : uf.getSigla(),
        regiaoIntermediaria == null ? null : regiaoIntermediaria.getId(),
        regiaoIntermediaria == null ? null : regiaoIntermediaria.getNome(),
        regiaoImediata == null ? null : regiaoImediata.getId(),
        regiaoImediata == null ? null : regiaoImediata.getNome(),
        distritoDTOs));
    }

    return result;
  }

  private static RegiaoIntermediaria getRegiaoIntermediaria(Municipio municipio) {
    RegiaoImediata r = municipio.getRegiaoImediata();
    return r == null ? null : r.getRegiaoIntermediaria();
  }

  private static Uf getUf(Municipio municipio) {
    RegiaoIntermediaria ri = getRegiaoIntermediaria(municipio);
    return ri == null ? null : ri.getUf();
  }
}
