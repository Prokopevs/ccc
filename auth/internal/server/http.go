package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/Prokopevs/ccc/auth/internal/core"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//go:generate mockgen -source=http.go -destination=mocks/mock.go
type Service interface {
	GetUserInfo(context.Context, string, int) (*core.UserInfo, core.Code, error)
	GetUserReferrals(context.Context, int) ([]*core.UserReferrals, core.Code, error)
	GetUsers(context.Context) ([]*core.User, core.Code, error)
}

type HTTP struct {
	innerServer *http.Server

	log     *zap.SugaredLogger
	service Service
	password string
}

func (h *HTTP) Run(ctx context.Context) {
	h.log.Infow("HTTP server starting.", "addr", h.innerServer.Addr)

	go func() {
		err := h.innerServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}

			h.log.Errorw("Listen and serve HTTP", "addr", h.innerServer.Addr, "err", err)
		}
	}()

	<-ctx.Done()

	h.log.Info("Graceful server shutdown.")
	h.innerServer.Shutdown(context.Background())
}

func NewHTTP(addr string, logger *zap.SugaredLogger, service Service, password string) *HTTP {
	h := &HTTP{
		log:     logger,
		service: service,
		password: password,
	}

	r := gin.Default()

	h.setRoutes(r)

	srv := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	h.innerServer = srv

	return h
}
