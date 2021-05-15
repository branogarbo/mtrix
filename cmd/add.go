package cmd

import (
	"fmt"

	"github.com/branogarbo/mtrix/add"
	"github.com/branogarbo/mtrix/util"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"+"},
	Example: "mtrix add mat1.txt mat2.txt",
	Short:   "Get the sum of two matrices",
	Args:    cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		mats, err := util.GetMatsFromFiles(args...)
		if err != nil {
			fmt.Println(err)
			return
		}

		resultMat, err := add.MatAdd(mats...)
		if err != nil {
			fmt.Println(err)
			return
		}

		util.PrintMat(resultMat)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
