#!/usr/bin/env python3
"""Apply global 'type-level' specification defaults.

Reads a JSON array of solution objects (same contract as other scripts) and
replaces missing/empty `specification` with the matching entry from:
`architecture/governance/config/global-types-solution.json`.

Matching key: {stack, platform, framework, type}.

Outputs (GHA-style):
- solutions: compact JSON array string
- paths: compact JSON array string

This script is intentionally permissive: it only requires each solution to have
`path` plus the four key fields (or infers them from the unified 5-level path).
"""

from __future__ import annotations

import argparse
import copy
import json
import os
import sys
from typing import Any, Dict, List, Optional, Tuple


DEFAULT_GLOBAL_TYPES_PATH = (
    "architecture/governance/config/global-types-solution.json"
)


def _is_empty_spec(spec: Any) -> bool:
    if spec is None:
        return True
    if not isinstance(spec, dict):
        return True
    return len(spec) == 0


def _parse_json_array(raw: str) -> List[Any]:
    try:
        data = json.loads(raw)
    except Exception as exc:  # noqa: BLE001
        raise ValueError("services JSON inválido") from exc
    if not isinstance(data, list):
        raise ValueError("services deve ser um JSON array")
    return data


def _ensure_solution_obj(item: Any) -> Dict[str, Any]:
    if isinstance(item, str):
        return {"path": item}
    if isinstance(item, dict):
        return item
    raise ValueError("Cada service deve ser string (path) ou objeto")


def _split_path(path: str) -> List[str]:
    return [p for p in path.strip().split("/") if p]


def _infer_key_from_solution(solution: Dict[str, Any]) -> Tuple[str, str, str, str]:
    stack = (solution.get("stack") or "").strip()
    platform = (solution.get("platform") or "").strip()
    framework = (solution.get("framework") or "").strip()
    solution_type = (solution.get("type") or "").strip()

    if stack and platform and framework and solution_type:
        return (stack, platform, framework, solution_type)

    path = (solution.get("path") or "").strip()
    parts = _split_path(path)
    # Unified structure: stack/platform/framework/type/name
    if len(parts) >= 4:
        return (parts[0], parts[1], parts[2], parts[3])

    raise ValueError(
        "Não foi possível inferir chave {stack,platform,framework,type} "
        f"para path='{path}'"
    )


def _load_global_types(global_types_path: str) -> Dict[Tuple[str, str, str, str], Dict[str, Any]]:
    try:
        with open(global_types_path, "r", encoding="utf-8") as f:
            data = json.load(f)
    except FileNotFoundError as exc:
        raise FileNotFoundError(
            f"Arquivo global types não encontrado: {global_types_path}"
        ) from exc

    inventory = data.get("inventory") if isinstance(data, dict) else None
    if not isinstance(inventory, list):
        raise ValueError(
            "global-types-solution.json deve ter um objeto com a chave 'inventory' (array)"
        )

    mapping: Dict[Tuple[str, str, str, str], Dict[str, Any]] = {}
    for idx, item in enumerate(inventory):
        if not isinstance(item, dict):
            raise ValueError(f"inventory[{idx}] deve ser um objeto")

        missing = [
            k
            for k in ("stack", "platform", "framework", "type", "specification")
            if k not in item
        ]
        if missing:
            raise ValueError(f"inventory[{idx}] faltando campos: {', '.join(missing)}")

        spec = item.get("specification")
        if not isinstance(spec, dict):
            raise ValueError(f"inventory[{idx}].specification deve ser um objeto")

        key = (
            str(item.get("stack", "")).strip(),
            str(item.get("platform", "")).strip(),
            str(item.get("framework", "")).strip(),
            str(item.get("type", "")).strip(),
        )
        if any(not p for p in key):
            raise ValueError(f"inventory[{idx}] chave inválida: {key}")

        mapping[key] = spec

    return mapping


def _apply_spec_defaults(
    services: List[Any],
    global_types: Dict[Tuple[str, str, str, str], Dict[str, Any]],
) -> List[Dict[str, Any]]:
    out: List[Dict[str, Any]] = []

    for item in services:
        s = _ensure_solution_obj(item)

        path = str(s.get("path", "")).strip()
        if not path:
            # keep consistent with other scripts: ignore empty entries
            continue

        try:
            key = _infer_key_from_solution(s)
        except ValueError:
            # If it doesn't have enough info, keep it as-is
            out.append(s)
            continue

        if _is_empty_spec(s.get("specification")):
            spec = global_types.get(key)
            if isinstance(spec, dict):
                s["specification"] = copy.deepcopy(spec)

                # Convenience fields (keep aligned with other scripts)
                if isinstance(s["specification"].get("platformVersion"), (str, int)):
                    s["platformVersion"] = str(s["specification"].get("platformVersion"))
                if isinstance(s["specification"].get("platformDistributor"), str):
                    s["platformDistributor"] = s["specification"].get("platformDistributor")
                if isinstance(s["specification"].get("docker"), dict):
                    s["docker"] = copy.deepcopy(s["specification"].get("docker"))

        out.append(s)

    return out


def _gha_set_output(name: str, value: str) -> None:
    github_output = os.environ.get("GITHUB_OUTPUT")
    if github_output:
        with open(github_output, "a", encoding="utf-8") as f:
            f.write(f"{name}={value}\n")
    else:
        # local usage
        print(f"{name}={value}")


def main(argv: Optional[List[str]] = None) -> int:
    parser = argparse.ArgumentParser(
        description=(
            "Preenche 'specification' ausente/vazio em services usando "
            "architecture/governance/config/global-types-solution.json (por tipo)."
        )
    )
    parser.add_argument(
        "--services",
        help="JSON array (string). Se omitido, lê de env SERVICES_JSON.",
        default=None,
    )
    parser.add_argument(
        "--global-types",
        help="Path para global-types-solution.json",
        default=DEFAULT_GLOBAL_TYPES_PATH,
    )

    args = parser.parse_args(argv)

    raw_services = args.services
    if raw_services is None:
        raw_services = os.environ.get("SERVICES_JSON")

    if not raw_services:
        raise ValueError("services não informado (nem --services, nem SERVICES_JSON)")

    services = _parse_json_array(raw_services)
    global_types = _load_global_types(args.global_types)

    enriched = _apply_spec_defaults(services, global_types)
    solutions_json = json.dumps(enriched, separators=(",", ":"), ensure_ascii=False)
    paths_json = json.dumps(
        [str(s.get("path", "")).strip() for s in enriched if str(s.get("path", "")).strip()],
        separators=(",", ":"),
        ensure_ascii=False,
    )

    _gha_set_output("solutions", solutions_json)
    _gha_set_output("paths", paths_json)

    return 0


if __name__ == "__main__":
    try:
        raise SystemExit(main())
    except Exception as exc:  # noqa: BLE001
        print(f"ERROR: {exc}", file=sys.stderr)
        raise SystemExit(2)
