package moco

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type BriMutationTestSuite struct {
	suite.Suite
	parser         MutationParser
	BriMutationRec []string
}

func (suite *BriMutationTestSuite) SetupTest() {
	suite.BriMutationRec = []string{"DATE", "TIME", "REMARK", "DEBET", "CREDIT", "TELLER ID", "26/11/19", "03:42:14", "DARI 020601084702507 KE  038501000860307", "0.00", "500,000.00", "1,438,000.00", "DDRAZN1"}
	suite.parser = NewBriParser()
}

func (suite *BriMutationTestSuite) TestLoadRecord() {
	err := suite.parser.LoadRecord(suite.BriMutationRec)
	assert.Nil(suite.T(), err, "Error should be nil")
}

func (suite *BriMutationTestSuite) TestGetAccountName() {
	err := suite.parser.LoadRecord(suite.BriMutationRec)
	v := suite.parser.GetAccountName()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "", v, "Account name is not empty")
}

func (suite *BriMutationTestSuite) TestGetAccountNumber() {
	err := suite.parser.LoadRecord(suite.BriMutationRec)
	v := suite.parser.GetAccountNumber()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "020601084702507", v, "Account number is  empty")
}

func (suite *BriMutationTestSuite) TestGetAmount() {
	err := suite.parser.LoadRecord(suite.BriMutationRec)
	v := suite.parser.GetAmount()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "500000", v, "Amount is wrong")
}

func (suite *BriMutationTestSuite) TestGetDescription() {
	err := suite.parser.LoadRecord(suite.BriMutationRec)
	v := suite.parser.GetDescription()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "DARI 020601084702507 KE  038501000860307", v, "Description is wrong")
}

func (suite *BriMutationTestSuite) TestGetDate() {
	err := suite.parser.LoadRecord(suite.BriMutationRec)
	v := suite.parser.GetDate()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "26/11/19", v, "Date is wrong")
}

func TestBriMutationTestSuite(t *testing.T) {
	suite.Run(t, new(BriMutationTestSuite))
}
