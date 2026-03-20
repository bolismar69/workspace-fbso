package com.fbso.geolocalidade.dto;

import com.fasterxml.jackson.annotation.JsonIgnoreProperties;
import java.util.List;

@JsonIgnoreProperties(ignoreUnknown = true)
public record AwesomeSearchResponseDTO(
    Integer found,
    List<AwesomeCepDTO> results
) {}
