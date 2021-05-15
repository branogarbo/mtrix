package util

import (
	"errors"
	"fmt"
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

func GetMatFromFile(path string) (Matrix, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return Matrix{}, err
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
				return Matrix{}, err
			}

			row = append(row, elFloat)
		}

		matrix.Value = append(matrix.Value, row)
	}

	matrix.RowsNum = len(matrix.Value)
	matrix.ColsNum = len(matrix.Value[0])

	err = CheckMats(matrix)
	if err != nil {
		return Matrix{}, err
	}

	return matrix, nil
}

func GetMatsFromFiles(paths ...string) ([]Matrix, error) {
	var mats []Matrix

	// use goroutines?
	for _, path := range paths {
		mat, err := GetMatFromFile(path)
		if err != nil {
			return nil, err
		}

		mats = append(mats, mat)
	}

	return mats, nil
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

	for i := 1; i < len(mats); i++ {
		if !((mats[i].RowsNum == mats[i-1].RowsNum) && (mats[i].ColsNum == mats[i-1].ColsNum)) {
			return errors.New("matrices are not the same size")
		}
	}

	return nil
}

func PopulateNewMat(mat Matrix, action func(mv MatrixValue, r, c int, secMvs ...MatrixValue) float64, secMats ...Matrix) Matrix {
	var (
		resultMat Matrix
		newMatVal MatrixValue
		argMvs    []MatrixValue
	)

	argMvs = append(argMvs, mat.Value)

	for _, m := range secMats {
		argMvs = append(argMvs, m.Value)
	}

	for r, row := range mat.Value {
		newMatVal = append(newMatVal, Row{})
		for c := range row {
			newEl := action(mat.Value, r, c, argMvs[1:]...)

			newMatVal[r] = append(newMatVal[r], newEl)
		}
	}

	resultMat = Matrix{
		RowsNum: len(newMatVal),
		ColsNum: len(newMatVal[0]),
		Value:   newMatVal,
	}

	return resultMat
}

func PrintMat(mat Matrix) {
	for _, row := range mat.Value {
		for _, el := range row {
			fmt.Printf("%v ", el)
		}
		fmt.Print("\n")
	}
}
