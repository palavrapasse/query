package http

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
