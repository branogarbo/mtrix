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
package mult

import u "github.com/branogarbo/mtrix/util"

func ScalarMult(s float64, mat u.Matrix) u.Matrix {
	return u.PopulateNewMat(mat, func(mv u.MatrixValue, r, c int, secMvs ...u.MatrixValue) float64 {
		return mv[r][c] * s
	})
}

func MatMult(mats ...u.Matrix) u.Matrix {
	return u.Matrix{} // for now
}
