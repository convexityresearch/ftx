package api

// Response is a type representing the FTX API Response format.
type Response struct {
	Success bool        `json:"success"`
	Result  interface{} `json:"result"`
}
