"""Jira issue linking and remote link operations."""
from jctk.config import jira_api


def link_issues(inward_key, outward_key, link_type="Relates"):
    """Create an issue link between two issues.

    Args:
        inward_key: Inward issue key (e.g., the Story)
        outward_key: Outward issue key (e.g., the Task)
        link_type: 'Blocks', 'Relates', 'Cloners', 'Duplicate'

    Returns:
        dict: Link creation result
    """
    return jira_api("/issueLink", method="POST", data={
        "type": {"name": link_type},
        "inwardIssue": {"key": inward_key},
        "outwardIssue": {"key": outward_key},
    })


def add_remote_links(issue_key, links):
    """Add web/remote links to a Jira issue.

    Each link appears as a clickable card in the issue's Links tab > Web Links section.

    Args:
        issue_key: Jira issue key (e.g., 'PRJSHIELD-7')
        links: List of (title, url) tuples

    Returns:
        list[dict]: Results for each link
    """
    results = []
    for title, url in links:
        result = jira_api(
            f"/issue/{issue_key}/remotelink",
            method="POST",
            data={"object": {"url": url, "title": title}}
        )
        results.append({
            "title": title,
            "url": url,
            "success": "id" in result,
            "response": result
        })
    return results


def link_issues_bidirectional(key_a, key_b, link_type="Relates"):
    """Create bidirectional link between two issues.

    Creates two links: key_a relates to key_b AND key_b relates to key_a.
    Effectively the same as one 'Relates' link (which is already bidirectional).

    Args:
        key_a: First issue key
        key_b: Second issue key
        link_type: Link type name

    Returns:
        dict: Link creation result
    """
    return link_issues(key_a, key_b, link_type)


def get_issue_links(issue_key):
    """Get all links for an issue.

    Args:
        issue_key: Issue key

    Returns:
        dict: Issue links data
    """
    return jira_api(f"/issue/{issue_key}?fields=issuelinks")
