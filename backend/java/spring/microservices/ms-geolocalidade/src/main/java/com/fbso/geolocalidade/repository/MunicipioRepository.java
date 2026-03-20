package com.fbso.geolocalidade.repository;

import com.fbso.geolocalidade.entity.Municipio;
import java.util.List;
import java.util.Optional;
import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.data.jpa.repository.EntityGraph;
import org.springframework.data.domain.Page;
import org.springframework.data.domain.Pageable;

public interface MunicipioRepository extends JpaRepository<Municipio, String> {

	@Override
	@EntityGraph(attributePaths = {
			"regiaoImediata",
			"regiaoImediata.regiaoIntermediaria",
			"regiaoImediata.regiaoIntermediaria.uf"
	})
	Optional<Municipio> findById(String id);

	@EntityGraph(attributePaths = {
			"regiaoImediata",
			"regiaoImediata.regiaoIntermediaria",
			"regiaoImediata.regiaoIntermediaria.uf"
	})
	List<Municipio> findAllByOrderByIdAsc();

	@EntityGraph(attributePaths = {
			"regiaoImediata",
			"regiaoImediata.regiaoIntermediaria",
			"regiaoImediata.regiaoIntermediaria.uf"
	})
	Page<Municipio> findAllByOrderByIdAsc(Pageable pageable);

	@EntityGraph(attributePaths = {
			"regiaoImediata",
			"regiaoImediata.regiaoIntermediaria",
			"regiaoImediata.regiaoIntermediaria.uf"
	})
	Page<Municipio> findByNomeContainingIgnoreCaseOrderByIdAsc(String nome, Pageable pageable);

	@EntityGraph(attributePaths = {
			"regiaoImediata",
			"regiaoImediata.regiaoIntermediaria",
			"regiaoImediata.regiaoIntermediaria.uf"
	})
	Page<Municipio> findByRegiaoImediata_RegiaoIntermediaria_Uf_SiglaIgnoreCaseOrderByIdAsc(
			String sigla,
			Pageable pageable);

	@EntityGraph(attributePaths = {
			"regiaoImediata",
			"regiaoImediata.regiaoIntermediaria",
			"regiaoImediata.regiaoIntermediaria.uf"
	})
	Page<Municipio> findByRegiaoImediata_RegiaoIntermediaria_Uf_SiglaIgnoreCaseAndNomeContainingIgnoreCaseOrderByIdAsc(
			String sigla,
			String nome,
			Pageable pageable);
}
