package transpose

import u "github.com/branogarbo/mtrix/util"

// MatTrans returns the transpose of mat.
func MatTrans(mat u.Matrix) (u.Matrix, error) {
	err := mat.SetSize()
	if err != nil {
		return u.Matrix{}, err
	}

	MPconf := u.MatPopConfig{
		MainMat: mat,
		NewRows: mat.ColsNum,
		NewCols: mat.RowsNum,
		Action: func(mv u.MatVal, r, c int, secMvs []u.MatVal) float64 {
			return mv[c][r]
		},
	}

	return u.PopulateNewMat(MPconf), nil
}
