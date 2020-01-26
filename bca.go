package moco

import (
	"regexp"
	"strings"
)

var BCABlacklist = []string{"TRSF", "E-BANKING", "CR", "DB", "Wallet",
	"User", "BIAYA", "SME", "MFTS", "Dana", "YAY", "KITA", "BISA", "Recurring",
	"Auto", "Debet"}

const (
	BcaMinimumRecordLength    = 2
	BcaIndexNumberDescription = 1
	BcaIndexNumberAmount      = 3
	BcaAmountCreditString     = "CR"
)

type bcaParser struct {
	record      []string
	accountName string
	description string
	amount      string
	date        string
	isCredit    bool
}

func (p *bcaParser) validRecordLength() bool {
	return len(p.record) >= BcaMinimumRecordLength
}

func (p *bcaParser) parseRecord() error {
	a := p.record[BcaIndexNumberAmount]
	if !p.isCreditLedger(a) {
		return nil
	}
	d := p.record[BcaIndexNumberDescription]
	p.description = d
	p.date = p.parseTransferDate(d)
	p.accountName = p.parseAccountName(d)
	p.amount = a

	return nil
}

func (p *bcaParser) isCreditLedger(s string) bool {
	return strings.Contains(s, BcaAmountCreditString)
}

func (p *bcaParser) parseTransferDate(s string) string {
	ns := WhitespaceSplit(s)
	ns = BlacklistTrim(ns, BCABlacklist)
	for _, v := range ns {
		if p.isTransferCode(v) {
			nsSplitted := strings.Split(s, "/")
			if len(nsSplitted) == 1 {
				return ""
			}

			return nsSplitted[0]
		}
	}

	return ""
}

func (p *bcaParser) isTransferCode(s string) bool {
	match, err := regexp.MatchString("[0-9]+\\/[A-Z]+\\/[A-Z0-9]*$", s)
	if err != nil {
		return false
	}

	return match
}

func (p *bcaParser) isAllNumber(s string) bool {
	match, err := regexp.MatchString("^[0-9]*$", s)
	if err != nil {
		return false
	}

	return match
}

func (p *bcaParser) isAmount(s string) bool {
	match, err := regexp.MatchString("^[0-9]*.00", s)
	if err != nil {
		return false
	}

	return match
}

func (p *bcaParser) parseAccountName(s string) string {
	var trimmed []string
	ns := WhitespaceSplit(s)
	ns = BlacklistTrim(ns, BCABlacklist)
	for _, v := range ns {
		if p.isTransferCode(v) {
			date := p.parseTransferDate(v)
			if date != "" {
				p.date = date
			}
			continue
		}

		if p.isAllNumber(v) || p.isAmount(v) {
			continue
		}

		trimmed = append(trimmed, v)
	}

	return strings.Join(trimmed, " ")
}

func NewBcaParser() MutationParser {
	return &bcaParser{}
}

func (p *bcaParser) GetAccountName() string {
	return p.accountName
}

func (p *bcaParser) GetAccountNumber() string {
	return ""
}

func (p *bcaParser) LoadRecord(record []string) error {
	p.cleanUp()
	p.record = record
	if !p.validRecordLength() {
		return nil
	}

	return p.parseRecord()
}

func (p *bcaParser) GetDescription() string {
	return p.description
}

func (p *bcaParser) GetAmount() string {
	a, err := NumericTrim(p.amount)

	if err != nil {
		return ""
	}

	return a
}

func (p *bcaParser) GetDate() string {
	return p.date
}

func (p *bcaParser) cleanUp() {
	p.accountName = ""
	p.description = ""
	p.amount = ""
	p.date = ""
}
