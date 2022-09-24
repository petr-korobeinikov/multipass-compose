package multipass

import (
	"context"
	"os"
	"os/exec"
)

func Execute(ctx context.Context, args ...string) error {
	cmd := exec.CommandContext(ctx, "multipass", args...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
