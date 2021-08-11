// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package github

type Repository struct {
	ID               int      `json:"id"`
	NodeID           string   `json:"node_id"`
	Name             string   `json:"name"`
	FullName         string   `json:"full_name"`
	Private          bool     `json:"private"`
	Owner            *Owner   `json:"owner"`
	HTMLURL          string   `json:"html_url"`
	Description      string   `json:"description"`
	Fork             bool     `json:"fork"`
	URL              string   `json:"url"`
	ForksURL         string   `json:"forks_url"`
	KeysURL          string   `json:"keys_url"`
	CollaboratorsURL string   `json:"collaborators_url"`
	TeamsURL         string   `json:"teams_url"`
	HooksURL         string   `json:"hooks_url"`
	IssueEventsURL   string   `json:"issue_events_url"`
	EventsURL        string   `json:"events_url"`
	AssigneesURL     string   `json:"assignees_url"`
	BranchesURL      string   `json:"branches_url"`
	TagsURL          string   `json:"tags_url"`
	BlobsURL         string   `json:"blobs_url"`
	GitTagsURL       string   `json:"git_tags_url"`
	GitRefsURL       string   `json:"git_refs_url"`
	TreesURL         string   `json:"trees_url"`
	StatusesURL      string   `json:"statuses_url"`
	LanguagesURL     string   `json:"languages_url"`
	StargazersURL    string   `json:"stargazers_url"`
	ContributorsURL  string   `json:"contributors_url"`
	SubscribersURL   string   `json:"subscribers_url"`
	SubscriptionURL  string   `json:"subscription_url"`
	CommitsURL       string   `json:"commits_url"`
	GitCommitsURL    string   `json:"git_commits_url"`
	CommentsURL      string   `json:"comments_url"`
	IssueCommentURL  string   `json:"issue_comment_url"`
	ContentsURL      string   `json:"contents_url"`
	CompareURL       string   `json:"compare_url"`
	MergesURL        string   `json:"merges_url"`
	ArchiveURL       string   `json:"archive_url"`
	DownloadsURL     string   `json:"downloads_url"`
	IssuesURL        string   `json:"issues_url"`
	PullsURL         string   `json:"pulls_url"`
	MilestonesURL    string   `json:"milestones_url"`
	NotificationsURL string   `json:"notifications_url"`
	LabelsURL        string   `json:"labels_url"`
	ReleasesURL      string   `json:"releases_url"`
	DeploymentsURL   string   `json:"deployments_url"`
	CreatedAt        int64    `json:"created_at"`
	UpdatedAt        string   `json:"updated_at"`
	PushedAt         int64    `json:"pushed_at"`
	GitURL           string   `json:"git_url"`
	SSHURL           string   `json:"ssh_url"`
	CloneURL         string   `json:"clone_url"`
	SVNURL           string   `json:"svn_url"`
	Homepage         string   `json:"homepage"`
	Size             int64    `json:"size"`
	StargazersCount  int      `json:"stargazers_count"`
	WatchersCount    int      `json:"watchers_count"`
	Language         string   `json:"language"`
	HasIssues        bool     `json:"has_issues"`
	HasProjects      bool     `json:"has_projects"`
	HasDownloads     bool     `json:"has_downloads"`
	HasWiki          bool     `json:"has_wiki"`
	HasPages         bool     `json:"has_pages"`
	ForksCount       int      `json:"forks_count"`
	MirrorURL        string   `json:"mirror_url"`
	Archived         bool     `json:"archived"`
	Disabled         bool     `json:"disabled"`
	OpenIssuesCount  int      `json:"open_issues_count"`
	License          *License `json:"license"`
	Forks            int      `json:"forks"`
	OpenIssues       int      `json:"open_issues"`
	Watchers         int      `json:"watchers"`
	DefaultBranch    string   `json:"default_branch"`
	Stargazers       int      `json:"stargazers"`
	MasterBranch     string   `json:"master_branch"`
	Organization     string   `json:"organization"`
}

/*
{
	"id": 101220964,
	"node_id": "MDEwOlJlcG9zaXRvcnkxMDEyMjA5NjQ=",
	"name": "my-repo",
	"full_name": "my-org/my-repo",
	"private": false,
	"owner": {
		"name": "my-org",
		"email": null,
		"login": "my-org",
		"id": 26484091,
		"node_id": "MDEyOk9yZ2FuaXphdGlvbjI2NDg0MDkx",
		"avatar_url": "https://avatars.githubusercontent.com/u/26484091?v=4",
		"gravatar_id": "",
		"url": "https://api.github.com/users/my-org",
		"html_url": "https://github.com/my-org",
		"followers_url": "https://api.github.com/users/my-org/followers",
		"following_url": "https://api.github.com/users/my-org/following{/other_user}",
		"gists_url": "https://api.github.com/users/my-org/gists{/gist_id}",
		"starred_url": "https://api.github.com/users/my-org/starred{/owner}{/repo}",
		"subscriptions_url": "https://api.github.com/users/my-org/subscriptions",
		"organizations_url": "https://api.github.com/users/my-org/orgs",
		"repos_url": "https://api.github.com/users/my-org/repos",
		"events_url": "https://api.github.com/users/my-org/events{/privacy}",
		"received_events_url": "https://api.github.com/users/my-org/received_events",
		"type": "Organization",
		"site_admin": false
	},
	"html_url": "https://github.com/my-org/my-repo",
	"description": "My description",
	"fork": false,
	"url": "https://github.com/my-org/my-repo",
	"forks_url": "https://api.github.com/repos/my-org/my-repo/forks",
	"keys_url": "https://api.github.com/repos/my-org/my-repo/keys{/key_id}",
	"collaborators_url": "https://api.github.com/repos/my-org/my-repo/collaborators{/collaborator}",
	"teams_url": "https://api.github.com/repos/my-org/my-repo/teams",
	"hooks_url": "https://api.github.com/repos/my-org/my-repo/hooks",
	"issue_events_url": "https://api.github.com/repos/my-org/my-repo/issues/events{/number}",
	"events_url": "https://api.github.com/repos/my-org/my-repo/events",
	"assignees_url": "https://api.github.com/repos/my-org/my-repo/assignees{/user}",
	"branches_url": "https://api.github.com/repos/my-org/my-repo/branches{/branch}",
	"tags_url": "https://api.github.com/repos/my-org/my-repo/tags",
	"blobs_url": "https://api.github.com/repos/my-org/my-repo/git/blobs{/sha}",
	"git_tags_url": "https://api.github.com/repos/my-org/my-repo/git/tags{/sha}",
	"git_refs_url": "https://api.github.com/repos/my-org/my-repo/git/refs{/sha}",
	"trees_url": "https://api.github.com/repos/my-org/my-repo/git/trees{/sha}",
	"statuses_url": "https://api.github.com/repos/my-org/my-repo/statuses/{sha}",
	"languages_url": "https://api.github.com/repos/my-org/my-repo/languages",
	"stargazers_url": "https://api.github.com/repos/my-org/my-repo/stargazers",
	"contributors_url": "https://api.github.com/repos/my-org/my-repo/contributors",
	"subscribers_url": "https://api.github.com/repos/my-org/my-repo/subscribers",
	"subscription_url": "https://api.github.com/repos/my-org/my-repo/subscription",
	"commits_url": "https://api.github.com/repos/my-org/my-repo/commits{/sha}",
	"git_commits_url": "https://api.github.com/repos/my-org/my-repo/git/commits{/sha}",
	"comments_url": "https://api.github.com/repos/my-org/my-repo/comments{/number}",
	"issue_comment_url": "https://api.github.com/repos/my-org/my-repo/issues/comments{/number}",
	"contents_url": "https://api.github.com/repos/my-org/my-repo/contents/{+path}",
	"compare_url": "https://api.github.com/repos/my-org/my-repo/compare/{base}...{head}",
	"merges_url": "https://api.github.com/repos/my-org/my-repo/merges",
	"archive_url": "https://api.github.com/repos/my-org/my-repo/{archive_format}{/ref}",
	"downloads_url": "https://api.github.com/repos/my-org/my-repo/downloads",
	"issues_url": "https://api.github.com/repos/my-org/my-repo/issues{/number}",
	"pulls_url": "https://api.github.com/repos/my-org/my-repo/pulls{/number}",
	"milestones_url": "https://api.github.com/repos/my-org/my-repo/milestones{/number}",
	"notifications_url": "https://api.github.com/repos/my-org/my-repo/notifications{?since,all,participating}",
	"labels_url": "https://api.github.com/repos/my-org/my-repo/labels{/name}",
	"releases_url": "https://api.github.com/repos/my-org/my-repo/releases{/id}",
	"deployments_url": "https://api.github.com/repos/my-org/my-repo/deployments",
	"created_at": 1503519213,
	"updated_at": "2021-08-04T08:04:38Z",
	"pushed_at": 1628634973,
	"git_url": "git://github.com/my-org/my-repo.git",
	"ssh_url": "git@github.com:my-org/my-repo.git",
	"clone_url": "https://github.com/my-org/my-repo.git",
	"svn_url": "https://github.com/my-org/my-repo",
	"homepage": "",
	"size": 1890,
	"stargazers_count": 1,
	"watchers_count": 1,
	"language": "Go",
	"has_issues": true,
	"has_projects": true,
	"has_downloads": true,
	"has_wiki": true,
	"has_pages": false,
	"forks_count": 1,
	"mirror_url": null,
	"archived": false,
	"disabled": false,
	"open_issues_count": 0,
	"license": {
		"key": "mit",
		"name": "MIT License",
		"spdx_id": "MIT",
		"url": "https://api.github.com/licenses/mit",
		"node_id": "MDc6TGljZW5zZTEz"
	},
	"forks": 1,
	"open_issues": 0,
	"watchers": 1,
	"default_branch": "develop",
	"stargazers": 1,
	"master_branch": "develop",
	"organization": "my-org"
}

*/
