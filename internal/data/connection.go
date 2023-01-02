package data

import (
	"os"

	"github.com/palavrapasse/damn/pkg/database"
)

const (
	leaksDbFilePathEnvKey = "leaksdb.fp"
)

var (
	leaksDbFilePath = os.Getenv(leaksDbFilePathEnvKey)
)

func Open() (database.DatabaseContext, error) {
	return database.NewDatabaseContext(leaksDbFilePath)
}

func Close(dbctx database.DatabaseContext) error {
	return dbctx.DB.Close()
}
