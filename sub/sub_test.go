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
package sub

import (
	"reflect"
	"testing"

	u "github.com/branogarbo/mtrix/util"
)

func TestMatSub(t *testing.T) {
	type args struct {
		m1 u.Matrix
		m2 u.Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    u.Matrix
		wantErr bool
	}{
		{
			name: "wanted",
			args: args{
				m1: u.Matrix{
					RowsNum: 3,
					ColsNum: 3,
					Value: u.MatrixValue{
						{3, 3, 3},
						{3, 3, 3},
						{3, 3, 3},
					},
				},
				m2: u.Matrix{
					RowsNum: 3,
					ColsNum: 3,
					Value: u.MatrixValue{
						{1, 1, 1},
						{1, 1, 1},
						{1, 1, 1},
					},
				},
			},
			want: u.Matrix{
				RowsNum: 3,
				ColsNum: 3,
				Value: u.MatrixValue{
					{2, 2, 2},
					{2, 2, 2},
					{2, 2, 2},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MatSub(tt.args.m1, tt.args.m2)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatSub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatSub() = %v, want %v", got, tt.want)
			}
		})
	}
}
