package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var projectBuildercmd = &cobra.Command{
	Use:   "new [folder-name]",
	Short: "Creates a new folder",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		folderName := args[0]
		err := os.MkdirAll(fmt.Sprintf("tmp/ %s", folderName), os.ModePerm)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
		fmt.Println("Folder created successfully:", folderName)
	},
}
