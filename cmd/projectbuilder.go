package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var projectBuildercmd = &cobra.Command{
	Use:   "new [folder-name] [user-name]",
	Short: "Creates a new folder",
	Args:  cobra.ExactArgs(2), // accept two args
	Run: func(cmd *cobra.Command, args []string) {
		// this run function handles the core logic of building CLI
		folderName := args[0]
		userName := args[1]
		fmt.Println(userName, args)
		err := os.MkdirAll(fmt.Sprintf("tmp/ %s", folderName), os.ModePerm)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
		fmt.Println("Folder created successfully:", folderName)
	},
}
