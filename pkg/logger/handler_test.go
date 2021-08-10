// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package logger

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	Handler(req, http.StatusOK, 123, 156*time.Millisecond)
	Handler(req, http.StatusFound, 123, 156*time.Millisecond)
	Handler(req, http.StatusNotFound, 123, 156*time.Millisecond)
	Handler(req, http.StatusInternalServerError, 123, 156*time.Millisecond)
}
