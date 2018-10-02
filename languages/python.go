package languages

import (
	"context"
)

type PythonRunner struct {
}

const runCmd = "python %s"

func (r *PythonRunner) Run(ctx context.Context) (context.Context, string, string) {
	return ctx, "python:latest", "python main.py"
}
