package http

import (
	"strings"

	"github.com/palavrapasse/damn/pkg/entity"
)

const (
	leakTypeQueryParam = "type"
	affectedQueryParam = "affected"
	targetQueryParam   = "target"
)

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

func ParseAffected(s string) []string {
	return strings.Split(strings.ReplaceAll(s, " ", ""), ",")
}

func AffectedToHashUser(aff []string) []entity.HashUser {
	laff := len(aff)
	hus := make([]entity.HashUser, laff)

	for i := 0; i < laff; i++ {
		hus[i] = entity.NewHashUser(entity.User{Email: entity.Email(aff[i])})
	}

	return hus
}
