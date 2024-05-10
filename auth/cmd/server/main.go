package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Prokopevs/ccc/auth/internal/core"
	"github.com/Prokopevs/ccc/auth/internal/server"
	"github.com/Prokopevs/ccc/schema"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	exitCodeInitError = 2
)

func run() error {
	cfg, err := loadEnvConfig()
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())

	conn, err := grpc.DialContext(ctx, cfg.usersGRPCAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		cancel()
		return err
	}

	client := schema.NewUsersClient(conn)

	service := core.NewServiceImpl(client, cfg.telegramToken)

	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugaredLogger := logger.Sugar()

	httpServer := server.NewHTTP(cfg.httpAddr, sugaredLogger, service)

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

//  @title Game API
//  @version 1.0
//	@description This is Auth server.
// @host localhost:4000
// @BasePath /api/v1/auth
func main() {
	err := run()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(exitCodeInitError)
	}
}