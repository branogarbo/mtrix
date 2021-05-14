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
						RowsNum: 1,
						ColsNum: 3,
					},
					{
						RowsNum: 5,
						ColsNum: 4,
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

func TestPopulateNewMatVal(t *testing.T) {
	type args struct {
		mv     MatrixValue
		action func(mv MatrixValue, r, c int, secMvs ...MatrixValue) float64
		secMvs []MatrixValue
	}
	tests := []struct {
		name string
		args args
		want MatrixValue
	}{
		{
			name: "add 1",
			args: args{
				mv: MatrixValue{
					{1, 2, 4},
					{1, 3, 1},
					{3, 3, 3},
				},
				action: func(mv MatrixValue, r, c int, secMvs ...MatrixValue) float64 {
					secMv := secMvs[0]

					return mv[r][c] + secMv[r][c]
				},
				secMvs: []MatrixValue{
					{
						{1, 1, 1},
						{1, 1, 1},
						{1, 1, 1},
					},
				},
			},
			want: MatrixValue{
				{2, 3, 5},
				{2, 4, 2},
				{4, 4, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopulateNewMatVal(tt.args.mv, tt.args.action, tt.args.secMvs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopulateNewMatVal() = %v, want %v", got, tt.want)
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
				paths: []string{"../sampleMats/mat1.txt", "../sampleMats/mat2.txt"},
			},
			want: []Matrix{
				{
					RowsNum: 3,
					ColsNum: 3,
					Value: MatrixValue{
						{1, 2, -3.9},
						{4.3, 5, 6},
						{5, -3, 4},
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
