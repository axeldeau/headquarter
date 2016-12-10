package main

import (
	"fmt"

	"./errors"
	"./logger"
)

func main() {
	fmt.Println("Hello World !")
	l := logger.NewLogger()
	l.SetLevel(logger.Warn)
	l.Trace(errors.Err{"test", "message"})
	l.Debug(errors.Err{"test", "message"})
	l.Info(errors.Err{"test", "message"})
	l.Warn(errors.Err{"test", "message"})
	l.Error(errors.Err{"test", "message"})
	l.Fatal(errors.Err{"test", "message"})
}
