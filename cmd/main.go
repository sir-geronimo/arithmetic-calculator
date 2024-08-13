package main

import (
	"log"

	"github.com/sir-geronimo/arithmetic-calculator/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatalf("%v", err)
	}
}
