// Copyright 2016 Massimiliano Dessi'.  All rights reserved.
// Use of this source code is governed by Apache
// license that can be found in the LICENSE file.
package middlewares

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"org.desmax/gdgsardegna/config"
)

func ConfigMidw(h fasthttprouter.Handle, conf config.Configuration) fasthttprouter.Handle {
	return fasthttprouter.Handle(func(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
		SetConfig(ctx, conf)
		h(ctx, ps)
	})
}

func GetConfig(ctx *fasthttp.RequestCtx) (config.Configuration, error) {
	configuration := ctx.UserValue(config.CONFIG)
	if configuration != nil {
		return configuration.(config.Configuration), nil //type assertion
	} else {
		return config.Configuration{}, fmt.Errorf("Config not found")
	}
}

func SetConfig(ctx *fasthttp.RequestCtx, configuration config.Configuration) {
	ctx.SetUserValue(config.CONFIG, configuration)
}
