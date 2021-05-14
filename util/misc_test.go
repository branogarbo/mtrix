package util

import (
	"reflect"
	"testing"
)

func TestGetMatrixFromFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    Matrix
		wantErr bool
	}{
		{
			name: "parse test",
			args: args{path: "./testMatrix.txt"},
			want: Matrix{
				Value: MatrixValue{
					Row{1, 2, 3},
					Row{4, 5, 6},
					Row{5, -3, 4},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMatrixFromFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMatrixFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMatrixFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
						Row{1, 1, 1},
						Row{1, 1, 1},
						Row{1, 1, 1},
					},
				},
				m2: Matrix{
					RowsNum: 2,
					ColsNum: 3,
					Value: MatrixValue{
						Row{1, 1, 1},
						Row{1, 1, 1},
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
						Row{1, 1, 1},
						Row{1, 1, 1},
						Row{1, 1, 1},
					},
				},
				m2: Matrix{
					RowsNum: 3,
					ColsNum: 4,
					Value: MatrixValue{
						Row{1, 1, 1, 1},
						Row{1, 1, 1, 1},
						Row{1, 1, 1, 1},
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
						Row{1, 1, 1, 1},
						Row{1, 1, 1, 1},
						Row{1, 1, 1},
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
						Row{1, 1, 1, 1},
						Row{1, 1, 1, 1},
						Row{1, 1, 1, 1},
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
