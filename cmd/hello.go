package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// create the hello command

var HelloCommand = &cobra.Command{
	Use:   "Hello",
	Short: "Prints a greeting message",
	Long:  "This command prints a personalized greeting message.",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			name = "there"
		}
		fmt.Printf("Hello, %s! Welcome to MyCLI.\n", name)
	},
}

func init() {
	HelloCommand.Flags().StringP("name", "n", "", "Your name")
}
