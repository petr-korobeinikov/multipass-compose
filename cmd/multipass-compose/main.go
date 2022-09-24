package main

import (
	"log"
	"os"

	"github.com/pkorobeinikov/multipass-compose/internal/app"
)

func main() {
	if err := app.New().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
