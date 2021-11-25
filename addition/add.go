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
package addition

import (
	"errors"

	u "github.com/branogarbo/mtrix/util"
)

// MatAdd adds matrices together.
func MatAdd(mats ...u.Matrix) (u.Matrix, error) {
	var (
		resultMat u.Matrix
		err       error
	)

	if len(mats) < 2 {
		return u.Matrix{}, errors.New("less than 2 args passed")
	}

	err = u.CheckMatSizes(mats...)
	if err != nil {
		return u.Matrix{}, err
	}

	MPconf := u.MatPopConfig{
		MainMat: mats[0],
		SecMats: mats[1:],
		Action: func(m u.Matrix, r, c int, secMats []u.Matrix) float64 {
			elSum := m[r][c]

			for _, secMv := range secMats {
				elSum += secMv[r][c]
			}

			return elSum
		},
	}

	resultMat = u.PopulateNewMat(MPconf)

	return resultMat, nil
}
