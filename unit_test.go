package gobitpanda

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	c, _ := NewClient(APIBase, "")

	time, err := c.GetTime()

	if err != nil {
		t.Errorf("GetTime failed")
	}
	fmt.Println("Time ISO: ", time.Iso)
	fmt.Println("UNIX Time ms: ", time.EpochMillis)
}

func TestGetInstruments(t *testing.T) {

	c, _ := NewClient(APIBase, "")

	instr, err := c.GetInstruments()

	if err != nil {
		t.Errorf("GetInstruments failed")
	}

	fmt.Println("Instruments:\n", instr)
	if len(*instr) == 0 {
		t.Errorf("Empty instruments")
	}
}

func TestGetFees(t *testing.T) {

	c, _ := NewClient(APIBase, "")

	fees, err := c.GetFees()

	if err != nil {
		t.Errorf("GetFees failed")
	}

	fmt.Println("Fees:\n", fees)
	if len(*fees) == 0 {
		t.Errorf("Empty fees")
	}
}

func TestGetCurrencies(t *testing.T) {

	c, _ := NewClient(APIBase, "")

	fees, err := c.GetCurrencies()

	if err != nil {
		t.Errorf("GetCurrencies failed")
	}

	fmt.Println("Currencies:\n", fees)
	if len(*fees) == 0 {
		t.Errorf("Empty currencies")
	}
}

func TestGetCandlesticks(t *testing.T) {

	c, _ := NewClient(APIBase, "")

	to := time.Now()
	fmt.Println(to)
	from := to.AddDate(0, 0, -1)
	fmt.Println(from)
	cndlstck, err := c.GetCandlesticks(InstrumentMIOTAEUR, UnitHours, 1, from, to)

	if err != nil {
		t.Errorf("GetCandlesticks failed: %s", err)
	}

	fmt.Println("Candlesticks:\n", cndlstck)
	if len(*cndlstck) == 0 {
		t.Errorf("Empty candlesticks")
	}
}

func TestGetOrderBook(t *testing.T) {
	c, _ := NewClient(APIBase, "")

	orderbook, err := c.GetOrderBook(InstrumentMIOTAEUR, LevelDefault)

	if err != nil {
		t.Errorf("GetOrderBook failed: %s", err)
	}

	fmt.Println("Order Book:\n", orderbook)
}

func TestGetOrderBookLvlOne(t *testing.T) {
	c, _ := NewClient(APIBase, "")

	orderbook, err := c.GetOrderBookLvlOne(InstrumentMIOTAEUR)

	if err != nil {
		t.Errorf("GetOrderBookLvlOne failed: %s", err)
	}

	fmt.Println("Order Book:\n", orderbook)
}
