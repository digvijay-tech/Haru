/*
Copyright © 2025 Digvijaysinh Padhiyar
*/

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run <file>",
	Short: "Run a Haru script",
	Long:  `Interprets and executes Haru code from a file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Haru Interpreter!")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
}
