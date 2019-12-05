package moco

import (
	"encoding/csv"
	"io"
)

type CsvReader interface {
	Read() ([]string, error)
	ReadAll() ([][]string, error)
}

type csvReader struct {
	Reader *csv.Reader
}

func NewCsvReader(r io.Reader, delimiter rune) CsvReader {

	reader := csv.NewReader(r)
	reader.Comma = delimiter

	return &csvReader{
		Reader: reader,
	}
}

func (cr *csvReader) Read() ([]string, error) {
	return cr.Reader.Read()
}

func (cr *csvReader) ReadAll() ([][]string, error) {
	return cr.Reader.ReadAll()
}
