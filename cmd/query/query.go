package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/palavrapasse/query/internal/data"
	"github.com/palavrapasse/query/internal/http"
)

func main() {

	e := echo.New()

	defer e.Close()

	dbctx, oerr := data.Open()

	if oerr != nil {

		log.Printf("Could not open DB connection on server start")
		// todo (#10): log.Printf(oerr.Error())

		panic("DB connection is required to operate query")
	}

	http.RegisterMiddlewares(e, dbctx)
	http.RegisterHandlers(e)

	e.Logger.Fatal(http.Start(e))
}
