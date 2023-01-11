package data

import (
	"strings"

	"github.com/palavrapasse/damn/pkg/database"
	"github.com/palavrapasse/damn/pkg/entity"
)

const affectedSeparator = ","

const (
	EmailLeakType    LeakType = "email"
	PasswordLeakType LeakType = "password"
	UnknownLeakType  LeakType = "unknown"
)

const (
	AllTarget    Target = "all"
	NewestTarget Target = "newest"
	OldestTarget Target = "oldest"
)

type LeakType string
type Target string

type QueryLeaksResult struct {
	entity.User
	entity.Leak
}

func ParseLeakType(s string) LeakType {
	lt := LeakType(s)

	switch lt {
	case EmailLeakType:
	case PasswordLeakType:
	default:
		lt = UnknownLeakType
	}

	return lt
}

func ParseTarget(s string) Target {
	t := Target(s)

	switch t {
	case AllTarget:
	case NewestTarget:
	case OldestTarget:
	default:
		t = AllTarget
	}

	return t
}

func ParseAffected(s string) []string {
	aff := strings.Split(strings.ReplaceAll(s, " ", ""), affectedSeparator)
	laff := len(aff)

	if aff[0] == "" {
		aff = []string{}
	} else if aff[laff-1] == "" {
		aff = aff[:laff-1]
	}

	return aff
}

func AffectedToHashUser(aff []string) []entity.HashUser {
	laff := len(aff)
	hus := make([]entity.HashUser, laff)

	for i := 0; i < laff; i++ {
		hus[i] = entity.NewHashUser(entity.User{Email: entity.Email(aff[i])})
	}

	return hus
}

func (t Target) ToSQLKeyword() string {
	if t == OldestTarget {
		return database.AscendingSortOrderKeyword
	}

	return database.DescendingSortOrderKeyword
}
