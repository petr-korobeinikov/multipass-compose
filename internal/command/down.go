package command

import (
	"context"
	"log"

	"github.com/petr-korobeinikov/multipass-compose/internal/cfg"
	"github.com/petr-korobeinikov/multipass-compose/internal/multipass"
	"github.com/petr-korobeinikov/multipass-compose/internal/spec"
)

func Down(ctx context.Context) error {
	s, err := spec.Load(cfg.DefaultMultipassComposeSpecFile)
	if err != nil {
		return err
	}

	for name := range s.Services {
		if err := multipass.Execute(ctx, "delete", "--purge", name); err != nil {
			log.Println(err)
		}
	}

	return nil
}
