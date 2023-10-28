package neuralnetwork

import (
	"errors"

	"github.com/mac-lawson/gotorch/utils"
)

type Token struct {
	isPosative bool
	Ammount    int64
}

func PositiveReinforcement(result float64, goal float64) (error, bool) {
	if utils.Similar(result, goal) {
		return errors.New("types are not the same"), false
	} else {
		// TODO implement logic here, if there is some effort twords the goal reward a token.
		if goal > result {
			return nil, true

		} else {

			return nil, false

		}
	}
}
