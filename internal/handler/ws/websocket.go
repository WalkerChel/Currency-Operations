package ws

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"
)

// {
// 	"Realtime Currency Exchange Rate": {
// 	  "1. From_Currency Code": "BTC",
// 	  "2. From_Currency Name": "Bitcoin",
// 	  "3. To_Currency Code": "USD",
// 	  "4. To_Currency Name": "United States Dollar",
// 	  "5. Exchange Rate": "68508.68000000",
// 	  "6. Last Refreshed": "2024-03-15 06:44:02",
// 	  "7. Time Zone": "UTC",
// 	  "8. Bid Price": "68508.67000000",
// 	  "9. Ask Price": "68508.68000000"
// 	}
//   }

type CryptoResponceJSON struct {
}

func GetInfoTest(key string) {
	currencyURL := fmt.Sprintf("https://www.alphavantage.co/query?function=CURRENCY_EXCHANGE_RATE&from_currency=BTC&to_currency=USD&apikey=%s", key)
	text, err := getJson(currencyURL)
	if err != nil {
		logrus.Fatalf("error parsing crypto: %s", err.Error())
	}

	fmt.Println(string(text))
}

func getJson(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
