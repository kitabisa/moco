package moco

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ParserTestSuite struct {
	suite.Suite
	parser Parser
}

type MutationParserMock struct {
	mock.Mock
}

func (m *MutationParserMock) LoadRecord(record []string) error {
	args := m.Called(record)
	return args.Error(0)

}

func (m *MutationParserMock) GetAccountName() string {
	args := m.Called()
	return args.String(0)

}

func (m *MutationParserMock) GetAccountNumber() string {
	args := m.Called()
	return args.String(0)

}

func (m *MutationParserMock) GetDescription() string {
	args := m.Called()
	return args.String(0)

}

func (m *MutationParserMock) GetAmount() string {
	args := m.Called()
	return args.String(0)

}

func (m *MutationParserMock) GetDate() string {
	args := m.Called()
	return args.String(0)

}

func MockNewParser(rec []string) Parser {
	mocked := new(MutationParserMock)
	mocked.On("LoadRecord", rec).Return(nil)
	mocked.On("GetAccountName").Return("some-name")
	mocked.On("GetAccountNumber").Return("some-number")
	mocked.On("GetDescription").Return("some-desc")
	mocked.On("GetAmount").Return("1000")
	mocked.On("GetDate").Return("01/01/1970")
	return &parser{
		mutationParser: mocked,
	}
}

func (suite *ParserTestSuite) SetupTest() {
	suite.parser = MockNewParser([]string{})
}

func (suite *ParserTestSuite) TestLoadRecord() {
	err := suite.parser.LoadRecord([]string{})

	assert.Nil(suite.T(), err, "Error should be nil")
}

func (suite *ParserTestSuite) TestGetMutation() {
	err := suite.parser.LoadRecord([]string{})
	mutation := suite.parser.GetMutation()
	assert.Nil(suite.T(), err, "Error should be nil")
	assert.NotNil(suite.T(), mutation, "Mutation is nil")
}

func TestParserTestSuite(t *testing.T) {
	suite.Run(t, new(ParserTestSuite))
}
