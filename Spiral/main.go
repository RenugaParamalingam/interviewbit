package main

import "fmt"

func main() {
	fmt.Println("Spiral order matrix")
	// null matrix
	// matrix := [][]int{}

	// 1 x 3 matrix
	matrix := [][]int{{1}, {2}, {3}}
	// 4 x 4 matrix
	// matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}

	// 5 x 4 matrix
	// matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}, {17, 18, 19, 20}}

	// 5 x 5 matrix
	// matrix := [][]int{{1, 2, 3, 4, 21}, {5, 6, 7, 8, 22}, {9, 10, 11, 12, 23}, {13, 14, 15, 16, 24}, {17, 18, 19, 20, 25}}

	fmt.Println("Left to Right Result: ", spiralOrderInLeftToRight(matrix))

	fmt.Println("Right to Left Result: ", spiralOrderInRightToLeft(matrix))
}

func spiralOrderInLeftToRight(matrix [][]int) []int {
	var result []int

	fmt.Println("Length of matrix: ", len(matrix))

	if len(matrix) != 0 {
		rowMin := 0
		rowMax := len(matrix)
		colMin := 0
		colMax := len(matrix[0])

		for (rowMin < rowMax) && (colMin < colMax) {
			//left to right
			for i := colMin; i <= colMax-1; i++ {
				// fmt.Printf("L to R: [%d][%d] is %d \n", colMin, i, matrix[colMin][i])
				result = append(result, matrix[colMin][i])
			}

			rowMin++

			//right to down
			for i := rowMin; i <= rowMax-1; i++ {
				// fmt.Printf("R to D: [%d][%d] is %d \n", i, colMax-1, matrix[i][colMax-1])
				result = append(result, matrix[i][colMax-1])
			}

			colMax--

			//right to left
			if rowMin < rowMax {
				for i := colMax - 1; i >= colMin; i-- {
					// fmt.Printf("R to L: [%d][%d] is %d \n", rowMax-1, i, matrix[rowMax-1][i])
					result = append(result, matrix[rowMax-1][i])
				}
				rowMax--
			}

			//left to top
			if colMin < colMax {
				for i := rowMax - 1; i >= rowMin; i-- {
					// fmt.Printf("L to T: [%d][%d] is %d \n", i, colMin, matrix[i][colMin])
					result = append(result, matrix[i][colMin])
				}
				colMin++
			}
		}
	}

	return result
}

func spiralOrderInRightToLeft(matrix [][]int) []int {
	var result []int

	if len(matrix) != 0 {
		rowMin := 0
		rowMax := len(matrix)
		colMin := 0
		colMax := len(matrix[0])

		for (rowMin < rowMax) && (colMin < colMax) {
			// Right to left
			for i := colMax - 1; i >= colMin; i-- {
				// fmt.Printf("R to L: [%d][%d is %d \n", colMin, i, matrix[colMin][i])
				result = append(result, matrix[colMin][i])
			}

			rowMin++

			// Left to down
			for i := rowMin; i <= rowMax-1; i++ {
				// fmt.Printf("L to D: [%d][%d] is %d \n", i, colMin, matrix[i][colMin])
				result = append(result, matrix[i][colMin])
			}

			colMin++

			if rowMax > rowMin {
				// left to right
				for i := colMin; i <= colMax-1; i++ {
					// fmt.Printf("L to R: [%d][%d] is %d \n", rowMax-1, i, matrix[rowMax-1][i])
					result = append(result, matrix[rowMax-1][i])
				}
				rowMax--
			}

			if colMax > colMin {

				//right to top
				for i := rowMax - 1; i >= rowMin; i-- {
					// fmt.Printf("R to T: [%d][%d] is %d \n", i, colMax-1, matrix[i][colMax-1])
					result = append(result, matrix[i][colMax-1])
				}
				colMax--
			}
		}
	}
	return result
}
