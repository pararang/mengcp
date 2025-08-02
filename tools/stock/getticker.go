package stock

import (
	"encoding/json"

	"github.com/pararang/emcp/apis"
	"github.com/pararang/emcp/claude"
)

type GetTickerInput struct {
	Symbol   string `json:"symbol" jsonschema_description:"The stock symbol to get ticker information for."`
	Period   string `json:"period" jsonschema_description:"The period or time range for which to get ticker information. Default is '1d'."`
	Interval string `json:"interval" jsonschema_description:"The interval for the ticker data. Default is '5m'."`
}

var GetTickerDefinition = claude.ToolDefinition{
	Name: "get_ticker",
	Description: `
	Get ticker information for a stock symbol. Use this tool to retrieve real-time or historical stock data based on the provided symbol, period, and interval.
	Use the stock symbol in the 'symbol' field. The 'period' field can be set to "1d","5d","1mo","3mo","6mo","1y","2y","5y","10y","ytd","max". 
	The 'interval' field can be set to "1m","2m","3m","5m","15m","30m","60m","4h","1d","1wk","1mo","1y". If not specified, defaults are used. 
	If the interval is greater than the period, it will be adjusted to be less than the period.
	}
	`,
	InputSchema: claude.GenerateSchema[GetTickerInput](),
	Function:    GetTicker,
}

func GetTicker(input json.RawMessage) (string, error) {
	var getTickerInput GetTickerInput
	err := json.Unmarshal(input, &getTickerInput)
	if err != nil {
		return "", err
	}

	var ticker any
	// ticker, err = apis.YFinanceGetTicker(getTickerInput.Symbol, getTickerInput.Period, getTickerInput.Interval)
	// if err != nil {
	// fallback to EODHD API if YFinance fails
	ticker, err = apis.EODHDGetTicker(getTickerInput.Symbol, getTickerInput.Period, getTickerInput.Interval)
	if err != nil {
		return "", err
	}
	// }

	result := struct {
		Param  GetTickerInput `json:"param"`
		Ticker any            `json:"detail"`
	}{
		Param:  getTickerInput,
		Ticker: ticker,
	}

	resultJSON, err := json.Marshal(result)
	if err != nil {
		return "", err
	}

	return string(resultJSON), nil
}
