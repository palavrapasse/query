package main

import (
	"github.com/labstack/echo/v4"
	as "github.com/palavrapasse/aspirador/pkg"
	"github.com/palavrapasse/query/internal/data"
	"github.com/palavrapasse/query/internal/http"
	"github.com/palavrapasse/query/internal/logging"
)

func main() {

	logging.Aspirador = as.WithClients(logging.CreateAspiradorClients())

	logging.Aspirador.Trace("Starting Query Service")

	e := echo.New()

	defer e.Close()

	dbctx, oerr := data.Open()

	if oerr != nil {

		logging.Aspirador.Warning("Could not open DB connection on server start")
		logging.Aspirador.Error(oerr.Error())
		logging.Aspirador.Error("DB connection is required to operate query")

		return
	}

	http.RegisterMiddlewares(e, dbctx)
	http.RegisterHandlers(e)

	e.Logger.Fatal(http.Start(e))
}
