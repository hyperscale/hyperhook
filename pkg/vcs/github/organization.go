// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package github

type Organization struct {
	Login            string `json:"login"`
	ID               int    `json:"id"`
	NodeID           string `json:"node_id"`
	URL              string `json:"url"`
	ReposURL         string `json:"repos_url"`
	EventsURL        string `json:"events_url"`
	HooksURL         string `json:"hooks_url"`
	IssuesURL        string `json:"issues_url"`
	MembersURL       string `json:"members_url"`
	PublicMembersURL string `json:"public_members_url"`
	AvatarURL        string `json:"avatar_url"`
	Description      string `json:"description"`
}

/*

{
	"login": "my-org",
	"id": 26484091,
	"node_id": "MDEyOk9yZ2FuaXphdGlvbjI2NDg0MDkx",
	"url": "https://api.github.com/orgs/my-org",
	"repos_url": "https://api.github.com/orgs/my-org/repos",
	"events_url": "https://api.github.com/orgs/my-org/events",
	"hooks_url": "https://api.github.com/orgs/my-org/hooks",
	"issues_url": "https://api.github.com/orgs/my-org/issues",
	"members_url": "https://api.github.com/orgs/my-org/members{/member}",
	"public_members_url": "https://api.github.com/orgs/my-org/public_members{/member}",
	"avatar_url": "https://avatars.githubusercontent.com/u/26484091?v=4",
	"description": ""
}
*/
