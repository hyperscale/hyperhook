// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package controller

import (
	"encoding/json"
	"net/http"

	"github.com/euskadi31/go-eventemitter"
	"github.com/euskadi31/go-server"
	"github.com/euskadi31/go-server/response"
	"github.com/hyperscale/hyperhook/pkg/webhook"
	"github.com/rs/zerolog/hlog"
)

var _ server.Controller = (*webhookController)(nil)

type webhookController struct {
	emitter eventemitter.EventEmitter
}

// NewWebhookController func.
func NewWebhookController(
	emitter eventemitter.EventEmitter,
) (server.Controller, error) {
	c := &webhookController{
		emitter: emitter,
	}

	return c, nil
}

// Mount endpoints.
func (c webhookController) Mount(r *server.Router) {
	r.HandleFunc("/v1/webhook/{id}/{provider}", c.postWebhookHandler).Methods(http.MethodPost)
}

// POST /v1/webhook/{id}/{provider} .
func (c webhookController) postWebhookHandler(w http.ResponseWriter, r *http.Request) {
	webhook := &webhook.Event{}

	log := hlog.FromRequest(r)

	var inputMap map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&inputMap); err != nil {
		log.Error().Err(err).Msg("Parsing json body failed")

		response.FailureFromError(w, http.StatusBadRequest, err)

		return
	}

	c.emitter.Dispatch("webkook.receive", webhook)

	response.Encode(w, r, http.StatusCreated, map[string]bool{"status": true})
}
