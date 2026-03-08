import { StatusBar } from 'expo-status-bar';
import { useMemo, useState } from 'react';
import { StyleSheet, Text, TextInput, View } from 'react-native';

import { formatCnpjMasked, isValidCnpj, normalizeCnpj } from './src/cnpj';

export default function App() {
  const [cnpjMasked, setCnpjMasked] = useState<string>('');

  const normalized = useMemo(() => normalizeCnpj(cnpjMasked), [cnpjMasked]);
  const valid = useMemo(() => (normalized === null ? null : isValidCnpj(normalized)), [normalized]);

  const inputValue = cnpjMasked.trim().length === 0 ? null : cnpjMasked;

  return (
    <View style={styles.container}>
      <Text style={styles.title}>Validação de CNPJ</Text>

      <View style={styles.form}>
        <Text style={styles.label}>CNPJ</Text>
        <TextInput
          nativeID="cnpj-input"
          style={styles.input}
          inputMode="numeric"
          keyboardType="numeric"
          autoCorrect={false}
          autoCapitalize="none"
          placeholder="99.999.999/9999-99"
          value={cnpjMasked}
          onChangeText={(text) => setCnpjMasked(formatCnpjMasked(text))}
        />

        <View style={styles.result}>
          <Text nativeID="result-input" style={styles.resultLine}>
            <Text style={styles.resultKey}>input:</Text> {display(inputValue)}
          </Text>
          <Text nativeID="result-normalized" style={styles.resultLine}>
            <Text style={styles.resultKey}>normalized:</Text> {display(normalized)}
          </Text>
          <Text nativeID="result-valid" style={styles.resultLine}>
            <Text style={styles.resultKey}>valid:</Text> {displayBool(valid)}
          </Text>
        </View>
      </View>

      <StatusBar style="auto" />
    </View>
  );
}

function display(value: string | null): string {
  return value ?? 'null';
}

function displayBool(value: boolean | null): string {
  return value === null ? 'null' : value ? 'true' : 'false';
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    padding: 24,
    paddingTop: 64,
  },
  title: {
    fontSize: 24,
    fontWeight: '700',
    marginBottom: 16,
  },
  form: {
    gap: 8,
    maxWidth: 420,
  },
  label: {
    fontSize: 14,
    fontWeight: '600',
  },
  input: {
    borderWidth: 1,
    borderColor: '#cbd5e1',
    borderRadius: 8,
    paddingHorizontal: 12,
    paddingVertical: 10,
    fontSize: 16,
  },
  result: {
    marginTop: 12,
    gap: 6,
  },
  resultLine: {
    fontSize: 14,
  },
  resultKey: {
    fontWeight: '700',
  },
});
