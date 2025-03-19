package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var projectBuildercmd = &cobra.Command{
	Use:   "new [folder-name] [user-name]",
	Short: "Creates a new folder",
	Args:  cobra.ExactArgs(2), // accept two args
	Run: func(cmd *cobra.Command, args []string) {
		// this run function handles the core logic of building CLI
		folderName := args[0]
		filePath := filepath.Join("tmp", folderName, "index.js")
		userName := args[1]
		fmt.Println(userName, args)
		err := os.MkdirAll(fmt.Sprintf("tmp/%s", folderName), os.ModePerm)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
		fmt.Println("Folder created successfully:", folderName)

		fmt.Println(filePath)
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		// wait to complete the execution of function then use defer to close file

		defer file.Close()

		jsCode := `console.log('Hello world this is js file created using golang cli tool');`
		_, err = file.WriteString(jsCode)

		if err != nil {
			fmt.Println("error writing file", err)
			return
		}
		fmt.Println("File created and content written successfully:", filePath)
	},
}
