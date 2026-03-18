package com.fbso.geolocalidade.service;

public class InvalidCepException extends RuntimeException {
  public InvalidCepException(String message) {
    super(message);
  }
}
