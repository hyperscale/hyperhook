// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package github

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhookParsing(t *testing.T) {
	body, err := os.ReadFile("./testdata/github_webhook.json")
	assert.NoError(t, err)

	var webhook *Webhook
	err = json.Unmarshal(body, &webhook)
	assert.NoError(t, err)

	assert.Equal(t, "refs/heads/my-branch", webhook.Ref)
	assert.Equal(t, "bd558d42c1e368bed4d6ac8c174bcab3e3dc5827", webhook.Before)
	assert.Equal(t, "0000000000000000000000000000000000000000", webhook.After)
	assert.Equal(t, "my-repo", webhook.Repository.Name)
	assert.False(t, webhook.Repository.Private)
	assert.Equal(t, "my-org/my-repo", webhook.Repository.FullName)
	assert.Equal(t, "git@github.com:my-org/my-repo.git", webhook.Repository.SSHURL)

	assert.Equal(t, "euskadi31", webhook.Pusher.Name)

	assert.Equal(t, "my-org", webhook.Organization.Login)

	assert.Equal(t, "euskadi31", webhook.Sender.Login)

}
