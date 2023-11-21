package command

import (
	"context"
	"fmt"

	"github.com/petr-korobeinikov/multipass-compose/internal/cfg"
	"github.com/petr-korobeinikov/multipass-compose/internal/multipass"
	"github.com/petr-korobeinikov/multipass-compose/internal/spec"
	"github.com/petr-korobeinikov/multipass-compose/internal/state"
)

func Ip(ctx context.Context, machineName string) error {
	s, err := spec.Load(cfg.DefaultMultipassComposeSpecFile)
	if err != nil {
		return err
	}

	for name := range s.Services {
		if name == machineName {
			b, err := multipass.ExecuteOutput(ctx, "info", name, "--format", "json")
			if err != nil {
				return err
			}

			info, err := state.ParseInfo(b)
			if err != nil {
				return err
			}

			fmt.Println(info.Info[name].Ipv4[0])
		}
	}

	return nil
}
