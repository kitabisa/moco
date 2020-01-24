package moco

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BcaMutationTestSuite struct {
	suite.Suite
	parser                MutationParser
	BcaMutationRec        []string
	InvalidBcaMutationRec []string
}

func (suite *BcaMutationTestSuite) SetupTest() {
	suite.BcaMutationRec = []string{"PEND", "TRSF E-BANKING CR 2711/ADSCY/0000100 19112700651710  Lesman Buulolo Wallet User  ", "0498", "100,000.00 CR", "585,000.00"}
	suite.InvalidBcaMutationRec = []string{"test"}
	suite.parser = NewBcaParser()
}

func (suite *BcaMutationTestSuite) TestLoadRecord() {
	err := suite.parser.LoadRecord(suite.BcaMutationRec)
	assert.Nil(suite.T(), err, "Error should be nil")
}

func (suite *BcaMutationTestSuite) TestLoadRecordInvalid() {
	err := suite.parser.LoadRecord(suite.InvalidBcaMutationRec)
	assert.Nil(suite.T(), err, "Error should be nil")
}

func (suite *BcaMutationTestSuite) TestGetAccountName() {
	err := suite.parser.LoadRecord(suite.BcaMutationRec)
	v := suite.parser.GetAccountName()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "Lesman Buulolo", v, "Account name is wrong")
}

func (suite *BcaMutationTestSuite) TestGetAccountNumber() {
	err := suite.parser.LoadRecord(suite.BcaMutationRec)
	v := suite.parser.GetAccountNumber()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "", v, "Account number is not empty")
}

func (suite *BcaMutationTestSuite) TestGetAmount() {
	err := suite.parser.LoadRecord(suite.BcaMutationRec)
	v := suite.parser.GetAmount()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "100000", v, "Amount is wrong")
}

func (suite *BcaMutationTestSuite) TestGetDescription() {
	err := suite.parser.LoadRecord(suite.BcaMutationRec)
	v := suite.parser.GetDescription()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "TRSF E-BANKING CR 2711/ADSCY/0000100 19112700651710  Lesman Buulolo Wallet User  ", v, "Description is wrong")
}

func (suite *BcaMutationTestSuite) TestGetDate() {
	err := suite.parser.LoadRecord(suite.BcaMutationRec)
	v := suite.parser.GetDate()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "2711", v, "Date is wrong")
}

func TestBcaMutationTestSuite(t *testing.T) {
	suite.Run(t, new(BcaMutationTestSuite))
}
