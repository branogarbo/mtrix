package power

import (
	"errors"

	m "github.com/branogarbo/mtrix/multiply"
	u "github.com/branogarbo/mtrix/util"
)

// MatPow returns m raised to the xth power.
func MatPow(mat u.Matrix, x int) (u.Matrix, error) {
	if x <= 0 {
		return u.Matrix{}, errors.New("invalid exponent passed")
	}

	var (
		resultMat = mat
		err       error
	)

	for i := 1; i < x; i++ {
		resultMat, err = m.MatMult(resultMat, mat)
		if err != nil {
			return u.Matrix{}, err
		}
	}

	return resultMat, nil
}
