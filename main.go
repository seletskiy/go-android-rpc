package main

//go:generate ./generate.bash

import (
	"golang.org/x/mobile/app"

	_ "github.com/seletskiy/go-android-rpc/android/go_android"
	_ "golang.org/x/mobile/bind/java"
)

func main() {
	app.Run(app.Callbacks{})
}
