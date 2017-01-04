// Copyright 2016 The Gem Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license
// that can be found in the LICENSE file.

/*
Package ratelimit is a HTTP middleware that limit API usage of each user.

Example

	package main

	import (
		"time"
		"log"

		"github.com/go-gem/gem"
		"github.com/go-gem/middleware-rate-limit"
	)

	func main() {
		router := gem.NewRouter()

		router.GET("/", func(ctx *gem.Context) {
			ctx.HTML(200, "hello")
		}, &gem.HandlerOption{
			Middlewares:[]gem.Middleware{
				ratelimit.New(60, time.Minute),
			},
		})

		log.Println(gem.ListenAndServe(":8080", router.Handler()))
	}
*/
package ratelimit

import (
	"time"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/config"
	"github.com/go-gem/gem"
)

// New returns a rate limiter.
//
// maximum number of requests during the given duration.
func New(max int64, ttl time.Duration) *Limiter {
	return &Limiter{tollbooth.NewLimiter(max, ttl)}
}

// Limiter is a HTTP middleware that limit
// API usage of each user.
type Limiter struct {
	*config.Limiter
}

// Wrap implements the Middleware interface.
func (l *Limiter) Wrap(next gem.Handler) gem.Handler {
	return gem.HandlerFunc(func(ctx *gem.Context) {
		err := tollbooth.LimitByRequest(l.Limiter, ctx.Request)
		if err != nil {
			ctx.HTML(err.StatusCode, err.Message)
			return
		}

		next.Handle(ctx)
	})
}
