package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/palavrapasse/query/internal/data"
)

// Tooled with json-go-struct mapper: https://mholt.github.io/json-to-go/

type QueryLeaksView []QueryLeaksResultView

type LeakView struct {
	Context          string `json:"context"`
	ShareDateMSEpoch int64  `json:"shareDateMSEpoch"`
}

type AffectedUserView struct {
	Email string `json:"email,omitempty"`
}

type QueryLeaksResultView struct {
	AffectedUserView
	LeakView
}

func ToQueryLeaksView(auls []data.QueryLeaksResult) QueryLeaksView {
	lls := len(auls)
	qlv := make(QueryLeaksView, lls)

	for i := 0; i < lls; i++ {
		qlv[i] = ToQueryLeaksResultView(auls[i])
	}

	return qlv
}

func ToQueryLeaksResultView(aul data.QueryLeaksResult) QueryLeaksResultView {
	return QueryLeaksResultView{
		LeakView: LeakView{
			Context:          string(aul.Context),
			ShareDateMSEpoch: int64(aul.ShareDateSC) * 1000,
		},
		AffectedUserView: AffectedUserView{
			Email: string(aul.Email),
		},
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
