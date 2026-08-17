"""Authentication and site configuration for Atlassian APIs."""
import os
import base64

# Path to API token file (one token per line or raw)
TOKEN_FILE = os.path.expanduser("/tmp/jira-token.txt")

# Default Atlassian site
DEFAULT_SITE = "bolismar69.atlassian.net"
DEFAULT_EMAIL = "bolismar69@gmail.com"
DEFAULT_ACCOUNT_ID = "712020:bed45ac6-4894-4b1d-909c-2a27c5c4653a"


def get_token():
    """Read API token from file."""
    with open(TOKEN_FILE) as f:
        return f.read().strip()


def get_auth_header():
    """Return Basic auth header value using Python's built-in base64 module."""
    token = get_token()
    credentials = f"{DEFAULT_EMAIL}:{token}"
    return base64.b64encode(credentials.encode()).decode()


def jira_api(path, method="GET", data=None):
    """Make a Jira REST API call (v3)."""
    import subprocess, json as j
    auth = get_auth_header()
    args = [
        "curl", "-s",
        "-H", f"Authorization: Basic {auth}",
        "-H", "Content-Type: application/json",
        "-X", method,
        f"https://{DEFAULT_SITE}/rest/api/3{path}"
    ]
    if data:
        args.extend(["-d", j.dumps(data)])
    r = subprocess.run(args, capture_output=True, text=True)
    try:
        return j.loads(r.stdout)
    except:
        return {"_raw": r.stdout[:200], "_error": True}


def confluence_api(path, method="GET", data=None):
    """Make a Confluence REST API call."""
    import subprocess, json as j
    auth = get_auth_header()
    args = [
        "curl", "-s",
        "-H", f"Authorization: Basic {auth}",
        "-H", "Content-Type: application/json",
        "-X", method,
        f"https://{DEFAULT_SITE}/wiki{path}"
    ]
    if data:
        args.extend(["-d", j.dumps(data)])
    r = subprocess.run(args, capture_output=True, text=True)
    try:
        return j.loads(r.stdout)
    except:
        return {"_raw": r.stdout[:200], "_error": True}


def get_site():
    """Return default site URL."""
    return DEFAULT_SITE


def get_account_id():
    """Return default lead account ID."""
    return DEFAULT_ACCOUNT_ID
