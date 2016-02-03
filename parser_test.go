package gcalc

import (
	. "github.com/Focinfi/gtester"
	"testing"
)

func TestParse(t *testing.T) {
	tokens, _ := getTokens("1 + 2 * 3 + 4 /2")
	stuck := NewTokenStuck()
	err := Parse(tokens, stuck)
	AssertNilError(t, err)
	AssertEqual(t, stuck.Ret().value, 9.0)
}
