package gcalc

import (
	"github.com/Focinfi/gtester"
)

func Caculate(exp string) (result float64, err error) {
	tokens := []Token{}
	stuck := NewTokenStuck()

	err = gtester.NewCheckQueue().Add(func() (err error) {
		tokens, err = getTokens(exp)
		return err
	}).Add(func() error {
		return Parse(tokens, stuck)
	}).Run()

	if err != nil {
		return
	}

	return stuck.Ret().value, nil
}
