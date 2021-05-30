package inverse

import (
	"reflect"
	"testing"

	u "github.com/branogarbo/mtrix/util"
)

func TestMatInv(t *testing.T) {
	type args struct {
		m u.Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    u.Matrix
		wantErr bool
	}{
		{
			name: "2x2, working",
			args: args{u.Matrix{
				Value: u.MatVal{
					{5, -1},
					{5, 1},
				},
			}},
			want: u.Matrix{
				RowsNum: 2,
				ColsNum: 2,
				Value: u.MatVal{
					{0.1, 0.1},
					{-0.5, 0.5},
				},
			},
			wantErr: false,
		},
		{
			name: "2x2, singular, error",
			args: args{u.Matrix{
				Value: u.MatVal{
					{5, 1},
					{5, 1},
				},
			}},
			want:    u.Matrix{},
			wantErr: true,
		},
		{
			name: "3x3, working",
			args: args{u.Matrix{
				Value: u.MatVal{
					{1, 0, 0},
					{0, 1, 0},
					{0, 0, 1},
				},
			}},
			want: u.Matrix{
				RowsNum: 3,
				ColsNum: 3,
				Value: u.MatVal{
					{1, 0, 0},
					{0, 1, 0},
					{0, 0, 1},
				},
			},
			wantErr: false,
		},
		{
			name: "3x3, working",
			args: args{u.Matrix{
				Value: u.MatVal{
					{3, 0, 2},
					{2, 0, -2},
					{0, 1, 1},
				},
			}},
			want: u.Matrix{
				RowsNum: 3,
				ColsNum: 3,
				Value: u.MatVal{
					{0.2, 0.2, 0},
					{-0.2, 0.30000000000000004, 1}, // uhhhhhhh
					{0.2, -0.30000000000000004, 0},
				},
			},
			wantErr: false,
		},
		{
			name: "3x3, working",
			args: args{u.Matrix{
				Value: u.MatVal{
					{1, 2, 3},
					{0, 1, 4},
					{5, 6, 0},
				},
			}},
			want: u.Matrix{
				RowsNum: 3,
				ColsNum: 3,
				Value: u.MatVal{
					{-24, 18, 5},
					{20, -15, -4},
					{-5, 4, 1},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MatInv(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatInv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatInv() = %v, want %v", got, tt.want)
			}
		})
	}
}
