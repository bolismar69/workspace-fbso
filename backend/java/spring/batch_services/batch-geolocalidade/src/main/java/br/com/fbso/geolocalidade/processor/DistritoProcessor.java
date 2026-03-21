package br.com.fbso.geolocalidade.processor;

import br.com.fbso.geolocalidade.dto.DistritoCsvDTO;
import br.com.fbso.geolocalidade.entity.Distrito;
import br.com.fbso.geolocalidade.entity.Municipio;
import br.com.fbso.geolocalidade.repository.MunicipioRepository;
import org.springframework.batch.item.ItemProcessor;
import org.springframework.stereotype.Component;

@Component
public class DistritoProcessor implements ItemProcessor<DistritoCsvDTO, Distrito> {

    private final MunicipioRepository municipioRepository;

    public DistritoProcessor(MunicipioRepository municipioRepository) {
        this.municipioRepository = municipioRepository;
    }

    @Override
    public Distrito process(DistritoCsvDTO item) {
        // Referência gerenciada pelo JPA (não faz INSERT; evita violar NOT NULL em campos do Município)
        Municipio municipioRef = municipioRepository.getReferenceById(item.municipioId());

        return new Distrito(
            item.distritoIdCompleto(),
            item.distritoCodCurto(),
            item.distritoNome(),
            municipioRef
        );
    }
}