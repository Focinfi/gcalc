package gcalc

import (
	"fmt"
	"strconv"
	"strings"
)

type Calculator struct {
	exp                  []rune
	position             int
	nsm                  *NumberSM
	currentRune          rune
	isCurrentRuneExists  bool
	currentToken         *Token
	isCurrentTokenExists bool
}

func NewCalculator(exp string) *Calculator {
	return &Calculator{exp: ([]rune)(strings.TrimSpace(exp)), nsm: &NumberSM{InitialState, ""}}
}

func (calc *Calculator) getRune() rune {
	if !calc.isCurrentRuneExists {
		calc.currentRune = calc.exp[calc.position]
		calc.position++
	}

	calc.isCurrentRuneExists = false
	return calc.currentRune
}

func (calc *Calculator) ungetRune(r rune) {
	calc.isCurrentRuneExists = true
	calc.currentRune = r
}

func (calc *Calculator) getToken() *Token {
	if !calc.isCurrentTokenExists {
		calc.currentToken = calc.nextToken()
	}

	calc.isCurrentTokenExists = false
	return calc.currentToken
}

func (calc *Calculator) ungetToken(token *Token) {
	calc.isCurrentTokenExists = true
	calc.currentToken = token
}

func (calc *Calculator) nextToken() *Token {
	var token *Token
	var err error
	for calc.position <= len(calc.exp) && token == nil {
		fmt.Println("[nextToken-before]", calc.position)
		r := calc.getRune()
		fmt.Println("[nextToken-after]", calc.position, string(r))
		if token, err = calc.checkRune(r); err != nil {
			panic(err.Error())
		}

		if token == nil && calc.position == len(calc.exp) {
			// fmt.Println("[records]", calc.position)
			if calc.nsm.IsNumber() {
				f64, _ := strconv.ParseFloat(calc.nsm.Records(), 64)
				token = &Token{kind: NumberToken, value: f64}
				calc.nsm.Reset()
			} else {
				panic(fmt.Sprintf("number %s is invlide", calc.nsm.Records()))
			}
		}
	}

	calc.currentToken = token
	return token
}

func (calc *Calculator) checkRune(c rune) (*Token, error) {
	if c >= '0' && c <= '9' || c == '.' {
		if !calc.nsm.Feed(string(c)) {
			return nil, fmt.Errorf("char %s can not append to the number %s", string(c), calc.nsm.Records())
		}
	} else if separators.Has(c) {
		if calc.nsm.IsNumber() {
			f64, _ := strconv.ParseFloat(calc.nsm.Records(), 64)
			calc.nsm.Reset()
			calc.ungetRune(c)
			return &Token{kind: NumberToken, value: f64}, nil
		}

		// operators and whitespace
		var token *Token
		switch c {
		case '+':
			token = &Token{kind: AddOperatorToken, str: string(c)}
		case '-':
			token = &Token{kind: SubOperatorToken, str: string(c)}
		case '*':
			token = &Token{kind: MulOperatorToken, str: string(c)}
		case '/':
			token = &Token{kind: DivOperatorToken, str: string(c)}
		case '(':
			token = &Token{kind: LPToken, str: string(c)}
		case ')':
			token = &Token{kind: RPToken, str: string(c)}
		case ' ':
		default:
			return nil, fmt.Errorf("unexpected char %s", string(c))
		}

		if token != nil {
			return token, nil
		}
	}

	return nil, nil
}

func (calc *Calculator) parseExpression() float64 {
	v1 := calc.parseTerm()

	for {
		if calc.position > len(calc.exp)-1 {
			break
		}

		fmt.Println("[parseExpression v1]", v1)
		token := calc.getToken()
		if token.kind == AddOperatorToken || token.kind == SubOperatorToken {
			v1 = token.calc(v1, calc.parseTerm())
		} else {
			calc.ungetToken(token)
			break
		}
	}

	return v1
}

func (calc *Calculator) parseTerm() float64 {
	v1 := calc.prasePrimaryExpression()

	for {
		if calc.position > len(calc.exp)-1 {
			break
		}

		token := calc.getToken()
		// fmt.Println("[parseTerm token]", token)
		if token.kind == MulOperatorToken || token.kind == DivOperatorToken {
			v1 = token.calc(v1, calc.prasePrimaryExpression())
		} else {
			calc.ungetToken(token)
			break
		}
	}

	return v1
}

func (calc *Calculator) prasePrimaryExpression() float64 {
	token := calc.getToken()
	if token.kind == SubOperatorToken {
		v1 := calc.prasePrimaryExpression()
		return -1 * v1
	} else if token.kind == NumberToken {
		return token.value
	} else if token.kind == LPToken {
		v1 := calc.parseExpression()
		token = calc.getToken()
		if token.kind == RPToken {
			return v1
		} else {
			panic(fmt.Sprintln("invalid expression:", token))
		}
	} else {
		panic(fmt.Sprintln("invalid expression:", token))
	}

	return float64(0)
}

func (calc *Calculator) Calculate() float64 {
	return calc.parseExpression()
}
