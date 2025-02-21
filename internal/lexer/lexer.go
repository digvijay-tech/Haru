package lexer

import (
	"unicode"

	"github.com/digvijay-tech/Haru/internal/lexer/tokenizer"
)

type Lexer struct {
	input     string
	position  int
	line, col int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: input, position: 0, line: 1, col: 1}
}

func (lex *Lexer) NextToken() tokenizer.Token {
	if lex.position >= len(lex.input) {
		return tokenizer.Token{
			Type:  tokenizer.EOF,
			Value: "",
			Line:  lex.line,
			Col:   lex.col,
		}
	}

	// skip whitespaces and handle newlines first
	for lex.position < len(lex.input) && unicode.IsSpace(rune(lex.input[lex.position])) {
		if lex.input[lex.position] == '\n' {
			lex.line++
			lex.col = 1
		} else {
			lex.col++
		}
		lex.position++
	}

	// reads only a character
	ch := lex.input[lex.position]

	// check for tuple type pattern: `:<...>` or `: <...>`
	if ch == ':' {
		start := lex.position
		lex.position++
		lex.col++

		// skip optional whitespace
		for lex.position < len(lex.input) && unicode.IsSpace(rune(lex.input[lex.position])) {
			if lex.input[lex.position] == '\n' {
				lex.line++
				lex.col = 1
			} else {
				lex.col++
			}
			lex.position++
		}

		if lex.position < len(lex.input) && lex.input[lex.position] == '<' {
			tupleStart := lex.position
			lex.position++ // consume '<'
			lex.col++
			stack := 1

			for lex.position < len(lex.input) {
				curr := lex.input[lex.position]
				if curr == '<' {
					stack++
				} else if curr == '>' {
					stack--
					if stack == 0 {
						lex.position++ // consume '>'
						lex.col++
						return tokenizer.Token{
							Type:  tokenizer.TUPLE_TYPE,
							Value: lex.input[tupleStart:lex.position],
							Line:  lex.line,
							Col:   tupleStart + 1,
						}
					}
				}
				lex.position++
				lex.col++
			}
		}

		// if no tuple found, return just the colon
		return tokenizer.Token{
			Type:  tokenizer.SYMBOL,
			Value: ":",
			Line:  lex.line,
			Col:   start + 1,
		}
	}

	switch {
	case unicode.IsLetter(rune(ch)) || ch == '_':
		return lex.readIdentifier()
	case unicode.IsDigit(rune(ch)):
		return lex.readNumber()
	case ch == '"':
		return lex.readString()
	case ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '=' || ch == '<' || ch == '>' || ch == '!' || ch == '&' || ch == '|':
		return lex.readOperator()
	case ch == '(' || ch == ')' || ch == '{' || ch == '}' || ch == '[' || ch == ']' || ch == ',' || ch == ':' || ch == ';':
		return lex.readSymbol()
	default:
		lex.position++
		return lex.NextToken()
	}
}

func (lex *Lexer) readIdentifier() tokenizer.Token {
	start := lex.position

	for lex.position < len(lex.input) && (unicode.IsLetter(rune(lex.input[lex.position])) || unicode.IsDigit(rune(lex.input[lex.position])) || lex.input[lex.position] == '_') {
		lex.position++
		lex.col++
	}

	word := lex.input[start:lex.position]

	if tokType, isKeyword := tokenizer.KeywordsTable[word]; isKeyword {
		return tokenizer.Token{
			Type:  tokType,
			Value: word,
			Line:  lex.line,
			Col:   lex.col - len(word),
		}
	}

	if tokType, isDatatype := tokenizer.DatatypesTable[word]; isDatatype {
		return tokenizer.Token{
			Type:  tokType,
			Value: word,
			Line:  lex.line,
			Col:   lex.col - len(word),
		}
	}

	return tokenizer.Token{
		Type:  tokenizer.IDENTIFIER,
		Value: word,
		Line:  lex.line,
		Col:   start + 1,
	}
}

func (lex *Lexer) readNumber() tokenizer.Token {
	start := lex.position
	hasDot := false

	for lex.position < len(lex.input) {
		ch := lex.input[lex.position]

		if ch == '.' {
			if hasDot {
				break // only one dot allowed
			}

			hasDot = true
		} else if !unicode.IsDigit(rune(ch)) {
			break
		}

		lex.position++
		lex.col++
	}

	return tokenizer.Token{
		Type:  tokenizer.NUMBER,
		Value: lex.input[start:lex.position],
		Line:  lex.line,
		Col:   start + 1,
	}
}

func (lex *Lexer) readString() tokenizer.Token {
	start := lex.position

	// skip opening "
	lex.position++
	lex.col++

	for lex.position < len(lex.input) && lex.input[lex.position] != '"' {
		lex.position++
		lex.col++
	}

	// skip closing "
	lex.position++
	lex.col++

	return tokenizer.Token{
		Type:  tokenizer.STRING,
		Value: lex.input[start+1 : lex.position-1],
		Line:  lex.line,
		Col:   start + 1,
	}
}

func (lex *Lexer) readOperator() tokenizer.Token {
	start := lex.position

	lex.position++
	lex.col++

	if lex.position < len(lex.input) && (lex.input[lex.position] == '=' || lex.input[lex.position] == '&' || lex.input[lex.position] == '|') {
		lex.position++
		lex.col++
	}

	return tokenizer.Token{
		Type:  tokenizer.OPERATOR,
		Value: lex.input[start:lex.position],
		Line:  lex.line,
		Col:   start + 1,
	}
}

func (lex *Lexer) readSymbol() tokenizer.Token {
	start := lex.position
	ch := lex.input[lex.position]
	lex.position++
	lex.col++

	return tokenizer.Token{
		Type:  tokenizer.SYMBOL,
		Value: string(ch),
		Line:  lex.line,
		Col:   start + 1,
	}
}
