package com.fbso.geolocalidade.utils;

import com.fbso.geolocalidade.exception.InvalidCepException;

import java.util.Objects;
import java.util.regex.Pattern;

public final class CepUtils {
  private static final Pattern NON_DIGITS = Pattern.compile("[^0-9]");

  private CepUtils() {}

  public static String normalizeCep(String cep) {
    Objects.requireNonNull(cep, "cep");
    String cleaned = NON_DIGITS.matcher(cep).replaceAll("");
    if (cleaned.length() != 8) {
      throw new InvalidCepException("CEP deve ter 8 dígitos: '" + cep + "'");
    }
    return cleaned;
  }
}
