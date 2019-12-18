package moco

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ReaderTestSuite struct {
	suite.Suite
	reader   Reader
	mockText string
}

func MockNewReader(f io.Reader, s string, rec []string) Reader {
	return &reader{
		ioReader: f,
		parser:   MockNewParser(rec),
		bankType: s,
	}
}

func (suite *ReaderTestSuite) SetupTest() {
	suite.mockText = "some,string,to,test"
	ioReader := strings.NewReader(suite.mockText)
	suite.reader = MockNewReader(ioReader, "some-bank", strings.Split(suite.mockText, ","))
}

func (suite *ReaderTestSuite) TestRead() {
	err := suite.reader.ReadMutation()

	assert.Nil(suite.T(), err, "Error should be nil")
}

func (suite *ReaderTestSuite) TestGetSuccess() {
	err := suite.reader.ReadMutation()
	mutations := suite.reader.GetSuccess()
	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), 1, len(mutations), "Success mutations is empty")
}

func (suite *ReaderTestSuite) TestGetFail() {
	err := suite.reader.ReadMutation()
	failures := suite.reader.GetFail()
	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), 0, len(failures), "Failure records is not empty")
}

func (suite *ReaderTestSuite) TestGetRaw() {
	err := suite.reader.ReadMutation()
	raws := suite.reader.GetRaw()
	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), 1, len(raws), "Raw records is empty")
}

func TestReaderTestSuite(t *testing.T) {
	suite.Run(t, new(ReaderTestSuite))
}
