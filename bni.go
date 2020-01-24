package moco

import (
	"strings"
)

const (
	BniMinimumRecordLength       = 5
	BniIndexNumberDescription    = 4
	BniIndexNumberValidationDate = 0
	BniIndexNumberAmount         = 6
)

var BniBlacklist = []string{"DARI", "TRANSFER", "Sdr", "|"}

type bniParser struct {
	record      []string
	accountName string
	description string
	amount      string
	date        string
}

func (p *bniParser) validRecordLength() bool {
	return len(p.record) >= BniMinimumRecordLength
}

func (p *bniParser) parseRecord() error {
	d := p.record[BniIndexNumberDescription]
	p.description = d
	ns := WhitespaceSplit(d)
	ns = BlacklistTrim(ns, BniBlacklist)
	p.accountName = p.parseAccountName(ns)
	p.amount = p.record[BniIndexNumberAmount]
	date := p.record[BniIndexNumberValidationDate]
	p.date = p.parseDate(date)

	return nil
}

func (p *bniParser) parseAccountName(ns []string) string {
	if len(ns) > 1 {
		return strings.Join(ns, " ")
	}

	return ""
}

func NewBniParser() MutationParser {
	return &bniParser{}
}

func (p *bniParser) GetAccountName() string {
	return p.accountName
}

func (p *bniParser) GetAccountNumber() string {
	//Not implemented
	return ""
}

func (p *bniParser) LoadRecord(record []string) error {
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
