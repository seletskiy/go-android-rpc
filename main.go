// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This is the Go entry point for the libhello app.
// It is invoked from Java.
//
// See README for details.
package main

//go:generate ./generate.bash

import (
	"golang.org/x/mobile/app"

	_ "github.com/seletskiy/java-go-rpc/rpc/go_rpc"
	_ "golang.org/x/mobile/bind/java"
)

func main() {
	app.Run(app.Callbacks{})
}
