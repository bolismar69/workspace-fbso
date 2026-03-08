import 'e2e_hook_stub.dart' if (dart.library.html) 'e2e_hook_web.dart' as impl;

void updateE2eResults(Map<String, Object?> data) {
  impl.updateE2eResults(data);
}
