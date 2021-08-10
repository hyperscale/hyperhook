// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package github

type Owner struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	Login             string `json:"login"`
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

/*
{
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
*/
