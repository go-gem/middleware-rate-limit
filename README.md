# Rate Limiting Middleware

[![Build Status](https://travis-ci.org/go-gem/middleware-rate-limit.svg?branch=master)](https://travis-ci.org/go-gem/middleware-rate-limit)
[![GoDoc](https://godoc.org/github.com/go-gem/middleware-rate-limit?status.svg)](https://godoc.org/github.com/go-gem/middleware-rate-limit)
[![Coverage Status](https://coveralls.io/repos/github/go-gem/middleware-rate-limit/badge.svg?branch=master)](https://coveralls.io/github/go-gem/middleware-rate-limit?branch=master)

Rate limiting middleware for [Gem](https://github.com/go-gem/gem) Web framework.

## Getting Started

**Install**

```
$ go get -u github.com/go-gem/middleware-rate-limit
```

**Example**

```
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
			ratelimit.New(1, time.Second),
		},
	})

	log.Println(gem.ListenAndServe(":8080", router.Handler()))
}
```