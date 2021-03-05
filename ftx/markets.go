package ftx

import (
	"log"
	"strconv"

	"github.com/convexityresearch/ftx/ftx/api"
)

// HistoricalPrices is an alias.
type HistoricalPrices api.HistoricalPrices

// Trades is an alias.
type Trades api.Trades

// GetHistoricalPrices returns a list of historical prices for a market.
func (client *Client) GetHistoricalPrices(market string, resolution int64,
	limit int64, startTime int64, endTime int64) (HistoricalPrices, error) {
	var historicalPrices HistoricalPrices
	resp, err := client._get(
		"markets/"+market+
			"/candles?resolution="+strconv.FormatInt(resolution, 10)+
			"&limit="+strconv.FormatInt(limit, 10)+
			"&start_time="+strconv.FormatInt(startTime, 10)+
			"&end_time="+strconv.FormatInt(endTime, 10),
		[]byte(""))
	if err != nil {
		log.Printf("Error GetHistoricalPrices", err)
		return historicalPrices, err
	}
	err = _processResponse(resp, &historicalPrices)
	return historicalPrices, err
}

// GetTrades returns a list of trades for a market.
func (client *Client) GetTrades(market string, limit int64, startTime int64, endTime int64) (Trades, error) {
	var trades Trades
	resp, err := client._get(
		"markets/"+market+"/trades?"+
			"&limit="+strconv.FormatInt(limit, 10)+
			"&start_time="+strconv.FormatInt(startTime, 10)+
			"&end_time="+strconv.FormatInt(endTime, 10),
		[]byte(""))
	if err != nil {
		log.Printf("Error GetTrades", err)
		return trades, err
	}
	err = _processResponse(resp, &trades)
	return trades, err
}
