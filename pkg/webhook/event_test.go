// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package webhook

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSourceFromString(t *testing.T) {
	for _, item := range []struct {
		expected SourceType
		value    string
	}{
		{
			expected: SourceTypeGithub,
			value:    "github",
		},
		{
			expected: SourceTypeBitbucket,
			value:    "bitbucket",
		},
		{
			expected: SourceTypeGitlab,
			value:    "gitlab",
		},
		{
			expected: SourceTypeUnknwon,
			value:    "bad",
		},
	} {
		assert.Equal(t, item.expected, SourceFromString(item.value))
	}
}
