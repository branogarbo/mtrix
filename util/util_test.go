package util

import (
	"reflect"
	"testing"
)

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
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
				m2: Matrix{
					{1, 1, 1},
					{1, 1, 1},
				},
			},
			want: false,
		},
		{
			name: "possible",
			args: args{
				m1: Matrix{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
				m2: Matrix{
					{1, 1, 1, 1},
					{1, 1, 1, 1},
					{1, 1, 1, 1},
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

func TestCheckMatsSizes(t *testing.T) {
	type args struct {
		mats []Matrix
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "no error",
			args: args{
				mats: []Matrix{
					{
						{0, 0, 0},
						{0, 0, 0},
					},
					{
						{0, 0, 0},
						{0, 0, 0},
					},
					{
						{0, 0, 0},
						{0, 0, 0},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				mats: []Matrix{
					{
						{0, 0, 0},
						{0, 0, 0},
						{0, 0, 0},
					},
					{
						{0, 0, 0},
						{0, 0, 0},
					},
					{
						{0, 0, 0},
						{0, 0, 0},
					},
				},
			},
			wantErr: true,
		},
		{
			name: "invalid args",
			args: args{
				mats: []Matrix{
					{
						{0, 0, 0},
						{0, 0, 0},
					},
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckMatSizes(tt.args.mats...); (err != nil) != tt.wantErr {
				t.Errorf("CheckMatSizes() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetMatsFromFiles(t *testing.T) {
	type args struct {
		paths []string
	}
	tests := []struct {
		name    string
		args    args
		want    []Matrix
		wantErr bool
	}{
		{
			name: "parse matrix files test",
			args: args{
				paths: []string{"../sampleMats/2x2_1.txt", "../sampleMats/3x3_2.txt", "../sampleMats/2x3_1.txt"},
			},
			want: []Matrix{
				{
					{1, 2},
					{3, -5},
				},
				{
					{2, 2, 2},
					{2, 2, 2},
					{2, 2, 2},
				},
				{
					{2, 3, 4},
					{-1, 4, 3},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMatsFromFiles(tt.args.paths)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMatsFromFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMatsFromFiles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPopulateNewMat(t *testing.T) {
	type args struct {
		c MatPopConfig
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{
			name: "add 1 to each el",
			args: args{
				c: MatPopConfig{
					MainMat: Matrix{
						{2, 2, 2},
						{2, 2, 2},
					},
					SecMats: []Matrix{
						{
							{1, 1, 1},
							{1, 1, 1},
						},
					},
					Action: func(m1 Matrix, r, c int, secMs []Matrix) float64 {
						mv2 := secMs[0]

						return m1[r][c] + mv2[r][c]
					},
				},
			},
			want: Matrix{
				{3, 3, 3},
				{3, 3, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PopulateNewMat(tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopulateNewMat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitMat(t *testing.T) {
	type args struct {
		rows int
		cols int
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{
			name: "2x3",
			args: args{
				rows: 2,
				cols: 3,
			},
			want: Matrix{
				{0, 0, 0},
				{0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitMat(tt.args.rows, tt.args.cols); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitMat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMakeIdentityMat(t *testing.T) {
	type args struct {
		wid int
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{
			name: "3x3 working",
			args: args{3},
			want: Matrix{
				{1, 0, 0},
				{0, 1, 0},
				{0, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeIdentityMat(tt.args.wid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MakeIdentityMat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMinor(t *testing.T) {
	type args struct {
		m   Matrix
		row int
		c   int
	}
	tests := []struct {
		name string
		args args
		want Matrix
	}{
		{
			name: "working",
			args: args{
				m: Matrix{
					{1, -4, 2},
					{0, 6, -7},
					{5, 8, 0},
				},
				row: 0,
				c:   1,
			},
			want: Matrix{
				{0, -7},
				{5, 0},
			},
		},
		{
			name: "working",
			args: args{
				m: Matrix{
					{1, -4, 2},
					{0, 6, -7},
					{5, 8, 0},
				},
				row: 0,
				c:   2,
			},
			want: Matrix{
				{0, 6},
				{5, 8},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMinor(tt.args.m, tt.args.row, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMinor() = %v, want %v", got, tt.want)
			}
		})
	}
}
