// Copyright 2016 The Gem Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

package ratelimit

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-gem/gem"
)

func TestRateLimit(t *testing.T) {
	max := int64(1)
	duration := 2 * time.Second
	limiter := New(max, duration)
	handler := limiter.Wrap(gem.HandlerFunc(func(ctx *gem.Context) {
		ctx.HTML(http.StatusOK, "success")
	}))

	req := httptest.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	ctx := &gem.Context{Request: req, Response: resp}

	handler.Handle(ctx)
	if resp.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.Code)
	}

	resp = httptest.NewRecorder()
	ctx.Response = resp
	handler.Handle(ctx)
	if resp.Code != limiter.StatusCode {
		t.Errorf("expected status code %d, got %d", limiter.StatusCode, resp.Code)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %q", err)
	}
	if string(body) != limiter.Message {
		t.Errorf("expected response body %q, got %q", limiter.Message, body)
	}

	time.Sleep(duration)

	resp = httptest.NewRecorder()
	ctx.Response = resp
	handler.Handle(ctx)
	if resp.Code != http.StatusOK {
		t.Errorf("expected status code %d, got %d", http.StatusOK, resp.Code)
	}
}
