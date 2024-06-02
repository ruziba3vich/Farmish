// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"Farmish/config"
	"Farmish/internal/controller/middleware"
	tokens "Farmish/pkg/token"
	"net/http"

	"github.com/casbin/casbin/v2"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "Farmish/docs"
	"Farmish/internal/usecase"
	"Farmish/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Golang dealer APP
// @description Documentation
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l logger.Interface, t usecase.Admin, e *casbin.Enforcer, cfg *config.Config) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	jwtHandler := tokens.JWTHandler{
		SigninKey: cfg.Casbin.SigningKey,
	}
	handler.Use(middleware.NewAuthorizer(e, jwtHandler, cfg, l))

	// Swagger
	swaggerHandler := ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, "DISABLE_SWAGGER_HTTP_HANDLER")
	handler.GET("/swagger/*any", swaggerHandler)

	// K8s probe
	handler.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	// Prometheus metrics
	handler.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Routers
	h := handler.Group("/v1")
	{
		newAdminRoutes(h, t, l)
	}
}
