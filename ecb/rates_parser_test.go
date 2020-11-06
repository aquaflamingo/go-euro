package ecb

import (
	"io/ioutil"
	"testing"
)

func xmlFixture(path string) []byte {
	f, err := ioutil.ReadFile("testdata/fixtures/" + path)

	if err != nil {
		panic(err)
	}

	return f
}

func setup() func() {
	return func() {}
}

type MockECB struct{}

func (m *MockECB) GetRatesXml(f Feed) ([]byte, error) {
	return xmlFixture("2020-11-06.xml"), nil
}

func Test_Today(t *testing.T) {
	teardown := setup()
	defer teardown()

	mockSource := MockECB{}

	ratesParser := NewRatesParser(&mockSource)

	// TODO Mock xml
	test
}
