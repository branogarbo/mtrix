package util

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// IsMatrixValid checks if m is missing any elements.
func IsMatrixValid(m Matrix) bool {
	var matLength int

	for _, row := range m.Value {
		for range row {
			matLength++
		}
	}

	return (m.RowsNum*m.ColsNum == matLength) && (m.RowsNum == len(m.Value))
}

// IsMultPossible checks if multiplication between m1 and m2 is possible.
func IsMultPossible(m1, m2 Matrix) bool {
	return m1.ColsNum == m2.RowsNum
}

// GetMatFromFile returns a matrix from a matrix file.
func GetMatFromFile(path string) (Matrix, error) {
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		return Matrix{}, err
	}

	var (
		fileStr = strings.TrimSpace(string(fileBytes))
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

// GetMatsFromFiles returns a slice of matrices from their matrix files.
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

// CheckMats checks if any mats are missing elements.
func CheckMats(mats ...Matrix) error {
	for _, m := range mats {
		if !IsMatrixValid(m) {
			return errors.New("invalid matrix passed")
		}
	}

	return nil
}

// CheckMatSizes checks if all mats are the same size.
func CheckMatSizes(mats ...Matrix) error {
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

// PopulateNewMat creates and fills a new matrix according to the action
// performed on each element of the passed matrices and configurations.
func PopulateNewMat(c MatPopConfig) Matrix {
	var (
		argMvs    []MatVal
		resultMat = InitMat(c.NewRows, c.NewCols)
	)

	argMvs = append(argMvs, c.MainMat.Value)

	for _, m := range c.SecMats {
		argMvs = append(argMvs, m.Value)
	}

	for rn, row := range resultMat.Value {
		for cn := range row {
			resultMat.Value[rn][cn] = c.Action(c.MainMat.Value, rn, cn, argMvs[1:])
		}
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

// InitMat creates a zero matrix with the passed size. Main purpose
// is to init matrix that can later be populated with PopulateNewMat.
func InitMat(rows, cols int) Matrix {
	resultMat := Matrix{
		RowsNum: rows,
		ColsNum: cols,
	}

	for i := 0; i < rows; i++ {
		row := Row{}
		for i := 0; i < cols; i++ {
			row = append(row, 0)
		}
		resultMat.Value = append(resultMat.Value, row)
	}

	return resultMat
}
