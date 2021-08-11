// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package config

import (
	"testing"

	"github.com/hyperscale/hyperhook/pkg/environment"
	"github.com/stretchr/testify/assert"
)

func TestConfiguration(t *testing.T) {
	c := NewConfiguration()

	assert.NotNil(t, c)
	assert.Equal(t, environment.Dev, c.Environment)
}
