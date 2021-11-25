package transpose

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
					{1, 2, 3, 4},
					{5, 6, 7, 8},
					{9, 10, 11, 12},
					{13, 14, 15, 16},
				},
			},
			want: u.Matrix{
				{1, 5, 9, 13},
				{2, 6, 10, 14},
				{3, 7, 11, 15},
				{4, 8, 12, 16},
			},
		},
		{
			name: "mismatch size trans",
			args: args{
				mat: u.Matrix{
					{1, 2, 3},
					{5, 6, 7},
				},
			},
			want: u.Matrix{
				{1, 5},
				{2, 6},
				{3, 7},
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
