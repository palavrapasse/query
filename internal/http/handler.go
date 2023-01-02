package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/palavrapasse/damn/pkg/entity"
)

func RegisterHandlers(e *echo.Echo) {

	e.GET(leaksRoute, QueryLeaks)

	echo.NotFoundHandler = useNotFoundHandler()
}

func QueryLeaks(ectx echo.Context) error {

	_, gmerr := GetMiddlewareContext(ectx)

	if gmerr != nil {
		return InternalServerError(ectx)
	}

	return Ok(ectx, ToQueryLeaksView([]entity.Leak{}))
}

func useNotFoundHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusNotFound)
	}
}
