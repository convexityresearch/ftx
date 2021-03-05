package ftx

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/convexityresearch/ftx/ftx/api"
)

// NewOrder is an alias.
type NewOrder api.NewOrder

// NewOrderResponse is an alias.
type NewOrderResponse api.NewOrderResponse

// OpenOrders is an alias.
type OpenOrders api.OpenOrders

// OrderHistory is an alias.
type OrderHistory api.OrderHistory

// NewTriggerOrder is an alias.
type NewTriggerOrder api.NewTriggerOrder

// NewTriggerOrderResponse is an alias.
type NewTriggerOrderResponse api.NewTriggerOrderResponse

// OpenTriggerOrders is an alias.
type OpenTriggerOrders api.OpenTriggerOrders

// TriggerOrderHistory is an alias.
type TriggerOrderHistory api.TriggerOrderHistory

// Triggers is an alias.
type Triggers api.Triggers

// GetOpenOrders will list all current posted orders that are not cancelled.
func (client *Client) GetOpenOrders(market string) (OpenOrders, error) {
	var openOrders OpenOrders
	resp, err := client._get("orders?market="+market, []byte(""))
	if err != nil {
		log.Printf("Error GetOpenOrders", err)
		return openOrders, err
	}
	err = _processResponse(resp, &openOrders)
	return openOrders, err
}

// GetOrderHistory will fetch the order history for a given period of time.
func (client *Client) GetOrderHistory(market string, startTime float64, endTime float64, limit int64) (OrderHistory, error) {
	var orderHistory OrderHistory
	requestBody, err := json.Marshal(map[string]interface{}{
		"market":     market,
		"start_time": startTime,
		"end_time":   endTime,
		"limit":      limit,
	})
	if err != nil {
		log.Printf("Error GetOrderHistory", err)
		return orderHistory, err
	}
	resp, err := client._get("orders/history?market="+market, requestBody)
	if err != nil {
		log.Printf("Error GetOrderHistory", err)
		return orderHistory, err
	}
	err = _processResponse(resp, &orderHistory)
	return orderHistory, err
}

// GetOpenTriggerOrders will fetch all open trigger orders.
func (client *Client) GetOpenTriggerOrders(market string, _type string) (OpenTriggerOrders, error) {
	var openTriggerOrders OpenTriggerOrders
	requestBody, err := json.Marshal(map[string]string{"market": market, "type": _type})
	if err != nil {
		log.Printf("Error GetOpenTriggerOrders", err)
		return openTriggerOrders, err
	}
	resp, err := client._get("conditional_orders?market="+market, requestBody)
	if err != nil {
		log.Printf("Error GetOpenTriggerOrders", err)
		return openTriggerOrders, err
	}
	err = _processResponse(resp, &openTriggerOrders)
	return openTriggerOrders, err
}

// GetTriggers will fetch the active triggers for a given order.
func (client *Client) GetTriggers(orderID string) (Triggers, error) {
	var trigger Triggers
	resp, err := client._get("conditional_orders/"+orderID+"/triggers", []byte(""))
	if err != nil {
		log.Printf("Error GetTriggers", err)
		return trigger, err
	}
	err = _processResponse(resp, &trigger)
	return trigger, err
}

// GetTriggerOrdersHistory will use the historical data API to fetch all trigger orders ever posted.
func (client *Client) GetTriggerOrdersHistory(market string, startTime float64, endTime float64, limit int64) (TriggerOrderHistory, error) {
	var triggerOrderHistory TriggerOrderHistory
	requestBody, err := json.Marshal(map[string]interface{}{
		"market":     market,
		"start_time": startTime,
		"end_time":   endTime,
	})
	if err != nil {
		log.Printf("Error GetTriggerOrdersHistory", err)
		return triggerOrderHistory, err
	}
	resp, err := client._get("conditional_orders/history?market="+market, requestBody)
	if err != nil {
		log.Printf("Error GetTriggerOrdersHistory", err)
		return triggerOrderHistory, err
	}
	err = _processResponse(resp, &triggerOrderHistory)
	return triggerOrderHistory, err
}

// PlaceOrder will post an order to market
func (client *Client) PlaceOrder(market string, side string, price float64,
	_type string, size float64, reduceOnly bool, ioc bool, postOnly bool) (NewOrderResponse, error) {
	var newOrderResponse NewOrderResponse
	requestBody, err := json.Marshal(NewOrder{
		Market:     market,
		Side:       side,
		Price:      price,
		Type:       _type,
		Size:       size,
		ReduceOnly: reduceOnly,
		Ioc:        ioc,
		PostOnly:   postOnly})
	if err != nil {
		log.Printf("Error PlaceOrder", err)
		return newOrderResponse, err
	}
	resp, err := client._post("orders", requestBody)
	if err != nil {
		log.Printf("Error PlaceOrder", err)
		return newOrderResponse, err
	}
	err = _processResponse(resp, &newOrderResponse)
	return newOrderResponse, err
}

// PlaceTriggerOrder will post a trigger order for an open order.
func (client *Client) PlaceTriggerOrder(market string, side string, size float64,
	_type string, reduceOnly bool, retryUntilFilled bool, triggerPrice float64,
	orderPrice float64, trailValue float64) (NewTriggerOrderResponse, error) {

	var newTriggerOrderResponse NewTriggerOrderResponse
	var newTriggerOrder NewTriggerOrder

	switch _type {
	case "stop":
		if orderPrice != 0 {
			newTriggerOrder = NewTriggerOrder{
				Market:       market,
				Side:         side,
				TriggerPrice: triggerPrice,
				Type:         _type,
				Size:         size,
				ReduceOnly:   reduceOnly,
				OrderPrice:   orderPrice,
			}
		} else {
			newTriggerOrder = NewTriggerOrder{
				Market:       market,
				Side:         side,
				TriggerPrice: triggerPrice,
				Type:         _type,
				Size:         size,
				ReduceOnly:   reduceOnly,
			}
		}
	case "trailingStop":
		newTriggerOrder = NewTriggerOrder{
			Market:     market,
			Side:       side,
			Type:       _type,
			Size:       size,
			ReduceOnly: reduceOnly,
			TrailValue: trailValue,
		}
	case "takeProfit":
		newTriggerOrder = NewTriggerOrder{
			Market:       market,
			Side:         side,
			TriggerPrice: triggerPrice,
			Type:         _type,
			Size:         size,
			ReduceOnly:   reduceOnly,
			OrderPrice:   orderPrice,
		}
	default:
		log.Printf("Trigger type is not valid")
	}
	requestBody, err := json.Marshal(newTriggerOrder)
	if err != nil {
		log.Printf("Error PlaceTriggerOrder", err)
		return newTriggerOrderResponse, err
	}
	resp, err := client._post("conditional_orders", requestBody)
	if err != nil {
		log.Printf("Error PlaceTriggerOrder", err)
		return newTriggerOrderResponse, err
	}
	err = _processResponse(resp, &newTriggerOrderResponse)
	return newTriggerOrderResponse, err
}

// CancelOrder will cancel an order by its ID.
func (client *Client) CancelOrder(orderID int64) (Response, error) {
	var deleteResponse Response
	id := strconv.FormatInt(orderID, 10)
	resp, err := client._delete("orders/"+id, []byte(""))
	if err != nil {
		log.Printf("Error CancelOrder", err)
		return deleteResponse, err
	}
	err = _processResponse(resp, &deleteResponse)
	return deleteResponse, err
}

// CancelTriggerOrder will cancel an open trigger order.
func (client *Client) CancelTriggerOrder(orderID int64) (Response, error) {
	var deleteResponse Response
	id := strconv.FormatInt(orderID, 10)
	resp, err := client._delete("conditional_orders/"+id, []byte(""))
	if err != nil {
		log.Printf("Error CancelTriggerOrder", err)
		return deleteResponse, err
	}
	err = _processResponse(resp, &deleteResponse)
	return deleteResponse, err
}

// CancelAllOrders will cancel all open orders.
func (client *Client) CancelAllOrders() (Response, error) {
	var deleteResponse Response
	resp, err := client._delete("orders", []byte(""))
	if err != nil {
		log.Printf("Error CancelAllOrders", err)
		return deleteResponse, err
	}
	err = _processResponse(resp, &deleteResponse)
	return deleteResponse, err
}
