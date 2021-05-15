package cmd

import (
	"fmt"

	"github.com/branogarbo/mtrix/sub"
	"github.com/branogarbo/mtrix/util"
	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:     "sub",
	Aliases: []string{"-"},
	Example: "mtrix sub mat1.txt mat2.txt",
	Short:   "Get the difference of two matrices",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		mats, err := util.GetMatsFromFiles(args...)
		if err != nil {
			fmt.Println(err)
			return
		}

		resultMat, err := sub.MatSub(mats[0], mats[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		util.PrintMat(resultMat)
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}
