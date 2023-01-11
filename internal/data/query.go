package data

import (
	"fmt"

	"github.com/palavrapasse/damn/pkg/database"
	"github.com/palavrapasse/damn/pkg/entity"
)

const leaksByUserHashPreparedQuery = `
SELECT L.*, U.email FROM Leak L, User U
WHERE U.userid IN (
	SELECT HU.userid FROM HashUser HU
	WHERE HU.hsha256 IN (%s)
)
AND L.leakid IN (
	SELECT AU.leakid FROM AffectedUsers AU
	WHERE AU.userid = U.userid
)
`

const leaksQuery = `
SELECT L.* FROM Leak L
ORDER BY L.sharedatesc %s
`

const platformsQuery = `SELECT P.name FROM Platform P`

const platformsByLeaksQuery = `
SELECT P.name FROM LeakPlatform LP
INNER JOIN Leak L ON L.leakid == LP.leakid
INNER JOIN Platform P on P.platid == LP.platid
ORDER BY L.sharedatesc %s
`

var leaksByUserQueryMapper = func() (*QueryLeaksResult, []any) {
	aul := QueryLeaksResult{}

	return &aul, []any{&aul.LeakId, &aul.ShareDateSC, &aul.Context, &aul.Email}
}

var leaksQueryMapper = func() (*QueryLeaksResult, []any) {
	aul := QueryLeaksResult{}

	return &aul, []any{&aul.LeakId, &aul.ShareDateSC, &aul.Context}
}

var platformsQueryMapper = func() (*QueryPlatformsResult, []any) {
	qpr := QueryPlatformsResult{}

	return &qpr, []any{&qpr.Name}
}

func QueryLeaksDB(dbctx database.DatabaseContext[database.Record], tt Target, hus ...entity.HashUser) ([]QueryLeaksResult, error) {
	ctx := database.Convert[database.Record, QueryLeaksResult](dbctx)

	if len(hus) > 0 {
		return queryLeaksThatAffectUser(ctx, hus)
	}

	return queryLeaks(ctx, tt)
}

func QueryPlaformsDB(dbctx database.DatabaseContext[database.Record], tt Target) ([]QueryPlatformsResult, error) {
	ctx := database.Convert[database.Record, QueryPlatformsResult](dbctx)

	if tt == AllTarget {
		return queryPlatforms(ctx)
	}

	return queryPlatformsByLeaks(ctx, tt)
}

func queryLeaksThatAffectUser(dbctx database.DatabaseContext[QueryLeaksResult], hus []entity.HashUser) ([]QueryLeaksResult, error) {
	q, m, vs := prepareAffectedUserQuery(hus)

	return dbctx.CustomQuery(q, m, vs...)
}

func queryLeaks(dbctx database.DatabaseContext[QueryLeaksResult], tt Target) ([]QueryLeaksResult, error) {
	q, m, vs := prepareLeaksQuery(tt)

	return dbctx.CustomQuery(q, m, vs...)
}

func queryPlatforms(dbctx database.DatabaseContext[QueryPlatformsResult]) ([]QueryPlatformsResult, error) {
	q, m, vs := preparePlatformsQuery()

	return dbctx.CustomQuery(q, m, vs...)
}

func queryPlatformsByLeaks(dbctx database.DatabaseContext[QueryPlatformsResult], tt Target) ([]QueryPlatformsResult, error) {
	q, m, vs := preparePlatformsByLeaksQuery(tt)

	return dbctx.CustomQuery(q, m, vs...)
}

func prepareAffectedUserQuery(hus []entity.HashUser) (string, database.TypedQueryResultMapper[QueryLeaksResult], []any) {
	lhus := len(hus)

	values := make([]any, lhus)

	for i := 0; i < lhus; i++ {
		values[i] = hus[i].HSHA256
	}

	return fmt.Sprintf(leaksByUserHashPreparedQuery, database.MultiplePlaceholder(lhus)), leaksByUserQueryMapper, values
}

func prepareLeaksQuery(tt Target) (string, database.TypedQueryResultMapper[QueryLeaksResult], []any) {
	return fmt.Sprintf(leaksQuery, tt.ToSQLKeyword()), leaksQueryMapper, []any{}
}

func preparePlatformsQuery() (string, database.TypedQueryResultMapper[QueryPlatformsResult], []any) {
	return platformsQuery, platformsQueryMapper, []any{}
}

func preparePlatformsByLeaksQuery(tt Target) (string, database.TypedQueryResultMapper[QueryPlatformsResult], []any) {
	return fmt.Sprintf(platformsByLeaksQuery, tt.ToSQLKeyword()), platformsQueryMapper, []any{}
}
