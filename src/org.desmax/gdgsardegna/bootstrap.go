// Copyright 2016 Massimiliano Dessi'.  All rights reserved.
// Use of this source code is governed by Apache
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"org.desmax/gdgsardegna/config"
	"org.desmax/gdgsardegna/handlers"
	mid "org.desmax/gdgsardegna/middlewares"
	"org.desmax/gdgsardegna/sys"
)

var Conf config.Configuration

func init() {
	fmt.Println("Ready to rumble")
	Conf = config.GetConfiguration()
	fmt.Println(config.Banner)
}

func main() {

	//Log to server (on linux nc -lk 1902)
	// logServer := sys.StartLogServer(Conf)
	//logger := sys.StartLogger(Conf, logServer)
	//defer logServer.Close()

	//Log to file
	//logger, logFile := sys.StartLoggerFS(Conf)
	//defer logFile.Close()

	//LOG to stdout
	logger := sys.StartLoggerStdOut(Conf, "INFO:")

	db := sys.StartDB(Conf, logger)// NOTE: embedded only for development, when deployed as a container use an external database
	defer db.Close()

	//goroutine to print stats
	//go sys.MonitorRuntime(logger, Conf)

	/*Routes*/
	router := fasthttprouter.New()
	router.GET("/", mid.DBMidw(mid.LogMidw(mid.ConfigMidw(handlers.List, Conf), logger), db))  //Middleware chain, decorator pattern
	router.GET("/:id", mid.DBMidw(mid.LogMidw(mid.ConfigMidw(handlers.Read, Conf), logger), db))
	router.POST("/", mid.DBMidw(mid.LogMidw(mid.ConfigMidw(handlers.SaveOrUpdateMsg, Conf), logger), db))
	router.PUT("/", mid.DBMidw(mid.LogMidw(mid.ConfigMidw(handlers.SaveOrUpdateMsg, Conf), logger), db))

	/*Server*/
	config.HttpServer(router, Conf) //HTTP
	//config.HttpsServer(router, conf) //HTTPS


}
