package moco

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MandiriMutationTestSuite struct {
	suite.Suite
	parser                              MutationParser
	MandiriMutationRec                  []string
	MandiriMutationRecWithAccountNumber []string
}

func (suite *MandiriMutationTestSuite) SetupTest() {
	suite.MandiriMutationRec = []string{"1270010264453", "25/11/19", "25/11/19", "2315", "SA OB CA No Book  DARI INDAH FEBRIANTY", "DONASI PT KITA BISA ", "", ".00", "350,000.00", ""}
	suite.MandiriMutationRecWithAccountNumber = []string{"1270010264453", "25/11/19", "25/11/19", "2315", "Transfer Otomatis", "DARI  9000006726914 KE  1270010264453 ", "", ".00", "100,000.00", ""}
	suite.parser = NewMandiriParser()
}

func (suite *MandiriMutationTestSuite) TestLoadRecord() {
	err := suite.parser.LoadRecord(suite.MandiriMutationRec)
	assert.Nil(suite.T(), err, "Error should be nil")
}

func (suite *MandiriMutationTestSuite) TestGetAccountName() {
	err := suite.parser.LoadRecord(suite.MandiriMutationRec)
	v := suite.parser.GetAccountName()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "INDAH FEBRIANTY", v, "Account name is empty")
}

func (suite *MandiriMutationTestSuite) TestGetAccountNumber() {
	err := suite.parser.LoadRecord(suite.MandiriMutationRecWithAccountNumber)
	v := suite.parser.GetAccountNumber()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "9000006726914", v, "Account number is empty")
}

func (suite *MandiriMutationTestSuite) TestGetAmount() {
	err := suite.parser.LoadRecord(suite.MandiriMutationRec)
	v := suite.parser.GetAmount()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "350000", v, "Amount is wrong")
}

func (suite *MandiriMutationTestSuite) TestGetDescription() {
	err := suite.parser.LoadRecord(suite.MandiriMutationRec)
	v := suite.parser.GetDescription()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "SA OB CA No Book  DARI INDAH FEBRIANTY", v, "Description is wrong")
}

func (suite *MandiriMutationTestSuite) TestGetDate() {
	err := suite.parser.LoadRecord(suite.MandiriMutationRec)
	v := suite.parser.GetDate()

	assert.Nil(suite.T(), err, "Error should be nil")
	assert.Equal(suite.T(), "25/11/19", v, "Date is wrong")
}

func TestMandiriutationTestSuite(t *testing.T) {
	suite.Run(t, new(MandiriMutationTestSuite))
}
