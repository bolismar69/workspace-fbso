"""Jira issue creation and update operations."""
from jctk.config import jira_api


def create_issue(project_key, issue_type, summary, description="", parent_key=None,
                 labels=None, priority="Medium", additional_fields=None):
    """Create any Jira issue. Generic low-level function.

    Args:
        project_key: Project key (e.g., 'PRJSHIELD')
        issue_type: Type name - 'Epic', 'História', 'Tarefa', 'Subtarefa', 'Bug'
        summary: Issue title
        description: Description text (markdown supported)
        parent_key: Parent issue key for Epic child or Subtask parent
        labels: List of label strings
        priority: 'Highest', 'High', 'Medium', 'Low', 'Lowest'
        additional_fields: Dict of extra Jira fields

    Returns:
        dict: Created issue with key, id, self URL
    """
    fields = {
        "project": {"key": project_key},
        "issuetype": {"name": issue_type},
        "summary": summary,
        "description": description,
        "priority": {"name": priority},
    }

    if labels:
        fields["labels"] = labels

    if parent_key:
        # For Subtasks, parent is the parent issue key
        # For Stories/Tasks under Epic, parent is the Epic key
        fields["parent"] = {"key": parent_key}

    if additional_fields:
        fields.update(additional_fields)

    return jira_api("/issue", method="POST", data={"fields": fields})


def create_epic(project_key, summary, description="", labels=None, priority="High"):
    """Create an Epic issue.

    Args:
        project_key: Project key
        summary: Epic name
        description: Description with markdown
        labels: List of labels
        priority: Priority level

    Returns:
        dict: Created issue
    """
    return create_issue(
        project_key=project_key,
        issue_type="Epic",
        summary=summary,
        description=description,
        labels=labels,
        priority=priority,
    )


def create_story(project_key, summary, description="", parent_epic_key=None,
                 labels=None, priority="Medium"):
    """Create a Story (História) issue.

    Args:
        project_key: Project key
        summary: Story title
        description: Description with markdown
        parent_epic_key: Epic key to parent under (e.g., 'PRJSHIELD-7')
        labels: List of labels
        priority: Priority level

    Returns:
        dict: Created issue
    """
    return create_issue(
        project_key=project_key,
        issue_type="História",
        summary=summary,
        description=description,
        parent_key=parent_epic_key,
        labels=labels,
        priority=priority,
    )


def create_task(project_key, summary, description="", parent_key=None,
                labels=None, priority="Medium"):
    """Create a Task (Tarefa) issue.

    Args:
        project_key: Project key
        summary: Task title
        description: Description with markdown
        parent_key: Parent Epic or Story key
        labels: List of labels
        priority: Priority level

    Returns:
        dict: Created issue
    """
    return create_issue(
        project_key=project_key,
        issue_type="Tarefa",
        summary=summary,
        description=description,
        parent_key=parent_key,
        labels=labels,
        priority=priority,
    )


def create_subtask(project_key, summary, description="", parent_key=None,
                   labels=None):
    """Create a Subtask (Subtarefa) under a parent issue.

    Args:
        project_key: Project key
        summary: Subtask title
        description: Description
        parent_key: REQUIRED - parent issue key
        labels: List of labels

    Returns:
        dict: Created issue
    """
    if not parent_key:
        raise ValueError("parent_key is required for subtasks")
    return create_issue(
        project_key=project_key,
        issue_type="Subtarefa",
        summary=summary,
        description=description,
        parent_key=parent_key,
        labels=labels,
        priority="Medium",
    )


def update_issue(issue_key, fields):
    """Update an existing issue's fields.

    Args:
        issue_key: Issue key (e.g., 'PRJSHIELD-7')
        fields: Dict of fields to update (e.g., {'description': '...'})

    Returns:
        dict: Updated issue
    """
    return jira_api(f"/issue/{issue_key}", method="PUT", data={"fields": fields})
