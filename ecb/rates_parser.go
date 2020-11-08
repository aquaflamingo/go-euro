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
	Daily      Feed = "daily"
	Historical Feed = "hist"
	NinetyDay  Feed = "hist-90d"
)

type ExchangeRateDataSource interface {
	GetRatesXml(f Feed) ([]byte, error)
}

type RatesParser struct {
	source ExchangeRateDataSource
}

func NewRatesParser(f ExchangeRateDataSource) *RatesParser {
	return &RatesParser{source: f}
}

type EuropeanCentralBankData struct{}

func (ecb *EuropeanCentralBankData) GetRatesXml(f Feed) ([]byte, error) {
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

func (p *RatesParser) GetRates(f Feed) (*GenericStatisticalRatesMessage, error) {
	xmlData, err := p.source.GetRatesXml(f)

	if err != nil {
		return &GenericStatisticalRatesMessage{}, err
	}

	var rates GenericStatisticalRatesMessage

	xml.Unmarshal(xmlData, &rates)

	return &rates, nil
}
