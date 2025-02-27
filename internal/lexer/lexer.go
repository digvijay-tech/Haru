package lexer

import (
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"

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
	case ch == '"' || ch == '\'' || ch == '`':
		return lex.readString()
	case unicode.IsNumber(rune(ch)):
		return lex.readNumbers()
	case ch == '_' || unicode.IsLetter(rune(ch)):
		return lex.classifyWord()
	case ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '%' || ch == '=' || ch == '<' || ch == '>' || ch == '!' || ch == '&' || ch == '|':
		return lex.readOperators()
	case unicode.IsPunct(rune(ch)):
		return lex.readPunctuations()
	default:
		lex.position++
		return lex.NextToken()
	}
}

func (lex *Lexer) readNumbers() tokenizer.Token {
	start := lex.position
	lex.position++
	lex.col++
	hasDot := false

	for lex.position < len(lex.input) {
		char := rune(lex.input[lex.position])

		// stop when a whitespace or non-numeric characters are encountered
		if unicode.IsSpace(char) || unicode.IsLetter(char) {
			break
		}

		// encountered a decimal point (only allowed once)
		if char == '.' {
			if hasDot {
				// invalid number (multiple dots)
				for lex.position < len(lex.input) && !unicode.IsSpace(rune(lex.input[lex.position])) {
					lex.position++
					lex.col++
				}

				return tokenizer.Token{Type: tokenizer.ERROR, Value: "Invalid Numeric Value", Line: lex.line, Col: lex.col}
			}

			// Edge case: A number cannot end with `.`
			if lex.position+1 >= len(lex.input) || !unicode.IsDigit(rune(lex.input[lex.position+1])) {
				return tokenizer.Token{Type: tokenizer.ERROR, Value: "Invalid Float Format", Line: lex.line, Col: lex.col}
			}

			hasDot = true
		}

		// encountered a punctuation other than `.`
		if unicode.IsPunct(char) && char != '.' {
			break
		}

		lex.position++
		lex.col++
	}

	return tokenizer.Token{Type: tokenizer.NUMBER, Value: lex.input[start:lex.position], Line: lex.line, Col: lex.col}
}

func (lex *Lexer) readString() tokenizer.Token {
	start := lex.position
	lex.position++
	lex.col++

	for lex.position < len(lex.input) {
		// found the same ending punctuation that initialized the string it could be (" or ')
		if lex.input[start] == lex.input[lex.position] {
			lex.position++ // move past the closing quote
			lex.col++
			break
		}

		// encountered a line-break
		if lex.input[lex.position] == '\n' {
			lex.position++
			lex.line++
			lex.col = 1
			continue
		}

		lex.position++
		lex.col++
	}

	// unterminated string
	if lex.position >= len(lex.input) || lex.input[lex.position-1] != lex.input[start] {
		return tokenizer.Token{Type: tokenizer.ERROR, Value: "Unterminated String", Line: lex.line, Col: lex.col}
	}

	// excluding quotations
	value := lex.input[start+1 : lex.position-1]
	return tokenizer.Token{Type: tokenizer.STRING, Value: value, Line: lex.line, Col: lex.col}
}

func (lex *Lexer) classifyWord() tokenizer.Token {
	start := lex.position
	lex.position++
	lex.col++

	for lex.position < len(lex.input) && !unicode.IsSpace(rune(lex.input[lex.position])) {
		// making exception to allow the use of `_` anywhere in identifer
		if lex.input[lex.position] == '_' {
			lex.position++
			lex.col++
			continue // jump to next iteration
		}

		if lex.input[lex.position] == '<' {
			break
		}

		if lex.input[lex.position] == '>' {
			break
		}

		// prevent any other punctuations in the identifier
		if unicode.IsPunct(rune(lex.input[lex.position])) {
			break
		}

		lex.position++
		lex.col++
	}

	word := lex.input[start:min(lex.position, len(lex.input))]

	// reserved keyword
	if _, exists := tokenizer.KeywordsTable[word]; exists {
		return tokenizer.Token{Type: tokenizer.KEYWORD, Value: word, Line: lex.line, Col: lex.col}
	}

	// datatype
	if _, exists := tokenizer.DatatypesTable[word]; exists {
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

func (lex *Lexer) readPunctuations() tokenizer.Token {
	// reached end of file
	if lex.position >= len(lex.input) {
		return tokenizer.Token{Type: tokenizer.EOF, Value: "", Line: lex.line, Col: lex.col}
	}

	_, size := utf8.DecodeLastRuneInString(lex.input[lex.position:])

	start := lex.position
	lex.position += size
	lex.col++

	return tokenizer.Token{Type: tokenizer.PUNCTUATION, Value: string(lex.input[start : start+size]), Line: lex.line, Col: lex.col}
}
