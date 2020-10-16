package main

import (
	"github.com/objque/go-app-template/internal/log"
)

func main() {
	log.SetLevel("INFO")
	log.SetWriters(log.GetConsoleWriter())

	log.Info("Hello, world")
}
