package br.com.fbso.geolocalidade.processor;

import br.com.fbso.geolocalidade.dto.SubdistritoCsvDTO;
import br.com.fbso.geolocalidade.entity.Distrito;
import br.com.fbso.geolocalidade.entity.Subdistrito;
import br.com.fbso.geolocalidade.repository.DistritoRepository;
import org.springframework.batch.item.ItemProcessor;
import org.springframework.stereotype.Component;

@Component
public class SubdistritoProcessor implements ItemProcessor<SubdistritoCsvDTO, Subdistrito> {

    private final DistritoRepository distritoRepository;

    public SubdistritoProcessor(DistritoRepository distritoRepository) {
        this.distritoRepository = distritoRepository;
    }

    @Override
    public Subdistrito process(SubdistritoCsvDTO item) {
        // Referência gerenciada pelo JPA (não faz INSERT; evita violar NOT NULL em campos do Distrito)
        Distrito distritoRef = distritoRepository.getReferenceById(item.distritoId());

        return new Subdistrito(
            item.subdistritoIdCompleto(),
            item.subdistritoCodCurto(),
            item.subdistritoNome(),
            distritoRef
        );
    }
}