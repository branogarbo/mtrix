package cmd

import (
	"github.com/branogarbo/mtrix/sub"
	"github.com/branogarbo/mtrix/util"
	"github.com/spf13/cobra"
)

var subCmd = &cobra.Command{
	Use:     "sub",
	Aliases: []string{"-", "s"},
	Example: "mtrix sub mat1.txt mat2.txt",
	Short:   "Get the difference of two matrices",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		mats, err := util.ParseCmdArgs(cmd, args)
		if err != nil {
			return err
		}

		resultMat, err := sub.MatSub(mats[0], mats[1])
		if err != nil {
			return err
		}

		util.PrintMat(resultMat)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(subCmd)
}
