package gobitpanda

import "fmt"

// GetInstruments gets the public exchange resource encompassing information about tradeable assets.
func (c *Client) GetInstruments() (*[]Instrument, error) {
	instruments := &[]Instrument{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.APIBase, "/v1/instruments"), nil)
	if err != nil {
		return instruments, err
	}

	if err = c.Send(req, instruments); err != nil {
		return instruments, err
	}

	return instruments, nil
}
