/*
Copyright © 2025 Digvijaysinh Padhiyar
*/

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/digvijay-tech/Haru/internal/interpreter"
	"github.com/digvijay-tech/Haru/internal/parser"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <file>",
	Short: "Run a Haru script",
	Long:  `Interprets and executes Haru code from a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		visitor := interpreter.NewHaruVisitor()
		file, err := os.Open(args[0])

		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}

		defer file.Close()
		scanner := bufio.NewScanner(file)

		var code strings.Builder
		for scanner.Scan() {
			code.WriteString(scanner.Text() + "\n")
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		input := antlr.NewInputStream(code.String())
		runScript(input, visitor)
		fmt.Println("Vars:", visitor.Vars)
	},
}

func runScript(input *antlr.InputStream, visitor *interpreter.HaruVisitor) {
	lexer := parser.NewharuLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewharuParser(stream)
	tree := p.Program()
	visitor.VisitProgram(tree.(*parser.ProgramContext))
}

func init() {
	rootCmd.AddCommand(runCmd)
}
