package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type EODHDTickerItem struct {
	Date          string  `json:"date"`
	Open          float64 `json:"open"`
	High          float64 `json:"high"`
	Low           float64 `json:"low"`
	Close         float64 `json:"close"`
	AdjustedClose float64 `json:"adjusted_close"`
	Volume        int64   `json:"volume"`
}

const eodhdHost = "https://eodhd.com/api"

func EODHDGetTicker(symbol string, period string, interval string) ([]EODHDTickerItem, error) {
	apiKey := os.Getenv("EODHD_API_KEY") //TODO: inject from the main app
	if apiKey == "" {
		return nil, fmt.Errorf("EODHD_API_KEY is not set in environment variables")
	}

	url := fmt.Sprintf("%s/eod/%s?api_token=%s&period=%s&interval=%s&fmt=json", eodhdHost, symbol, apiKey, period, interval)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get ticker data: %s", resp.Status)
	}
	var ticker []EODHDTickerItem
	err = json.NewDecoder(resp.Body).Decode(&ticker)
	if err != nil {
		return ticker, err
	}

	return ticker, nil
}
