package com.fbso.geolocalidade.service;

import com.fbso.geolocalidade.dto.MunicipioDTO.UfDTO;
import com.fbso.geolocalidade.repository.UfRepository;
import java.util.List;
import org.springframework.data.domain.Sort;
import org.springframework.stereotype.Service;

@Service
public class UfService {

  private final UfRepository ufRepository;

  public UfService(UfRepository ufRepository) {
    this.ufRepository = ufRepository;
  }

  public List<UfDTO> buscarUfs() {
    return ufRepository.findAll(Sort.by("sigla").ascending()).stream()
        .map(uf -> new UfDTO(uf.getId(), uf.getSigla(), uf.getNome()))
        .toList();
  }
}
