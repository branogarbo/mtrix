package cmd

import (
	"strconv"

	p "github.com/branogarbo/mtrix/power"
	u "github.com/branogarbo/mtrix/util"
	"github.com/spf13/cobra"
)

var powCmd = &cobra.Command{
	Use:     "pow",
	Aliases: []string{"p", "power"},
	Example: "mtrix pow mat.txt 3",
	Short:   "Raise a matrix to the nth power",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		exp, err := strconv.Atoi(args[1])
		if err != nil {
			return err
		}

		mats, err := u.ParseCmdArgs(cmd, args[:1])
		if err != nil {
			return err
		}

		resultMat, err := p.MatPow(mats[0], exp)
		if err != nil {
			return err
		}

		u.PrintMat(resultMat)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(powCmd)
}
