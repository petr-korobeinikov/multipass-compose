package command

import (
	"context"
	"log"

	"github.com/pkorobeinikov/multipass-compose/internal/cfg"
	"github.com/pkorobeinikov/multipass-compose/internal/multipass"
	"github.com/pkorobeinikov/multipass-compose/internal/spec"
)

func Status(ctx context.Context) error {
	s, err := spec.Load(cfg.DefaultMultipassComposeSpecFile)
	if err != nil {
		return err
	}

	for name := range s.Services {
		if err := multipass.Execute(ctx, "info", name); err != nil {
			log.Println(err)
		}
	}

	return nil
}
