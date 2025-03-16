package tests

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/digvijay-tech/Haru/internal/lexer"
	"github.com/digvijay-tech/Haru/internal/lexer/tokenizer"
	"github.com/digvijay-tech/Haru/internal/parser"
	"github.com/digvijay-tech/Haru/internal/parser/ast"
	"github.com/digvijay-tech/Haru/internal/preprocessor"
)

func formatExpression(expr ast.Expression) string {
	if expr == nil {
		return "nil"
	}

	switch v := expr.(type) {
	case *ast.NumberLiteral:
		return v.Value
	case *ast.Identifier:
		return v.Name
	case *ast.BinaryExpression:
		// Recursively format binary expressions
		return fmt.Sprintf("(%s %s %s)", formatExpression(v.Left), v.Operator, formatExpression(v.Right))
	default:
		return "Unknown Expression"
	}
}

func TestParsingVarStatements(t *testing.T) {
	// loading source code file
	data, err := os.ReadFile("./temp/VariableParsing/test.haru")
	if err != nil {
		fmt.Printf("Error Reading File: \n%s", err)
		os.Exit(1)
	}

	// preprocessing
	processedSource := preprocessor.PreProcess(string(data))

	// lexing
	lex := lexer.NewLexer(processedSource)
	tokens := []tokenizer.Token{}

	// collecting all tokens
	for {
		tok := lex.NextToken()
		tokens = append(tokens, tok)
		if tok.Type == tokenizer.EOF {
			break
		}
	}

	// parsing
	parser := &parser.Parser{Tokens: tokens, Current: 0}
	var builder strings.Builder

	for {
		statement := parser.ParseLetAndMutStatement()
		if statement == nil {
			// ensuring the loop doesn't exit too early
			if tokens[parser.Current].Type == tokenizer.EOF {
				break
			}

			continue
		}

		exprStr := formatExpression(statement.Value)

		row := fmt.Sprintf("Identifier: %s, Type: %s, Value: %s, Mutable: %t, Line: %d, Col: %d\n",
			statement.Identifier, statement.Type, exprStr, statement.Mutable, statement.Line, statement.Col)

		builder.WriteString(row)

		fmt.Printf("%s", row)

		// stop only when EOF is actually reached
		if tokens[parser.Current].Type == tokenizer.EOF {
			break
		}
	}

	// writing the buffered output
	writeErr := os.WriteFile("./temp/VariableParsing/result.log", []byte(builder.String()), 0666)

	if writeErr != nil {
		fmt.Printf("Error Writing Results: \n%s", err)
		os.Exit(1)
	}
}
