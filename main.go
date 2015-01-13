package main

import (
	"golang.org/x/mobile/app"

	_ "github.com/seletskiy/go-android-rpc/rpc/go_rpc"
	_ "golang.org/x/mobile/bind/java"
)

func main() {
	app.Run(app.Callbacks{})
}
