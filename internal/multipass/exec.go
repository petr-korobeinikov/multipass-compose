package multipass

import (
	"bytes"
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

func ExecuteOutput(ctx context.Context, args ...string) (output []byte, err error) {
	var buf bytes.Buffer

	cmd := exec.CommandContext(ctx, "multipass", args...)

	cmd.Stdout = &buf
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
