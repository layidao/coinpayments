package coinpayments

import "net/url"

//Request is the representation of a '' api request
type BalancesRequest struct {
	All string
}

func (r *BalancesRequest) command() string {
	return "balances"
}

func (r *BalancesRequest) values() url.Values {
	values := url.Values{}
	if r.All != "" {
		values.Set("all", r.All)
	}

	return values
}

type BalancesResponse map[string]struct {
	Balance  int    `json:"balance"`
	Balancef string `json:"balancef"`
	Status   string `json:"status"`
}

type balancesResult struct {
	errResponse
	Result *BalancesResponse `json:"result"`
}

func (c *Client) Balances(request *BalancesRequest) (*BalancesResponse, error) {
	var resp balancesResult
	if err := c.call(request, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
