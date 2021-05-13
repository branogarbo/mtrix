package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var multCmd = &cobra.Command{
	Use:     "mult",
	Aliases: []string{"*"},
	Example: "mtrix mult mat1.txt mat2.txt",
	Short:   "Multiply two matrices together.",
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mult called")
	},
}

func init() {
	rootCmd.AddCommand(multCmd)
}
