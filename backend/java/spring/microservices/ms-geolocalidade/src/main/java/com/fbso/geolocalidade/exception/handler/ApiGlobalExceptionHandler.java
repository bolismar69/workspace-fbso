package com.fbso.geolocalidade.exception.handler;

import com.fbso.geolocalidade.dto.PageResponseDTO;
import com.fbso.geolocalidade.exception.AwesomeApiException;
import com.fbso.geolocalidade.exception.InvalidCepException;
import com.fbso.geolocalidade.exception.ResourceNotFoundException;
import java.util.stream.Collectors;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.MethodArgumentNotValidException;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

@RestControllerAdvice
public class ApiGlobalExceptionHandler {

  @ExceptionHandler(InvalidCepException.class)
  public ResponseEntity<PageResponseDTO<Object>> handleInvalidCep(InvalidCepException ex) {
    String message = (ex.getMessage() == null || ex.getMessage().isBlank())
        ? "CEP inválido"
        : "CEP inválido: " + ex.getMessage();

    var response = PageResponseDTO.error(HttpStatus.BAD_REQUEST.value(), message);
    return ResponseEntity.badRequest().body(response);
  }

  @ExceptionHandler(AwesomeApiException.class)
  public ResponseEntity<PageResponseDTO<Object>> handleAwesomeApi(AwesomeApiException ex) {
    String message = (ex.getMessage() == null || ex.getMessage().isBlank())
        ? "Falha ao consultar AwesomeAPI"
        : "Falha ao consultar AwesomeAPI: " + ex.getMessage();

    var response = PageResponseDTO.error(HttpStatus.BAD_GATEWAY.value(), message);
    return ResponseEntity.status(HttpStatus.BAD_GATEWAY).body(response);
  }

  @ExceptionHandler(ResourceNotFoundException.class)
  public ResponseEntity<PageResponseDTO<Object>> handleNotFound(ResourceNotFoundException ex) {
      var response = PageResponseDTO.error(HttpStatus.NOT_FOUND.value(), ex.getMessage());
      return ResponseEntity.status(HttpStatus.NOT_FOUND).body(response);
  }

  @ExceptionHandler(MethodArgumentNotValidException.class)
  public ResponseEntity<PageResponseDTO<Object>> handleValidation(MethodArgumentNotValidException ex) {
      String errors = ex.getBindingResult().getFieldErrors().stream()
              .map(error -> error.getField() + ": " + error.getDefaultMessage())
              .collect(Collectors.joining(", "));
      
      var response = PageResponseDTO.error(HttpStatus.BAD_REQUEST.value(), "Erro de validação: " + errors);
      return ResponseEntity.badRequest().body(response);
  }

  @ExceptionHandler(Exception.class)
  public ResponseEntity<PageResponseDTO<Object>> handleGeneric(Exception ex) {
      var response = PageResponseDTO.error(HttpStatus.INTERNAL_SERVER_ERROR.value(), "Erro interno no servidor");
      return ResponseEntity.internalServerError().body(response);
  }
}
