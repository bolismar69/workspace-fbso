package com.fbso.geolocalidade.exception;

public class AwesomeApiException extends RuntimeException {
  public AwesomeApiException(String message) {
    super(message);
  }

  public AwesomeApiException(String message, Throwable cause) {
    super(message, cause);
  }
}
