package inverse

import (
	"errors"
	"math"

	det "github.com/branogarbo/mtrix/determinant"
	mult "github.com/branogarbo/mtrix/multiply"
	trans "github.com/branogarbo/mtrix/transpose"
	u "github.com/branogarbo/mtrix/util"
)

// MatInv returns the inverse of m.
func MatInv(m u.Matrix) (u.Matrix, error) {
	var (
		err       = m.SetSize()
		resultMat u.Matrix
		mv        = m.Value
	)
	if err != nil {
		return u.Matrix{}, err
	}

	if m.ColsNum == 2 && m.RowsNum == 2 {
		nm := u.Matrix{
			RowsNum: 2,
			ColsNum: 2,
			Value: u.MatVal{
				{mv[1][1], -mv[0][1]},
				{-mv[1][0], mv[0][0]},
			},
		}

		detM := mv[0][0]*mv[1][1] - mv[0][1]*mv[1][0]
		if detM == 0 {
			return u.Matrix{}, errors.New("matrix is singular, does not have an inverse")
		}

		return mult.ScalarMult(1/detM, nm), nil
	}

	// matrix of minors
	resultMat = u.PopulateNewMat(u.MatPopConfig{
		MainMat: m,
		NewRows: m.RowsNum,
		NewCols: m.ColsNum,
		Action: func(mv u.MatVal, r, c int, secMvs []u.MatVal) float64 {
			minor := u.GetMinor(m, r, c)
			detMinor, _ := det.MatDet(minor) // add error handling in the future

			return math.Pow(-1, float64(r+c)) * detMinor
		},
	})

	resultMat, err = trans.MatTrans(resultMat)
	if err != nil {
		return u.Matrix{}, err
	}

	detM, err := det.MatDet(m)
	if err != nil {
		return u.Matrix{}, err
	} else if detM == 0 {
		return u.Matrix{}, errors.New("matrix is singular, does not have an inverse")
	}

	resultMat = mult.ScalarMult(1/detM, resultMat)

	return resultMat, nil
}
