package app

import (
	"errors"

	"github.com/urfave/cli/v2"

	"github.com/petr-korobeinikov/multipass-compose/internal/command"
)

func New() *cli.App {
	return &cli.App{
		Commands: []*cli.Command{
			{
				Name: "up",
				Action: func(ctx *cli.Context) error {
					return command.Up(ctx.Context)
				},
			},
			{
				Name: "down",
				Action: func(ctx *cli.Context) error {
					return command.Down(ctx.Context)
				},
			},
			{
				Name: "status",
				Action: func(ctx *cli.Context) error {
					return command.Status(ctx.Context)
				},
			},
			{
				Name: "ip",
				Action: func(ctx *cli.Context) error {
					if ctx.NArg() == 0 {
						return ErrMachineNameNotSpecified
					}

					return command.Ip(ctx.Context, ctx.Args().First())
				},
			},
			{
				Name: "init",
				Action: func(ctx *cli.Context) error {
					return command.Init(ctx.Context)
				},
			},
			{
				Name: "restart",
				Action: func(ctx *cli.Context) error {
					if err := command.Down(ctx.Context); err != nil {
						return err
					}

					if err := command.Up(ctx.Context); err != nil {
						return err
					}

					return nil
				},
			},
		},
	}
}

var (
	ErrMachineNameNotSpecified = errors.New("machine name not specified")
)
