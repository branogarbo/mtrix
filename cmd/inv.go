package cmd

import (
	i "github.com/branogarbo/mtrix/inverse"
	u "github.com/branogarbo/mtrix/util"
	"github.com/spf13/cobra"
)

var invCmd = &cobra.Command{
	Use:     "inv",
	Aliases: []string{"i"},
	Example: "mtrix inv mat.txt",
	Short:   "Get the inverse of a matrix",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		mats, err := u.ParseCmdArgs(cmd, args)
		if err != nil {
			return err
		}

		resultMat, err := i.MatInv(mats[0])
		if err != nil {
			return err
		}

		u.PrintMat(resultMat)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(invCmd)
}
