package authorization

import (
	"errors"
)

// Level is the type representing an authorization level.
type Level int

const (
	// Bypass bypass level.
	Bypass Level = iota
	// OneFactor one factor level.
	OneFactor Level = iota
	// TwoFactor two factor level.
	TwoFactor Level = iota
	// Denied denied level.
	Denied Level = iota
)

const userPrefix = "user:"
const groupPrefix = "group:"

var errNoMatchingRule = errors.New("no rule matching the request was found")
