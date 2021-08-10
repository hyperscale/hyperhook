// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package config

import (
	"github.com/euskadi31/go-server"
	"github.com/hyperscale/hyperhook/pkg/environment"
	"github.com/hyperscale/hyperhook/pkg/logger"
)

// Configuration struct.
type Configuration struct {
	Environment environment.Env
	Logger      *logger.Configuration
	Server      *server.Configuration
}

// NewConfiguration constructor.
func NewConfiguration() *Configuration {
	return &Configuration{
		Environment: environment.Dev,
		Logger:      &logger.Configuration{},
		Server:      &server.Configuration{},
	}
}
