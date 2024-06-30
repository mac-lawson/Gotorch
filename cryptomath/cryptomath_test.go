package cryptomath_test

import (
	"reflect"
	"testing"

	"github.com/mac-lawson/gotorch/cryptomath"
	"github.com/mac-lawson/gotorch/testhelper"
)

func TestRandomInt8(t *testing.T) {
	ri8, err := cryptomath.CryptoRandomInt8()
	testhelper.Ok(t, err)
	testhelper.Equals(t, "int8", reflect.TypeOf(ri8).String())
}

func TestRandomInt64(t *testing.T) {
	ri8, err := cryptomath.CryptoRandomInt64()
	testhelper.Ok(t, err)
	testhelper.Equals(t, "int64", reflect.TypeOf(ri8).String())
}
