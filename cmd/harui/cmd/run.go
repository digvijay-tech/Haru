/*
Copyright © 2025 Digvijaysinh Padhiyar
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/digvijay-tech/Haru/internal/preprocessor"
	"github.com/spf13/cobra"
)

// Run command
var rumCmd = &cobra.Command{
	Use:   "run <file>",
	Short: "Run a Haru source file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filename := args[0]
		data, err := os.ReadFile(filename)

		if err != nil {
			fmt.Println("Error: Unable to read file:", err)
			os.Exit(1)
		}

		// Preprocessing source code
		processedCode := preprocessor.PreProcess(string(data))

		fmt.Println("Processed Code:")
		fmt.Println(processedCode)
	},
}

func init() {
	rootCmd.AddCommand(rumCmd)
}
