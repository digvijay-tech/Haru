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
	case ch == '_' || unicode.IsLetter(rune(ch)):
		return lex.classifyWord()
	case ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '%' || ch == '=' || ch == '<' || ch == '>' || ch == '!' || ch == '&' || ch == '|':
		return lex.readOperators()
	default:
		lex.position++
		return lex.NextToken()
	}
}

func (lex *Lexer) classifyWord() tokenizer.Token {
	start := lex.position
	lex.position++
	lex.col++

	for lex.position < len(lex.input) && !unicode.IsSpace(rune(lex.input[lex.position])) {

		// making exception for `_`
		if lex.input[lex.position] == '_' {
			lex.position++
			lex.col++
			continue // jump to next iteration
		}

		// prevent any other punctuations in the identifier
		if unicode.IsPunct(rune(lex.input[lex.position])) {
			break
		}

		lex.position++
		lex.col++
	}

	word := lex.input[start:lex.position]

	// reserved keyword
	keyword := tokenizer.KeywordsTable[word]
	if keyword != "" {
		return tokenizer.Token{Type: tokenizer.KEYWORD, Value: word, Line: lex.line, Col: lex.col}
	}

	// datatype
	datatype := tokenizer.DatatypesTable[word]
	if datatype != "" {
		return tokenizer.Token{Type: tokenizer.DATATYPE, Value: word, Line: lex.line, Col: lex.col}
	}

	// user defined identifier
	return tokenizer.Token{Type: tokenizer.IDENTIFIER, Value: word, Line: lex.line, Col: lex.col}
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
