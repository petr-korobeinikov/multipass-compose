package command

import (
	"log"

	"github.com/urfave/cli/v2"

	"github.com/pkorobeinikov/multipass-compose/internal/cfg"
	"github.com/pkorobeinikov/multipass-compose/internal/multipass"
	"github.com/pkorobeinikov/multipass-compose/internal/spec"
)

func Down(ctx *cli.Context) error {
	s, err := spec.Load(cfg.DefaultMultipassComposeSpecFile)
	if err != nil {
		return err
	}

	for name := range s.Services {
		if err := multipass.Execute(ctx.Context, "delete", "--purge", name); err != nil {
			log.Println(err)
		}
	}

	return nil
}
