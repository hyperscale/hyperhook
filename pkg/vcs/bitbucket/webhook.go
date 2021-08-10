// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package bitbucket

type Webhook struct {
	Actor      *Owner       `json:"actor"`
	Repository *Repository  `json:"repository"`
	Push       *PushPayload `json:"push"`
}
