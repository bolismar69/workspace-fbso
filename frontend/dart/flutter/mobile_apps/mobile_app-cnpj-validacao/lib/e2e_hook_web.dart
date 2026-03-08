import 'dart:convert';

import 'dart:js_util' as js_util;

void updateE2eResults(Map<String, Object?> data) {
  // Most reliable shape for Playwright to read: a plain JSON string.
  final json = jsonEncode(data);
  js_util.setProperty(js_util.globalThis, '__e2eResultsJson', json);

  // Best-effort object form (can fail if values are not jsify-compatible).
  try {
    js_util.setProperty(js_util.globalThis, '__e2eResults', js_util.jsify(data));
  } catch (_) {
    // ignore
  }
}
