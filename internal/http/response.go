package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/palavrapasse/damn/pkg/entity"
)

// Tooled with json-go-struct mapper: https://mholt.github.io/json-to-go/

type QueryLeaksView []LeakView

type LeakView struct {
	Context          string `json:"context"`
	ShareDateMSEpoch int64  `json:"shareDateMSEpoch"`
}

func ToQueryLeaksView(ls []entity.Leak) QueryLeaksView {
	lls := len(ls)
	qlv := make(QueryLeaksView, lls)

	for i := 0; i < lls; i++ {
		qlv[i] = ToLeakView(ls[i])
	}

	return qlv
}

func ToLeakView(l entity.Leak) LeakView {
	return LeakView{
		Context:          string(l.Context),
		ShareDateMSEpoch: int64(l.ShareDateSC) * 1000,
	}
}

func NoContent(ectx echo.Context) error {
	return ectx.NoContent(http.StatusNoContent)
}

func Ok(ectx echo.Context, i interface{}) error {
	return ectx.JSON(http.StatusOK, i)
}

func InternalServerError(ectx echo.Context) error {
	return ectx.NoContent(http.StatusInternalServerError)
}
