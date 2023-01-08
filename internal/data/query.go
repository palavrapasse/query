package data

import (
	"github.com/palavrapasse/damn/pkg/database"
	"github.com/palavrapasse/damn/pkg/entity"
)

const leaksByUserHashPreparedQuery = `
SELECT * FROM Leak L
WHERE L.leakid IN (
	SELECT AU.leakid FROM AffectedUsers AU
	WHERE AU.userid = (
		SELECT HU.userid FROM HashUser HU
		WHERE HU.hsha256 = ?
	)
)`

var leaksQueryMapper = func() (*entity.Leak, []any) {
	l := entity.Leak{}
	return &l, []any{&l.LeakId, &l.ShareDateSC, &l.Context}
}

func QueryLeaksDB(dbctx database.DatabaseContext[database.Record], hu []entity.HashUser) ([]entity.Leak, error) {
	ctx := database.Convert[database.Record, entity.Leak](dbctx)

	return ctx.CustomQuery(leaksByUserHashPreparedQuery, leaksQueryMapper, hu[0].HSHA256)
}
