package sub

import (
	"github.com/branogarbo/mtrix/add"
	"github.com/branogarbo/mtrix/mult"
	u "github.com/branogarbo/mtrix/util"
)

func MatSub(m1, m2 u.Matrix) (u.Matrix, error) {
	negM2 := mult.ScalarMult(-1, m2)

	resultMat, err := add.MatAdd(m1, negM2)
	if err != nil {
		return u.Matrix{}, err
	}

	return resultMat, nil
}
