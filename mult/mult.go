package mult

import (
	"errors"

	u "github.com/branogarbo/mtrix/util"
)

func ScalarMult(s float64, mat u.Matrix) u.Matrix {
	return u.PopulateNewMat(mat, func(mv u.MatrixValue, r, c int, secMvs ...u.MatrixValue) float64 {
		return mv[r][c] * s
	})
}

func MatMult(m1, m2 u.Matrix) (u.Matrix, error) {
	if !u.IsMultPossible(m1, m2) {
		return u.Matrix{}, errors.New("matrix multiplication not possible")
	}

	resultMat := u.PopulateNewMat(m1, func(mv1 u.MatrixValue, r, c int, secMvs ...u.MatrixValue) float64 {
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
	})

	return resultMat, nil
}
