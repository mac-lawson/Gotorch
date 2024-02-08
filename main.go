package main

import (
	"fmt"

	"github.com/mac-lawson/gotorch/matrix"
)

func main() {
	m1 := matrix.Matrix{
		Internal: [][]float64{
			{1.0, 2},
			{3.5, 9.0},
		},
	}
	m2 := matrix.Matrix{
		Internal: [][]float64{
			{1.0, 2},
			{3.5, 9.0},
		},
	}

	m3 := matrix.Add(&m1, &m2)
	fmt.Println(m3)
}
