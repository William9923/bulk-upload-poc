package csv

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

type ICsvBuilder interface {
	Build(header []string, content [][]string) ICsv
	FromFile(file io.Reader) (ICsv, error)
}

type ICsv interface {
	Export() (bytes.Buffer, error)
	Save(path string) error
	GetHeader() []string
	GetContents() [][]string
}

type Csv struct {
	Header  []string
	Content [][]string
}

type CsvBuilder struct{}

func NewCsvBuilder() ICsvBuilder {
	return &CsvBuilder{}
}

func (c *CsvBuilder) Build(header []string, contents [][]string) ICsv {

	return &Csv{
		Header:  header,
		Content: contents,
	}
}

func (c *CsvBuilder) FromFile(file io.Reader) (ICsv, error) {

	// read the file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	// check empty file
	if err != nil || len(records) < 1 {
		return nil, err
	}

	// get csv file header
	header := records[0]

	// handle empty csv files
	contents := [][]string{}
	if len(records) > 1 {
		contents = records[1:]
	}

	return c.Build(header, contents), nil
}

func (c *Csv) Export() (bytes.Buffer, error) {
	commaDelimiter := ','
	return c.ExportWithCustomDelimiter(&commaDelimiter, true)
}

func (c Csv) GetHeader() []string {
	return c.Header
}

func (c Csv) GetContents() [][]string {
	return c.Content
}

func (c *Csv) ExportWithCustomDelimiter(delimiter *rune, mustHeader bool) (bytes.Buffer, error) {
	var data [][]string

	if mustHeader || len(c.Header) > 0 {
		data = append(data, c.Header)
	}

	if len(c.Content) > 0 {
		data = append(data, c.Content...)
	}

	var buf bytes.Buffer
	write := csv.NewWriter(&buf)

	// if no delimiter set, comma is default
	if delimiter != nil {
		write.Comma = *delimiter
	}

	defer write.Flush()

	err := write.WriteAll(data)
	if err != nil {
		return buf, err
	}
	return buf, nil
}

func (c *Csv) Save(path string) error {
	buffer, err := c.Export()
	if err != nil {
		return err
	}

	// If the file doesn't exist, create the file
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	_, err = buffer.WriteTo(f)
	if err != nil {
		return err
	}
	return nil
}
