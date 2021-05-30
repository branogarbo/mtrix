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
	"reflect"
	"testing"

	u "github.com/branogarbo/mtrix/util"
)

func TestMatPow(t *testing.T) {
	type args struct {
		mat u.Matrix
		x   int
	}
	tests := []struct {
		name    string
		args    args
		want    u.Matrix
		wantErr bool
	}{
		{
			name: "working",
			args: args{
				mat: u.Matrix{
					Value: u.MatVal{
						{1, 2, -3.9},
						{4.3, 5, 6},
						{5, -3, 4},
					},
				},
				x: 2,
			},
			want: u.Matrix{
				RowsNum: 3,
				ColsNum: 3,
				Value: u.MatVal{
					{-9.9, 23.7, -7.5},
					{55.8, 15.600000000000001, 37.230000000000004},
					{12.100000000000001, -17, -21.5},
				},
			},
			wantErr: false,
		},
		{
			name: "size error",
			args: args{
				mat: u.Matrix{
					Value: u.MatVal{
						{1, 2, 3},
						{1, 2, 3},
					},
				},
			},
			want:    u.Matrix{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MatPow(tt.args.mat, tt.args.x)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatPow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
