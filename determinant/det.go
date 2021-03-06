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
package determinant

import (
	"errors"
	"math"

	u "github.com/branogarbo/mtrix/util"
)

// MatDet returns the determinant of m.
func MatDet(m u.Matrix) (float64, error) {
	var det float64

	if m.Rows() != m.Cols() {
		return 0, errors.New("argument is not a square matrix")
	}

	if m.Rows() == 2 {
		return m[0][0]*m[1][1] - m[0][1]*m[1][0], nil
	}

	for c, elVal := range m[0] {
		minor := u.GetMinor(m, 0, c)

		minorDet, err := MatDet(minor)
		if err != nil {
			return 0, err
		}

		det += math.Pow(-1, float64(c)) * elVal * minorDet
	}

	return det, nil
}
