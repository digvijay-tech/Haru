/*
Copyright © 2025 Digvijaysinh Padhiyar
*/

package cmd

import (
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
	"github.com/digvijay-tech/Haru/internal/interpreter"
	"github.com/digvijay-tech/Haru/internal/parser"
	"github.com/digvijay-tech/Haru/internal/preprocessor"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <file>",
	Short: "Run a Haru script",
	Long:  `Interprets and executes Haru code from a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Usage: haru run <file.haru>")
			return
		}

		filePath := args[0] // use args[0], not os.Args[1]
		content, err := os.ReadFile(filePath)
		if err != nil {
			fmt.Printf("Error reading file: %v\n", err)
			return
		}

		code := preprocessor.PreProcess(string(content))

		input := antlr.NewInputStream(code)
		lexer := parser.NewharuLexer(input)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewharuParser(stream)

		tree := p.Program()
		visitor := interpreter.NewHaruVisitor()

		visitor.Visit(tree)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
