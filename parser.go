package moco

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

type MutationBank struct {
	AccountName   string
	AccountNumber string
	Amount        int
	Date          string
	Description   string
	Hash          string
}

type parser struct {
	mutationParser MutationParser
	record         []string
}

type Parser interface {
	LoadRecord(record []string) error
	GetMutation() *MutationBank
}

type MutationParser interface {
	LoadRecord(record []string) error
	GetAccountName() string
	GetAccountNumber() string
	GetDescription() string
	GetAmount() string
	GetDate() string
}

func NewParser(bankName string) Parser {
	if bankName == BankMandiri {
		return &parser{
			mutationParser: NewMandiriParser(),
		}
	}
	if bankName == BankBCA {
		return &parser{
			mutationParser: NewBcaParser(),
		}
	}
	if bankName == BankBRI {
		return &parser{
			mutationParser: NewBriParser(),
		}
	}

	return nil
}

func (p *parser) LoadRecord(record []string) error {
	p.record = record
	return p.mutationParser.LoadRecord(record)
}

func (p *parser) isValidMutation(mutation *MutationBank) bool {
	if mutation.AccountName == "" && mutation.AccountNumber == "" {
		return false
	}

	if mutation.Amount <= 0 {
		return false
	}

	return true
}

func (p *parser) getHash(ns []string) string {
	data := []byte(strings.Join(ns, " "))
	return fmt.Sprintf("%x", md5.Sum(data))
}

func (p *parser) GetMutation() *MutationBank {
	mutation := new(MutationBank)

	mutation.AccountName = p.mutationParser.GetAccountName()
	mutation.AccountNumber = p.mutationParser.GetAccountNumber()
	amount, err := strconv.Atoi(p.mutationParser.GetAmount())
	if err != nil {
		amount = 0
	}
	mutation.Amount = amount
	mutation.Description = p.mutationParser.GetDescription()
	mutation.Date = p.mutationParser.GetDate()
	mutation.Hash = p.getHash(p.record)

	if p.isValidMutation(mutation) {
		return mutation
	}

	return nil
}
