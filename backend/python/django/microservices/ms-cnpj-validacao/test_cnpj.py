import json
import threading
import time
import unittest
from http.client import HTTPConnection

from app import Handler, ThreadingHTTPServer, is_valid_cnpj, normalize_cnpj


class CnpjValidatorTest(unittest.TestCase):
    def test_should_validate_known_valid_cnpj(self) -> None:
        self.assertTrue(is_valid_cnpj("04.252.011/0001-10"))
        self.assertTrue(is_valid_cnpj("40.688.134/0001-61"))

    def test_should_reject_invalid_cnpj(self) -> None:
        self.assertFalse(is_valid_cnpj("04.252.011/0001-11"))
        self.assertFalse(is_valid_cnpj("00000000000000"))
        self.assertFalse(is_valid_cnpj("11111111111111"))
        self.assertFalse(is_valid_cnpj(""))
        self.assertFalse(is_valid_cnpj(None))

    def test_should_normalize_digits(self) -> None:
        self.assertEqual("04252011000110", normalize_cnpj("04.252.011/0001-10"))
        self.assertEqual("123", normalize_cnpj("  1-2-3  "))
        self.assertIsNone(normalize_cnpj(""))
        self.assertIsNone(normalize_cnpj("   "))
        self.assertIsNone(normalize_cnpj(None))


class _TestServer:
    def __init__(self) -> None:
        self.httpd = ThreadingHTTPServer(("127.0.0.1", 0), Handler)
        self.port = self.httpd.server_address[1]
        self.thread = threading.Thread(target=self.httpd.serve_forever, daemon=True)

    def __enter__(self):
        self.thread.start()
        # Give the server a moment.
        time.sleep(0.05)
        return self

    def __exit__(self, exc_type, exc, tb):
        self.httpd.shutdown()
        self.httpd.server_close()


class CnpjResourceTest(unittest.TestCase):
    def test_validate_get_should_return_valid_true(self) -> None:
        with _TestServer() as s:
            conn = HTTPConnection("127.0.0.1", s.port, timeout=2)
            conn.request("GET", "/cnpj/validate?value=04.252.011/0001-10")
            res = conn.getresponse()
            self.assertEqual(200, res.status)
            body = json.loads(res.read().decode("utf-8"))
            self.assertEqual(True, body["valid"])
            self.assertEqual("04252011000110", body["normalized"])

    def test_validate_post_should_return_valid_false(self) -> None:
        with _TestServer() as s:
            conn = HTTPConnection("127.0.0.1", s.port, timeout=2)
            payload = json.dumps({"cnpj": "04.252.011/0001-11"}).encode("utf-8")
            conn.request(
                "POST",
                "/cnpj/validate",
                body=payload,
                headers={"Content-Type": "application/json"},
            )
            res = conn.getresponse()
            self.assertEqual(200, res.status)
            body = json.loads(res.read().decode("utf-8"))
            self.assertEqual(False, body["valid"])
            self.assertEqual("04252011000111", body["normalized"])


if __name__ == "__main__":
    unittest.main()
