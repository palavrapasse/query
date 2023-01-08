package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/palavrapasse/query/internal/data"
	"github.com/palavrapasse/query/internal/logging"
)

func RegisterHandlers(e *echo.Echo) {

	e.GET(leaksRoute, QueryLeaks)

	echo.NotFoundHandler = useNotFoundHandler()
}

func QueryLeaks(ectx echo.Context) error {

	logging.Aspirador.Trace("Querying leaks")

	mwctx, gmerr := GetMiddlewareContext(ectx)

	if gmerr != nil {
		return InternalServerError(ectx)
	}

	affp := ectx.QueryParam(affectedQueryParam)
	aff := ParseAffected(affp)
	hus := AffectedToHashUser(aff)

	ls, err := data.QueryLeaksDB(mwctx.DB, hus)

	if err != nil {
		logging.Aspirador.Error(fmt.Sprintf("Error while querying Leaks from DB: %s", err))

		return InternalServerError(ectx)
	}

	logging.Aspirador.Trace(fmt.Sprintf("Success in querying leaks. Found %d leaks", len(ls)))
	return Ok(ectx, ToQueryLeaksView(ls))
}

func useNotFoundHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusNotFound)
	}
}
