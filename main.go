package main

import (
	"log"

	"github.com/praveenmahasena/aiclient/internal"
)

func main() {
	if err := internal.Run(); err != nil {
		log.Fatalln(err)
	}
}
