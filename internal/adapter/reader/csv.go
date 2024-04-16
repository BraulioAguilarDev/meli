package reader

import (
	"encoding/csv"
	"mime/multipart"
)

// Implementing for CSV files
type CSVReader struct {
	file      *multipart.FileHeader
	separator rune
	encoding  string
}

func NewCSVReader(filePath *multipart.FileHeader, separator rune, encoding string) *CSVReader {
	return &CSVReader{
		file:      filePath,
		separator: separator,
		encoding:  encoding,
	}
}

func (cr *CSVReader) Read() ([][]string, error) {
	file, err := cr.file.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.Comma = cr.separator

	if cr.encoding != "" {
		r.FieldsPerRecord = -1
	}

	records, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
