#!/usr/bin/env python3
import argparse
import json
import os
import re
import subprocess
import sys
from pathlib import Path
from typing import Any, Dict, List, Optional, Set, Tuple

_MANAGER_SOLUTIONS_PATH = Path("architecture/governance/config/manager-solutions-inventory.json")

STACKS = {"backend", "frontend", "orchestration"}
SOLUTION_CHANGED_RE = re.compile(r"^(backend|frontend|orchestration)/[^/]+/[^/]+/[^/]+/[^/]+/")

# Hash especial do Git para a "árvore vazia". Útil para push inicial
# (quando o event.before vem como 40 zeros).
GIT_EMPTY_TREE = "4b825dc642cb6eb9a060e54bf8d69288fbee4904"


def _run(cmd: List[str], check: bool = True) -> subprocess.CompletedProcess:
    return subprocess.run(cmd, check=check, text=True, capture_output=True)


def _read_event_payload() -> Dict[str, Any]:
    event_path = os.environ.get("GITHUB_EVENT_PATH", "")
    if not event_path or not os.path.exists(event_path):
        return {}
    try:
        with open(event_path, "r", encoding="utf-8") as f:
            return json.load(f)
    except Exception:
        return {}


def _infer_base_head(base: Optional[str], head: Optional[str]) -> Tuple[str, str]:
    base = (base or "").strip()
    head = (head or "").strip()

    if base and head:
        return base, head

    event_name = (os.environ.get("GITHUB_EVENT_NAME") or "").strip()
    event = _read_event_payload()

    if (not base or not head) and event_name == "pull_request":
        pr = event.get("pull_request") or {}
        if not base:
            base = (((pr.get("base") or {}).get("sha")) or "").strip()
        if not head:
            head = (((pr.get("head") or {}).get("sha")) or "").strip()
    else:
        if not base:
            base = (event.get("before") or "").strip()
        if not head:
            head = (os.environ.get("GITHUB_SHA") or "").strip()

    if not base or not head:
        raise SystemExit(
            f"Não foi possível resolver BASE/HEAD. "
            f"base='{base}' head='{head}' event_name='{event_name}'"
        )
    return base, head


def _get_changed_files(base: str, head: str) -> List[str]:
    # equivalente ao workflow atual
    try:
        _run(["git", "fetch", "--no-tags", "--prune", "--depth=0", "origin"], check=False)
    except Exception:
        pass

    # Push inicial pode trazer base=0000...0000. Nesse caso, fazemos diff contra a árvore vazia
    # para obter todos os paths presentes em HEAD.
    if re.fullmatch(r"0+", base) and len(base) >= 40:
        base = GIT_EMPTY_TREE

    try:
        proc = _run(["git", "diff", "--name-only", base, head], check=False)
        out = (proc.stdout or "").splitlines()
        return [line.strip() for line in out if line.strip()]
    except Exception:
        return []


def _roots_from_changed_files(changed: List[str]) -> Set[str]:
    roots: Set[str] = set()

    for f in changed:
        if SOLUTION_CHANGED_RE.match(f):
            parts = f.split("/")
            if len(parts) >= 6:
                root = "/".join(parts[:5])
                roots.add(root)

    return roots


def _build_outputs(roots: Set[str]) -> Tuple[List[Dict[str, str]], List[str]]:
    solutions: List[Dict[str, str]] = []
    paths: List[str] = []

    for root in roots:
        root = root.rstrip("/")

        parts = root.split("/")
        if len(parts) != 5:
            continue
        stack, platform, framework, typ, name = parts
        if stack not in STACKS:
            continue

        item = {
            "stack": stack,
            "platform": platform,
            "framework": framework,
            "type": typ,
            "name": name,
            "path": root,
        }
        solutions.append(item)
        paths.append(root)

    # Sort + unique for stability (igual ao jq)
    solutions = sorted({s["path"]: s for s in solutions}.values(), key=lambda x: x["path"])
    paths = sorted(set(paths))

    return solutions, paths


def _json_compact(value: Any) -> str:
    return json.dumps(value, separators=(",", ":"), ensure_ascii=False)


def _load_manager_inventory_by_path():
    """
    Retorna dict[path -> inventoryItem]
    """
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

    inventory = data.get("solutions") or []
    if not isinstance(inventory, list):
        raise SystemExit("Inventory inválido: esperado array em solutions")

    def _missing_base_fields(sol: Dict[str, Any]) -> List[str]:
        missing: List[str] = []
        for k in ("path", "stack", "platform", "framework", "type", "name", "status"):
            v = sol.get(k)
            if v is None or (isinstance(v, str) and not v.strip()):
                missing.append(k)
        return missing

    by_path: Dict[str, Dict[str, Any]] = {}
    for idx, raw in enumerate(inventory):
        item = dict(raw or {})
        p = (item.get("path") or "").strip()
        if not p:
            raise SystemExit(f"Inventory inválido: solution sem 'path' (index={idx})")

        parts = [seg for seg in p.split("/") if seg]
        if len(parts) != 5:
            raise SystemExit(
                f"Inventory inválido: path deve ter 5 níveis stack/platform/framework/type/name: '{p}'"
            )

        missing = _missing_base_fields(item)
        if missing:
            raise SystemExit(
                f"Inventory inválido: solution '{p}' com campos obrigatórios ausentes: {', '.join(missing)}"
            )

        # Mantém a estrutura exatamente como definida no inventory.
        # Não achata (flatten) campos de `specification` para o topo.

        by_path[p] = item

    return by_path


def _enrich_solutions_with_inventory(solutions):
    """
    solutions: list[dict] com pelo menos {"path": "..."}
    """
    inv_by_path = _load_manager_inventory_by_path()
    enriched = []
    missing = []

    for s in (solutions or []):
        s = dict(s or {})
        p = (s.get("path") or "").strip()
        inv = inv_by_path.get(p)

        if not inv:
            missing.append(p)
            continue

        # Fonte de verdade: inventory (já contém specification + flatten compat)
        merged = dict(inv)
        enriched.append(merged)

    if missing:
        missing_list = "\n".join([f"- {p}" for p in sorted(set(missing))])
        raise SystemExit(
            "As seguintes solutions foram alteradas, mas não existem no inventory:\n"
            f"{missing_list}\n"
            "Atualize architecture/governance/config/manager-solutions-inventory.json"
        )

    return enriched


def main() -> int:
    parser = argparse.ArgumentParser(description="Detect modified backend/frontend solutions in a monorepo.")
    parser.add_argument("--base", default=os.environ.get("INPUT_BASE_SHA", ""), help="Base commit SHA (optional)")
    parser.add_argument("--head", default=os.environ.get("INPUT_HEAD_SHA", ""), help="Head commit SHA (optional)")
    parser.add_argument(
        "--stack",
        default=os.environ.get("INPUT_STACK", ""),
        help="Optional stack filter (backend|frontend|orchestration)",
    )
    args = parser.parse_args()

    base, head = _infer_base_head(args.base, args.head)

    print(f"Using BASE={base}")
    print(f"Using HEAD={head}")

    changed = _get_changed_files(base, head)
    roots = _roots_from_changed_files(changed)

    stack_filter = (args.stack or "").strip().lower()
    if stack_filter:
        if stack_filter not in STACKS:
            raise SystemExit(f"Stack inválida: '{args.stack}'. Esperado backend|frontend|orchestration")
        roots = {r for r in roots if r.startswith(f"{stack_filter}/")}

    solutions, paths = _build_outputs(roots)

    solutions = _enrich_solutions_with_inventory(solutions)

    solutions_json = _json_compact(solutions)
    paths_json = _json_compact(paths)

    github_output = os.environ.get("GITHUB_OUTPUT", "")
    if github_output:
        with open(github_output, "a", encoding="utf-8") as f:
            f.write(f"solutions={solutions_json}\n")
            f.write(f"paths={paths_json}\n")

    print("Discovered solutions:")
    print(json.dumps(solutions, indent=2, ensure_ascii=False))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
