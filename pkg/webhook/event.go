// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package webhook

import "encoding/json"

type SourceType string

const (
	SourceTypeUnknwon   SourceType = "unknown"
	SourceTypeGithub    SourceType = "github"
	SourceTypeBitbucket SourceType = "bitbucket"
	SourceTypeGitlab    SourceType = "gitlab"
)

type Event struct {
	ID      string          `json:"id"`
	Source  SourceType      `json:"source"`
	Payload json.RawMessage `json:"payload"`
}

func SourceFromString(source string) SourceType {
	switch source {
	case "github":
		return SourceTypeGithub
	case "bitbucket":
		return SourceTypeBitbucket
	case "gitlab":
		return SourceTypeGitlab
	default:
		return SourceTypeUnknwon
	}
}
