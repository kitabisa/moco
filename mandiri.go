package moco

import (
	"strings"
)

const (
	MandiriMinimumRecordLength       = 5
	MandiriIndexNumberSecondDescription = 5
	MandiriIndexNumberDescription    = 4
	MandiriIndexNumberValidationDate = 2
	MandiriIndexNumberAmount         = 8
)

var MandiriBlacklist = []string{"SA", "OB", "CA", "No", "Book", "DARI", "Transfer", "Otomatis", "KE"}

type mandiriParser struct {
	record        []string
	accountName   string
	accountNumber string
	description   string
	amount        string
	date          string
}

func (p *mandiriParser) validRecordLength() bool {
	return len(p.record) >= MandiriMinimumRecordLength
}

func (p *mandiriParser) parseRecord() error {
	desc := p.record[MandiriIndexNumberDescription]
	account := p.parseAccountName(desc)

	if account == "" {
		desc = p.record[MandiriIndexNumberSecondDescription]
		p.accountNumber = p.parseAccountNumber(desc)
	} else {
		p.accountName = account
	}

	p.description = desc
	p.amount = p.record[MandiriIndexNumberAmount]
	p.date = p.record[MandiriIndexNumberValidationDate]

	return nil
}

func (p *mandiriParser) isAllNumber(s string) bool {
	match, err := regexp.MatchString("^[0-9]*$", s)
	if err != nil {
		return false
	}

	return match
}

func (p *mandiriParser) parseAccountName(s string) string {
	ns := WhitespaceSplit(s)
	ns = BlacklistTrim(ns, MandiriBlacklist)

	return strings.Join(ns, " ")
}

func (p *mandiriParser) parseAccountNumber(s string) string {
	if !p.isContainAccountNumber(s) {
		return ""
	}

	ns := WhitespaceSplit(s)
	ns = BlacklistTrim(ns, MandiriBlacklist)

	if len(ns) > 1 {
		return ns[0]
	}

	return ""
}

func NewMandiriParser() MutationParser {
	return &mandiriParser{}
}

func (p *mandiriParser) GetAccountName() string {
	return p.accountName
}

func (p *mandiriParser) GetAccountNumber() string {
	return p.accountNumber
}

func (p *mandiriParser) LoadRecord(record []string) error {
	p.record = record
	if !p.validRecordLength() {
		return nil
	}
	return p.parseRecord()
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
