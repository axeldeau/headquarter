package main

import (
	"fmt"

	"./app"
)

func init() {
	fmt.Println("Initializing Main.")
}

func main() {
	// logger := logger.NewLogger()
	// logger.SetLogLevel("Info")
	app := app.NewApp()
	// app.SetLogger(logger)
	// app.SetConfig(conf)
	app.Run()
}
