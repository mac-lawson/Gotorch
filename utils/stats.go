package utils

import (
	"errors"
	"fmt"

	"github.com/mac-lawson/gotorch/tensor"
)

// MeanFloat64 calculates the mean of a tensor containing float64 values.
// It prints the column number and its corresponding mean value.
// If the input tensor is not numerical, it returns 0.
func MeanFloat64(tensor tensor.Gotensor_dtypefloat64) (error, []float64) {
	result := []float64{}
	if Numerical(tensor) {
		fmt.Println("Column | Mean")
		for index, _ := range tensor.Data {
			sum := 0.0
			total := 0.0
			for _, value1 := range tensor.Data[index] {
				sum += value1
				total += 1
			}
			sum /= total
			fmt.Println(index, "\t", sum)
			result = append(result, sum)
		}
	} else {
		return errors.New("the data provided was not numerical or an issue was encountered with the data"), result
	}
	return nil, result
}

// MeanInt64 calculates the mean of a tensor containing int64 values.
// It prints the column number and its corresponding mean value.
// If the input tensor is not numerical, it returns 0.
func MeanInt64(tensor tensor.Gotensor_dtypeint64) any {
	if Numerical(tensor) {
		fmt.Println("Column | Mean")
		for index, _ := range tensor.Data {
			var sum int64 = 0
			var total int64 = 0
			for _, value1 := range tensor.Data[index] {
				sum += value1
				total += 1
			}
			sum /= total
			fmt.Println(index, sum)
		}
	} else {
		return 0
	}
	return 0
}
