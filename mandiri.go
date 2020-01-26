package moco

import (
	"regexp"
	"strings"
)

const (
	MandiriMinimumRecordLength          = 5
	MandiriIndexNumberSecondDescription = 5
	MandiriIndexNumberDescription       = 4
	MandiriIndexNumberValidationDate    = 2
	MandiriIndexNumberAmount            = 8
)

var MandiriBlacklist = []string{
	"SA", "OB", "CA", "No", "Book", "DARI", "Transfer",
	"Otomatis", "KE", "MCM", "InhouseTrf", "KITA", "BISA",
	"Auto", "Overbooking"}

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

func (p *mandiriParser) isContainAccountNumber(s string) bool {
	match, err := regexp.MatchString("^\\bDARI\\b\\s*[0-9]+\\s*\\bKE\\b\\s*[0-9]+", s)
	if err != nil {
		return false
	}

	return match
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
	p.amount = p.parseAmount(p.record[MandiriIndexNumberAmount])
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

func (p *mandiriParser) parseAmount(s string) string {
	a, err := NumericTrim(s)

	if err != nil {
		return ""
	}

	return a
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
	p.cleanUp()
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
	return p.amount
}

func (p *mandiriParser) GetDate() string {
	return p.date
}

func (p *mandiriParser) cleanUp() {
	p.accountNumber = ""
	p.accountName = ""
	p.description = ""
	p.amount = ""
	p.date = ""
}
