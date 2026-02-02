#!/usr/bin/env python3
import argparse
import json
import os
import re
import subprocess
from typing import Any, Dict, List, Optional, Set, Tuple


STACKS = {"backend", "frontend", "orchestration"}

# Unified structure: stack/platform/framework/type/name
# We detect file changes that happen inside a solution folder and normalize to exactly 5 levels.
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
    # Output contract for downstream jobs: list of objects with only {path}.
    # We intentionally do NOT validate the paths against the inventory here.
    paths = sorted({r.rstrip("/") for r in roots if r.rstrip("/")})
    solutions: List[Dict[str, str]] = [{"path": p} for p in paths]
    return solutions, paths


def _json_compact(value: Any) -> str:
    return json.dumps(value, separators=(",", ":"), ensure_ascii=False)


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

    solutions_json = _json_compact(solutions)
    paths_json = _json_compact(paths)

    github_output = os.environ.get("GITHUB_OUTPUT", "")
    if github_output:
        with open(github_output, "a", encoding="utf-8") as f:
            f.write(f"solutions={solutions_json}\n")
            f.write(f"paths={paths_json}\n")

    print("Detected solutions (paths only):")
    print(json.dumps(solutions, indent=2, ensure_ascii=False))
    return 0


if __name__ == "__main__":
    raise SystemExit(main())
