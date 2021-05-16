package trans

import u "github.com/branogarbo/mtrix/util"

// MatTrans returns the transpose of mat.
func MatTrans(mat u.Matrix) u.Matrix {
	MPconf := u.MatPopConfig{
		MainMat: mat,
		NewRows: mat.ColsNum,
		NewCols: mat.RowsNum,
		Action: func(mv u.MatVal, r, c int, secMvs []u.MatVal) float64 {
			return mv[c][r]
		},
	}

	return u.PopulateNewMat(MPconf)
}
