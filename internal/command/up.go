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
		args := []string{
			"launch",
			svc.Image,
			"--name",
			name,
		}

		if svc.CPUs != "" {
			args = append(args, "--cpus", svc.CPUs)
		}

		if svc.Mem != "" {
			args = append(args, "--mem", svc.Mem)
		}

		if svc.Disk != "" {
			args = append(args, "--disk", svc.Disk)
		}

		if err := multipass.Execute(ctx.Context, args...); err != nil {
			return err
		}
	}

	return nil
}
