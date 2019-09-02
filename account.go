package gobitpanda

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// GetAccountBalances get the balance details for an account.
func (c *Client) GetAccountBalances() (*Account, error) {
	acc := &Account{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.APIBase, "/v1/account/balances"), nil)
	if err != nil {
		return acc, err
	}

	if err = c.SendWithAuth(req, acc); err != nil {
		return acc, err
	}

	return acc, nil
}

// GetAccountFees gets the fee details for an account.
func (c *Client) GetAccountFees() (*FeeGroup, error) {
	fees := &FeeGroup{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.APIBase, "/v1/account/fees"), nil)
	if err != nil {
		return fees, err
	}

	if err = c.SendWithAuth(req, fees); err != nil {
		return fees, err
	}

	return fees, nil
}

// GetAccountOrders gets a paginated report on currently open orders, sorted by timestamp (newest first).
// Use query parameters and filters to specify if historical orders should be reported as well.
// If no query filters are defined it returns all orders which are currently active.
// If you want to query specific time frame parameters, FROM and TO are mandatory, otherwise it will start from the newest orders.
// The maximum time-frame you can query at one time is 100 days.
func (c *Client) GetAccountOrders(
	from time.Time,
	to time.Time,
	instrumentCode string,
	withCancelledAndRejected bool,
	withJustFilledInactive bool,
	maxPageSize string,
	cursor string,
) (*OrderHistory, error) {
	orders := &OrderHistory{}
	var params []string
	paramsString := ""

	if !from.IsZero() {
		params = append(params, "from="+from.UTC().Format(time.RFC3339))
	}

	if !to.IsZero() {
		params = append(params, "to="+to.UTC().Format(time.RFC3339))
	}

	if instrumentCode != "" {
		params = append(params, "instrument_code="+instrumentCode)
	}

	if maxPageSize != "" {
		params = append(params, "max_page_size="+maxPageSize)
	}

	if cursor != "" {
		params = append(params, "cursor="+cursor)
	}

	params = append(params, "with_cancelled_and_rejected="+strconv.FormatBool(withCancelledAndRejected))
	params = append(params, "with_just_filled_inactive="+strconv.FormatBool(withJustFilledInactive))

	for i, p := range params {
		if i == 0 {
			paramsString += "?" + p
		} else {
			paramsString += "&" + p
		}
	}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s", c.APIBase, "/v1/account/orders", paramsString), nil)
	if err != nil {
		return orders, err
	}

	if err = c.SendWithAuth(req, orders); err != nil {
		return orders, err
	}

	return orders, nil
}

// GetAccountOrderByID gets information for an order by it's ID
func (c *Client) GetAccountOrderByID(ID string) (*OrderHistoryEntry, error) {
	if ID == "" {
		return nil, errors.New("Order ID can not be empty")
	}

	order := &OrderHistoryEntry{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s", c.APIBase, "/v1/account/orders/", ID), nil)
	if err != nil {
		return order, err
	}

	if err = c.SendWithAuth(req, order); err != nil {
		return order, err
	}

	return order, nil
}

// NewOrder creates a new order
func (c *Client) NewOrder(o *CreateOrder) (*Order, error) {
	if o == nil {
		return nil, errors.New("Invalid input")
	}

	if o.InstrumentCode == "" || o.Side == "" || o.Type == "" || o.Amount == "" {
		return nil, errors.New("InstrumentCode, Side, Type and Ammount can not be empty")
	}

	if o.Type == OrderTypeLimit && o.Price == "" {
		return nil, errors.New("Price can not be empty")
	}

	if o.Type == OrderTypeStop && (o.Price == "" || o.TriggerPrice == "") {
		return nil, errors.New("Price and TriggerPrice can not be empty")
	}

	order := &Order{}

	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.APIBase, "/v1/account/orders"), o)
	if err != nil {
		return order, err
	}

	if err = c.SendWithAuth(req, order); err != nil {
		return order, err
	}

	return order, nil
}

// CloseOrders closes all orders. If an instrument code is given, only orders in this market will be closed.
// Returns an array with closed order IDs
func (c *Client) CloseOrders(m ...string) ([]string, error) {
	if len(m) > 1 {
		return nil, errors.New("Too manny arguments")
	}

	var orderIDs []string
	marketLimit := "?instrument_code="

	if len(m) == 1 {
		marketLimit += m[0]
	} else {
		marketLimit = ""
	}

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s%s", c.APIBase, "/v1/account/orders", marketLimit), nil)
	if err != nil {
		return orderIDs, err
	}

	if err = c.SendWithAuth(req, &orderIDs); err != nil {
		return orderIDs, err
	}

	return orderIDs, nil
}

// CloseOrderByID closes an order by it's ID
func (c *Client) CloseOrderByID(ID string) error {
	if ID == "" {
		return errors.New("Order ID can not be empty")
	}

	req, err := c.NewRequest("DELETE", fmt.Sprintf("%s%s%s", c.APIBase, "/v1/account/orders/", ID), nil)
	if err != nil {
		return err
	}

	if err = c.SendWithAuth(req, nil); err != nil {
		return err
	}

	return nil
}

// GetAccountTrades gets a paginated report on past trades, sorted by timestamp (newest first).
// If no query parameters are defined, it returns the last 100 trades.
func (c *Client) GetAccountTrades(
	from time.Time,
	to time.Time,
	instrumentCode string,
	maxPageSize string,
	cursor string,
) (*TradeHistory, error) {
	trades := &TradeHistory{}
	var params []string
	paramsString := ""

	if !from.IsZero() {
		params = append(params, "from="+from.UTC().Format(time.RFC3339))
	}

	if !to.IsZero() {
		params = append(params, "to="+to.UTC().Format(time.RFC3339))
	}

	if instrumentCode != "" {
		params = append(params, "instrument_code="+instrumentCode)
	}

	if maxPageSize != "" {
		params = append(params, "max_page_size="+maxPageSize)
	}

	if cursor != "" {
		params = append(params, "cursor="+cursor)
	}

	for i, p := range params {
		if i == 0 {
			paramsString += "?" + p
		} else {
			paramsString += "&" + p
		}
	}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s", c.APIBase, "/v1/account/trades", paramsString), nil)
	if err != nil {
		return trades, err
	}

	if err = c.SendWithAuth(req, trades); err != nil {
		return trades, err
	}

	return trades, nil
}

// GetAccountTradeByID gets information for an trade by it' ID
func (c *Client) GetAccountTradeByID(ID string) (*TradeHistoryEntry, error) {
	if ID == "" {
		return nil, errors.New("Trade ID can not be empty")
	}

	trade := &TradeHistoryEntry{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s", c.APIBase, "/v1/account/trades/", ID), nil)
	if err != nil {
		return trade, err
	}

	if err = c.SendWithAuth(req, trade); err != nil {
		return trade, err
	}

	return trade, nil
}

// GetAccountTradesByOrderID gets trade information for a specific order by it's order ID
func (c *Client) GetAccountTradesByOrderID(ID string) (*TradeHistory, error) {
	if ID == "" {
		return nil, errors.New("Order ID can not be empty")
	}

	trade := &TradeHistory{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s%s", c.APIBase, "/v1/account/orders/", ID, "/trades"), nil)
	if err != nil {
		return trade, err
	}

	if err = c.SendWithAuth(req, trade); err != nil {
		return trade, err
	}

	return trade, nil
}

// GetAccountTradingVolume gets the running trading volume for this account.
// It is calculated over a 30 day running window and updated once every 24hrs.
func (c *Client) GetAccountTradingVolume() (*TradingVolume, error) {
	tradingVolume := &TradingVolume{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.APIBase, "/v1/account/trading-volume"), nil)
	if err != nil {
		return tradingVolume, err
	}

	if err = c.SendWithAuth(req, tradingVolume); err != nil {
		return tradingVolume, err
	}

	return tradingVolume, nil
}
