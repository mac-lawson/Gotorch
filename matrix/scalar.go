package matrix

func Scalar(scalarValue float64, matrix *Matrix) *Matrix {
	returnMatrix := Matrix{
		Data: [][]float64{},
	}
	for index, _ := range matrix.Data {
		var innerArray []float64
		for _, innervalue := range matrix.Data[index] {
			innerArray = append(innerArray, scalarValue*innervalue)
		}
		returnMatrix.Data = append(returnMatrix.Data, innerArray)
	}
	return &returnMatrix
}
