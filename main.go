package main

import (
	"fmt"

	"github.com/mac-lawson/gotorch/tensor"
	"github.com/mac-lawson/gotorch/utils"
)

func main() {
	values := [][]int64{{1, 2, 3}, {2, 3, 4}}
	tn := tensor.Gotensor_dtypeint64{
		Data: values,
	}

	// utils.Mean(tn)
	fmt.Println(utils.Numerical(tn))
}
