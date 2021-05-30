package util

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// IsMatrixValid checks if m is missing any elements.
func IsMatrixValid(m Matrix) bool {
	mv := m.Value

	for i := 1; i < len(mv); i++ {
		if len(mv[i]) != len(mv[i-1]) {
			return false
		}
	}

	return true
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

	return StringToMat(string(fileBytes))
}

// StringToMat parses a string to a matrix.
func StringToMat(ms string) (Matrix, error) {
	var (
		matStr  = strings.ReplaceAll(strings.TrimSpace(ms), "\r", "")
		rowStrs = strings.Split(matStr, "\n")
		matrix  = Matrix{}
		err     error
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

	err = matrix.SetSize()
	if err != nil {
		return Matrix{}, err
	}

	return matrix, nil
}

// StringsToMats parses strings to matrices.
func StringsToMats(mats []string) ([]Matrix, error) {
	var matStrs []Matrix

	for _, m := range mats {
		str, err := StringToMat(m)
		if err != nil {
			return nil, err
		}

		matStrs = append(matStrs, str)
	}

	return matStrs, nil
}

// GetMatsFromFiles returns a slice of matrices from their matrix files.
func GetMatsFromFiles(paths []string) ([]Matrix, error) {
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

// MatToString returns mat in its string form.
func MatToString(mat Matrix) string {
	var matStr string

	for _, row := range mat.Value {
		for c, el := range row {
			matStr += fmt.Sprint(el)

			if c != len(row)-1 {
				matStr += " "
			}
		}
		matStr += "\n"
	}

	return matStr
}

// PrintMat prints mat to the command line.
func PrintMat(mat Matrix) {
	matStr := MatToString(mat)

	fmt.Print(matStr)
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

// ParseCmdArgs parses args according to the command raw-input flag.
func ParseCmdArgs(cmd *cobra.Command, args []string) ([]Matrix, error) {
	var mats []Matrix

	isRaw, err := cmd.Flags().GetBool("raw-input")
	if err != nil {
		return nil, err
	}

	if isRaw {
		mats, err = StringsToMats(args)
	} else {
		mats, err = GetMatsFromFiles(args)
	}
	if err != nil {
		return nil, err
	}

	return mats, nil
}

// MakeIdentityMat creates an indentity matrix with wid as the number of rows and cols.
func MakeIdentityMat(w int) Matrix {
	m := InitMat(w, w)

	for r, row := range m.Value {
		for c := range row {
			if r == c {
				m.Value[r][c] = 1
			}
		}
	}

	return m
}

// GetMinor returns the minor of m according to row at column c.
func GetMinor(m Matrix, row, c int) Matrix {
	newMv := append(MatVal{}, m.Value[:row]...)
	newMv = append(newMv, m.Value[row+1:]...)

	minor := InitMat(m.RowsNum-1, 0)
	minor.ColsNum = m.ColsNum - 1

	for r, row := range newMv {
		minor.Value[r] = append(minor.Value[r], row[:c]...)
		minor.Value[r] = append(minor.Value[r], row[c+1:]...)
	}

	return minor
}

// SetSize sets RowsNum and ColsNum of m based on its value.
func (m *Matrix) SetSize() error {
	err := CheckMats(*m)
	if err != nil {
		return err
	}

	m.RowsNum = len(m.Value)
	m.ColsNum = len(m.Value[0])

	return nil
}

// SetSizes sets RowsNum and ColsNum of each matrix in ms based on its value.
func SetSizes(ms ...Matrix) ([]Matrix, error) {
	var (
		mats []Matrix
		err  error
	)

	for _, m := range ms {
		err = m.SetSize()
		if err != nil {
			return nil, err
		}

		mats = append(mats, m)
	}

	return mats, nil
}
