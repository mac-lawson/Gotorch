package optparser

import "os"

type OptParser struct {
	CommandType uint8
}

type FlagMap struct {
	FullMap map[int]string
}

const (
	MODEL       int = 0
	HELP        int = 1
	TENSOR_MATH int = 2
)

func Flags() FlagMap {
	ArgsMap := make(map[int]string)
	argsWithoutProg := os.Args[1:]
	for index, arg := range argsWithoutProg {
		ArgsMap[index] = arg
	}
	return FlagMap{FullMap: ArgsMap}
}
