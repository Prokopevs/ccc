package main

import (
	"fmt"
	"os"
)

type envConfig struct {
	pgConnString    string
	httpAddr        string
	usersGRPCAddr   string
	key             string
	iv              string
	redisConnString string
}

func loadEnvConfig() (*envConfig, error) {
	const (
		provideEnvErrorMsg = `please provide "%s" environment variable`

		pgConnStringEnv    = "PG_CONN"
		httpAddrEnv        = "HTTP_ADDR"
		usersGRPCAddrEnv   = "USERS_GRPC_ADDR"
		keyEnv             = "KEY"
		ivEnv              = "IV"
		redisConnStringEnv = "RD_CONN"
	)

	var ok bool

	cfg := &envConfig{}

	cfg.pgConnString, ok = os.LookupEnv(pgConnStringEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, pgConnStringEnv)
	}

	cfg.httpAddr, ok = os.LookupEnv(httpAddrEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, httpAddrEnv)
	}

	cfg.usersGRPCAddr, ok = os.LookupEnv(usersGRPCAddrEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, usersGRPCAddrEnv)
	}

	cfg.key, ok = os.LookupEnv(keyEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, keyEnv)
	}

	cfg.iv, ok = os.LookupEnv(ivEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, ivEnv)
	}

	cfg.redisConnString, ok = os.LookupEnv(redisConnStringEnv)
	if !ok {
		return nil, fmt.Errorf(provideEnvErrorMsg, redisConnStringEnv)
	}

	return cfg, nil
}
