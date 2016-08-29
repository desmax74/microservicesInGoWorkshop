// Copyright 2016 Massimiliano Dessi'.  All rights reserved.
// Use of this source code is governed by Apache
// license that can be found in the LICENSE file.
package main

import (
	"fmt"
	"log"
)

var conf config.Configuration

func init() {
	log.Println("Ready to rumble")
	fmt.Println(config.Banner)
	conf = config.GetConfiguration()
}

func main() {
	fmt.Println("The best is yet to come")
}
