"""Jira project operations."""
from jctk.config import jira_api, get_account_id


def get_project(project_key):
    """Get project details by key. Returns dict with keys: id, key, name, issueTypes."""
    return jira_api(f"/project/{project_key}")


def create_project(key, name, project_type="software", lead_account_id=None):
    """Create a new Jira project (company-managed).

    Args:
        key: Project key (e.g., 'SHIELD')
        name: Project display name
        project_type: 'software', 'business', or 'service_desk'
        lead_account_id: Atlassian account ID for project lead

    Returns:
        dict: API response with project details or error
    """
    if lead_account_id is None:
        lead_account_id = get_account_id()

    # Try common template keys
    templates = [
        "com.pyxis.greenhopper.jira:gh-simplified-kanban",
        "com.pyxis.greenhopper.jira:gh-simplified-scrum",
        "com.atlassian.jira-core-project-templates:jira-core-project-management",
    ]

    for template in templates:
        result = jira_api("/project", method="POST", data={
            "key": key,
            "name": name,
            "projectTypeKey": project_type,
            "leadAccountId": lead_account_id,
            "projectTemplateKey": template,
        })
        if "id" in result and "errors" not in result:
            return result

    # Try without template
    return jira_api("/project", method="POST", data={
        "key": key,
        "name": name,
        "projectTypeKey": project_type,
        "leadAccountId": lead_account_id,
    })


def list_projects():
    """List all visible Jira projects."""
    result = jira_api("/project/search")
    return result.get("values", [])


def get_issue_types(project_key):
    """Get available issue types for a project. Returns list of {id, name, subtask, hierarchyLevel}."""
    project = get_project(project_key)
    return project.get("issueTypes", [])
