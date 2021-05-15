package util

type Row []float64

type MatrixValue []Row

type Matrix struct {
	RowsNum int
	ColsNum int
	Value   MatrixValue
}
