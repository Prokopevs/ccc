package server

import (
	"context"
	"errors"
	"net/http"

	"github.com/Prokopevs/ccc/game/internal/model"
	"github.com/Prokopevs/ccc/game/internal/core"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Service interface {
	GetGame(context.Context, int) (*model.Game, error)
	UpdateScore(context.Context, *model.Score) (core.Code, error)
	UpdateMultiplicator(context.Context, *model.MultipUpdate) (core.Code, error)
}

type HTTP struct {
	innerServer *http.Server

	log     *zap.SugaredLogger
	service Service
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

func NewHTTP(addr string, logger *zap.SugaredLogger, service Service) *HTTP {
	h := &HTTP{
		log:     logger,
		service: service,
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