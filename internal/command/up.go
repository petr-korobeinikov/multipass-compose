package command

import (
	"github.com/urfave/cli/v2"

	"github.com/pkorobeinikov/multipass-compose/internal/cfg"
	"github.com/pkorobeinikov/multipass-compose/internal/multipass"
	"github.com/pkorobeinikov/multipass-compose/internal/spec"
)

func Up(ctx *cli.Context) error {
	s, err := spec.Load(cfg.DefaultMultipassComposeSpecFile)
	if err != nil {
		return err
	}

	for name, svc := range s.Services {
		if err := multipass.Execute(ctx.Context, "launch", svc.Image, "--name", name); err != nil {
			return err
		}
	}

	return nil
}
