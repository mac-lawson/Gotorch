package utils

import (
	"reflect"
)

/*Purpose of this file: is too
* detirmine the type of params passed to ulity functions such as mean.
 */

func Numerical(tensor any) bool {
	if reflect.TypeOf(tensor).Name() != "Gotensor_dtypefloat64" {
		return false
	} else {
		return true
	}
}
