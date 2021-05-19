package cmd

import (
	"strconv"

	m "github.com/branogarbo/mtrix/multiply"
	u "github.com/branogarbo/mtrix/util"
	"github.com/spf13/cobra"
)

var smultCmd = &cobra.Command{
	Use:     "smult",
	Aliases: []string{"sm", "ms", "mults"},
	Example: "mtrix smult 2 mat.txt",
	Short:   "Multiply a matrix by a scalar",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		scal, err := strconv.ParseFloat(args[0], 64)
		if err != nil {
			return err
		}

		mats, err := u.ParseCmdArgs(cmd, args[1:])
		if err != nil {
			return err
		}

		resultMat := m.ScalarMult(scal, mats[0])

		u.PrintMat(resultMat)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(smultCmd)
}
