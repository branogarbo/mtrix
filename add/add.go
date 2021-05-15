package add

import u "github.com/branogarbo/mtrix/util"

func MatAdd(mats ...u.Matrix) (u.Matrix, error) {
	var (
		resultMat u.Matrix
		err       error
	)

	err = u.CheckMatsSizes(mats...)
	if err != nil {
		return u.Matrix{}, err
	}

	resultMat = u.PopulateNewMat(mats[0], func(mv u.MatrixValue, r, c int, secMats ...u.MatrixValue) float64 {
		elSum := mv[r][c]

		for _, secMv := range secMats {
			elSum += secMv[r][c]
		}

		return elSum
	}, mats[1:]...)

	return resultMat, nil
}
