package main

import (
	"fmt"

	"./app"
	"./logger"
)

func init() {
	fmt.Println("Initializing main")
}
func main() {
	logger := logger.NewLogger()
	logger.SetLogLevel("Info")
	app := app.NewApp()
	app.SetLogger(logger)
	app.Run()
}
