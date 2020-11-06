package ecb

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

const EuropeanCentralBankRatesUrl = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-%s.xml"

type Feed string

const (
	Daily               Feed = "daily"
	Historical          Feed = "hist"
	NinetyDayHistorical Feed = "hist-90d"
)

type XMLFetcher interface {
	GetRatesXml(f Feed) ([]byte, error)
}

type RatesParser struct {
	fetcher XMLFetcher
}

func NewRatesParser(f XMLFetcher) *RatesParser {
	return &RatesParser{fetcher: f}
}

func (p *RatesParser) Today() (*RatesMessage, error) {
	xmlData, err := p.fetcher.GetRatesXml(Daily)

	if err != nil {
		return &RatesMessage{}, err
	}

	var rates RatesMessage

	xml.Unmarshal(xmlData, &rates)

	return &rates, nil
}

type EuropeanCentralBankXMLFetcher struct{}

func (ecb *EuropeanCentralBankXMLFetcher) GetRatesXml(f Feed) ([]byte, error) {
	response, err := http.Get(fmt.Sprintf(EuropeanCentralBankRatesUrl, f))

	if err != nil {
		return []byte{}, fmt.Errorf("failed to get xml: %v", err)
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("failed to get xml bad status error: %v", response.StatusCode)
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to get xml could not read body: %v", err)
	}

	return data, nil
}
