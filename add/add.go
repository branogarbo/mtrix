/*
Copyright © 2021 Brian Longmore brianl.ext@gmail.com

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
package util

import u "github.com/branogarbo/mtrix/util"

func Add(mats ...u.Matrix) (u.Matrix, error) {
	var (
		err       error
		resultMat u.Matrix
		newMatVal u.MatrixValue
		argMvs    []u.MatrixValue
	)

	err = u.CheckMats(mats...)
	if err != nil {
		return u.Matrix{}, err
	}
	err = u.CheckMatsSizes(mats...)
	if err != nil {
		return u.Matrix{}, err
	}

	for _, mat := range mats {
		argMvs = append(argMvs, mat.Value)
	}

	newMatVal = u.PopulateNewMat(argMvs[0], func(mv u.MatrixValue, r, c int, secMvs ...u.MatrixValue) float64 {
		elSum := mv[r][c]

		for _, secMv := range secMvs {
			elSum += secMv[r][c]
		}

		return elSum
	}, argMvs[1:]...)

	resultMat = u.Matrix{
		RowsNum: len(newMatVal),
		ColsNum: len(newMatVal[0]),
		Value:   newMatVal,
	}

	return resultMat, nil
}
