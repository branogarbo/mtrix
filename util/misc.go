package util

import (
	"errors"
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

func CheckMats(mats ...Matrix) error {
	for _, m := range mats {
		if !IsMatrixValid(m) {
			return errors.New("invalid matrix passed")
		}
	}

	return nil
}

func CheckMatsSizes(mats ...Matrix) error {
	if len(mats) <= 1 {
		return errors.New("2 or more matrices must be passed")
	}

	for i := 1; i < len(mats)-1; i++ {
		if !((mats[i].RowsNum == mats[i-1].RowsNum) && (mats[i].ColsNum == mats[i-1].ColsNum)) {
			return errors.New("matrices are not the same size")
		}
	}

	return nil
}

func PopulateNewMat(mv MatrixValue, action func(mv MatrixValue, r, c int, secMvs ...MatrixValue) float64, secMvs ...MatrixValue) MatrixValue {
	resultMatVal := MatrixValue{}

	for r, row := range mv {
		resultMatVal = append(resultMatVal, Row{})
		for c := range row {
			newEl := action(mv, r, c, secMvs...)

			resultMatVal[r] = append(resultMatVal[r], newEl)
		}
	}

	return resultMatVal
}
