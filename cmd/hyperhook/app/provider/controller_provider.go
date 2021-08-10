// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package provider

import (
	appprovider "github.com/euskadi31/go-application/provider"
	"github.com/euskadi31/go-eventemitter"
	"github.com/euskadi31/go-service"
	"github.com/hyperscale/hyperhook/cmd/hyperhook/app/controller"
	"github.com/rs/zerolog/log"
)

// Controller Services keys.
const (
	WebhookControllerKey = "controller.webhook"
)

// ControllerServiceProvider struct.
type ControllerServiceProvider struct {
}

// NewControllerServiceProvider constructor.
func NewControllerServiceProvider() *ControllerServiceProvider {
	return &ControllerServiceProvider{}
}

// Register implements application.ServiceProvider.
func (p ControllerServiceProvider) Register(app service.Container) {
	app.Set(WebhookControllerKey, func(c service.Container) interface{} {
		var emitter eventemitter.EventEmitter

		c.Fill(appprovider.EventDispatcherKey, emitter)

		// emitter := c.Get(appprovider.EventDispatcherKey).(eventemitter.EventEmitter)

		controller, err := controller.NewWebhookController(emitter)
		if err != nil {
			log.Fatal().Err(err).Msg(WebhookControllerKey)
		}

		return controller // server.Controller
	})
}
