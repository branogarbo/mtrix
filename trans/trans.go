package trans

import u "github.com/branogarbo/mtrix/util"

func MatTrans(mat u.Matrix) u.Matrix {
	return u.PopulateNewMat(mat, func(mv u.MatrixValue, r, c int, secMvs ...u.MatrixValue) float64 {
		return mv[c][r]
	})
}
