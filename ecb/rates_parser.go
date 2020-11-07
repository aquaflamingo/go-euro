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

type DataSource interface {
	GetRatesXml(f Feed) ([]byte, error)
}

type RatesParser struct {
	source DataSource
}

func NewRatesParser(f DataSource) *RatesParser {
	return &RatesParser{source: f}
}

func (p *RatesParser) NinetyDay() (*RatesMessage, error) {
	xmlData, err := p.source.GetRatesXml(NinetyDayHistorical)

	if err != nil {
		return &RatesMessage{}, err
	}

	var rates RatesMessage

	xml.Unmarshal(xmlData, &rates)

	return &rates, nil
}

func (p *RatesParser) Historical() (*RatesMessage, error) {
	xmlData, err := p.source.GetRatesXml(Historical)

	if err != nil {
		return &RatesMessage{}, err
	}

	var rates RatesMessage

	xml.Unmarshal(xmlData, &rates)

	return &rates, nil
}

func (p *RatesParser) Today() (*RatesMessage, error) {
	xmlData, err := p.source.GetRatesXml(Daily)

	if err != nil {
		return &RatesMessage{}, err
	}

	var rates RatesMessage

	xml.Unmarshal(xmlData, &rates)

	return &rates, nil
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
