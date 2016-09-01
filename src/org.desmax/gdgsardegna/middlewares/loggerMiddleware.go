// Copyright 2016 Massimiliano Dessi'.  All rights reserved.
// Use of this source code is governed by Apache
// license that can be found in the LICENSE file.
package middlewares

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
	"org.desmax/gdgsardegna/config"
)

func LogMidw(h fasthttprouter.Handle, logger log.Logger) fasthttprouter.Handle {
	return fasthttprouter.Handle(func(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
		SetLog(ctx, logger)
		h(ctx, ps)
	})
}

func GetLog(ctx *fasthttp.RequestCtx) (log.Logger, error) {
	logger := ctx.UserValue(config.LOGGER)
	if logger != nil {
		return logger.(log.Logger), nil //type assertion
	} else {
		return log.Logger{}, fmt.Errorf("Logger not found")
	}
}

func SetLog(ctx *fasthttp.RequestCtx, logger log.Logger) {
	ctx.SetUserValue(config.LOGGER, logger)
}
