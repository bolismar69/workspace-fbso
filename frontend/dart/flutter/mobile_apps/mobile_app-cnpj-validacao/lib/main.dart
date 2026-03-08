import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import 'cnpj.dart';
import 'e2e_hook.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Validação de CNPJ',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: const CnpjPage(),
    );
  }
}

class CnpjPage extends StatefulWidget {
  const CnpjPage({super.key});

  @override
  State<CnpjPage> createState() => _CnpjPageState();
}

class _CnpjPageState extends State<CnpjPage> {
  final _controller = TextEditingController();
  final _formatter = const CnpjMaskFormatter();

  String _masked = '';

  @override
  void initState() {
    super.initState();
    _controller.addListener(() {
      setState(() {
        _masked = _controller.text;
      });
    });
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final input = _masked.trim().isEmpty ? null : _masked;
    final normalized = Cnpj.normalize(_masked);
    final valid = normalized == null ? null : Cnpj.isValid(normalized);

    final inputLine = 'input: ${display(input)}';
    final normalizedLine = 'normalized: ${display(normalized)}';
    final validLine = 'valid: ${displayBool(valid)}';

    updateE2eResults({
      'input': input,
      'normalized': normalized,
      'valid': valid,
    });

    return Scaffold(
      appBar: AppBar(
        title: Semantics(
          header: true,
          label: 'Validação de CNPJ',
          child: const Text('Validação de CNPJ'),
        ),
      ),
      body: Padding(
        padding: const EdgeInsets.all(24),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            const Text(
              'CNPJ',
              style: TextStyle(fontWeight: FontWeight.w600),
            ),
            const SizedBox(height: 8),
            TextField(
              key: const ValueKey('cnpj-input'),
              controller: _controller,
              keyboardType: TextInputType.number,
              decoration: const InputDecoration(
                labelText: 'CNPJ',
                hintText: '99.999.999/9999-99',
                border: OutlineInputBorder(),
              ),
              inputFormatters: <TextInputFormatter>[_formatter],
            ),
            const SizedBox(height: 16),
            Semantics(
              label: inputLine,
              child: Text(
                inputLine,
                key: const ValueKey('result-input'),
              ),
            ),
            Semantics(
              label: normalizedLine,
              child: Text(
                normalizedLine,
                key: const ValueKey('result-normalized'),
              ),
            ),
            Semantics(
              label: validLine,
              child: Text(
                validLine,
                key: const ValueKey('result-valid'),
              ),
            ),
          ],
        ),
      ),
    );
  }
}

String display(String? value) => value ?? 'null';

String displayBool(bool? value) => value == null ? 'null' : (value ? 'true' : 'false');
