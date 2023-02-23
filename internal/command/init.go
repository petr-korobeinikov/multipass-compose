package command

import (
	"errors"
	"os"

	"github.com/urfave/cli/v2"

	"github.com/pkorobeinikov/multipass-compose/internal/cfg"
	"github.com/pkorobeinikov/multipass-compose/internal/spec"
)

func Init(ctx *cli.Context) error {
	if _, err := os.Stat(cfg.DefaultMultipassComposeSpecFile); err == nil {
		return ErrSpecFileExists
	}

	s := &spec.Spec{
		Services: map[string]spec.Service{
			"instance-1": {
				Image: "focal",
			},
		},
	}

	b, err := spec.Encode(s)
	if err != nil {
		return err
	}

	if err := os.WriteFile(cfg.DefaultMultipassComposeSpecFile, b, os.ModePerm); err != nil {
		return err
	}

	return nil
}

var (
	ErrSpecFileExists = errors.New("spec file already exists")
)
