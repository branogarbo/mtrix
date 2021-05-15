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
package util

import (
	"reflect"
	"testing"
)

func TestIsMultPossible(t *testing.T) {
	type args struct {
		m1 Matrix
		m2 Matrix
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "not possible",
			args: args{
				m1: Matrix{
					RowsNum: 3,
					ColsNum: 3,
					Value: MatrixValue{
						{1, 1, 1},
						{1, 1, 1},
						{1, 1, 1},
					},
				},
				m2: Matrix{
					RowsNum: 2,
					ColsNum: 3,
					Value: MatrixValue{
						{1, 1, 1},
						{1, 1, 1},
					},
				},
			},
			want: false,
		},
		{
			name: "possible",
			args: args{
				m1: Matrix{
					RowsNum: 3,
					ColsNum: 3,
					Value: MatrixValue{
						{1, 1, 1},
						{1, 1, 1},
						{1, 1, 1},
					},
				},
				m2: Matrix{
					RowsNum: 3,
					ColsNum: 4,
					Value: MatrixValue{
						{1, 1, 1, 1},
						{1, 1, 1, 1},
						{1, 1, 1, 1},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMultPossible(tt.args.m1, tt.args.m2); got != tt.want {
				t.Errorf("IsMultPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMatrixValid(t *testing.T) {
	type args struct {
		m Matrix
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "invalid",
			args: args{
				m: Matrix{
					RowsNum: 3,
					ColsNum: 4,
					Value: MatrixValue{
						{1, 1, 1, 1},
						{1, 1, 1, 1},
						{1, 1, 1},
					},
				},
			},
			want: false,
		},
		{
			name: "valid",
			args: args{
				m: Matrix{
					RowsNum: 3,
					ColsNum: 4,
					Value: MatrixValue{
						{1, 1, 1, 1},
						{1, 1, 1, 1},
						{1, 1, 1, 1},
					},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMatrixValid(tt.args.m); got != tt.want {
				t.Errorf("IsMatrixValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckMatsSizes(t *testing.T) {
	type args struct {
		mats []Matrix
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no error",
			args: args{
				mats: []Matrix{
					{
						RowsNum: 3,
						ColsNum: 4,
					},
					{
						RowsNum: 3,
						ColsNum: 4,
					},
					{
						RowsNum: 3,
						ColsNum: 4,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				mats: []Matrix{
					{
						RowsNum: 3,
						ColsNum: 3,
					},
					{
						RowsNum: 2,
						ColsNum: 2,
					},
					{
						RowsNum: 2,
						ColsNum: 2,
					},
					{
						RowsNum: 3,
						ColsNum: 7,
					},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid args",
			args: args{
				mats: []Matrix{
					{
						RowsNum: 1,
						ColsNum: 3,
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckMatsSizes(tt.args.mats...); (err != nil) != tt.wantErr {
				t.Errorf("CheckMatsSizes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetMatsFromFiles(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name    string
		args    args
		want    []Matrix
		wantErr bool
	}{
		{
			name: "parse matrix files test",
			args: args{
				paths: []string{"../sampleMats/2square1.txt", "../sampleMats/3square2.txt", "../sampleMats/2x3_1.txt"},
			},
			want: []Matrix{
				{
					RowsNum: 2,
					ColsNum: 2,
					Value: MatrixValue{
						{1, 2},
						{3, -5},
					},
				},
				{
					RowsNum: 3,
					ColsNum: 3,
					Value: MatrixValue{
						{2, 2, 2},
						{2, 2, 2},
						{2, 2, 2},
					},
				},
				{
					RowsNum: 2,
					ColsNum: 3,
					Value: MatrixValue{
						{2, 3, 4},
						{-1, 4, 3},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMatsFromFiles(tt.args.paths...)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMatsFromFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMatsFromFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPopulateNewMat(t *testing.T) {
	type args struct {
		mat     Matrix
		action  func(mv MatrixValue, r, c int, secMvs ...MatrixValue) float64
		secMats []Matrix
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{
			name: "add 1 to each el",
			args: args{
				mat: Matrix{
					RowsNum: 3,
					ColsNum: 3,
					Value: MatrixValue{
						{2, 2, 2},
						{2, 2, 2},
						{2, 2, 2},
					},
				},
				action: func(mv MatrixValue, r, c int, secMvs ...MatrixValue) float64 {
					secMv := secMvs[0]

					return mv[r][c] + secMv[r][c]
				},
				secMats: []Matrix{
					{
						RowsNum: 3,
						ColsNum: 3,
						Value: MatrixValue{
							{1, 1, 1},
							{1, 1, 1},
							{1, 1, 1},
						},
					},
				},
			},
			want: Matrix{
				RowsNum: 3,
				ColsNum: 3,
				Value: MatrixValue{
					{3, 3, 3},
					{3, 3, 3},
					{3, 3, 3},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopulateNewMat(tt.args.mat, tt.args.action, tt.args.secMats...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopulateNewMat() = %v, want %v", got, tt.want)
			}
		})
	}
}
