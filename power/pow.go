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
