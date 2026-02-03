#!/usr/bin/env python3
import argparse
import json
import os
import sys
from typing import Any, Dict, List, Optional, Set, Tuple


def _json_compact(value: Any) -> str:
    return json.dumps(value, separators=(",", ":"), ensure_ascii=False)


def _parse_bool(value: str, default: bool) -> bool:
    v = (value or "").strip().lower()
    if v in {"1", "true", "yes", "y", "on"}:
        return True
    if v in {"0", "false", "no", "n", "off"}:
        return False
    return default


def _parse_services(raw: str) -> Tuple[List[Dict[str, Any]], List[str]]:
    raw = (raw or "").strip()
    if not raw or raw == "null":
        raise SystemExit("input.services está vazio")

    try:
        parsed = json.loads(raw)
    except Exception:
        raise SystemExit("input.services não é JSON válido")

    if not isinstance(parsed, list):
        raise SystemExit("input.services deve ser um JSON array")

    # Nota: tratamos array vazio como 'sem override' aqui para não travar o workflow.
    # O job já tende a ser skipped pelo workflow quando services == '[]'.
    if len(parsed) == 0:
        return [], []

    paths: List[str] = []
    entries: List[Dict[str, Any]] = []

    for entry in parsed:
        if isinstance(entry, str):
            raise SystemExit("layout inválido: esperado array de objetos {path}, não strings")

        if not isinstance(entry, dict):
            raise SystemExit("layout inválido: esperado array de objetos com 'path'")

        p = (entry.get("path") or "").strip()
        if not p:
            raise SystemExit("cada item de input.services deve conter 'path'")

        paths.append(p)
        entries.append({"path": p})

    return entries, paths


def _split_path(path: str) -> List[str]:
    return [p for p in (path or "").strip().split("/") if p]


def main() -> int:
    parser = argparse.ArgumentParser(
        description=(
            "Valida APENAS o layout de input.services (workflow_dispatch/workflow_call). "
            "Formato esperado: JSON array não-vazio de objetos {path}. "
            "Não valida contra o inventory (isso ocorre em job posterior)."
        )
    )
    parser.add_argument(
        "--services",
        default=os.environ.get("INPUT_SERVICES", ""),
        help="JSON array (string) de objetos com 'path'.",
    )
    parser.add_argument(
        "--stack",
        default=os.environ.get("INPUT_STACK", ""),
        help="Stack esperada (backend|frontend|orchestration). Opcional.",
    )
    parser.add_argument(
        "--strict",
        default=os.environ.get("INPUT_STRICT", "false"),
        help=(
            "(Compat) Ignorado nesta versão: validação de inventory/specification ocorre em job posterior."
        ),
    )
    args = parser.parse_args()

    _ = _parse_bool(str(args.strict), default=True)
    stack = (args.stack or "").strip().lower()
    if stack and stack not in {"backend", "frontend", "orchestration"}:
        raise SystemExit(f"Stack inválida: '{args.stack}'. Esperado backend|frontend|orchestration")

    provided_entries, provided_paths = _parse_services(args.services)

    seen: Set[str] = set()
    dedup_paths: List[str] = []
    for p in provided_paths:
        if p in seen:
            continue
        seen.add(p)
        dedup_paths.append(p)

    errors: List[str] = []
    normalized: List[Dict[str, Any]] = []
    normalized_paths: List[str] = []

    for p in dedup_paths:
        if stack and not p.startswith(f"{stack}/"):
            errors.append(f"path inválido para stack '{stack}': {p}")
            continue

        parts = _split_path(p)
        if len(parts) != 5:
            errors.append(f"path inválido (esperado 5 níveis stack/platform/framework/type/name): {p}")
            continue

        normalized.append({"path": p})
        normalized_paths.append(p)

    if errors:
        print("Itens inválidos em input.services (serão ignorados):", file=sys.stderr)
        for e in errors:
            print(f"- {e}", file=sys.stderr)

    if not normalized:
        print(
            "Nenhum item válido restou após a validação; seguindo com lista vazia (nenhuma solution será processada).",
            file=sys.stderr,
        )

    paths = sorted(set([str(p).strip() for p in normalized_paths if str(p).strip()]))
    normalized = sorted({str(s.get("path")): s for s in normalized}.values(), key=lambda x: str(x.get("path")))

    solutions_json = _json_compact(normalized)
    paths_json = _json_compact(paths)

    github_output = os.environ.get("GITHUB_OUTPUT", "")
    if github_output:
        with open(github_output, "a", encoding="utf-8") as f:
            f.write(f"solutions={solutions_json}\n")
            f.write(f"paths={paths_json}\n")

    print("Validated input.services layout (paths only):")
    print(json.dumps(normalized, indent=2, ensure_ascii=False))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
