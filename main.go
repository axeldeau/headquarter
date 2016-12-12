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
	a := app.NewApp()
	// app.SetLogger(logger)
	// app.SetConfig(conf)
	a.Register("/", app.HelloWorld, "get")
	a.Run()
}
