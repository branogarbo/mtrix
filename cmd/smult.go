package cmd

import (
	"fmt"
	"strconv"

	"github.com/branogarbo/mtrix/mult"
	"github.com/branogarbo/mtrix/util"
	"github.com/spf13/cobra"
)

var smultCmd = &cobra.Command{
	Use:     "smult",
	Aliases: []string{"sm", "mults"},
	Example: "mtrix smult 2 mat.txt",
	Short:   "Multiply a matrix by a scalar",
	Args:    cobra.ExactArgs(2), //cobra.MinimumNArgs(2) later
	Run: func(cmd *cobra.Command, args []string) {
		var (
			mat       util.Matrix
			matArg    = args[1]
			scal, err = strconv.ParseFloat(args[0], 64)
		)
		if err != nil {
			fmt.Println(err)
			return
		}

		isRaw, err := cmd.Flags().GetBool("raw-input")
		if err != nil {
			fmt.Println(err)
			return
		}

		if isRaw {
			mat, err = util.StringToMat(matArg)
		} else {
			mat, err = util.GetMatFromFile(matArg)
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		resultMat := mult.ScalarMult(scal, mat)

		util.PrintMat(resultMat)
	},
}

func init() {
	rootCmd.AddCommand(smultCmd)
}
