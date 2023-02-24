package command

import (
	"context"
	"errors"
	"os"

	"github.com/pkorobeinikov/multipass-compose/internal/cfg"
	"github.com/pkorobeinikov/multipass-compose/internal/fsext"
	"github.com/pkorobeinikov/multipass-compose/internal/spec"
)

func Init(_ context.Context) error {
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

	if err := os.WriteFile(cfg.DefaultMultipassComposeSpecFile, b, fsext.ModeFile); err != nil {
		return err
	}

	return nil
}

var (
	ErrSpecFileExists = errors.New("spec file already exists")
)
