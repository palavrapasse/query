package data

import (
	"database/sql"

	"github.com/palavrapasse/damn/pkg/database"
	"github.com/palavrapasse/damn/pkg/entity"
)

func QueryLeaksDB(dbctx database.DatabaseContext, hu []entity.HashUser) ([]entity.Leak, error) {

	var tx *sql.Tx
	var rs *sql.Rows
	var stmt *sql.Stmt
	var err error
	var ls []entity.Leak

	tx, err = dbctx.DB.Begin()

	if err != nil {
		return ls, err
	}

	stmt, err = tx.Prepare(
		`SELECT * FROM Leak L
			WHERE L.leakid IN (
				SELECT AU.leakid FROM AffectedUsers AU
				WHERE AU.userid = (
					SELECT HU.userid FROM HashUser HU
					WHERE HU.hsha256 = ?
				)
			)`,
	)

	if err != nil {
		return ls, err
	}

	rs, err = stmt.Query(hu[0].HSHA256)

	if err != nil {
		return ls, err
	}

	for rs.Next() {
		l := entity.Leak{}

		err = rs.Scan(&l.LeakId, &l.ShareDateSC, &l.Context)

		if err != nil {
			return ls, err
		} else {
			ls = append(ls, l)
		}
	}

	return ls, err
}
