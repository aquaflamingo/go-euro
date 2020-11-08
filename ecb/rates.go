package ecb

import "fmt"

type EuroExchangeRate struct {
	Date  string
	Rates map[string]string
}

func AllCurrentRates() (*EuroExchangeRate, error) {
	rp := NewRatesParser(&EuropeanCentralBankData{})

	rates, err := rp.GetRates(Daily)

	if err != nil {
		return &EuroExchangeRate{}, err
	}

	today := rates.BaseCube.TimeCubes[0]

	exchangeRates := buildEuroExchangeRateMap(&today)

	if err != nil {
		return &EuroExchangeRate{}, err
	}

	return exchangeRates, nil
}

func GetCurrentRate(iso string) (string, error) {
	rates, err := AllCurrentRates()

	if err != nil {
		return "", err
	}

	v, ok := rates.Rates[iso]

	if !ok {
		return "", fmt.Errorf("No rate available for %s. Perhaps it is not a valid ISO Currency?", iso)
	}
	return v, nil
}

func buildEuroExchangeRateMap(tc *TimeCube) *EuroExchangeRate {
	rate := &EuroExchangeRate{
		Date:  tc.Date,
		Rates: make(map[string]string),
	}

	for _, r := range tc.RateCubes {
		rate.Rates[r.IsoCurrency] = r.Rate
	}

	return rate
}
