package gcalc

import (
	"fmt"
	. "github.com/Focinfi/gset"
)

var separators = NewSetSimple('+', '-', '*', '/', ' ', '(', ')')

type TokenKind int64

const (
	NumberToken TokenKind = iota << (iota + 32)
	AddOperatorToken
	SubOperatorToken
	MulOperatorToken
	DivOperatorToken
	LPToken
	RPToken
)

type Token struct {
	kind  TokenKind
	value float64
	str   string
}

func (t *Token) calc(firstV, secondV float64) float64 {
	switch t.kind {
	case AddOperatorToken:
		return firstV + secondV
	case SubOperatorToken:
		return firstV - secondV
	case MulOperatorToken:
		return firstV * secondV
	case DivOperatorToken:
		return firstV / secondV
	}
	return float64(0)
}

func (t Token) String() string {
	var fmtStr string
	if t.str != "" {
		fmtStr = t.str
	} else {
		fmtStr = fmt.Sprintf("%f", t.value)
	}

	return fmtStr
}
