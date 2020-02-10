package moco

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BniMutationTestSuite struct {
	suite.Suite
	parser                MutationParser
	BniMutationRec        []string
	InvalidBniMutationRec []string
}

func (suite *BniMutationTestSuite) SetupTest() {
	suite.BniMutationRec = []string{"02/12/19 06.25.21", "02/12/19 06.25.21", "0996", "932902", "TRANSFER DARI | PEMINDAHAN DARI 719147165 Sdr JOHN DOE", ".00", "200,000.00"}
	suite.InvalidBniMutationRec = []string{"test", "invalid", "record"}
	suite.parser = NewBniParser()
}

func (suite *BniMutationTestSuite) TestLoadRecord() {
	err := suite.parser.LoadRecord(suite.BniMutationRec)
	assert.Nil(suite.T(), err, "Error should be nil")
}

func (suite *BniMutationTestSuite) TestLoadRecordInvalid() {
	err := suite.parser.LoadRecord(suite.InvalidBniMutationRec)
	assert.Nil(suite.T(), err, "Error should be nil")
}

func (suite *BniMutationTestSuite) TestGetAccountName() {
	err := suite.parser.LoadRecord(suite.BniMutationRec)
	v := suite.parser.GetAccountName()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "JOHN DOE", v, "Account name is empty")
}

func (suite *BniMutationTestSuite) TestGetAccountNumber() {
	err := suite.parser.LoadRecord(suite.BniMutationRec)
	v := suite.parser.GetAccountNumber()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "719147165", v, "Account number is empty")
}

func (suite *BniMutationTestSuite) TestGetAmount() {
	err := suite.parser.LoadRecord(suite.BniMutationRec)
	v := suite.parser.GetAmount()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "200000", v, "Amount is wrong")
}

func (suite *BniMutationTestSuite) TestGetDescription() {
	err := suite.parser.LoadRecord(suite.BniMutationRec)
	v := suite.parser.GetDescription()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "TRANSFER DARI | PEMINDAHAN DARI 719147165 Sdr JOHN DOE", v, "Description is wrong")
}

func (suite *BniMutationTestSuite) TestGetDate() {
	err := suite.parser.LoadRecord(suite.BniMutationRec)
	v := suite.parser.GetDate()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "02/12/19", v, "Date is wrong")
}

func TestBniMutationTestSuite(t *testing.T) {
	suite.Run(t, new(BniMutationTestSuite))
}
