// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package github

import "encoding/json"

type Webhook struct {
	Ref          string            `json:"ref"`
	Before       string            `json:"before"`
	After        string            `json:"after"`
	Repository   *Repository       `json:"repository"`
	Pusher       *Pusher           `json:"pusher"`
	Organization *Organization     `json:"organization"`
	Sender       *Owner            `json:"sender"`
	Created      bool              `json:"created"`
	Deleted      bool              `json:"deleted"`
	Forced       bool              `json:"forced"`
	BaseRef      json.RawMessage   `json:"base_ref"`
	Compare      string            `json:"compare"`
	Commits      []json.RawMessage `json:"commits"`
	HeadCommit   json.RawMessage   `json:"head_commit"`
}
