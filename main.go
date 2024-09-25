package main

import (
	"errors"
	"fmt"
	"sync"

	"math/rand/v2"
)

var wg sync.WaitGroup

const MAT1_ROWS, MAT2_ROWS = 100, 100
const MAT1_COLS, MAT2_COLS = 100, 100

func parallenCompute(idx1 int, idx2 int, a [MAT1_ROWS][MAT1_COLS]int, b [MAT2_ROWS][MAT2_COLS]int, intChan chan int) {
	defer wg.Done()
	var res int
	for idx3 := range len(a[0]) {
		res += a[idx1][idx3] * b[idx3][idx2]
	}
    intChan <- res
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

	resChannels := make([]chan int, rRows*rCols)

	for idx := range resChannels {
		resChannels[idx] = make(chan int, rCols)
	}  
    
	var cidx int = 0
	for idx := range a {
		for idx2 := range b[0] {
			wg.Add(1)
			go parallenCompute(idx, idx2, a, b, resChannels[cidx])
			cidx += 1
		}
	}
    cidx = 0
	for tidx := range a {
        for tidx2 := range b[0] {
			result[tidx][tidx2] = <- resChannels[cidx]
			cidx += 1
		}
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

	res, err := matrixMultiplication(a, b)
    
	if err != nil {
		panic(err)
	}
	fmt.Println(res)

}