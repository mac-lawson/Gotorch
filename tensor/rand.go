package tensor

import (
	"github.com/mac-lawson/gotorch/cryptomath"
)

/* TODO
* Rewrite the RandomFloat64 random array func to return the struct now (they are public)
*
*
 */
func RandomFloat64(xsize int32, ysize int32) *Tensorfloat64 {
	tensorarray := make(Tensorfloat64, xsize)
	for i := range tensorarray {
		tensorarray[i] = make([]float64, ysize)
		for j := range tensorarray[i] {
			tensorarray[i][j], _ = cryptomath.CryptoRandomFloat64() // Generates a random float between 0 and 1
		}
	}
	return &tensorarray
}
