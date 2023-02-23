package command

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/pkorobeinikov/multipass-compose/internal/cfg"
	"github.com/pkorobeinikov/multipass-compose/internal/multipass"
	"github.com/pkorobeinikov/multipass-compose/internal/spec"
	"github.com/pkorobeinikov/multipass-compose/internal/state"
)

func Ip(ctx *cli.Context, machineName string) error {
	s, err := spec.Load(cfg.DefaultMultipassComposeSpecFile)
	if err != nil {
		return err
	}

	for name := range s.Services {
		if name == machineName {
			b, err := multipass.ExecuteOutput(ctx.Context, "info", name, "--format", "json")
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
