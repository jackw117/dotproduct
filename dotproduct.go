package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
	"strings"
)

//row (x) and column (y) sizes for the two matrices
var x1 int = 0
var x2 int = 0
var y1 int = 0
var y2 int = 0

//default sizes for matrix rows (X) and columns (Y)
var maxY int = 10
var maxX int = 10

func main() {
	if (len(os.Args) < 3) {
		fmt.Println("Please include the locations of the matrix files as command line arguments.")
	} else {
		//open both matrix files
	  f1, err := os.Open(os.Args[1]) //"../tmp/matrix1.txt"
	  check(err)
		defer f1.Close()

		f2, err := os.Open(os.Args[2]) //"../tmp/matrix2.txt"
	  check(err)
		defer f2.Close()

		//create matrices from the files
		matrix1 := getMatrix(f1, true)
		matrix2 := getMatrix(f2, false)

		//checks to see if the sizes of the matrices are correct for a dot product
		if (x1 != y2) {
			fmt.Println("Invalid matrix format. The number of columns in the first matrix must match the number of rows in the second matrix.")
		} else {
			dotproduct := getDotProduct(matrix1, matrix2)
			fmt.Println(dotproduct)
		}
	}
}

//error handling function
func check(e error) {
  if e != nil {
      panic(e)
  }
}

//takes a file containing a matrix as input, and a boolean to tell whether the matrix is the first or second
//updates global matrix sizes, and doubles the matrix if the input files are larger than the default max sizes
//returns a matrix built from the input file
func getMatrix(f *os.File, first bool) [][]int {
	matrix := make([][]int, maxY)
	scanner := bufio.NewScanner(f)
	y := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if (len(line) > maxX) {
			maxX = len(line)
		}

		matrix[y] = make([]int, maxX)
		var i int
		for i = range line {
			num, err := strconv.Atoi(line[i])
			check(err)
			matrix[y][i] = num
		}
		y++

		if (y == maxY) {
			matrix = doubleLength(matrix)
		}

		//increment row and column sizes based on what matrix is currently being read
		if (first) {
			x1 = i + 1
			y1++
		} else {
			x2 = i + 1
			y2++
		}
	}

	if err := scanner.Err(); err != nil {
		check(err)
	}

	return matrix
}

//takes a matrix as input and returns a new matrix that is double the length + 1 in case initial size is 0
//also updates the global maxY value to new doubled size
func doubleLength(matrix [][]int) [][]int {
	var newMax = 2 * len(matrix) + 1
	newMatrix := make([][]int, newMax)
	copy(newMatrix, matrix)
	maxY = newMax
	return newMatrix
}

//takes two matrices as input and returns a matrix that corresponds to the dot product of those two matrices
func getDotProduct(m1 [][]int, m2 [][]int) [][]int {
	dp := make([][]int, y1)

	for i := range dp {
		dp[i] = make([]int, x2)
	}

	for i := 0; i < y1; i++ {
		for k := 0; k < x2; k++ {
			var sum int = 0
			for j := 0; j < x1; j++ {
				sum += m1[i][j] * m2[j][k]
			}
			dp[i][k] = sum
		}
	}

	return dp
}
