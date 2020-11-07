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

	result, err := ratesParser.Today()

	if err != nil {
		t.Fatal(err)
	}

	timeRates := result.BaseCube.TimeCubes

	assertIntMatch(len(timeRates), 1, t)

	rates := timeRates[0].RateCubes

	assertIntMatch(len(rates), 32, t)

	usdRate := rates[0]
	assertMatch(usdRate.IsoCurrency, "USD", t)
	assertMatch(usdRate.Rate, "1.1870", t)
}

func assertIntMatch(got int, expect int, t *testing.T) {
	if expect != got {
		t.Fatalf("expected %d got %d", expect, got)
	}
}

func assertMatch(got string, expect string, t *testing.T) {
	if expect != got {
		t.Fatalf("expected %s got %s", expect, got)
	}
}
