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

const leaksQuery = `SELECT * FROM Leak`

var leaksByUserQueryMapper = func() (*QueryLeaksResult, []any) {
	aul := QueryLeaksResult{}

	return &aul, []any{&aul.LeakId, &aul.ShareDateSC, &aul.Context, &aul.Email}
}

var leaksQueryMapper = func() (*QueryLeaksResult, []any) {
	aul := QueryLeaksResult{}

	return &aul, []any{&aul.LeakId, &aul.ShareDateSC, &aul.Context}
}

func QueryLeaksDB(dbctx database.DatabaseContext[database.Record], tt Target, hus ...entity.HashUser) ([]QueryLeaksResult, error) {
	ctx := database.Convert[database.Record, QueryLeaksResult](dbctx)

	if len(hus) > 0 {
		return queryLeaksThatAffectUser(ctx, hus)
	} else {
		return queryLeaks(ctx, tt)
	}
}

func queryLeaksThatAffectUser(dbctx database.DatabaseContext[QueryLeaksResult], hus []entity.HashUser) ([]QueryLeaksResult, error) {
	q, m, vs := prepareAffectedUserQuery(hus)

	return dbctx.CustomQuery(q, m, vs...)
}

func queryLeaks(dbctx database.DatabaseContext[QueryLeaksResult], tt Target) ([]QueryLeaksResult, error) {
	q, m, vs := prepareLeaksQuery(tt)

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
	return leaksQuery, leaksQueryMapper, []any{}
}
