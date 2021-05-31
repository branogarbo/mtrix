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
package subtract

import (
	a "github.com/branogarbo/mtrix/addition"
	m "github.com/branogarbo/mtrix/multiply"
	u "github.com/branogarbo/mtrix/util"
)

// MatSub subtracts m2 from m1.
func MatSub(m1, m2 u.Matrix) (u.Matrix, error) {
	negM2 := m.ScalarMult(-1, m2)

	resultMat, err := a.MatAdd(m1, negM2)
	if err != nil {
		return u.Matrix{}, err
	}

	return resultMat, nil
}
