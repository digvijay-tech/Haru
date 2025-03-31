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

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "harui",
	Short: "Haru Interpreter",
	Long:  `Harui is the command-line tool for running and managing Haru scripts.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Haru REPL - type 'exit' to quit")
		visitor := interpreter.NewHaruVisitor()
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

			if !strings.HasSuffix(line, ";") && !strings.HasSuffix(line, "}") {
				line += ";"
			}

			input := antlr.NewInputStream(line)
			lexer := parser.NewharuLexer(input)

			stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
			p := parser.NewharuParser(stream)

			tree := p.Program()

			visitor.Visit(tree)
			fmt.Println("Vars:", visitor.Vars)
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error:", err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.harui.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
