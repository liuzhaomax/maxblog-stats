package main

import (
	"context"
	"github.com/liuzhaomax/maxblog-stats/internal/app"
	"github.com/liuzhaomax/maxblog-stats/internal/core"
)

func main() {
	app.Launch(
		context.Background(),
		app.SetConfigFile(core.LoadEnv()),
		app.SetWWWDir("www"),
	)
}
