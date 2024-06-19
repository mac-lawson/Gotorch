package datasets

const (
	CSV int = 0
	PDF int = 1
	JSON int = 2
	YAML int = 3

)

type Dataset struct {
	
	Type int64
	Description string
	Source string

}



