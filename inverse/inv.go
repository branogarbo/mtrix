/*
Copyright © 2021 Brian Longmore branodev@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
		resultMat u.Matrix
		err       error
	)

	if m.Cols() == 2 && m.Rows() == 2 {
		nm := u.Matrix{
			{m[1][1], -m[0][1]},
			{-m[1][0], m[0][0]},
		}

		detM := m[0][0]*m[1][1] - m[0][1]*m[1][0]
		if detM == 0 {
			return u.Matrix{}, errors.New("matrix is singular, does not have an inverse")
		}

		return mult.ScalarMult(1/detM, nm), nil
	}

	// matrix of minors
	resultMat = u.PopulateNewMat(u.MatPopConfig{
		MainMat: m,
		Action: func(m u.Matrix, r, c int, secMs []u.Matrix) float64 {
			minor := u.GetMinor(m, r, c)
			detMinor, _ := det.MatDet(minor) // add error handling in the future

			return math.Pow(-1, float64(r+c)) * detMinor
		},
	})

	resultMat = trans.MatTrans(resultMat)

	detM, err := det.MatDet(m)
	if err != nil {
		return u.Matrix{}, err
	} else if detM == 0 {
		return u.Matrix{}, errors.New("matrix is singular, does not have an inverse")
	}

	resultMat = mult.ScalarMult(1/detM, resultMat)

	return resultMat, nil
}
