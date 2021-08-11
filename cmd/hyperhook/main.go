// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"errors"

	"github.com/euskadi31/go-application"
	appprovider "github.com/euskadi31/go-application/provider"
	"github.com/hyperscale/hyperhook/cmd/hyperhook/app/provider"
	"github.com/rs/zerolog/log"
)

func main() {
	app := application.New()

	app.Register(appprovider.NewEventDispatcherServiceProvider())
	app.Register(appprovider.NewHTTPServiceProvider())
	app.Register(provider.NewFlagServiceProvider())
	app.Register(provider.NewConfigServiceProvider())
	app.Register(provider.NewLoggerServiceProvider())
	app.Register(provider.NewControllerServiceProvider())
	app.Register(provider.NewRouterServiceProvider())

	defer func() {
		if err := app.Close(); err != nil {
			log.Error().Err(err).Msg("hyperhook close failed")
		}
	}()

	if err := app.Run(); err != nil {
		if errors.Is(err, context.Canceled) {
			log.Debug().Err(err).Msg("ignore error since context is cancelled")
		} else {
			log.Error().Err(err).Msg("hyperhook run failed")
		}
	}
}
