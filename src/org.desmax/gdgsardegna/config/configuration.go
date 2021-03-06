// Copyright 2016 Massimiliano Dessi'.  All rights reserved.
// Use of this source code is governed by Apache
// license that can be found in the LICENSE file.
package config

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"strconv"
)

type Configuration struct {
	AppName     string
	ServerPort  string
	Address     string
	Development bool
	Certfile    string
	KeyFile     string
	LogPath     string
	LogServer   string
}

func GetConfiguration() Configuration {

	development := true
	serverPort := os.Getenv(PORT)
	if(serverPort == "") {serverPort = "8080"}

	appName := os.Getenv(APPNAME)
	if(appName == "") {appName = "GDG"}

	dev := os.Getenv(DEVELOPMENT)
	if(dev == "") {dev = "true"}

	certFile := os.Getenv(CERTFILE)
	keyFile := os.Getenv(KEYFILE)
	logPath := os.Getenv(LOGPATH)
	logServer := os.Getenv(LOGSERVER)
	development, _ = strconv.ParseBool(dev)

	conf := Configuration{
		Development: development,
		ServerPort:  serverPort,
		AppName:     appName,
		Certfile:    certFile,
		KeyFile:     keyFile,
		LogPath:     logPath,
		LogServer:   logServer,
	}

	//if conf.Development {
		fmt.Println("Conf:", conf)
	//}
	return conf
}

func HttpServer(router *fasthttprouter.Router, conf Configuration) {
	log.Fatal(fasthttp.ListenAndServe(":"+conf.ServerPort, router.Handler))
}

func HttpsServer(router *fasthttprouter.Router, conf Configuration) {
	log.Fatal(fasthttp.ListenAndServeTLS(conf.ServerPort, conf.Certfile, conf.KeyFile, router.Handler))
}
