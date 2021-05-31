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
