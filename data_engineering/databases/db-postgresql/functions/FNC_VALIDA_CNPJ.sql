CREATE OR REPLACE FUNCTION FNC_VALIDA_CNPJ(P_CNPJ TEXT) 
RETURNS BOOLEAN AS $$
DECLARE
    V_CNPJ_CLEAN TEXT;
    V_SOMA1 INT := 0;
    V_SOMA2 INT := 0;
    V_DIGITO1 INT;
    V_DIGITO2 INT;
    V_PESO1 INT[] := ARRAY[5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2];
    V_PESO2 INT[] := ARRAY[6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2];
BEGIN
    -- Limpeza e verificação básica
    V_CNPJ_CLEAN := REGEXP_REPLACE(P_CNPJ, '[^0-9]', '', 'g');
    
    IF LENGTH(V_CNPJ_CLEAN) != 14 THEN
        RETURN FALSE;
    END IF;

    -- Cálculo do 1º Dígito
    FOR i IN 1..12 LOOP
        V_SOMA1 := V_SOMA1 + (SUBSTRING(V_CNPJ_CLEAN FROM i FOR 1)::INT * V_PESO1[i]);
    END LOOP;
    V_DIGITO1 := 11 - (V_SOMA1 % 11);
    IF V_DIGITO1 >= 10 THEN V_DIGITO1 := 0; END IF;

    -- Cálculo do 2º Dígito
    FOR i IN 1..13 LOOP
        IF i < 13 THEN
            V_SOMA2 := V_SOMA2 + (SUBSTRING(V_CNPJ_CLEAN FROM i FOR 1)::INT * V_PESO2[i]);
        ELSE
            V_SOMA2 := V_SOMA2 + (V_DIGITO1 * V_PESO2[i]);
        END IF;
    END LOOP;
    V_DIGITO2 := 11 - (V_SOMA2 % 11);
    IF V_DIGITO2 >= 10 THEN V_DIGITO2 := 0; END IF;

    RETURN V_DIGITO1 = SUBSTRING(V_CNPJ_CLEAN FROM 13 FOR 1)::INT 
       AND V_DIGITO2 = SUBSTRING(V_CNPJ_CLEAN FROM 14 FOR 1)::INT;
END;
$$ LANGUAGE plpgsql;