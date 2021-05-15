package cmd

import (
	"fmt"

	"github.com/branogarbo/mtrix/trans"
	"github.com/branogarbo/mtrix/util"
	"github.com/spf13/cobra"
)

var transCmd = &cobra.Command{
	Use:     "trans",
	Aliases: []string{"*"},
	Example: "mtrix trans mat.txt",
	Short:   "Get the transpose of a matrix",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mat, err := util.GetMatFromFile(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		resultMat := trans.MatTrans(mat)
		if err != nil {
			fmt.Println(err)
			return
		}

		util.PrintMat(resultMat)
	},
}

func init() {
	rootCmd.AddCommand(transCmd)
}
