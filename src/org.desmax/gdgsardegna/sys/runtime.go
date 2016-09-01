// Copyright 2016 Massimiliano Dessi'.  All rights reserved.
// Use of this source code is governed by Apache
// license that can be found in the LICENSE file.
package sys

import (
	"log"
	"org.desmax/gdgsardegna/config"
	"runtime"
	"time"
)

func MonitorRuntime(logger log.Logger, conf config.Configuration) {
	logger.Println("Number of CPUs:", runtime.NumCPU())
	if conf.Development {
		m := &runtime.MemStats{}
		for {
			logger.Println("ARCH", runtime.GOARCH)
			logger.Println("GOOS", runtime.GOOS)
			logger.Println("Number of goroutines", runtime.NumGoroutine())
			runtime.ReadMemStats(m) //this had a performance impact
			logger.Println("Allocated memory", m.Alloc)
			time.Sleep(10 * time.Second)
		}
	}
}
