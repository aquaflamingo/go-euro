# Go Euro

Go library to parse the European Central Bank statistics for daily foreign currency rates

## Usage
* Get the current EUR exchange rate
* Get the past 90 day EUR historic exchange rate
* Get the historic EUR exchange rate

```
import (
	 rates "github.com/robertsimoes/go-euro/rates"
)

func main() {
	 exchangeRate, err := rates.AllCurrentRates()
	 
	 usd := exchangeRate.Rates["USD"]
	 fmt.Printf("today 1 EUR = %s USD", usd)
}
```

# Author 
Robert Simoes

# Licence
This repository is licenced under MIT Open Source Licence
