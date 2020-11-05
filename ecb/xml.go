package ecb

import "encoding/xml"

type Sender struct {
	XMLName xml.Name `xml:"gesmes:Sender"`
	Name    string   `xml:"gesmes:name"`
}

type RateCube struct {
	XMLName  xml.Name `xml:"Cube"`
	Currency string   `xml:"currency,attr"`
	// Rate in EUR
	Rate string `xml:"rate,attr"`
}

type TimeCube struct {
	XMLName   xml.Name   `xml:"Cube"`
	Time      string     `xml:"time,attr"`
	RateCubes []RateCube `xml:"Cube"`
}

type BaseCube struct {
	XMLName   xml.Name   `xml:"Cube"`
	TimeCubes []TimeCube `xml:"Cube"`
}

type EcbRatesMessage struct {
	XMLName  xml.Name `xml:"gesmes:Envelope"`
	Subject  string   `xml:"gesmes:subject"`
	Sender   Sender   `xml:"gesmes:Sender"`
	BaseCube BaseCube `xml:"Cube"`
}
