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
		name    string
		args    args
		want    u.Matrix
		wantErr bool
	}{
		{
			name: "transpose test",
			args: args{
				mat: u.Matrix{
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
			wantErr: false,
		},
		{
			name: "mismatch size trans",
			args: args{
				mat: u.Matrix{
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
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MatTrans(tt.args.mat)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatTrans() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatTrans() = %v, want %v", got, tt.want)
			}
		})
	}
}
