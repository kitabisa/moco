package moco

const (
	BriMinimumRecordLength       = 9
	BriIndexNumberDescription    = 8
	BriIndexNumberValidationDate = 6
	BriIndexNumberAmount         = 10
)

var BriBlacklist = []string{"DARI", "KE", "DATE", "TIME", "REMARK", "DEBET", "CREDIT", "TELLER ID"}

type briParser struct {
	record        []string
	accountNumber string
	description   string
	amount        string
	date          string
}

func (p *briParser) validRecordLength() bool {
	return len(p.record) >= BriMinimumRecordLength
}

func (p *briParser) parseRecord() error {
	d := p.record[BriIndexNumberDescription]
	p.description = d
	ns := WhitespaceSplit(d)
	ns = BlacklistTrim(ns, BriBlacklist)
	p.accountNumber = p.parseAccountNumber(ns)
	p.amount = p.record[BriIndexNumberAmount]
	p.date = p.record[BriIndexNumberValidationDate]

	return nil
}

func (p *briParser) parseAccountNumber(ns []string) string {
	if len(ns) > 1 {
		return ns[0]
	}

	return ""
}

func NewBriParser() MutationParser {
	return &briParser{}
}

func (p *briParser) GetAccountName() string {
	//Not implemented
	return ""
}

func (p *briParser) GetAccountNumber() string {
	return p.accountNumber
}

func (p *briParser) LoadRecord(record []string) error {
	p.record = record
	if !p.validRecordLength() {
		return nil
	}
	return p.parseRecord()
}

func (p *briParser) GetDescription() string {
	return p.description
}

func (p *briParser) GetAmount() string {
	a, err := NumericTrim(p.amount)

	if err != nil {
		return ""
	}

	return a
}

func (p *briParser) GetDate() string {
	return p.date
}
