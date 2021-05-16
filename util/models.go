package util

type Row []float64

type MatVal []Row

type Matrix struct {
	RowsNum int
	ColsNum int
	Value   MatVal
}

type MatPopConfig struct {
	MainMat Matrix
	SecMats []Matrix
	NewRows int
	NewCols int
	Action  func(mv MatVal, r, c int, secMvs []MatVal) float64
}
