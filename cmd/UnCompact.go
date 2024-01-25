package cmd

import (
	"fmt"
	"compactor/Pkg"
	"github.com/spf13/cobra"
)

var uncompactor = &cobra.Command{
	Use:   "uncompact",
	Short: "You can check uncompactor",
	Long:  ``,
	Run: func(cmd *cobra.Command, args[]string) {
		var name, path string

		fmt.Println(Blue +"📁 File & Folder uncompactor", Reset + "\n")

		fmt.Print("File name: ")
		fmt.Scan(&name)

		fmt.Print("Directory path: ")
		fmt.Scan(&path)

		err := 	Pkg.UnCompact(name + ".cmt", path)
	 	if err != nil {
			fmt.Println(Red + "🐞 Error :", err , Reset)
		}

		fmt.Println(Green + "✨ Process successful" + Reset)
	
	},
}

func init() {
	rootCmd.AddCommand(uncompactor)
}