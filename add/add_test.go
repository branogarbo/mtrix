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
package add

import (
	"reflect"
	"testing"

	u "github.com/branogarbo/mtrix/util"
)

func TestMatAdd(t *testing.T) {
	type args struct {
		mats []u.Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    u.Matrix
		wantErr bool
	}{
		{
			name: "add matrices",
			args: args{
				mats: []u.Matrix{
					{
						RowsNum: 3,
						ColsNum: 3,
						Value: u.MatrixValue{
							{1, 2, 3},
							{4, 5, 6},
							{7, 8, 9},
						},
					},
					{
						RowsNum: 3,
						ColsNum: 3,
						Value: u.MatrixValue{
							{1, 1, 1},
							{1, 1, 1},
							{1, 1, 1},
						},
					},
				},
			},
			want: u.Matrix{
				RowsNum: 3,
				ColsNum: 3,
				Value: u.MatrixValue{
					{2, 3, 4},
					{5, 6, 7},
					{8, 9, 10},
				},
			},
			wantErr: false,
		},
		{
			name: "nonmatching sizes",
			args: args{
				mats: []u.Matrix{
					{
						RowsNum: 3,
						ColsNum: 3,
						Value: u.MatrixValue{
							{1, 2, 3},
							{4, 5, 6},
							{7, 8, 9},
						},
					},
					{
						RowsNum: 2,
						ColsNum: 2,
						Value: u.MatrixValue{
							{1, 1},
							{1, 1},
						},
					},
					{
						RowsNum: 2,
						ColsNum: 3,
						Value: u.MatrixValue{
							{1, 1, 1},
							{1, 1, 1},
						},
					},
				},
			},
			want:    u.Matrix{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MatAdd(tt.args.mats...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatAdd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
