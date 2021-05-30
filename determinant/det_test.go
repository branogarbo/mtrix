package determinant

import (
	"testing"

	u "github.com/branogarbo/mtrix/util"
)

func TestMatDet(t *testing.T) {
	type args struct {
		m u.Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name: "det = 0, working",
			args: args{
				m: u.Matrix{
					Value: u.MatVal{
						{1, 8, 3},
						{-2, 2, -6},
						{3, 0, 9},
					},
				},
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "2x2, det = 0, working",
			args: args{
				m: u.Matrix{
					Value: u.MatVal{
						{1, 3},
						{-2, -6},
					},
				},
			},
			want:    0,
			wantErr: false,
		},
		{
			name: "2x3 error",
			args: args{
				m: u.Matrix{
					Value: u.MatVal{
						{1, 3, 8},
						{-2, -6, 2},
					},
				},
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "4x4, det = 0, working",
			args: args{
				m: u.Matrix{
					Value: u.MatVal{
						{1, 3, 2, 2},
						{-2, -6, 3, -4},
						{7, 6, 2, 14},
						{0, 1, 3, 0},
					},
				},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MatDet(tt.args.m)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatDet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MatDet() = %v, want %v", got, tt.want)
			}
		})
	}
}
