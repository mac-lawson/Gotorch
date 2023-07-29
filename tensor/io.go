package tensor

import (
	"fmt"
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
	for i := 0; i < len(tensor.data); i++ {
		for si := 0; si < len(tensor.data[i]); si++ {
			fmt.Print(tensor.data[i][si], "\t")
		}
		fmt.Println()
	}
}
