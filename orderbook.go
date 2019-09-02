package gobitpanda

import (
	"errors"
	"fmt"
	"strconv"
)

// GetOrderBook gets a given instrument's order book.
func (c *Client) GetOrderBook(instrumentCode string, level int) (*OrderBook, error) {
	orderbook := &OrderBook{}

	if instrumentCode == "" {
		return nil, errors.New("instrumentCode can not be empty")
	}

	if level == 0 {
		level = LevelDefault
	} else if level > 3 {
		return nil, errors.New("level can not be > 3")
	} else if level < 0 {
		return nil, errors.New("level can not be < 0")
	} else if level == 1 {
		return nil, errors.New("For leve one please use GetOrderBookLvlOne()")
	}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s%s", c.APIBase, "/v1/order-book/", instrumentCode, "?level="+strconv.Itoa(level)), nil)
	if err != nil {
		return orderbook, err
	}

	if err = c.Send(req, orderbook); err != nil {
		return orderbook, err
	}

	return orderbook, nil
}

// GetOrderBookLvlOne gets a given instrument's order book at level one
func (c *Client) GetOrderBookLvlOne(instrumentCode string) (*OrderBookLvlOne, error) {
	orderbook := &OrderBookLvlOne{}

	if instrumentCode == "" {
		return nil, errors.New("instrumentCode can not be empty")
	}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s%s", c.APIBase, "/v1/order-book/", instrumentCode, "?level=1"), nil)
	if err != nil {
		return orderbook, err
	}

	if err = c.Send(req, orderbook); err != nil {
		return orderbook, err
	}

	return orderbook, nil
}
