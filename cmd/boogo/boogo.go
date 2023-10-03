package main

import (
	"log"

	"github.com/realfabecker/boogo/internal/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
