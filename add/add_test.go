package add

import (
	"reflect"
	"testing"

	u "github.com/branogarbo/mtrix/util"
)

func TestAdd(t *testing.T) {
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
						RowsNum: 3,
						ColsNum: 3,
						Value: u.MatrixValue{
							{1, 2, 3},
							{4, 5, 6},
							{7, 8, 9},
						},
					},
					{
						RowsNum: 3,
						ColsNum: 3,
						Value: u.MatrixValue{
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
				Value: u.MatrixValue{
					{2, 3, 4},
					{5, 6, 7},
					{8, 9, 10},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Add(tt.args.mats...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
