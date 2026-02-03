#!/usr/bin/env python3
"""Validate identified solutions against the manager inventory.

This script recINPUT_SERVICESeives a JSON array of services (usually produced by the detector
or by validate_input_services) and validates them against:
  architecture/governance/config/manager-solutions-inventory.json

Rules:
- Input is expected as a JSON array of objects with {"path": "..."}.
  (Strings are accepted for backwards compatibility and will be converted.)
- Path must be the unified 5-level structure: stack/platform/framework/type/name.
- If a path is not present in the inventory: report it (do not fail).
- If present but status != "active": report it (do not fail).
- Only solutions that exist and are status=="active" are returned, using the
  full object as defined in the inventory (source of truth).

Outputs (GHA-style):
- solutions: compact JSON array string with full inventory objects
- paths: compact JSON array string with validated active paths

Exit code:
- 0 on success (even if there are missing/inactive paths).
- 2 on hard errors (invalid inventory file or invalid input JSON).
"""

from __future__ import annotations

import argparse
import json
import os
import sys
from pathlib import Path
from typing import Any, Dict, List, Optional, Set, Tuple

_MANAGER_SOLUTIONS_PATH = Path("architecture/governance/config/manager-solutions-inventory.json")
_ALLOWED_STACKS = {"backend", "frontend", "orchestration"}


def _json_compact(value: Any) -> str:
    return json.dumps(value, separators=(",", ":"), ensure_ascii=False)


def _split_path(path: str) -> List[str]:
    return [p for p in (path or "").strip().split("/") if p]


def _parse_services(raw: str) -> List[str]:
    raw = (raw or "").strip()
    if not raw or raw == "null":
        return []

    try:
        parsed = json.loads(raw)
    except Exception as exc:  # noqa: BLE001
        raise ValueError("services não é JSON válido") from exc

    if not isinstance(parsed, list):
        raise ValueError("services deve ser um JSON array")

    out: List[str] = []
    for entry in parsed:
        if isinstance(entry, str):
            p = entry.strip()
            if p:
                out.append(p)
            continue
        if isinstance(entry, dict):
            p = str(entry.get("path") or "").strip()
            if p:
                out.append(p)
            continue
        raise ValueError("services deve conter objetos {path} ou strings")

    # unique preserve order
    seen: Set[str] = set()
    dedup: List[str] = []
    for p in out:
        if p in seen:
            continue
        seen.add(p)
        dedup.append(p)
    return dedup


def _load_inventory_by_path() -> Dict[str, Dict[str, Any]]:
    if not _MANAGER_SOLUTIONS_PATH.exists():
        raise FileNotFoundError(
            f"Inventory não encontrado em '{_MANAGER_SOLUTIONS_PATH}'. Verifique o path e commit do arquivo."
        )

    data = json.loads(_MANAGER_SOLUTIONS_PATH.read_text(encoding="utf-8"))
    raw_items = data.get("solutions")
    if not isinstance(raw_items, list):
        raise ValueError("Inventory inválido: esperado objeto com a chave 'solutions' (array)")

    by_path: Dict[str, Dict[str, Any]] = {}
    for idx, raw in enumerate(raw_items):
        if not isinstance(raw, dict):
            raise ValueError(f"Inventory inválido: solutions[{idx}] deve ser um objeto")

        item = dict(raw)
        p = str(item.get("path") or "").strip()
        if not p:
            raise ValueError(f"Inventory inválido: solution sem 'path' (index={idx})")

        # Basic schema validation (do not over-validate)
        parts = _split_path(p)
        if len(parts) != 5:
            raise ValueError(
                f"Inventory inválido: path deve ter 5 níveis stack/platform/framework/type/name: '{p}'"
            )

        if p in by_path:
            # Keep the first occurrence but warn.
            print(f"WARNING: path duplicado no inventory (mantendo o primeiro): {p}", file=sys.stderr)
            continue

        by_path[p] = item

    return by_path


def _gha_set_output(name: str, value: str) -> None:
    github_output = os.environ.get("GITHUB_OUTPUT")
    if github_output:
        with open(github_output, "a", encoding="utf-8") as f:
            f.write(f"{name}={value}\n")


def main(argv: Optional[List[str]] = None) -> int:
    parser = argparse.ArgumentParser(
        description=(
            "Valida services identificados/reportados contra o inventory manager-solutions-inventory.json. "
            "Reporta paths não cadastrados ou inativos sem falhar e retorna apenas solutions válidas (active)."
        )
    )
    parser.add_argument(
        "--services",
        default=os.environ.get("INPUT_SERVICES", ""),
        help="JSON array (string) contendo objetos {path} (ou strings por compat).",
    )
    parser.add_argument(
        "--stack",
        default=os.environ.get("INPUT_STACK", ""),
        help="Opcional: stack esperada (backend|frontend|orchestration).",
    )
    args = parser.parse_args(argv)

    stack = str(args.stack or "").strip().lower()
    if stack and stack not in _ALLOWED_STACKS:
        raise ValueError(f"Stack inválida: '{args.stack}'. Esperado backend|frontend|orchestration")

    paths = _parse_services(args.services)
    if not paths:
        # No inputs: produce empty outputs but succeed.
        _gha_set_output("solutions", "[]")
        _gha_set_output("paths", "[]")
        print("INFO: Nenhuma solution informada para validação (services vazio).")
        return 0

    inv_by_path = _load_inventory_by_path()

    missing: List[str] = []
    inactive: List[str] = []
    invalid_format: List[str] = []

    validated: List[Dict[str, Any]] = []

    for p in paths:
        if stack and not p.startswith(f"{stack}/"):
            invalid_format.append(f"path fora da stack '{stack}': {p}")
            continue

        parts = _split_path(p)
        if len(parts) != 5:
            invalid_format.append(f"path inválido (esperado 5 níveis stack/platform/framework/type/name): {p}")
            continue

        inv = inv_by_path.get(p)
        if not inv:
            missing.append(p)
            continue

        status = str(inv.get("status") or "").strip().lower()
        if status != "active":
            inactive.append(f"{p} (status='{inv.get('status')}')")
            continue

        validated.append(dict(inv))

    # Reporting (non-blocking)
    if invalid_format:
        print("NOTICE: Alguns paths foram ignorados por layout/stack inválidos:")
        for e in invalid_format:
            print(f"- {e}")

    if missing:
        print("NOTICE: Alguns paths não estão cadastrados no inventory (ignorados):")
        for p in sorted(set(missing)):
            print(f"- {p}")

    if inactive:
        print("NOTICE: Alguns paths estão cadastrados, mas não estão active (ignorados):")
        for p in sorted(set(inactive)):
            print(f"- {p}")

    # Stable + unique
    validated = sorted({str(s.get("path")): s for s in validated}.values(), key=lambda x: str(x.get("path")))
    validated_paths = sorted({str(s.get("path")) for s in validated if str(s.get("path") or "").strip()})

    solutions_json = _json_compact(validated)
    paths_json = _json_compact(validated_paths)

    _gha_set_output("solutions", solutions_json)
    _gha_set_output("paths", paths_json)

    print(f"INFO: Solutions válidas (active) para processar: {len(validated)}")
    return 0


if __name__ == "__main__":
    try:
        raise SystemExit(main())
    except Exception as exc:  # noqa: BLE001
        print(f"ERROR: {exc}", file=sys.stderr)
        raise SystemExit(2)
