// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package provider

import (
	"flag"
	"os"

	"github.com/euskadi31/go-service"
)

// Flag Services keys.
const (
	FlagKey = "flag"
)

// FlagServiceProvider struct.
type FlagServiceProvider struct {
}

// NewFlagServiceProvider constructor.
func NewFlagServiceProvider() *FlagServiceProvider {
	return &FlagServiceProvider{}
}

// Register implements application.ServiceProvider.
func (p FlagServiceProvider) Register(app service.Container) {
	app.Set(FlagKey, func(c service.Container) interface{} {
		cmd := flag.NewFlagSet(os.Args[0], flag.ExitOnError)

		return cmd // *flag.FlagSet
	})
}
