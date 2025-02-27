package tests

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/digvijay-tech/Haru/internal/lexer"
	"github.com/digvijay-tech/Haru/internal/lexer/tokenizer"
	"github.com/digvijay-tech/Haru/internal/preprocessor"
)

func TestLexing(t *testing.T) {
	// loading source code file
	data, err := os.ReadFile("./temp/final_test.haru")

	if err != nil {
		fmt.Printf("Error Reading File: \n%s", err)
		os.Exit(1)
	}

	// preprocessing for removing comments and non-significant whitespaces
	source := preprocessor.PreProcess(string(data))

	fmt.Println("--------------- BEFORE ---------------")
	fmt.Println(source)
	fmt.Println("--------------- AFTER ---------------")

	// lexing the processed source code
	lexer := lexer.NewLexer(source)

	// storing output
	var builder strings.Builder

	for {
		tok := lexer.NextToken()
		row := fmt.Sprintf("%s: '%s' (Line: %d, Column: %d)\n", tok.Type, tok.Value, tok.Line, tok.Col)

		// appends to builder
		builder.WriteString(row)

		fmt.Printf("%s", row)

		if tok.Type == tokenizer.EOF {
			break
		}
	}

	// write whole buffered output at once
	writeErr := os.WriteFile("./temp/result.log", []byte(builder.String()), 0666)

	if writeErr != nil {
		fmt.Printf("Error Writing Results: \n%s", err)
		os.Exit(1)
	}
}
