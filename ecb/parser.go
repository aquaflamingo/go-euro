package ecb

const DailyRatesSourceUrl = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

type Parser struct {
	SourceUrl string
}

func (p *Parser) GetRates() (*EcbRatesMessage, error) {
	// todo
	return &EcbRatesMessage{}, nil
}
