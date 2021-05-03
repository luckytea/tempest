package main

import (
	"log"
	"os"

	"github.com/LuckyTea/tempest/cfg"
)

var version = "dev"

func main() {
	_, err := cfg.Init().Validate()
	if err != nil {
		log.Println("can't initiate application cause:", err)
		os.Exit(1)
	}
}
