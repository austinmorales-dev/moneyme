package logic

import (
	"log"

	"github.com/piquette/finance-go/quote"
)

func StockPrice(symbol string) float64 {
	if ValidStock(symbol) {
		q, err := quote.Get(symbol)
		if err != nil {
			log.Fatal(err)
		}
		return q.RegularMarketPrice
	} else {
		log.Fatalf("symbol %v returned an error", symbol)
	}
	return 0.0
}

func ValidStock(symbol string) bool {
	q, err := quote.Get(symbol)
	if q == nil {
		return false
	}
	if err != nil {
		return false
	}
	return true
}
