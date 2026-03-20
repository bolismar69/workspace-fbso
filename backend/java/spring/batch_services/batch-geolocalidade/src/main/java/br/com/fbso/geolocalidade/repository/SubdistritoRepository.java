package br.com.fbso.geolocalidade.repository;

import br.com.fbso.geolocalidade.entity.Subdistrito;
import java.util.List;
import java.util.Optional;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.Query;
import org.springframework.data.repository.query.Param;

public interface SubdistritoRepository extends JpaRepository<Subdistrito, String> {

  @Query("select s.nome from Subdistrito s where s.id like concat(:codigoIBGE, '%')")
  Optional<String> findNomeByCodigo(@Param("codigoIBGE") String codigoIBGE);

  List<Subdistrito> findByDistritoIdInOrderByIdAsc(List<String> distritoIds);
}
