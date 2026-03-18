package com.fbso.geolocalidade.controller;

import com.fbso.geolocalidade.service.AwesomeApiException;
import com.fbso.geolocalidade.service.InvalidCepException;
import org.springframework.http.HttpStatus;
import org.springframework.http.ProblemDetail;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

@RestControllerAdvice
public class ApiExceptionHandler {

  @ExceptionHandler(InvalidCepException.class)
  public ProblemDetail handleInvalidCep(InvalidCepException ex) {
    var pd = ProblemDetail.forStatus(HttpStatus.BAD_REQUEST);
    pd.setTitle("CEP inválido");
    pd.setDetail(ex.getMessage());
    return pd;
  }

  @ExceptionHandler(AwesomeApiException.class)
  public ProblemDetail handleAwesomeApi(AwesomeApiException ex) {
    var pd = ProblemDetail.forStatus(HttpStatus.BAD_GATEWAY);
    pd.setTitle("Falha ao consultar AwesomeAPI");
    pd.setDetail(ex.getMessage());
    return pd;
  }
}
