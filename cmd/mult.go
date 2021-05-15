package cmd

import (
	"fmt"

	"github.com/branogarbo/mtrix/mult"
	"github.com/branogarbo/mtrix/util"
	"github.com/spf13/cobra"
)

var multCmd = &cobra.Command{
	Use:     "mult",
	Aliases: []string{"*"},
	Example: "mtrix mult mat1.txt mat2.txt",
	Short:   "Multiply two matrices together",
	Args:    cobra.ExactArgs(2), //cobra.MinimumNArgs(2) later
	Run: func(cmd *cobra.Command, args []string) {
		mats, err := util.GetMatsFromFiles(args...)
		if err != nil {
			fmt.Println(err)
			return
		}

		resultMat, err := mult.MatMult(mats[0], mats[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		util.PrintMat(resultMat)
	},
}

func init() {
	rootCmd.AddCommand(multCmd)
}
