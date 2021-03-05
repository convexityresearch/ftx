package api

import (
	"time"
)

// HistoricalPrices represents a response from the historical prices API.
type HistoricalPrices struct {
	Success bool `json:"success"`
	Result  []struct {
		Close     float64   `json:"close"`
		High      float64   `json:"high"`
		Low       float64   `json:"low"`
		Open      float64   `json:"open"`
		StartTime time.Time `json:"startTime"`
		Volume    float64   `json:"volume"`
	} `json:"result"`
}

// Trades represent a response from the trades API.
type Trades struct {
	Success bool `json:"success"`
	Result  []struct {
		ID          int64     `json:"id"`
		Liquidation bool      `json:"liquidation"`
		Price       float64   `json:"price"`
		Side        string    `json:"side"`
		Size        float64   `json:"size"`
		Time        time.Time `json:"time"`
	} `json:"result"`
}
