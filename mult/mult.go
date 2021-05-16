package mult

import (
	"errors"

	u "github.com/branogarbo/mtrix/util"
)

// ScalarMult multiplies the matrix mat by the scalar s.
func ScalarMult(s float64, mat u.Matrix) u.Matrix {
	MPconf := u.MatPopConfig{
		MainMat: mat,
		NewRows: mat.RowsNum,
		NewCols: mat.ColsNum,
		Action: func(mv u.MatVal, r, c int, secMvs []u.MatVal) float64 {
			return mv[r][c] * s
		},
	}

	return u.PopulateNewMat(MPconf)
}

// MatMult multiplies two matrices together.
func MatMult(m1, m2 u.Matrix) (u.Matrix, error) {
	if !u.IsMultPossible(m1, m2) {
		return u.Matrix{}, errors.New("matrix multiplication not possible")
	}

	MPconf := u.MatPopConfig{
		MainMat: m1,
		NewRows: m1.RowsNum,
		NewCols: m2.ColsNum,
		Action: func(mv1 u.MatVal, r, c int, secMvs []u.MatVal) float64 {
			var (
				newEl   float64
				m1ElRow = m1.Value[r]
				m2ElCol []float64
			)

			for _, row := range m2.Value {
				m2ElCol = append(m2ElCol, row[c])
			}

			for i := 0; i < m1.ColsNum; i++ {
				newEl += m1ElRow[i] * m2ElCol[i]
			}

			return newEl
		},
	}

	resultMat := u.PopulateNewMat(MPconf)

	return resultMat, nil
}
