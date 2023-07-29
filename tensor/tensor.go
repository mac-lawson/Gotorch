package tensor

type Gotensor_dtypefloat64 struct {
	Data [][]float64
}

type Gotensor_dtypeint64 struct {
	data [][]int64
}

type Gotensor_dtypestring struct {
	data [][]string
}

/* TODO Detirmine if the structs for the tensors should be public/private
* I kept the init methods below for each of the data types.
*
*
 */

// func GoTensorFloat64(data [][]float64) *gotensor_dtypefloat64 {
// 	return &gotensor_dtypefloat64{
// 		data: data,
// 	}
// }

// func GoTensorInt64(data [][]int64) *gotensor_dtypeint64 {
// 	return &gotensor_dtypeint64{
// 		data: data,
// 	}
// }

// func GoTensorString(data [][]string) *gotensor_dtypestring {
// 	return &gotensor_dtypestring{
// 		data: data,
// 	}
// }
