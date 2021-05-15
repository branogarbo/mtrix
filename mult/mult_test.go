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

func TestMatMult(t *testing.T) {
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
			name: "no error",
			args: args{
				m1: u.Matrix{
					RowsNum: 2,
					ColsNum: 2,
					Value: u.MatrixValue{
						{2, 3},
						{3, -5},
					},
				},
				m2: u.Matrix{
					RowsNum: 2,
					ColsNum: 2,
					Value: u.MatrixValue{
						{1, 0},
						{0, 1},
					},
				},
			},
			want: u.Matrix{
				RowsNum: 2,
				ColsNum: 2,
				Value: u.MatrixValue{
					{2, 3},
					{3, -5},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MatMult(tt.args.m1, tt.args.m2)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatMult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatMult() = %v, want %v", got, tt.want)
			}
		})
	}
}
