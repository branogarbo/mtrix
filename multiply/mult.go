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
package multiply

import (
	"errors"

	u "github.com/branogarbo/mtrix/util"
)

// ScalarMult multiplies the matrix mat by the scalar s.
func ScalarMult(s float64, mat u.Matrix) u.Matrix {
	MPconf := u.MatPopConfig{
		MainMat: mat,
		Action: func(mv u.Matrix, r, c int, secMvs []u.Matrix) float64 {
			return mv[r][c] * s
		},
	}

	return u.PopulateNewMat(MPconf)
}

// MatMult multiplies the passed matrices together.
func MatMult(mats ...u.Matrix) (u.Matrix, error) {
	var (
		resultMat = mats[0]
		err       error
	)

	for _, mat := range mats[1:] {
		resultMat, err = UnitMatMult(resultMat, mat)
		if err != nil {
			return u.Matrix{}, err
		}
	}

	return resultMat, nil
}

// UnitMatMult multiplies two matrices together.
func UnitMatMult(m1, m2 u.Matrix) (u.Matrix, error) {
	if !u.IsMultPossible(m1, m2) {
		return u.Matrix{}, errors.New("matrix multiplication not possible")
	}

	MPconf := u.MatPopConfig{
		MainMat: m1,
		NewRows: m1.Rows(),
		NewCols: m2.Cols(),
		Action: func(mv1 u.Matrix, r, c int, secMvs []u.Matrix) float64 {
			var (
				newEl   float64
				m1ElRow = m1[r]
				m2ElCol []float64
			)

			for _, row := range m2 {
				m2ElCol = append(m2ElCol, row[c])
			}

			for i := 0; i < m1.Cols(); i++ {
				newEl += m1ElRow[i] * m2ElCol[i]
			}

			return newEl
		},
	}

	resultMat := u.PopulateNewMat(MPconf)

	return resultMat, nil
}
