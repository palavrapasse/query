package http

import (
	"fmt"
	"math"
	"net/http"
	"runtime"

	"github.com/labstack/echo/v4"
	"github.com/palavrapasse/query/internal/data"
	"github.com/palavrapasse/query/internal/logging"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func RegisterHandlers(e *echo.Echo) {

	e.GET(leaksRoute, QueryLeaks)
	e.GET(platformsRoute, QueryPlatforms)
	e.GET(healthCheckRoute, QueryHealthCheck)

	echo.NotFoundHandler = useNotFoundHandler()
}

func QueryLeaks(ectx echo.Context) error {
	var affu []data.QueryLeaksResult
	var err error

	logging.Aspirador.Trace("Querying leaks")

	mwctx, gmerr := GetMiddlewareContext(ectx)

	if gmerr != nil {
		return InternalServerError(ectx)
	}

	affp := ectx.QueryParam(affectedQueryParam)
	ttp := ectx.QueryParam(targetQueryParam)

	aff := data.ParseAffected(affp)
	tt := data.ParseTarget(ttp)
	hus := data.AffectedToHashUser(aff)

	affu, err = data.QueryLeaksDB(mwctx.DB, tt, hus...)

	if err != nil {
		logging.Aspirador.Error(fmt.Sprintf("Error while querying Leaks from DB: %s", err))

		return InternalServerError(ectx)
	}

	logging.Aspirador.Trace(fmt.Sprintf("Success in querying leaks. Found %d leaks", len(affu)))

	return Ok(ectx, ToQueryLeaksView(affu))
}

func QueryPlatforms(ectx echo.Context) error {
	var qprs []data.QueryPlatformsResult
	var err error

	logging.Aspirador.Trace("Querying platforms")

	mwctx, gmerr := GetMiddlewareContext(ectx)

	if gmerr != nil {
		return InternalServerError(ectx)
	}

	ttp := ectx.QueryParam(targetQueryParam)

	tt := data.ParseTarget(ttp)

	qprs, err = data.QueryPlaformsDB(mwctx.DB, tt)

	if err != nil {
		logging.Aspirador.Error(fmt.Sprintf("Error while querying Platforms from DB: %s", err))

		return InternalServerError(ectx)
	}

	logging.Aspirador.Trace(fmt.Sprintf("Success in querying platforms. Found %d platforms", len(qprs)))

	return Ok(ectx, ToQueryPlatformsView(qprs))
}

func QueryHealthCheck(ectx echo.Context) error {

	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return InternalServerError(ectx)
	}

	cpuPercentages, err := cpu.Percent(0, false)
	if err != nil {
		return InternalServerError(ectx)
	}

	cpu := runtime.NumCPU()
	cpuPercentage := math.Round(cpuPercentages[0]*100) / 100

	ram := vmStat.Total
	ramMax := ram + vmStat.Free
	ramPercentage := math.Round(vmStat.UsedPercent*100) / 100

	result := HealthStatusView{
		CPU:           float64(cpu),
		CPUPercentage: float64(cpuPercentage),
		RAM:           float64(ram),
		RAMMax:        float64(ramMax),
		RAMPercentage: float64(ramPercentage),
	}
	logging.Aspirador.Trace(fmt.Sprintf("%v", result))

	return Ok(ectx, result)
}

func useNotFoundHandler() func(c echo.Context) error {
	return func(c echo.Context) error {
		return c.NoContent(http.StatusNotFound)
	}
}
