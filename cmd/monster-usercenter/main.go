package main

import (
	_ "go.uber.org/automaxprocs"

	"github.com/rosas99/monster/cmd/monster-usercenter/app"
)

func main() {
	app.NewApp("monster-usercenter").Run()
}
