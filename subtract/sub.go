package subtract

import (
	a "github.com/branogarbo/mtrix/addition"
	m "github.com/branogarbo/mtrix/multiply"
	u "github.com/branogarbo/mtrix/util"
)

// MatSub subtracts m2 from m1.
func MatSub(m1, m2 u.Matrix) (u.Matrix, error) {
	negM2 := m.ScalarMult(-1, m2)

	resultMat, err := a.MatAdd(m1, negM2)
	if err != nil {
		return u.Matrix{}, err
	}

	return resultMat, nil
}
