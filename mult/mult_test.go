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

import (
	"reflect"
	"testing"

	u "github.com/branogarbo/mtrix/util"
)

func TestScalarMult(t *testing.T) {
	type args struct {
		s   float64
		mat u.Matrix
	}
	tests := []struct {
		name string
		args args
		want u.Matrix
	}{
		{
			name: "wanted",
			args: args{
				s: -2,
				mat: u.Matrix{
					RowsNum: 2,
					ColsNum: 2,
					Value: u.MatrixValue{
						{1, 1},
						{1, 1},
					},
				},
			},
			want: u.Matrix{
				RowsNum: 2,
				ColsNum: 2,
				Value: u.MatrixValue{
					{-2, -2},
					{-2, -2},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ScalarMult(tt.args.s, tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ScalarMult() = %v, want %v", got, tt.want)
			}
		})
	}
}
