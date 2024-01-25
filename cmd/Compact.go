package cmd

import (
	"fmt"
	"compactor/Pkg"
	"github.com/spf13/cobra"
)

var compactor = &cobra.Command{
	Use:   "compact",
	Short: "You can check compactor",
	Long:  ``,
	Run: func(cmd *cobra.Command, args[]string) {
		var name, path string

		fmt.Println(Blue +"📁 File & Folder compactor", Reset + "\n")

		fmt.Print("File name: ")
		fmt.Scan(&name)

		fmt.Print("Directory path: ")
		fmt.Scan(&path)


		err := 	Pkg.Compact(name+ ".cmt", path)
	 	if err != nil {
			fmt.Println(Red + "🐞 Error :", err , Reset)
		}

		fmt.Println(Green + "✨ Process successful" + Reset)
	
	},
}

func init() {
	rootCmd.AddCommand(compactor)
}