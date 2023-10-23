package tensor

type Gotensor_dtypefloat64 struct {
	Data [][]float64
}

type Gotensor_dtypeint64 struct {
	Data [][]int64
}

type Gotensor_dtypestring struct {
	Data [][]string
}

/* TODO Detirmine if the structs for the tensors should be public/private
* I kept the init methods below for each of the data types.
*
*
 */

func GoTensorFloat64(data [][]float64) *Gotensor_dtypefloat64 {
	return &Gotensor_dtypefloat64{
		Data: data,
	}
}

func GoTensorInt64(data [][]int64) *Gotensor_dtypeint64 {
	return &Gotensor_dtypeint64{
		Data: data,
	}
}

func GoTensorString(data [][]string) *Gotensor_dtypestring {
	return &Gotensor_dtypestring{
		Data: data,
	}
}
