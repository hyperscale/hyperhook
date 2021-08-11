// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package provider

import (
	"fmt"
	"net/http"

	appprovider "github.com/euskadi31/go-application/provider"
	"github.com/euskadi31/go-server"
	"github.com/euskadi31/go-server/locale"
	"github.com/euskadi31/go-server/response"
	"github.com/euskadi31/go-service"
	hlogger "github.com/hyperscale/hyperhook/pkg/logger"
	"github.com/rs/cors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
)

// RouterServiceProvider struct.
type RouterServiceProvider struct {
}

// NewRouterServiceProvider constructor.
func NewRouterServiceProvider() *RouterServiceProvider {
	return &RouterServiceProvider{}
}

// Register implements application.ServiceProvider.
func (p RouterServiceProvider) Register(app service.Container) {
	app.Extend(appprovider.HTTPServerKey, func(router *server.Server, c service.Container) interface{} {
		var logger zerolog.Logger
		c.Fill(LoggerKey, logger)

		// logger := c.Get(LoggerKey).(zerolog.Logger)

		var webhookController server.Controller

		c.Fill(WebhookControllerKey, webhookController)

		// webhookController := c.Get(WebhookControllerKey).(server.Controller)

		router.Use(hlog.NewHandler(logger))
		router.Use(hlog.AccessHandler(hlogger.Handler))
		router.Use(hlog.RemoteAddrHandler("ip"))
		router.Use(hlog.UserAgentHandler("user_agent"))
		router.Use(hlog.RefererHandler("referer"))
		router.Use(hlog.RequestIDHandler("req_id", "Request-Id"))
		router.Use(locale.HandlerWithConfig([]string{"en", "fr"}))

		router.EnableCorsWithOptions(cors.Options{
			AllowCredentials: true,
			AllowedOrigins:   []string{"*"},
			AllowedMethods: []string{
				http.MethodGet,
				http.MethodOptions,
				http.MethodPost,
				http.MethodPut,
				http.MethodDelete,
			},
			AllowedHeaders: []string{
				"Authorization",
				"Content-Type",
				"X-Requested-With",
			},
		})

		router.SetNotFoundFunc(func(w http.ResponseWriter, r *http.Request) {
			response.Encode(w, r, http.StatusNotFound, map[string]interface{}{
				"error": map[string]interface{}{
					"message": fmt.Sprintf("%s %s not found", r.Method, r.URL.Path),
				},
			})
		})

		router.AddController(webhookController)

		return router // *router *server.Server
	})
}
