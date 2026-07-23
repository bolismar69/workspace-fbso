package com.fbso.platform.admin.utils;

/**
 * Validador de CNPJ (Cadastro Nacional da Pessoa Jurídica).
 *
 * <p>Implementa o algoritmo oficial conforme Instrução Normativa RFB nº 2.119/2022,
 * Anexo XV, que passou a vigorar em Julho/2026.</p>
 *
 * <h3>Formatos suportados (unificado)</h3>
 * <ul>
 *   <li>14 posições: 12 caracteres base (A-Z, 0-9) + 2 dígitos verificadores (0-9)</li>
 *   <li>Máscara: {@code SS.SSS.SSS/SSSS-DD} (S = letra ou número, D = dígito)</li>
 *   <li>CNPJs numéricos legados são um subconjunto natural (S = apenas dígitos)</li>
 *   <li>Rejeita sequências com todos os caracteres iguais a '0'</li>
 * </ul>
 *
 * <h3>Algoritmo</h3>
 * <p>Valor de cada caractere = código ASCII do caractere − código ASCII de '0' (48).
 * Isso produz 0-9 para dígitos e 17-42 para letras A-Z. Os pesos são aplicados
 * da direita para esquerda usando o array fixo {6,5,4,3,2,9,8,7,6,5,4,3,2}.</p>
 *
 * <h3>Uso</h3>
 * <pre>{@code
 *   CnpjValidator.isValid("11.222.333/0001-81")    // true (numérico legado)
 *   CnpjValidator.isValid("12.ABC.345/01DE-35")    // true (alfanumérico novo)
 *   CnpjValidator.isValid("00.000.000/0000-00")    // false (todos zeros)
 * }</pre>
 *
 * <p><b>DT-129 (Sprint 6):</b> Criado para substituir validação superficial
 * no {@code OnboardingService.isValidCnpj()}.</p>
 */
public final class CnpjValidator {

    private static final int TAMANHO_CNPJ_SEM_DV = 12;
    private static final String REGEX_CARACTERES_FORMATACAO = "[./-]";
    private static final String REGEX_FORMACAO_BASE_CNPJ = "[A-Z\\d]{12}";
    private static final String REGEX_FORMACAO_DV = "[\\d]{2}";
    private static final String REGEX_VALOR_ZERADO = "^[0]+$";

    /** Código ASCII do caractere '0' — base para cálculo do valor de cada posição. */
    private static final int VALOR_BASE = (int) '0';

    /**
     * Pesos para o cálculo dos dígitos verificadores.
     * Aplicados da direita para esquerda: para uma base de 12 caracteres,
     * usam-se as últimas 12 posições; para 13, o array completo.
     */
    private static final int[] PESOS_DV = { 6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2 };

    private CnpjValidator() {
        // classe utilitária — não instanciável
    }

    /**
     * Valida um CNPJ (numérico ou alfanumérico), com ou sem máscara.
     *
     * @param cnpj CNPJ a validar (pode ser nulo, com ou sem máscara, maiúsculo ou minúsculo)
     * @return {@code true} se o CNPJ for válido
     */
    public static boolean isValid(String cnpj) {
        if (cnpj != null) {
            cnpj = removeCaracteresFormatacao(cnpj);
            if (isCnpjFormacaoValidaComDV(cnpj)) {
                String dvInformado = cnpj.substring(TAMANHO_CNPJ_SEM_DV);
                String dvCalculado = calculaDV(cnpj.substring(0, TAMANHO_CNPJ_SEM_DV));
                return dvCalculado.equals(dvInformado);
            }
        }
        return false;
    }

    /**
     * Calcula os 2 dígitos verificadores para uma base de CNPJ (12 caracteres).
     *
     * @param baseCnpj 12 caracteres alfanuméricos (A-Z, 0-9), com ou sem máscara
     * @return string com os 2 dígitos verificadores
     * @throws IllegalArgumentException se a base não for válida
     */
    public static String calculaDV(String baseCnpj) {
        if (baseCnpj != null) {
            baseCnpj = removeCaracteresFormatacao(baseCnpj);
            if (isCnpjFormacaoValidaSemDV(baseCnpj)) {
                String dv1 = String.format("%d", calculaDigito(baseCnpj));
                String dv2 = String.format("%d", calculaDigito(baseCnpj.concat(dv1)));
                return dv1.concat(dv2);
            }
        }
        throw new IllegalArgumentException(
                String.format("CNPJ '%s' não é válido para o cálculo do DV", baseCnpj));
    }

    /**
     * Remove caracteres de formatação (., /, -) e converte para maiúsculas.
     */
    public static String strip(String cnpj) {
        if (cnpj == null) {
            return null;
        }
        return removeCaracteresFormatacao(cnpj);
    }

    /**
     * Aplica máscara de CNPJ: {@code SS.SSS.SSS/SSSS-DD}.
     *
     * @param raw 14 caracteres alfanuméricos sem máscara
     * @return CNPJ formatado
     * @throws IllegalArgumentException se não forem 14 caracteres no formato esperado
     */
    public static String format(String raw) {
        if (raw == null || !raw.matches("^[A-Z0-9]{14}$")) {
            throw new IllegalArgumentException(
                    "São necessários 14 caracteres alfanuméricos (A-Z, 0-9)");
        }
        return raw.substring(0, 2) + "." + raw.substring(2, 5) + "."
             + raw.substring(5, 8) + "/" + raw.substring(8, 12) + "-"
             + raw.substring(12);
    }

    // ---- Algoritmo interno ----

    /**
     * Calcula um único dígito verificador.
     * Itera da direita para esquerda, aplicando os pesos do array {@link #PESOS_DV}.
     */
    private static int calculaDigito(String cnpj) {
        int soma = 0;
        for (int indice = cnpj.length() - 1; indice >= 0; indice--) {
            int valorCaracter = (int) cnpj.charAt(indice) - VALOR_BASE;
            soma += valorCaracter * PESOS_DV[PESOS_DV.length - cnpj.length() + indice];
        }
        return soma % 11 < 2 ? 0 : 11 - (soma % 11);
    }

    private static String removeCaracteresFormatacao(String cnpj) {
        return cnpj.trim().toUpperCase().replaceAll(REGEX_CARACTERES_FORMATACAO, "");
    }

    private static boolean isCnpjFormacaoValidaSemDV(String cnpj) {
        return cnpj.matches(REGEX_FORMACAO_BASE_CNPJ) && !cnpj.matches(REGEX_VALOR_ZERADO);
    }

    private static boolean isCnpjFormacaoValidaComDV(String cnpj) {
        return cnpj.matches(REGEX_FORMACAO_BASE_CNPJ.concat(REGEX_FORMACAO_DV))
                && !cnpj.matches(REGEX_VALOR_ZERADO);
    }
}
