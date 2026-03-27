package com.fbso.geolocalidade.dto;

public record SingleResponseDTO<T>(
    T content,
    ResponseStatusDTO status,
    PageInfoDTO pageInfo
) {
  public static <T> SingleResponseDTO<T> of(T content, PageInfoDTO pageInfo) {
    return new SingleResponseDTO<>(
        content,
        ResponseStatusDTO.success(200, "OK"),
        pageInfo);
  }
}
