# Go Euro
[![Project Status: Active – The project has reached a stable, usable state and is being actively developed.](https://www.repostatus.org/badges/latest/active.svg)](https://www.repostatus.org/#active)

Go library to parse the European Central Bank statistics for daily foreign currency rates

## Usage
* Get the current EUR exchange rate
* Get the past 90 day EUR historic exchange rate
* Get the historic EUR exchange rate

```

func main() {
	 exchangeRate, err := goeuro.AllCurrentRates()
	 
	 usd := exchangeRate.Rates["USD"]
	 fmt.Printf("today 1 EUR = %s USD", usd)
}
```

# License
This repository is licensed under MIT Open Source Licence
