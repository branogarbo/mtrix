package multiply

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
					Value: u.MatVal{
						{1, 1},
						{1, 1},
					},
				},
			},
			want: u.Matrix{
				RowsNum: 2,
				ColsNum: 2,
				Value: u.MatVal{
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
					Value: u.MatVal{
						{2, 3},
						{3, -5},
					},
				},
				m2: u.Matrix{
					Value: u.MatVal{
						{1, 0},
						{0, 1},
					},
				},
			},
			want: u.Matrix{
				RowsNum: 2,
				ColsNum: 2,
				Value: u.MatVal{
					{2, 3},
					{3, -5},
				},
			},
			wantErr: false,
		},
		{
			name: "mismatch sizes",
			args: args{
				m1: u.Matrix{
					Value: u.MatVal{
						{1, 0, 0},
						{0, 1, 0},
						{0, 0, 1},
					},
				},
				m2: u.Matrix{
					Value: u.MatVal{
						{2, 3},
						{0, 1},
						{3, -4},
					},
				},
			},
			want: u.Matrix{
				RowsNum: 3,
				ColsNum: 2,
				Value: u.MatVal{
					{2, 3},
					{0, 1},
					{3, -4},
				},
			},
			wantErr: false,
		},
		{
			name: "mismatch sizes error",
			args: args{
				m1: u.Matrix{
					Value: u.MatVal{
						{1, 0, 0},
						{0, 1, 0},
						{0, 0, 1},
					},
				},
				m2: u.Matrix{
					Value: u.MatVal{
						{2, 3, 1},
						{0, 1, -2},
					},
				},
			},
			want:    u.Matrix{},
			wantErr: true,
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
