package com.fbso.geolocalidade.repository;

import com.fbso.geolocalidade.entity.Distrito;
import java.util.List;
import org.springframework.data.jpa.repository.JpaRepository;

public interface DistritoRepository extends JpaRepository<Distrito, String> {
  List<Distrito> findByMunicipioIdInOrderByIdAsc(List<String> municipioIds);
}
