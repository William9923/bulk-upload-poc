package mocks

import (
	bytes "bytes"

	"github.com/stretchr/testify/mock"
)

func New(err error) *ICsv {

	mockCsv := &ICsv{}
	mockCsv.On("Export").Return(*new(bytes.Buffer), err)
	mockCsv.On("Save", mock.Anything).Return(err)
	mockCsv.On("GetHeader").Return([]string{})
	mockCsv.On("GetContents").Return([][]string{})
	return mockCsv
}

func NewBuilder(err error) *ICsvBuilder {

	mockCsvBuilder := &ICsvBuilder{}
	mockCsvBuilder.On("Build", mock.Anything, mock.Anything).Return(New(err))
	mockCsvBuilder.On("FromFile", mock.Anything).Return(New(err), err)
	return mockCsvBuilder
}
