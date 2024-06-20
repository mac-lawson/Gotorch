package main

import (
	"github.com/mac-lawson/gotorch/datasets"
)

/*
func main() {
	n2 := tensor.RandomFloat64(4, 4)
	wi := neuralnetwork.RandomWeights(3)
	neuralnetwork.BinaryClassification(2, 1, *wi, *n2, true)
	m1 := matrix.Matrixf{
		Internal: [][]float64{
			{1.0, 2},
			{3.5, 9.0},
		},
	}
	m2 := matrix.Matrixf{
		Internal: [][]float64{
			{1.0, 2},
			{3.5, 9.0},
		},
	}

	m3 := matrix.Add(&m1, &m2)
	fmt.Println(m3)
}
*/

func main() {
	data, err := datasets.FromCSVFile("datasets/samples/data.csv")
	if err != nil {
		panic(err)
	}
	// fmt.Println(data)

	//new := data.Pop(0)
	//fmt.Println(new)
	//
	d2, err2 := data.Tokenize()
	if err2 != nil {
		panic(err2)
	}
	// fmt.Println(d2)
	// d2.Helper()
	d2.Dump()
}
