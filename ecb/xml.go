package ecb

import "encoding/xml"

type Sender struct {
	XMLName xml.Name `xml:"Sender"`
	Name    string   `xml:"name"`
}

type RateCube struct {
	XMLName     xml.Name `xml:"Cube"`
	IsoCurrency string   `xml:"currency,attr"`
	// Rate in EUR
	Rate string `xml:"rate,attr"`
}

type TimeCube struct {
	XMLName   xml.Name   `xml:"Cube"`
	Date      string     `xml:"time,attr"`
	RateCubes []RateCube `xml:"Cube"`
}

type BaseCube struct {
	XMLName   xml.Name   `xml:"Cube"`
	TimeCubes []TimeCube `xml:"Cube"`
}

type GenericStatisticalRatesMessage struct {
	XMLName  xml.Name `xml:"Envelope"`
	Subject  string   `xml:"subject"`
	Sender   Sender   `xml:"Sender"`
	BaseCube BaseCube `xml:"Cube"`
}
