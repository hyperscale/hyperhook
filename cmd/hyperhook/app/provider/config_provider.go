// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package provider

import (
	"flag"
	"os"
	"strings"
	"time"

	appprovider "github.com/euskadi31/go-application/provider"
	"github.com/euskadi31/go-server"
	"github.com/euskadi31/go-service"
	"github.com/hyperscale/hyperhook/cmd/hyperhook/app/config"
	"github.com/hyperscale/hyperhook/pkg/environment"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// HTTP Services keys.
const (
	ConfigKey = "app.config"
)

const name = "hyperhook"

// ConfigServiceProvider struct.
type ConfigServiceProvider struct {
}

// NewConfigServiceProvider constructor.
func NewConfigServiceProvider() *ConfigServiceProvider {
	return &ConfigServiceProvider{}
}

// Register implements application.ServiceProvider.
func (p *ConfigServiceProvider) Register(app service.Container) {
	app.Set(ConfigKey, func(c service.Container) interface{} {
		cfg := config.NewConfiguration()

		var cfgFile string
		var env string

		var cmd *flag.FlagSet

		c.Fill(FlagKey, cmd)

		// cmd := c.Get(FlagKey).(*flag.FlagSet)

		cmd.StringVar(&cfgFile, "config", "", "config file (default is $HOME/config.yaml)")
		cmd.StringVar(&env, "env", "dev", "Environment {dev, preprod, prod}")

		// Ignore errors; cmd is set for ExitOnError.
		// nolint:gosec
		_ = cmd.Parse(os.Args[1:])

		options := viper.New()

		options.SetDefault("environment", "dev")
		options.SetDefault("logger.level", "debug")
		options.SetDefault("logger.prefix", name)
		options.SetDefault("server.http.host", "")
		options.SetDefault("server.http.port", 8088)
		options.SetDefault("server.read_header_timeout", 10*time.Millisecond)
		options.SetDefault("server.read_timeout", 10*time.Second)
		options.SetDefault("server.shutdown_timeout", 10*time.Second)
		options.SetDefault("server.write_timeout", 10*time.Second)
		options.SetDefault("server.metrics", true)
		options.SetDefault("server.healthcheck", true)

		options.SetConfigName("config") // name of config file (without extension)

		options.AddConfigPath("/etc/" + name + "/")   // path to look for the config file in
		options.AddConfigPath("$HOME/." + name + "/") // call multiple times to add many search paths
		options.AddConfigPath(".")

		if cfgFile != "" { // enable ability to specify config file via flag
			options.SetConfigFile(cfgFile)
		}

		if environment := os.Getenv("ENV"); environment != "" {
			if err := os.Setenv("HYPERHOOK_ENVIRONMENT", environment); err != nil {
				log.Error().Err(err).Msg("Setenv failed")
			}
		}

		options.SetEnvPrefix("HYPERHOOK")
		options.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		options.AutomaticEnv() // read in environment variables that match

		// If a config file is found, read it in.
		if err := options.ReadInConfig(); err == nil {
			log.Info().Msgf("Using config file: %s", options.ConfigFileUsed())
		}

		if err := options.Unmarshal(&cfg); err != nil {
			log.Fatal().Err(err).Msg(ConfigKey)
		}

		cfg.Environment = environment.FromString(env)

		return cfg
	})

	app.Extend(appprovider.HTTPServerConfigKey, func(old *server.Configuration, c service.Container) interface{} {
		var cfg *config.Configuration

		c.Fill(ConfigKey, cfg)

		return cfg.Server
	})
}
