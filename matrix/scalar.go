package matrix

func Scalar(scalarValue float64, matrix *Matrixf) *Matrixf {
	returnMatrix := Matrixf{
		Internal: [][]float64{},
	}
	for index := range matrix.Internal {
		var innerArray []float64
		for _, innervalue := range matrix.Internal[index] {
			innerArray = append(innerArray, scalarValue*innervalue)
		}
		returnMatrix.Internal = append(returnMatrix.Internal, innerArray)
	}
	return &returnMatrix
}
