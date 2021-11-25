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
package transpose

import u "github.com/branogarbo/mtrix/util"

// MatTrans returns the transpose of mat.
func MatTrans(mat u.Matrix) u.Matrix {
	MPconf := u.MatPopConfig{
		MainMat: mat,
		NewRows: mat.Cols(),
		NewCols: mat.Rows(),
		Action: func(mv u.Matrix, r, c int, secMvs []u.Matrix) float64 {
			return mv[c][r]
		},
	}

	return u.PopulateNewMat(MPconf)
}
