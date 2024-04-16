package reader

import (
	"fmt"
	"mime/multipart"
)

type Format int

const (
	CSV_FORMAT Format = iota
	TEXT_FORMAT
	XML_FORMAT
)

func (f Format) String() string {
	switch f {
	case CSV_FORMAT:
		return "text/csv"
	case TEXT_FORMAT:
		return "text/plain"
	case XML_FORMAT:
		return "application/xml"
		// Adds new formats
	default:
		return "unknown"
	}
}

// FileReader is an interface for interacting with file process acording ext file
type FileReader interface {
	Read() ([][]string, error)
}

// ReadFileByType executes reading process according file type
func ReadFileByType(file *multipart.FileHeader) ([][]string, error) {
	var reader FileReader

	switch f := file.Header.Get("Content-Type"); f {
	case CSV_FORMAT.String():
		reader = NewCSVReader(file, ',', "")
	case TEXT_FORMAT.String():
		reader = NewTextFileReader(file)
	case XML_FORMAT.String():
		reader = NewXMLFileReader(file)
		// Adds new elements
	default:
		return nil, fmt.Errorf("format not available")
	}

	return reader.Read()
}
