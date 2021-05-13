package util

func IsMatrixComplete(m Matrix) bool {
	var matLength int

	for _, row := range m.Value {
		for range row {
			matLength++
		}
	}

	return m.Width*m.Height == matLength
}

func IsMultPossible(m1, m2 Matrix) bool {
	return m1.Height == m2.Width
}
