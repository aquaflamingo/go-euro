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

type MockECB struct {
	filePath string
}

func (m *MockECB) GetRatesXml(f Feed) ([]byte, error) {
	return xmlFixture(m.filePath), nil
}

func Test_Today(t *testing.T) {
	teardown := setup()
	defer teardown()

	mockSource := MockECB{filePath: "2020-11-06.xml"}

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

func Test_NinetyDay(t *testing.T) {
	teardown := setup()
	defer teardown()

	mockSource := MockECB{filePath: "hist-90d.xml"}

	ratesParser := NewRatesParser(&mockSource)

	result, err := ratesParser.Today()

	if err != nil {
		t.Fatal(err)
	}

	timeRates := result.BaseCube.TimeCubes

	// Only 65 rates
	assertIntMatch(len(timeRates), 65, t)

	rates := timeRates[0].RateCubes

	assertIntMatch(len(rates), 32, t)

	usdRate := rates[0]
	assertMatch(usdRate.IsoCurrency, "USD", t)
	assertMatch(usdRate.Rate, "1.187", t)
}

func Test_Historical(t *testing.T) {
	teardown := setup()
	defer teardown()

	mockSource := MockECB{filePath: "historical.xml"}

	ratesParser := NewRatesParser(&mockSource)

	result, err := ratesParser.Today()

	if err != nil {
		t.Fatal(err)
	}

	timeRates := result.BaseCube.TimeCubes

	if len(timeRates) <= 90 {
		t.Fatalf("expected > %d got %d", 90, len(timeRates))
	}

	rates := timeRates[0].RateCubes

	assertIntMatch(len(rates), 32, t)

	usdRate := rates[0]
	assertMatch(usdRate.IsoCurrency, "USD", t)
	assertMatch(usdRate.Rate, "1.187", t)
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
