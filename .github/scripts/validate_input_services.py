#!/usr/bin/env python3
import argparse
import json
import os
import sys
from pathlib import Path
from typing import Any, Dict, List, Optional, Set, Tuple

_MANAGER_SOLUTIONS_PATH = Path("architecture/governance/config/manager-solutions-inventory.json")


def _json_compact(value: Any) -> str:
    return json.dumps(value, separators=(",", ":"), ensure_ascii=False)


def _parse_bool(value: str, default: bool) -> bool:
    v = (value or "").strip().lower()
    if v in {"1", "true", "yes", "y", "on"}:
        return True
    if v in {"0", "false", "no", "n", "off"}:
        return False
    return default


def _load_manager_inventory_by_path() -> Dict[str, Dict[str, Any]]:
    if not _MANAGER_SOLUTIONS_PATH.exists():
        raise SystemExit(
            f"Inventory não encontrado em '{_MANAGER_SOLUTIONS_PATH}'. "
            "Verifique o path e commit do arquivo."
        )

    data = json.loads(_MANAGER_SOLUTIONS_PATH.read_text(encoding="utf-8"))
    if "solutions" not in data:
        raise SystemExit(
            "Inventory inválido: esperado objeto com a chave 'solutions' (array). "
            "O formato antigo não é mais aceito."
        )

    raw_items = data.get("solutions") or []
    if not isinstance(raw_items, list):
        raise SystemExit("Inventory inválido: esperado array em solutions")

    by_path: Dict[str, Dict[str, Any]] = {}
    for raw in raw_items:
        item = dict(raw or {})
        spec = item.get("specification")
        if not isinstance(spec, dict):
            raise SystemExit(
                "Inventory inválido: cada solution deve conter 'specification' (objeto)."
            )

        # Backward-compat dos outputs: expõe alguns atalhos no topo,
        # mas mantém o objeto specification original.
        if item.get("platformVersion") is None and spec.get("platformVersion") is not None:
            item["platformVersion"] = spec.get("platformVersion")
        if item.get("platformDistributor") is None and spec.get("platformDistributor") is not None:
            item["platformDistributor"] = spec.get("platformDistributor")
        if item.get("docker") is None and spec.get("docker") is not None:
            item["docker"] = spec.get("docker")

        p = (item.get("path") or "").strip()
        if p:
            by_path[p] = item

    return by_path


def _parse_services(raw: str) -> Tuple[List[Dict[str, Any]], List[str]]:
    raw = (raw or "").strip()
    if not raw or raw == "null":
        raise SystemExit("input.services está vazio")

    try:
        parsed = json.loads(raw)
    except Exception:
        raise SystemExit("input.services não é JSON válido")

    if not isinstance(parsed, list) or len(parsed) == 0:
        raise SystemExit("input.services deve ser um JSON array não-vazio")

    paths: List[str] = []
    entries: List[Dict[str, Any]] = []

    for entry in parsed:
        if isinstance(entry, str):
            p = entry.strip()
            if not p:
                raise SystemExit("input.services contém um path vazio")
            paths.append(p)
            entries.append({"path": p})
            continue

        if not isinstance(entry, dict):
            raise SystemExit("input.services deve conter strings (paths) ou objetos com 'path'")

        p = (entry.get("path") or "").strip()
        if not p:
            raise SystemExit("cada item de input.services deve conter 'path'")

        paths.append(p)
        entries.append(dict(entry))

    return entries, paths


def _validate_required_fields(item: Dict[str, Any], strict: bool) -> List[str]:
    missing: List[str] = []

    for key in ("stack", "platform", "framework", "type", "name", "status", "path"):
        v = item.get(key)
        if v is None or (isinstance(v, str) and not v.strip()):
            missing.append(key)

    if strict:
        spec = item.get("specification")
        if not isinstance(spec, dict):
            missing.append("specification")
            return missing

        for key in ("platformVersion", "platformDistributor"):
            v = spec.get(key)
            if v is None or (isinstance(v, str) and not str(v).strip()):
                missing.append(f"specification.{key}")

        docker = spec.get("docker") or {}
        if not isinstance(docker, dict) or not str(docker.get("argCompilationMode") or "").strip():
            missing.append("specification.docker.argCompilationMode")

    return missing


def main() -> int:
    parser = argparse.ArgumentParser(
        description=(
            "Valida/enriquece input.services (workflow_dispatch/workflow_call) usando o inventory de soluções. "
            "Retorna outputs compatíveis com o detector: solutions e paths (JSON compact)."
        )
    )
    parser.add_argument(
        "--services",
        default=os.environ.get("INPUT_SERVICES", ""),
        help="JSON array (string). Pode conter objetos com 'path' ou strings (paths).",
    )
    parser.add_argument(
        "--stack",
        default=os.environ.get("INPUT_STACK", ""),
        help="Stack esperada (backend|frontend|orchestration). Opcional.",
    )
    parser.add_argument(
        "--strict",
        default=os.environ.get("INPUT_STRICT", "false"),
        help="Se true, exige specification.platformVersion/specification.platformDistributor/specification.docker.argCompilationMode no inventory.",
    )
    args = parser.parse_args()

    strict = _parse_bool(str(args.strict), default=True)
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

    inv_by_path = _load_manager_inventory_by_path()

    enriched: List[Dict[str, Any]] = []
    errors: List[str] = []

    for p in dedup_paths:
        if stack and not p.startswith(f"{stack}/"):
            errors.append(f"path inválido para stack '{stack}': {p}")
            continue

        # Estrutura unificada: stack/platform/framework/type/name
        parts = [seg for seg in p.split("/") if seg]
        if len(parts) != 5:
            errors.append(
                f"path inválido (esperado 5 níveis stack/platform/framework/type/name): {p}"
            )
            continue

        inv = inv_by_path.get(p)
        if not inv:
            errors.append(f"path não encontrado no inventory: {p}")
            continue

        missing = _validate_required_fields(inv, strict=strict)
        if missing:
            errors.append(f"inventory incompleto para '{p}' (faltando: {', '.join(missing)})")
            continue

        # Usa o inventory como fonte de verdade para o shape final
        enriched.append(dict(inv))

    if errors:
        msg = "\n".join(["Falha na validação de input.services:"] + [f"- {e}" for e in errors])
        raise SystemExit(msg)

    paths = [s.get("path") for s in enriched if (s.get("path") or "").strip()]
    paths = sorted(set([str(p).strip() for p in paths if str(p).strip()]))

    enriched = sorted({str(s.get("path")): s for s in enriched}.values(), key=lambda x: str(x.get("path")))

    solutions_json = _json_compact(enriched)
    paths_json = _json_compact(paths)

    github_output = os.environ.get("GITHUB_OUTPUT", "")
    if github_output:
        with open(github_output, "a", encoding="utf-8") as f:
            f.write(f"solutions={solutions_json}\n")
            f.write(f"paths={paths_json}\n")

    print("Validated/enriched services:")
    print(json.dumps(enriched, indent=2, ensure_ascii=False))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
