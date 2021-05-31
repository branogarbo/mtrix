package cmd

import (
	a "github.com/branogarbo/mtrix/addition"
	u "github.com/branogarbo/mtrix/util"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"+", "a"},
	Example: "mtrix add mat1.txt mat2.txt mat3.txt",
	Short:   "Get the sum of matrices",
	Args:    cobra.MinimumNArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		mats, err := u.ParseCmdArgs(cmd, args)
		if err != nil {
			return err
		}

		resultMat, err := a.MatAdd(mats...)
		if err != nil {
			return err
		}

		u.PrintMat(resultMat)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
