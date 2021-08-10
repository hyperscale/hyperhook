// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package controller

import "github.com/euskadi31/go-eventemitter"

//go:generate mockery --name=EventEmitter --inpackage --case underscore
type EventEmitter interface {
	eventemitter.EventEmitter
}
