package main

import (
	"fmt"
	"os"
	"time"

	"github.com/antlr4-go/antlr/v4"
	"github.com/digvijay-tech/Haru/internal/interpreter"
	"github.com/digvijay-tech/Haru/internal/parser"
	"github.com/digvijay-tech/Haru/internal/preprocessor"
)

func main() {
	filePath := os.Args[1]
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Preprocess the code (remove comments, trim whitespace etc.)
	code := preprocessor.PreProcess(string(content))

	// Setup ANTLR input stream and lexer
	input := antlr.NewInputStream(code)
	lexer := parser.NewharuLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewharuParser(stream)

	// Catch parse errors
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	// Parse program and walk tree
	tree := p.Program()
	visitor := interpreter.NewHaruVisitor()

	start := time.Now()

	visitor.Visit(tree)

	elapsed := time.Since(start)
	fmt.Printf("Haru took: %s\n", elapsed)
}
