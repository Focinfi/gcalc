package gcalc

import (
	"github.com/Focinfi/gtester"
)

func sum(tokens []*Token) float64 {
	length := len(tokens)
	if length == 1 {
		return tokens[0].value
	}

	if length < 3 || length%2 != 1 {
		return float64(0)
	}

	var sum float64 = tokens[0].value
	for i := 1; i < length; i += 2 {
		if tokens[i].kind == NumberToken || tokens[i+1].kind != NumberToken {
			return float64(0)
		}

		sum = tokens[i].calc(sum, tokens[i+1].value)
	}

	return sum
}

func Calculate(exp string) (result float64, err error) {
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

	return sum(stuck.ToSlice()), nil
}
