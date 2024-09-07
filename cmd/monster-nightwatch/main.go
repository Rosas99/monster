package main

import (
	_ "go.uber.org/automaxprocs"

	"github.com/rosas99/monster/cmd/monster-nightwatch/app"
)

func main() {
	app.NewApp("monster-nightwatch").Run()
}
