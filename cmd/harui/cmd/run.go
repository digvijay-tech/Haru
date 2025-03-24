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
	"github.com/digvijay-tech/Haru/internal/interpreter/listener"
	"github.com/digvijay-tech/Haru/internal/parser"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a Haru script",
	Long:  `Interprets and executes Haru code provided as a string or file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Haru REPL - type 'exit' to quit")
		listener := listener.NewHaruListener()
		scanner := bufio.NewScanner(os.Stdin)
		for {
			fmt.Print("> ")

			if !scanner.Scan() {
				break
			}

			line := scanner.Text()
			if strings.TrimSpace(line) == "exit" {
				break
			}

			if !strings.HasSuffix(line, ";") {
				line += ";"
			}

			input := antlr.NewInputStream(line)
			lexer := parser.NewharuLexer(input)

			stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

			p := parser.NewharuParser(stream)
			tree := p.Program()

			antlr.ParseTreeWalkerDefault.Walk(listener, tree)
			fmt.Println("Vars:", listener.Vars)
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

type HaruListener struct {
	*parser.BaseharuListener
}

func init() {
	rootCmd.AddCommand(runCmd)
}
