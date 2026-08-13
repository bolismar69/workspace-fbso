"""Confluence page creation and update operations."""
import html
import os
from jctk.config import confluence_api, get_site


def create_page(space_key, title, body, parent_id=None, representation="storage"):
    """Create a Confluence page.

    Args:
        space_key: Space key (e.g., 'NEPF')
        title: Page title
        body: Page content (HTML for storage, wiki markup for wiki)
        parent_id: Parent page ID (optional)
        representation: 'storage' (HTML) or 'wiki' (Confluence wiki markup)

    Returns:
        dict: Created page with id, _links
    """
    data = {
        "type": "page",
        "title": title,
        "space": {"key": space_key},
        "body": {
            "storage": {
                "value": body,
                "representation": representation
            }
        }
    }
    if parent_id:
        data["ancestors"] = [{"id": parent_id}]

    return confluence_api("/rest/api/content", method="POST", data=data)


def get_page(page_id):
    """Get a Confluence page by ID.

    Args:
        page_id: Page ID

    Returns:
        dict: Page with body.storage.value, version, title
    """
    return confluence_api(
        f"/rest/api/content/{page_id}?expand=body.storage,version"
    )


def update_page(page_id, title, body, representation="storage"):
    """Update an existing Confluence page.

    Args:
        page_id: Page ID
        title: Page title (must match existing or will rename)
        body: New page content
        representation: 'storage' (HTML) or 'wiki'

    Returns:
        dict: Updated page
    """
    # Get current version first
    current = get_page(page_id)
    if "_error" in current:
        return {"error": f"Cannot get page {page_id}: {current.get('_raw', '?')}"}

    version = current.get("version", {}).get("number", 1)

    return confluence_api(
        f"/rest/api/content/{page_id}",
        method="PUT",
        data={
            "version": {"number": version + 1},
            "type": "page",
            "title": title,
            "body": {
                "storage": {
                    "value": body,
                    "representation": representation
                }
            }
        }
    )


def upload_markdown_file(filepath, space_key, title, parent_id):
    """Upload a markdown file to Confluence as HTML storage page.

    Escapes HTML entities and wraps in a styled div/pre to avoid
    Confluence macro interpretation errors.

    Args:
        filepath: Path to .md file
        space_key: Confluence space key
        title: Page title
        parent_id: Parent page ID

    Returns:
        dict: Created page
    """
    if not os.path.exists(filepath):
        return {"error": f"File not found: {filepath}"}

    with open(filepath) as f:
        raw = f.read()

    # Truncate if too large
    if len(raw) > 80000:
        raw = raw[:80000] + '\n\n... (truncated)'

    # Escape HTML and wrap
    escaped = html.escape(raw)
    has_code_blocks = raw.count('```') > 3
    tag = 'pre' if has_code_blocks else 'div'
    style = 'white-space: pre-wrap; font-family: monospace; font-size: 13px; line-height: 1.5;'
    html_body = f'<{tag} style="{style}">{escaped}</{tag}>'

    return create_page(space_key, title, html_body, parent_id, representation="storage")


def add_jira_footer(page_id, jira_key, jira_summary):
    """Add a 🔗 Jira footer link to the bottom of a Confluence page.

    Gets current page content, appends footer, updates page.

    Args:
        page_id: Confluence page ID
        jira_key: Jira issue key (e.g., 'PRJSHIELD-2')
        jira_summary: Human-readable summary (e.g., 'F1 — Negócios & Discovery')

    Returns:
        dict: Update result
    """
    site = get_site()
    footer = (
        f'<hr /><p>🔗 <strong>Jira:</strong> '
        f'<a href="https://{site}/browse/{jira_key}">{jira_key} — {jira_summary}</a></p>'
    )

    current = get_page(page_id)
    if "_error" in current:
        return {"error": f"Cannot get page {page_id}"}

    title = current.get("title", "")
    current_body = current.get("body", {}).get("storage", {}).get("value", "")

    # Skip if footer already exists
    if '🔗 <strong>Jira:</strong>' in current_body:
        return {"status": "skipped", "reason": "footer already exists"}

    new_body = current_body + '\n' + footer
    return update_page(page_id, title, new_body)


def list_child_pages(parent_id, space_key=None):
    """List child pages of a parent page.

    Args:
        parent_id: Parent page ID
        space_key: Optional space key filter

    Returns:
        list: Child pages
    """
    # Use CQL-like search via REST API
    result = confluence_api(
        f"/rest/api/content/{parent_id}/child/page"
    )
    return result.get("results", [])
