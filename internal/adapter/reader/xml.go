package reader

import (
	"encoding/xml"
	"fmt"
	"io"
	"mime/multipart"
)

type XMLFileReader struct {
	file *multipart.FileHeader
}

func NewXMLFileReader(file *multipart.FileHeader) *XMLFileReader {
	return &XMLFileReader{file: file}
}

type Row struct {
	Site string `xml:"site"`
	ID   int    `xml:"id"`
	// Adds new columns
}

type Root struct {
	XMLName xml.Name `xml:"root"`
	Rows    []Row    `xml:"row"`
}

func (xfr *XMLFileReader) Read() ([][]string, error) {
	file, err := xfr.file.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var rows Root

	if err := xml.Unmarshal(byteValue, &rows); err != nil {
		return nil, err
	}

	var records [][]string

	for _, row := range rows.Rows {
		fields := []string{row.Site, fmt.Sprintf("%d", row.ID)}
		records = append(records, fields)
	}

	return records, nil
}
