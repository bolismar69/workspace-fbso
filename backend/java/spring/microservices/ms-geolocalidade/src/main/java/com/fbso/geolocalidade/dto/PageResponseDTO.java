package com.fbso.geolocalidade.dto;

import org.springframework.data.domain.Page;
import java.util.List;

public record PageResponseDTO<T>(
    List<T> content,
    ResponseStatusDTO status,
    PageInfoDTO pageInfo
) {
    public static <T> PageResponseDTO<T> of(Page<T> page) {
        return new PageResponseDTO<>(
            page.getContent(), // Aqui vai a sua lista de MunicipioDTO
            ResponseStatusDTO.success(200, "OK"),
            PageInfoDTO.of(page)
        );
    }

    public static PageResponseDTO<Object> ofObject(Page<?> page) {
        return PageResponseDTO.of(page.map(it -> (Object) it));
    }

    public static PageResponseDTO<Object> success(Object item) {
        return new PageResponseDTO<>(
            item == null ? List.of() : List.of(item),
            ResponseStatusDTO.success(200, "OK"),
            null
        );
    }

    public static PageResponseDTO<Object> successList(List<?> items) {
        return new PageResponseDTO<>(
            items == null ? List.of() : List.copyOf(items),
            ResponseStatusDTO.success(200, "OK"),
            null
        );
    }

    // Método para Erros (sem itens e sem paginação)
    public static <T> PageResponseDTO<T> error(int code, String message) {
        return new PageResponseDTO<>(
            List.of(),
            ResponseStatusDTO.error(code, message),
            null
        );
    }
}
