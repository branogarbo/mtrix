package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:     "sub",
	Aliases: []string{"-"},
	Example: "mtrix sub mat1.txt mat2.txt",
	Short:   "Get the difference between two matrices.",
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sub called")
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}
