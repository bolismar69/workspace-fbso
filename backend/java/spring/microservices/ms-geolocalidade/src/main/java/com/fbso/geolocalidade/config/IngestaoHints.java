package com.fbso.geolocalidade.config;

import com.fbso.geolocalidade.dto.AwesomeCepDTO;
import com.fbso.geolocalidade.dto.AwesomeCepResponseDTO;
import com.fbso.geolocalidade.dto.CepInfoDTO;
import com.fbso.geolocalidade.dto.LocalidadeDetalhadaDTO;
import com.fbso.geolocalidade.dto.RespostaCompletaDTO;
import com.fbso.geolocalidade.dto.SingleResponseDTO;
import com.fbso.geolocalidade.dto.VizinhoEnriquecidoDTO;
import com.fbso.geolocalidade.entity.Municipio;
import com.fbso.geolocalidade.entity.Subdistrito;
import org.springframework.aot.hint.RuntimeHints;
import org.springframework.aot.hint.RuntimeHintsRegistrar;
import org.springframework.aot.hint.TypeReference;

public class IngestaoHints implements RuntimeHintsRegistrar {
  @Override
  public void registerHints(RuntimeHints hints, ClassLoader classLoader) {
    hints.reflection().registerTypes(
        java.util.List.of(
            TypeReference.of(AwesomeCepDTO.class),
            TypeReference.of(AwesomeCepResponseDTO.class),
            TypeReference.of(LocalidadeDetalhadaDTO.class),
            TypeReference.of(RespostaCompletaDTO.class),
            TypeReference.of(CepInfoDTO.class),
            TypeReference.of(SingleResponseDTO.class),
            TypeReference.of(VizinhoEnriquecidoDTO.class),
            TypeReference.of(Municipio.class),
            TypeReference.of(Subdistrito.class)
        ),
        builder -> builder.withMembers());
  }
}
