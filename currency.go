package gobitpanda

import "fmt"

// GetCurrencies gets a list of all available currencies.
func (c *Client) GetCurrencies() (*[]Currency, error) {
	curr := &[]Currency{}

	req, err := c.NewRequest("GET", fmt.Sprintf("%s%s", c.APIBase, "/v1/currencies"), nil)
	if err != nil {
		return curr, err
	}

	if err = c.Send(req, curr); err != nil {
		return curr, err
	}

	return curr, nil
}
