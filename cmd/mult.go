package cmd

import (
	m "github.com/branogarbo/mtrix/multiply"
	u "github.com/branogarbo/mtrix/util"
	"github.com/spf13/cobra"
)

var multCmd = &cobra.Command{
	Use:     "mult",
	Aliases: []string{"x", "m"},
	Example: "mtrix mult mat1.txt mat2.txt",
	Short:   "Multiply two matrices together",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		mats, err := u.ParseCmdArgs(cmd, args)
		if err != nil {
			return err
		}

		resultMat, err := m.MatMult(mats[0], mats[1])
		if err != nil {
			return err
		}

		u.PrintMat(resultMat)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(multCmd)
}
