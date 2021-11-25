package util

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

// IsMultPossible checks if multiplication between m1 and m2 is possible.
func IsMultPossible(m1, m2 Matrix) bool {
	return m1.Cols() == m2.Rows()
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

		matrix = append(matrix, row)
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

// CheckMatSizes checks if all mats are the same size.
func CheckMatSizes(mats ...Matrix) error {
	if len(mats) <= 1 {
		return errors.New("2 or more matrices must be passed")
	}

	for i := 1; i < len(mats); i++ {
		if !((mats[i].Rows() == mats[i-1].Rows()) && (mats[i].Cols() == mats[i-1].Cols())) {
			return errors.New("matrices are not the same size")
		}
	}

	return nil
}

// PopulateNewMat creates and fills a new matrix according to the action
// performed on each element of the passed matrices and configurations.
// Omit NewRows and NewCols fields in c if you want them to be the same as
// the rows and cols of MainMat.
func PopulateNewMat(c MatPopConfig) Matrix {
	if c.NewRows == 0 && c.NewCols == 0 {
		c.NewRows = c.MainMat.Rows()
		c.NewCols = c.MainMat.Cols()
	}

	var (
		wg        sync.WaitGroup
		argMs     []Matrix
		resultMat = InitMat(c.NewRows, c.NewCols)
	)

	argMs = append(argMs, c.MainMat)
	argMs = append(argMs, c.SecMats...)

	wg.Add(c.NewRows)

	for rn, row := range resultMat {
		go func(rn int, row Row) {
			for cn := range row {
				resultMat[rn][cn] = c.Action(c.MainMat, rn, cn, argMs[1:])
			}
			wg.Done()
		}(rn, row)
	}

	wg.Wait()

	return resultMat
}

// MatToString returns mat in its string form.
func MatToString(mat Matrix) string {
	var matStr string

	for _, row := range mat {
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

// PrintMat prints mat to the command line. Returns the
// number of bytes written and any write error encountered.
func PrintMat(mat Matrix) (n int, err error) {
	matStr := MatToString(mat)

	return fmt.Print(matStr)
}

// InitMat creates an empty matrix with the passed size. Main purpose
// is to init matrix that can later be populated with PopulateNewMat.
func InitMat(rows, cols int) Matrix {
	m := make(Matrix, rows)

	for r := 0; r < rows; r++ {
		m[r] = make(Row, cols)
	}

	return m
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

	for r, row := range m {
		for c := range row {
			if r == c {
				m[r][c] = 1
			}
		}
	}

	return m
}

// GetMinor returns the minor of m according to row at column c.
func GetMinor(m Matrix, row, c int) Matrix {
	newM := append(Matrix{}, m[:row]...)
	newM = append(newM, m[row+1:]...)

	minor := InitMat(m.Rows()-1, 0)
	// minor.ColsNum = m.ColsNum - 1

	for r, row := range newM {
		minor[r] = append(minor[r], row[:c]...)
		minor[r] = append(minor[r], row[c+1:]...)
	}

	return minor
}

// Get the number of rows of m.
func (m Matrix) Rows() int {
	return len(m)
}

// Get the number of columns of m.
func (m Matrix) Cols() int {
	return len(m[0])
}
