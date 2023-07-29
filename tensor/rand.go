package tensor

import "math/rand"

/* TODO
* Rewrite the RandomFloat64 random array func to return the struct now (they are public)
*
*
 */

func RandomFloat64(xsize int32, ysize int32) *Gotensor_dtypefloat64 {
	tensorarray := make([][]float64, xsize)
	for i := range tensorarray {
		tensorarray[i] = make([]float64, ysize)
		for j := range tensorarray[i] {
			tensorarray[i][j] = rand.Float64() // Generates a random float between 0 and 1
		}
	}
	return &Gotensor_dtypefloat64{
		Data: tensorarray,
	}
}
