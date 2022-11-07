package biz

import "time"

// IssueEventAction enumerates the triggers for this
// webhook payload type. See also:
// https://developer.github.com/v3/activity/events/types/#issuesevent
type IssueEventAction string

// RepoPermissions describes which permission level an entity has in a
// repo. At most one of the booleans here should be true.
type RepoPermissions struct {
	// Pull is equivalent to "Read" permissions in the web UI
	Pull   bool `json:"pull"`
	Triage bool `json:"triage"`
	// Push is equivalent to "Edit" permissions in the web UI
	Push     bool `json:"push"`
	Maintain bool `json:"maintain"`
	Admin    bool `json:"admin"`
}

// User is a GitHub user account.
type User struct {
	Login       string          `json:"login"`
	Name        string          `json:"name"`
	Email       string          `json:"email"`
	ID          int             `json:"id"`
	HTMLURL     string          `json:"html_url"`
	Permissions RepoPermissions `json:"permissions"`
	Type        string          `json:"type"`
}

// Label describes a GitHub label.
type Label struct {
	URL         string `json:"url"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Color       string `json:"color"`
}

// Milestone is a milestone defined on a github repository
type Milestone struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
}

// Issue represents general info about an issue.
type Issue struct {
	ID        int       `json:"id"`
	NodeID    string    `json:"node_id"`
	User      User      `json:"user"`
	Number    int       `json:"number"`
	Title     string    `json:"title"`
	State     string    `json:"state"`
	HTMLURL   string    `json:"html_url"`
	Labels    []Label   `json:"labels"`
	Assignees []User    `json:"assignees"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Milestone Milestone `json:"milestone"`

	// This will be non-nil if it is a pull request.
	PullRequest *struct{} `json:"pull_request,omitempty"`
}

// ParentRepo contains a small subsection of general repository information: it
// just includes the information needed to confirm that a parent repo exists
// and what the name of that repo is.
type ParentRepo struct {
	Owner    User   `json:"owner"`
	Name     string `json:"name"`
	FullName string `json:"full_name"`
	HTMLURL  string `json:"html_url"`
}

// Repo contains general repository information: it includes fields available
// in repo records returned by GH "List" methods but not those returned by GH
// "Get" method.
// See also https://developer.github.com/v3/repos/#list-organization-repositories
type Repo struct {
	Owner         User   `json:"owner"`
	Name          string `json:"name"`
	FullName      string `json:"full_name"`
	HTMLURL       string `json:"html_url"`
	Fork          bool   `json:"fork"`
	DefaultBranch string `json:"default_branch"`
	Archived      bool   `json:"archived"`
	Private       bool   `json:"private"`
	Description   string `json:"description"`
	Homepage      string `json:"homepage"`
	HasIssues     bool   `json:"has_issues"`
	HasProjects   bool   `json:"has_projects"`
	HasWiki       bool   `json:"has_wiki"`
	NodeID        string `json:"node_id"`
	// Permissions reflect the permission level for the requester, so
	// on a repository GET call this will be for the user whose token
	// is being used, if listing a team's repos this will be for the
	// team's privilege level in the repo
	Permissions RepoPermissions `json:"permissions"`
	Parent      ParentRepo      `json:"parent"`
}

// IssueEvent represents an issue event from a webhook payload (not from the events API).
type IssueEvent struct {
	Action IssueEventAction `json:"action"`
	Issue  Issue            `json:"issue"`
	Repo   Repo             `json:"repository"`
	// Label is specified for IssueActionLabeled and IssueActionUnlabeled events.
	Label  Label `json:"label"`
	Sender User  `json:"sender"`

	// GUID is included in the header of the request received by GitHub.
	GUID string
}
