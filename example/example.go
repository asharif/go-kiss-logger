package main

import (
	logger "github.com/asharif/go-kiss-logger"
)

func main() {
	logger := logger.GetInstance()
	logger.Info("foo")
	logger.Error("foo")
	logger.Fatal("foo", 1)

}
