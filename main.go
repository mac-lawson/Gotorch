package main

import (
	"fmt"

	"github.com/mac-lawson/gotorch/tensor"
	"github.com/mac-lawson/gotorch/utils"
)

func main() {

	tn := tensor.Gotensor_dtypefloat64{
		Data: [][]float64{
			{1.3, 1.4},
			{1.5},
			{2.3, 5.6, 10},
		},
	}

	// utils.Mean(tn)
	fmt.Println(utils.Numerical(tn))
}
