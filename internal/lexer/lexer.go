package lexer

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/digvijay-tech/Haru/internal/lexer/tokenizer"
)

type Lexer struct {
	input    string
	position int
	line     int
	col      int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input, position: 0, line: 1, col: 1}
}

func (lex *Lexer) peek() string {
	// when EOF is reached
	if lex.position >= len(lex.input)-1 {
		return ""
	}

	return string(lex.input[lex.position+1])
}

func (lex *Lexer) NextToken() tokenizer.Token {
	// reached end of file
	if lex.position >= len(lex.input) {
		return tokenizer.Token{Type: tokenizer.EOF, Value: "", Line: lex.line, Col: lex.col}
	}

	// skip whitespaces
	for lex.position < len(lex.input) && unicode.IsSpace(rune(lex.input[lex.position])) {
		if lex.input[lex.position] == '\n' {
			lex.line++
			lex.col = 0
		}
		lex.position++
		lex.col++
	}

	// reads one character at a time
	ch := lex.input[lex.position]

	// classify operators
	switch {
	case ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '%' || ch == '=' || ch == '<' || ch == '>' || ch == '!' || ch == '&' || ch == '|':
		return lex.readOperators()
	default:
		lex.position++
		return lex.NextToken()
	}
}

func (lex *Lexer) readOperators() tokenizer.Token {
	start := lex.position // previous position
	lex.position++        // new position
	lex.col++

	if lex.position < len(lex.input) && strings.TrimSpace(string(lex.input[lex.position])) != "" {
		pattern := `==|!=|>=|<=|=>|\+=|-=|\*=|%=|/=|//=|>>|<<|\|\||&&|--|\+\+|//=|//|\*\*`
		regex := regexp.MustCompile(pattern)

		pairedOperator := string(lex.input[start]) + string(lex.input[lex.position])

		if regex.MatchString(pairedOperator) {
			lex.position++ // move past the second character of the operator
			lex.col++
			return tokenizer.Token{Type: tokenizer.OPERATOR, Value: string(pairedOperator), Line: lex.line, Col: lex.col}
		}
	}

	// regular operator
	return tokenizer.Token{Type: tokenizer.OPERATOR, Value: string(lex.input[start]), Line: lex.line, Col: lex.col}
}
