package app

import (
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
		},
	}
}
