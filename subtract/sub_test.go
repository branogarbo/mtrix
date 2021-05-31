package subtract

import (
	"reflect"
	"testing"

	u "github.com/branogarbo/mtrix/util"
)

func TestMatSub(t *testing.T) {
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
			name: "wanted",
			args: args{
				m1: u.Matrix{
					Value: u.MatVal{
						{3, 3, 3},
						{3, 3, 3},
						{3, 3, 3},
					},
				},
				m2: u.Matrix{
					Value: u.MatVal{
						{1, 1, 1},
						{1, 1, 1},
						{1, 1, 1},
					},
				},
			},
			want: u.Matrix{
				RowsNum: 3,
				ColsNum: 3,
				Value: u.MatVal{
					{2, 2, 2},
					{2, 2, 2},
					{2, 2, 2},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MatSub(tt.args.m1, tt.args.m2)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatSub() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatSub() = %v, want %v", got, tt.want)
			}
		})
	}
}
