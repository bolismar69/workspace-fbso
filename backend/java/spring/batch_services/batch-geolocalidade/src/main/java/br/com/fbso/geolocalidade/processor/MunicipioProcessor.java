package br.com.fbso.geolocalidade.processor;

import br.com.fbso.geolocalidade.dto.MunicipioCsvDTO;
import br.com.fbso.geolocalidade.entity.RegiaoImediata;
import br.com.fbso.geolocalidade.entity.RegiaoIntermediaria;
import br.com.fbso.geolocalidade.entity.Uf;
import br.com.fbso.geolocalidade.entity.Municipio;
import org.springframework.batch.item.ItemProcessor;
import org.springframework.stereotype.Component;
import java.util.HashMap;
import java.util.Map;

@Component
public class MunicipioProcessor implements ItemProcessor<MunicipioCsvDTO, Municipio> {

    // Cache simples para evitar recriar instâncias idênticas no mesmo Job
    private final Map<String, Uf> ufCache = new HashMap<>();
    private final Map<String, RegiaoIntermediaria> interCache = new HashMap<>();
    private final Map<String, RegiaoImediata> imedCache = new HashMap<>();
    
    // Mapa auxiliar para as siglas que não constam no CSV do IBGE
    private static final Map<String, String> UF_SIGLAS = Map.ofEntries(
        Map.entry("11", "RO"), Map.entry("12", "AC"), Map.entry("13", "AM"),
        Map.entry("14", "RR"), Map.entry("15", "PA"), Map.entry("16", "AP"),
        Map.entry("17", "TO"), Map.entry("21", "MA"), Map.entry("22", "PI"),
        Map.entry("23", "CE"), Map.entry("24", "RN"), Map.entry("25", "PB"),
        Map.entry("26", "PE"), Map.entry("27", "AL"), Map.entry("28", "SE"),
        Map.entry("29", "BA"), Map.entry("31", "MG"), Map.entry("32", "ES"),
        Map.entry("33", "RJ"), Map.entry("35", "SP"), Map.entry("41", "PR"),
        Map.entry("42", "SC"), Map.entry("43", "RS"), Map.entry("50", "MS"),
        Map.entry("51", "MT"), Map.entry("52", "GO"), Map.entry("53", "DF")
    );

    @Override
    public Municipio process(MunicipioCsvDTO item) {
        // 1. Criar UF
        Uf uf = ufCache.computeIfAbsent(item.ufId(), id -> 
            new Uf(id, UF_SIGLAS.get(id), item.ufNome()));

        // 2. Criar Região Intermediária
        RegiaoIntermediaria intermediaria = interCache.computeIfAbsent(item.regiaoInterId(), id -> 
            new RegiaoIntermediaria(id, item.regiaoInterNome(), uf));

        // 3. Criar Região Imediata
        RegiaoImediata imediata = imedCache.computeIfAbsent(item.regiaoImedId(), id -> 
            new RegiaoImediata(id, item.regiaoImedNome(), intermediaria));

        // 4. Retornar Município (o Writer cuidará da persistência)
        return new Municipio(
            item.municipioIdCompleto(),
            item.municipioCodCurto(),
            item.municipioNome(),
            imediata
        );
    }
}