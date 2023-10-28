package neuralnetwork

func MeanOfNetworkResultPerceptron(result NeuronOutputArray) *NeuronOutput {
	finresult := NeuronOutput{
		Y: 0,
	}

	sum := 0.0
	numvalues := 0
	for index, value := range result.Y {
		sum += value
		numvalues += index
	}
	final := sum / float64(numvalues)
	finresult.Y = final

	return &finresult
}
