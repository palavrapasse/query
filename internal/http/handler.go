package http

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/palavrapasse/query/internal/data"
)

func RegisterHandlers(e *echo.Echo) {

	e.GET(leaksRoute, QueryLeaks)

	echo.NotFoundHandler = useNotFoundHandler()
}

func QueryLeaks(ectx echo.Context) error {

	mwctx, gmerr := GetMiddlewareContext(ectx)

	if gmerr != nil {
		return InternalServerError(ectx)
	}

	affp := ectx.QueryParam(affectedQueryParam)
	aff := ParseAffected(affp)
	hus := AffectedToHashUser(aff)

	ls, err := data.QueryLeaksDB(mwctx.DB, hus)

	if err != nil {
		log.Printf("wtf happened: %v\n", err)

		return InternalServerError(ectx)
	}

	return Ok(ectx, ToQueryLeaksView(ls))
}

func useNotFoundHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusNotFound)
	}
}
