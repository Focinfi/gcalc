package gcalc

import (
	"fmt"
	"strconv"
)

var stLine string
var stLinePos int

type TokenKind int64

const (
	NumberToken TokenKind = iota << (iota + 32)
	AddOperatorToken
	SubOperatorToken
	MulOperatorToken
	DivOperatorToken
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

func setLine(line string) {

}

func getTokens(calcExp string) ([]Token, error) {
	sm := NumberSM{InitialState, ""}
	tokens := []Token{}
	for i, c := range calcExp {
		// singular number rune
		if c >= '0' && c <= '9' || c == '.' {
			if !sm.Feed(string(c)) {
				return tokens, fmt.Errorf("char %s can not append to the number %s", string(c), sm.Records())
			}
		} else {
			// operators and whitespace
			var token Token
			switch c {
			case '+':
				token = Token{kind: AddOperatorToken, str: string(c)}
			case '-':
				token = Token{kind: SubOperatorToken, str: string(c)}
			case '*':
				token = Token{kind: MulOperatorToken, str: string(c)}
			case '/':
				token = Token{kind: DivOperatorToken, str: string(c)}
			case ' ':
			default:
				return tokens, fmt.Errorf("unexpected char %s", string(c))
			}

			if sm.IsNumber() {
				numStr := sm.Records()
				f64, _ := strconv.ParseFloat(numStr, 64)
				tokens = append(tokens, Token{kind: NumberToken, value: f64})
				sm.Reset()
			}

			if token.str != "" {
				tokens = append(tokens, token)
			}
		}

		// last rune
		if i == len(calcExp)-1 {
			if sm.IsNumber() {
				numStr := sm.Records()
				f64, _ := strconv.ParseFloat(numStr, 64)
				tokens = append(tokens, Token{kind: NumberToken, value: f64})
				sm.Reset()
			} else if sm.Records() != "" {
				return tokens, fmt.Errorf("number %s is invlide", sm.Records())
			}
		}
	}

	return tokens, nil
}
