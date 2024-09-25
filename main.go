package main

import (
	"errors"
	"fmt"
	"sync"

	"math/rand/v2"
)

var wg sync.WaitGroup

const MAT1_ROWS, MAT1_COLS = 150, 150
const MAT2_ROWS, MAT2_COLS = 150, 150

func parallenCompute(idx1 int, idx2 int, a [MAT1_ROWS][MAT1_COLS]int, b [MAT2_ROWS][MAT2_COLS]int, result *[][]int) {
	colsA := len(a[0]) //either columns of matrix a or rows of matrix b
	for idx3 := range colsA {
		(*result)[idx1][idx2] += a[idx1][idx3] * b[idx3][idx2]
	}
}

func validate(a [MAT1_ROWS][MAT1_COLS]int, b [MAT2_ROWS][MAT2_COLS]int) (bool, error) {
	aCols := len(a[0])
	bRows := len(b)

	if aCols != bRows {
		return false, errors.New("matrix dimensions are violationg the rules")
	}
	return true, nil
}


func matrixMultiplication(a [MAT1_ROWS][MAT1_COLS]int, b [MAT2_ROWS][MAT2_COLS]int) ([][]int, error){
	validMatrix, err := validate(a, b)
	if !validMatrix {
        return [][]int{}, err
	}

    rRows := len(a)
	rCols := len(b[0])

	result := make([][]int, rRows)
	for idx := range rRows {
		result[idx] = make([]int, rCols)
	} 

    
	for idx := range a {
		worker := func() {
			for idx2 := range b[0] {
			    parallenCompute(idx, idx2, a, b, &result)
			}
			defer wg.Done()
		};

		wg.Add(1)

		go worker()
	}

    wg.Wait()
    return result, nil
}


func main() {
	a := [MAT1_ROWS][MAT1_COLS]int{}

	b := [MAT2_ROWS][MAT2_COLS]int{}

	for i := range a {
		for j := range a[0] {
            a[i][j] = rand.IntN(10)
		}
	} 

	for i := range b {
		for j := range b[0] {
            b[i][j] = rand.IntN(10)
		}
	} 

	fmt.Println("Original Matrix A : ", a)
	fmt.Println("Original Matrix B : ", b)

	res, err := matrixMultiplication(a, b)
    
	if err != nil {
		panic(err)
	}
	fmt.Println("Result Matrix ")
	fmt.Println(res)

}