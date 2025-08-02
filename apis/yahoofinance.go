package apis

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type YFinanceClient struct{}

type Ticker struct {
	Chart Chart `json:"chart"`
}

type Chart struct {
	Result []Result    `json:"result"`
	Error  interface{} `json:"error"`
}

type Result struct {
	Meta       Meta       `json:"meta"`
	Timestamp  []int64    `json:"timestamp"`
	Indicators Indicators `json:"indicators"`
}

type Indicators struct {
	Quote []Quote `json:"quote"`
}

type Quote struct {
	High   []float64 `json:"high"`
	Volume []int64   `json:"volume"`
	Open   []float64 `json:"open"`
	Low    []float64 `json:"low"`
	Close  []float64 `json:"close"`
}

type Meta struct {
	Currency             string               `json:"currency"`
	Symbol               string               `json:"symbol"`
	ExchangeName         string               `json:"exchangeName"`
	FullExchangeName     string               `json:"fullExchangeName"`
	InstrumentType       string               `json:"instrumentType"`
	FirstTradeDate       int64                `json:"firstTradeDate"`
	RegularMarketTime    int64                `json:"regularMarketTime"`
	HasPrePostMarketData bool                 `json:"hasPrePostMarketData"`
	Gmtoffset            int64                `json:"gmtoffset"`
	Timezone             string               `json:"timezone"`
	ExchangeTimezoneName string               `json:"exchangeTimezoneName"`
	RegularMarketPrice   float64              `json:"regularMarketPrice"`
	FiftyTwoWeekHigh     float64              `json:"fiftyTwoWeekHigh"`
	FiftyTwoWeekLow      float64              `json:"fiftyTwoWeekLow"`
	RegularMarketDayHigh float64              `json:"regularMarketDayHigh"`
	RegularMarketDayLow  float64              `json:"regularMarketDayLow"`
	RegularMarketVolume  int64                `json:"regularMarketVolume"`
	LongName             string               `json:"longName"`
	ShortName            string               `json:"shortName"`
	ChartPreviousClose   float64              `json:"chartPreviousClose"`
	PreviousClose        float64              `json:"previousClose"`
	Scale                int64                `json:"scale"`
	PriceHint            int64                `json:"priceHint"`
	CurrentTradingPeriod CurrentTradingPeriod `json:"currentTradingPeriod"`
	TradingPeriods       [][]Post             `json:"tradingPeriods"`
	DataGranularity      string               `json:"dataGranularity"`
	Range                string               `json:"range"`
	ValidRanges          []string             `json:"validRanges"`
}

type CurrentTradingPeriod struct {
	Pre     Post `json:"pre"`
	Regular Post `json:"regular"`
	Post    Post `json:"post"`
}

type Post struct {
	Timezone  string `json:"timezone"`
	End       int64  `json:"end"`
	Start     int64  `json:"start"`
	Gmtoffset int64  `json:"gmtoffset"`
}

func YFinanceGetTicker(symbol, period, interval string) (Ticker, error) {
	ticker := Ticker{}
	url := fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?range=%s&interval=%s", symbol, period, interval)

	resp, err := http.Get(url)
	if err != nil {
		return ticker, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ticker, fmt.Errorf("failed to get ticker data: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&ticker)
	if err != nil {
		return ticker, err
	}

	return ticker, nil
}
