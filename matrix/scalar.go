package matrix

func Scalar(scalarValue float64, matrix *Matrix) *Matrix {
	returnMatrix := Matrix{
		Internal: [][]float64{},
	}
	for index, _ := range matrix.Internal {
		var innerArray []float64
		for _, innervalue := range matrix.Internal[index] {
			innerArray = append(innerArray, scalarValue*innervalue)
		}
		returnMatrix.Internal = append(returnMatrix.Internal, innerArray)
	}
	return &returnMatrix
}
