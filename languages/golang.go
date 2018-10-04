package languages

import (
	"context"
	"fmt"
	"time"
)

const maxRunTime = time.Second * 10
const buildCmd = "go build -o %s %s && ./%s"
const image = "golang:1.11.1-alpine3.7"

type GolangRunner struct {
}

func (r *GolangRunner) Run(ctx context.Context) (context.Context, string, string) {
	ctx, _ = context.WithTimeout(context.Background(), maxRunTime)
	return ctx, image, fmt.Sprintf(buildCmd, "main", ".", "main")
}
