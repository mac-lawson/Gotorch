package errorhandling

/* TensorGo: an error first platform! ;)
*	At TensorGo, I'm trying to make errors as friendly as possible
*	Once the MKDOCS are up, each error will be raised with a link to a *helpfull* debug assistance page
*
*
 */

type documentationLink struct {
	url string
}

func TensorNotMatching() string {
	return "[TensorGo Error] The length of the tensor and the weights are not the same."
}

func ActiviationFunctionOptionNotProvided(detail string) string {
	return "[TensorGo Error] The option 32-bit Integer provided was not 1, 2, or 3. \nThe y value is now 0.0. \n 1: Sigmoid \t 2: Tanh \t 3: ReLu"
}

// probally won't need this, the Activation function already returns an error
func InnerNeuronError() string {
	return "[TensorGo Error] An error occured inside of the "
}
