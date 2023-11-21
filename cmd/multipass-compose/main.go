package main

import (
	"log"
	"os"

	"github.com/petr-korobeinikov/multipass-compose/internal/app"
)

func main() {
	if err := app.New().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
