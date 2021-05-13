package util

func IsMatrixValid(m Matrix) bool {
	var matLength int

	for _, row := range m.Value {
		for range row {
			matLength++
		}
	}

	return (m.RowsNum*m.ColsNum == matLength) && (m.RowsNum == len(m.Value))
}

func IsMultPossible(m1, m2 Matrix) bool {
	return m1.RowsNum == m2.ColsNum
}

func GetMatrixFromFile(path string) Matrix {
	return Matrix{} //for now
}
