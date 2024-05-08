package main

import (
	"fmt"
	"os"
)

type envConfig struct {
	httpAddr      string
	usersGRPCAddr string
	telegramToken string
	key string
}

func loadEnvConfig() (*envConfig, error) {
	var err error

	cfg := &envConfig{}

	cfg.httpAddr, err = lookupEnv("HTTP_SERVER_ADDR")
	if err != nil {
		return nil, err
	}

	cfg.usersGRPCAddr, err = lookupEnv("USERS_GRPC_ADDR")
	if err != nil {
		return nil, err
	}

	cfg.telegramToken, err = lookupEnv("TELEGRAM_TOKEN")
	if err != nil {
		return nil, err
	}

	cfg.key, err = lookupEnv("KEY")
	if err != nil {
		return nil, err
	}
	fmt.Println(cfg.key)

	return cfg, nil
}

func lookupEnv(name string) (string, error) {
	const provideEnvErrorMsg = `please provide "%s" environment variable`

	val, ok := os.LookupEnv(name)
	if !ok {
		return "", fmt.Errorf(provideEnvErrorMsg, name)
	}

	return val, nil
}
