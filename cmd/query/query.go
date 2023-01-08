package main

import (
	"github.com/labstack/echo/v4"
	"github.com/palavrapasse/query/internal"
	"github.com/palavrapasse/query/internal/data"
	"github.com/palavrapasse/query/internal/http"
)

func main() {

	internal.Aspirador.Trace("Starting Query Service")
	e := echo.New()

	defer e.Close()

	dbctx, oerr := data.Open()

	if oerr != nil {

		internal.Aspirador.Warning("Could not open DB connection on server start")
		internal.Aspirador.Error(oerr.Error())
		internal.Aspirador.Error("DB connection is required to operate query")

		return
	}

	http.RegisterMiddlewares(e, dbctx)
	http.RegisterHandlers(e)

	e.Logger.Fatal(http.Start(e))
}
