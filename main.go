package main

import (
	"clitool/cmd"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "clitool",
		Short: "A basic cli tool in go",
		Long:  "This is a simple CLI tool built using Golang and Cobra.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to MyCLI! Use --help to see available commands.")
		},
	}
	// add the hello command
	rootCmd.AddCommand(cmd.HelloCommand)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
