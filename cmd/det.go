package cmd

import (
	"fmt"

	"github.com/branogarbo/mtrix/det"
	"github.com/branogarbo/mtrix/util"
	"github.com/spf13/cobra"
)

var detCmd = &cobra.Command{
	Use:     "det",
	Aliases: []string{"d"},
	Example: "mtrix det mat.txt",
	Short:   "Compute the determinant of a matrix",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		mats, err := util.ParseCmdArgs(cmd, args)
		if err != nil {
			return err
		}

		det, err := det.MatDet(mats[0])
		if err != nil {
			return err
		}

		fmt.Println(det)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(detCmd)
}
