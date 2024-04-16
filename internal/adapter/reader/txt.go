package reader

import (
	"bufio"
	"mime/multipart"
	"strings"
)

type TextFileReader struct {
	file *multipart.FileHeader
}

func NewTextFileReader(file *multipart.FileHeader) *TextFileReader {
	return &TextFileReader{file: file}
}

func (tfr *TextFileReader) Read() ([][]string, error) {
	file, err := tfr.file.Open()
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var records [][]string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := strings.Split(scanner.Text(), ",")
		records = append(records, []string{line[0], line[1]})
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
