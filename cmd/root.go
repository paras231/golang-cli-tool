package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd is the base command for the CLI
var rootCmd = &cobra.Command{
	Use:   "clitool",
	Short: "A basic CLI tool in Go",
	Long:  "This is a simple CLI tool built using Golang and Cobra.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to MyCLI! Use --help to see available commands.")
	},
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Add subcommands in init()
func init() {
	rootCmd.AddCommand(projectBuildercmd) // Add HelloCommand here
}
