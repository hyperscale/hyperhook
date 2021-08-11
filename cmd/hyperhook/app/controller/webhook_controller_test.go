// Copyright 2021 Hyperhook. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package controller

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/euskadi31/go-eventemitter"
	"github.com/euskadi31/go-server"
	"github.com/google/uuid"
	"github.com/hyperscale/hyperhook/pkg/webhook"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestWebhookBadIDHandler(t *testing.T) {
	bodyReq, err := os.ReadFile("./testdata/bitbucket_webhook.json")
	assert.NoError(t, err)

	server := server.NewRouter()

	c, err := NewWebhookController(nil)
	assert.NoError(t, err)

	server.AddController(c)

	provider := "bad"

	id := "bad"

	req := httptest.NewRequest(http.MethodPost, "https://example.com/v1/webhook/"+id+"/"+provider, bytes.NewReader(bodyReq))

	req.Header.Set("X-Event-Key", "repo:push")
	req.Header.Set("X-Hook-UUID", "f4fd20d3-7088-4597-9b60-28265aa76bbe")
	req.Header.Set("X-B3-SpanId", "3033da5659cfc94d")
	req.Header.Set("X-Event-Time", "Tue, 10 Aug 2021 21:50:44 GMT")
	req.Header.Set("User-Agent", "Bitbucket-Webhooks/2.0")
	req.Header.Set("X-B3-Sampled", "1")
	req.Header.Set("X-B3-TraceId", "3033da5659cfc94d")
	req.Header.Set("X-Attempt-Number", "1")
	req.Header.Set("X-Request-UUID", "dbddf50d-fca5-4845-8c95-1593ff233364")
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	err = resp.Body.Close()
	assert.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Equal(t, "text/plain; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.Equal(t, "404 page not found\n", string(body))
}

func TestWebhookBadProviderHandler(t *testing.T) {
	bodyReq, err := os.ReadFile("./testdata/bitbucket_webhook.json")
	assert.NoError(t, err)

	server := server.NewRouter()

	c, err := NewWebhookController(nil)
	assert.NoError(t, err)

	server.AddController(c)

	provider := "bad"

	id := uuid.New().String()

	req := httptest.NewRequest(http.MethodPost, "https://example.com/v1/webhook/"+id+"/"+provider, bytes.NewReader(bodyReq))

	req.Header.Set("X-Event-Key", "repo:push")
	req.Header.Set("X-Hook-UUID", "f4fd20d3-7088-4597-9b60-28265aa76bbe")
	req.Header.Set("X-B3-SpanId", "3033da5659cfc94d")
	req.Header.Set("X-Event-Time", "Tue, 10 Aug 2021 21:50:44 GMT")
	req.Header.Set("User-Agent", "Bitbucket-Webhooks/2.0")
	req.Header.Set("X-B3-Sampled", "1")
	req.Header.Set("X-B3-TraceId", "3033da5659cfc94d")
	req.Header.Set("X-Attempt-Number", "1")
	req.Header.Set("X-Request-UUID", "dbddf50d-fca5-4845-8c95-1593ff233364")
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	err = resp.Body.Close()
	assert.NoError(t, err)

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
	assert.Equal(t, "application/json; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.Equal(t, `{"error":{"code":404,"message":"provider not found"}}`+"\n", string(body))
}

type mockReadCloser struct {
	mock.Mock
}

func (m *mockReadCloser) Read(p []byte) (n int, err error) {
	args := m.Called(p)

	return args.Int(0), args.Error(1)
}

func (m *mockReadCloser) Close() error {
	args := m.Called()

	return args.Error(0)
}

func TestWebhookBitbucketHandlerWithBadBody(t *testing.T) {
	mockReadCloser := &mockReadCloser{}
	// if Read is called, it will return error
	mockReadCloser.On("Read", mock.AnythingOfType("[]uint8")).Return(0, fmt.Errorf("error reading"))
	// if Close is called, it will return error
	mockReadCloser.On("Close").Return(fmt.Errorf("error closing"))

	server := server.NewRouter()

	c, err := NewWebhookController(nil)
	assert.NoError(t, err)

	server.AddController(c)

	provider := "bitbucket"

	id := uuid.New().String()

	req := httptest.NewRequest(http.MethodPost, "https://example.com/v1/webhook/"+id+"/"+provider, mockReadCloser)

	req.Header.Set("X-Event-Key", "repo:push")
	req.Header.Set("X-Hook-UUID", "f4fd20d3-7088-4597-9b60-28265aa76bbe")
	req.Header.Set("X-B3-SpanId", "3033da5659cfc94d")
	req.Header.Set("X-Event-Time", "Tue, 10 Aug 2021 21:50:44 GMT")
	req.Header.Set("User-Agent", "Bitbucket-Webhooks/2.0")
	req.Header.Set("X-B3-Sampled", "1")
	req.Header.Set("X-B3-TraceId", "3033da5659cfc94d")
	req.Header.Set("X-Attempt-Number", "1")
	req.Header.Set("X-Request-UUID", "dbddf50d-fca5-4845-8c95-1593ff233364")
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	err = resp.Body.Close()
	assert.NoError(t, err)

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
	assert.Equal(t, "application/json; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.Equal(t, `{"error":{"code":400,"message":"error reading"}}`+"\n", string(body))
}

func TestWebhookBitbucketHandler(t *testing.T) {
	provider := "bitbucket"

	id := uuid.New().String()

	bodyReq, err := os.ReadFile("./testdata/bitbucket_webhook.json")
	assert.NoError(t, err)

	emitterMock := &MockEventEmitter{}

	emitterMock.On("Dispatch", "webkook.receive", mock.MatchedBy(func(webhook *webhook.Event) bool {
		if webhook.ID != id {
			t.Errorf("webhook.ID %q is not match %q", webhook.ID, id)

			return false
		}

		if string(webhook.Source) != provider {
			t.Errorf("webhook.Source %q is not match %q", webhook.Source, provider)

			return false
		}

		if !bytes.Equal(webhook.Payload, bodyReq) {
			t.Error("webhook.Payload is not valid")

			return false
		}

		return true
	})).Return()

	server := server.NewRouter()

	c, err := NewWebhookController(emitterMock)
	assert.NoError(t, err)

	server.AddController(c)

	req := httptest.NewRequest(http.MethodPost, "https://example.com/v1/webhook/"+id+"/"+provider, bytes.NewReader(bodyReq))

	req.Header.Set("X-Event-Key", "repo:push")
	req.Header.Set("X-Hook-UUID", "f4fd20d3-7088-4597-9b60-28265aa76bbe")
	req.Header.Set("X-B3-SpanId", "3033da5659cfc94d")
	req.Header.Set("X-Event-Time", "Tue, 10 Aug 2021 21:50:44 GMT")
	req.Header.Set("User-Agent", "Bitbucket-Webhooks/2.0")
	req.Header.Set("X-B3-Sampled", "1")
	req.Header.Set("X-B3-TraceId", "3033da5659cfc94d")
	req.Header.Set("X-Attempt-Number", "1")
	req.Header.Set("X-Request-UUID", "dbddf50d-fca5-4845-8c95-1593ff233364")
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	err = resp.Body.Close()
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.Equal(t, `{"status":true}`+"\n", string(body))

	emitterMock.AssertExpectations(t)
}

func TestWebhookGithubHandler(t *testing.T) {
	provider := "github"

	id := uuid.New().String()

	bodyReq, err := os.ReadFile("./testdata/github_webhook.json")
	assert.NoError(t, err)

	emitterMock := &MockEventEmitter{}

	emitterMock.On("Dispatch", "webkook.receive", mock.MatchedBy(func(webhook *webhook.Event) bool {
		if webhook.ID != id {
			t.Errorf("webhook.ID %q is not match %q", webhook.ID, id)

			return false
		}

		if string(webhook.Source) != provider {
			t.Errorf("webhook.Source %q is not match %q", webhook.Source, provider)

			return false
		}

		if !bytes.Equal(webhook.Payload, bodyReq) {
			t.Error("webhook.Payload is not valid")

			return false
		}

		return true
	})).Return()

	server := server.NewRouter()

	c, err := NewWebhookController(emitterMock)
	assert.NoError(t, err)

	server.AddController(c)

	req := httptest.NewRequest(http.MethodPost, "https://example.com/v1/webhook/"+id+"/"+provider, bytes.NewReader(bodyReq))

	req.Header.Set("Accept", "*/*")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-GitHub-Delivery", "5ec81240-fa2b-11eb-9bd5-b3ae6778bdbf")
	req.Header.Set("X-GitHub-Event", "push")
	req.Header.Set("X-GitHub-Hook-ID", "77063815")
	req.Header.Set("User-Agent", "GitHub-Hookshot/f9d234e")
	req.Header.Set("X-GitHub-Hook-Installation-Target-ID", "101220964")
	req.Header.Set("X-GitHub-Hook-Installation-Target-Type", "repository")

	w := httptest.NewRecorder()

	server.ServeHTTP(w, req)

	resp := w.Result()
	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)

	err = resp.Body.Close()
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.Equal(t, `{"status":true}`+"\n", string(body))

	emitterMock.AssertExpectations(t)
}

func BenchmarkWebhook(b *testing.B) {
	provider := "bitbucket"

	bodyReq, err := os.ReadFile("./testdata/bitbucket_webhook.json")
	assert.NoError(b, err)

	emitter := eventemitter.New()

	server := server.NewRouter()

	c, err := NewWebhookController(emitter)
	assert.NoError(b, err)

	server.AddController(c)

	for i := 0; i < b.N; i++ {
		id := uuid.New().String()

		req := httptest.NewRequest(http.MethodPost, "https://example.com/v1/webhook/"+id+"/"+provider, bytes.NewReader(bodyReq))

		req.Header.Set("X-Event-Key", "repo:push")
		req.Header.Set("X-Hook-UUID", "f4fd20d3-7088-4597-9b60-28265aa76bbe")
		req.Header.Set("X-B3-SpanId", "3033da5659cfc94d")
		req.Header.Set("X-Event-Time", "Tue, 10 Aug 2021 21:50:44 GMT")
		req.Header.Set("User-Agent", "Bitbucket-Webhooks/2.0")
		req.Header.Set("X-B3-Sampled", "1")
		req.Header.Set("X-B3-TraceId", "3033da5659cfc94d")
		req.Header.Set("X-Attempt-Number", "1")
		req.Header.Set("X-Request-UUID", "dbddf50d-fca5-4845-8c95-1593ff233364")
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()

		server.ServeHTTP(w, req)
	}
}
