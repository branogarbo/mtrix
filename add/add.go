package add

import (
	"errors"

	u "github.com/branogarbo/mtrix/util"
)

// MatAdd adds matrices together.
func MatAdd(mats ...u.Matrix) (u.Matrix, error) {
	var (
		resultMat u.Matrix
		err       error
	)

	if len(mats) < 2 {
		return u.Matrix{}, errors.New("less than 2 args passed")
	}

	err = u.CheckMatSizes(mats...)
	if err != nil {
		return u.Matrix{}, err
	}

	MPconf := u.MatPopConfig{
		MainMat: mats[0],
		SecMats: mats[1:],
		NewRows: mats[0].RowsNum,
		NewCols: mats[0].ColsNum,
		Action: func(mv u.MatVal, r, c int, secMats []u.MatVal) float64 {
			elSum := mv[r][c]

			for _, secMv := range secMats {
				elSum += secMv[r][c]
			}

			return elSum
		},
	}

	resultMat = u.PopulateNewMat(MPconf)

	return resultMat, nil
}
