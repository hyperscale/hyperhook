// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package controller

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/euskadi31/go-eventemitter"
	"github.com/euskadi31/go-server"
	"github.com/euskadi31/go-server/response"
	"github.com/gorilla/mux"
	"github.com/hyperscale/hyperhook/pkg/webhook"
	"github.com/rs/zerolog/hlog"
)

var errProviderNotFound = errors.New("provider not found")

const uuidRegxp = `\b[0-9a-f]{8}\b-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-\b[0-9a-f]{12}\b`

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
	r.HandleFunc("/v1/webhook/{id:"+uuidRegxp+"}/{provider}", c.postWebhookHandler).Methods(http.MethodPost)
}

// POST /v1/webhook/{id}/{provider} .
func (c webhookController) postWebhookHandler(w http.ResponseWriter, r *http.Request) {
	log := hlog.FromRequest(r)

	vars := mux.Vars(r)

	provider := webhook.SourceFromString(vars["provider"])
	if provider == webhook.SourceTypeUnknwon {
		log.Error().Err(errProviderNotFound).Str("provider", vars["provider"]).Msg("")

		response.FailureFromError(w, http.StatusNotFound, errProviderNotFound)

		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error().Err(err).Str("provider", vars["provider"]).Msg("read body failed")

		response.FailureFromError(w, http.StatusBadRequest, err)

		return
	}

	webhook := &webhook.Event{
		ID:      vars["id"],
		Source:  provider,
		Payload: json.RawMessage(body),
	}

	c.emitter.Dispatch("webkook.receive", webhook)

	response.Encode(w, r, http.StatusOK, map[string]bool{"status": true})
}
