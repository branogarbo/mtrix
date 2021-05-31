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
package transpose

import u "github.com/branogarbo/mtrix/util"

// MatTrans returns the transpose of mat.
func MatTrans(mat u.Matrix) (u.Matrix, error) {
	err := mat.SetSize()
	if err != nil {
		return u.Matrix{}, err
	}

	MPconf := u.MatPopConfig{
		MainMat: mat,
		NewRows: mat.ColsNum,
		NewCols: mat.RowsNum,
		Action: func(mv u.MatVal, r, c int, secMvs []u.MatVal) float64 {
			return mv[c][r]
		},
	}

	return u.PopulateNewMat(MPconf), nil
}
