package main

import (
	"fmt"

	"github.com/mac-lawson/tensorgo/neuralnetwork"
	"github.com/mac-lawson/tensorgo/tensor"
	// "math/rand"
)

func main() {
	tn := tensor.Gotensor_dtypefloat64{
		Data: [][]float64{
			{5.3, 6.6, 7.8},
			{4.4, 6.6, 7.9},
			{3.5, 2.6, 7.8},
			{7.0, 6.6, 7.7},
			{8.5, 4.6, 7.6},
			{9.3, 6.6, 7.4},
			{5.7, 1.6, 7.8},
		},
	}
	resp, err := neuralnetwork.SimpleNeuralNetwork(2, 1, tn)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resp)
	}

}
