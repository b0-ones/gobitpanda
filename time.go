package gobitpanda

import "fmt"

// GetTime gets the time
func (c *Client) GetTime() (*Time, error) {
	time := &Time{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.APIBase, "/v1/time"), nil)
	if err != nil {
		return time, err
	}

	if err = c.Send(req, time); err != nil {
		return time, err
	}

	return time, nil
}
