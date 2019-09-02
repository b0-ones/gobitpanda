package gobitpanda

import "fmt"

// GetFees gets details of all publicly-visible Fee Groups.
func (c *Client) GetFees() (*[]FeeGroup, error) {
	fees := &[]FeeGroup{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.APIBase, "/v1/fees"), nil)
	if err != nil {
		return fees, err
	}

	if err = c.Send(req, fees); err != nil {
		return fees, err
	}

	return fees, nil
}
