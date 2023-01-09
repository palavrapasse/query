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

var leaksByUserQueryMapper = func() (*AffectedUserLeak, []any) {
	aul := AffectedUserLeak{}

	return &aul, []any{&aul.LeakId, &aul.ShareDateSC, &aul.Context, &aul.Email}
}

type AffectedUserLeak struct {
	entity.User
	entity.Leak
}

func QueryLeaksDB(dbctx database.DatabaseContext[database.Record], hus []entity.HashUser) ([]AffectedUserLeak, error) {
	q, vs := prepareAffectedUserQuery(hus)

	ctx := database.Convert[database.Record, AffectedUserLeak](dbctx)

	return ctx.CustomQuery(q, leaksByUserQueryMapper, vs...)
}

func prepareAffectedUserQuery(hus []entity.HashUser) (string, []any) {
	lhus := len(hus)

	values := make([]any, lhus)

	for i := 0; i < lhus; i++ {
		values[i] = hus[i].HSHA256
	}

	return fmt.Sprintf(leaksByUserHashPreparedQuery, database.MultiplePlaceholder(lhus)), values
}
