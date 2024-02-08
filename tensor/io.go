package tensor

import (
	"fmt"

	"github.com/mac-lawson/gotorch/matrix"
)

// Outputs a formatted 64-bit float tensor
func Readfloat64(tensor Gotensor_dtypefloat64) {
	for i := 0; i < len(tensor.Data); i++ {
		for si := 0; si < len(tensor.Data[i]); si++ {
			fmt.Print(tensor.Data[i][si], "\t")
		}
		fmt.Println()
	}
}

// Outputs a formatted 64-bit interger tensor
func Readint64(tensor Gotensor_dtypefloat64) {
	for i := 0; i < len(tensor.Data); i++ {
		for si := 0; si < len(tensor.Data[i]); si++ {
			fmt.Print(tensor.Data[i][si], "\t")
		}
		fmt.Println()
	}
}

// Outputs a formatted string tensor
func Readstring(tensor Gotensor_dtypestring) {
	for i := 0; i < len(tensor.Data); i++ {
		for si := 0; si < len(tensor.Data[i]); si++ {
			fmt.Print(tensor.Data[i][si], "\t")
		}
		fmt.Println()
	}
}

func Matrix_to_Tensor(matrix matrix.Matrix) *Gotensor_dtypefloat64 {
	return &Gotensor_dtypefloat64{
		Data: matrix.Internal,
	}
}
