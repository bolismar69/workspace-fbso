package br.com.fbso.geolocalidade.processor;

import br.com.fbso.geolocalidade.dto.DistritoCsvDTO;
import br.com.fbso.geolocalidade.entity.Distrito;
import br.com.fbso.geolocalidade.entity.Municipio;
import org.springframework.batch.item.ItemProcessor;
import org.springframework.stereotype.Component;

@Component
public class DistritoProcessor implements ItemProcessor<DistritoCsvDTO, Distrito> {

    @Override
    public Distrito process(DistritoCsvDTO item) {
        // Estratégia de Proxy: Criamos o Município apenas com o ID 
        // para satisfazer a FK sem precisar de um Select no banco.
        Municipio municipioProxy = new Municipio(
            item.municipioId(), 
            null, // código não necessário para o Proxy
            null, // nome não necessário para o Proxy
            null  // região não necessária para o Proxy
        );

        return new Distrito(
            item.distritoIdCompleto(),
            item.distritoCodCurto(),
            item.distritoNome(),
            municipioProxy
        );
    }
}