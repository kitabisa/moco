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
	suite.BniMutationRec = []string{"04/05/20 12.28.07", "04/05/20 12.28.07", "0989", "359823", "TRANSFER DARI | |7878078081 20200504945299342 | Sdr EKA DANA KRISTANTO   ", ".00", "250,000.00"}
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
	assert.Equal(suite.T(), "EKA DANA KRISTANTO", v, "Account name is empty")
}

// New BNI CSV doesn't include bank account number
// func (suite *BniMutationTestSuite) TestGetAccountNumber() {
// 	err := suite.parser.LoadRecord(suite.BniMutationRec)
// 	v := suite.parser.GetAccountNumber()

// 	assert.Nil(suite.T(), err, "Error should be nil")
// 	assert.Equal(suite.T(), "719147165", v, "Account number is empty")
// }

func (suite *BniMutationTestSuite) TestGetAmount() {
	err := suite.parser.LoadRecord(suite.BniMutationRec)
	v := suite.parser.GetAmount()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "250000", v, "Amount is wrong")
}

func (suite *BniMutationTestSuite) TestGetDescription() {
	err := suite.parser.LoadRecord(suite.BniMutationRec)
	v := suite.parser.GetDescription()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "TRANSFER DARI | |7878078081 20200504945299342 | Sdr EKA DANA KRISTANTO   ", v, "Description is wrong")
}

func (suite *BniMutationTestSuite) TestGetDate() {
	err := suite.parser.LoadRecord(suite.BniMutationRec)
	v := suite.parser.GetDate()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "04/05/20", v, "Date is wrong")
}

func TestBniMutationTestSuite(t *testing.T) {
	suite.Run(t, new(BniMutationTestSuite))
}
