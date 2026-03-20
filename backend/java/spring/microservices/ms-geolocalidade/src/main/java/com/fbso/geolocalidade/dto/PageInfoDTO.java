package com.fbso.geolocalidade.dto;

import org.springframework.data.domain.Page;

public record PageInfoDTO(
    int page,
    int size,
    long totalElements,
    int totalPages,
    boolean first,
    boolean last,
    boolean empty,
    int numberOfElements,
    String sortDirection
) {
    public static PageInfoDTO of(Page<?> page) {
        return new PageInfoDTO(
            page.getNumber(),
            page.getSize(),
            page.getTotalElements(),
            page.getTotalPages(),
            page.isFirst(),
            page.isLast(),
            page.isEmpty(),
            page.getNumberOfElements(),
            page.getSort().isSorted()
                ? page.getSort().iterator().next().getDirection().name()
                : "NONE"
        );
    }
}
