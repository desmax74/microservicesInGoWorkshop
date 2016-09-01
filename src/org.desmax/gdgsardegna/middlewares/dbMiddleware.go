// Copyright 2016 Massimiliano Dessi'.  All rights reserved.
// Use of this source code is governed by Apache
// license that can be found in the LICENSE file.
package middlewares

import (
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"org.desmax/gdgsardegna/config"
)

func DBMidw(h fasthttprouter.Handle, db bolt.DB) fasthttprouter.Handle {
	return fasthttprouter.Handle(func(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
		SetDB(ctx, db)
		h(ctx, ps)
	})
}

func GetDB(ctx *fasthttp.RequestCtx) (bolt.DB, error) {
	db := ctx.UserValue(config.DB)
	if db != nil {
		return db.(bolt.DB), nil //type assertion
	} else {
		return bolt.DB{}, fmt.Errorf("Config not found")
	}
}

func SetDB(ctx *fasthttp.RequestCtx, db bolt.DB) {
	ctx.SetUserValue(config.DB, db)
}
