package main

import (
	_ "go.uber.org/automaxprocs"

	"github.com/rosas99/monster/cmd/monster-sms/app"
)

func main() {
	app.NewApp("monster-sms").Run()
}
