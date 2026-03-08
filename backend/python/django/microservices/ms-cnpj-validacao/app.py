import json
import os
import re
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer
from urllib.parse import parse_qs, urlparse


def normalize_cnpj(value: str | None) -> str | None:
    if value is None:
        return None
    digits_only = re.sub(r"\D", "", value)
    if digits_only.strip() == "":
        return None
    return digits_only


def _all_digits_same(digits: str) -> bool:
    first = digits[0]
    return all(ch == first for ch in digits)


def _calculate_check_digit(digits: str, length: int) -> int:
    if length == 12:
        weights = [5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2]
    else:
        weights = [6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2]

    total = 0
    for i in range(length):
        total += int(digits[i]) * weights[i]

    mod = total % 11
    return 0 if mod < 2 else 11 - mod


def is_valid_cnpj(value: str | None) -> bool:
    normalized = normalize_cnpj(value)
    if normalized is None:
        return False

    if len(normalized) != 14:
        return False

    if _all_digits_same(normalized):
        return False

    d1 = _calculate_check_digit(normalized, 12)
    d2 = _calculate_check_digit(normalized, 13)
    return normalized[12] == str(d1) and normalized[13] == str(d2)


class Handler(BaseHTTPRequestHandler):
    server_version = "ms-cnpj-validacao/py"

    def _send_json(self, status_code: int, body: dict) -> None:
        encoded = json.dumps(body).encode("utf-8")
        self.send_response(status_code)
        self.send_header("Content-Type", "application/json")
        self.send_header("Content-Length", str(len(encoded)))
        self.end_headers()
        self.wfile.write(encoded)

    def do_GET(self) -> None:  # noqa: N802
        parsed = urlparse(self.path)
        if parsed.path != "/cnpj/validate":
            self.send_error(404)
            return

        qs = parse_qs(parsed.query, keep_blank_values=True)
        input_value = qs.get("value", [None])[0]

        response = {
            "input": input_value,
            "normalized": normalize_cnpj(input_value),
            "valid": is_valid_cnpj(input_value),
        }
        self._send_json(200, response)

    def do_POST(self) -> None:  # noqa: N802
        parsed = urlparse(self.path)
        if parsed.path != "/cnpj/validate":
            self.send_error(404)
            return

        length = int(self.headers.get("Content-Length") or "0")
        raw = self.rfile.read(length) if length > 0 else b""

        input_value = None
        if raw.strip() != b"":
            try:
                payload = json.loads(raw.decode("utf-8"))
            except Exception:
                self._send_json(400, {"error": "invalid json"})
                return

            if isinstance(payload, dict):
                input_value = payload.get("cnpj")

        response = {
            "input": input_value,
            "normalized": normalize_cnpj(input_value),
            "valid": is_valid_cnpj(input_value),
        }
        self._send_json(200, response)

    def log_message(self, fmt: str, *args) -> None:
        # Keep logs concise and consistent.
        super().log_message(fmt, *args)


def main() -> None:
    port = int(os.environ.get("PORT", "8080"))
    host = "0.0.0.0"

    httpd = ThreadingHTTPServer((host, port), Handler)
    print(f"listening on {host}:{port}")
    httpd.serve_forever()


if __name__ == "__main__":
    main()
