package com.fbso.geolocalidade.service;

import com.fbso.geolocalidade.dto.MunicipioDTO;
import com.fbso.geolocalidade.dto.MunicipioDTO.DistritoDTO;
import com.fbso.geolocalidade.dto.MunicipioDTO.LocalidadeDTO;
import com.fbso.geolocalidade.dto.MunicipioDTO.MunicipioInfoDTO;
import com.fbso.geolocalidade.dto.MunicipioDTO.RegiaoImediataDTO;
import com.fbso.geolocalidade.dto.MunicipioDTO.RegiaoIntermediariaDTO;
import com.fbso.geolocalidade.dto.MunicipioDTO.SubdistritoDTO;
import com.fbso.geolocalidade.dto.MunicipioDTO.UfDTO;
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

  private List<MunicipioDTO> mapMunicipios(List<Municipio> municipios) {
    List<String> municipioIds = municipios.stream().map(Municipio::getId).toList();

    Map<String, Distrito> primeiroDistritoPorMunicipioId = new HashMap<>();
    if (!municipioIds.isEmpty()) {
      for (Distrito distrito : distritoRepository.findByMunicipioIdInOrderByIdAsc(municipioIds)) {
        String municipioId = distrito.getMunicipio().getId();
        primeiroDistritoPorMunicipioId.putIfAbsent(municipioId, distrito);
      }
    }

    List<String> distritoIds = primeiroDistritoPorMunicipioId.values().stream().map(Distrito::getId).toList();

    Map<String, Subdistrito> primeiroSubdistritoPorDistritoId = new HashMap<>();
    if (!distritoIds.isEmpty()) {
      for (Subdistrito subdistrito : subdistritoRepository.findByDistritoIdInOrderByIdAsc(distritoIds)) {
        String distritoId = subdistrito.getDistrito().getId();
        primeiroSubdistritoPorDistritoId.putIfAbsent(distritoId, subdistrito);
      }
    }

    List<MunicipioDTO> result = new ArrayList<>(municipios.size());
    for (Municipio municipio : municipios) {
      Distrito distrito = primeiroDistritoPorMunicipioId.get(municipio.getId());
      Subdistrito subdistrito = distrito == null ? null : primeiroSubdistritoPorDistritoId.get(distrito.getId());

      result.add(new MunicipioDTO(new LocalidadeDTO(
          toUfDTO(municipio),
          toRegiaoIntermediariaDTO(municipio),
          toRegiaoImediataDTO(municipio),
          new MunicipioInfoDTO(municipio.getId(), municipio.getCodigo(), municipio.getNome()),
          distrito == null ? null : new DistritoDTO(distrito.getId(), distrito.getCodigo(), distrito.getNome()),
          subdistrito == null ? null : new SubdistritoDTO(subdistrito.getId(), subdistrito.getCodigo(), subdistrito.getNome())
      )));
    }

    return result;
  }

  private static UfDTO toUfDTO(Municipio municipio) {
    Uf uf = getUf(municipio);
    if (uf == null) {
      return null;
    }
    return new UfDTO(uf.getId(), uf.getSigla(), uf.getNome());
  }

  private static RegiaoIntermediariaDTO toRegiaoIntermediariaDTO(Municipio municipio) {
    RegiaoIntermediaria ri = getRegiaoIntermediaria(municipio);
    if (ri == null) {
      return null;
    }
    return new RegiaoIntermediariaDTO(ri.getId(), ri.getNome());
  }

  private static RegiaoImediataDTO toRegiaoImediataDTO(Municipio municipio) {
    RegiaoImediata r = municipio.getRegiaoImediata();
    if (r == null) {
      return null;
    }
    return new RegiaoImediataDTO(r.getId(), r.getNome());
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
