package br.com.fbso.geolocalidade.processor;

import br.com.fbso.geolocalidade.dto.SubdistritoCsvDTO;
import br.com.fbso.geolocalidade.entity.Distrito;
import br.com.fbso.geolocalidade.entity.Subdistrito;
import org.springframework.batch.item.ItemProcessor;
import org.springframework.stereotype.Component;

@Component
public class SubdistritoProcessor implements ItemProcessor<SubdistritoCsvDTO, Subdistrito> {

    @Override
    public Subdistrito process(SubdistritoCsvDTO item) {
        // Estratégia de Proxy: Associamos ao Distrito usando apenas o ID completo (9 dígitos)
        Distrito distritoProxy = new Distrito(
            item.distritoId(),
            null,
            null,
            null
        );

        return new Subdistrito(
            item.subdistritoIdCompleto(),
            item.subdistritoCodCurto(),
            item.subdistritoNome(),
            distritoProxy
        );
    }
}