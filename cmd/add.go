package cmd

import (
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
	RunE: func(cmd *cobra.Command, args []string) error {
		mats, err := util.ParseCmdArgs(cmd, args)
		if err != nil {
			return err
		}

		resultMat, err := add.MatAdd(mats...)
		if err != nil {
			return err
		}

		util.PrintMat(resultMat)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
