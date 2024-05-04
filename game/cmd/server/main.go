package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Prokopevs/ccc/game/internal/core"
	"github.com/Prokopevs/ccc/game/internal/pg"
	"go.uber.org/zap"
)

const (
	exitCodeInitError = 2
)

func run() error {
	cfg, err := loadEnvConfig()
	if err != nil {
		return err
	}

	d, err := pg.Connect(context.Background(), cfg.pgConnString) // in
	if err != nil {
		return err
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugaredLogger := logger.Sugar()

	ctx, cancel := context.WithCancel(context.Background())

	service := core.NewService(d) // in
}

func main() {
	err := run()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(exitCodeInitError)
	}
}