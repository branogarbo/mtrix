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
package trans

import (
	"reflect"
	"testing"

	u "github.com/branogarbo/mtrix/util"
)

func TestMatTrans(t *testing.T) {
	type args struct {
		mat u.Matrix
	}
	tests := []struct {
		name string
		args args
		want u.Matrix
	}{
		{
			name: "transpose test",
			args: args{
				mat: u.Matrix{
					RowsNum: 4,
					ColsNum: 4,
					Value: u.MatVal{
						{1, 2, 3, 4},
						{5, 6, 7, 8},
						{9, 10, 11, 12},
						{13, 14, 15, 16},
					},
				},
			},
			want: u.Matrix{
				RowsNum: 4,
				ColsNum: 4,
				Value: u.MatVal{
					{1, 5, 9, 13},
					{2, 6, 10, 14},
					{3, 7, 11, 15},
					{4, 8, 12, 16},
				},
			},
		},
		{
			name: "mismatch size trans",
			args: args{
				mat: u.Matrix{
					RowsNum: 2,
					ColsNum: 3,
					Value: u.MatVal{
						{1, 2, 3},
						{5, 6, 7},
					},
				},
			},
			want: u.Matrix{
				RowsNum: 3,
				ColsNum: 2,
				Value: u.MatVal{
					{1, 5},
					{2, 6},
					{3, 7},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MatTrans(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatTrans() = %v, want %v", got, tt.want)
			}
		})
	}
}
