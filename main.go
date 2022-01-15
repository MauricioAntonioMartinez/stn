package main

import (
	"fmt"
	"log"

	"github.com/tonymackay/finnhub-go"
)

func main() {
	client := finnhub.NewClient("c79req2ad3iel691oe20")

	symbols, err := client.StockSymbols(finnhub.ExchangeUS)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range symbols {
		fmt.Println("Symbol: ", s.Symbol)

	}

}
