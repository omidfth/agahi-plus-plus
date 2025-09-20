package main

import (
	"agahi-plus-plus/handler/app"
	"agahi-plus-plus/internal/constant"
	"context"
	"flag"
)

func main() {
	path := flag.String("e", constant.DefaultEnvPath, "env file path")
	flag.Parse()
	config, err := app.SetupViper(*path)
	if err != nil {
		panic(err)
	}

	application := app.NewApplication(config)
	ctx := context.Background()

	err = application.Setup(ctx)
	if err != nil {
		panic(err)
	}
}
