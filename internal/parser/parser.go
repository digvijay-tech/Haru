package parser

import (
	"fmt"

	"github.com/digvijay-tech/Haru/internal/lexer/tokenizer"
	"github.com/digvijay-tech/Haru/internal/parser/ast"
)

type Parser struct {
	Tokens  []tokenizer.Token
	Current int
}

/*
Expect consumes and returns the next token, assuming that it must be of a specified type.
If the expectation is not met, it can either return an error or handle the situation gracefully.
*/
func (p *Parser) expect(tokenType tokenizer.TokenType) tokenizer.Token {
	if p.Current >= len(p.Tokens) {
		return tokenizer.Token{
			Type:  tokenizer.EOF,
			Value: "",
			Line:  -1,
			Col:   -1,
		}
	}

	currentToken := p.Tokens[p.Current]

	// ensure if the token is of expected type
	if currentToken.Type != tokenType {
		fmt.Printf("Syntax Error: Expected %s but got %s at Line %d, Col %d\n", tokenType, currentToken.Type, currentToken.Line, currentToken.Col)

		return tokenizer.Token{
			Type:  tokenizer.ERROR,
			Value: "Unexpected Token",
			Line:  currentToken.Line,
			Col:   currentToken.Col,
		}
	}

	p.Current++
	return currentToken
}

// Match checks if the next token matches the given type and advances if true
func (p *Parser) match(tokenType tokenizer.TokenType, value string) bool {
	if p.Current >= len(p.Tokens) {
		return false
	}

	if p.Tokens[p.Current].Type == tokenType && p.Tokens[p.Current].Value == value {
		p.Current++
		return true
	}

	return false
}

// Parsing Let/Mut statements
func (p *Parser) ParseLetAndMutStatement() *ast.VarStatement {
	// ensuring first token is always `let` or `mut`
	token := p.expect(tokenizer.KEYWORD)
	if token.Value != "let" && token.Value != "mut" {
		// not a variable declaration
		return nil
	}

	mutable := (token.Value == "mut")

	// expecting an identifier (a variable name)
	ident := p.expect(tokenizer.IDENTIFIER)
	if ident.Type != tokenizer.IDENTIFIER {
		fmt.Printf("Syntax Error: Expected identifier at Line %d, Col %d\n", ident.Line, ident.Col)
		return nil
	}

	// optional type annotation
	typeName := ""

	if p.match(tokenizer.PUNCTUATION, ":") {
		typeToken := p.expect(tokenizer.DATATYPE)
		if typeToken.Type != tokenizer.DATATYPE {
			fmt.Printf("Syntax Error: Expected data type after ':' at Line %d, Col %d\n", typeToken.Line, typeToken.Col)
			return nil
		}

		typeName = typeToken.Value
	}

	// expecting assignment
	if !p.match(tokenizer.OPERATOR, "=") {
		fmt.Printf("Syntax Error: Expected '=' after variable declaration at Line %d, Col %d\n", p.Tokens[p.Current].Line, p.Tokens[p.Current].Col)
		return nil
	}

	// parsing expression
	expr := p.parseExpression()
	if expr == nil {
		fmt.Printf("Syntax Error: Invalid expression after '=' at Line %d, Col %d\n", p.Tokens[p.Current].Line, p.Tokens[p.Current].Col)
		return nil
	}

	// skiping the semicolon
	if !p.match(tokenizer.PUNCTUATION, ";") {
		fmt.Printf("Syntax Error: Expected ';' after variable declaration at Line %d, Col %d\n", p.Tokens[p.Current].Line, p.Tokens[p.Current].Col)
		return nil
	}

	return &ast.VarStatement{
		Identifier: ident.Value,
		Type:       typeName,
		Value:      expr,
		Mutable:    mutable,
		Line:       ident.Line,
		Col:        ident.Col,
	}
}

// Handling `+` and `-`
func (p *Parser) parseExpression() ast.Expression {
	// handles higher precedence first
	left := p.parseTerm()

	for p.Current < len(p.Tokens) {
		tok := p.Tokens[p.Current]

		if tok.Type != tokenizer.OPERATOR || (tok.Value != "+" && tok.Value != "-") {
			break
		}

		operator := tok.Value
		p.Current++

		right := p.parseTerm()

		left = &ast.BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    right,
			Line:     tok.Line,
			Col:      tok.Col,
		}
	}

	return left
}

// Handling `*` and `/`
func (p *Parser) parseTerm() ast.Expression {
	left := p.parseFactor()

	for p.Current < len(p.Tokens) {
		tok := p.Tokens[p.Current]

		if tok.Type != tokenizer.OPERATOR || (tok.Value != "*" && tok.Value != "/") {
			break
		}

		operator := tok.Value
		p.Current++

		right := p.parseFactor()

		left = &ast.BinaryExpression{
			Left:     left,
			Operator: operator,
			Right:    right,
			Line:     tok.Line,
			Col:      tok.Col,
		}
	}

	return left
}

// Handling numbers, identifiers, and parentheses
func (p *Parser) parseFactor() ast.Expression {
	if p.Current >= len(p.Tokens) {
		return nil
	}

	tok := p.Tokens[p.Current]

	// handle numeric literals
	if tok.Type == tokenizer.NUMBER {
		p.Current++
		return &ast.NumberLiteral{Value: tok.Value, Line: tok.Line, Col: tok.Col}
	}

	// handle variable identifiers
	if tok.Type == tokenizer.IDENTIFIER {
		p.Current++
		return &ast.Identifier{Name: tok.Value, Line: tok.Line, Col: tok.Col}
	}

	// Handle parentheses `(expression)`
	if tok.Type == tokenizer.PUNCTUATION && tok.Value == "(" {
		p.Current++ // consume '('

		// recursively parse expression inside parentheses
		expr := p.parseExpression()

		if !p.match(tokenizer.PUNCTUATION, ")") {
			fmt.Printf("Syntax Error: Expected ')' at Line %d, Col %d\n", tok.Line, tok.Col)
			return nil
		}

		return expr
	}

	return nil
}
