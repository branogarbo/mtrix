package cmd

import (
	t "github.com/branogarbo/mtrix/transpose"
	u "github.com/branogarbo/mtrix/util"
	"github.com/spf13/cobra"
)

var transCmd = &cobra.Command{
	Use:     "trans",
	Aliases: []string{"t"},
	Example: "mtrix trans mat.txt",
	Short:   "Get the transpose of a matrix",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		mats, err := u.ParseCmdArgs(cmd, args)
		if err != nil {
			return err
		}

		resultMat, err := t.MatTrans(mats[0])
		if err != nil {
			return err
		}

		u.PrintMat(resultMat)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(transCmd)
}
