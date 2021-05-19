package det

import (
	"errors"
	"math"

	u "github.com/branogarbo/mtrix/util"
)

// MatDet returns the determinant of m.
func MatDet(m u.Matrix) (float64, error) {
	var (
		det float64
		mv  = m.Value
	)

	if m.RowsNum != m.ColsNum {
		return 0, errors.New("argument is not a square matrix")
	}

	if m.RowsNum == 2 {
		return mv[0][0]*mv[1][1] - mv[0][1]*mv[1][0], nil
	}

	for c, elVal := range mv[0] {
		minor := u.GetMinor(m, 0, c)

		minorDet, err := MatDet(minor)
		if err != nil {
			return 0, err
		}

		det += math.Pow(-1, float64(c)) * elVal * minorDet
	}

	return det, nil
}
