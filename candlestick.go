package gobitpanda

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

// GetCandlesticks gets instrument's candlesticks for a closed time period.
func (c *Client) GetCandlesticks(instrumentCode string, unit string, period int, from time.Time, to time.Time) (*[]Candlestick, error) {
	cdlstcks := &[]Candlestick{}

	if instrumentCode == "" || unit == "" || period == 0 || from.IsZero() || to.IsZero() {
		return nil, errors.New("instrumentCode, unit, period, from or to can not be empty")
	}

	switch unit {
	case UnitHours:
		if !(period == 1 || period == 4) {
			return nil, errors.New("Unit HOURS only supports period 1 or 4")
		}
	case UnitDays:
		if period != 1 {
			return nil, errors.New("Unit DAYS only supports period 1")
		}
	case UnitMinutes:
		if !(period == 1 || period == 5 || period == 15 || period == 30) {
			return nil, errors.New("Unit MINUTES only supports period 1, 5, 15 or 30")
		}
	case UnitMonths:
		if period != 1 {
			return nil, errors.New("Unit MONTHS only supports period 1")
		}
	case UnitWeeks:
		if period != 1 {
			return nil, errors.New("Unit WEEKS only supports period 1")
		}
	default:
		return nil, errors.New("Unsupported unit")
	}

	f := from.UTC().Format(time.RFC3339)
	t := to.UTC().Format(time.RFC3339)

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s%s%s%s%s%s", c.APIBase, "/v1/candlesticks/", instrumentCode, "?from="+f, "&to="+t, "&unit="+unit, "&period="+strconv.Itoa(period)), nil)
	if err != nil {
		return cdlstcks, err
	}

	if err = c.Send(req, cdlstcks); err != nil {
		return cdlstcks, err
	}

	return cdlstcks, nil
}
