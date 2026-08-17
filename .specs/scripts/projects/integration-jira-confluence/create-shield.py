#!/usr/bin/env python3
"""Orchestrator: Create full SHIELD project structure in Jira + Confluence.

Reads config.json, structure.json, and doc-mapping.json to:
1. Upload documents to Confluence
2. Create Jira issues (Epics, Stories, Tasks)
3. Create issue links (relates to, blocks)
4. Add remote links (Jira → Confluence docs)

Usage:
    cd /home/bolismar/work/workspace-fbso/.specs/scripts
    python3 -m projects.shield.create-shield [--step upload|issues|links|remote]
"""
import sys
import os
import json

# Add parent to path for jctk import
sys.path.insert(0, os.path.dirname(os.path.dirname(os.path.dirname(os.path.abspath(__file__)))))

from jctk.confluence.pages import upload_markdown_file, add_jira_footer
from jctk.jira.links import add_remote_links
from jctk.jira.issues import create_epic, create_story, create_task
from jctk.config import get_site

BASE_DIR = os.path.dirname(os.path.abspath(__file__))

def load_json(name):
    with open(os.path.join(BASE_DIR, name)) as f:
        return json.load(f)

def step_upload_docs():
    """Step 1: Upload all documents to Confluence."""
    config = load_json("config.json")
    mapping = load_json("doc-mapping.json")
    doc_base = config["doc_base_path"]
    space = config["confluence_space"]
    folders = mapping["confluence_folders"]

    all_docs = mapping.get("documents", []) + mapping.get("waterfall_estimation", [])
    print(f"Uploading {len(all_docs)} documents to Confluence...\n")

    ok = 0
    for doc in all_docs:
        filepath = os.path.join(doc_base, doc["file"])
        folder_id = folders[doc["folder"]]["id"]
        title = doc["confluence_title"]

        result = upload_markdown_file(filepath, space, title, folder_id)
        if "id" in result:
            print(f"  ✅ {result['id']} — {title}")
            ok += 1
        elif "error" in result:
            print(f"  ❌ {title}: {result['error']}")
        else:
            print(f"  ⚠️  {title}: {str(result)[:100]}")

    print(f"\nUpload complete: {ok}/{len(all_docs)}")

def step_add_jira_footers():
    """Step 2: Add Jira footer links to Confluence pages."""
    config = load_json("config.json")
    mapping = load_json("doc-mapping.json")
    site = get_site()
    space = config["confluence_space"]

    # Map folder to Jira epic
    folder_to_epic = {
        "F1": ("PRJSHIELD-2", "F1 — Negócios & Discovery"),
        "F2": ("PRJSHIELD-3", "F2 — Estimativa UPSTREAM"),
        "F3": ("PRJSHIELD-4", "F3 — Detalhamento Técnico"),
        "F4": ("PRJSHIELD-5", "F4 — Estimativa DOWNSTREAM"),
        "F5": ("PRJSHIELD-6", "F5 — GO/NO-GO & Planejamento"),
        "F6": ("PRJSHIELD-7", "F6 — Execução Técnica"),
        "F7": ("PRJSHIELD-8", "F7 — Documentação & Encerramento"),
    }

    print("Adding Jira footers to Confluence pages...\n")
    ok = 0
    for entry in mapping.get("remote_links", {}).items():
        # This is a simplified version - in practice, we'd look up page IDs
        pass

    print(f"\nFooters: {ok} pages updated")

def step_remote_links():
    """Step 3: Add Jira remote links to Confluence docs."""
    mapping = load_json("doc-mapping.json")
    remote_links = mapping.get("remote_links", {})

    print(f"Adding remote links for {len(remote_links)} issues...\n")
    ok = 0
    total = 0
    for issue_key, links in remote_links.items():
        results = add_remote_links(issue_key, links)
        for r in results:
            total += 1
            if r["success"]:
                print(f"  ✅ {issue_key} ← {r['title']}")
                ok += 1
            else:
                print(f"  ❌ {issue_key} ← {r['title']}")

    print(f"\nRemote links: {ok}/{total}")

def print_summary():
    """Print project summary."""
    config = load_json("config.json")
    struct = load_json("structure.json")
    mapping = load_json("doc-mapping.json")

    print(f"\n=== {config['project_name']} ===")
    print(f"Jira: https://{get_site()}/jira/software/c/projects/{config['project_key']}")
    print(f"Confluence: https://{get_site()}/wiki/spaces/{config['confluence_space']}/pages/{config['confluence_parent_id']}")
    print(f"\nPhases: {len(struct['phases'])}")
    print(f"Deliverables: {len(struct['deliverables'])}")
    print(f"Documents: {len(mapping['documents'])} + {len(mapping.get('waterfall_estimation',[]))} estimation")

if __name__ == "__main__":
    step = sys.argv[1] if len(sys.argv) > 1 else "summary"

    if step == "upload":
        step_upload_docs()
    elif step == "footers":
        step_add_jira_footers()
    elif step == "remote":
        step_remote_links()
    elif step == "all":
        step_upload_docs()
        step_remote_links()
    else:
        print_summary()
        print("\nUsage: python3 -m projects.shield.create-shield [upload|footers|remote|all|summary]")
