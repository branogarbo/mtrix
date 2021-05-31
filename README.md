# 🔮 **mtrix**
**A tool for performing basic matrix operations.**

---

## 🧪 **Setup**
Download and compile from sources:
```
go get github.com/branogarbo/mtrix
```
Install just the binary with Go:
```
go install github.com/branogarbo/mtrix@latest
```

Or get the pre-compiled binaries for your platform on the [releases page](https://github.com/branogarbo/mtrix/releases)


## 🧙‍♂️ **CLI Usage**
```
mtrix

A tool for performing basic matrix operations.

Usage:
  mtrix [command]

Available Commands:
  add         Get the sum of matrices
  det         Compute the determinant of a matrix
  help        Help about any command
  inv         Get the inverse of a matrix
  mult        Multiply two matrices together
  pow         Raise a matrix to the nth power
  smult       Multiply a matrix by a scalar
  sub         Get the difference of two matrices
  trans       Get the transpose of a matrix

Flags:
  -h, --help        help for mtrix
  -r, --raw-input   whether or not the command takes matrices as strings

Use "mtrix [command] --help" for more information about a command.
```

---

## 🌌 **Package Usage**
Get the packages you need for your project:
```
go get github.com/branogarbo/mtrix/addition
go get github.com/branogarbo/mtrix/multiply
go get github.com/branogarbo/mtrix/util
...
```

### **Examples:**
Adding two matrices and printing the result:
``` go
package main

import (
	"log"

	add "github.com/branogarbo/mtrix/addition"
	u "github.com/branogarbo/mtrix/util"
)

func main() {
	mat1 := u.Matrix{
		Value: u.MatVal{
			{1, -4, 9},
			{2.7, 3, 7},
			{0, -3.3, 8},
		},
	}

	mat2 := u.Matrix{
		Value: u.MatVal{
			{1, 1, 1},
			{1, 1, 1},
			{1, 1, 1},
		},
	}

	resultMat, err := add.MatAdd(mat1, mat2)
	if err != nil {
		log.Fatal(err)
	}

	u.PrintMat(resultMat)
}
```

Multiplying two matrices together and printing the result:
``` go
package main

import (
	"log"

	mult "github.com/branogarbo/mtrix/multiply"
	u "github.com/branogarbo/mtrix/util"
)

func main() {
	mat1 := u.Matrix{
		Value: u.MatVal{
			{1, -4},
			{2.7, 3},
			{0, -3.3},
		},
	}

	mat2 := u.Matrix{
		Value: u.MatVal{
			{1, 0},
			{0, 1},
		},
	}

	resultMat, err := mult.MatMult(mat1, mat2)
	if err != nil {
		log.Fatal(err)
	}

	u.PrintMat(resultMat)
}
```