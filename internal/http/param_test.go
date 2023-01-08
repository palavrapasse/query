package http

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseLeakTypeReturnsEmailLeakTypeIfStringMatchesEmail(t *testing.T) {
	s := "email"

	lt := ParseLeakType(s)
	elt := EmailLeakType

	if lt != elt {
		t.Fatalf("function should have returned (%s), but got: (%s)", elt, lt)
	}
}

func TestParseLeakTypeReturnsPasswordLeakTypeIfStringMatchesPassword(t *testing.T) {
	s := "password"

	lt := ParseLeakType(s)
	elt := PasswordLeakType

	if lt != elt {
		t.Fatalf("function should have returned (%s), but got: (%s)", elt, lt)
	}
}

func TestParseLeakTypeReturnsUnknownLeakTypeIfStringCannotBeMatched(t *testing.T) {
	s := "unsupported leak type"

	lt := ParseLeakType(s)
	elt := UnknownLeakType

	if lt != elt {
		t.Fatalf("function should have returned (%s), but got: (%s)", elt, lt)
	}
}

func TestParseTargetReturnsAllTargetIfStringMatchesAll(t *testing.T) {
	s := "all"

	tt := ParseTarget(s)
	ett := AllTarget

	if tt != ett {
		t.Fatalf("function should have returned (%s), but got: (%s)", ett, tt)
	}
}

func TestParseTargetReturnsNewestTargetIfStringMatchesNewest(t *testing.T) {
	s := "newest"

	tt := ParseTarget(s)
	ett := NewestTarget

	if tt != ett {
		t.Fatalf("function should have returned (%s), but got: (%s)", ett, tt)
	}
}

func TestParseTargetReturnsOldestTargetIfStringMatchesOldest(t *testing.T) {
	s := "oldest"

	tt := ParseTarget(s)
	ett := OldestTarget

	if tt != ett {
		t.Fatalf("function should have returned (%s), but got: (%s)", ett, tt)
	}
}

func TestParseTargetReturnsAllTargetIfStringCannotBeMatched(t *testing.T) {
	s := "unsupported target"

	tt := ParseTarget(s)
	ett := AllTarget

	if tt != ett {
		t.Fatalf("function should have returned (%s), but got: (%s)", ett, tt)
	}
}

func TestParseAffectedReturnsSingleEntrySliceIfAffectedContainsOneEntry(t *testing.T) {
	ae := "my.email@email.com"
	s := ae

	aff := ParseAffected(s)
	eaff := []string{ae}

	if !reflect.DeepEqual(aff, eaff) {
		t.Fatalf("function should have returned (%v: len=%d), but got: (%v: len=%d)", eaff, len(eaff), aff, len(aff))
	}
}

func TestParseAffectedReturnsSingleEntrySliceIfAffectedContainsOneEntryAndIsFinishedWithSeparator(t *testing.T) {
	ae := "my.email@email.com"
	s := fmt.Sprintf("%s%s", ae, affectedSeparator)

	aff := ParseAffected(s)
	eaff := []string{ae}

	if !reflect.DeepEqual(aff, eaff) {
		t.Fatalf("function should have returned (%v: len=%d), but got: (%v: len=%d)", eaff, len(eaff), aff, len(aff))
	}
}

func TestParseAffectedReturnsMultipleEntriesSliceIfAffectedContainsMultipleEntries(t *testing.T) {
	ae := "my.email@email.com"
	s := fmt.Sprintf("%s%s%s", ae, affectedSeparator, ae)

	aff := ParseAffected(s)
	eaff := []string{ae, ae}

	if !reflect.DeepEqual(aff, eaff) {
		t.Fatalf("function should have returned (%v: len=%d), but got: (%v: len=%d)", eaff, len(eaff), aff, len(aff))
	}
}

func TestParseAffectedReturnsEmptySliceIfAffectedContainsSpaces(t *testing.T) {
	s := "           "

	aff := ParseAffected(s)
	eaff := []string{}

	if !reflect.DeepEqual(aff, eaff) {
		t.Fatalf("function should have returned (%v: len=%d), but got: (%v: len=%d)", eaff, len(eaff), aff, len(aff))
	}
}

func TestParseAffectedReturnsEmptySliceIfAffectedContainsOnlySeparators(t *testing.T) {
	s := fmt.Sprintf("%s%s", affectedSeparator, affectedSeparator)

	aff := ParseAffected(s)
	eaff := []string{}

	if !reflect.DeepEqual(aff, eaff) {
		t.Fatalf("function should have returned (%v: len=%d), but got: (%v: len=%d)", eaff, len(eaff), aff, len(aff))
	}
}

func TestParseAffectedReturnsEmptySliceIfAffectedContainsSpacesAndSeparators(t *testing.T) {
	s := fmt.Sprintf("     %s  %s    ", affectedSeparator, affectedSeparator)

	aff := ParseAffected(s)
	eaff := []string{}

	if !reflect.DeepEqual(aff, eaff) {
		t.Fatalf("function should have returned (%v: len=%d), but got: (%v: len=%d)", eaff, len(eaff), aff, len(aff))
	}
}
