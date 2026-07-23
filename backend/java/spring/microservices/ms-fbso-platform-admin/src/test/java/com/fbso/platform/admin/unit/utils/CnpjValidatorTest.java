package com.fbso.platform.admin.unit.utils;

import com.fbso.platform.admin.utils.CnpjValidator;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Nested;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.NullAndEmptySource;
import org.junit.jupiter.params.provider.ValueSource;

import static org.assertj.core.api.Assertions.assertThat;
import static org.assertj.core.api.Assertions.assertThatThrownBy;

@DisplayName("CnpjValidator")
class CnpjValidatorTest {

    @Nested
    @DisplayName("CNPJs válidos — numéricos (legado)")
    class NumericValid {

        @Test
        @DisplayName("com máscara: 11.222.333/0001-81")
        void validWithMask() {
            assertThat(CnpjValidator.isValid("11.222.333/0001-81")).isTrue();
        }

        @Test
        @DisplayName("sem máscara: 11222333000181")
        void validWithoutMask() {
            assertThat(CnpjValidator.isValid("11222333000181")).isTrue();
        }

        @ParameterizedTest
        @ValueSource(strings = {
            "20.577.209/0001-20",
            "12.337.865/0001-27",
            "12.979.537/0001-24"
        })
        @DisplayName("CNPJs reais validados")
        void realWorldCnpjs(String cnpj) {
            assertThat(CnpjValidator.isValid(cnpj)).isTrue();
        }
    }

    @Nested
    @DisplayName("CNPJs válidos — alfanuméricos (IN RFB 2.119/2022)")
    class AlphanumericValid {

        @Test
        @DisplayName("exemplo oficial: 12.ABC.345/01DE-35")
        void officialExample() {
            assertThat(CnpjValidator.isValid("12.ABC.345/01DE-35"))
                    .describedAs("Exemplo da IN RFB 2.119/2022, Anexo XV")
                    .isTrue();
        }

        @Test
        @DisplayName("exemplo oficial sem máscara: 12ABC34501DE35")
        void officialExampleNoMask() {
            assertThat(CnpjValidator.isValid("12ABC34501DE35")).isTrue();
        }

        @Test
        @DisplayName("minúsculas convertidas: 12.abc.345/01de-35")
        void lowercaseConverted() {
            assertThat(CnpjValidator.isValid("12.abc.345/01de-35")).isTrue();
        }

        @ParameterizedTest
        @ValueSource(strings = {
            "3X.BLP.J0D/0001-00",
            "BC.PZH.24J/0001-67",
            "BE.AA3.NLH/0001-17",
            "MB.V7S.45K/0001-06",
            "KY.BDN.AHS/0001-97",
            "P2.DWE.PLZ/0001-98",
            "20.N09.L1L/0001-15",
            "VV.VTT.B6G/0001-09"
        })
        @DisplayName("CNPJs alfanuméricos reais validados")
        void realWorldAlphanumericCnpjs(String cnpj) {
            assertThat(CnpjValidator.isValid(cnpj)).isTrue();
        }
    }

    @Nested
    @DisplayName("CNPJs inválidos — numéricos")
    class NumericInvalid {

        @Test
        @DisplayName("dígitos verificadores errados")
        void wrongCheckDigits() {
            assertThat(CnpjValidator.isValid("11222333000199")).isFalse();
        }

        @Test
        @DisplayName("todos zeros — 00.000.000/0000-00")
        void allZeros() {
            assertThat(CnpjValidator.isValid("00.000.000/0000-00")).isFalse();
        }

        @Test
        @DisplayName("todos zeros sem máscara")
        void allZerosNoMask() {
            assertThat(CnpjValidator.isValid("00000000000000")).isFalse();
        }

        @ParameterizedTest
        @ValueSource(strings = {"123", "12.345.678/0001-9", "123456789012345"})
        @DisplayName("comprimento incorreto")
        void wrongLength(String cnpj) {
            assertThat(CnpjValidator.isValid(cnpj)).isFalse();
        }
    }

    @Nested
    @DisplayName("CNPJs inválidos — alfanuméricos")
    class AlphanumericInvalid {

        @Test
        @DisplayName("DV errado: 12.ABC.345/01DE-99")
        void wrongDv() {
            assertThat(CnpjValidator.isValid("12.ABC.345/01DE-99")).isFalse();
        }

        @Test
        @DisplayName("comprimento incorreto (13 chars)")
        void tooShort() {
            assertThat(CnpjValidator.isValid("12ABC34501DE3")).isFalse();
        }

        @Test
        @DisplayName("DV não numérico — último char é letra")
        void dvNotNumeric() {
            assertThat(CnpjValidator.isValid("12ABC34501DE3A")).isFalse();
        }

        @Test
        @DisplayName("caractere inválido — símbolo $")
        void invalidCharSymbol() {
            assertThat(CnpjValidator.isValid("12.ABC.345/01$$-35")).isFalse();
        }

        @Test
        @DisplayName("caractere inválido — ç (cedilha)")
        void invalidCharCedilla() {
            assertThat(CnpjValidator.isValid("12.ABÇ.345/01DE-35")).isFalse();
        }
    }

    @Nested
    @DisplayName("Entradas nulas/vazias")
    class NullEmpty {

        @ParameterizedTest
        @NullAndEmptySource
        @ValueSource(strings = {"   ", "\t", "\n"})
        @DisplayName("retorna false")
        void returnsFalse(String cnpj) {
            assertThat(CnpjValidator.isValid(cnpj)).isFalse();
        }
    }

    @Nested
    @DisplayName("calculaDV")
    class CalculaDv {

        @Test
        @DisplayName("base numérica: 112223330001 → 81")
        void numericBase() {
            assertThat(CnpjValidator.calculaDV("112223330001")).isEqualTo("81");
        }

        @Test
        @DisplayName("base alfanumérica: 12ABC34501DE → 35")
        void alphanumericBase() {
            assertThat(CnpjValidator.calculaDV("12ABC34501DE")).isEqualTo("35");
        }

        @Test
        @DisplayName("base com máscara: 11.222.333/0001 → 81")
        void baseWithMask() {
            assertThat(CnpjValidator.calculaDV("11.222.333/0001")).isEqualTo("81");
        }

        @Test
        @DisplayName("base apenas zeros → exceção")
        void allZerosThrows() {
            assertThatThrownBy(() -> CnpjValidator.calculaDV("000000000000"))
                    .isInstanceOf(IllegalArgumentException.class);
        }

        @Test
        @DisplayName("base inválida → exceção")
        void invalidThrows() {
            assertThatThrownBy(() -> CnpjValidator.calculaDV("123"))
                    .isInstanceOf(IllegalArgumentException.class);
        }

        @Test
        @DisplayName("null → exceção")
        void nullThrows() {
            assertThatThrownBy(() -> CnpjValidator.calculaDV(null))
                    .isInstanceOf(IllegalArgumentException.class);
        }
    }

    @Nested
    @DisplayName("strip")
    class Strip {

        @Test
        @DisplayName("remove máscara numérica")
        void numericMask() {
            assertThat(CnpjValidator.strip("11.222.333/0001-81"))
                    .isEqualTo("11222333000181");
        }

        @Test
        @DisplayName("remove máscara alfanumérica e uppercase")
        void alphanumericMask() {
            assertThat(CnpjValidator.strip("12.abc.345/01de-35"))
                    .isEqualTo("12ABC34501DE35");
        }

        @Test
        @DisplayName("null retorna null")
        void nullReturnsNull() {
            assertThat(CnpjValidator.strip(null)).isNull();
        }
    }

    @Nested
    @DisplayName("format")
    class Format {

        @Test
        @DisplayName("formata CNPJ numérico")
        void numeric() {
            assertThat(CnpjValidator.format("11222333000181"))
                    .isEqualTo("11.222.333/0001-81");
        }

        @Test
        @DisplayName("formata CNPJ alfanumérico")
        void alphanumeric() {
            assertThat(CnpjValidator.format("12ABC34501DE35"))
                    .isEqualTo("12.ABC.345/01DE-35");
        }

        @Test
        @DisplayName("entrada inválida → exceção")
        void invalidThrows() {
            assertThatThrownBy(() -> CnpjValidator.format("123"))
                    .isInstanceOf(IllegalArgumentException.class);
        }

        @Test
        @DisplayName("null → exceção")
        void nullThrows() {
            assertThatThrownBy(() -> CnpjValidator.format(null))
                    .isInstanceOf(IllegalArgumentException.class);
        }
    }
}
