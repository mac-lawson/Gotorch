package optparser_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/mac-lawson/gotorch/optparser"
)

func TestFlags(t *testing.T) {
	fmt.Println(os.Args)
	os.Args = append(os.Args, "build")
	os.Args = append(os.Args, "--perceptron")
	os.Args = append(os.Args, `-o`)
	fmt.Println(os.Args)
	result := optparser.Flags()
	fmt.Println(result)
}
