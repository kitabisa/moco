package moco

import (
	"strings"
)

const (
	MandiriIndexNumberDescription    = 4
	MandiriIndexNumberValidationDate = 2
	MandiriIndexNumberAmount         = 8
)

var MandiriBlacklist = []string{"SA", "OB", "CA", "No", "Book", "DARI"}

type mandiriParser struct {
	record      []string
	accountName string
	description string
	amount      string
	date        string
}

func (p *mandiriParser) ParseRecord() error {
	p.description = p.record[MandiriIndexNumberDescription]
	p.accountName = p.parseAccountName(p.description)
	p.amount = p.record[MandiriIndexNumberAmount]
	p.date = p.record[MandiriIndexNumberValidationDate]

	return nil
}

func (p *mandiriParser) parseAccountName(s string) string {
	ns := WhitespaceSplit(s)
	ns = BlacklistTrim(ns, MandiriBlacklist)

	return strings.Join(ns, " ")
}

func NewMandiriParser() MutationParser {
	return &mandiriParser{}
}

func (p *mandiriParser) GetAccountName() string {
	return p.accountName
}

func (p *mandiriParser) GetAccountNumber() string {
	return ""
}

func (p *mandiriParser) LoadRecord(record []string) error {
	p.record = record
	return p.ParseRecord()
}

func (p *mandiriParser) GetDescription() string {
	return p.description
}

func (p *mandiriParser) GetAmount() string {
	a, err := NumericTrim(p.amount)

	if err != nil {
		return ""
	}

	return a
}

func (p *mandiriParser) GetDate() string {
	return p.date
}
