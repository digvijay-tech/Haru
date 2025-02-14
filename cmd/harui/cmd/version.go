/*
Copyright © 2025 Digvijaysinh Padhiyar
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const version = "0.1.0"

// Version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Haru",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Haru Interpreter v" + version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
