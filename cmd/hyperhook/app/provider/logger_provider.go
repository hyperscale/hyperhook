// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package provider

import (
	stdlog "log"
	"os"

	"github.com/euskadi31/go-service"
	"github.com/hyperscale/hyperhook/cmd/hyperhook/app/config"
	"github.com/hyperscale/hyperhook/pkg/version"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// Logger Services keys.
const (
	LoggerKey = "logger"
)

// LoggerServiceProvider struct.
type LoggerServiceProvider struct {
}

// NewLoggerServiceProvider constructor.
func NewLoggerServiceProvider() *LoggerServiceProvider {
	return &LoggerServiceProvider{}
}

// Register implements application.ServiceProvider.
func (p LoggerServiceProvider) Register(app service.Container) {
	app.Set(LoggerKey, func(c service.Container) interface{} {
		var cfg *config.Configuration

		c.Fill(ConfigKey, cfg)

		zerolog.SetGlobalLevel(cfg.Logger.Level())

		logger := zerolog.New(os.Stdout).With().
			Timestamp().
			Str("role", cfg.Logger.Prefix).
			Str("version", version.Get().Version).
			Str("env", cfg.Environment.String()).
			Logger()

		fi, err := os.Stdin.Stat()
		if err != nil {
			log.Fatal().Err(err).Msg("Stdin.Stat failed")
		}

		if (fi.Mode() & os.ModeCharDevice) != 0 {
			logger = logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		}

		stdlog.SetFlags(0)
		stdlog.SetOutput(logger)

		log.Logger = logger

		return logger
	})

	_ = app.Get(LoggerKey)
}
