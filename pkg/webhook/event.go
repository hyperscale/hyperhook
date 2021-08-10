// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package webhook

import "encoding/json"

type SourceType string

const (
	SourceTypeGithub    SourceType = "github"
	SourceTypeBitbucket SourceType = "bitbucket"
	SourceTypeGitlab    SourceType = "gitlab"
)

type Event struct {
	ID      string          `json:"id"`
	Source  SourceType      `json:"source"`
	Payload json.RawMessage `json:"payload"`
}
