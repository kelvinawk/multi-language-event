package api

import (
	"net/http"

	"github.com/labstack/echo"
)

func (api *Api) healthCheckHandler(e echo.Context) error {
	return e.String(http.StatusOK, "OK")
}
