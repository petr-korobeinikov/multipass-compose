package app

import (
	"errors"

	"github.com/urfave/cli/v2"

	"github.com/pkorobeinikov/multipass-compose/internal/command"
)

func New() *cli.App {
	return &cli.App{
		Commands: []*cli.Command{
			{
				Name: "up",
				Action: func(ctx *cli.Context) error {
					return command.Up(ctx)
				},
			},
			{
				Name: "down",
				Action: func(ctx *cli.Context) error {
					return command.Down(ctx)
				},
			},
			{
				Name: "status",
				Action: func(ctx *cli.Context) error {
					return command.Status(ctx)
				},
			},
			{
				Name: "ip",
				Action: func(ctx *cli.Context) error {
					if ctx.NArg() == 0 {
						return ErrMachineNameNotSpecified
					}

					return command.Ip(ctx, ctx.Args().First())
				},
			},
			{
				Name: "init",
				Action: func(ctx *cli.Context) error {
					return command.Init(ctx)
				},
			},
		},
	}
}

var (
	ErrMachineNameNotSpecified = errors.New("machine name not specified")
)
