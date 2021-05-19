package cmd

import (
	"github.com/branogarbo/mtrix/inv"
	"github.com/branogarbo/mtrix/util"
	"github.com/spf13/cobra"
)

var invCmd = &cobra.Command{
	Use:     "inv",
	Aliases: []string{"i"},
	Example: "mtrix inv mat.txt",
	Short:   "Get the inverse of a matrix",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		mats, err := util.ParseCmdArgs(cmd, args)
		if err != nil {
			return err
		}

		resultMat, err := inv.MatInv(mats[0])
		if err != nil {
			return err
		}

		util.PrintMat(resultMat)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(invCmd)
}
