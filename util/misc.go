package util

import (
	"os"
	"strconv"
	"strings"
)

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
	return m1.ColsNum == m2.RowsNum
}

func GetMatrixFromFile(path string) (Matrix, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return Matrix{}, nil
	}

	var (
		fileStr = string(fileBytes)
		rowStrs = strings.Split(fileStr, "\n")
		matrix  = Matrix{}
	)

	for _, rowStr := range rowStrs {
		elStrs := strings.Split(rowStr, " ")
		row := Row{}

		for _, elStr := range elStrs {
			elFloat, err := strconv.ParseFloat(elStr, 64)
			if err != nil {
				return Matrix{}, nil
			}

			row = append(row, elFloat)
		}

		matrix.Value = append(matrix.Value, row)
	}

	return matrix, nil
}
