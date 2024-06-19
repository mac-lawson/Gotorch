package matrix

import "github.com/mac-lawson/gotorch/tensor"

type Matrixf struct {
	Internal [][]float64
}

func (m Matrixf) ToTensor() (tensor.Gotensor_dtypefloat64, error) {
	return tensor.Gotensor_dtypefloat64{
		Data: m.Internal,
	}, nil
}

type Matrixi struct {
	Internal [][]int64
}

func (m Matrixi) ToTensor() (tensor.Gotensor_dtypeint64, error) {
	return tensor.Gotensor_dtypeint64{
		Data: m.Internal,
	}, nil
}
