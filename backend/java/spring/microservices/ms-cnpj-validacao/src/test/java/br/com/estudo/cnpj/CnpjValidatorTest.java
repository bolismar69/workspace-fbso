package br.com.estudo.cnpj;

import static org.junit.jupiter.api.Assertions.*;

import org.junit.jupiter.api.Test;

class CnpjValidatorTest {

    @Test
    void shouldValidateKnownValidCnpj() {
        assertTrue(CnpjValidator.isValid("04.252.011/0001-10"));
        assertTrue(CnpjValidator.isValid("40.688.134/0001-61"));
    }

    @Test
    void shouldRejectInvalidCnpj() {
        assertFalse(CnpjValidator.isValid("04.252.011/0001-11"));
        assertFalse(CnpjValidator.isValid("00000000000000"));
        assertFalse(CnpjValidator.isValid("11111111111111"));
        assertFalse(CnpjValidator.isValid(""));
        assertFalse(CnpjValidator.isValid(null));
    }

    @Test
    void shouldNormalizeDigits() {
        assertEquals("04252011000110", CnpjValidator.normalize("04.252.011/0001-10"));
        assertEquals("123", CnpjValidator.normalize("  1-2-3  "));
        assertNull(CnpjValidator.normalize(""));
        assertNull(CnpjValidator.normalize("   "));
    }
}
