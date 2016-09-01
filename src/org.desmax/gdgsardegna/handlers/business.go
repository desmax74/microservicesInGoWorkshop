package handlers

import (
	"errors"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	m "org.desmax/gdgsardegna/middlewares"
	"strconv"
	"org.desmax/gdgsardegna/repos"
	"org.desmax/gdgsardegna/domain"
	"encoding/json"
)

var ErrTimeout = errors.New("The request timed out")
var ErrRejected = errors.New("The request was rejected")

func Index(ctx *fasthttp.RequestCtx, params fasthttprouter.Params) {

	conf, _ := m.GetConfig(ctx)
	logger, _ := m.GetLog(ctx)
	logger.Println("Welcome to {}", conf.AppName)
}

func Read(ctx *fasthttp.RequestCtx, params fasthttprouter.Params) {
	logger, _ := m.GetLog(ctx)
	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		logger.Print(err)
	}
	logger.Print("ID received ", id)
	db, err := m.GetDB(ctx)
	if err != nil {
		logger.Fatal("Can't access the DB")
		ctx.SetStatusCode(500)
	}
	v := repos.ReadMsg(db, strconv.Itoa(id), logger)
	logger.Println(v)
}

func List(ctx *fasthttp.RequestCtx, params fasthttprouter.Params) {
	conf, _ := m.GetConfig(ctx)
	logger, _ := m.GetLog(ctx)
	fmt.Fprint(ctx, "Welcome to ", conf.AppName, " ")
	db, err:= m.GetDB(ctx)
	if err != nil {
		logger.Fatal("Can't access the DB")
		ctx.SetStatusCode(500)
	}
	msgs := repos.ReadAll(db,repos.GdgBucket, logger)
	//write on the outputsream in to ways:
	fmt.Fprint(ctx, msgs) //short way
	//ctx.Response.BodyWriter().Write([]byte(fmt.Sprintf("%v", msgs))) long way
}

func SaveOrUpdateMsg(ctx *fasthttp.RequestCtx, params fasthttprouter.Params) {
	logger, _ := m.GetLog(ctx)
	db, err := m.GetDB(ctx)
	if err != nil {
		logger.Fatal("Can't access the DB")
		ctx.SetStatusCode(500)
	}
	body := string(ctx.Request.Body())
	logger.Printf("String received %s", body)
	var msg domain.Message
	errM := json.Unmarshal( ctx.Request.Body(),&msg)
	if errM != nil {
		logger.Printf("Error %s", errM)
		ctx.SetStatusCode(500)
	}
	b := repos.SaveOrUpdateMsg(db, msg, logger)
	//logger.Printf("key %s", b)
	ctx.Response.BodyWriter().Write(b)
}


