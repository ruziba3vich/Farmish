// Package app configures and runs application.
package app

import (
	"Farmish/pkg/redis"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/casbin/casbin/v2"

	"github.com/gin-gonic/gin"

	"Farmish/config"
	v1 "Farmish/internal/controller/http/v1"
	"Farmish/internal/usecase"
	"Farmish/internal/usecase/repo"
	"Farmish/pkg/httpserver"
	"Farmish/pkg/logger"
	"Farmish/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Redis
	RedisClient, err := redis.NewRedisDB(cfg)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - redis.New: %w", err))
	}

	//Casbin RBAC
	casbinEnforcer, err := casbin.NewEnforcer(cfg.Casbin.ConfigFilePath, cfg.Casbin.CSVFilePath)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - casbin.NewEnforcer: %w", err))
	}

	// Use case
	adminUseCase := usecase.NewAdminUseCase(
		repo.NewAdminRepo(pg), cfg, RedisClient,
		//webapi.New(),
	)

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, adminUseCase, casbinEnforcer, cfg)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}

}
