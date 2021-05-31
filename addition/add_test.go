package addition

import (
	"reflect"
	"testing"

	u "github.com/branogarbo/mtrix/util"
)

func TestMatAdd(t *testing.T) {
	type args struct {
		mats []u.Matrix
	}
	tests := []struct {
		name    string
		args    args
		want    u.Matrix
		wantErr bool
	}{
		{
			name: "add matrices",
			args: args{
				mats: []u.Matrix{
					{
						Value: u.MatVal{
							{1, 2, 3},
							{4, 5, 6},
							{7, 8, 9},
						},
					},
					{
						Value: u.MatVal{
							{1, 1, 1},
							{1, 1, 1},
							{1, 1, 1},
						},
					},
				},
			},
			want: u.Matrix{
				RowsNum: 3,
				ColsNum: 3,
				Value: u.MatVal{
					{2, 3, 4},
					{5, 6, 7},
					{8, 9, 10},
				},
			},
			wantErr: false,
		},
		{
			name: "nonmatching sizes",
			args: args{
				mats: []u.Matrix{
					{
						Value: u.MatVal{
							{1, 2, 3},
							{4, 5, 6},
							{7, 8, 9},
						},
					},
					{
						Value: u.MatVal{
							{1, 1},
							{1, 1},
						},
					},
					{
						Value: u.MatVal{
							{1, 1, 1},
							{1, 1, 1},
						},
					},
				},
			},
			want:    u.Matrix{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MatAdd(tt.args.mats...)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatAdd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MatAdd() = %v, want %v", got, tt.want)
			}
		})
	}
}
