package utils

type ErrorCodes struct {
	Code []uint64
	ErrInfo []string

}
// ErrorDirectory: Return the error directory
// codes: An array of the error codes
// info: The corresponding information for debugging the error codes

func ErrorDirectory(codes []uint64, info []string) *ErrorCodes {
	
	return &ErrorCodes{
		Code: codes, 
		ErrInfo: info,

	}

}
