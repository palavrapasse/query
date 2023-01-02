package http

import (
	"errors"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/palavrapasse/damn/pkg/database"
)

const (
	dbMiddlewareKey = "db"
)

type MiddlewareContext struct {
	DB database.DatabaseContext
}

func RegisterMiddlewares(e *echo.Echo, dbctx database.DatabaseContext) {
	e.Use(dbAccessMiddleware(dbctx))
	e.Use(loggingMiddleware())
}

func GetMiddlewareContext(ectx echo.Context) (MiddlewareContext, error) {
	db, dok := ectx.Get(dbMiddlewareKey).(database.DatabaseContext)
	var err error

	if !dok {
		err = errors.New("DB not available in middleware")

		log.Printf("%v\n", err)
	}

	return MiddlewareContext{
		DB: db,
	}, err
}

func dbAccessMiddleware(dbctx database.DatabaseContext) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {
			ectx.Set(dbMiddlewareKey, dbctx)
			next(ectx)
			return nil
		}
	}
}

func loggingMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ectx echo.Context) error {

			req := ectx.Request()

			log.Printf(fmt.Sprintf("Host: %s | Method: %s | Path: %s", req.Host, req.Method, req.URL.RequestURI()))

			next(ectx)

			return nil
		}
	}
}
