package moco

import (
	"regexp"
	"strings"
)

const (
	BniMinimumRecordLength       = 5
	BniIndexNumberDescription    = 4
	BniIndexNumberValidationDate = 0
	BniIndexNumberAmount         = 6
)

var BniBlacklist = []string{"DARI", "TRANSFER", "Sdr", "Sdri", "|", "PEMINDAHAN"}

type bniParser struct {
	record        []string
	accountName   string
	accountNumber string
	description   string
	amount        string
	date          string
}

func (p *bniParser) validRecordLength() bool {
	return len(p.record) >= BniMinimumRecordLength
}

func (p *bniParser) parseRecord() error {
	d := p.record[BniIndexNumberDescription]
	p.description = d
	p.accountName = p.parseAccountName(d)
	// p.accountNumber = p.parseAccountNumber(d) // new bni csv doesn't include bank account number
	p.amount = p.record[BniIndexNumberAmount]
	date := p.record[BniIndexNumberValidationDate]
	p.date = p.parseDate(date)

	return nil
}

func (p *bniParser) parseAccountName(s string) string {
	ns := WhitespaceSplit(s)
	ns = BlacklistTrim(ns, BniBlacklist)

	an := make([]string, 0)

	for _, v := range ns {
		v = strings.ReplaceAll(v, "|", "")
		if !p.isAllNumber(v) {
			an = append(an, v)
		}
	}

	if len(an) > 1 {
		return strings.Join(an, " ")
	}

	return ""
}

func (p *bniParser) parseAccountNumber(s string) string {
	ns := WhitespaceSplit(s)
	ns = BlacklistTrim(ns, BniBlacklist)
	for _, v := range ns {
		if p.isAllNumber(v) {
			return v
		}
	}

	return ""
}

func (p *bniParser) isAllNumber(s string) bool {
	match, err := regexp.MatchString("^[0-9]*$", s)
	if err != nil {
		return false
	}

	return match
}

func NewBniParser() MutationParser {
	return &bniParser{}
}

func (p *bniParser) GetAccountName() string {
	return p.accountName
}

func (p *bniParser) GetAccountNumber() string {
	return p.accountNumber
}

func (p *bniParser) LoadRecord(record []string) error {
	p.cleanUp()
	p.record = record
	if !p.validRecordLength() {
		return nil
	}

	return p.parseRecord()
}

func (p *bniParser) GetDescription() string {
	return p.description
}

func (p *bniParser) GetAmount() string {
	a, err := NumericTrim(p.amount)

	if err != nil {
		return ""
	}

	return a
}

func (p *bniParser) GetDate() string {
	return p.date
}

func (p *bniParser) parseDate(s string) string {
	ns := WhitespaceSplit(s)
	if len(ns) > 1 {
		return ns[0]
	}

	return ""
}

func (p *bniParser) cleanUp() {
	p.accountName = ""
	p.accountNumber = ""
	p.description = ""
	p.amount = ""
	p.date = ""
}
