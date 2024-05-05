package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Prokopevs/ccc/game/internal/core"
	"github.com/Prokopevs/ccc/game/internal/pg"
	"github.com/Prokopevs/ccc/game/internal/server"
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

	httpServer := server.NewHTTP(cfg.httpAddr, sugaredLogger, service) //in

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(ctx context.Context) {
		httpServer.Run(ctx)
		wg.Done()
	}(ctx)

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-termChan
	cancel()

	return nil

}

func main() {
	err := run()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(exitCodeInitError)
	}
}