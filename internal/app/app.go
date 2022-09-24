package app

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func New() *cli.App {
	return &cli.App{
		Commands: []*cli.Command{
			{
				Name: "up",
				Action: func(ctx *cli.Context) error {
					fmt.Println("up")
					return nil
				},
			},
			{
				Name: "down",
				Action: func(ctx *cli.Context) error {
					fmt.Println("down")
					return nil
				},
			},
		},
	}
}
