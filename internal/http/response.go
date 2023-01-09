package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/palavrapasse/query/internal/data"
)

// Tooled with json-go-struct mapper: https://mholt.github.io/json-to-go/

type QueryAffectedUserLeaksView []AffectedUserLeakView

type LeakView struct {
	Context          string `json:"context"`
	ShareDateMSEpoch int64  `json:"shareDateMSEpoch"`
}

type AffectedUserLeakView struct {
	Email string `json:"email"`
	LeakView
}

func ToQueryAffectedUserLeaksView(auls []data.AffectedUserLeak) QueryAffectedUserLeaksView {
	lls := len(auls)
	qlv := make(QueryAffectedUserLeaksView, lls)

	for i := 0; i < lls; i++ {
		qlv[i] = ToAffectedUserLeakView(auls[i])
	}

	return qlv
}

func ToAffectedUserLeakView(aul data.AffectedUserLeak) AffectedUserLeakView {
	return AffectedUserLeakView{
		LeakView: LeakView{
			Context:          string(aul.Context),
			ShareDateMSEpoch: int64(aul.ShareDateSC) * 1000,
		},
		Email: string(aul.Email),
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
