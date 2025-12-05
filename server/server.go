package server

import (
	"fmt"
	"multi-language-event/config"
	"multi-language-event/internal/api/v1"
	"multi-language-event/internal/service"
	store "multi-language-event/internal/storage"

	"github.com/labstack/echo"
)

type Server struct {
	router *echo.Echo
	config *config.Config
}

func NewHTTPServer(store *store.Store, cfg *config.Config) *Server {
	e := echo.New()

	service := service.NewAPIService(store, cfg)
	api := api.NewApi(service, cfg)
	api.Register(e)

	s := &Server{
		router: e,
		config: cfg,
	}

	return s
}

func (s Server) Start() error {
	return s.router.Start(fmt.Sprint(":", s.config.Port))
}
