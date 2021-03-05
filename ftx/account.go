package ftx

import (
	"log"

	"github.com/convexityresearch/ftx/ftx/api"
)

// Positions is an alias.
type Positions api.Positions

// GetPositions returns a list of positions for the client.
func (client *Client) GetPositions(showAvgPrice bool) (Positions, error) {
	var positions Positions
	resp, err := client._get("positions", []byte(""))
	if err != nil {
		log.Println("Error GetPositions", err)
		return positions, err
	}
	err = _processResponse(resp, &positions)
	return positions, err
}
