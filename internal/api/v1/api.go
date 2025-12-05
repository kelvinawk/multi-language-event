package api

import (
	"multi-language-event/config"
	"multi-language-event/internal/service"

	"github.com/labstack/echo"
)

type Api struct {
	config  *config.Config
	service *service.APIService
}

func NewApi(service *service.APIService, cfg *config.Config) *Api {
	return &Api{
		config:  cfg,
		service: service,
	}
}

func (api *Api) Register(e *echo.Echo) {
	v1 := e.Group("/api/v1")
	v1.GET("/healthz", api.healthCheckHandler)
}
