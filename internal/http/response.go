package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/palavrapasse/query/internal/data"
)

// Tooled with json-go-struct mapper: https://mholt.github.io/json-to-go/

type QueryLeaksView []QueryLeaksResultView
type QueryPlatformsView []QueryPlatformsResultView

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

type QueryPlatformsResultView string

type HealthStatusView struct {
	CPU           float64 `json:"cpu"`
	CPUPercentage float64 `json:"cpuPercentage"`
	RAM           float64 `json:"ram"`
	RAMPercentage float64 `json:"ramPercentage"`
	RAMMax        float64 `json:"ramMax"`
}

func ToQueryLeaksView(auls []data.QueryLeaksResult) QueryLeaksView {
	lls := len(auls)
	qlv := make(QueryLeaksView, lls)

	for i := 0; i < lls; i++ {
		qlv[i] = ToQueryLeaksResultView(auls[i])
	}

	return qlv
}

func ToQueryPlatformsView(qprs []data.QueryPlatformsResult) QueryPlatformsView {
	lps := len(qprs)
	qpv := make(QueryPlatformsView, lps)

	for i := 0; i < lps; i++ {
		qpv[i] = ToQueryPlatformsResultView(qprs[i])
	}

	return qpv
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

func ToQueryPlatformsResultView(qpr data.QueryPlatformsResult) QueryPlatformsResultView {
	return QueryPlatformsResultView(qpr.Name)
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
