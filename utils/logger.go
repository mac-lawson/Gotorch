package utils

type LogError struct {
	Er error
	Message string

} 

type LogWarning struct {
	Message string	
	Code uint32

}
