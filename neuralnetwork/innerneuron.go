package neuralnetwork

/*
* TODO Inner Neuron functions
*	function should be able to take an array of inputs and an array of weights
*	to should run the Activiation function and return the Y value for each epoch
*	- Phase out the weightX input. For now, I will just set it as 1 if I don't need it
*
 */

func InnerNeuron(tensorX float64, weightX float64, b float64, activiationFunctionType uint8) (float64, error) {
	xInputValue := float64((tensorX * weightX) + b)
	functionReturnValue, err := Activation(activiationFunctionType, xInputValue)
	if err != nil {
		return 0.0, err
	} else {
		return functionReturnValue, err
	}

}

/* For those who are confused....
* This function is public (indicated by the uppercase func name)
* Its purpose is to return a clean, single float64 y value
* The regular innerneuron function returns both a float64
 */
// func InnerNeuron(tensorX float32, weightX float64, b float64, activiationFunctionType int32) float64 {
// }
