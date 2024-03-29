package command

import (
	"context"

	"github.com/petr-korobeinikov/multipass-compose/internal/cfg"
	"github.com/petr-korobeinikov/multipass-compose/internal/multipass"
	"github.com/petr-korobeinikov/multipass-compose/internal/spec"
)

func Up(ctx context.Context) error {
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

		if svc.CloudInit != "" {
			args = append(args, "--cloud-init", svc.CloudInit)
		}

		if err := multipass.Execute(ctx, args...); err != nil {
			return err
		}
	}

	return nil
}
