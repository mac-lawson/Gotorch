package datasets

type CSVdata struct {
	Data [][]string
}

type PDFdata struct {
	Data []byte
}

type Dataset struct {
	Type        int64
	Description string
	Source      string
}

const (
	CSV  int = 0
	PDF  int = 1
	JSON int = 2
	YAML int = 3
)
