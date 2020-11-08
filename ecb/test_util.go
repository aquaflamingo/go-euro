package ecb

import (
	"io/ioutil"
)

func xmlFixture(path string) []byte {
	f, err := ioutil.ReadFile("testdata/fixtures/" + path)

	if err != nil {
		panic(err)
	}

	return f
}

var ratesParser *RatesParser

func TestSetup(filePath string) func() {
	mockSource := MockECB{filePath: filePath}

	ratesParser = NewRatesParser(&mockSource)

	return func() {
		// Reset for next test
		ratesParser = nil
	}
}

type MockECB struct {
	filePath string
}

func (m *MockECB) GetRatesXml(_ Feed) ([]byte, error) {
	return xmlFixture(m.filePath), nil
}
