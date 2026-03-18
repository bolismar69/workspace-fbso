package com.fbso.geolocalidade.repository;

import com.fbso.geolocalidade.entity.Subdistrito;
import java.util.Optional;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

public interface SubdistritoRepository extends JpaRepository<Subdistrito, String> {

  @Query("select s.nomeSubdistrito from Subdistrito s where s.codigoSubdistrito11 like concat(:codigoIBGE, '%')")
  Optional<String> findNomeByCodigo(@Param("codigoIBGE") String codigoIBGE);
}
