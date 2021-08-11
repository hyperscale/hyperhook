// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package bitbucket

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhookParsing(t *testing.T) {
	body, err := os.ReadFile("./testdata/bitbucket_webhook.json")
	assert.NoError(t, err)

	var webhook *Webhook
	err = json.Unmarshal(body, &webhook)
	assert.NoError(t, err)

	assert.Equal(t, "my-repo", webhook.Repository.Name)
	assert.Equal(t, "git", webhook.Repository.SCM)
	assert.Equal(t, "my-org/my-repo", webhook.Repository.FullName)
	assert.Equal(t, "repository", webhook.Repository.Type)
	assert.True(t, webhook.Repository.IsPrivate)
	assert.Equal(t, "user", webhook.Actor.Type)
	assert.Equal(t, 1, len(webhook.Push.Changes))
	assert.Equal(t, "develop", webhook.Push.Changes[0].Old.Name)
	assert.Equal(t, "develop", webhook.Push.Changes[0].New.Name)
}
